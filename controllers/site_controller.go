package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"regexp"
	"time"

	"spurt-cms/models"
	"strconv"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/categories"
	chn "github.com/spurtcms/channels"
	listing "github.com/spurtcms/listing"
	"github.com/spurtcms/membership"
	"github.com/spurtcms/menu"
	"github.com/spurtcms/team"
	csrf "github.com/utrack/gin-csrf"
)

type ChannelTemplate struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	TemplateType string `json:"templatetype"`
}

func ChannelEntriesList(c *gin.Context) {

	channelname := strings.ToLower(c.Param("slug"))

	categoryname := c.Query("filter")

	User, website, _ := GetTenantByHost(c)

	templatedetails, err := MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	if err != nil {
		fmt.Println(err)
	}

	template_name := strings.ToLower(templatedetails.TemplateName)
	template_name = strings.ReplaceAll(template_name, " ", "_")
	// templatedetails, err := MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	newmenulist, merr := MenuConfig.GetmenusByTenantId(website.Id, User.TenantId)

	if merr != nil {

		fmt.Println(merr)
	}

	for i := range newmenulist {
		newmenulist[i].Description = StripHTML(newmenulist[i].Description)
	}

	var (
		limt, offset int
	)

	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	if pageno != 0 {
		offset = (pageno - 1) * Limit
	}

	searchKey := c.Query("search")

	var listingfilter listing.Filter

	if searchKey != "" {
		listingfilter.Keyword = searchKey
	}

	channelDetail, err := ChannelConfigWP.ChannelDetail(chn.Channels{Slug: channelname, TenantId: User.TenantId})
	if channelDetail.Id == 0 || err != nil {
		c.Redirect(301, "/404-TemplatePage")
		return
	}

	_, FinalSelectedCategories, cerr := ChannelConfig.GetChannelsById(channelDetail.Id, User.TenantId)

	if cerr != nil {
		ErrorLog.Printf("getchannel error: %s", cerr)
	}

	uniqueCats := make(map[int]categories.CatgoriesOrd)
	order := []int{}

	for _, item := range FinalSelectedCategories {
		for _, cat := range item.Categories {
			if _, exists := uniqueCats[cat.Id]; !exists {
				uniqueCats[cat.Id] = cat
				order = append(order, cat.Id)
			}
		}
	}

	var categories []categories.CatgoriesOrd
	for _, id := range order {
		categories = append(categories, uniqueCats[id])
	}

	profile, member_details, ProfileName := GetProfileName(c, User.TenantId)
	chnentry, ecount, _, err := ChannelConfigWP.FlexibleChannelEntriesList(chn.EntriesInputs{Profile: profile, NoDirectAccess: true, UserRoleId: member_details.MemberGroupId, SlugName: channelname, CategorySlug: categoryname, Status: "Publish", TenantId: User.TenantId, Keyword: listingfilter.Keyword, GetAuthorDetails: true})
	if err != nil {
		fmt.Println(err.Error())
	}

	Count := int64(ecount)
	for i := range chnentry {

		first := chnentry[i].AuthorDetail.FirstName
		last := chnentry[i].AuthorDetail.LastName

		var firstn string
		if first != "" {
			firstn = strings.ToUpper(first[:1])
		}

		var lastn string
		if last != "" {
			lastn = strings.ToUpper(last[:1])
		}

		name := firstn + lastn
		chnentry[i].NameString = name

		chnentry[i].TagsArray = strings.FieldsFunc(chnentry[i].Tags, func(r rune) bool {
			return r == ','
		})

		chnentry[i].CreatedDate = chnentry[i].CreatedOn.In(TZONE).Format(Datelayout)

		if chnentry[i].ProfileImagePath != "" {
			chnentry[i].ProfileImagePath = "/image-resize?name=" + chnentry[i].ProfileImagePath
		}
		categoryid, _ := strconv.Atoi(chnentry[i].CategoriesId)
		categorydetails, _ := CategoryConfig.GetSubCategoryDetails(categoryid, User.TenantId)
		chnentry[i].CategoryGroup = categorydetails.CategoryName
		chnentry[i].Description = StripHTML(chnentry[i].Description)
		channelinfo, _, _ := ChannelConfigWP.GetChannelsById(chnentry[i].ChannelId, User.TenantId)

		chnentry[i].ChannelName = strings.ToLower(strings.ReplaceAll(channelinfo.ChannelName, " ", "-"))

		chnentry[i].Slug = "/" + channelDetail.SlugName + "/" + chnentry[i].Slug

		chnentry[i].AuthorDetail.NameString = chnentry[i].NameString

	}

	menudata, _ := MenuConfig.GetMenuBySlugName(channelname, website.Id, User.TenantId)

	seodetail, err := MenuConfig.SeoDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	seodetail.PageTitle = channelDetail.SeoTitle
	seodetail.PageDescription = channelDetail.SeoDescription
	seodetail.PageKeyword = channelDetail.SeoKeyword

	settingsdetail, err := MenuConfig.SettingsDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	var templates []ChannelTemplate

	err = json.Unmarshal(settingsdetail.TemplateType, &templates)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		// return
	}

	Template := template_name

	var tmpl *template.Template
	var erre error
	var TemplateHTML string

	TemplateHTML = "default.html"

	for _, item := range templates {

		if item.ID == channelDetail.Id {

			TemplateHTML = item.TemplateType

		}
	}

	tmpl, erre = template.ParseFiles(
		"websites/themes/"+Template+"/layouts/partials/header.html",
		"websites/themes/"+Template+"/layouts/partials/footer.html",
		"websites/themes/"+Template+"/layouts/partials/head.html",
		"websites/themes/"+Template+"/layouts/channels/"+TemplateHTML+"",
	)

	if erre != nil {
		fmt.Println(err, "templateerr")
	}

	var paginationendcount = len(chnentry) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(pageno, int(Count), limt)

	UserDetailsFunction(c)

	memberdet, _ := c.Get("userdetails")

	AllEntryList, _ := AllEntryList(User.TenantId, website.Id)

	websitemenu, _ := c.Cookie("websitemenu")
	if websitemenu == "" {
		websitemenu = "false"
	}
	PageHeading := menudata.Name
	RenderTemplate(c, tmpl, TemplateHTML, gin.H{"Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3}, "menulist": newmenulist, "PageHeading": PageHeading, "searchlist": AllEntryList, "Entries": chnentry, "template_name": template_name, "seodetail": seodetail, "settingsdetail": settingsdetail, "ChannelName": channelDetail.ChannelName, "Count": Count, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Limit": limt, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "memberdet": memberdet, "profile": profile, "websitemenu": websitemenu, "memberprofile": member_details, "profilename": ProfileName, "Route": "/channel/" + channelname, "Filter": listingfilter, "FinalCategories": categories, "categoryslugname": categoryname, "Datelayout": Datelayout})

}

