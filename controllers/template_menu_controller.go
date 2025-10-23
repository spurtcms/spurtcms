package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
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
	webid, _ := strconv.Atoi(c.Query("webid"))
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

	Menulist, Total_count, err := MenuConfig.MenuList(limt, offset, menu.Filter(filter), TenantId, webid)

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

	if webid != 0 {
		pages = true

	}

	Websitedet, _ := MenuConfig.GetWebsiteById(webid, TenantId)
	templatedetail, err := MenuConfig.GetTemplateById(Websitedet.TemplateId, TenantId)

	if err != nil {

		fmt.Println(err)
	}

	templatedetail.DateString = templatedetail.CreatedOn.In(TZONE).Format(Datelayout)

	c.HTML(200, "menulist.html", gin.H{"Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3,
	}, "Searchtrue": filterflag, "webbanner": webbanner, "websitedetail": Websitedet, "gotemplatepageheader": pages, "templatedetail": templatedetail, "Count": Total_count, "Limit": limt, "csrf": csrf.GetToken(c), "title": ModuleName, "Tabmenu": TabName, "translate": translate, "Menulist": FinalMenusList, "Menu": NewMenuController(c), "linktitle": "Menus", "Paginationendcount": paginationendcount, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Filter": filter, "Paginationstartcount": paginationstartcount, "HeadTitle": "Menu"})

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

	fmt.Println(recentChannels, "channellistsdf")
	menuid, _ := strconv.Atoi(c.Param("id"))

	menuitemslist, err := MenuConfig.GetMenusByParentid(menuid, TenantId)

	menudetails, _ := MenuConfig.GetmenyById(menuid, TenantId)

	ModuleName, TabName, _ := ModuleRouteName(c)

	translate, _ := TranslateHandler(c)

	webid, _ := strconv.Atoi(c.Query("webid"))

	var pages bool

	if webid != 0 {
		pages = true

	}

	Websitedet, _ := MenuConfig.GetWebsiteById(webid, TenantId)
	templatedetail, err := MenuConfig.GetTemplateById(Websitedet.TemplateId, TenantId)

	if err != nil {

		fmt.Println(err)
	}

	Entrylist, _, _, _ := ChannelConfig.ChannelEntriesList(chn.Entries{}, TenantId)

	pagelist, _, err := MenuConfig.GetTemplatePageList(100, 0, menu.Filter{Status: "Active"}, TenantId, webid)

	listinglist, _, err := ListingConfig.ListingsList(0, 0, listing.Filter{}, TenantId)
	if err != nil {
		fmt.Printf("Unable to fetch Listings List:%v", err)
	}

	c.HTML(200, "menuitems.html", gin.H{"csrf": csrf.GetToken(c), "websitedetail": Websitedet, "Entrylist": Entrylist, "pagelist": pagelist, "listinglist": listinglist, "gotemplatepageheader": pages, "templatedetail": templatedetail, "recentChannels": recentChannels, "recentCourses": recentCourses, "recentForms": recentForms, "menudetails": menudetails, "menuitemslist": menuitemslist, "title": ModuleName, "Tabmenu": TabName, "translate": translate, "Menu": NewMenuController(c), "linktitle": "Menus Items", "channelist": channelist, "courselist": courseslist, "Formlist": Formlist, "menuid": menuid})

}

