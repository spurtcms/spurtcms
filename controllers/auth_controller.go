package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"spurt-cms/logger"
	"spurt-cms/models"
	storagecontroller "spurt-cms/storage-controller"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	mem "github.com/spurtcms/member"
	csrf "github.com/utrack/gin-csrf"
	"gorm.io/gorm"
)

var Store = sessions.NewCookieStore([]byte("!@#$%"))

var Store1 = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY1")))

func init() {

	ErrorLog = logger.ErrorLog()

	WarnLog = logger.WarnLog()

	// SetInitialGeneralValues()
}

/*login page view*/
func Login(c *gin.Context) {

	session, _ := Store.Get(c.Request, os.Getenv("SESSION_KEY"))
	tkn := session.Values["token"]

	if tkn != "" && tkn != nil {

		tkn := session.Values["token"].(string)
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tkn, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			c.Redirect(http.StatusFound, "/")
			return
		}

		c.Writer.Header().Set("Pragma", "no-cache")

		c.Redirect(301, "/admin/dashboard")

		return
	}

	c.HTML(200, "login.html", gin.H{"csrf": csrf.GetToken(c)})
}

/*forgetpassword page view*/
func ForgetPassword(c *gin.Context) {
	c.HTML(200, "forgotpassword.html", gin.H{"csrf": csrf.GetToken(c)})
}

func NewPassword(c *gin.Context) {

	token := c.Query("token")

	Claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, Claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ErrorLog.Printf("Invalid token change password error: %s", err)
			return
		}
	}

	userid := Claims["user_id"]

	c.HTML(200, "change_password.html", gin.H{"csrf": csrf.GetToken(c), "userid": userid})

}