func EntryDetailsPage(c *gin.Context) {

	// template_name := c.Query("template_name")

	var slug string
	var EntryCreateOn time.Time

	slug = strings.ToLower(c.Param("entryslug"))

	channelname := strings.ToLower(c.Param("slug"))

	User, website, err := GetTenantByHost(c)
	if err != nil {
		fmt.Println(err)
	}
	channelDetail, err := ChannelConfigWP.ChannelDetail(chn.Channels{Slug: channelname, TenantId: User.TenantId})
	if channelDetail.Id == 0 || err != nil {
		c.Redirect(301, "/404-TemplatePage")
		return
	}

	templatedetails, err := MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	fmt.Println(templatedetails, "detailtemplate")

	if err != nil {
		fmt.Println(err)
	}

	template_name := strings.ToLower(templatedetails.TemplateName)
	template_name = strings.ReplaceAll(template_name, " ", "_")
	// templatedetails, _ := MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	newmenulist, _ := MenuConfig.GetmenusByTenantId(website.Id, User.TenantId)

	for i := range newmenulist {
		newmenulist[i].Description = StripHTML(newmenulist[i].Description)
	}
	channelEntry, _, err := ChannelConfigWP.FetchChannelEntryDetail(chn.EntriesInputs{Slug: slug, ChannelName: channelname, TenantId: User.TenantId, GetAuthorDetails: true, GetLinkedCategories: true, GetAdditionalFields: true}, nil)

	if err != nil || channelEntry.Id == 0 {

		c.Redirect(301, "/404-TemplatePage")
		return

	}
	// plainText := StripHTML(channelEntry.Description)

	// channelEntry.Description = plainText

	channelEntry.CreatedDate = channelEntry.CreatedOn.In(TZONE).Format("Jan 2 2006")
	if channelEntry.AuthorDetail.ProfileImagePath != "" {
		channelEntry.AuthorDetail.ProfileImagePath = "/image-resize?name=" + channelEntry.AuthorDetail.ProfileImagePath
	}

	if len(channelEntry.Categories) > 0 && len(channelEntry.Categories[0]) > 0 {
		channelEntry.CategoryGroup = channelEntry.Categories[0][0].CategoryName
	} else {
		channelEntry.CategoryGroup = ""
	}
	var desc = channelEntry.Description
	var content template.HTML
	var chnentry []chn.Tblchannelentries

	content = template.HTML(desc)
	chnentry, _, _, _ = ChannelConfigWP.ChannelEntriesList(chn.Entries{SlugName: channelname, Status: "Published"}, User.TenantId)

	var filteredEntries []chn.Tblchannelentries

	for i := range chnentry {

		chnentry[i].CreatedDate = chnentry[i].CreatedOn.In(TZONE).Format("Jan 2 2006")

		if chnentry[i].ProfileImagePath != "" {
			chnentry[i].ProfileImagePath = "/image-resize?name=" + chnentry[i].ProfileImagePath
		}
		categoryid, _ := strconv.Atoi(chnentry[i].CategoriesId)
		categorydetails, _ := CategoryConfig.GetSubCategoryDetails(categoryid, User.TenantId)
		chnentry[i].CategoryGroup = categorydetails.CategoryName
		chnentry[i].Description = StripHTML(chnentry[i].Description)
		channelinfo, _, _ := ChannelConfigWP.GetChannelsById(chnentry[i].ChannelId, User.TenantId)

		chnentry[i].ChannelName = strings.ToLower(strings.ReplaceAll(channelinfo.ChannelName, " ", "-"))

		if channelEntry.AuthorDetail.ProfileImagePath == "" {

			channelEntry.AuthorDetail.NameString = chnentry[i].NameString
		}

		if chnentry[i].Id != channelEntry.Id {
			filteredEntries = append(filteredEntries, chnentry[i])
		}

	}

	seodetail, err := MenuConfig.SeoDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	seodetail.PageKeyword = channelEntry.Keyword

	seodetail.PageTitle = channelEntry.MetaTitle

	seodetail.PageDescription = channelEntry.Description

	settingsdetail, err := MenuConfig.SettingsDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}
	UserDetailsFunction(c)

	memberdet, _ := c.Get("userdetails")

	AllEntryList, _ := AllEntryList(User.TenantId, website.Id)

	var templates []ChannelTemplate

	err = json.Unmarshal(settingsdetail.TemplateType, &templates)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		// return
	}

	Template := template_name

	var tmpl *template.Template
	var erre error
	var TemplateHTML string

	TemplateHTML = "default.html"

	for _, item := range templates {

		if item.ID == channelDetail.Id {

			TemplateHTML = item.TemplateType

		}
	}

	tmpl, erre = template.ParseFiles(
		"websites/themes/"+Template+"/layouts/partials/header.html",
		"websites/themes/"+Template+"/layouts/partials/footer.html",
		"websites/themes/"+Template+"/layouts/partials/head.html",
		"websites/themes/"+Template+"/layouts/channel_entries/"+TemplateHTML+"",
	)

	if erre != nil {

		fmt.Println(err, "templateerr")
	}

	profile, member_details, ProfileName := GetProfileName(c, User.TenantId)

	websitemenu, _ := c.Cookie("websitemenu")
	if websitemenu == "" {
		websitemenu = "false"
	}

	PageHeading := channelDetail.ChannelName

	RenderTemplate(c, tmpl, TemplateHTML, gin.H{"entrydetail": channelEntry, "PageHeading": PageHeading, "searchlist": AllEntryList, "menulist": newmenulist, "relatedentry": chnentry, "template_name": template_name, "seodetail": seodetail, "settingsdetail": settingsdetail, "ChannelName": channelDetail.ChannelName, "filteredEntries": filteredEntries, "memberdet": memberdet, "title": channelEntry.Title, "content": content, "profile": profile, "websitemenu": websitemenu, "memberprofile": member_details, "profilename": ProfileName, "CreatedDate": channelEntry.CreatedDate, "ReadingTime": channelEntry.ReadingTime, "Datelayout": Datelayout, "createon": EntryCreateOn})

}
func CategoryBaseEntryList(c *gin.Context) {

	categoryname := strings.ToLower(c.Param("categoryname"))

	channelname := strings.ToLower(c.Param("slug"))

	User, website, _ := GetTenantByHost(c)

	channelDetail, err := ChannelConfigWP.ChannelDetail(chn.Channels{Slug: channelname, TenantId: User.TenantId})

	fmt.Println(channelDetail, c.Param("slug"), channelDetail.Id, "channeldetildfd")
	if channelDetail.Id == 0 || err != nil {

		c.Redirect(301, "/404-TemplatePage")
		return
	}

	templatedetails, err := MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	if err != nil {
		fmt.Println(err)
	}

	template_name := strings.ToLower(templatedetails.TemplateName)
	template_name = strings.ReplaceAll(template_name, " ", "_")

	newmenulist, merr := MenuConfig.GetmenusByTenantId(website.Id, User.TenantId)

	if merr != nil {

		fmt.Println(merr)
	}
	for i := range newmenulist {
		newmenulist[i].Description = StripHTML(newmenulist[i].Description)
	}
	fullurl := strings.Trim(c.Request.URL.Path, "")

	menudata, _ := MenuConfig.GetMenuByUrlPath(fullurl, website.Id, User.TenantId)

	// if menudata.Id == 0 || err != nil {
	//  c.Redirect(301, "/404-TemplatePage")
	//  return
	// }

	entries, count, _, err := ChannelConfigWP.FlexibleChannelEntriesList(chn.EntriesInputs{
		Status:              "Publish",
		CategorySlug:        categoryname,
		TenantId:            User.TenantId,
		GetAdditionalFields: true,
		SlugName:            channelname,
		GetAuthorDetails:    true,
	})
	for i := range entries {
		entries[i].TagsArray = strings.FieldsFunc(
			entries[i].Tags,
			func(r rune) bool { return r == ',' },
		)
	}

	for i := range entries {

		first := entries[i].AuthorDetail.FirstName
		last := entries[i].AuthorDetail.LastName

		var firstn string
		if first != "" {
			firstn = strings.ToUpper(first[:1])
		}

		var lastn string
		if last != "" {
			lastn = strings.ToUpper(last[:1])
		}

		name := firstn + lastn

		entries[i].NameString = name

		entries[i].CreatedDate = entries[i].CreatedOn.In(TZONE).Format("Jan 2 2006")

		if entries[i].ProfileImagePath != "" {
			entries[i].ProfileImagePath = "/image-resize?name=" + entries[i].ProfileImagePath
		}
		categoryid, _ := strconv.Atoi(entries[i].CategoriesId)
		categorydetails, _ := CategoryConfig.GetSubCategoryDetails(categoryid, User.TenantId)
		entries[i].CategoryGroup = categorydetails.CategoryName
		entries[i].Description = StripHTML(entries[i].Description)

		channelinfo, _, _ := ChannelConfigWP.GetChannelsById(channelDetail.Id, User.TenantId)

		entries[i].ChannelName = strings.ToLower(strings.ReplaceAll(channelinfo.ChannelName, " ", "-"))

		entries[i].Slug = "/" + channelinfo.SlugName + "/" + entries[i].Slug

	}

	seodetail, err := MenuConfig.SeoDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	categoryinfo, err := CategoryConfig.GetSubCategoryDetails(menudata.TypeId, User.TenantId)
	if err != nil {
		fmt.Println(err)
	}

	seodetail.PageTitle = categoryinfo.SeoTitle
	seodetail.PageDescription = categoryinfo.SeoDescription
	seodetail.PageKeyword = categoryinfo.SeoKeyword

	settingsdetail, err := MenuConfig.SettingsDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}
	entrylist, _ := AllEntryList(User.TenantId, website.Id)

	var templates []ChannelTemplate

	err = json.Unmarshal(settingsdetail.TemplateType, &templates)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		// return
	}
	_, FinalSelectedCategories, cerr := ChannelConfig.GetChannelsById(channelDetail.Id, User.TenantId)

	if cerr != nil {
		ErrorLog.Printf("getchannel error: %s", cerr)
	}

	uniqueCats := make(map[int]categories.CatgoriesOrd)
	order := []int{}

	for _, item := range FinalSelectedCategories {
		for _, cat := range item.Categories {
			if _, exists := uniqueCats[cat.Id]; !exists {
				uniqueCats[cat.Id] = cat
				order = append(order, cat.Id)
			}
		}
	}

	var categories []categories.CatgoriesOrd
	for _, id := range order {
		categories = append(categories, uniqueCats[id])
	}

	Template := template_name

	var TemplateHTML string

	TemplateHTML = "default.html"

	for _, item := range templates {

		if item.ID == channelDetail.Id {

			TemplateHTML = item.TemplateType

		}
	}
	tmpl, erre := template.ParseFiles(
		"websites/themes/"+Template+"/layouts/partials/header.html",
		"websites/themes/"+Template+"/layouts/partials/footer.html",
		"websites/themes/"+Template+"/layouts/partials/head.html",
		"websites/themes/"+Template+"/layouts/channels/"+TemplateHTML+"",
	)

	// tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/themes/"+Template+"/layouts/_default/category_list.html")

	if erre != nil {

		fmt.Println(err, "templateerr")
	}

	UserDetailsFunction(c)

	memberid := c.GetInt("member_id")

	profile, member_details, ProfileName := GetProfileName(c, User.TenantId)

	var member models.TblMembers

	member, err = models.GetMemberById(memberid, User.TenantId)

	if err != nil {
		fmt.Println(err)
	}

	websitemenu, _ := c.Cookie("websitemenu")
	if websitemenu == "" {
		websitemenu = "false"
	}

	PageHeading := menudata.Name

	RenderTemplate(c, tmpl, TemplateHTML, gin.H{"websitemenu": websitemenu, "searchlist": entrylist, "menulist": newmenulist, "Count": count, "Entries": entries, "menudata": menudata, "template_name": template_name, "seodetail": seodetail, "settingsdetail": settingsdetail, "ProfileImagePath": member.ProfileImagePath, "profile": profile, "memberprofile": member_details, "profilename": ProfileName, "PageHeading": PageHeading, "FinalCategories": categories, "ChannelName": channelDetail.ChannelName, "categoryslugname": categoryname, "ChannelSlugName": channelDetail.SlugName})

}

