package templatecontroller

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"regexp"
	"spurt-cms/controllers"
	"spurt-cms/models"
	"strconv"

	"strings"

	"github.com/gin-gonic/gin"
	chn "github.com/spurtcms/channels"
	"github.com/spurtcms/membership"
	"github.com/spurtcms/menu"
	"github.com/spurtcms/team"
	"golang.org/x/net/html"
)

func ListingsList(c *gin.Context) {

	menuname := c.Param("menuname")

	tag := c.Query("tag")

	User, website, _ := GetTenantByHost(c)

	templatedetails, err := controllers.MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	if err != nil {
		fmt.Println(err)
	}

	template_name := strings.ToLower(templatedetails.TemplateName)
	template_name = strings.ReplaceAll(template_name, " ", "_")

	newmenulist, merr := GetMenuItemsListByTenantID(User.TenantId, website.Id)

	if merr != nil {

		fmt.Println(merr)
	}

	menudata, _ := controllers.MenuConfig.GetMenuBySlugName(menuname, website.Id, User.TenantId)

	listingIds := strings.Split(menudata.ListingsIds, ",")

	listings, _ := controllers.ListingConfig.GetListingsByIds(listingIds, tag, User.TenantId)

	seodetail, err := controllers.MenuConfig.SeoDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	settingsdetail, err := controllers.MenuConfig.SettingsDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	Template := template_name

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/themes/"+Template+"/layouts/_default/content_list.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	UserDetailsFunction(c)

	memberid := c.GetInt("member_id")

	var profile bool

	if memberid != 0 {

		profile = true

	} else {

		profile = false

	}

	var member models.TblMembers

	member, err = models.GetMemberById(memberid, User.TenantId)

	if err != nil {
		fmt.Println(err)
	}

	PageHeading := menudata.Name

	controllers.RenderTemplate(c, tmpl, "content_list.html", gin.H{"menulist": newmenulist, "listingslist": listings, "menudata": menudata, "template_name": template_name, "seodetail": seodetail, "settingsdetail": settingsdetail, "ProfileImagePath": member.ProfileImagePath, "profile": profile, "PageHeading": PageHeading})

}

func ListingsDetailsPage(c *gin.Context) {

	menuname := c.Param("menuname")

	listingname := c.Param("listingname")

	User, website, _ := GetTenantByHost(c)

	templatedetails, err := controllers.MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	if err != nil {
		fmt.Println(err)
	}

	template_name := strings.ToLower(templatedetails.TemplateName)
	template_name = strings.ReplaceAll(template_name, " ", "_")

	newmenulist, merr := GetMenuItemsListByTenantID(User.TenantId, website.Id)

	if merr != nil {

		fmt.Println(merr)
	}

	menudata, _ := controllers.MenuConfig.GetMenuBySlugName(menuname, website.Id, User.TenantId)

	listingIds := strings.Split(menudata.ListingsIds, ",")

	listingData, _ := controllers.ListingConfig.GetListingBySlugName(listingIds, listingname, User.TenantId)

	var entries chn.Tblchannelentries

	var content template.HTML

	if listingData.Slug == listingname {

		entries, _ = controllers.ChannelConfig.EntryDetailsById(chn.IndivEntriesReq{EntryId: listingData.EntryId}, User.TenantId)

		content = template.HTML(entries.Description)

	}

	seodetail, err := controllers.MenuConfig.SeoDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	settingsdetail, err := controllers.MenuConfig.SettingsDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	Template := template_name

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/themes/"+Template+"/layouts/_default/content_detail.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	UserDetailsFunction(c)

	memberid := c.GetInt("member_id")

	var profile bool

	if memberid != 0 {

		profile = true

	} else {

		profile = false

	}

	var member models.TblMembers

	member, err = models.GetMemberById(memberid, User.TenantId)

	if err != nil {
		fmt.Println(err)
	}

	PageHeading := menudata.Name

	controllers.RenderTemplate(c, tmpl, "content_detail.html", gin.H{"menulist": newmenulist, "menudata": menudata, "listingData": listingData, "entries": entries, "content": content, "menuname": menuname, "template_name": template_name, "seodetail": seodetail, "settingsdetail": settingsdetail, "ProfileImagePath": member.ProfileImagePath, "profile": profile, "PageHeading": PageHeading})

}