func SendLinkForForgotPass(c *gin.Context) {

	email := c.PostForm("emailid")

	user, _, err := NewTeam.CheckEmail(email, 0, TenantId)

	id := user.Id
	token, _ := CreateTokenWithExpireTime(id)

	if err != nil {
		ErrorLog.Printf("Send Link forget password error: %s", err)
		c.SetCookie("log-toast", "You are not registered with us", 3600, "", "", false, false)
		c.Redirect(301, "/admin/forgot")
		return
	}

	var url_prefix = os.Getenv("BASE_URL")
	linkedin := os.Getenv("LINKEDIN")
	facebook := os.Getenv("FACEBOOK")
	twitter := os.Getenv("TWITTER")
	youtube := os.Getenv("YOUTUBE")
	insta := os.Getenv("INSTAGRAM")

	data := map[string]interface{}{
		"fname":         user.Username,
		"resetpassword": url_prefix + "admin/change-password?token=" + token,
		"restpassurl":   url_prefix + "admin/change-password",
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
	}

	// fmt.Println(user.Username, "isernameee")
	var wg sync.WaitGroup

	wg.Add(1)

	Chan := make(chan string, 1)

	go ForgetPasswordEmail(Chan, &wg, data, email, "Forgot Password", "1")

	close(Chan)

	c.SetCookie("success", "Reset password email send successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, "/admin/forgot")

}

func SetNewPassword(c *gin.Context) {

	newPassword := c.PostForm("pass")
	confirmPassword := c.PostForm("cpass")

	userid, _ := strconv.Atoi(c.PostForm("userId"))

	if newPassword == confirmPassword {

		_, err := NewTeamWP.ChangeYourPassword(newPassword, userid, TenantId)
		if err != nil {
			ErrorLog.Printf("set new password error: %s", err)
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			return
		}

		c.SetCookie("get-toast", "Password Updated Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, "/admin")
	}
}

/*checklogin*/
func CheckLogin(c *gin.Context) {

	fmt.Println("checklogin")

	//get values from html
	uname := c.PostForm("username")
	remember := c.PostForm("rememberme")

	token, userid, err := NewAuth.Checklogin(uname, c.Request.PostFormValue("pass"), TenantId)

	if err != nil {
		ErrorLog.Printf("CheckLogin error: %s", err)
		c.Redirect(301, "/admin")
		// return
	}

	if err != nil && err.Error() == "user disabled please contact admin" {
		c.SetCookie("log-toast", "This account is inactive please contact the admin", 3600, "", "", false, false)
		c.Redirect(301, "/admin")
		return
	}

	userdata, _, _ := NewTeam.GetUserById(userid, []int{})

	if !userdata.LastLogin.IsZero() {
		Lastactive := userdata.LastLogin.In(TZONE).Format(Datelayout)
		Lastlogin[userdata.Id] = Lastactive
	} else {
		Lastlogin[userdata.Id] = "--"
	}

	if c.Request.PostFormValue("username") == "" || c.Request.PostFormValue("pass") == "" {
		c.Redirect(301, "/admin")
		return
	}

	if gorm.ErrRecordNotFound == err {
		c.SetCookie("log-toast", "You are not registered with us", 3600, "", "", false, false)
		c.Redirect(301, "/admin")
		return
	}

	if strings.Contains(fmt.Sprint(err), "invalid password") {
		c.SetCookie("username", uname, 3600, "", "", false, false)
		c.SetCookie("pass-toast", "Invalid Password", 3600, "", "", false, false)
		// c.Redirect(301, "/")
		return
	}

	if remember == "1" {
		Session1, _ := Store.Get(c.Request, "REMEMBER_ME")
		Session1.Values["rememberme"] = remember
		Session1.Save(c.Request, c.Writer)
	}

	Session, _ := Store.Get(c.Request, os.Getenv("SESSION_KEY"))
	Session.Values["token"] = token
	Session.Save(c.Request, c.Writer)

	c.Redirect(301, "/admin/dashboard")

}

// logout and clear the sessions
func Logout(c *gin.Context) {

	userid := c.GetInt("userid")

	// models.LastLoginActivity(userid, TenantId)

	NewTeamWP.LastLoginActivity(userid, TenantId)

	session, err := Store.Get(c.Request, os.Getenv("SESSION_KEY"))
	if err != nil {
		ErrorLog.Printf("Logout session get error: %s", err)
	}

	session.Values["token"] = ""
	session.Options.MaxAge = -1
	er := session.Save(c.Request, c.Writer)
	if er != nil {
		ErrorLog.Printf("Logout session save error: %s", er)
	}

	c.Writer.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "Thu, 01 Jan 1970 00:00:00 GMT") // date in the past ensures expiration

	// cookieNames := []string{"myspurtcms", "lang", "channelbanner", "blockbanner", "tempbanner", "ctadesc", "ctabanner"}
	// for _, name := range cookieNames {
	// 	c.SetCookie(name, "", -1, "/", "spurtcms.com", false, false)
	// 	c.SetCookie(name, "", -1, "/", ".lvh.me", false, false)
	// }
	c.Writer.Header().Set("Pragma", "no-cache")
	for _, cookie := range c.Request.Cookies() {
		c.SetCookie(cookie.Name, "", -1, "/", "", false, true)
	}
	c.Redirect(301, "/")

}

func LastActive(c *gin.Context) {
	userid := c.GetInt("userid")
	NewTeamWP.LastLoginActivity(userid, TenantId)
}

func CreateTokenWithExpireTime(userid int) (string, error) {

	var err error

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Second * 300).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}
func GetTemplateNamee(c *gin.Context, TemplateName string) string {

	templateName := c.Query("template_name")
	if templateName == "" {
		templateName = TemplateName
	}
	templateName = strings.ReplaceAll(strings.ToLower(templateName), " ", "_")

	return templateName
}
func MyProfiles(c *gin.Context) {

	User, website, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	newmenulist, err := MenuConfig.GetmenusByTenantId(website.Id, User.TenantId)

	if err != nil {
		fmt.Println(err)
	}
	UserDetailsFunction(c)

	AllEntryList, _ := AllEntryList(User.TenantId, website.Id)

	templatedetails, _ := MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	Template := GetTemplateNamee(c, templatedetails.TemplateName)

	seodetail, err := MenuConfig.SeoDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	settingsdetail, err := MenuConfig.SettingsDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/themes/"+Template+"/layouts/_default/my_profiles.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	memberid := c.GetInt("member_id")

	profile, member_details, ProfileName := GetProfileName(c, User.TenantId)

	var member models.TblMembers

	member, err = models.GetMemberById(memberid, User.TenantId)

	if err != nil {
		fmt.Println(err)
	}

	PageHeading := "My Profile"

	websitemenu, _ := c.Cookie("websitemenu")
	if websitemenu == "" {
		websitemenu = "false"
	}

	RenderTemplate(c, tmpl, "my_profiles.html", gin.H{"menulist": newmenulist, "searchlist": AllEntryList, "template_name": c.Query("template_name"), "seodetail": seodetail, "settingsdetail": settingsdetail, "member": member, "ProfileImagePath": member.ProfileImagePath, "profile": profile, "PageHeading": PageHeading, "websitemenu": websitemenu, "memberprofile": member_details, "profilename": ProfileName})

}

