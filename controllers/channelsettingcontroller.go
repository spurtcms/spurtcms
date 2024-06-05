package controllers

import (
	"spurt-cms/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func ChannelSettingView(c *gin.Context) {

	//assign need varibles
	var (
		limit         int
		offset        int
		filter        models.Filter
		templates     []models.TblEmailTemplate
		template      []models.TblEmailTemplate
		templatelists []models.TblEmailTemplate
	)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")

	if c.Query("limit") == "" {
		limit = Limit
	} else {
		limit, _ = strconv.Atoi(c.Query("limit"))
	}

	if page != 0 {
		offset = (page - 1) * limit
	}

	flag := false

	_, Totaltemplates := models.GetTemplateList(&templates, 0, 0, filter, flag)

	models.GetTemplateList(&template, limit, offset, filter, flag)

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
	ModuleName, TabName, moduleid := ModuleRouteName(c)

	templatelist, _ := models.GetTemplatesByModuleId(moduleid)

	c.HTML(200, "channelsettings.html", gin.H{"Pagination": PaginationData{
		NextPage:     page + 1,
		PreviousPage: page - 1,
		TotalPages:   PageCount,
		TwoAfter:     page + 2,
		TwoBelow:     page - 2,
		ThreeAfter:   page + 3,
	}, "Menu": menu, "translate": translate, "Page": Page, "Count": Totaltemplates, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": page, "Limit": limit, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "Filter": filter, "Template": templatelist, "csrf": csrf.GetToken(c), "title": ModuleName, "Emailsmenu": true, "Tooltiptitle": translate.Setting.Emailtooltip, "Tabmenu": TabName, "HeadTitle": translate.Channell.Channels})
}
