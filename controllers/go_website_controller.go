package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	chn "github.com/spurtcms/channels"
	listing "github.com/spurtcms/listing"
	menu "github.com/spurtcms/menu"
	csrf "github.com/utrack/gin-csrf"
	"gopkg.in/yaml.v2"
)

type TemplateInfo struct {
	TemplateName  string `yaml:"template_name"`
	TemplateImage string `yaml:"template_image"`
	Description   string `yaml:"description"`
	Type          string `yaml:"type"`
	Channel       string `yaml:"channel"`
	HtmlTemplates string `yaml:"html_templates"`
}

func WebsiteList(c *gin.Context) {

	var (
		limt   int
		offset int
		filter menu.Filter
	)

	limit := c.Query("limit")

	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")
	filter.Status = strings.Trim(c.DefaultQuery("status", ""), " ")
	filter.ToDate = strings.Trim(c.DefaultQuery("lastupdated", ""), " ")
	var filterflag bool
	if filter.Keyword != "" {
		filterflag = true
	} else {
		filterflag = false
	}

	fmt.Println(filterflag)

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}
	permisison, perr := NewAuth.IsGranted("Menu", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("menu authorization error: %s", perr)
	}
	if !permisison {
		ErrorLog.Printf("Menu authorization error")
		c.Redirect(301, "/403-page")
		return
	}
	channelslist, _, err := ChannelConfig.ListChannel(chn.Channels{Limit: 100, Offset: 0, IsActive: false, TenantId: TenantId})

	if err != nil {

		fmt.Println(err)
	}

	goTemplateList, err := ReadYamlFile("websites/themes")
	if err != nil {
		fmt.Println("Failed to read YAML files:", err)

	}

	host := c.Request.Host
	isLocal := false

	if strings.Contains(host, "localhost") {
		isLocal = true
	}

	baseurl := os.Getenv("BASE_URL")
	baseurl = strings.TrimPrefix(baseurl, "https://")
	baseurl = strings.TrimPrefix(baseurl, "http://")
	baseurl = strings.TrimSuffix(baseurl, "/")

	var FinalwebsiteList []menu.TblWebsite
	websitelist, count, err := MenuConfig.WebsiteList(limt, offset, menu.Filter{Keyword: filter.Keyword}, TenantId)

	for _, val := range websitelist {
		val.CreatedDate = val.CreatedOn.In(TZONE).Format(Datelayout)
		if !val.ModifiedOn.IsZero() {
			val.DateString = val.ModifiedOn.In(TZONE).Format(Datelayout)
		} else {
			val.DateString = val.CreatedOn.In(TZONE).Format(Datelayout)
		}

		if strings.Contains(host, "localhost") {
			val.Subdomain = "http://" + val.Name + "." + baseurl
		} else {
			val.Subdomain = "https://" + val.Name + "." + baseurl
		}
		FinalwebsiteList = append(FinalwebsiteList, val)

	}

	go_template_default, _ := c.Get("go_template_default")
	webbanner, _ := c.Cookie("webbanner")
	if webbanner == "" {
		webbanner = "true"
	}
	webbanner = "true"
	ModuleName, _, _ := ModuleRouteName(c)
	translate, _ := TranslateHandler(c)

	//pagination calc
	paginationendcount := len(websitelist) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(pageno, int(count), limt)

	c.HTML(200, "websitelist.html", gin.H{"Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3},
		"csrf":                csrf.GetToken(c),
		"webbanner":           webbanner,
		"Menu":                NewMenuController(c),
		"linktitle":           "Website",
		"translate":           translate,
		"title":               ModuleName,
		"baseurl":             baseurl,
		"go_template_default": go_template_default,
		"channelslist":        channelslist,
		"templatelist":        goTemplateList,
		"count":               count,
		"websitelist":         FinalwebsiteList,
		"isLocal":             isLocal,
		"baseurlpath":         os.Getenv("BASE_URL"),
		"Count":               count, "Limit": limt, "Paginationendcount": paginationendcount, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Filter": filter, "Paginationstartcount": paginationstartcount,
	})
}