func UpdateprofileData(c *gin.Context) {

	var MemberGroupId, _ = strconv.Atoi(c.PostForm("membergroupvalue"))

	memberbid, _ := strconv.Atoi(c.PostForm("memberid"))

	fmt.Println("mem grp id", MemberGroupId)

	imagedata := c.PostForm("crop_data")

	User, _, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	var imageName, imagePath string

	if imagedata != "" {

		fmt.Println("checkprofileimagepath")

		var (
			imageByte []byte
			err       error
		)

		imageName, imagePath, imageByte, err = ConvertBase64toByte(imagedata, "member")
		if err != nil {
			ErrorLog.Printf("convert base 64 to byte error : %s", err)
		}

		imagePath = User.S3FolderName + imagePath

		uerr := storagecontroller.UploadCropImageS3(imageName, imagePath, imageByte)
		if uerr != nil {
			c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
			c.Redirect(301, "/")
			return
		}
	}

	Member := map[string]interface{}{
		"first_name":   c.PostForm("mem_fname"),
		"username":     c.PostForm("mem_name"),
		"email":        c.PostForm("mem_email"),
		"mobile_no":    c.PostForm("mem_mobile"),
		"is_active":    1,
		"storage_type": "aws",
	}

	removeImage := c.PostForm("remove_image")

	if removeImage == "1" {
		Member["profile_image"] = ""
		Member["profile_image_path"] = ""
	} else if imageName != "" && imagePath != "" {

		Member["profile_image"] = imageName
		Member["profile_image_path"] = imagePath
	}

	err = MemberConfigWP.MemberFlexibleUpdate(Member, memberbid, User.Id, User.TenantId)
	if err != nil {
		fmt.Println("update member error:", err)
	}

	c.Redirect(301, "/")

}

func SignUp(c *gin.Context) {

	User, website, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	newmenulist, err := GetMenuItemsListByTenantID(User.TenantId, website.Id)

	if err != nil {
		fmt.Println(err)
	}

	templatedetails, _ := MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	Template := GetTemplateNamee(c, templatedetails.TemplateName)

	AllEntryList, _ := AllEntryList(User.TenantId, website.Id)

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/common/layouts/signup.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	RenderTemplate(c, tmpl, "signup.html", gin.H{"menulist": newmenulist, "csrf": csrf.GetToken(c), "searchlist": AllEntryList, "template_name": c.Query("template_name"), "userid": User.Id})
}

func Signin(c *gin.Context) {

	User, webisite, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}
	newmenulist, err := GetMenuItemsListByTenantID(User.TenantId, webisite.Id)

	if err != nil {
		fmt.Println(err)
	}

	templatedetails, err := MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	Template := GetTemplateNamee(c, templatedetails.TemplateName)

	AllEntryList, _ := AllEntryList(User.TenantId, webisite.Id)

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/common/layouts/signin.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	RenderTemplate(c, tmpl, "signin.html", gin.H{"menulist": newmenulist, "csrf": csrf.GetToken(c), "searchlist": AllEntryList, "template_name": c.Query("template_name")})
}

func ForgotPassword(c *gin.Context) {

	User, webisite, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	newmenulist, err := GetMenuItemsListByTenantID(User.TenantId, webisite.Id)

	if err != nil {
		fmt.Println(err)
	}

	AllEntryList, _ := AllEntryList(User.TenantId, webisite.Id)

	templatedetails, err := MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	Template := GetTemplateNamee(c, templatedetails.TemplateName)

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/common/layouts/forgetPassword.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	RenderTemplate(c, tmpl, "forgetPassword.html", gin.H{"menulist": newmenulist, "csrf": csrf.GetToken(c), "searchlist": AllEntryList, "template_name": c.Query("template_name"), "userid": User.Id})

}

//Myprofile//

func TemplateMyProfile(c *gin.Context) {

	User, webisite, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	newmenulist, err := GetMenuItemsListByTenantID(User.TenantId, webisite.Id)

	if err != nil {
		fmt.Println(err)
	}
	UserDetailsFunction(c)

	memberdet, _ := c.Get("userdetails")

	AllEntryList, _ := AllEntryList(User.TenantId, webisite.Id)

	templatedetails, _ := MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	Template := GetTemplateNamee(c, templatedetails.TemplateName)

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/common/layouts/myprofile.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	RenderTemplate(c, tmpl, "myprofile.html", gin.H{"menulist": newmenulist, "csrf": csrf.GetToken(c), "searchlist": AllEntryList, "template_name": c.Query("template_name"), "memberdet": memberdet})

}

