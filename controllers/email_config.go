package controllers

import (
	"log"
	"os"
	"spurt-cms/models"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"gorm.io/datatypes"
)

func EmailConfig(c *gin.Context) {

	menu := NewMenuController(c)

	translate, _ := TranslateHandler(c)

	var mail models.TblEmailConfigurations

	err := models.GetMail(&mail)

	if err != nil {
		log.Println(err)
	}
	var email models.StmpConfig

	if mail.StmpConfig != nil {

		email.Mail = mail.StmpConfig["Mail"].(string)
		email.Password = mail.StmpConfig["Password"].(string)
		email.Host = mail.StmpConfig["Host"].(string)
		email.Port = mail.StmpConfig["Port"].(string)

	}

	c.HTML(200, "emailconfig.html", gin.H{"csrf": csrf.GetToken(c), "Smtpdiv": false, "Maildetails": mail, "EnvHost": os.Getenv("MAIL_HOST"), "EnvMail": os.Getenv("MAIL_USERNAME"), "EnvPassword": os.Getenv("MAIL_PASSWORD"), "EnvPort": os.Getenv("MAIL_PORT"), "Email": email, "Menu": menu, "translate": translate, "HeadTitle": translate.Email, "SettingsHead": true, "Emailconfigmenu": true, "title": "Email Configuration"})

}

func CreateEmailConfig(c *gin.Context) {

	//get data form html form data
	mail := c.PostForm("email")
	password := c.PostForm("password")
	host := c.PostForm("host")
	port := c.PostForm("port")
	seltype := c.PostForm("seltype")

	var Email models.TblEmailConfigurations

	Email.Id = 1
	Email.SelectedType = seltype
	Email.StmpConfig = datatypes.JSONMap{"Mail": mail, "Password": password, "Host": host, "Port": port}
	createmail, err := models.UpdateMail(&Email)

	var updateemail models.StmpConfig
	updateemail.Mail = createmail.StmpConfig["Mail"].(string)
	updateemail.Password = createmail.StmpConfig["Password"].(string)
	updateemail.Host = createmail.StmpConfig["Host"].(string)
	updateemail.Port = createmail.StmpConfig["Port"].(string)

	if err != nil {
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
		c.Redirect(301, "/settings/emails/")
		return
	}

	menu := NewMenuController(c)

	translate, _ := TranslateHandler(c)

	c.HTML(200, "emailconfig.html", gin.H{"csrf": csrf.GetToken(c), "Email": updateemail, "Maildetails": createmail, "HeadTitle": translate.Email, "Menu": menu, "translate": translate, "SettingsHead": true, "Emailconfigmenu": true, "title": "Email", "Activecls": "active show", "emailtemp": "hide"})

}