func ChannelEntriesList(c *gin.Context) {

	channelname := c.Param("cname")

	// template_name := c.Query("template_name")

	User, website, _ := GetTenantByHost(c)

	templatedetails, err := controllers.MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	if err != nil {
		fmt.Println(err)
	}

	template_name := strings.ToLower(templatedetails.TemplateName)
	template_name = strings.ReplaceAll(template_name, " ", "_")
	// templatedetails, err := controllers.MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	newmenulist, merr := GetMenuItemsListByTenantID(User.TenantId, website.Id)

	if merr != nil {

		fmt.Println(merr)
	}

	var (
		limt, offset int
	)

	//get values from html form data
	// limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	fmt.Println("template_name::", template_name)

	switch template_name {
	case "Content chronicle":

		limt = 12

	case "Content Verse":

		limt = 15
	case "support_sphere":

		limt = 0
	default:
		limt = controllers.Limit
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}
	chnentry, _, ecount, _ := controllers.ChannelConfigWP.ChannelEntriesList(chn.Entries{ChannelName: channelname, Status: "Published", Limit: limt, Offset: offset}, User.TenantId)

	Count := int64(ecount)
	for i := range chnentry {

		chnentry[i].CreatedDate = chnentry[i].CreatedOn.In(controllers.TZONE).Format("Jan 2 2006")

		if chnentry[i].ProfileImagePath != "" {
			chnentry[i].ProfileImagePath = "/image-resize?name=" + chnentry[i].ProfileImagePath
		}
		categoryid, _ := strconv.Atoi(chnentry[i].CategoriesId)
		categorydetails, _ := controllers.CategoryConfig.GetSubCategoryDetails(categoryid, User.TenantId)
		chnentry[i].CategoryGroup = categorydetails.CategoryName
		chnentry[i].Description = StripHTML(chnentry[i].Description)

	}
	seodetail, err := controllers.MenuConfig.SeoDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	settingsdetail, err := controllers.MenuConfig.SettingsDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	Template := template_name

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/themes/"+Template+"/layouts/_default/content_list.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}
	fmt.Println(Count, "countddf")
	var paginationendcount = len(chnentry) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := controllers.Pagination(pageno, int(Count), limt)

	UserDetailsFunction(c)

	memberdet, _ := c.Get("userdetails")

	AllEntryList, _ := AllEntryList(User.TenantId, website.Id)

	controllers.RenderTemplate(c, tmpl, "content_list.html", gin.H{"Pagination": controllers.PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3}, "menulist": newmenulist, "searchlist": AllEntryList, "entries": chnentry, "template_name": template_name, "seodetail": seodetail, "settingsdetail": settingsdetail, "ChannelName": channelname, "Count": Count, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Limit": limt, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "memberdet": memberdet})

}

func StripHTML(input string) string {
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		return input
	}
	var buf bytes.Buffer
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			buf.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return buf.String()
}