//signin functionlity//

func SignIn(c *gin.Context) {

	User, website, err := GetTenantByHost(c)

	templatedetails, _ := MenuConfig.GetTemplateById(website.TemplateId, User.TenantId)

	Template := GetTemplateNamee(c, templatedetails.TemplateName)

	if err != nil {
		fmt.Println(err)
	}

	emailid := c.PostForm("emailid")

	password := c.PostForm("password")

	token, pass, email, _ := models.CheckMemberLogin(emailid, password, User.Id, User.TenantId)

	redirectURL := "/signin"
	if Template != "" {
		redirectURL += "?template_name=" + (Template)
	}
	if !email {
		c.SetCookie("email-toast", "You are not registered with us", 3600, "", "", false, false)
		c.Redirect(301, redirectURL)
		return
	}
	if !pass {
		c.SetCookie("pass-toast", "Invalid Password", 3600, "", "", false, false)
		// c.SetCookie("email", url.QueryEscape(emailid), 3600, "", "", false, false)
		c.Redirect(301, redirectURL)
		return
	}

	Session, _ := Store.Get(c.Request, "gotemplates")
	Session.Values["token"] = token
	Session.Save(c.Request, c.Writer)

	c.Redirect(301, "/")

}

func CheckNameInUser(c *gin.Context) {

	userid, _ := strconv.Atoi(c.PostForm("id"))

	name := c.PostForm("name")

	User, _, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	flg, err := MemberConfig.CheckNameInMember(userid, name, User.TenantId)

	if err != nil {
		fmt.Printf("User checkname error: %s", err)
		json.NewEncoder(c.Writer).Encode(false)
		return
	}

	json.NewEncoder(c.Writer).Encode(flg)

}

func TemplateCheckEmailInMember(c *gin.Context) {

	userid, _ := strconv.Atoi(c.PostForm("id"))

	email := c.PostForm("email")

	User, _, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	flg, err := MemberConfig.CheckEmailInMember(userid, email, User.TenantId)

	if err != nil {
		fmt.Printf("User check email error: %s", err)
		json.NewEncoder(c.Writer).Encode(false)
		return
	}

	json.NewEncoder(c.Writer).Encode(flg)

}

func NewUserSignUp(c *gin.Context) {

	template_name := c.Query("template_name")

	username := c.PostForm("username")

	email := c.PostForm("email")

	password := c.PostForm("password")

	// fmt.Println("name::", username)
	// fmt.Println("email::", email)
	// fmt.Println("password::", password)

	User, website, err := GetTenantByHost(c)

	templatedetails, _ := MenuConfig.GetTemplateById(website.TemplateId, User.TenantId)

	Template := GetTemplateNamee(c, templatedetails.TemplateName)

	fmt.Println(Template)
	if err != nil {
		fmt.Println(err)
	}

	Member := mem.MemberCreationUpdation{

		FirstName: username,
		Email:     email,
		// GroupId:   MemberGroupId,
		IsActive:  1,
		Username:  username,
		Password:  password,
		CreatedBy: User.Id,
		TenantId:  User.TenantId,
	}

	userdata, Uerr := MemberConfig.CreateMember(Member)

	Memberprofile := mem.MemberprofilecreationUpdation{
		MemberId:    userdata.Id,
		ProfileSlug: username,
		ProfileName: username,
		CreatedBy:   User.Id,
		TenantId:    User.TenantId,
	}

	merr := MemberConfig.CreateMemberProfile(Memberprofile)
	if merr != nil {
		fmt.Printf("memberprofile create error: %s", merr)
	}

	if Uerr != nil {
		fmt.Println(Uerr)
	}

	if template_name != "" {

		c.Redirect(301, "/signin?template_name="+template_name)

	} else if template_name == "" {

		c.Redirect(301, "/signin")

	}

}