//membership page//

func MembershipList(c *gin.Context) {

	user, website, err := GetTenantByHost(c)
	if err != nil {
		fmt.Println(err)
	}

	newmenulist, err := GetMenuItemsListByTenantID(user.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	membershipLevels, _, _ := MembershipConfig.MembershipLevelsList(0, 10, membership.Filter{}, user.TenantId)

	templatedetails, err := MenuConfig.GetTemplateById(user.GoTemplateDefault, user.TenantId)
	if err != nil {
		fmt.Println(err)
	}

	templateName := GetTemplateNamee(c, templatedetails.TemplateName)

	tmpl, err := template.ParseFiles("websites/themes/"+templateName+"/layouts/partials/header.html", "websites/themes/"+templateName+"/layouts/partials/footer.html", "websites/themes/"+templateName+"/layouts/partials/head.html", "websites/common/layouts/membership.html")
	if err != nil {
		fmt.Println(err, "templateerr")
	}

	RenderTemplate(c, tmpl, "membership.html", gin.H{"menulist": newmenulist, "membership": membershipLevels, "template_name": c.Query("template_name")})
}

func MembershiDetail(c *gin.Context) {

	User, website, err := GetTenantByHost(c)
	if err != nil {
		fmt.Println(err)
	}
	newmenulist, err := GetMenuItemsListByTenantID(User.TenantId, website.Id)

	if err != nil {
		fmt.Println(err)
	}

	templatedetails, _ := MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	Template := GetTemplateNamee(c, templatedetails.TemplateName)

	membershipid, _ := strconv.Atoi(c.Param("id"))


	GetEditMembership, err := MembershipConfig.MembershiplevelEdit(membershipid, User.TenantId)

	if err != nil {
		log.Fatal("Edit membership level error", err)
		// c.AbortWithError(500, err)
	}

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/common/layouts/checkout.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	RenderTemplate(c, tmpl, "checkout.html", gin.H{"menulist": newmenulist, "template_name": c.Query("template_name"), "membershipdetail": GetEditMembership})

}

func GetMenuItemsListByTenantID(tenantId string, websiteid int) ([]menu.TblMenus, error) {

	newmenulist, _ := MenuConfig.GetmenusByTenantId(websiteid, tenantId)

	return newmenulist, nil
}
func GetTenantByHost(c *gin.Context) (team.TblUser, menu.TblWebsite, error) {

	TENANTID := os.Getenv("TENANTID")

	host := c.Request.Host

	subdomain := GetSubdomain(host)

	var website menu.TblWebsite

	if TENANTID == "" || subdomain != "" {

		website, _ = MenuConfig.GetWebsiteByName(subdomain)
	} else {

		website, _ = MenuConfig.GetWebsiteById(1, TENANTID)

	}

	if website.Id == 0 {
		c.HTML(404, "nodata.html", nil)
		c.Abort()
		return team.TblUser{}, menu.TblWebsite{}, fmt.Errorf("website not found")
	}

	User, _, _ := NewTeamWP.GetUserById(website.CreatedBy, []int{})
	User.GoTemplateDefault = website.TemplateId

	return User, website, nil
}

func AllEntryList(tenatid string, webid int) ([]chn.Tblchannelentries, error) {

	menulist, _ := GetMenuItemsListByTenantID(tenatid, webid)

	var entries []chn.Tblchannelentries

	for _, val := range menulist {

		if val.ParentId != 0 {

			entries, _, _, _ = ChannelConfigWP.ChannelEntriesList(chn.Entries{ChannelId: val.TypeId, Status: "Published"}, tenatid)

			for i := range entries {

				entries[i].CreatedDate = entries[i].CreatedOn.In(TZONE).Format("Jan 2 2006")

				if entries[i].ProfileImagePath != "" {
					entries[i].ProfileImagePath = "/image-resize?name=" + entries[i].ProfileImagePath
				}
				categoryid, _ := strconv.Atoi(entries[i].CategoriesId)
				categorydetails, _ := CategoryConfig.GetSubCategoryDetails(categoryid, tenatid)
				entries[i].CategoryGroup = categorydetails.CategoryName
				entries[i].Description = StripHTML(entries[i].Description)

			}
		}
	}

	return entries, nil
}

func CategoryDetailsPage(c *gin.Context) {

	slug := c.Param("dynamicname")
	if len(slug) > 0 {
		slug = slug[1:]
	}

	User, website, err := GetTenantByHost(c)
	if err != nil {
		fmt.Println(err)
	}

	templatedetails, err := MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	if err != nil {
		fmt.Println(err)
	}

	template_name := strings.ToLower(templatedetails.TemplateName)
	template_name = strings.ReplaceAll(template_name, " ", "_")
	// templatedetails, _ := MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	newmenulist, _ := MenuConfig.GetmenusByTenantId(website.Id, User.TenantId)
	for i := range newmenulist {
		newmenulist[i].Description = StripHTML(newmenulist[i].Description)
	}

	channelEntry, _, err := ChannelConfigWP.FetchChannelEntryDetail(chn.EntriesInputs{Slug: slug, TenantId: User.TenantId, GetAuthorDetails: true, GetLinkedCategories: true, GetAdditionalFields: true}, nil)

	if err != nil || channelEntry.Id == 0 {

		c.Redirect(301, "/404-TemplatePage")
		return

	}
	// plainText := StripHTML(channelEntry.Description)

	// channelEntry.Description = plainText

	channelEntry.CreatedDate = channelEntry.CreatedOn.In(TZONE).Format("Jan 2 2006")
	if channelEntry.AuthorDetail.ProfileImagePath != "" {
		channelEntry.AuthorDetail.ProfileImagePath = "/image-resize?name=" + channelEntry.AuthorDetail.ProfileImagePath
	}

	if len(channelEntry.Categories) > 0 && len(channelEntry.Categories[0]) > 0 {
		channelEntry.CategoryGroup = channelEntry.Categories[0][0].CategoryName
	} else {
		channelEntry.CategoryGroup = ""
	}
	var desc = channelEntry.Description
	var content template.HTML
	var chnentry []chn.Tblchannelentries
	if template_name == "jobs" {
		re := regexp.MustCompile(`class="header1"`)
		desc = re.ReplaceAllString(desc, "")
		desc = strings.ReplaceAll(desc, "  ", " ")
		content = template.HTML(desc)
		chnentry, _, _, _ = ChannelConfigWP.ChannelEntriesList(chn.Entries{SlugName: slug, CategoryName: channelEntry.CategoryGroup, AdditionalFields: true, MemberProfile: true, Status: "Published"}, User.TenantId)

	} else {

		content = template.HTML(desc)
		chnentry, _, _, _ = ChannelConfigWP.ChannelEntriesList(chn.Entries{SlugName: slug, Status: "Published"}, User.TenantId)

	}
	var filteredEntries []chn.Tblchannelentries

	for i := range chnentry {

		chnentry[i].CreatedDate = chnentry[i].CreatedOn.In(TZONE).Format("Jan 2 2006")

		if chnentry[i].ProfileImagePath != "" {
			chnentry[i].ProfileImagePath = "/image-resize?name=" + chnentry[i].ProfileImagePath
		}
		categoryid, _ := strconv.Atoi(chnentry[i].CategoriesId)
		categorydetails, _ := CategoryConfig.GetSubCategoryDetails(categoryid, User.TenantId)
		chnentry[i].CategoryGroup = categorydetails.CategoryName
		chnentry[i].Description = StripHTML(chnentry[i].Description)
		channelinfo, _, _ := ChannelConfigWP.GetChannelsById(chnentry[i].ChannelId, User.TenantId)

		chnentry[i].ChannelName = strings.ToLower(strings.ReplaceAll(channelinfo.ChannelName, " ", "-"))

		if channelEntry.AuthorDetail.ProfileImagePath == "" {

			channelEntry.AuthorDetail.NameString = chnentry[i].NameString
		}

		if chnentry[i].Id != channelEntry.Id {
			filteredEntries = append(filteredEntries, chnentry[i])
		}

	}

	seodetail, err := MenuConfig.SeoDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	seodetail.PageKeyword = channelEntry.Keyword

	seodetail.PageTitle = channelEntry.MetaTitle

	seodetail.PageDescription = channelEntry.Description

	settingsdetail, err := MenuConfig.SettingsDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}
	UserDetailsFunction(c)

	memberdet, _ := c.Get("userdetails")

	entrylist, _ := AllEntryList(User.TenantId, website.Id)

	Template := template_name

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/themes/"+Template+"/layouts/_default/category_detail.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	profile, member_details, ProfileName := GetProfileName(c, User.TenantId)

	websitemenu, _ := c.Cookie("websitemenu")
	if websitemenu == "" {
		websitemenu = "false"
	}

	RenderTemplate(c, tmpl, "category_detail.html", gin.H{"entrydetail": channelEntry, "searchlist": entrylist, "menulist": newmenulist, "relatedentry": chnentry, "template_name": template_name, "seodetail": seodetail, "settingsdetail": settingsdetail, "filteredEntries": filteredEntries, "memberdet": memberdet, "content": content, "profile": profile, "websitemenu": websitemenu, "memberprofile": member_details, "profilename": ProfileName})

}

