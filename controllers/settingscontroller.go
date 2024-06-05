package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"spurt-cms/models"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/pkgcore/teams"
	csrf "github.com/utrack/gin-csrf"
	"golang.org/x/crypto/bcrypt"
)

var Team teams.TeamAuth

func SettingView(c *gin.Context) {

	menu := NewMenuController(c)

	translate, _ := TranslateHandler(c)

	c.HTML(200, "settings.html", gin.H{"Menu": menu, "csrf": csrf.GetToken(c), "translate": translate, "HeadTitle": translate.Settings, "title": "Settings"})

}

func MyProfile(c *gin.Context) {

	users, err := NewTeam.GetUserById(c.GetInt("userid"))
	if err != nil {
		ErrorLog.Printf("get myprofile error : %s", err)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
	}

	var first = users.FirstName
	var last = users.LastName
	var firstn = strings.ToUpper(first[:1])
	var lastn string
	if users.LastName != "" {
		lastn = strings.ToUpper(last[:1])
	}

	var Name = firstn + lastn
	users.NameString = Name

	Role, err := NewRole.GetRoleById(users.RoleId)
	if err != nil {
		log.Println(err)
	}

	translate, _ := TranslateHandler(c)
	menu := NewMenuController(c)
	Folder, File, Media, _ := GetMedia()
	selectedtype, _ := GetSelectedType()

	c.HTML(200, "myaccount.html", gin.H{"Menu": menu, "title": "My Account", "csrf": csrf.GetToken(c), "HeadTitle": translate.Myprofile, "translate": translate, "user": users, "rolename": Role.Name, "SettingsHead": true, "Myprofmenu": true, "Tooltiptitle": translate.Setting.Myprofiletooltip, "Folder": Folder, "File": File, "Media": Media, "StorageType": selectedtype.SelectedType})
}

func ChangePassword(c *gin.Context) {

	translate, _ := TranslateHandler(c)

	menu := NewMenuController(c)

	c.HTML(200, "security.html", gin.H{"Menu": menu, "title": "Security", "csrf": csrf.GetToken(c), "translate": translate, "HeadTitle": translate.Setting.Security, "SettingsHead": true, "Changepassmenu": true, "Tooltiptitle": translate.Setting.Securitytooltip})
}

func UpdateProfile(c *gin.Context) {

	User.Authority = &AUTH

	if c.PostForm("user_fname") == "" || c.PostForm("user_email") == "" || c.PostForm("user_mob") == "" || c.PostForm("user_name") == "" {
		c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
		c.Redirect(301, "/settings/myprofile/")
		return
	}

	imagedata := c.PostForm("crop_data")

	var (
		cl                   = c.PostForm("color-change")
		imagePath1           string
		expandimagePath      string
		personalize_data     models.TblUserPersonalize
		imageName, imagePath string
	)

	if imagedata != "" {
		imageName, imagePath, _ = ConvertBase64(imagedata, "storage/user")
	}

	userid, _ := strconv.Atoi(c.PostForm("id"))

	Newuser := teams.TeamCreate{

		FirstName:        c.PostForm("user_fname"),
		LastName:         c.PostForm("user_lname"),
		Email:            c.PostForm("user_email"),
		MobileNo:         c.PostForm("user_mob"),
		Username:         c.PostForm("user_name"),
		ProfileImage:     imageName,
		ProfileImagePath: imagePath,
	}

	err := User.UpdateMyUser(Newuser)

	imagePath1 = c.PostForm("logo_imgpath")
	expandimagePath = c.PostForm("expandlogo_imgpath")

	err1 := models.GetPersonalize(&personalize_data, userid)

	if err1 != nil {

		personalize_data.UserId = userid
		personalize_data.MenuBackgroundColor = cl
		personalize_data.LogoPath = imagePath1
		personalize_data.ExpandLogoPath = expandimagePath
		personalize_data.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		err2 := models.CreatePersonalize(&personalize_data)
		if err2 != nil {
			ErrorLog.Printf("update profile error : %s", err2)
		}

	} else {

		personalize_data.MenuBackgroundColor = cl
		personalize_data.LogoPath = imagePath1
		personalize_data.ExpandLogoPath = expandimagePath
		personalize_data.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		err3 := models.UpdatePersonalize(&personalize_data, userid)
		if err3 != nil {
			ErrorLog.Printf("update personalize error : %s", err3)
		}

	}

	if strings.Contains(fmt.Sprint(err), "given some values is empty") {
		c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
		c.Redirect(301, "/settings/myprofile/")
		return
	}

	if err != nil {
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
		c.Redirect(301, "/settings/myprofile/")
		return
	}

	c.SetCookie("get-toast", "My Profile Updated Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, "/settings/myprofile/")

}

func UptPassword(c *gin.Context) {

	Team.Authority = &AUTH

	pswd := c.PostForm("pass")

	cpswd := c.PostForm("cpass")

	if pswd == "" && cpswd == "" && pswd != cpswd {

		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)

		c.Redirect(301, "/settings/changepassword")

		return

	}

	err, _ := Team.ChangeYourPassword(pswd)

	if strings.Contains(fmt.Sprint(err), "given some values is empty") {

		c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)

		c.Redirect(301, "/settings/changepassword")

		return

	} else {

		c.SetCookie("get-toast", "Password Updated Successfully", 3600, "", "", false, false)

		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)

		c.Redirect(301, "/settings/changepassword")

	}
	id := c.GetInt("userid")

	userdet, _ := Team.GetUserDetails(id)

	fname := userdet.FirstName

	email := userdet.Email

	var url_prefix = os.Getenv("BASE_URL")

	// if os.Getenv("BASE_URL") != "" {

	// 	url_prefix = os.Getenv("BASE_URL")

	// } else {

	// 	url_prefix = os.Getenv("DOMAIN_URL")

	// }

	data := map[string]interface{}{

		"fname":         fname,
		"admin_logo":    url_prefix + "public/img/spurtcms.png",
		"fb_logo":       url_prefix + "public/img/facebook.png",
		"linkedin_logo": url_prefix + "public/img/linkedin.png",
		"twitter_logo":  url_prefix + "public/img/twitter.png",
	}

	var wg sync.WaitGroup

	wg.Add(1)

	Chan := make(chan string, 1)

	go ChangePasswordEmail(Chan, &wg, data, email, "Changepassword")

	close(Chan)

}
func CheckPassword(c *gin.Context) {

	var Password bool

	Team.Authority = &AUTH

	id := c.GetInt("userid")

	pswd := c.PostForm("pass")

	userdet, _ := Team.GetUserDetails(id)

	if strings.Contains(fmt.Sprint(userdet), "given some values is empty") {

		json.NewEncoder(c.Writer).Encode(false)

		return
	}

	passerr := bcrypt.CompareHashAndPassword([]byte(userdet.Password), []byte(pswd))

	if passerr == bcrypt.ErrMismatchedHashAndPassword {

		Password = false
	} else {
		Password = true
	}

	c.JSON(200, gin.H{"pass": Password})

}
