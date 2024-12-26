package controllers

import (
	"fmt"
	"spurt-cms/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/channels"
	csrf "github.com/utrack/gin-csrf"
)

func ListTemplates(c *gin.Context) {

	var limit, offset int

	keyword := strings.TrimSpace(c.Request.URL.Query().Get("keyword"))

	channelId := c.Request.URL.Query().Get("channel")

	userLimit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	if userLimit == "" {
		limit = Limit

	} else {
		limit, _ = strconv.Atoi(userLimit)
	}

	if pageno != 0 {

		offset = (pageno - 1) * limit
	}

	var (
		templateList       []models.TblTemplates
		templateModuleList []models.TblTemplateModules
		mainCount          int64
		tempCountList      []int
		err                error
		channelDetails     models.TblChannel
	)

	templateModuleList, mainCount, err = models.GetTemplateModuleList(keyword, TenantId, channelId)
	if err != nil {
		ErrorLog.Printf("%v: %v", ErrGettingTemplateModuleList, err)

	}

	channelDetails, err = models.GetChannelDetailsWithTemplateCount(channelId, TenantId)
	if err != nil {
		ErrorLog.Printf("%v: %v", ErrGettingTemplateModuleList, err)

	}

	var imagePath string

	var finalList []models.TblTemplateModules

	for _, templateModule := range templateModuleList {

		tempCountList = append(tempCountList, len(templateModule.Templates))

		moduleName := strings.ToTitle(templateModule.TemplateModuleName)
		moduleSlug := strings.ToTitle(templateModule.TemplateModuleSlug)

		templateModule.TemplateModuleName = moduleName
		templateModule.TemplateModuleSlug = moduleSlug

		var finalTemplateList []models.TblTemplates
		var count = 0
		for _, template := range templateModule.Templates {

			changedTempName := strings.ToTitle(template.TemplateName)
			changedTempSlug := strings.ToTitle(template.TemplateSlug)

			if template.ImagePath != "" {

				imagePath = "/image-resize?name=" + template.ImagePath

			} else {
				imagePath = ""
			}

			count++
			template.ImagePath = imagePath
			template.TemplateName = changedTempName
			template.TemplateSlug = changedTempSlug
			finalTemplateList = append(finalTemplateList, template)
		}

		templateModule.Templates = finalTemplateList
		templateModule.TemplateCount = count
		finalList = append(finalList, templateModule)

	}

	channelList, _, err := ChannelConfig.ListChannel(channels.Channels{Limit: 200, Offset: 0, IsActive: true, TenantId: TenantId})
	if err != nil {
		ErrorLog.Printf("%v: %v", ErrGettingTemplateModuleList, err)

	}

	var finalChannelList []channels.Tblchannel

	for _, chnList := range channelList {

		templateCount, err := models.GetChannelBasedTemplateCount(strconv.Itoa(chnList.Id), TenantId)
		if err != nil {
			ErrorLog.Printf("%v: %v", ErrGettingTemplateModuleList, err)

		}

		chnList.EntriesCount = templateCount.TemplateCount
		finalChannelList = append(finalChannelList, chnList)
	}

	var allZeros bool

	if len(tempCountList) > 0 {
		for _, val := range tempCountList {
			if val != 0 {
				allZeros = false
				break
			}
			allZeros = true
		}
	} else {
		allZeros = true
	}

	var paginationendcount = len(templateList) + offset
	var paginationstartcount = offset + 1
	previous, next, pageCount, page := Pagination(pageno, int(mainCount), limit)

	menu := NewMenuController(c)
	translate, _ := TranslateHandler(c)

	c.HTML(200, "template.html", gin.H{"csrf": csrf.GetToken(c), "linktitle": "Template", "Menu": menu, "translate": translate, "SettingsHead": true, "title": "Template", "TemplateModules": finalList, "Count": mainCount, "IsCountZero": allZeros, "Filter": keyword, "Previous": previous, "Next": next, "PageCount": pageCount, "CurrentPage": pageno, "Page": page, "Limit": limit, "Paginationendcount": paginationendcount, "ChannelList": finalChannelList, "ChannelDetail": channelDetails, "Paginationstartcount": paginationstartcount, "Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   pageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3,
	}})
}

func Templates(c *gin.Context) {

	fmt.Println("Dhanush")
	c.HTML(200, "template.html", gin.H{"Name": "Dhanush"})

}