func EntryDetailsPage(c *gin.Context) {

	// template_name := c.Query("template_name")

	var slug string

	slug = c.Param("dynamicname")

	arr := strings.Split(slug, "-")
	uuid := arr[len(arr)-1]

	lastHyphen := strings.LastIndex(slug, "-")
	if lastHyphen == -1 {
		c.String(400, "Invalid slug format")
		return
	}

	slug = slug[:lastHyphen]

	channelname := c.Param("cname")
	if len(slug) > 1 {
		slug = slug[1:]
	} else {
		slug = ""
	}
	User, website, err := GetTenantByHost(c)
	if err != nil {
		fmt.Println(err)
	}

	templatedetails, err := controllers.MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	if err != nil {
		fmt.Println(err)
	}

	template_name := strings.ToLower(templatedetails.TemplateName)
	template_name = strings.ReplaceAll(template_name, " ", "_")
	// templatedetails, _ := controllers.MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	newmenulist, _ := GetMenuItemsListByTenantID(User.TenantId, website.Id)

	channelEntry, _, err := controllers.ChannelConfigWP.FetchChannelEntryDetail(chn.EntriesInputs{Slug: slug, ChannelName: channelname, TenantId: User.TenantId, GetAuthorDetails: true, GetLinkedCategories: true, Uuid: uuid, GetAdditionalFields: true}, nil)

	if err != nil || channelEntry.Id == 0 {

		c.HTML(404, "nodata.html", nil)

		return

	}
	// plainText := StripHTML(channelEntry.Description)

	// channelEntry.Description = plainText

	channelEntry.CreatedDate = channelEntry.CreatedOn.In(controllers.TZONE).Format("Jan 2 2006")
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
		fmt.Println(channelEntry.CategoryGroup, channelname, "checkerepladf")
		chnentry, _, _, _ = controllers.ChannelConfigWP.ChannelEntriesList(chn.Entries{ChannelName: channelname, CategoryName: channelEntry.CategoryGroup, AdditionalFields: true, MemberProfile: true, Status: "Published"}, User.TenantId)

	} else {

		content = template.HTML(desc)
		chnentry, _, _, _ = controllers.ChannelConfigWP.ChannelEntriesList(chn.Entries{ChannelName: channelname, Status: "Published"}, User.TenantId)

	}
	var filteredEntries []chn.Tblchannelentries

	for i := range chnentry {

		chnentry[i].CreatedDate = chnentry[i].CreatedOn.In(controllers.TZONE).Format("Jan 2 2006")

		if chnentry[i].ProfileImagePath != "" {
			chnentry[i].ProfileImagePath = "/image-resize?name=" + chnentry[i].ProfileImagePath
		}
		categoryid, _ := strconv.Atoi(chnentry[i].CategoriesId)
		categorydetails, _ := controllers.CategoryConfig.GetSubCategoryDetails(categoryid, User.TenantId)
		chnentry[i].CategoryGroup = categorydetails.CategoryName
		chnentry[i].Description = StripHTML(chnentry[i].Description)

		if channelEntry.AuthorDetail.ProfileImagePath == "" {

			channelEntry.AuthorDetail.NameString = chnentry[i].NameString
		}

		if chnentry[i].Id != channelEntry.Id {
			filteredEntries = append(filteredEntries, chnentry[i])
		}

	}

	seodetail, err := controllers.MenuConfig.SeoDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	seodetail.PageKeyword = channelEntry.Keyword

	seodetail.PageTitle = channelEntry.MetaTitle

	seodetail.PageDescription = channelEntry.Description

	settingsdetail, err := controllers.MenuConfig.SettingsDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}
	UserDetailsFunction(c)

	memberdet, _ := c.Get("userdetails")

	AllEntryList, _ := AllEntryList(User.TenantId, website.Id)

	Template := template_name

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/themes/"+Template+"/layouts/_default/content_detail.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	controllers.RenderTemplate(c, tmpl, "content_detail.html", gin.H{"entrydetail": channelEntry, "searchlist": AllEntryList, "menulist": newmenulist, "relatedentry": chnentry, "template_name": template_name, "seodetail": seodetail, "settingsdetail": settingsdetail, "ChannelName": channelname, "filteredEntries": filteredEntries, "memberdet": memberdet, "content": content})

}

