package controllers

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"spurt-cms/models"
	storagecontroller "spurt-cms/storage-controller"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	"github.com/spurtcms/blocks"
	chn "github.com/spurtcms/channels"
	forms "github.com/spurtcms/forms-builders"
	listing "github.com/spurtcms/listing"
	menu "github.com/spurtcms/menu"
	csrf "github.com/utrack/gin-csrf"
	"gorm.io/datatypes"
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
	} else if filter.ToDate != "" {

		filterflag = true
	} else if filter.Status != "" {

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

	// webid, _ := strconv.Atoi(c.Param("id"))

	Websitedet, _ := MenuConfig.GetWebsiteById(0, TenantId)

	templatedetail, err := MenuConfig.GetTemplateById(Websitedet.TemplateId, TenantId)

	if err != nil {

		fmt.Println(err)
	}

	templatedetail.DateString = templatedetail.CreatedOn.In(TZONE).Format(Datelayout)

	templatepagelist, count, err := MenuConfig.GetTemplatePageList(100, 0, filter, TenantId, 0)
	var pagelist []menu.TblTemplatePages

	for _, page := range templatepagelist {

		page.CreatedDate = page.CreatedOn.In(TZONE).Format(Datelayout)
		if !page.ModifiedOn.IsZero() {
			page.ModifiedDate = page.ModifiedOn.In(TZONE).Format(Datelayout)
		} else {
			page.ModifiedDate = page.CreatedOn.In(TZONE).Format(Datelayout)
		}

		// GetmenuNames, err := MenuConfig.GetMenusByPageId(page.Id, TenantId)

		// if err != nil {

		// 	fmt.Println(err)
		// }

		page.MenuNames = "/pages/" + page.Slug

		pagelist = append(pagelist, page)
	}

	if err != nil {

		fmt.Println(err)
	}
	baseurl := os.Getenv("BASE_URL")
	paginationendcount := len(pagelist) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(pageno, int(count), limt)

	ModuleName, _, _ := ModuleRouteName(c)

	translate, _ := TranslateHandler(c)
	Websitedet.DateString = Websitedet.ModifiedOn.In(TZONE).Format(Datelayout)
	url := os.Getenv("BASE_URL")
	c.HTML(200, "go_template_edit.html", gin.H{"Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3}, "csrf": csrf.GetToken(c), "baseurl": baseurl, "Searchtrue": filterflag, "Count": count, "Limit": limt, "templatepagelist": pagelist, "templatedetail": templatedetail, "web": Websitedet, "Menu": NewMenuController(c), "linktitle": "Website Page", "translate": translate, "title": ModuleName, "Paginationendcount": paginationendcount, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Filter": filter, "Paginationstartcount": paginationstartcount, "Url": url})

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
	// webid, _ := strconv.Atoi(c.Query("webid"))

	PageDetail, _ := MenuConfig.GetPageById(pageid, TenantId)

	fmt.Println(PageDetail.PageDescription, "descriptionnnn")

	Websitedet, _ := MenuConfig.GetWebsiteById(0, TenantId)

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

	Formlist, _, _, err := FormConfig.FormBuildersList(100, 0, forms.Filter{}, TenantId, 3, 0, "", 0)
	publicFORMList := map[string]interface{}{"data": Formlist}
	publicFORMListbyte, _ := json.Marshal(publicFORMList)

	baseurl := os.Getenv("BASE_URL")

	urlpath := map[string]interface{}{"path": baseurl + "uploadb64image", "payload": "imagedata", "success": map[string]interface{}{"imagepath": "imagepath", "imagename": "imagename"}}

	newpath := os.Getenv("BASE_URL")

	ubyte, _ := json.Marshal(urlpath)

	fmt.Println("demonstration print")

	templatepagelist, _, err := MenuConfig.GetTemplatePageList(100, 0, menu.Filter{}, TenantId, 0)

	ModuleName, _, _ := ModuleRouteName(c)

	translate, _ := TranslateHandler(c)

	allTemplateInfos, _ := ReadYamlFile("websites/themes")
	WEBSITE_DATA_ID := os.Getenv("WEBSITE_DATA_ID")
	websiteId, _ := strconv.Atoi(WEBSITE_DATA_ID)
	website, _ := MenuConfig.GetWebsiteById(websiteId, TenantId)
	Template, err := MenuConfig.GetTemplateById(website.TemplateId, TenantId)

	custompagelist, _ := getCustomPagesList(strings.ToLower(strings.ReplaceAll(templatedetail.TemplateName, " ", "_")))
	BlockURL := os.Getenv("S3_ENDPOINT_URL")
	landingpagelist, _ := GetLandingPagesList(strings.ToLower(strings.ReplaceAll(templatedetail.TemplateName, " ", "_")))
	c.HTML(200, "go_template_addpage.html", gin.H{"csrf": csrf.GetToken(c), "custompagelist": custompagelist, "landingpagelist": landingpagelist, "channellist": channelist, "websitedetail": Websitedet, "pagelist": templatepagelist, "PageDetail": PageDetail, "editpage": editpage, "Storagepath": string(ubyte), "uploadurl": newpath, "StorageType": selectedtype.SelectedType, "blocks": string(bytedata1), "defaultblocks": string(bytedata2), "Mode": mode, "templatedetail": templatedetail, "Menu": NewMenuController(c), "linktitle": "Website", "translate": translate, "title": ModuleName, "publicdata": string(publicFORMListbyte), "BlockURL": BlockURL, "templatelist": allTemplateInfos, "TemplateName": Template.TemplateName}) // ✅ ADD THIS

}

func GetLandingPagesList(name string) ([]string, error) {

	basePath := "websites/themes/" + name + "/pages/landing-pages"

	entries, err := os.ReadDir(basePath)
	if err != nil {
		return nil, err
	}

	var pages []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		// only .html files
		if filepath.Ext(entry.Name()) == ".html" {
			pages = append(pages, entry.Name())
		}
	}

	return pages, nil
}

