package controllers

import (
	"encoding/json"
	"log"
	"spurt-cms/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func MemberSettings(c *gin.Context) {

	menu := NewMenuController(c)

	moulename, TabName, moduleid := ModuleRouteName(c)

	templatelist, _ := models.GetTemplatesByModuleId(moduleid, TenantId)

	membersetttings, _ := models.GetMemberSettings(TenantId)

	Adminroles, _ := NewRoleWP.GetRoleByName(TenantId)

	var Adminroleids []int

	for _, val := range Adminroles {

		Adminroleids = append(Adminroleids, val.Id)

	}

	Adminmember, _ := NewTeamWP.GetAdminRoleUsers(Adminroleids, TenantId)

	translate, _ := TranslateHandler(c)

	c.HTML(200, "membersettings.html", gin.H{"Menu": menu, "linktitle": "Member Settings", "Cmsmenu": true, "title": moulename, "Tabmenu": TabName, "HeadTitle": translate.Memberss.Members, "Templatelist": templatelist, "csrf": csrf.GetToken(c), "Membersettings": membersetttings, "Adminmembers": Adminmember, "translate": translate})
}

func MemberSettingUpdate(c *gin.Context) {

	var updatedetails models.TblMemberSetting
	updatedetails.AllowRegistration, _ = strconv.Atoi(c.PostForm("allowregistration"))
	updatedetails.MemberLogin = strings.ToLower(c.PostForm("memberlogin"))
	updatedetails.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	updatedetails.ModifiedBy = c.GetInt("userid")
	updatedetails.Id, _ = strconv.Atoi(c.PostForm("membersettingid"))
	updatedetails.NotificationUsers = c.PostForm("multiselectuser")

	var templatedata []map[string]string
	if err := json.Unmarshal([]byte(c.Request.PostFormValue("templatestatus")), &templatedata); err != nil {
		log.Println(err)
	}

	userid := c.GetInt("userid")

	var template models.TblEmailTemplate

	for _, val := range templatedata {

		template.Id, _ = strconv.Atoi(val["templateid"])
		template.IsActive, _ = strconv.Atoi(val["status"])
		err1 := models.TemplateStatus(template.Id, template.IsActive, userid, TenantId)
		if err1 != nil {
			log.Println(err1)
		}

	}

	err := models.UpdateMemberSetting(&updatedetails, TenantId)

	if err != nil {

		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)

		return
	}

	c.SetCookie("get-toast", "Member Settings Updated Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, "/member/settings/")

}

func EditMemberSettingTemplate(c *gin.Context) {

	var temp models.TblEmailTemplate

	var id, _ = strconv.Atoi(c.Query("id"))

	models.GetTempdetail(&temp, id, TenantId)

	c.JSON(200, temp)

}

func UpdateMemberSettingTemplate(c *gin.Context) {

	if c.PostForm("tempname") == "" || c.PostForm("tempdesc") == "" || c.PostForm("tempsub") == "" || c.PostForm("temcont") == "" {
		c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
		c.Redirect(301, "/member/settings/")

		return
	}

	var template models.TblEmailTemplate
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

		c.Redirect(301, "/member/settings/")

		return
	}

	c.SetCookie("get-toast", "Templateupdatedsuccessfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, "/member/settings/")

}