func CoursesDetails(c *gin.Context) {

	uuid := c.Param("uuid")

	lessonid, _ := strconv.Atoi(c.Param("lessonid"))

	// template_name := c.Query("template_name")

	User, _, err := GetTenantByHost(c)
	if err != nil {
		fmt.Println(err)
	}

	templatedetails, err := controllers.MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	if err != nil {
		fmt.Println(err)
	}

	template_name := strings.ToLower(templatedetails.TemplateName)
	template_name = strings.ReplaceAll(template_name, " ", "_")

	courses, _ := controllers.CoursesConfig.UuidtoCourseID(uuid, User.TenantId)

	sectionlist, serr := controllers.CoursesConfig.ListSections(courses.Id, User.TenantId)
	if serr != nil {
		fmt.Println(serr)
	}

	lessonlist, lerr := controllers.CoursesConfig.ListLessons(courses.Id, User.TenantId)
	if lerr != nil {
		fmt.Println(lerr)
	}

	lesson, err := controllers.CoursesConfig.EditLessons(lessonid, courses.Id, User.TenantId)
	if err != nil {
		fmt.Println(err)
	}

	var count = 0

	var title, data, Lesson string

	if lessonid == 0 {

		for _, val := range lessonlist {

			if count == 0 {

				lesson = val
			}

			count++
		}

		title = lesson.Title

		data = string(lesson.Content)

		Lesson = lesson.LessonType

	} else if lessonid != 0 {

		title = lesson.Title

		data = string(lesson.Content)

		Lesson = lesson.LessonType
	}

	switch Lesson {
	case "embed":

		Lesson = "AddEmbed"

	case "text":

		Lesson = "AddText"

	case "quiz":

		Lesson = "AddQuizs"

	case "file":

		Lesson = "AddFile"

	}

	Template := "courses"

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/themes/"+Template+"/layouts/_default/content_detail.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	controllers.RenderTemplate(c, tmpl, "content_detail.html", gin.H{"template_name": template_name, "sectionlist": sectionlist, "lessonlist": lessonlist, "lesson": lesson, "data": data, "coursename": courses.Title, "uuid": uuid, "Lesson": Lesson, "title": title})

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

	membershipLevels, _, _ := controllers.MembershipConfig.MembershipLevelsList(0, 10, membership.Filter{}, user.TenantId)

	templatedetails, err := controllers.MenuConfig.GetTemplateById(user.GoTemplateDefault, user.TenantId)
	if err != nil {
		fmt.Println(err)
	}

	templateName := GetTemplateName(c, templatedetails.TemplateName)

	tmpl, err := template.ParseFiles("websites/themes/"+templateName+"/layouts/partials/header.html", "websites/themes/"+templateName+"/layouts/partials/footer.html", "websites/themes/"+templateName+"/layouts/partials/head.html", "websites/common/layouts/membership.html")
	if err != nil {
		fmt.Println(err, "templateerr")
	}

	controllers.RenderTemplate(c, tmpl, "membership.html", gin.H{"menulist": newmenulist, "membership": membershipLevels, "template_name": c.Query("template_name")})
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

	templatedetails, err := controllers.MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	Template := GetTemplateName(c, templatedetails.TemplateName)

	membershipid, _ := strconv.Atoi(c.Param("id"))

	fmt.Println(membershipid, "membeershipid")

	GetEditMembership, err := controllers.MembershipConfig.MembershiplevelEdit(membershipid, User.TenantId)

	if err != nil {
		log.Fatal("Edit membership level error", err)
		// c.AbortWithError(500, err)
	}

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/common/layouts/checkout.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	controllers.RenderTemplate(c, tmpl, "checkout.html", gin.H{"menulist": newmenulist, "template_name": c.Query("template_name"), "membershipdetail": GetEditMembership})

}

func GetMenuItemsListByTenantID(tenantId string, websiteid int) ([]menu.TblMenus, error) {

	menulist, _, err := controllers.MenuConfig.MenuList(100, 0, menu.Filter{}, tenantId, websiteid)
	if err != nil {
		return nil, err
	}

	var parentid int
	for _, val := range menulist {
		parentid = val.Id
	}

	menuitemslist, err := controllers.MenuConfig.GetMenusByParentid(parentid, tenantId)
	if err != nil {
		return nil, err
	}

	var newmenulist []menu.TblMenus
	for _, val := range menuitemslist {
		if val.ParentId != 0 {
			newmenulist = append(newmenulist, val)
		}
	}
	return newmenulist, nil
}
func GetTemplateName(c *gin.Context, TemplateName string) string {

	templateName := c.Query("template_name")
	if templateName == "" {
		templateName = TemplateName
	}
	templateName = strings.ReplaceAll(strings.ToLower(templateName), " ", "_")

	return templateName
}
func GetTenantByHost(c *gin.Context) (team.TblUser, menu.TblWebsite, error) {

	var website menu.TblWebsite

	if os.Getenv("TENANTID") == "1" {

		website, _ = controllers.MenuConfig.GetWebsiteById(1, "1")
	}
	if website.Id == 0 {
		c.HTML(404, "nodata.html", nil)
		c.Abort()
		return team.TblUser{}, menu.TblWebsite{}, fmt.Errorf("website not found")
	}

	User, _, _ := controllers.NewTeamWP.GetUserById(website.CreatedBy, []int{})
	User.GoTemplateDefault = website.TemplateId

	return User, website, nil
}