func getCustomPagesList(name string) ([]string, error) {
	basePath := "websites/themes/" + name + "/pages/static-pages"

	entries, err := os.ReadDir(basePath)
	if err != nil {
		return nil, err
	}

	var pages []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		// only .html files
		if filepath.Ext(entry.Name()) == ".html" {
			pages = append(pages, entry.Name())
		}
	}

	return pages, nil
}

//save page Name//

func SavePage(c *gin.Context) {
	fmt.Println("checkpagee")

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
	// webid, _ := strconv.Atoi(c.PostForm("webid"))
	pageid, _ := strconv.Atoi(c.PostForm("pageid"))
	pagedata := c.PostForm("pagedata")
	pagetype := c.PostForm("pagetype")
	pagepath := c.PostForm("pagepath")
	metatitle := c.PostForm("meta_title")
	metadesc := c.PostForm("meta_description")
	metakeywords := c.PostForm("meta_keywords")
	metaslug := c.PostForm("page_slug")
	fmt.Println(metaslug, "pageinfooo")
	var pagedet menu.TblTemplatePages
	var err error

	var slugname string

	slugname = strings.ToLower(strings.ReplaceAll(metaslug, " ", "-"))

	if pagetype == "block-page" {

		slugname = ""
	}
	templatepagelist, _, err := MenuConfig.GetTemplatePageList(100, 0, menu.Filter{}, TenantId, 0)
	if pageid != 0 {

		pageinfo := menu.TblTemplatePages{
			Id:              pageid,
			Name:            pagename,
			TenantId:        TenantId,
			ModifiedBy:      userid,
			PageDescription: pagedata,
			// WebsiteId:       webid,
			Slug:            slugname,
			PageType:        pagetype,
			CustomPagePath:  pagepath,
			MetaTitle:       metatitle,
			MetaDescription: metadesc,
			MetaKeywords:    metakeywords,
		}
		pagedet, err = MenuConfig.EditTemplatePage(&pageinfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if pagetype != "block-page" {

			var routeinput chn.TblRouteSlugs
			routeinput.ProductId = pagedet.Id
			routeinput.Slug = pagedet.Slug
			routeinput.TenantId = TenantId
			routeinput.ModuleName = "Pages"
			routeinput.ModifiedBy = userid
			err := ChannelConfig.UpdateGenericRouteslug(routeinput)

			if err != nil {

				fmt.Println(err)
			}

		}

	} else {

		for _, value := range templatepagelist {
			value.OrderIndex = value.OrderIndex + 1

			_, err := MenuConfig.UpdatePageOrderIndex(value.OrderIndex, value.Id, userid, TenantId)
			fmt.Println("err", err)

		}

		pageinfo := menu.TblTemplatePages{
			Name:      pagename,
			TenantId:  TenantId,
			CreatedBy: userid,
			IsDeleted: 0,
			Status:    1,
			// WebsiteId:       webid,
			PageType:        pagetype,
			CustomPagePath:  pagepath,
			MetaTitle:       metatitle,
			MetaDescription: metadesc,
			MetaKeywords:    metakeywords,
			Slug:            slugname,
			PageDescription: pagedata,
			OrderIndex:      1,
		}
		pagedet, err = MenuConfig.CreateTemplatePage(&pageinfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if pagetype != "block-page" {
			var routeinput chn.TblRouteSlugs
			routeinput.ProductId = pagedet.Id
			routeinput.Slug = pagedet.Slug
			routeinput.TenantId = TenantId
			routeinput.ModuleName = "Pages"
			routeinput.CreatedBy = userid
			_, err := ChannelConfig.CreateGenetricRouteslug(routeinput)

			if err != nil {

				fmt.Println(err)
			}
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
	// webid := c.Query("webid")

	var url string
	if pageno != "" {
		url = "/admin/website/pages/" + "?page=" + pageno
	} else {
		url = "/admin/website/pages/"

	}

	err := MenuConfig.DeletePage(pageid, userid, TenantId)

	err = ChannelConfig.DeleteGenericRouteslug("Pages", pageid, TenantId, userid)
	if err != nil {
		fmt.Println("Channel slug Delete Error : ", err)
	}

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

	// fmt.Println(pageid, "pageiddd")
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
		ErrorLog.Printf("menu status change error: %s", err)
		json.NewEncoder(c.Writer).Encode(flg)

	} else {
		json.NewEncoder(c.Writer).Encode(flg)
	}

}

//SEO Page

func Seo(c *gin.Context) {

	// webid, _ := strconv.Atoi(c.Query("webid"))

	seodetail, err := MenuConfig.SeoDetail(TenantId, 0)
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

	// if webid != 0 {
	// 	pages = true

	// }
	websiteid := os.Getenv("WEBSITE_DATA_ID")
	webint, _ := strconv.Atoi(websiteid)
	Websitedet, _ := MenuConfig.GetWebsiteById(webint, TenantId)

	//  Set CreatedDate like Templatesetting does
	Websitedet.CreatedDate = Websitedet.CreatedOn.In(TZONE).Format(Datelayout)
	if !Websitedet.ModifiedOn.IsZero() {
		Websitedet.DateString = Websitedet.ModifiedOn.In(TZONE).Format(Datelayout)
	} else {
		Websitedet.DateString = Websitedet.CreatedOn.In(TZONE).Format(Datelayout)
	}

	templatedetail, err := MenuConfig.GetTemplateById(Websitedet.TemplateId, TenantId)

	if err != nil {

		fmt.Println(err)
	}

	templatedetail.DateString = templatedetail.CreatedOn.In(TZONE).Format(Datelayout)

	Websitedet.TemplateName = templatedetail.TemplateName

	UserDetails, _, _ := NewTeam.GetUserById(Websitedet.CreatedBy, []int{})

	ModuleName, _, _ := ModuleRouteName(c)

	translate, _ := TranslateHandler(c)

	url := os.Getenv("BASE_URL")

	c.HTML(200, "seo.html", gin.H{"csrf": csrf.GetToken(c), "web": Websitedet, "gotemplatepageheader": pages, "templatedetail": templatedetail, "webbanner": webbanner, "Menu": NewMenuController(c), "linktitle": "SEO", "translate": translate, "title": ModuleName, "seodetail": seodetail, "Url": url, "UserDetails": UserDetails, "tabactive": "seo"})

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

		// REMOVED: var imageByte []byte  ← This was causing the error

		tenantDetails, err := NewTeam.GetTenantDetails(TenantId)
		if err != nil {
			ErrorLog.Printf("error get storage type error: %s", err)
		}

		// In your SeoPage function - imagedata section:
		tenantDetails, err = NewTeam.GetTenantDetails(TenantId)
		if err != nil {
			ErrorLog.Printf("error get storage type error: %s", err)
		}

		var imageName, imagePath string

		if imagedata != "" {
			// Check if it's XML or image
			if strings.Contains(imagedata, "data:application/xml") || strings.Contains(imagedata, "data:text/xml") {
				// XML processing
				imageName, imagePath, _, err = ConvertBase64toXML(imagedata, "seo")
			} else {
				// Image processing (original)
				imageName, imagePath, _, err = ConvertBase64toByte(imagedata, "seo")
			}

			if err != nil {
				ErrorLog.Printf("File processing error: %v", err)
				c.SetCookie("Alert-msg", "File processing failed", 3600, "", "", false, false)
				c.Redirect(301, "/admin/website/seo"+url)
				return
			}

			// Add tenant prefix
			imagePath = tenantDetails.S3FolderName + imagePath
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

func readXMLFile(filePath string) ([]byte, error) {
	xmlFilePath := strings.TrimSuffix(filePath, filepath.Ext(filePath)) + ".xml"

	// Check if XML file exists first
	if _, err := os.Stat(xmlFilePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("XML file not found: %s", xmlFilePath)
	}

	data, err := os.ReadFile(xmlFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read XML file %s: %w", xmlFilePath, err)
	}

	return data, nil
}

func storeXMLDataToDB(xmlData []byte, imageName string) error {
	// Example: Parse XML and store in your database
	var xmlDoc struct {
		XMLName xml.Name `xml:"image_metadata"`
		Width   int      `xml:"width"`
		Height  int      `xml:"height"`
		Title   string   `xml:"title"`
	}

	if err := xml.Unmarshal(xmlData, &xmlDoc); err != nil {
		return fmt.Errorf("failed to parse XML: %w", err)
	}

	// Replace with your actual DB logic:
	// db.Exec("INSERT INTO image_metadata (image_name, width, height, title) VALUES (?, ?, ?, ?)",
	//     imageName, xmlDoc.Width, xmlDoc.Height, xmlDoc.Title)

	log.Printf("Stored XML metadata for %s - Width: %d, Height: %d, Title: %s",
		imageName, xmlDoc.Width, xmlDoc.Height, xmlDoc.Title)

	return nil
}

//Settings Page

func Settings(c *gin.Context) {

	webid, _ := strconv.Atoi(c.Query("webid"))

	TENANTID := os.Getenv("TENANTID")

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

	channelslist, _, _ := ChannelConfig.ListChannel(chn.Channels{Limit: Limit, Offset: 0, Keyword: "", IsActive: false, TenantId: TenantId})

	website, _ := MenuConfig.GetWebsiteById(1, TENANTID)

	templatedet, _ := MenuConfig.GetTemplateById(website.TemplateId, website.TenantId)

	template_name := strings.ToLower(strings.ReplaceAll(templatedet.TemplateName, " ", "_"))

	goTemplateList, err := ReadYamlFileForWebsite("websites/themes", template_name)
	if err != nil {
		fmt.Println("Failed to read YAML files:", err)
		c.AbortWithStatus(500)
	}

	HtmlTemplate := strings.Split(goTemplateList.HtmlTemplates, ",")

	c.HTML(200, "template_setting.html", gin.H{"csrf": csrf.GetToken(c), "websitedetail": Websitedet, "gotemplatepageheader": pages, "templatedetail": templatedetail, "webbanner": webbanner, "Menu": NewMenuController(c), "linktitle": "Setting", "translate": translate, "title": ModuleName, "settingsdetail": settingsdetail, "HtmlTemplate": HtmlTemplate, "channelslist": channelslist, "TemplateTypeJSON": string(settingsdetail.TemplateType), "sociallink": string(settingsdetail.SocialMediaLink), "headerthame": settingsdetail.HeaderThame})
}

func SettingPage(c *gin.Context) {

	siteName := c.PostForm("siteName")

	sitelogoimage := c.PostForm("sitelogoimage")

	sitelogoDlt := c.PostForm("sitelogoDlt")

	sitefaviconimage := c.PostForm("sitefaviconimage")

	sitefaviconDlt := c.PostForm("sitefaviconDlt")

	websiteInput := c.PostForm("websiteInput")

	templateId := c.Param("template")

	websiteid, _ := strconv.Atoi(os.Getenv("WEBSITE_DATA_ID"))

	channel_template_jsonData := c.PostForm("channel_template_data")

	social_media_jsonData := c.PostForm("social_media_data")

	Headerthame := c.PostForm("headertheme")

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

		var favimageByte []byte

		tenantDetails, err := NewTeam.GetTenantDetails(TenantId)
		if err != nil {
			ErrorLog.Printf("error get storage type error: %s", err)
		}

		if sitelogoDlt == "1" {
			imagePath = ""
			imageName = ""
		} else if sitelogoimage != "" {

			// imageName, imagePath, imageByte, _ = ConvertBase64toByte(sitelogoimage, "setting")

			// fmt.Println("imimim",imageName,imagePath)

			// imagePath = tenantDetails.S3FolderName + imagePath

			// uerr := storagecontroller.UploadCropImageS3(imageName, imagePath, imageByte)
			// if uerr != nil {
			//  c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
			//  c.Redirect(301, "/admin/website/setting"+url)
			//  return
			// }

			if strings.HasPrefix(sitelogoimage, "data:image") {
				imageName, imagePath, err = ConvertBase64(sitelogoimage, strings.TrimPrefix("storage/website", "/"))

				// fmt.Println("imageName", imageName, "newimagepath", imagePath)

				if err != nil {
					ErrorLog.Printf("error get storage type error: %s", err)
				}
			} else {
				imagePath = sitelogoimage
				if len(imagePath) > 0 {
					parts := strings.Split(imagePath, "/")
					imageName = parts[len(parts)-1]
				}
			}

		}

		if sitefaviconDlt == "1" {

			favimageName = ""
			favimagePath = ""

		} else if sitefaviconimage != "" {

			if strings.HasPrefix(sitefaviconimage, "data:image") {

				favimageName, favimagePath, favimageByte, _ = ConvertBase64toByte(sitefaviconimage, "setting")

				favimagePath = tenantDetails.S3FolderName + favimagePath

				uerr := storagecontroller.UploadCropImageS3(favimageName, favimagePath, favimageByte)
				if uerr != nil {
					c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
					// c.Redirect(301, "/admin/browsetheme/templatesetting/"+templateId)

				}
			} else {
				favimagePath = sitefaviconimage
				if len(imagePath) > 0 {
					parts := strings.Split(imagePath, "/")
					favimageName = parts[len(parts)-1]
				}
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
			TemplateType:    datatypes.JSON([]byte(channel_template_jsonData)),
			SocialMediaLink: datatypes.JSON([]byte(social_media_jsonData)),
			HeaderThame:     Headerthame,
			TemplateID:      templateId,
		}

		webdata, err := MenuConfig.GetWebsiteById(websiteid, TenantId)
		if err != nil {
			ErrorLog.Printf("error getting website data: %s", err)
		}
		_, err = MenuConfig.GetTemplateById(webdata.TemplateId, TenantId)
		if err != nil {
			ErrorLog.Printf("error getting template data: %s", err)
		}

		// isTemplateProvided := channel_template_jsonData != "" &&
		//  channel_template_jsonData != "null" &&
		//  channel_template_jsonData != "[]"

		// isSociallinks := social_media_jsonData != "" &&
		//  social_media_jsonData != "null" &&
		//  social_media_jsonData != "[]"

		// if siteName != "" ||
		//  imageName != "" ||
		//  favimageName != "" ||
		//  websiteInput != "" ||
		//  isSociallinks ||
		//  isTemplateProvided ||
		//  Headerthame != "" {
		err = MenuConfig.SettingUpdate(setting)
		if err != nil {
			fmt.Println(err)
		}
		c.SetCookie("get-toast", "Settings Details Updated Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		// }
		c.Redirect(301, "/admin/browsetheme/templatesetting/"+templateId)

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

//widget page functionality//

func WidgetsList(c *gin.Context) {

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
	if filter.Keyword != "" || filter.ToDate != "" || filter.Status != "" {
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

	webid, _ := strconv.Atoi(c.Query("webid"))

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
	Websitedet, err := MenuConfig.GetWebsiteById(webid, TenantId)

	if Websitedet.Id == 0 {

		fmt.Println(err)
	}
	templatedetail, err := MenuConfig.GetTemplateById(Websitedet.TemplateId, TenantId)

	if err != nil {

		fmt.Println(err)
	}

	templatedetail.DateString = templatedetail.CreatedOn.In(TZONE).Format(Datelayout)

	Widgetlist, count, err := MenuConfig.GetWidgetList(limt, offset, menu.Filter{Keyword: filter.Keyword, Status: filter.Status, ToDate: filter.ToDate}, TenantId, webid)

	if err != nil {

		fmt.Println(err)
	}

	var AllWidgetList []menu.TblWidgets
	for _, widget := range Widgetlist {

		widget.CreatedDate = widget.CreatedOn.In(TZONE).Format(Datelayout)
		if !widget.ModifiedOn.IsZero() {
			widget.ModifiedDate = widget.ModifiedOn.In(TZONE).Format(Datelayout)
		} else {
			widget.ModifiedDate = widget.CreatedOn.In(TZONE).Format(Datelayout)
		}

		AllWidgetList = append(AllWidgetList, widget)
	}

	paginationendcount := len(Widgetlist) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(pageno, int(count), limt)

	ModuleName, _, _ := ModuleRouteName(c)

	translate, _ := TranslateHandler(c)

	c.HTML(200, "configure.html", gin.H{"csrf": csrf.GetToken(c), "Filter": filter, "Widgetlist": AllWidgetList, "Searchtrue": filterflag, "Count": count, "Limit": limt, "websitedetail": Websitedet, "gotemplatepageheader": pages, "templatedetail": templatedetail, "templateid": Websitedet.TemplateId, "webbanner": webbanner, "Menu": NewMenuController(c), "linktitle": "Widgets", "translate": translate, "title": ModuleName, "Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3,
	}, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page})

}

func CreateWidget(c *gin.Context) {
	webid, _ := strconv.Atoi(c.Query("webid"))
	tempid, _ := strconv.Atoi(c.Query("tempid"))

	template, _ := MenuConfig.GetTemplateById(tempid, TenantId)
	templateName := template.TemplateName
	if templateName == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
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

	widgetid, _ := strconv.Atoi(c.Param("id"))

	var widgetdetail menu.TblWidgets
	var productids []string
	var action string
	if widgetid != 0 {
		action = "/admin/website/widgets/updatewidget"

		detail, product, err := MenuConfig.GetWidgetById(widgetid, TenantId)
		if err != nil {
			fmt.Println(err)
		}

		for _, val := range product {
			productids = append(productids, strconv.Itoa(val.ProductId))
		}
		if detail.Id != 0 {
			widgetdetail = detail
		}

	} else {
		action = "/admin/website/widgets/savewidget"
	}

	productIdsStr := strings.Join(productids, ",")

	channelist, _, _ := ChannelConfig.ListChannel(chn.Channels{Limit: 100, Offset: 0, IsActive: true, TenantId: TenantId})
	Categorylist, _ := CategoryConfig.AllCategoriesWithSubList(TenantId)
	Entrylist, _, _, _ := ChannelConfig.ChannelEntriesList(chn.Entries{Status: "Published"}, TenantId)
	pagelist, _, _ := MenuConfig.GetTemplatePageList(100, 0, menu.Filter{Status: "Active"}, TenantId, webid)
	listinglist, _, _ := ListingConfig.ListingsList(100, 0, listing.Filter{}, TenantId)

	templatedetail.DateString = templatedetail.CreatedOn.In(TZONE).Format(Datelayout)

	ModuleName, _, _ := ModuleRouteName(c)
	translate, _ := TranslateHandler(c)

	templates, err := ReadYamlFile("websites/themes")
	if err != nil {
		fmt.Println("Failed to read YAML files:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	settingsdetail, err := MenuConfig.SettingsDetail(TenantId, 1)
	if err != nil {
		fmt.Println(err)
	}

	var templateInfo TemplateInfo
	for _, t := range templates {
		if t.TemplateName == templateName {
			templateInfo = t
			break
		}
	}

	if c.GetHeader("X-Requested-With") == "XMLHttpRequest" {
		c.JSON(200, gin.H{
			"widgetdetail": widgetdetail,
			"productids":   productIdsStr,
			"action":       action,
			"templateid":   tempid,
		})
		return
	}

	c.HTML(200, "configure.html", gin.H{
		"csrf":                 csrf.GetToken(c),
		"action":               action,
		"websitedetail":        Websitedet,
		"productids":           productIdsStr,
		"gotemplatepageheader": pages,
		"templatedetail":       templatedetail,
		"webbanner":            webbanner,
		"Menu":                 NewMenuController(c),
		"linktitle":            "Widgets",
		"translate":            translate,
		"title":                ModuleName,
		"Categorylist":         Categorylist,
		"channelist":           channelist,
		"listinglist":          listinglist,
		"Entrylist":            Entrylist,
		"pagelist":             pagelist,
		"widgetdetail":         widgetdetail,
		"settingsdetail":       settingsdetail,
		"templateInfo":         templateInfo,
		"templateid":           tempid,
	})
}

func SaveWidget(c *gin.Context) {

	permisison, perr := NewAuth.IsGranted("Menu", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("menu authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("menu authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	title := c.PostForm("title")

	longtitle := c.PostForm("long_title")

	position := c.PostForm("position")

	widget_slug := c.PostForm("widget_slug")

	sort_order, _ := strconv.Atoi(c.PostForm("sort_order"))

	status, _ := strconv.Atoi(c.PostForm("status"))

	productids := c.PostForm("productids")

	fmt.Println(productids, "productiddd", sort_order)

	product_type := c.PostForm("widget_type")

	metatitle := c.PostForm("meta_title")

	metadescription := c.PostForm("meta_description")

	metakeyword := c.PostForm("meta_keyword")

	templaId := c.PostForm("templaId")
	TemplateId, _ := strconv.Atoi(templaId)

	widget_limit, _ := strconv.Atoi(c.PostForm("widget_limit"))

	if product_type == "Entries" && productids == "" {
		c.SetCookie("get-toast", "Please select at least one Entry", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "error", 3600, "", "", false, false)
		c.Redirect(301, "/admin/browsetheme/configure/"+templaId)
		return
	}

	userid := c.GetInt("userid")

	webid, _ := strconv.Atoi(c.Query("webid"))

	WidgetData := menu.TblWidgets{

		Title:           title,
		LongTitle:       longtitle,
		Slug:            widget_slug,
		Position:        position,
		SortOrder:       sort_order,
		Status:          status,
		WidgetType:      product_type,
		MetaTitle:       metatitle,
		MetaDescription: metadescription,
		MetaKeywords:    metakeyword,
		TenantId:        TenantId,
		CreatedBy:       userid,
		WebsiteId:       webid,
		ProductIds:      productids,
		WidgetLimit:     widget_limit,
		TemplateId:      TemplateId,
	}

	widgetid, _ := strconv.Atoi(c.PostForm("widget_id"))

	if widgetid != 0 {

		_, err := MenuConfig.UpdateWidget(&WidgetData, widgetid)

		if err != nil {

			fmt.Println(err)
		}

		DB.Table("tbl_widgets").Where("id = ? AND tenant_id = ?", widgetid, TenantId).UpdateColumn("template_id", TemplateId)

	} else {

		widget, err := MenuConfig.CreateWidget(&WidgetData)

		if err != nil {

			fmt.Println(err)
		}

		fmt.Println(widget, "widget")
	}

	pageno := c.Query("page")

	var url string
	var toastMsg string
	if pageno != "" {
		url = "/admin/browsetheme/configure/" + strconv.Itoa(TemplateId)
	} else {
		url = "/admin/browsetheme/configure/" + strconv.Itoa(TemplateId)
	}

	if widgetid != 0 {
		toastMsg = "Widget Updated Successfully"
	} else {
		toastMsg = "Widget Created Successfully"
	}

	c.SetCookie("get-toast", toastMsg, 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, url)

}

func DeleteWidget(c *gin.Context) {
	permisison, perr := NewAuth.IsGranted("Menu", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("menu authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("menu authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	widgetid, _ := strconv.Atoi(c.Param("id"))

	pageno := c.Query("page")
	userid := c.GetInt("userid")

	widgetdetail, _, _ := MenuConfig.GetWidgetById(widgetid, TenantId)
	WebSiteData, _ := MenuConfig.GetWebsiteById(widgetdetail.WebsiteId, TenantId)
	tempid := strconv.Itoa(WebSiteData.TemplateId)

	var url string

	err := MenuConfig.DeleteWidgetById(widgetid, userid, TenantId)

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
	if pageno != "" {
		url = "/admin/browsetheme/configure/" + tempid
	} else {
		url = "/admin/browsetheme/configure/" + tempid
	}

	c.SetCookie("get-toast", "Widget Deleted Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, url)

}

//Status Change

func WidgetStatusChange(c *gin.Context) {

	widgetid, _ := strconv.Atoi(c.PostForm("id"))

	userid := c.GetInt("userid")
	val, _ := strconv.Atoi(c.Request.PostFormValue("isactive"))

	permisison, perr := NewAuth.IsGranted("Menu", auth.Update, TenantId)
	if perr != nil {
		ErrorLog.Printf("menu authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("menu authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	flg, err := MenuConfig.WidgetStatusChange(widgetid, val, userid, TenantId)

	if err != nil {
		ErrorLog.Printf("menu status change error: %s", err)
		json.NewEncoder(c.Writer).Encode(flg)

	} else {
		json.NewEncoder(c.Writer).Encode(flg)
	}

}

func CheckPageName(c *gin.Context) {

	pageid, _ := strconv.Atoi(c.PostForm("pageid"))

	name := c.PostForm("page_name")

	websiteid, _ := strconv.Atoi(c.PostForm("webid"))

	permisison, perr := NewAuth.IsGranted("Website", auth.Read, TenantId)
	if perr != nil {
		ErrorLog.Printf("page check name authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("page authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		flg, err := MenuConfig.CheckPageNameIsExits(pageid, name, websiteid, TenantId)

		if err != nil {
			ErrorLog.Printf("page checkname error: %s", err)
			json.NewEncoder(c.Writer).Encode(false)
			return
		}

		json.NewEncoder(c.Writer).Encode(flg)

	}

}
func UpdatePagesOrder(c *gin.Context) {

	var orderData []menu.OrderItem
	err := json.Unmarshal([]byte(c.Request.PostFormValue("orderData")), &orderData)
	if err != nil {
		fmt.Println(err)
	}

	userid := c.GetInt("userid")

	merr := MenuConfig.UpdatePagesOrder(orderData, userid, TenantId)

	if merr != nil {
		fmt.Println(merr)
	}

	c.JSON(200, true)
}

//clone page

func ClonePage(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	permisison, perr := NewAuth.IsGranted("Website", auth.CRUD, TenantId)

	if perr != nil {
		ErrorLog.Printf("Website authorization error :%s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Website authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	pagedetails, _ := MenuConfig.GetPageById(id, TenantId)

	CloneCount := pagedetails.CloneCount + 1

	page := menu.TblTemplatePages{
		Id:         pagedetails.Id,
		CloneCount: CloneCount,
	}

	cerr := MenuConfig.CloneCountUpdate(page)
	if cerr != nil {

		ErrorLog.Printf("channel Type error: %s", cerr)
	}

	var Count string

	if pagedetails.CloneCount != 0 {

		Count = "(" + strconv.Itoa(pagedetails.CloneCount) + ")"

	}
	var slugname string

	slugname = strings.ToLower(strings.ReplaceAll("Clone of "+pagedetails.Name+Count, " ", "-"))

	if pagedetails.PageType == "block-page" {

		slugname = ""
	}
	templatepagelist, _, err := MenuConfig.GetTemplatePageList(100, 0, menu.Filter{}, TenantId, pagedetails.WebsiteId)

	if err != nil {

		fmt.Println(err)
	}
	if permisison {

		userid := c.GetInt("userid")

		if pagedetails.ParentId == 0 {

			for _, value := range templatepagelist {
				value.OrderIndex = value.OrderIndex + 1

				_, err := MenuConfig.UpdatePageOrderIndex(value.OrderIndex, value.Id, userid, TenantId)

				if err != nil {
					fmt.Println("err", err)
				}
			}
		}
		pageinfo := menu.TblTemplatePages{
			Name:            "Clone of " + pagedetails.Name + Count,
			TenantId:        TenantId,
			CreatedBy:       userid,
			IsDeleted:       0,
			Status:          1,
			WebsiteId:       pagedetails.WebsiteId,
			PageType:        pagedetails.PageType,
			CustomPagePath:  pagedetails.CustomPagePath,
			MetaTitle:       pagedetails.MetaTitle,
			MetaDescription: pagedetails.MetaDescription,
			MetaKeywords:    pagedetails.MetaKeywords,
			Slug:            slugname,
			PageDescription: pagedetails.PageDescription,
			ParentId:        pagedetails.ParentId,
			OrderIndex:      1,
		}
		pagedet, err := MenuConfig.CreateTemplatePage(&pageinfo)

		if pagedetails.PageType != "block-page" {
			var routeinput chn.TblRouteSlugs
			routeinput.ProductId = pagedet.Id
			routeinput.Slug = pagedet.Slug
			routeinput.TenantId = TenantId
			routeinput.ModuleName = "Pages"
			routeinput.CreatedBy = userid
			_, err := ChannelConfig.CreateGenetricRouteslug(routeinput)

			if err != nil {

				fmt.Println(err)
			}
		}
		if err != nil {
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
			return
		}

		c.SetCookie("get-toast", "Collection Cloned Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, "/admin/website/pages/")

	}
}

// latest

func Templatesetting(c *gin.Context) {
	tempid := c.Param("template")
	websiteid, _ := strconv.Atoi(os.Getenv("WEBSITE_DATA_ID"))
	// websiteid, _ := strconv.Atoi(c.Query("website_id"))

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

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}
	fmt.Println("filterflagfilterflagfilterflag:", filterflag)
	// ==================================

	settingsdetail, err := MenuConfig.SettingDetailBasedONTemp(tempid, TenantId, websiteid)
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
	webid := 0
	if webid != 0 {
		pages = true

	}
	// Websitedet, _ := MenuConfig.GetWebsiteById(webid, TenantId)
	tempidInt, _ := strconv.Atoi(tempid)
	templatedetail, err := MenuConfig.GetTemplateById(tempidInt, TenantId)

	if err != nil {

		fmt.Println(err)
	}

	templatedetail.DateString = templatedetail.CreatedOn.In(TZONE).Format(Datelayout)

	ModuleName, _, _ := ModuleRouteName(c)

	translate, _ := TranslateHandler(c)

	channelslist, _, _ := ChannelConfig.ListChannel(chn.Channels{Limit: 0, Offset: 0, Keyword: "", IsActive: false, TenantId: TenantId})

	template_name := strings.ToLower(strings.ReplaceAll(templatedetail.TemplateName, " ", "_"))

	goTemplateList, err := ReadYamlFileForWebsite("websites/themes", template_name)
	if err != nil {
		fmt.Println("Failed to read YAML files:", err)
		c.AbortWithStatus(500)
	}

	HtmlTemplate := strings.Split(goTemplateList.HtmlTemplates, ",")

	websitelist, _, err := MenuConfig.WebsiteList(limt, offset, menu.Filter{Keyword: filter.Keyword}, TenantId)

	baseurl := os.Getenv("BASE_URL")
	baseurl = strings.TrimPrefix(baseurl, "https://")
	baseurl = strings.TrimPrefix(baseurl, "http://")
	baseurl = strings.TrimSuffix(baseurl, "/")
	host := c.Request.Host
	isLocal := false
	var data menu.TblWebsite

	var FinalwebsiteList []menu.TblWebsite

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

	tabActive := "templatesetting"
	// c.HTML(200, "templatesetting.html", gin.H{"csrf": csrf.GetToken(c), "baseurl": baseurl, "gotemplatepageheader": pages, "templateInfo": templatedetail, "Menu": NewMenuController(c), "linktitle": "Template setting", "translate": translate, "title": ModuleName, "HtmlTemplate": HtmlTemplate, "channelslist": channelslist, "TemplateTypeJSON": string(settingsdetail.TemplateType), "sociallink": string(settingsdetail.SocialMediaLink), "headerthame": settingsdetail.HeaderThame, "settingsdetail": settingsdetail, "webbanner": webbanner, "tabactive": tabActive, "web": FinalwebsiteList, "data": data, "isLocal": isLocal})

	websiteInfo, _ := MenuConfig.GetWebsiteById(websiteid, TenantId)
	websiteInfo.CreatedDate = websiteInfo.CreatedOn.In(TZONE).Format(Datelayout)
	if !websiteInfo.ModifiedOn.IsZero() {
		websiteInfo.DateString = websiteInfo.ModifiedOn.In(TZONE).Format(Datelayout)
	} else {
		websiteInfo.DateString = websiteInfo.CreatedOn.In(TZONE).Format(Datelayout)
	}

	websiteInfo.TemplateName = templatedetail.TemplateName
	data.TemplateName = templatedetail.TemplateName

	UserDetails, _, _ := NewTeam.GetUserById(websiteInfo.CreatedBy, []int{})
	baseurl = strings.TrimPrefix(baseurl, "www.")

	var subdomainURLDemo string
	if strings.Contains(host, "localhost") {
		subdomainURLDemo = "http://" + websiteInfo.Name + "." + baseurl
	} else {
		subdomainURLDemo = "https://" + websiteInfo.Name + "." + baseurl
	}
	c.HTML(200, "templatesetting.html", gin.H{
		"hideConfigBtn":        true,
		"csrf":                 csrf.GetToken(c),
		"baseurl":              baseurl,
		"gotemplatepageheader": pages,
		"templateInfo":         templatedetail,
		"templatedetail":       templatedetail,
		"web":                  websiteInfo,
		"data":                 data,
		"UserDetails":          UserDetails,
		"Menu":                 NewMenuController(c),
		"linktitle":            "Template setting",
		"translate":            translate,
		"title":                ModuleName,
		"HtmlTemplate":         HtmlTemplate,
		"subdomainurldemo":     subdomainURLDemo,
		"channelslist":         channelslist,
		"TemplateTypeJSON":     string(settingsdetail.TemplateType),
		"sociallink":           string(settingsdetail.SocialMediaLink),
		"headerthame":          settingsdetail.HeaderThame,
		"settingsdetail":       settingsdetail,
		"webbanner":            webbanner,
		"tabactive":            tabActive,
		"websitelist":          FinalwebsiteList,
		"isLocal":              isLocal,
	})
}

// func DataSettinng(c *gin.Context) {

// 	tempid := c.Param("template")
// 	var (
// 		limt   int
// 		offset int
// 		filter menu.Filter
// 	)

// 	limit := c.Query("limit")

// 	if limit == "" {
// 		limt = Limit
// 	} else {
// 		limt, _ = strconv.Atoi(limit)
// 	}

// 	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
// 	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")
// 	filter.Status = strings.Trim(c.DefaultQuery("status", ""), " ")
// 	filter.ToDate = strings.Trim(c.DefaultQuery("lastupdated", ""), " ")
// 	var filterflag bool
// 	if filter.Keyword != "" {
// 		filterflag = true
// 	} else {
// 		filterflag = false
// 	}

// 	if limit == "" {
// 		limt = Limit
// 	} else {
// 		limt, _ = strconv.Atoi(limit)
// 	}

// 	if pageno != 0 {
// 		offset = (pageno - 1) * limt
// 	}
// 	fmt.Println("filterflagfilterflagfilterflag:", filterflag)

// 	webid, _ := strconv.Atoi(c.Query("templateid"))

// 	settingsdetail, err := MenuConfig.SettingsDetail(TenantId, webid)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	webbanner, _ := c.Cookie("webbanner")

// 	if webbanner == "" {

// 		webbanner = "true"
// 	}

// 	permisison, perr := NewAuth.IsGranted("Menu", auth.CRUD, TenantId)
// 	if perr != nil {
// 		ErrorLog.Printf("menu authorization error: %s", perr)
// 	}

// 	if !permisison {
// 		ErrorLog.Printf("Menu authorization error")
// 		c.Redirect(301, "/403-page")
// 		return
// 	}

// 	var pages bool

// 	if webid != 0 {
// 		pages = true

// 	}
// 	// Websitedet, _ := MenuConfig.GetWebsiteById(webid, TenantId)
// 	tempidInt, _ := strconv.Atoi(tempid)
// 	templatedetail, err := MenuConfig.GetTemplateById(tempidInt, TenantId)

// 	if err != nil {

// 		fmt.Println(err)
// 	}
// 	templatedetail.DateString = templatedetail.CreatedOn.In(TZONE).Format(Datelayout)

// 	ModuleName, _, _ := ModuleRouteName(c)

// 	translate, _ := TranslateHandler(c)

// 	channelslist, _, _ := ChannelConfig.ListChannel(chn.Channels{Limit: Limit, Offset: 0, Keyword: "", IsActive: false, TenantId: TenantId})

// 	goTemplateList, err := ReadYamlFile("websites/themes")
// 	if err != nil {
// 		fmt.Println("Failed to read YAML files:", err)

// 	}

// 	HtmlTemplate := strings.Split(goTemplateList[0].HtmlTemplates, ",")

// 	websitelist, _, err := MenuConfig.WebsiteList(limt, offset, menu.Filter{Keyword: filter.Keyword}, TenantId)

// 	baseurl := os.Getenv("BASE_URL")
// 	baseurl = strings.TrimPrefix(baseurl, "https://")
// 	baseurl = strings.TrimPrefix(baseurl, "http://")
// 	baseurl = strings.TrimSuffix(baseurl, "/")
// 	host := c.Request.Host
// 	isLocal := false
// 	var data menu.TblWebsite

// 	var FinalwebsiteList []menu.TblWebsite

// 	for _, val := range websitelist {
// 		val.CreatedDate = val.CreatedOn.In(TZONE).Format(Datelayout)
// 		if !val.ModifiedOn.IsZero() {
// 			val.DateString = val.ModifiedOn.In(TZONE).Format(Datelayout)
// 		} else {
// 			val.DateString = val.CreatedOn.In(TZONE).Format(Datelayout)
// 		}
// 		if strings.Contains(host, "localhost") {
// 			val.Subdomain = "http://" + val.Name + "." + baseurl
// 		} else {
// 			val.Subdomain = "https://" + val.Name + "." + baseurl
// 		}
// 		FinalwebsiteList = append(FinalwebsiteList, val)
// 		data = val
// 	}

// 	tabActive := "data"
// 	// c.HTML(200, "browsethemedata.html", gin.H{"csrf": csrf.GetToken(c), "baseurl": baseurl, "gotemplatepageheader": pages, "templateInfo": templatedetail, "Menu": NewMenuController(c), "linktitle": "Data", "translate": translate, "title": ModuleName, "HtmlTemplate": HtmlTemplate, "channelslist": channelslist, "TemplateTypeJSON": string(settingsdetail.TemplateType), "sociallink": string(settingsdetail.SocialMediaLink), "headerthame": settingsdetail.HeaderThame, "settingsdetail": settingsdetail, "webbanner": webbanner, "tabactive": tabActive, "web": FinalwebsiteList, "data": data, "isLocal": isLocal})
// 	websiteId, _ := strconv.Atoi(os.Getenv("WEBSITE_DATA_ID"))
// 	websiteInfo, _ := MenuConfig.GetWebsiteById(websiteId, TenantId)
// 	websiteInfo.CreatedDate = websiteInfo.CreatedOn.In(TZONE).Format(Datelayout)
// 	websiteInfo.DateString = websiteInfo.ModifiedOn.In(TZONE).Format(Datelayout)

// 	websiteInfo.TemplateName = templatedetail.TemplateName

// 	UserDetails, _, _ := NewTeam.GetUserById(websiteInfo.CreatedBy, []int{})

// 	c.HTML(200, "browsethemedata.html", gin.H{
// 		"hideConfigBtn":        true,
// 		"csrf":                 csrf.GetToken(c),
// 		"baseurl":              baseurl,
// 		"gotemplatepageheader": pages,
// 		"templatedetail":       templatedetail,
// 		"templateInfo":         templatedetail,
// 		"web":                  websiteInfo,
// 		"data":                 data,
// 		"UserDetails":          UserDetails,
// 		"Menu":                 NewMenuController(c),
// 		"linktitle":            "Data",
// 		"translate":            translate,
// 		"title":                ModuleName,
// 		"HtmlTemplate":         HtmlTemplate,
// 		"channelslist":         channelslist,
// 		"TemplateTypeJSON":     string(settingsdetail.TemplateType),
// 		"sociallink":           string(settingsdetail.SocialMediaLink),
// 		"headerthame":          settingsdetail.HeaderThame,
// 		"settingsdetail":       settingsdetail,
// 		"webbanner":            webbanner,
// 		"tabactive":            tabActive,
// 		"isLocal":              isLocal,
// 		"websitelist":          FinalwebsiteList,
// 	})
// }