func GlobalSearch(c *gin.Context) {

	User, website, _ := GetTenantByHost(c)

	templatedetails, err := MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	if err != nil {
		fmt.Println(err)
	}

	var route string

	if c.Request.Referer() == "" || strings.Contains(c.Request.Referer(), "/search") {
		route = "/"
	} else {
		route = c.Request.Referer()
	}
	template_name := strings.ToLower(templatedetails.TemplateName)
	template_name = strings.ReplaceAll(template_name, " ", "_")

	newmenulist, merr := MenuConfig.GetmenusByTenantId(website.Id, User.TenantId)

	for i := range newmenulist {
		newmenulist[i].Description = StripHTML(newmenulist[i].Description)
	}

	if merr != nil {

		fmt.Println(merr)
	}
	memberid := c.GetInt("member_id")

	profile, member_details, ProfileName := GetProfileName(c, User.TenantId)

	seodetail, err := MenuConfig.SeoDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	settingsdetail, err := MenuConfig.SettingsDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	Template := template_name

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/themes/"+Template+"/layouts/_default/search.html")

	if err != nil {
		fmt.Println(err, "templateerr")
	}

	UserDetailsFunction(c)

	var member models.TblMembers

	member, err = models.GetMemberById(memberid, User.TenantId)

	if err != nil {
		fmt.Println(err)
	}

	PageHeading := "Search Results"

	websitemenu, _ := c.Cookie("websitemenu")
	if websitemenu == "" {
		websitemenu = "false"
	}

	searchKeyWord := strings.TrimSpace(c.Query("keyword"))

	//  filter the listing data using the search keyword
	var (
		listingfilter listing.Filter
		entriesFilter chn.Entries
	)

	if searchKeyWord != "" {
		listingfilter.Keyword = searchKeyWord
		entriesFilter.Keyword = searchKeyWord

	}

	// getting the list from the listing based on the keyword
	listings, err := ListingConfig.GetListingsList(listing.ListingInput{TenantId: User.TenantId, Filter: listingfilter})
	if err != nil {
		fmt.Println("error getting the list data :", err)
		return

	}

	if len(listings) > 0 {
		for index := range listings {
			listings[index].Slug = "/listing/" + listings[index].EntrySlug
		}
	}

	// // getting the list from the entries based on the keyword
	// entriesList, _, _, err := ChannelConfigWP.ChannelEntriesList(entriesFilter, User.TenantId)
	// if err != nil {
	// 	fmt.Println("error getting the list data :", err)
	// 	return

	// }

	// fmt.Println("lslslls", len(listings), len(entriesList))
	// count := len(listings) + len(entriesList)
	count := len(listings)

	RenderTemplate(c, tmpl, "search.html", gin.H{"menulist": newmenulist, "alllist": true, "template_name": template_name, "seodetail": seodetail, "settingsdetail": settingsdetail, "ProfileImagePath": member.ProfileImagePath, "profile": profile, "PageHeading": PageHeading, "url": "/search", "websitemenu": websitemenu, "memberprofile": member_details, "profilename": ProfileName, "Filter": listingfilter, "Route": route, "listingslist": listings, "Count": count})
}

