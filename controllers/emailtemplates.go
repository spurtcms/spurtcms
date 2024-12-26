package controllers

import (
	"fmt"
	"spurt-cms/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func EmailTemplate(c *gin.Context) {

	var (
		limit, offset int
		filter        models.Filter
		templates     []models.TblEmailTemplate
		template      []models.TblEmailTemplate
		templatelists []models.TblEmailTemplate
	)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	if c.Query("limit") == "" {
		limit = Limit
	} else {
		limit, _ = strconv.Atoi(c.Query("limit"))
	}

	if page != 0 {
		offset = (page - 1) * limit
	}

	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")

	flag := false

	_, Totaltemplates := models.GetTemplateList(&templates, 0, 0, filter, flag, TenantId)

	models.GetTemplateList(&template, limit, offset, filter, flag, TenantId)

	for _, val := range template {
		if !val.ModifiedOn.IsZero() {
			val.DateString = val.ModifiedOn.Format(Datelayout)
		} else {
			val.DateString = val.CreatedOn.Format(Datelayout)
		}
		templatelists = append(templatelists, val)
	}

	paginationendcount := len(template) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(page, int(Totaltemplates), limit)

	menu := NewMenuController(c)
	translate, _ := TranslateHandler(c)

	c.HTML(200, "emailtemplate.html", gin.H{"Pagination": PaginationData{
		NextPage:     page + 1,
		PreviousPage: page - 1,
		TotalPages:   PageCount,
		TwoAfter:     page + 2,
		TwoBelow:     page - 2,
		ThreeAfter:   page + 3,
	}, "Menu": menu, "HeadTitle": translate.Memberss.Email, "translate": translate, "Page": Page, "Count": Totaltemplates, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": page, "Limit": limit, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "Filter": filter, "Template": templatelists, "csrf": csrf.GetToken(c), "SettingsHead": true, "title": "Email Templates", "linktitle": "Email Templates", "Emailsmenu": true, "Tooltiptitle": translate.Setting.Emailtooltip})
}

func EditTemplate(c *gin.Context) {

	var (
		temp models.TblEmailTemplate
		err  error
		id   int
	)

	id, err = strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(500, "")
	}

	models.GetTempdetail(&temp, id, TenantId)

	c.JSON(200, temp)

}

func UpdateTemplate(c *gin.Context) {

	if c.PostForm("tempdesc") == "" || c.PostForm("tempsub") == "" || c.PostForm("temcont") == "" {
		c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
		c.Redirect(301, "/settings/emails/")
		return
	}

	var (
		template models.TblEmailTemplate
		url      string
	)

	fmt.Println("Test")

	pageno := c.PostForm("pageno")
	pathname := c.PostForm("pathname")

	if pathname != "" {
		url = pathname
	} else if pageno != "" {
		url = "/settings/emails?page=" + pageno
	} else {
		url = "/settings/emails/"
	}

	template.TemplateName = c.PostForm("tempname")
	template.TemplateSubject = (c.PostForm("tempsub"))
	template.TemplateMessage = c.PostForm("temcont")
	template.TemplateDescription = c.PostForm("tempdesc")
	template.ModifiedBy = c.GetInt("userid")
	template.Id, _ = strconv.Atoi(c.PostForm("userid"))
	template.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().In(TZONE).Format("2006-01-02 15:04:05"))

	err := models.UpdateTemplate(&template, TenantId)
	if err != nil {
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.Redirect(301, url)
		return
	}

	c.SetCookie("get-toast", "Templateupdatedsuccessfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, url)

}

/* template status*/
func TempIsActive(c *gin.Context) {

	id, _ := strconv.Atoi(c.Request.PostFormValue("id"))
	val, _ := strconv.Atoi(c.Request.PostFormValue("isactive"))

	userid := c.GetInt("userid")

	err := models.TemplateStatus(id, val, userid, TenantId)
	if err != nil {
		c.JSON(500, gin.H{"status": false})
	}

	c.JSON(200, gin.H{"status": true})

}
