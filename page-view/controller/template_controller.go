package controller

import (
	"fmt"
	"html/template"
	"os"
	"spurt-cms/controllers"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	forms "github.com/spurtcms/forms-builders"
)

func PageView(c *gin.Context) {

	uuid := c.Param("slugstring")

	arr := strings.Split(uuid, "-")
	entries, _, err := controllers.ChannelConfigWP.EntryPreview(arr[len(arr)-1])
	if err != nil {
		controllers.ErrorLog.Printf("preview  details api error: %s", err)
	}

	content := template.HTML(entries.Description)

	formdata, err := controllers.FormConfig.GetCtaById(entries.CtaId)
	newpath := os.Getenv("BASE_URL")
	if err != nil {
		controllers.ErrorLog.Printf("preview  details api error: %s", err)
	}
	viewurl := os.Getenv("VIEW_BASE_URL") + "/cta/formresponse"
	c.HTML(200, "page-view.html", gin.H{"Entries": content, "Entry": entries, "URL": viewurl, "FormData": formdata.FormData, "FormId": formdata.Id, "Formtitle": formdata.FormTitle, "Userid": formdata.CreatedBy, "uploadurl": newpath})
}

func FormsView(c *gin.Context) {
	uuid := c.Param("id")
	viewurl := os.Getenv("VIEW_BASE_URL") + "/cta/formresponse"
	arr := strings.Split(uuid, "-")
	Forms, err := controllers.FormConfig.FormPreview(arr[len(arr)-1])
	newpath := os.Getenv("BASE_URL")

	if err != nil {
		controllers.ErrorLog.Printf("preview  details api error: %s", err)
	}
	c.HTML(200, "formpreview.html", gin.H{"FormData": Forms.FormData, "URL": viewurl, "FormId": Forms.Id, "Formtitle": Forms.FormTitle, "Userid": Forms.CreatedBy, "uploadurl": newpath})
}

func FormResponse(c *gin.Context) {

	data := c.PostForm("data")
	Id := c.PostForm("Id")
	userid := c.PostForm("UserId")
	fmt.Println("useriduserid:",userid)
	entryidForm := c.PostForm("EntryId")
	entryidQuery := c.Query("id")
	id, _ := strconv.Atoi(Id)
	Userid, _ := strconv.Atoi(userid)

	var Entryid int
	var err error
	fmt.Println("entidform", entryidForm)
	if entryidQuery != "" {
		Entryid, err = strconv.Atoi(entryidQuery)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		Entryid, err = strconv.Atoi(entryidForm)
		if err != nil {
			fmt.Println(err)
		}
	}

	response := forms.TblFormResponse{
		FormId:       id,
		FormResponse: data,
		UserId:       Userid,
		TenantId:     controllers.TenantId,
		EntryId:      Entryid,
	}

	err1 := controllers.FormConfig.CreateFormResponse(response)
	if err1 != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{"status": 1, "message": "form submitted successfully"})

}

func FileNotFound(c *gin.Context) {

	c.HTML(200, "404pageview.html", gin.H{"page not found": true})
}
