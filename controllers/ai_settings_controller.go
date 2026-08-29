package controllers

import (
	"fmt"
	"spurt-cms/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	csrf "github.com/utrack/gin-csrf"
)

func AiSettings(c *gin.Context) {

	var limt int
	var offset int
	var filter models.Filter

	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")

	var searchflag bool
	if filter.Keyword != "" {
		searchflag = true
	} else {
		searchflag = false
	}

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	permisison, perr := NewAuth.IsGranted("AI Settings", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("AI Settings authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("AI Settings authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	list, _, _ := models.ListAiModule(limt, offset, filter, TenantId)

	_, count, err := models.ListAiModule(0, 0, filter, TenantId)
	if err != nil {
		fmt.Println(err)
	}

	var aimodulelist []models.TblAiSettingsModule

	for _, val := range list {

		if !val.ModifiedOn.IsZero() {
			val.DateString = val.ModifiedOn.In(TZONE).Format(Datelayout)
		} else {
			val.DateString = val.CreatedOn.In(TZONE).Format(Datelayout)
		}

		aimodulelist = append(aimodulelist, val)

	}

	//pagination calc
	paginationendcount := len(aimodulelist) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(pageno, int(count), limt)

	menu := NewMenuController(c)

	translate, _ := TranslateHandler(c)

	c.HTML(200, "ai-settings.html", gin.H{"Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3,
	}, "Menu": menu, "linktitle": "AI Module settings", "Searchtrue": searchflag, "title": "AI Module settings", "csrf": csrf.GetToken(c), "HeadTitle": "AI Module settings", "translate": translate, "SettingsHead": true, "list": aimodulelist, "Count": count, "Paginationendcount": paginationendcount, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Filter": filter, "Paginationstartcount": paginationstartcount, "Limit": limt})

}

func CreateAIApiKeySetting(c *gin.Context) {

	aimodule := c.PostForm("aimodule")

	apikey := c.PostForm("apikey")

	desc := c.PostForm("desc")

	aimodel := c.PostForm("apimodel")

	userid := c.GetInt("userid")

	createdon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	permisison, perr := NewAuth.IsGranted("AI Settings", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("AI Settings authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("AI Settings authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	create := models.TblAiSettingsModule{
		AIModule:    aimodule,
		ApiKey:      apikey,
		Description: desc,
		AiModel:     aimodel,
		IsActive:    0,
		CreatedOn:   createdon,
		CreatedBy:   userid,
		IsDeleted:   0,
		TenantId:    TenantId,
	}

	_, err := models.CreateAiModule(create)
	if err != nil {
		fmt.Println(err)
	}

	c.SetCookie("get-toast", "AI Module Created Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, "/admin/settings/aisettings/")

}

func StatusAction(c *gin.Context) {

	userid := c.GetInt("userid")

	id, _ := strconv.Atoi(c.PostForm("id"))

	isactive, _ := strconv.Atoi(c.PostForm("isactive"))

	permisison, perr := NewAuth.IsGranted("AI Settings", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("AI Settings authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("AI Settings authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	value, err := models.StatusChange(id, isactive, userid, TenantId)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{"value": value, "url": "/admin/settings/aisettings/"})
}

func UpdateAIApiKeySetting(c *gin.Context) {

	id, _ := strconv.Atoi(c.PostForm("moduleid"))

	aimodule := c.PostForm("aimodule")

	apikey := c.PostForm("apikey")

	desc := c.PostForm("desc")

	aimodel := c.PostForm("apimodel")

	userid := c.GetInt("userid")

	modifiedon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	permisison, perr := NewAuth.IsGranted("AI Settings", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("AI Settings authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("AI Settings authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	update := models.TblAiSettingsModule{
		Id:          id,
		AIModule:    aimodule,
		ApiKey:      apikey,
		Description: desc,
		AiModel:     aimodel,
		ModifiedBy:  userid,
		ModifiedOn:  modifiedon,
		TenantId:    TenantId,
	}

	_, err := models.UpdateAiModule(update)
	if err != nil {
		fmt.Println(err)
	}

	c.SetCookie("get-toast", "AI Module Updated Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, "/admin/settings/aisettings/")

}

func DeleteModule(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	pageno := c.Query("page")

	userid := c.GetInt("userid")

	var url string

	permisison, perr := NewAuth.IsGranted("AI Settings", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("AI Settings authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("AI Settings authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	_, err := models.DeleteModule(id, userid, TenantId)
	if err != nil {
		fmt.Println(err)
	}

	_, count, _ := models.ListAiModule(0, 0, models.Filter{}, TenantId)

	if pageno != "" {
		page, _ := strconv.Atoi(pageno)
		page = page - 1
		multi := page * 10
		multiInt64 := int64(multi)
		if count > multiInt64 {
			url = "/admin/settings/aisettings/?page=" + pageno
		} else {
			pagee, _ := strconv.Atoi(pageno)
			_page := pagee - 1
			pages := strconv.Itoa(_page)
			url = "/admin/settings/aisettings/?page=" + pages
		}
	} else {
		url = "/admin/settings/aisettings/"
	}

	c.SetCookie("get-toast", "AI Module Deleted Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, url)

}

func MultiselectDeleteModule(c *gin.Context) {

	moduleids := c.PostFormArray("moduleids[]")
	pageno := c.PostForm("page")
	userid := c.GetInt("userid")
	var url string

	moduleIntIds := make([]int, len(moduleids))
	for i, id := range moduleids {
		intId, _ := strconv.Atoi(id)
		moduleIntIds[i] = intId
	}

	permisison, perr := NewAuth.IsGranted("AI Settings", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("AI Settings authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("AI Settings authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	value, err := models.MultiSelectDelete(moduleIntIds, userid, TenantId)
	if err != nil {
		fmt.Println(err)
	}

	_, count, _ := models.ListAiModule(0, 0, models.Filter{}, TenantId)

	if pageno != "" {
		page, _ := strconv.Atoi(pageno)
		page = page - 1
		multi := page * 10
		multiInt64 := int64(multi)
		if count > multiInt64 {
			url = "/admin/settings/aisettings/?page=" + pageno
		} else {
			pagee, _ := strconv.Atoi(pageno)
			_page := pagee - 1
			pages := strconv.Itoa(_page)
			url = "/admin/settings/aisettings/?page=" + pages
		}
	} else {
		url = "/admin/settings/aisettings/"
	}

	c.JSON(200, gin.H{"value": value, "url": url})

}