func AllEntryList(tenatid string, webid int) ([]chn.Tblchannelentries, error) {

	fmt.Println("checkthisssss")

	menulist, _ := GetMenuItemsListByTenantID(tenatid, webid)

	var entries []chn.Tblchannelentries

	for _, val := range menulist {

		entries, _, _, _ = controllers.ChannelConfigWP.ChannelEntriesList(chn.Entries{ChannelId: val.TypeId, Status: "Published", Limit: 100, Offset: 0}, tenatid)

		for i := range entries {

			entries[i].CreatedDate = entries[i].CreatedOn.In(controllers.TZONE).Format("Jan 2 2006")

			if entries[i].ProfileImagePath != "" {
				entries[i].ProfileImagePath = "/image-resize?name=" + entries[i].ProfileImagePath
			}
			categoryid, _ := strconv.Atoi(entries[i].CategoriesId)
			categorydetails, _ := controllers.CategoryConfig.GetSubCategoryDetails(categoryid, tenatid)
			entries[i].CategoryGroup = categorydetails.CategoryName
			entries[i].Description = StripHTML(entries[i].Description)

		}
	}

	return entries, nil
}

func StaticPageData(c *gin.Context) {

	// template_name := c.Query("template_name")

	pagename := c.Param("pagename")

	User, website, err := GetTenantByHost(c)
	if err != nil {
		fmt.Println(err)
	}

	templatedetails, err := controllers.MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	if err != nil {
		fmt.Println(err)
	}

	template_name := strings.ToLower(templatedetails.TemplateName)
	template_name = strings.ReplaceAll(template_name, " ", "_")
	// templatedetails, _ := controllers.MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	newmenulist, _ := GetMenuItemsListByTenantID(User.TenantId, website.Id)

	pagedetails, err := controllers.MenuConfig.GetPageBySlug(pagename, User.TenantId)
	if err != nil || pagedetails.Id == 0 {

		c.HTML(404, "nodata.html", nil)

		return

	}

	seodetail, err := controllers.MenuConfig.SeoDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	seodetail.PageKeyword = pagedetails.MetaKeywords

	seodetail.PageTitle = pagedetails.MetaTitle

	seodetail.PageDescription = pagedetails.MetaDescription

	settingsdetail, err := controllers.MenuConfig.SettingsDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}
	UserDetailsFunction(c)

	memberdet, _ := c.Get("userdetails")

	Template := GetTemplateName(c, templatedetails.TemplateName)

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/themes/"+Template+"/layouts/_default/content_detail.html", "websites/themes/"+Template+"/layouts/_default/static_page.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}
	content := template.HTML(pagedetails.PageDescription)

	entrylist, _ := AllEntryList(User.TenantId, website.Id)

	controllers.RenderTemplate(c, tmpl, "static_page.html", gin.H{"PageDescription": content, "seodetail": seodetail, "searchlist": entrylist, "menulist": newmenulist, "template_name": template_name, "settingsdetail": settingsdetail, "memberdet": memberdet})

}

func StaticEntryData(c *gin.Context) {

	slug := c.Param("entryname")

	User, website, err := GetTenantByHost(c)
	if err != nil {
		fmt.Println(err)
	}

	templatedetails, err := controllers.MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	if err != nil {
		fmt.Println(err)
	}

	template_name := strings.ToLower(templatedetails.TemplateName)
	template_name = strings.ReplaceAll(template_name, " ", "_")

	newmenulist, _ := GetMenuItemsListByTenantID(User.TenantId, website.Id)

	channelEntry, _, err := controllers.ChannelConfigWP.FetchChannelEntryDetail(chn.EntriesInputs{Slug: slug, TenantId: User.TenantId, GetAuthorDetails: true, GetLinkedCategories: true}, nil)

	if err != nil || channelEntry.Id == 0 {

		c.HTML(404, "nodata.html", nil)

		return

	}

	UserDetailsFunction(c)

	memberdet, _ := c.Get("userdetails")

	Template := GetTemplateName(c, templatedetails.TemplateName)

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/themes/"+Template+"/layouts/_default/content_detail.html", "websites/themes/"+Template+"/layouts/_default/static_page.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}
	content := template.HTML(channelEntry.Description)

	entrylist, _ := AllEntryList(User.TenantId, website.Id)

	seodetail, err := controllers.MenuConfig.SeoDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	seodetail.PageKeyword = channelEntry.Keyword

	seodetail.PageTitle = channelEntry.MetaTitle

	seodetail.PageDescription = channelEntry.Description

	settingsdetail, err := controllers.MenuConfig.SettingsDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	controllers.RenderTemplate(c, tmpl, "static_page.html", gin.H{"PageDescription": content, "seodetail": seodetail, "searchlist": entrylist, "menulist": newmenulist, "template_name": template_name, "settingsdetail": settingsdetail, "memberdet": memberdet})

}
