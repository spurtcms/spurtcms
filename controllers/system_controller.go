package controllers

import (
	"spurt-cms/models"
	"time"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func Personalization(c *gin.Context) {

	var id = c.GetInt("userid")
	var personalize_data models.TblUserPersonalize
	models.GetPersonalize(&personalize_data, id,TenantId)
	var logo = personalize_data.LogoPath

	menu := NewMenuController(c)
	Folder, File, Media, _, _ := GetMedia()
	translate, _ := TranslateHandler(c)
	selectedtype, _ := GetSelectedType()

	c.HTML(200, "personalize.html", gin.H{"Menu": menu,"linktitle":"Personalize", "title": "Personalize", "csrf": csrf.GetToken(c), "HeadTitle": translate.Setting.Personalize, "translate": translate, "Folder": Folder, "File": File, "Media": Media, "logo": logo, "SettingsHead": true, "Personalizemenu": true, "Tooltiptitle": translate.Setting.Personalizetooltip, "StorageType": selectedtype.SelectedType})
}

func PersonalizationUpdate(c *gin.Context) {

	var (
		id               = c.GetInt("userid")
		cl               = c.PostForm("color-change")
		imagePath        string
		expandimagePath  string
		personalize_data models.TblUserPersonalize
	)

	imagePath = c.PostForm("logo_imgpath")
	expandimagePath = c.PostForm("expandlogo_imgpath")

	err := models.GetPersonalize(&personalize_data, id,TenantId)
	if err != nil {

		personalize_data.UserId = id
		personalize_data.MenuBackgroundColor = cl
		personalize_data.LogoPath = imagePath
		personalize_data.ExpandLogoPath = expandimagePath
		personalize_data.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		err1 := models.CreatePersonalize(&personalize_data)
		if err1 != nil {
			ErrorLog.Printf("create personalize error : %s", err)
		}

	} else {

		personalize_data.MenuBackgroundColor = cl
		personalize_data.LogoPath = imagePath
		personalize_data.ExpandLogoPath = expandimagePath
		personalize_data.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		err2 := models.UpdatePersonalize(&personalize_data, id,TenantId)
		if err2 != nil {
			ErrorLog.Printf("update personalize error : %s", err)
		}

	}

	c.SetCookie("get-toast", "Personalize Updated Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, "/settings/personalize/")
}