// CreateMenu
func CreateMenu(c *gin.Context) {

	menuname := c.PostForm("menu_name")

	menudesc := c.PostForm("menu_desc")

	userid := c.GetInt("userid")

	websiteid := c.Query("webid")

	fmt.Println("websiteiddd", websiteid)

	webid, _ := strconv.Atoi(c.Query("webid"))

	menustatus, _ := strconv.Atoi(c.PostForm("menustatus"))

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

	_, err := MenuConfig.CreateMenus(menu.MenuCreate{MenuName: menuname, WebsiteId: webid, Description: menudesc, Status: menustatus, TenantId: TenantId, CreatedBy: userid, ParentId: 0})

	if strings.Contains(fmt.Sprint(err), "given some values is empty") {
		ErrorLog.Printf("Menu mandatory field error: %s", err)
		c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
		c.Redirect(301, "/admin/website/menu?webid="+websiteid)
		return
	}

	if err != nil {
		ErrorLog.Printf("Menu error: %s", perr)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.Redirect(http.StatusMovedPermanently, "/admin/website/menu?webid="+websiteid)
		return
	}

	c.SetCookie("get-toast", "Menu Created Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(http.StatusMovedPermanently, "/admin/website/menu?webid="+websiteid)

}

func UpdateMenu(c *gin.Context) {

	userid := c.GetInt("userid")
	menuid, _ := strconv.Atoi(c.PostForm("menu_id"))
	menustatus, _ := strconv.Atoi(c.PostForm("menustatus"))
	pageno := c.Request.PostFormValue("menupageno")
	menupage := c.PostForm("menupage")
	templateid := c.Query("webid")
	webid, _ := strconv.Atoi(c.Query("webid"))
	menudetails := menu.MenuCreate{
		MenuName:    c.PostForm("menu_name"),
		Description: c.PostForm("menu_desc"),
		Status:      menustatus,
		ModifiedBy:  userid,
		TenantId:    TenantId,
		Id:          menuid,
		WebsiteId:   webid,
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

	if templateid != "" && templateid != "0" {
		if strings.Contains(url, "?") {
			url += "&webid=" + templateid
		} else {
			url += "?webid=" + templateid
		}
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

		flg, err := MenuConfig.CheckMenuName(menu_id, menu_name, websiteid, TenantId)
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

	fmt.Println(menuId, "menuiddd")

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
			ErrorLog.Printf("menu status change error: %s", perr)
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

	webid := c.Query("webid")

	_, err := MenuConfig.MenuStatusChange(menuid, 1, userid, TenantId)

	if err != nil {

		fmt.Println(err)
	}
	c.SetCookie("get-toast", "Menu Updated Successfully", 3600, "", "", false, false)
	c.Redirect(301, "/admin/website/menu/?webid="+webid)
}

//Menu Item Creation//

func CreateMenuItem(c *gin.Context) {

	labelname := c.PostForm("menu_name")

	urlpath := c.PostForm("urlpath")

	parentmenuid, _ := strconv.Atoi(c.PostForm("menu_id"))

	mode := c.PostForm("formmode")

	mtype := c.PostForm("type")

	typeid, _ := strconv.Atoi(c.PostForm("cid"))

	listingsids := c.PostForm("listingsids")

	userid := c.GetInt("userid")

	websiteid := c.Query("webid")

	webid, _ := strconv.Atoi(websiteid)

	if urlpath == "" && mtype == "channel" {

		urlpath = "/channel/" + labelname
	} else if urlpath == "/listings/" {

		urlpath = "/listings/" + strings.ToLower(strings.ReplaceAll(labelname, " ", "-"))
		mtype = "listings"
	}

	menu, err := MenuConfig.CreateMenus(menu.MenuCreate{MenuName: labelname, Status: 1, ParentId: parentmenuid, WebsiteId: webid, UrlPath: urlpath, TenantId: TenantId, CreatedBy: userid, Type: mtype, TypeId: typeid, ListingsIds: listingsids})
	if strings.Contains(fmt.Sprint(err), "given some values is empty") {
		ErrorLog.Printf("Menu mandatory field error: %s", err)
		if mode == "1" {
			c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
			c.Redirect(301, "/admin/website/menu/menuitems/"+c.PostForm("menu_id")+"?webid="+websiteid)
			return
		}
		c.JSON(200, false)
	}

	if err != nil {
		ErrorLog.Printf("Menu error: %s", err)

		if mode == "1" {
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			c.Redirect(http.StatusMovedPermanently, "/admin/website/menu/menuitems/"+c.PostForm("menu_id")+"?webid="+websiteid)
			return
		}
		c.JSON(200, false)

	}
	if mode == "1" {
		c.SetCookie("get-toast", "Menu Created Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(http.StatusMovedPermanently, "/admin/website/menu/menuitems/"+c.PostForm("menu_id")+"?webid="+websiteid)
	} else {
		c.JSON(200, menu)
	}

}

func UpdateMenuItem(c *gin.Context) {

	labelname := c.PostForm("menu_name")

	urlpath := c.PostForm("urlpath")

	urlpath = "/listings/" + strings.ToLower(strings.ReplaceAll(labelname, " ", "-"))

	parentmenuid, _ := strconv.Atoi(c.PostForm("parentmenu_id"))

	menuid, _ := strconv.Atoi(c.PostForm("id"))

	mtype := c.PostForm("type")

	typeid, _ := strconv.Atoi(c.PostForm("cid"))

	userid := c.GetInt("userid")

	webid, _ := strconv.Atoi(c.PostForm("webid"))

	menudetails := menu.MenuCreate{
		MenuName:   labelname,
		Status:     1,
		ModifiedBy: userid,
		TenantId:   TenantId,
		Id:         menuid,
		ParentId:   parentmenuid,
		UrlPath:    urlpath,
		Type:       mtype,
		TypeId:     typeid,
		WebsiteId:  webid,
	}
	fmt.Println("checkupdate", menudetails)
	permisison, perr := NewAuth.IsGranted("Menu", auth.Create, TenantId)
	if perr != nil {
		ErrorLog.Printf("Update Menu authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Update Menu authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	updatemenu, err := MenuConfig.UpdateMenu(menudetails)

	if err != nil {

		c.JSON(200, false)
	}

	c.JSON(200, updatemenu)
}

func DeleteMenuItem(c *gin.Context) {

	menuId, _ := strconv.Atoi(c.Param("id"))

	fmt.Println(menuId, "menuiddd")

	templateid := c.Query("webid")
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

	var parentid int

	fmt.Println("menudetadfd", menudeta.ParentId)

	if menudeta.ParentId != 0 {

		menudet, _ := MenuConfig.GetmenyById(menudeta.ParentId, TenantId)

		if menudet.ParentId == 0 {

			parentid = menudet.Id
		} else {

			parentid = menudet.ParentId
		}

	} else {

		parentid = menudeta.ParentId
	}

	if permisison {

		err := MenuConfig.DeleteMenu(menuId, userid, TenantId)

		// err := MenuConfig.DeleteMenuItem(menuId, userid, TenantId)

		if strings.Contains(fmt.Sprint(err), "given some values is empty") {
			ErrorLog.Printf("deletemenu mandatory field error: %s", perr)
			c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
			c.Redirect(301, "/admin/website/menu/menuitems/"+strconv.Itoa(parentid)+"?webid="+templateid)
			return
		}

		if err != nil {
			ErrorLog.Printf("deletemenu error: %s", perr)
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			c.Redirect(301, "/admin/website/menu/menuitems/"+strconv.Itoa(parentid)+"?webid="+templateid)
			return
		}

		c.SetCookie("get-toast", "Menu Deleted Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, "/admin/website/menu/menuitems/"+strconv.Itoa(parentid)+"?webid="+templateid)

	}
}
