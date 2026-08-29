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

	err := models.GetMail(&mail, TenantId)

	if err != nil {
		log.Println(err)
	}
	var email models.SmtpConfig

	if mail.SmtpConfig != nil {

		email.Mail = mail.SmtpConfig["Mail"].(string)
		email.Password = mail.SmtpConfig["Password"].(string)
		email.Host = mail.SmtpConfig["Host"].(string)
		email.Port = mail.SmtpConfig["Port"].(string)

	}

	c.HTML(200, "emailconfig.html", gin.H{"csrf": csrf.GetToken(c), "Smtpdiv": false, "Maildetails": mail, "EnvHost": os.Getenv("MAIL_HOST"), "EnvMail": os.Getenv("MAIL_USERNAME"), "EnvPassword": os.Getenv("MAIL_PASSWORD"), "EnvPort": os.Getenv("MAIL_PORT"), "Email": email, "Menu": menu, "translate": translate, "HeadTitle": translate.Email, "SettingsHead": true, "Emailconfigmenu": true, "title": "Email Configuration", "linktitle": "Email Configuration"})

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
	Email.SmtpConfig = datatypes.JSONMap{"Mail": mail, "Password": password, "Host": host, "Port": port}
	_, err := models.UpdateMail(&Email, TenantId)

	

	if err != nil {
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
		c.Redirect(301, "/settings/emails/emailconfig/")
		return
	}

	c.SetCookie("get-toast", "Email Configuration Updated Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, "/settings/emails/emailconfig/")

}