func LoadPages(c *gin.Context) {

	// template_name := c.Query("template_name")

	pagename := strings.ToLower(c.Param("slug"))
	subpage := strings.ToLower(c.Param("entryslug"))
	childsubpage := strings.ToLower(c.Param("subpage"))

	User, website, err := GetTenantByHost(c)
	if err != nil {
		fmt.Println(err)
		return // Exit early on tenant error
	}

	templatedetails, err := MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)
	if err != nil {
		fmt.Println(err)
		return
	}

	template_name := strings.ToLower(templatedetails.TemplateName)
	template_name = strings.ReplaceAll(template_name, " ", "_")

	newmenulist, merr := MenuConfig.GetmenusByTenantId(website.Id, User.TenantId)
	if merr != nil {
		fmt.Println(merr)
	}

	for i := range newmenulist {
		newmenulist[i].Description = StripHTML(newmenulist[i].Description)
	}

	if pagename != "" {

		pagedetails, perr := MenuConfig.GetPageBySlug(pagename, User.TenantId)

		if perr != nil || pagedetails.Id == 0 {
			c.Redirect(301, "/404-TemplatePage")
			return

		}
	}

	if subpage != "" {

		pagedetails, perr := MenuConfig.GetPageBySlug(subpage, User.TenantId)

		if perr != nil || pagedetails.Id == 0 {
			c.Redirect(301, "/404-TemplatePage")
			return

		}
	}

	// PRIORITY: childsubpage > subpage > pagename
	finalSlug := pagename
	if childsubpage != "" {
		finalSlug = childsubpage
	} else if subpage != "" {
		finalSlug = subpage
	}

	var pagedetails menu.TblTemplatePages

	var perr error
	pagedetails, perr = MenuConfig.GetPageBySlug(finalSlug, User.TenantId)

	if perr != nil || pagedetails.Id == 0 {
		c.Redirect(301, "/404-TemplatePage")
		return

	}

	if pagedetails.Status == 0 {

		c.Redirect(301, "/404-TemplatePage")
		return
	}

	pagelist, _, err := MenuConfigWp.GetTemplatePageList(100, 0, menu.Filter{PageId: pagedetails.Id}, User.TenantId, website.Id)

	for i := range pagelist {

		if subpage != "" {
			pagelist[i].Slug = "/" + pagename + "/" + subpage + "/" + pagelist[i].Slug
		} else {

			pagelist[i].Slug = "/" + pagename + "/" + pagelist[i].Slug
		}

	}

	var pagepath string

	seodetail, err := MenuConfig.SeoDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	seodetail.PageKeyword = pagedetails.MetaKeywords

	seodetail.PageTitle = pagedetails.MetaTitle

	seodetail.PageDescription = pagedetails.MetaDescription

	settingsdetail, err := MenuConfig.SettingsDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}
	UserDetailsFunction(c)

	memberid := c.GetInt("member_id")

	profile, member_details, ProfileName := GetProfileName(c, User.TenantId)

	var member models.TblMembers

	member, err = models.GetMemberById(memberid, User.TenantId)

	if err != nil {
		fmt.Println(err)
	}
	memberdet, _ := c.Get("userdetails")
	Template := GetTemplateNamee(c, templatedetails.TemplateName)

	content := template.HTML(pagedetails.PageDescription)

	entrylist, _ := AllEntryList(User.TenantId, website.Id)

	websitemenu, _ := c.Cookie("websitemenu")
	if websitemenu == "" {
		websitemenu = "false"
	}

	var tmpl *template.Template

	switch pagedetails.PageType {
	case "static-page":

		pagepath = pagedetails.CustomPagePath

		tmpl, err = template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/themes/"+Template+"/pages/static-pages/"+pagedetails.CustomPagePath)

	case "landing-page":

		pagepath = "landing_list.html"

		tmpl, err = template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/themes/"+Template+"/pages/landing-pages/landing_list.html")

	default:

		pagepath = "static_page.html"

		tmpl, err = template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/themes/"+Template+"/pages/editor-pages/static_page.html")

	}

	if err != nil {

		fmt.Println(err, "templateerr")
	}
	// =====piccosoft contact us page
	CountryList, _ := models.GetCountryLIst()

	reCAPTCHA_SITE_KEY := os.Getenv("reCAPTCHA_SITE_KEY")
	// ======
	RenderTemplate(c, tmpl, pagepath, gin.H{"PageDescription": content, "pagelist": pagelist, "seodetail": seodetail, "searchlist": entrylist, "menulist": newmenulist, "template_name": template_name, "settingsdetail": settingsdetail, "memberdet": memberdet, "profile": profile, "memberprofile": member_details, "profilename": ProfileName, "ProfileImagePath": member.ProfileImagePath, "websitemenu": websitemenu, "CountryList": CountryList, "CAPTCHASITEKEY": reCAPTCHA_SITE_KEY, "csrf": csrf.GetToken(c), "ChannelName": pagedetails.Name})
}
