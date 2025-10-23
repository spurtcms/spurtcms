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

	// goTemplateList, _, err := MenuConfig.GoTemplatesList("")
	// if err != nil {
	// 	fmt.Println(err)
	// }
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

	fmt.Println(goTemplateList, "templatelist")

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

//Create Website Function//

func CreateWebsite(c *gin.Context) {

	sitename := c.PostForm("site_name")
	channelname := c.PostForm("channel_name")
	template_name := c.PostForm("template_name")
	template_desc := c.PostForm("template_desc")
	template_image := c.PostForm("template_img")
	template_channel := c.PostForm("template_type")
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

	_, err1 := MenuConfig.CreateWebsite(websiteinfo)

	if err1 != nil {

		fmt.Println(err1)
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
	fmt.Println(sitename, channelname, "checkdata")
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
	c.SetCookie("get-toast", "Website Updated Successfully", 3600, "", "", false, false)
	c.Redirect(301, "/admin/website/")

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
