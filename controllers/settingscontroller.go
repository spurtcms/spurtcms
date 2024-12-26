package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	storagecontroller "spurt-cms/storage-controller"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/team"
	csrf "github.com/utrack/gin-csrf"
	"golang.org/x/crypto/bcrypt"
)

func SettingView(c *gin.Context) {

	menu := NewMenuController(c)
	// 	fmt.Println("valmenemodule",menu.TblModule)

	// 	for _,val:= range menu.TblModule{
	// 		if val.ModuleName=="Settings"{
	// 			for _,val2:=range val.SubModule{

	// fmt.Println("valmenemodule",val2)
	// 			}
	// 		}
	// 	}

	translate, _ := TranslateHandler(c)

	c.HTML(200, "settings.html", gin.H{"Menu": menu, "linktitle": "Settings", "csrf": csrf.GetToken(c), "translate": translate, "HeadTitle": translate.Settings, "title": "Settings"})

}

func MyProfile(c *gin.Context) {

	users, _, err := NewTeamWP.GetUserById(c.GetInt("userid"), []int{})
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

	if users.ProfileImagePath != "" {

		if users.StorageType == "local" {
			users.ProfileImagePath = "/" + users.ProfileImagePath

		} else if users.StorageType == "aws" {
			users.ProfileImagePath = "/image-resize?name=" + users.ProfileImagePath
		}
	}

	var Name = firstn + lastn
	users.NameString = Name

	Role, err := NewRole.GetRoleById(users.RoleId, 0)
	if err != nil {
		log.Println(err)
	}

	translate, _ := TranslateHandler(c)
	menu := NewMenuController(c)
	Folder, File, Media, _, _ := GetMedia()
	selectedtype, _ := GetSelectedType()

	c.HTML(200, "myaccount.html", gin.H{"Menu": menu, "title": "My Account", "linktitle": "My Account", "csrf": csrf.GetToken(c), "HeadTitle": translate.Myprofile, "translate": translate, "user": users, "rolename": Role.Name, "SettingsHead": true, "Myprofmenu": true, "Tooltiptitle": translate.Setting.Myprofiletooltip, "Folder": Folder, "File": File, "Media": Media, "StorageType": selectedtype.SelectedType})
}

func ChangePassword(c *gin.Context) {

	translate, _ := TranslateHandler(c)

	menu := NewMenuController(c)

	c.HTML(200, "security.html", gin.H{"Menu": menu, "title": "Security", "linktitle": "Security", "csrf": csrf.GetToken(c), "translate": translate, "HeadTitle": translate.Setting.Security, "SettingsHead": true, "Changepassmenu": true, "Tooltiptitle": translate.Setting.Securitytooltip})
}

func UpdateProfile(c *gin.Context) {

	if c.PostForm("user_fname") == "" || c.PostForm("user_email") == "" || c.PostForm("user_mob") == "" || c.PostForm("user_name") == "" {
		c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
		c.Redirect(301, "/settings/myprofile/")
		return
	}

	var (
		imgPath, imgName string
		err              error
	)

	storageType, err := GetSelectedType()
	if err != nil {
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
		c.Redirect(301, "/settings/myprofile/")
		return
	}

	imagedata := c.PostForm("crop_data")

	if strings.Contains(imagedata, "data:image/jpeg;base64") || strings.Contains(imagedata, "data:image/png;base64") || strings.Contains(imagedata, "data:image/svg+xml;base64") {

		tenantDetails, err := NewTeam.GetTenantDetails(TenantId)
		if err != nil {
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
			c.Redirect(301, "/settings/myprofile/")
			return
		}

		if storageType.SelectedType == "aws" {
			var (
				tempString, imageName string
				imageByte             []byte
			)

			imageName, tempString, imageByte, err = ConvertBase64toByte(imagedata, "user")
			if err != nil {
				c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
				c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
				c.Redirect(301, "/settings/myprofile/")
				return
			}

			tempString = tenantDetails.S3FolderName + tempString

			err = storagecontroller.UploadCropImageS3(imageName, tempString, imageByte)
			if err != nil {

				c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
				c.Redirect(301, "/settings/myprofile/")
				return
	
			}

			imgPath = tempString
			imgName = imageName

		} else if storageType.SelectedType == "local" {
			var tempImgPath, imageName string

			imageName, tempImgPath, err = ConvertBase64(imagedata, "storage/user")
			if err != nil {
				c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
				c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
				c.Redirect(301, "/settings/myprofile/")
				return
			}

			imgPath = "/" + tempImgPath
			imgName = imageName

		}

	} else {
		imgPath = imagedata
	}

	userid, _ := strconv.Atoi(c.PostForm("id"))

	Newuser := team.TeamCreate{

		FirstName:        c.PostForm("user_fname"),
		LastName:         c.PostForm("user_lname"),
		Email:            c.PostForm("user_email"),
		MobileNo:         c.PostForm("user_mob"),
		Username:         c.PostForm("user_name"),
		ProfileImage:     imgName,
		ProfileImagePath: imgPath,
		StorageType:      storageType.SelectedType,
	}

	fmt.Println("Dha", Newuser)

	err = NewTeamWP.UpdateMyUser(Newuser, userid, TenantId)
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

	fmt.Println("checkpass")

	userid := c.GetInt("userid")

	pswd := c.PostForm("pass")

	cpswd := c.PostForm("cpass")

	if pswd == "" && cpswd == "" && pswd != cpswd {

		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)

		c.Redirect(301, "/settings/changepassword")

		return

	}

	err, _ := NewTeamWP.ChangeYourPassword(pswd, userid, TenantId)

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

	userdet, _, _ := NewTeamWP.GetUserById(id, []int{})

	fname := userdet.FirstName

	email := userdet.Email

	var url_prefix = os.Getenv("BASE_URL")

	// if os.Getenv("BASE_URL") != "" {

	// 	url_prefix = os.Getenv("BASE_URL")

	// } else {

	// 	url_prefix = os.Getenv("DOMAIN_URL")

	// }

	linkedin := os.Getenv("LINKEDIN")
	facebook := os.Getenv("FACEBOOK")
	twitter := os.Getenv("TWITTER")
	youtube := os.Getenv("YOUTUBE")
	insta := os.Getenv("INSTAGRAM")

	data := map[string]interface{}{

		"fname":         fname,
		"admin_logo":    url_prefix + "public/img/SpurtCMSlogo.png",
		"fb_logo":       url_prefix + "public/img/email-icons/facebook.png",
		"linkedin_logo": url_prefix + "public/img/email-icons/linkedin.png",
		"twitter_logo":  url_prefix + "public/img/email-icons/x.png",
		"youtube_logo":  url_prefix + "public/img/email-icons/youtube.png",
		"insta_log":     url_prefix + "public/img/email-icons/instagram.png",
		"facebook":      facebook,
		"instagram":     insta,
		"youtube":       youtube,
		"linkedin":      linkedin,
		"twitter":       twitter,
		"password":      cpswd,
	}

	var wg sync.WaitGroup

	wg.Add(1)

	Chan := make(chan string, 1)

	go ChangePasswordEmail(Chan, &wg, data, email, "changepassword")

	close(Chan)

}

func CheckPassword(c *gin.Context) {

	var Password bool

	id := c.GetInt("userid")

	pswd := c.PostForm("pass")

	userdet, _, _ := NewTeamWP.GetUserById(id, []int{})

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
