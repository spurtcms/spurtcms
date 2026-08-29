package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	storagecontroller "spurt-cms/storage-controller"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	"github.com/spurtcms/categories"
	chn "github.com/spurtcms/channels"
	"github.com/spurtcms/courses"
	forms "github.com/spurtcms/forms-builders"
	"github.com/spurtcms/listing"
	menu "github.com/spurtcms/menu"
	csrf "github.com/utrack/gin-csrf"
)

// ContentMenuList
func MenuList(c *gin.Context) {

	var (
		limt           int
		offset         int
		filter         menu.Filter
		FinalMenusList []menu.TblMenus
	)

	//get values from url query
	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")
	filter.Status = strings.Trim(c.DefaultQuery("status", ""), " ")
	filter.ToDate = strings.Trim(c.DefaultQuery("lastupdated", ""), " ")
	// webid, _ := strconv.Atoi(c.Query("webid"))
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
	permission, perr := NewAuth.IsGranted("Menu", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("menu authorization error:%s", perr)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if !permission {
		ErrorLog.Printf("menu authorization error")
		c.Redirect(http.StatusMovedPermanently, "/403-page")
		return
	}

	MenuConfig.DataAccess = c.GetInt("dataaccess")
	MenuConfig.UserId = c.GetInt("userid")

	Menulist, Total_count, err := MenuConfig.MenuList(limt, offset, menu.Filter(filter), TenantId, 0)

	if err != nil {
		ErrorLog.Printf("menu list  error: %s", err)
	}
	for _, val := range Menulist {

		if !val.ModifiedOn.IsZero() {
			val.DateString = val.ModifiedOn.In(TZONE).Format(Datelayout)
		} else {
			val.DateString = val.CreatedOn.In(TZONE).Format(Datelayout)
		}

		menuitemslist, _ := MenuConfig.GetMenusByParentid(val.Id, TenantId)

		val.MenuitemCount = len(menuitemslist) - 1

		fmt.Println("lengdffd", len(menuitemslist))
		FinalMenusList = append(FinalMenusList, val)

	}

	//pagination calc
	paginationendcount := len(Menulist) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(pageno, int(Total_count), limt)

	webbanner, _ := c.Cookie("webbanner")

	if webbanner == "" {

		webbanner = "true"
	}

	ModuleName, TabName, _ := ModuleRouteName(c)

	translate, _ := TranslateHandler(c)

	var pages bool

	// if webid != 0 {
	// 	pages = true

	// }

	Websitedet, _ := MenuConfig.GetWebsiteById(0, TenantId)
	templatedetail, err := MenuConfig.GetTemplateById(Websitedet.TemplateId, TenantId)

	if err != nil {

		fmt.Println(err)
	}

	templatedetail.DateString = templatedetail.CreatedOn.In(TZONE).Format(Datelayout)
	Websitedet.DateString = Websitedet.ModifiedOn.In(TZONE).Format(Datelayout)

	url := os.Getenv("BASE_URL")

	c.HTML(200, "menulist.html", gin.H{"Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3,
	}, "Searchtrue": filterflag, "webbanner": webbanner, "web": Websitedet, "gotemplatepageheader": pages, "templatedetail": templatedetail, "Count": Total_count, "Limit": limt, "csrf": csrf.GetToken(c), "title": ModuleName, "Tabmenu": TabName, "translate": translate, "Menulist": FinalMenusList, "Menu": NewMenuController(c), "linktitle": "Menus", "Paginationendcount": paginationendcount, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Filter": filter, "Paginationstartcount": paginationstartcount, "HeadTitle": "Menu", "Url": url})

}

// Content MenuItemsList
func MenuIemsList(c *gin.Context) {

	permission, perr := NewAuth.IsGranted("Menu", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("menu authorization error:%s", perr)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if !permission {
		ErrorLog.Printf("menu authorization error")
		c.Redirect(http.StatusMovedPermanently, "/403-page")
		return
	}
	var recentChannels []chn.Tblchannel
	var recentCourses []courses.TblCourses
	var recentForms []forms.TblForms
	var recentCategories []categories.Arrangecategories
	var recentPages []menu.TblTemplatePages
	TenMinutesAgo := time.Now().Add(-60 * time.Minute)
	channelist, _, clerr := ChannelConfig.ListChannel(chn.Channels{Limit: 100, Offset: 0, IsActive: true, TenantId: TenantId})
	if clerr != nil {
		ErrorLog.Printf("channellist error :%s", clerr)
	}

	courseslist, _, err := CoursesConfig.CoursesList(100, 0, courses.Filter{}, TenantId)
	if err != nil {
		ErrorLog.Printf("courselist error :%s", err)
	}

	Formlist, _, _, err := FormConfig.FormBuildersList(100, 0, forms.Filter{}, TenantId, 1, 0, "", 0)

	if err != nil {
		ErrorLog.Printf("formslist error :%s", err)
	}
	for _, val := range channelist {
		if val.CreatedOn.After(TenMinutesAgo) {
			recentChannels = append(recentChannels, val)
		}
	}
	for _, val := range courseslist {
		if val.CreatedOn.After(TenMinutesAgo) {
			recentCourses = append(recentCourses, val)
		}
	}
	for _, val := range Formlist {
		if val.CreatedOn.After(TenMinutesAgo) {
			recentForms = append(recentForms, val)
		}
	}

	menuid, _ := strconv.Atoi(c.Param("id"))

	menuitemslist, err := MenuConfig.GetMenusByParentid(menuid, TenantId)

	var newmenuitemlist []menu.TblMenus

	for _, item := range menuitemslist {

		item.HtmlDescription = template.HTML(item.Description)

		newmenuitemlist = append(newmenuitemlist, item)
	}

	menudetails, _ := MenuConfig.GetmenyById(menuid, TenantId)

	Categorylist, _ := CategoryConfig.AllCategoriesWithSubList(TenantId)

	for _, val := range Categorylist {
		for _, val1 := range val.Categories {
			if val1.CreatedOn.After(TenMinutesAgo) {
				recentCategories = append(recentCategories, val)
			}
		}
	}

	var selectedcategoryids []int

	var selectedlistingids []int

	for _, val := range Categorylist {

		for _, val1 := range val.Categories {

			for _, val2 := range menuitemslist {

				if val2.Type == "categories" && val2.TypeId == val1.Id {

					selectedcategoryids = append(selectedcategoryids, val1.Id)
				}
				if val2.Type == "listings" && val2.TypeId == val1.Id {

					selectedlistingids = append(selectedlistingids, val1.Id)
				}
			}
		}
	}

	ModuleName, TabName, _ := ModuleRouteName(c)

	translate, _ := TranslateHandler(c)

	webid, _ := strconv.Atoi(c.Query("webid"))

	var pages bool

	if webid != 0 {
		pages = true

	}

	Websitedet, _ := MenuConfig.GetWebsiteById(0, TenantId)
	templatedetail, err := MenuConfig.GetTemplateById(Websitedet.TemplateId, TenantId)

	if err != nil {

		fmt.Println(err)
	}

	Entrylist, _, _, _ := ChannelConfig.ChannelEntriesList(chn.Entries{}, TenantId)

	pagelist, _, err := MenuConfig.GetTemplatePageList(100, 0, menu.Filter{Status: "Active"}, TenantId, 0)

	for _, val := range pagelist {
		if val.CreatedOn.After(TenMinutesAgo) {
			recentPages = append(recentPages, val)
		}
	}

	listinglist, _, err := ListingConfig.ListingsList(100, 0, listing.Filter{}, TenantId)
	if err != nil {
		fmt.Printf("Unable to fetch Listings List:%v", err)
	}

	c.HTML(200, "menuitems.html", gin.H{"csrf": csrf.GetToken(c), "recentPages": recentPages, "recentCategories": recentCategories, "selectedcategoryids": selectedcategoryids, "selectedlistingids": selectedlistingids, "websitedetail": Websitedet, "Entrylist": Entrylist, "pagelist": pagelist, "listinglist": listinglist, "gotemplatepageheader": pages, "templatedetail": templatedetail, "recentChannels": recentChannels, "recentCourses": recentCourses, "recentForms": recentForms, "menudetails": menudetails, "menuitemslist": newmenuitemlist, "title": ModuleName, "Tabmenu": TabName, "translate": translate, "Menu": NewMenuController(c), "linktitle": "Menus Items", "channelist": channelist, "courselist": courseslist, "Formlist": Formlist, "menuid": menuid, "Categorylist": Categorylist})

}

// CreateMenu
func CreateMenu(c *gin.Context) {
	menuname := c.PostForm("menu_name")
	menutitle := c.PostForm("menu_title")

	menudesc := c.PostForm("menu_desc")
	userid := c.GetInt("userid")

	websiteid := c.Query("webid")

	webid, _ := strconv.Atoi(c.Query("webid"))

	menustatus, _ := strconv.Atoi(c.PostForm("menustatus"))
	menugroup := c.PostForm("menu_group")
	menuOrder, _ := strconv.Atoi(c.PostForm("menu_order"))

	permission, perr := NewAuth.IsGranted("Menu", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("menu authorization error:%s", perr)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if !permission {
		ErrorLog.Printf("menu authorization error")
		c.Redirect(http.StatusMovedPermanently, "/403-page")
		return
	}

	_, err := MenuConfig.CreateMenus(menu.MenuCreate{MenuName: menuname, MenuTitle: menutitle, WebsiteId: webid, Description: menudesc, Status: menustatus, MenuGroup: menugroup, OrderIndex: menuOrder, TenantId: TenantId, CreatedBy: userid, ParentId: 0})

	if strings.Contains(fmt.Sprint(err), "given some values is empty") {
		ErrorLog.Printf("Menu mandatory field error: %s", err)
		c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
		c.Redirect(301, "/admin/website/menu?webid="+websiteid)
		return
	}

	if err != nil {
		ErrorLog.Printf("Menu error: %s", perr)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.Redirect(http.StatusMovedPermanently, "/admin/website/menu")
		return
	}

	c.SetCookie("get-toast", "Menu Created Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(http.StatusMovedPermanently, "/admin/website/menu")

}

func UpdateMenu(c *gin.Context) {
	web := os.Getenv("WEBSITE_DATA_ID")
	userid := c.GetInt("userid")
	menuid, _ := strconv.Atoi(c.PostForm("menu_id"))
	menustatus, _ := strconv.Atoi(c.PostForm("menustatus"))
	pageno := c.Request.PostFormValue("menupageno")
	menupage := c.PostForm("menupage")
	websiteId, _ := strconv.Atoi(web)
	// webid, _ := strconv.Atoi(c.Query("webid"))
	menugroup := c.PostForm("menu_group")
	menuOrder, _ := strconv.Atoi(c.PostForm("menu_order"))

	menudetails := menu.MenuCreate{
		MenuName:    c.PostForm("menu_name"),
		MenuTitle:   c.PostForm("menu_title"),
		Description: c.PostForm("menu_desc"),
		Status:      menustatus,
		ModifiedBy:  userid,
		TenantId:    TenantId,
		Id:          menuid,
		WebsiteId:   websiteId,
		MenuGroup:   menugroup,
		OrderIndex:  menuOrder,
	}
	var url string
	if menupage == "menuitems" {
		url = "/admin/website/menu/menuitems/" + c.PostForm("menu_id")
	} else {
		if pageno != "" {
			url = "/admin/website/menu?page=" + pageno
		} else {
			url = "/admin/website/menu"
		}
	}
	// if web != "" && web != "0" {
	// 	if strings.Contains(url, "?") {
	// 		url += "&webid=" + web
	// 	} else {
	// 		url += "?webid=" + web
	// 	}
	// }
	permisison, perr := NewAuth.IsGranted("Menu", auth.Create, TenantId)
	if perr != nil {
		ErrorLog.Printf("Update Menu authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Update Menu authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	_, err := MenuConfig.UpdateMenu(menudetails)
	if strings.Contains(fmt.Sprint(err), "given some values is empty") {
		ErrorLog.Printf("UpdateMenu mandatory field error: %s", err)
		c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
		c.Redirect(301, url)
		return
	}

	if err != nil {
		ErrorLog.Printf("UpdateMenu error: %s", perr)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.Redirect(http.StatusMovedPermanently, url)
		return
	}

	c.SetCookie("get-toast", "Menu Updated Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(http.StatusMovedPermanently, url)

}

//Delete Menu function//

func DeleteMenu(c *gin.Context) {

	menuId, _ := strconv.Atoi(c.Param("id"))

	fmt.Println(menuId, "menuiddd")
	pageno := c.Query("page")
	userid := c.GetInt("userid")
	webid := c.Query("webid")

	var url string
	if pageno != "" {
		url = "/admin/website/menu?page=" + pageno
	} else {
		url = "/admin/website/menu"

	}
	if webid != "" && webid != "0" {
		if strings.Contains(url, "?") {
			url += "&webid=" + webid
		} else {
			url += "?webid=" + webid
		}
	}
	permisison, perr := NewAuth.IsGranted("Menu", auth.Delete, TenantId)
	if perr != nil {
		ErrorLog.Printf("delete menu authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Menu authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		err := MenuConfig.DeleteMenu(menuId, userid, TenantId)

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

		c.SetCookie("get-toast", "Menu Deleted Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, url)

	}

}

//Duplicate Menu Name check function//

func CheckMenuName(c *gin.Context) {

	// get value from html form data
	menu_id, _ := strconv.Atoi(c.PostForm("menu_id"))
	menu_name := c.PostForm("menu_name")
	websiteid, _ := strconv.Atoi(c.PostForm("webid"))
	parentmenu_id, _ := strconv.Atoi(c.PostForm("parentmenu_id"))

	permisison, perr := NewAuth.IsGranted("Menu", auth.Read, TenantId)
	if perr != nil {
		ErrorLog.Printf("checkmenugroupname authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("menugroup authorization error")
		c.Redirect(301, "/403-page")
		return
	}
	if permisison {

		flg, err := MenuConfig.CheckMenuName(menu_id, menu_name, parentmenu_id, websiteid, TenantId)
		if err != nil {
			ErrorLog.Printf("checkmenugroupname  error: %s", err)
			json.NewEncoder(c.Writer).Encode(false)
			return
		}

		json.NewEncoder(c.Writer).Encode(flg)
	}

}

//Status change function//

func MenuStatusChange(c *gin.Context) {

	menuId, _ := strconv.Atoi(c.PostForm("id"))

	userid := c.GetInt("userid")
	val, _ := strconv.Atoi(c.Request.PostFormValue("isactive"))

	permisison, perr := NewAuth.IsGranted("Menu", auth.Update, TenantId)
	if perr != nil {
		ErrorLog.Printf("delete menu authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Menu authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		flg, err := MenuConfig.MenuStatusChange(menuId, val, userid, TenantId)

		if err != nil {
			ErrorLog.Printf("menu status change error: %s", err)
			json.NewEncoder(c.Writer).Encode(flg)

		} else {
			json.NewEncoder(c.Writer).Encode(flg)
		}
	}

}
func MenuPublish(c *gin.Context) {

	permisison, perr := NewAuth.IsGranted("Menu", auth.Update, TenantId)
	if perr != nil {
		ErrorLog.Printf("delete menu authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Menu authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	menuid, _ := strconv.Atoi(c.Param("id"))

	userid := c.GetInt("userid")

	// webid := c.Query("webid")

	_, err := MenuConfig.MenuStatusChange(menuid, 1, userid, TenantId)

	if err != nil {

		fmt.Println(err)
	}
	c.SetCookie("get-toast", "Menu Updated Successfully", 3600, "", "", false, false)
	c.Redirect(301, "/admin/website/menu/menuitems/"+c.Param("id"))
}

//Menu Item Creation//

func CreateMenuItem(c *gin.Context) {
	labelname := c.PostForm("menu_name")
	urlpath := c.PostForm("urlpath")
	seperatewindow := c.PostForm("separatewindow")
	parentmenuid, _ := strconv.Atoi(c.PostForm("menu_id"))
	menuGroupType := c.PostForm("parentmenu_grouptype")

	mode := c.PostForm("formmode")

	mtype := c.PostForm("type")

	typeid, _ := strconv.Atoi(c.PostForm("menu_typeid"))

	listingsids := c.PostForm("listingsids")

	categoryids := c.PostForm("categoryids")

	userid := c.GetInt("userid")

	websiteid := c.Query("webid")

	webid, _ := strconv.Atoi(websiteid)

	imagedata := c.PostForm("svgHidden")

	metatitle := c.PostForm("meta_title")

	metadesc := c.PostForm("meta_description")

	metakeywords := c.PostForm("meta_keywords")

	slugname := c.PostForm("slug_name")

	userDetails, err := GetRequestScopedTenantDetails(c)

	if err != nil {
		fmt.Println(err, "usdererr")
	}

	var imageName, imagePath, newimagepath string

	//svgimage upload in s3//
	if imagedata != "" {

		var (
			imageByte []byte
			err       error
		)

		imageName, imagePath, imageByte, err = ConvertBase64toByte(imagedata, "website")
		if err != nil {
			ErrorLog.Printf("convert base 64 to byte error : %s", err)
		}

		imagePath = path.Join(userDetails.S3FolderName, imagePath)

		newimagepath = strings.Replace(imagePath, "+", "%2B", -1)

		uerr := storagecontroller.UploadCropImageS3(imageName, imagePath, imageByte)
		if uerr != nil {
			c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
			c.Redirect(301, "/admin/website/menu/menuitems/"+c.PostForm("menu_id"))
			return
		}
	}
	//end

	if urlpath == "" && mtype == "channel" {

		urlpath = "/" + strings.ToLower(strings.ReplaceAll(slugname, " ", "-"))

	} else if mtype == "listings" {

		urlpath = "/listing/category/" + strings.ToLower(strings.ReplaceAll(labelname, " ", "-"))
		mtype = "listings"

	} else if urlpath == "/entries/" {

		urlpath = "/entries/" + strings.ToLower(strings.ReplaceAll(labelname, " ", "-"))
		mtype = "entries"
	}

	if mtype == "categories" {

		channelinfo, _ := ChannelConfig.GetChannelByCategoryId(typeid)
		urlpath = "/" + channelinfo.SlugName + "/category/" + strings.ToLower(strings.ReplaceAll(labelname, " ", "-"))
		mtype = "categories"
	}
	var mainpareninfo menu.TblTemplatePages
	if mtype == "pages" {

		pageinfo, err := MenuConfig.GetPageById(typeid, TenantId)

		if err != nil {

			fmt.Println(err)
		}

		if pageinfo.ParentId != 0 {

			parentinfo, _ := MenuConfig.GetPageById(pageinfo.ParentId, TenantId)

			if parentinfo.ParentId != 0 {

				mainpareninfo, _ = MenuConfig.GetPageById(parentinfo.ParentId, TenantId)

				urlpath = "/" + mainpareninfo.Slug + "/" + parentinfo.Slug + "/" + strings.ToLower(strings.ReplaceAll(slugname, " ", "-"))
			} else {

				urlpath = "/" + parentinfo.Slug + "/" + strings.ToLower(strings.ReplaceAll(slugname, " ", "-"))
			}
		} else {

			urlpath = "/" + strings.ToLower(strings.ReplaceAll(slugname, " ", "-"))
		}

		mtype = "pages"
	}
	var Checkbox int
	if seperatewindow == "on" {
		Checkbox = 1
	} else {
		Checkbox = 0
	}

	menuitemslist, err := MenuConfig.GetDirectSubMenusByParentID(parentmenuid, TenantId)

	for _, value := range menuitemslist {
		value.OrderIndex = value.OrderIndex + 1

		_, err := MenuConfig.UpdateMenuOrderIndexes(value.OrderIndex, value.Id, parentmenuid, userid, TenantId)
		fmt.Println("err", err)

	}

	menu, err := MenuConfig.CreateMenus(menu.MenuCreate{MenuName: labelname, Status: 1, ParentId: parentmenuid, WebsiteId: webid,
		UrlPath: urlpath, TenantId: TenantId, CreatedBy: userid, Type: mtype, TypeId: typeid, ListingsIds: listingsids,
		CategoryIds: categoryids, ImagePath: newimagepath, ImageName: imageName, MetaTitle: metatitle,
		MetaDescription: metadesc,
		MetaKeywords:    metakeywords, SeperateWindow: Checkbox, OrderIndex: 1, MenuGroup: menuGroupType})
	if strings.Contains(fmt.Sprint(err), "given some values is empty") {
		ErrorLog.Printf("Menu mandatory field error: %s", err)
		if mode == "1" {
			c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
			c.Redirect(301, "/admin/website/menu/menuitems/"+c.PostForm("menu_id"))
			return
		}
		c.JSON(200, false)
	}

	if err != nil {
		ErrorLog.Printf("Menu error: %s", err)

		if mode == "1" {
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			c.Redirect(http.StatusMovedPermanently, "/admin/website/menu/menuitems/"+c.PostForm("redirectid"))
			return
		}
		c.JSON(200, false)

	}
	if mode == "1" {
		c.SetCookie("get-toast", "Menu Created Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(http.StatusMovedPermanently, "/admin/website/menu/menuitems/"+c.PostForm("redirectid"))
	} else {
		c.JSON(200, menu)
	}

}

func UpdateMenuItem(c *gin.Context) {

	labelname := c.PostForm("menu_name")

	urlpath := c.PostForm("urlpath")

	menugroup := c.PostForm("menu_group")

	seperatewindow := c.PostForm("separatewindow")
	fmt.Println(seperatewindow, "seperatewindowskbs")
	if strings.Contains(urlpath, "/listings/") {

		urlpath = "/listings/" + strings.ToLower(strings.ReplaceAll(labelname, " ", "-"))

	}

	if strings.Contains(urlpath, "/categories/") {

		urlpath = "/categories/" + strings.ToLower(strings.ReplaceAll(labelname, " ", "-"))

	}

	// if strings.Contains(urlpath, "/channel/") {

	// 	urlpath = "/channel/" + strings.ToLower(strings.ReplaceAll(labelname, " ", "-"))

	// }
	parentmenuid, _ := strconv.Atoi(c.PostForm("parentmenu_id"))

	redirectid, _ := strconv.Atoi(c.PostForm("redirectid"))

	menuid, _ := strconv.Atoi(c.PostForm("menuitem_id"))

	mtype := c.PostForm("type")

	typeid, _ := strconv.Atoi(c.PostForm("menu_typeid"))

	userid := c.GetInt("userid")

	webid, _ := strconv.Atoi(c.PostForm("webid"))

	imagedata := c.PostForm("svgHidden")

	imagedelete := c.PostForm("svgDelete")

	listingsids := c.PostForm("listingsids")

	categoryids := c.PostForm("categoryids")

	metatitle := c.PostForm("meta_title")

	metadesc := c.PostForm("meta_description")

	metakeywords := c.PostForm("meta_keywords")

	editorContent := c.PostForm("editorContent")
	var Checkbox int
	seperatewindow1, _ := strconv.Atoi(seperatewindow)
	if seperatewindow1 == 1 {
		Checkbox = 1
	} else {
		Checkbox = 0
	}

	var imageName, newimagepath string

	var menuinfo menu.TblMenus

	menuinfo, _ = MenuConfig.GetmenyById(menuid, TenantId)
	if imagedelete == "1" {

		imageName = ""
		newimagepath = ""
	} else if imagedata != "" {
		var err error

		imageName, newimagepath, err = ConvertBase64(imagedata, strings.TrimPrefix("storage/website", "/"))

		fmt.Println("imageName", imageName, "newimagepath", newimagepath)

		if err != nil {
			ErrorLog.Printf("error get storage type error: %s", err)
		}

	} else {

		imageName = menuinfo.ImageName
		newimagepath = menuinfo.ImagePath
	}
	metainfo := c.PostForm("metainfo")
	if metainfo == "false" {
		metatitle = menuinfo.MetaTitle

		metadesc = menuinfo.MetaDescription

		metakeywords = menuinfo.MetaKeywords
	}

	menudetails := menu.MenuCreate{
		MenuName:        labelname,
		Status:          1,
		ModifiedBy:      userid,
		TenantId:        TenantId,
		Id:              menuid,
		ParentId:        parentmenuid,
		UrlPath:         urlpath,
		Type:            mtype,
		TypeId:          typeid,
		WebsiteId:       webid,
		ImagePath:       newimagepath,
		ImageName:       imageName,
		ListingsIds:     listingsids,
		CategoryIds:     categoryids,
		MetaTitle:       metatitle,
		MetaDescription: metadesc,
		MetaKeywords:    metakeywords,
		SeperateWindow:  Checkbox,
		Description:     editorContent,
		MenuGroup:       menugroup,
	}

	permisison, perr := NewAuth.IsGranted("Menu", auth.Create, TenantId)
	if perr != nil {
		ErrorLog.Printf("Update Menu authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Update Menu authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	_, err1 := MenuConfig.UpdateMenu(menudetails)

	if err1 != nil {

		c.JSON(200, false)
	}

	c.SetCookie("get-toast", "Menu Updated Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, "/admin/website/menu/menuitems/"+strconv.Itoa(redirectid)+"?webid="+c.PostForm("webid"))

	// c.JSON(200, updatemenu)
}

func DeleteMenuItem(c *gin.Context) {

	menuId, _ := strconv.Atoi(c.Param("id"))

	fmt.Println(menuId, "menuiddd")

	userid := c.GetInt("userid")
	menudeta, _ := MenuConfig.GetmenyById(menuId, TenantId)

	permisison, perr := NewAuth.IsGranted("Menu", auth.Delete, TenantId)
	if perr != nil {
		ErrorLog.Printf("delete menu authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Menu authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	// FIX: Get the current view's menu ID from URL query parameter
	// This ensures we redirect back to the same page the user was on
	redirectId := c.Query("redirectid")
	if redirectId == "" {
		// Fallback: use the item's immediate parent
		if menudeta.ParentId != 0 {
			redirectId = strconv.Itoa(menudeta.ParentId)
		} else {
			redirectId = strconv.Itoa(menudeta.Id)
		}
	}

	if permisison {

		err := MenuConfig.DeleteMenu(menuId, userid, TenantId)

		if strings.Contains(fmt.Sprint(err), "given some values is empty") {
			ErrorLog.Printf("deletemenu mandatory field error: %s", perr)
			c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
			c.Redirect(301, "/admin/website/menu/menuitems/"+redirectId)
			return
		}

		if err != nil {
			ErrorLog.Printf("deletemenu error: %s", perr)
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			c.Redirect(301, "/admin/website/menu/menuitems/"+redirectId)
			return
		}

		c.SetCookie("get-toast", "Menu Deleted Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, "/admin/website/menu/menuitems/"+redirectId)

	}
}

func EditMenuItem(c *gin.Context) {

	menuid, _ := strconv.Atoi(c.Param("id"))

	menuinfo, _ := MenuConfig.GetmenyById(menuid, TenantId)

	c.JSON(200, menuinfo)

}

func UpdateMenuitemOrder(c *gin.Context) {

	var orderData []menu.OrderItem
	err := json.Unmarshal([]byte(c.Request.PostFormValue("orderData")), &orderData)
	if err != nil {
		fmt.Println(err)
	}

	userid := c.GetInt("userid")

	merr := MenuConfig.UpdateMenuItemOrder(orderData, userid, TenantId)

	if merr != nil {
		fmt.Println(merr)
	}

	c.JSON(200, true)
}

// Content MenuItemsList
func MenuIemsListForWebsite(c *gin.Context) {

	menuid, _ := strconv.Atoi(c.Param("id"))

	newmenuitemlist, _ := MenuConfig.GetMenusByParentId(menuid)

	for i := range newmenuitemlist {
		newmenuitemlist[i].Description = StripHTML(newmenuitemlist[i].Description)
	}

	menudetails, _ := MenuConfig.GetmenyByIdForWebsite(menuid)

	ModuleName, TabName, _ := ModuleRouteName(c)

	responseData := gin.H{
		"menudetails": menudetails, "menuitemslist": newmenuitemlist, "title": ModuleName, "Tabmenu": TabName, "Menu": NewMenuController(c), "linktitle": "Menus Items", "menuid": menuid,
	}
	fmt.Println("responseData")
	c.JSON(200, gin.H{
		"status": true,
		"data":   responseData,
	})

}
