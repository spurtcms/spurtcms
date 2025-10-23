package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"spurt-cms/models"
	storagecontroller "spurt-cms/storage-controller"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	"github.com/spurtcms/blocks"
	menu "github.com/spurtcms/menu"
	csrf "github.com/utrack/gin-csrf"
)

func GoTemplates(c *gin.Context) {
	permisison, perr := NewAuth.IsGranted("Menu", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("menu authorization error: %s", perr)
	}
	if !permisison {
		ErrorLog.Printf("Menu authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	var (
		goTemplateList []menu.TblGoTemplates
		count          int
		err            error
		count2         int64
	)
	goTemplateList, count1, err := MenuConfig.GoTemplatesList(TenantId)
	if err != nil {
		fmt.Println(err)
	}
	count = int(count1)
	if count == 0 {
		goTemplateList, count2, err = MenuConfig.GoTemplatesList("")
		if err != nil {
			fmt.Println(err)
		}
		count = int(count2)
	}

	var modulenames []string
	moduleNameSet := make(map[string]struct{})

	for _, tpl := range goTemplateList {
		if tpl.TemplateModuleName != "" {
			if _, exists := moduleNameSet[tpl.TemplateModuleName]; !exists {
				modulenames = append(modulenames, tpl.TemplateModuleName)
				moduleNameSet[tpl.TemplateModuleName] = struct{}{}
			}
		}
	}

	go_template_default, _ := c.Get("go_template_default")
	webbanner, _ := c.Cookie("webbanner")
	if webbanner == "" {
		webbanner = "true"
	}
	ModuleName, _, _ := ModuleRouteName(c)
	translate, _ := TranslateHandler(c)

	c.HTML(200, "go_template.html", gin.H{
		"csrf":                csrf.GetToken(c),
		"modulenames":         modulenames, // always sorted!
		"webbanner":           webbanner,
		"Menu":                NewMenuController(c),
		"linktitle":           "Website",
		"translate":           translate,
		"title":               ModuleName,
		"goTemplateList":      goTemplateList,
		"count":               count,
		"go_template_default": go_template_default,
	})
}

func GoTemplateUpdate(c *gin.Context) {

	templateId, _ := strconv.Atoi(c.Param("id"))

	permisison, perr := NewAuth.IsGranted("Menu", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("menu authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Menu authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		userid := c.GetInt("userid")

		err := NewTeam.UpdateGoTemplate(templateId, userid, TenantId)
		if err != nil {
			fmt.Println(err)
		}

		c.SetCookie("get-toast", "Template Updated Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, "/admin/website/")
	}
}

//Template edit Page//

func TemplateEditPage(c *gin.Context) {
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

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}
	permisison, perr := NewAuth.IsGranted("Website", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("website authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("website authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	webid, _ := strconv.Atoi(c.Param("id"))

	Websitedet, _ := MenuConfig.GetWebsiteById(webid, TenantId)

	templatedetail, err := MenuConfig.GetTemplateById(Websitedet.TemplateId, TenantId)

	if err != nil {

		fmt.Println(err)
	}

	templatedetail.DateString = templatedetail.CreatedOn.In(TZONE).Format(Datelayout)

	templatepagelist, count, err := MenuConfig.GetTemplatePageList(limt, offset, filter, TenantId, webid)
	var pagelist []menu.TblTemplatePages

	for _, page := range templatepagelist {

		page.CreatedDate = page.CreatedOn.In(TZONE).Format(Datelayout)
		if !page.ModifiedOn.IsZero() {
			page.ModifiedDate = page.ModifiedOn.In(TZONE).Format(Datelayout)
		} else {
			page.ModifiedDate = page.CreatedOn.In(TZONE).Format(Datelayout)
		}

		pagelist = append(pagelist, page)
	}

	if err != nil {

		fmt.Println(err)
	}

	paginationendcount := len(pagelist) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(pageno, int(count), limt)

	ModuleName, _, _ := ModuleRouteName(c)

	translate, _ := TranslateHandler(c)

	c.HTML(200, "go_template_edit.html", gin.H{"Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3}, "csrf": csrf.GetToken(c), "Searchtrue": filterflag, "Count": count, "Limit": limt, "templatepagelist": pagelist, "templatedetail": templatedetail, "websitedetail": Websitedet, "Menu": NewMenuController(c), "linktitle": "Website Page", "translate": translate, "title": ModuleName, "Paginationendcount": paginationendcount, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Filter": filter, "Paginationstartcount": paginationstartcount})

}

// Template Add Page//
func AddPageInWebsite(c *gin.Context) {

	permisison, perr := NewAuth.IsGranted("Website", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("menu authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Menu authorization error")
		c.Redirect(301, "/403-page")
		return
	}
	var editpage bool
	var mode string
	pageid, _ := strconv.Atoi(c.Param("id"))

	if pageid != 0 {

		editpage = true

		mode = "edit"
	} else {

		mode = "create"
	}
	webid, _ := strconv.Atoi(c.Query("webid"))

	PageDetail, _ := MenuConfig.GetPageById(pageid, TenantId)

	fmt.Println(PageDetail.PageDescription, "descriptionnnn")

	Websitedet, _ := MenuConfig.GetWebsiteById(webid, TenantId)

	templatedetail, err := MenuConfig.GetTemplateById(Websitedet.TemplateId, TenantId)

	if err != nil {

		fmt.Println(err)
	}

	channelist, clerr := ChannelConfig.GetPermissionChannel(TenantId)
	if clerr != nil {
		ErrorLog.Printf("create Entry listchannel error: %s", clerr)
	}

	templatedetail.DateString = templatedetail.CreatedOn.In(TZONE).Format(Datelayout)

	var bytedata1 []byte

	var bytedata2 []byte

	var finalblocklist []blocks.TblBlock

	blocklist, _, err := BlockConfig.BlockList(0, 0, blocks.Filter{}, TenantId)

	if err != nil {
		fmt.Println("collection list", err)
	}

	for _, val := range blocklist {
		var first = val.FirstName
		var last = val.LastName
		var firstn = strings.ToUpper(first[:1])
		var lastn string
		if val.LastName != "" {
			lastn = strings.ToUpper(last[:1])
		}

		val.ChannelNames = strings.Split(val.ChannelSlugname, ",")

		var Name = firstn + lastn
		val.NameString = Name

		tagname := strings.Split(val.TagValue, ",")

		val.TagValueArr = append(val.TagValueArr, tagname...)
		img := val.CoverImage
		imgcontain := "/image-resize?name="
		flag := strings.Contains(img, imgcontain)
		if !flag {
			val.CoverImage = "/image-resize?name=" + val.CoverImage
		}
		if val.ProfileImagePath != "" {
			userimg := val.ProfileImagePath
			imgflag := strings.Contains(userimg, imgcontain)
			if !imgflag {
				val.ProfileImagePath = "/image-resize?name=" + val.ProfileImagePath
			}
		}

		finalblocklist = append(finalblocklist, val)

	}

	data := map[string]interface{}{"data": finalblocklist}

	bytedata1, _ = json.Marshal(data)

	if permisison {

		endurl := os.Getenv("MASTER_BLOCKS_ENDPOINTURL")
		req, err := http.NewRequest("GET", endurl, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request: " + err.Error()})
			return
		}

		query := req.URL.Query()

		req.URL.RawQuery = query.Encode()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		masterconnect := true

		if err != nil || resp.StatusCode != http.StatusOK {
			fmt.Println("Error connecting to master server:", err)
			masterconnect = false
		} else {
			defer resp.Body.Close()
		}

		var responseData ResponseData
		if masterconnect {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err == nil {
				fmt.Println("Error response:", err)
				resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
				err = json.NewDecoder(resp.Body).Decode(&responseData)
				if err != nil {
					masterconnect = false
				}
			} else {
				masterconnect = false
			}
		}

		if !masterconnect {
			responseData = ResponseData{
				DefaultList:    []models.TblMstrBlocks{},
				AllList:        []models.TblMstrBlocks{},
				FinalblockList: []models.TblMstrBlocks{},
				BlockCount:     0,
			}
		}

		permisison, perr := NewAuth.IsGranted("Blocks", auth.CRUD, TenantId)
		if perr != nil {
			ErrorLog.Printf("block collection list authorization error: %s", perr)
			c.Redirect(301, "/403-page")
			return
		}
		if !permisison {

			ErrorLog.Printf("Block authorization error: %s", perr)
			c.Redirect(301, "/403-page")
			return
		}

		blockBannerValue, _ := c.Cookie("blockbanner")
		if blockBannerValue == "" {
			blockBannerValue = "true"
		}

		// fmt.Println(responseData.FinalblockList,"checkdata")
		data := map[string]interface{}{"data": responseData.FinalblockList}

		bytedata2, _ = json.Marshal(data)
	}
	selectedtype, _ := GetSelectedType()

	baseurl := os.Getenv("BASE_URL")

	urlpath := map[string]interface{}{"path": baseurl + "uploadb64image", "payload": "imagedata", "success": map[string]interface{}{"imagepath": "imagepath", "imagename": "imagename"}}

	newpath := os.Getenv("BASE_URL")

	ubyte, _ := json.Marshal(urlpath)

	templatepagelist, _, err := MenuConfig.GetTemplatePageList(100, 0, menu.Filter{}, TenantId, webid)

	ModuleName, _, _ := ModuleRouteName(c)

	translate, _ := TranslateHandler(c)

	c.HTML(200, "go_template_addpage.html", gin.H{"csrf": csrf.GetToken(c), "channellist": channelist, "websitedetail": Websitedet, "pagelist": templatepagelist, "PageDetail": PageDetail, "editpage": editpage, "Storagepath": string(ubyte), "uploadurl": newpath, "StorageType": selectedtype.SelectedType, "blocks": string(bytedata1), "defaultblocks": string(bytedata2), "Mode": mode, "templatedetail": templatedetail, "Menu": NewMenuController(c), "linktitle": "Website", "translate": translate, "title": ModuleName})

}

//save page Name//

func SavePage(c *gin.Context) {

	permisison, perr := NewAuth.IsGranted("Website", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("website authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("website authorization error")
		c.Redirect(301, "/403-page")
		return
	}
	pagename := c.PostForm("page_name")
	userid := c.GetInt("userid")
	webid, _ := strconv.Atoi(c.PostForm("webid"))
	pageid, _ := strconv.Atoi(c.PostForm("pageid"))
	pagedata := c.PostForm("pagedata")
	var pagedet menu.TblTemplatePages
	var err error

	if pageid != 0 {

		pageinfo := menu.TblTemplatePages{
			Id:              pageid,
			Name:            pagename,
			TenantId:        TenantId,
			ModifiedBy:      userid,
			PageDescription: pagedata,
			WebsiteId:       webid,
		}
		pagedet, err = MenuConfig.EditTemplatePage(&pageinfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	} else {

		pageinfo := menu.TblTemplatePages{
			Name:      pagename,
			TenantId:  TenantId,
			CreatedBy: userid,
			IsDeleted: 0,
			Status:    1,
			WebsiteId: webid,
		}
		pagedet, err = MenuConfig.CreateTemplatePage(&pageinfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(200, gin.H{"page": pagedet})
}

//Delete Page

func DeletePage(c *gin.Context) {
	permisison, perr := NewAuth.IsGranted("Website", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("website authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("website authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	pageid, _ := strconv.Atoi(c.Param("id"))

	pageno := c.Query("page")
	userid := c.GetInt("userid")
	webid := c.Query("webid")

	var url string
	if pageno != "" {
		url = "/admin/website/pages/" + webid + "?page=" + pageno
	} else {
		url = "/admin/website/pages/" + webid

	}

	err := MenuConfig.DeletePage(pageid, userid, TenantId)

	if strings.Contains(fmt.Sprint(err), "given some values is empty") {
		ErrorLog.Printf("deletemenu mandatory field error: %s", perr)
		c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
		c.Redirect(301, url)
		return
	}

	if err != nil {
		ErrorLog.Printf("deletemenu error: %s", perr)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.Redirect(301, url)
		return
	}

	c.SetCookie("get-toast", "Page Deleted Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, url)

}

//Status Change

func PageStatusChange(c *gin.Context) {

	pageid, _ := strconv.Atoi(c.PostForm("id"))

	fmt.Println(pageid, "pageiddd")
	userid := c.GetInt("userid")
	val, _ := strconv.Atoi(c.Request.PostFormValue("isactive"))

	permisison, perr := NewAuth.IsGranted("Website", auth.Update, TenantId)
	if perr != nil {
		ErrorLog.Printf("delete menu authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Website authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	flg, err := MenuConfig.PageStatusChange(pageid, val, userid, TenantId)

	if err != nil {
		ErrorLog.Printf("menu status change error: %s", perr)
		json.NewEncoder(c.Writer).Encode(flg)

	} else {
		json.NewEncoder(c.Writer).Encode(flg)
	}

}

//SEO Page

func Seo(c *gin.Context) {

	webid, _ := strconv.Atoi(c.Query("webid"))

	seodetail, err := MenuConfig.SeoDetail(TenantId, webid)
	if err != nil {
		fmt.Println(err)
	}
	webbanner, _ := c.Cookie("webbanner")

	if webbanner == "" {

		webbanner = "true"
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

	var pages bool

	if webid != 0 {
		pages = true

	}
	Websitedet, _ := MenuConfig.GetWebsiteById(webid, TenantId)
	templatedetail, err := MenuConfig.GetTemplateById(Websitedet.TemplateId, TenantId)

	if err != nil {

		fmt.Println(err)
	}

	templatedetail.DateString = templatedetail.CreatedOn.In(TZONE).Format(Datelayout)

	ModuleName, _, _ := ModuleRouteName(c)

	translate, _ := TranslateHandler(c)

	c.HTML(200, "seo.html", gin.H{"csrf": csrf.GetToken(c), "websitedetail": Websitedet, "gotemplatepageheader": pages, "templatedetail": templatedetail, "webbanner": webbanner, "Menu": NewMenuController(c), "linktitle": "SEO", "translate": translate, "title": ModuleName, "seodetail": seodetail})

}

func SeoPage(c *gin.Context) {

	pagetitle := c.PostForm("pagetitle")

	pagedescription := c.PostForm("pagedescription")

	pagekeyword := c.PostForm("pagekeyword")

	storetitle := c.PostForm("storetitle")

	storedescription := c.PostForm("storedescription")

	storekeyword := c.PostForm("storekeyword")

	imagedata := c.PostForm("sitemapimage")

	webid := c.Query("webid")

	websiteid, _ := strconv.Atoi(c.Query("webid"))

	var url string

	if webid != "" && webid != "0" {

		url = "?webid=" + webid

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

	if permisison {

		var imageName, imagePath string

		var imageByte []byte

		tenantDetails, err := NewTeam.GetTenantDetails(TenantId)
		if err != nil {
			ErrorLog.Printf("error get storage type error: %s", err)
		}

		if imagedata != "" {
			imageName, imagePath, imageByte, _ = ConvertBase64toByte(imagedata, "seo")

			imagePath = tenantDetails.S3FolderName + imagePath

			uerr := storagecontroller.UploadCropImageS3(imageName, imagePath, imageByte)
			if uerr != nil {
				c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
				c.Redirect(301, "/admin/website/seo"+url)
				return
			}
		}

		seo := menu.TblGoTemplateSeo{
			PageTitle:        pagetitle,
			PageDescription:  pagedescription,
			PageKeyword:      pagekeyword,
			StoreTitle:       storetitle,
			StoreDescription: storedescription,
			StoreKeyword:     storekeyword,
			SiteMapName:      imageName,
			SiteMapPath:      imagePath,
			TenantId:         TenantId,
			WebsiteId:        websiteid,
		}

		if pagetitle != "" || storetitle != "" || imageName != "" {

			err := MenuConfig.SeoUpdate(seo)
			if err != nil {
				fmt.Println(err)
			}
			c.SetCookie("get-toast", "SEO Details Updated Successfully", 3600, "", "", false, false)
			c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)

		}

		c.Redirect(301, "/admin/website/seo"+url)

	}
}

//Settings Page

func Settings(c *gin.Context) {

	webid, _ := strconv.Atoi(c.Query("webid"))

	settingsdetail, err := MenuConfig.SettingsDetail(TenantId, webid)
	if err != nil {
		fmt.Println(err)
	}
	webbanner, _ := c.Cookie("webbanner")

	if webbanner == "" {

		webbanner = "true"
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

	var pages bool

	if webid != 0 {
		pages = true

	}
	Websitedet, _ := MenuConfig.GetWebsiteById(webid, TenantId)
	templatedetail, err := MenuConfig.GetTemplateById(Websitedet.TemplateId, TenantId)

	if err != nil {

		fmt.Println(err)
	}

	templatedetail.DateString = templatedetail.CreatedOn.In(TZONE).Format(Datelayout)

	ModuleName, _, _ := ModuleRouteName(c)

	translate, _ := TranslateHandler(c)

	c.HTML(200, "template_setting.html", gin.H{"csrf": csrf.GetToken(c), "websitedetail": Websitedet, "gotemplatepageheader": pages, "templatedetail": templatedetail, "webbanner": webbanner, "Menu": NewMenuController(c), "linktitle": "Setting", "translate": translate, "title": ModuleName, "settingsdetail": settingsdetail})

}

func SettingPage(c *gin.Context) {

	siteName := c.PostForm("siteName")

	sitelogoimage := c.PostForm("sitelogoimage")

	sitefaviconimage := c.PostForm("sitefaviconimage")

	websiteInput := c.PostForm("websiteInput")

	webid := c.Query("webid")

	websiteid, _ := strconv.Atoi(c.Query("webid"))

	var url string

	if webid != "" && webid != "0" {

		url = "?webid=" + webid

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

	if permisison {

		var imageName, imagePath, favimageName, favimagePath string

		var imageByte, favimageByte []byte

		tenantDetails, err := NewTeam.GetTenantDetails(TenantId)
		if err != nil {
			ErrorLog.Printf("error get storage type error: %s", err)
		}

		if sitelogoimage != "" {

			imageName, imagePath, imageByte, _ = ConvertBase64toByte(sitelogoimage, "setting")

			imagePath = tenantDetails.S3FolderName + imagePath

			uerr := storagecontroller.UploadCropImageS3(imageName, imagePath, imageByte)
			if uerr != nil {
				c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
				c.Redirect(301, "/admin/website/setting"+url)
				return
			}
		}

		if sitefaviconimage != "" {
			favimageName, favimagePath, favimageByte, _ = ConvertBase64toByte(sitefaviconimage, "setting")

			favimagePath = tenantDetails.S3FolderName + favimagePath

			uerr := storagecontroller.UploadCropImageS3(favimageName, favimagePath, favimageByte)
			if uerr != nil {
				c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
				c.Redirect(301, "/admin/website/setting"+url)
				return
			}
		}

		setting := menu.TblGoTemplateSettings{
			SiteName:        siteName,
			SiteLogo:        imageName,
			SiteLogoPath:    imagePath,
			SiteFavIcon:     favimageName,
			SiteFavIconPath: favimagePath,
			WebsiteUrl:      websiteInput,
			TenantId:        TenantId,
			WebsiteId:       websiteid,
		}

		if siteName != "" || imageName != "" || favimageName != "" || websiteInput != "" {

			err := MenuConfig.SettingUpdate(setting)
			if err != nil {
				fmt.Println(err)
			}
			c.SetCookie("get-toast", "Settings Details Updated Successfully", 3600, "", "", false, false)
			c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)

		}

		c.Redirect(301, "/admin/website/setting"+url)

	}
}

func CheckDomainName(c *gin.Context) {

	subdomain := c.PostForm("subdomain")

	userid := c.GetInt("userid")

	err := NewTeamWP.CheckDomainName(subdomain, userid, TenantId)

	if err != nil {

		c.JSON(200, gin.H{"status": true})

		return
	}

	c.JSON(200, gin.H{"status": false})
}