func BrowseTheme(c *gin.Context) {

	var (
		limt   int
		offset int
		filter menu.Filter
	)

	limit := c.Query("limit")

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")
	filter.Status = strings.Trim(c.DefaultQuery("status", ""), " ")
	filter.ToDate = strings.Trim(c.DefaultQuery("lastupdated", ""), " ")
	var filterflag bool
	if filter.Keyword != "" {
		filterflag = true
	} else {
		filterflag = false
	}

	fmt.Println(filterflag)

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	goTemplateList, err := ReadYamlFile("websites/themes")
	if err != nil {
		fmt.Println("Failed to read YAML files:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// get pagination theme list

	start := (pageno - 1) * limt
	end := start + limt
	total := len(goTemplateList)

	if start > total {
		start = total
	}
	if end > total {
		end = total
	}

	paginatedList := goTemplateList[start:end]

	webbanner, _ := c.Cookie("webbanner")
	if webbanner == "" {
		webbanner = "true"
	}

	ModuleName, _, _ := ModuleRouteName(c)
	translate, _ := TranslateHandler(c)

	// Get website_id from query param
	websiteId, _ := strconv.Atoi(c.Query("website_id"))

	Dataid := os.Getenv("WEBSITE_DATA_ID")
	websiteId, _ = strconv.Atoi(Dataid) // Fetch single website by ID — only if websiteId is valid

	var websiteInfo menu.TblWebsite
	if websiteId != 0 {
		websiteInfo, err = MenuConfig.GetWebsiteById(websiteId, TenantId)
		if err != nil {
			fmt.Println("Failed to get website:", err)
		}
	}

	// get template details

	templatedetails, _ := MenuConfig.GetTemplateById(websiteInfo.TemplateId, TenantId)

	// Find matching template detail
	var templateDetail TemplateInfo
	for _, t := range goTemplateList {
		if t.TemplateName == websiteInfo.TemplateName {
			templateDetail = t
			break
		}
	}

	// fallback to first template if no match
	if templateDetail.TemplateName == "" && len(goTemplateList) > 0 {
		templateDetail = goTemplateList[0]
	}
	var FinalwebsiteList []menu.TblWebsite
	websitelist, _, err := MenuConfig.WebsiteList(limt, offset, menu.Filter{Keyword: filter.Keyword}, TenantId)
	UserDetails, _, _ := NewTeam.GetUserById(websiteInfo.CreatedBy, []int{})
	baseurl := os.Getenv("BASE_URL")
	baseurl = strings.TrimPrefix(baseurl, "https://")
	baseurl = strings.TrimPrefix(baseurl, "http://")
	baseurl = strings.TrimSuffix(baseurl, "/")
	host := c.Request.Host
	isLocal := false
	var data menu.TblWebsite

	for _, val := range websitelist {
		val.CreatedDate = val.CreatedOn.In(TZONE).Format(Datelayout)
		if !val.ModifiedOn.IsZero() {
			val.DateString = val.ModifiedOn.In(TZONE).Format(Datelayout)
		} else {
			val.DateString = val.CreatedOn.In(TZONE).Format(Datelayout)
		}
		if strings.Contains(host, "localhost") {
			val.Subdomain = "http://" + val.Name + "." + baseurl
		} else {
			val.Subdomain = "https://" + val.Name + "." + baseurl
		}
		FinalwebsiteList = append(FinalwebsiteList, val)
		data = val
	}

	//pagination calc

	paginationendcount := end
	paginationstartcount := start
	Previous, Next, PageCount, Page := Pagination(pageno, len(goTemplateList), limt)

	Allwidget, _, err := MenuConfig.GetWidgetList(limt, offset, menu.Filter{Keyword: filter.Keyword, Status: filter.Status, ToDate: filter.ToDate}, TenantId, 1)

	websitdataid := os.Getenv("WEBSITE_DATA_ID")
	url := os.Getenv("BASE_URL")

	c.HTML(200, "browsetheme.html", gin.H{"Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3,
	},
		"Menu":           NewMenuController(c),
		"translate":      translate,
		"linktitle":      "Browse theme",
		"title":          ModuleName,
		"csrf":           csrf.GetToken(c),
		"templatelist":   paginatedList,
		"templatedetail": templatedetails,
		"website":        websiteInfo,
		"webbanner":      webbanner,
		"isLocal":        isLocal,
		"websitelist":    FinalwebsiteList,
		"web":            data,
		"websitedataid":  websitdataid,
		"Allwidgets":     Allwidget,
		"UserDetails":    UserDetails,
		"Url":            url,
		"Count":          len(goTemplateList), "Paginationendcount": paginationendcount, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Filter": filter, "Paginationstartcount": paginationstartcount, "Limit": limt,
	})
}
func Configuration(c *gin.Context) {

	tempid := c.Param("template")
	templateid, _ := strconv.Atoi(tempid)
	template, _ := MenuConfig.GetTemplateById(templateid, TenantId)
	templateName := template.TemplateName
	if templateName == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var (
		limt   int
		offset int
		filter menu.Filter
	)

	limit := c.Query("limit")

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")
	filter.Status = strings.Trim(c.DefaultQuery("status", ""), " ")
	filter.ToDate = strings.Trim(c.DefaultQuery("lastupdated", ""), " ")
	var filterflag bool
	if filter.Keyword != "" {
		filterflag = true
	} else {
		filterflag = false
	}

	fmt.Println(filterflag)

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	websiteId, _ := strconv.Atoi(c.Query("website_id"))

	websiteInfo, err := MenuConfig.GetWebsiteById(websiteId, TenantId)

	if err != nil {
		fmt.Println("Failed to get website:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	webbanner, _ := c.Cookie("webbanner")
	if webbanner == "" {
		webbanner = "true"
	}

	ModuleName, _, _ := ModuleRouteName(c)
	translate, _ := TranslateHandler(c)

	var FinalwebsiteList []menu.TblWebsite
	websitelist, _, err := MenuConfig.WebsiteList(limt, offset, menu.Filter{Keyword: filter.Keyword}, TenantId)

	baseurl := os.Getenv("BASE_URL")
	baseurl = strings.TrimPrefix(baseurl, "https://")
	baseurl = strings.TrimPrefix(baseurl, "http://")
	baseurl = strings.TrimSuffix(baseurl, "/")
	host := c.Request.Host
	isLocal := false
	var data menu.TblWebsite

	for _, val := range websitelist {
		val.CreatedDate = val.CreatedOn.In(TZONE).Format(Datelayout)
		if !val.ModifiedOn.IsZero() {
			val.DateString = val.ModifiedOn.In(TZONE).Format(Datelayout)
		} else {
			val.DateString = val.CreatedOn.In(TZONE).Format(Datelayout)
		}
		if strings.Contains(host, "localhost") {
			val.Subdomain = "http://" + val.Name + "." + baseurl
		} else {
			val.Subdomain = "https://" + val.Name + "." + baseurl
		}
		FinalwebsiteList = append(FinalwebsiteList, val)
		data = val
	}
	data.TemplateName = template.TemplateName

	var Allwidgetlist []menu.TblWidgets

	Allwidget, totalwidgetCount, err := MenuConfig.GetWidgetList(limt, offset, menu.Filter{Keyword: filter.Keyword, Status: filter.Status, ToDate: filter.ToDate}, TenantId, templateid)
	for _, val := range Allwidget {
		val.CreatedDate = val.CreatedOn.In(TZONE).Format(Datelayout)
		if !val.ModifiedOn.IsZero() {
			val.CreatedDate = val.ModifiedOn.In(TZONE).Format(Datelayout)
		} else {
			val.CreatedDate = val.CreatedOn.In(TZONE).Format(Datelayout)
		}
		Allwidgetlist = append(Allwidgetlist, val)
	}

	channelist, _, clerr := ChannelConfig.ListChannel(chn.Channels{Limit: 100, Offset: 0, IsActive: true, TenantId: TenantId})
	if clerr != nil {
		ErrorLog.Printf("channellist error :%s", clerr)
	}

	Categorylist, _ := CategoryConfig.AllCategoriesWithSubList(TenantId)

	Entrylist, _, _, _ := ChannelConfig.ChannelEntriesList(chn.Entries{Status: "Published"}, TenantId)

	pagelist, _, err := MenuConfig.GetTemplatePageList(100, 0, menu.Filter{Status: "Active"}, TenantId, 0)

	listinglist, _, err := ListingConfig.ListingsList(100, 0, listing.Filter{}, TenantId)
	paginationendcount := len(Allwidgetlist) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(pageno, int(totalwidgetCount), limt)

	UserDetails, _, _ := NewTeam.GetUserById(websiteInfo.CreatedBy, []int{})

	c.HTML(200, "configure.html", gin.H{
		"hideCancelBtn":  true,
		"Menu":           NewMenuController(c),
		"translate":      translate,
		"linktitle":      "Widgets",
		"title":          ModuleName,
		"csrf":           csrf.GetToken(c),
		"templatedetail": template,
		"webbanner":      webbanner,
		"website":        websiteInfo,
		"isLocal":        isLocal,
		"UserDetails":    UserDetails,
		"websitelist":    FinalwebsiteList,
		"web":            data,
		"Allwidget":      Allwidgetlist,
		"templateid":     templateid,
		"Categorylist":   Categorylist, "channelist": channelist, "listinglist": listinglist, "Entrylist": Entrylist, "pagelist": pagelist,
		"Filter":     filter,
		"tabactive":  "widgets",
		"Searchtrue": filterflag,
		"Count":  totalwidgetCount,
		"Limit":  limt,
		"Pagination": PaginationData{
			NextPage:     pageno + 1,
			PreviousPage: pageno - 1,
			TotalPages:   PageCount,
			TwoAfter:     pageno + 2,
			TwoBelow:     pageno - 2,
			ThreeAfter:   pageno + 3,
		},
		"Paginationendcount":  paginationendcount,
		"Paginationstartcount": paginationstartcount,
		"Previous":            Previous,
		"Next":                Next,
		"PageCount":           PageCount,
		"CurrentPage":         pageno,
		"Page":                Page,
	})
}

func ReadYamlFile(path string) ([]TemplateInfo, error) {
	var allTemplateInfos []TemplateInfo

	themes, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, theme := range themes {
		if !theme.IsDir() {
			continue
		}
		themeDir := filepath.Join(path, theme.Name())

		files, err := ioutil.ReadDir(themeDir)
		if err != nil {
			log.Println("Error reading theme dir:", themeDir, err)
			continue
		}
		for _, file := range files {
			if file.IsDir() || filepath.Ext(file.Name()) != ".yaml" {
				continue
			}
			data, err := ioutil.ReadFile(filepath.Join(themeDir, file.Name()))
			if err != nil {
				log.Println("Error reading file:", file.Name(), err)
				continue
			}
			var TemplateInfo TemplateInfo
			if err := yaml.Unmarshal(data, &TemplateInfo); err != nil {
				log.Println("Error parsing YAML:", file.Name(), err)
				continue
			}
			allTemplateInfos = append(allTemplateInfos, TemplateInfo)
		}
	}
	return allTemplateInfos, nil
}

func ReadYamlFileForWebsite(path string, templateName string) (TemplateInfo, error) {
	var templateInfo TemplateInfo

	themeDir := filepath.Join(path, templateName)

	entries, err := os.ReadDir(themeDir)
	if err != nil {
		return templateInfo, err
	}

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".yaml" {
			continue
		}

		data, err := os.ReadFile(filepath.Join(themeDir, entry.Name()))
		if err != nil {
			return templateInfo, err
		}
		if err := yaml.Unmarshal(data, &templateInfo); err != nil {
			return templateInfo, err
		}

		// Return on first valid YAML file
		return templateInfo, nil
	}

	return templateInfo, fmt.Errorf("no yaml file found in %s", themeDir)
}

//Create Website Function//

func CreateWebsite(c *gin.Context) {

	sitename := c.PostForm("site_name")
	channelname := c.PostForm("channel_name")
	template_name := c.PostForm("template_name")
	template_desc := c.PostForm("template_desc")
	template_image := c.PostForm("template_img")
	template_channel := c.PostForm("template_type")
	// templateid, _ := strconv.Atoi(c.PostForm("template_id"))
	userid := c.GetInt("userid")
	fmt.Println(sitename, channelname, template_name, "checkdata")

	temp := menu.TblGoTemplates{

		TemplateName:        template_name,
		TemplateDescription: template_desc,
		TemplateImage:       template_image,
		IsDeleted:           0,
		TenantId:            TenantId,
		CreatedBy:           userid,
		ChannelSlugName:     template_channel,
	}

	templateinfo, _ := MenuConfig.CreateTemplate(temp)

	websiteinfo := menu.TblWebsite{
		Name:         sitename,
		ChannelNames: channelname,
		TemplateId:   templateinfo.Id,
		CreatedBy:    userid,
		TenantId:     TenantId,
		IsDeleted:    0,
		Status:       0,
	}
	web, err1 := MenuConfig.CreateWebsite(websiteinfo)
	if err1 != nil {
		fmt.Println(err1)
	}
	_, errFooter := MenuConfig.CreateMenus(menu.MenuCreate{
		MenuName:    "Footer",
		Description: "The menu at the bottom of the webpage, commonly listing support, privacy, and contact pages",
		MenuSlug:    "footer",
		Status:      1,
		ParentId:    0,
		TenantId:    TenantId,
		CreatedBy:   userid,
		WebsiteId:   web.Id,
	})
	if errFooter != nil {
		fmt.Println(errFooter)
	}

	_, errHeader := MenuConfig.CreateMenus(menu.MenuCreate{
		MenuName:    "Headers",
		Description: "The primary menu at the top of a webpage, guiding users to major site sections",
		MenuSlug:    "headers",
		Status:      1,
		ParentId:    0,
		TenantId:    TenantId,
		CreatedBy:   userid,
		WebsiteId:   web.Id,
	})
	if errHeader != nil {
		fmt.Println(errHeader)
	}
	_, erraside := MenuConfig.CreateMenus(menu.MenuCreate{
		MenuName:    "Aside",
		Description: "The primary menu at the left of a webpage, guiding users to major site sections",
		MenuSlug:    "aside",
		Status:      1,
		ParentId:    0,
		TenantId:    TenantId,
		CreatedBy:   userid,
		WebsiteId:   web.Id,
	})
	if erraside != nil {
		fmt.Println(errHeader)
	}
	c.SetCookie("get-toast", "Website Created Successfully", 3600, "", "", false, false)
	c.Redirect(301, "/admin/website/")
}

func UpdateWebsite(c *gin.Context) {
	webid, _ := strconv.Atoi(c.PostForm("website_id"))
	sitename := c.PostForm("site_name")
	channelname := c.PostForm("channel_name")
	template_id, _ := strconv.Atoi(c.PostForm("template_id"))
	template_name := c.PostForm("template_name")
	template_desc := c.PostForm("template_desc")
	template_image := c.PostForm("template_img")
	template_channel := c.PostForm("template_type")
	userid := c.GetInt("userid")
	var templateid int
	temp := menu.TblGoTemplates{

		TemplateName:        template_name,
		TemplateDescription: template_desc,
		TemplateImage:       template_image,
		IsDeleted:           0,
		TenantId:            TenantId,
		CreatedBy:           userid,
		ChannelSlugName:     template_channel,
	}

	if template_id == 0 {

		templateinfo, _ := MenuConfig.CreateTemplate(temp)
		templateid = templateinfo.Id
	} else {

		templateid = template_id
	}

	websiteinfo := menu.TblWebsite{
		Id:           webid,
		Name:         sitename,
		ChannelNames: channelname,
		TemplateId:   templateid,
		ModifiedBy:   userid,
		TenantId:     TenantId,
	}
	_, err1 := MenuConfig.UpdateWebsite(websiteinfo)

	if err1 != nil {

		fmt.Println(err1)
	}
	c.SetCookie("get-toast", "Theme Updated Successfully", 3600, "", "", false, false)
	c.Redirect(301, "/admin/browsetheme/")

}

func DeleteWebsite(c *gin.Context) {

	websiteid, _ := strconv.Atoi(c.Param("id"))

	pageno := c.Query("page")

	var url string

	if pageno != "" {
		url = "/admin/website/?page=" + pageno
	} else {
		url = "/admin/website/"
	}

	userid := c.GetInt("userid")

	permisison, perr := NewAuth.IsGranted("Menu", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("Delete authorization error: %s", perr)
	}

	if !permisison {
		c.Redirect(301, "/403-page")
		return
	}

	err := MenuConfig.DeleteWebsite(websiteid, userid, TenantId)
	if err != nil {
		ErrorLog.Printf("Delete error: %s", perr)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		return
	}

	c.SetCookie("get-toast", "Website Deleted Successfully", 3600, "", "", false, false)
	c.Redirect(http.StatusMovedPermanently, url)
}
func CheckSiteName(c *gin.Context) {

	subdomain := c.PostForm("sitename")

	id, _ := strconv.Atoi(c.PostForm("webid"))

	err := MenuConfig.CheckSiteName(subdomain, id)

	if err != nil {

		c.JSON(200, false)

		return
	}

	c.JSON(200, true)
}