func Updateprofile(c *gin.Context) {

	var MemberGroupId, _ = strconv.Atoi(c.PostForm("membergroupvalue"))

	memberbid, _ := strconv.Atoi(c.PostForm("memberid"))

	// fmt.Println("mem grp id", MemberGroupId)

	imagedata := c.PostForm("crop_data")

	User, _, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	var imageName, imagePath string

	if imagedata != "" {

		fmt.Println("checkprofileimagepath")

		var (
			imageByte []byte
			err       error
		)

		imageName, imagePath, imageByte, err = ConvertBase64toByte(imagedata, "member")
		if err != nil {
			ErrorLog.Printf("convert base 64 to byte error : %s", err)
		}

		imagePath = User.S3FolderName + imagePath

		uerr := storagecontroller.UploadCropImageS3(imageName, imagePath, imageByte)
		if uerr != nil {
			c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
			c.Redirect(301, "/myprofile")
			return
		}
	}

	Member := map[string]interface{}{
		"first_name":      c.PostForm("mem_name"),
		"last_name":       c.PostForm("mem_lname"),
		"email":           c.PostForm("mem_email"),
		"mobile_no":       c.PostForm("mem_mobile"),
		"member_group_id": MemberGroupId,
		"is_active":       1,
		"storage_type":    "aws",
	}

	removeImage := c.PostForm("remove_image")

	if removeImage == "1" {
		Member["profile_image"] = ""
		Member["profile_image_path"] = ""
	} else if imageName != "" && imagePath != "" {

		Member["profile_image"] = imageName
		Member["profile_image_path"] = imagePath
	}

	err = MemberConfigWP.MemberFlexibleUpdate(Member, memberbid, User.Id, User.TenantId)
	if err != nil {
		fmt.Println("update member error:", err)
	}

	c.Redirect(301, "/myprofile")

}

// check number in member
func TemplateCheckNumberInMember(c *gin.Context) {

	userid, _ := strconv.Atoi(c.PostForm("id"))

	number := c.PostForm("number")

	User, _, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	flg, err := MemberConfig.CheckNumberInMember(userid, number, User.TenantId)
	if err != nil {
		ErrorLog.Printf("member check number error: %s", err)
		json.NewEncoder(c.Writer).Encode(false)
		return
	}
	json.NewEncoder(c.Writer).Encode(flg)

}

//Change password page//

func TemplateChangePassword(c *gin.Context) {

	token := c.Query("token")

	// fmt.Println("tokendfdfdsffd", token)

	Claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, Claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ErrorLog.Printf("Invalid token change password error: %s", err)
			return
		}
	}

	userid := Claims["user_id"]

	User, webisite, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	newmenulist, err := GetMenuItemsListByTenantID(User.TenantId, webisite.Id)

	if err != nil {
		fmt.Println(err)
	}

	AllEntryList, _ := AllEntryList(User.TenantId, webisite.Id)

	templatedetails, err := MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	Template := GetTemplateNamee(c, templatedetails.TemplateName)

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/common/layouts/changepassword.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	RenderTemplate(c, tmpl, "changepassword.html", gin.H{"userid": userid, "csrf": csrf.GetToken(c), "menulist": newmenulist, "searchlist": AllEntryList, "template_name": c.Query("template_name")})

}

//New Password function//

func SetNewpassword(c *gin.Context) {

	newPassword := c.PostForm("pass")

	confirmPassword := c.PostForm("cpass")

	userid, _ := strconv.Atoi(c.PostForm("userId"))

	User, _, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	if newPassword == confirmPassword {

		err := MemberConfigWP.MemberPasswordUpdate(newPassword, confirmPassword, "", userid, User.Id, User.TenantId)

		if err != nil {
			ErrorLog.Printf("set new password error: %s", err)
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			return
		}

		c.SetCookie("get-toast", "Password Updated Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)

		c.Redirect(301, "/signin")
	}
}

func TemplateLogout(c *gin.Context) {
	// Clear the session first
	// TEMPLATE_SESSION_KEY

	session, err := Store.Get(c.Request, os.Getenv("TEMPLATE_SESSION_KEY"))
	if err == nil {
		fmt.Println("inside if condition")
		session.Values["token"] = ""
		session.Options.MaxAge = -1
		if err := session.Save(c.Request, c.Writer); err != nil {
			ErrorLog.Printf("Logout session save error: %s", err)
		}
	} else {
		fmt.Println("inside else condition")
		ErrorLog.Printf("Logout session get error: %s", err)
	}

	// Set cache control headers before any other output
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")

	c.SetCookie(
		"gotemplates",
		"",
		-1,
		"/",
		"",    // empty domain to cover all subdomains
		false, // secure
		true,  // httpOnly
	)

	// Redirect to the home page
	c.Redirect(http.StatusFound, "/")
}
