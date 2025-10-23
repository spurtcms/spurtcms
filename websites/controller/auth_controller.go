package templatecontroller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"spurt-cms/controllers"
	"spurt-cms/graphql/controller"
	"spurt-cms/graphql/model"
	"spurt-cms/models"
	storagecontroller "spurt-cms/storage-controller"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	mem "github.com/spurtcms/member"
)

func Login(c *gin.Context) {

	User, website, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	newmenulist, err := GetMenuItemsListByTenantID(User.TenantId, website.Id)

	if err != nil {
		fmt.Println(err)
	}

	templatedetails, _ := controllers.MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	Template := GetTemplateName(c, templatedetails.TemplateName)

	template_name := strings.ToLower(templatedetails.TemplateName)
	template_name = strings.ReplaceAll(template_name, " ", "_")

	tmpl, err := template.ParseFiles("websites/themes/" + Template + "/layouts/_default/template_login.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	controllers.RenderTemplate(c, tmpl, "template_login.html", gin.H{"menulist": newmenulist, "template_name": template_name})
}

func EmailVerification(c *gin.Context) {

	emailid := c.PostForm("emailid")

	User, website, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	var memberdetails models.TblMembers

	member, err := models.GetMemberByEmail(emailid, User.TenantId)

	if err != nil {

		FirstName, LastName := strings.Split(emailid, "@")[0], ""

		createdon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		newMember := models.TblMembers{
			FirstName: FirstName,
			LastName:  LastName,
			Email:     emailid,
			Username:  FirstName,
			IsActive:  1,
			CreatedOn: createdon,
			CreatedBy: User.Id,
			IsDeleted: 0,
			TenantId:  User.TenantId,
		}

		members, _ := models.CreateMember(&newMember)

		Memberprofile := mem.MemberprofilecreationUpdation{
			MemberId:    members.Id,
			ProfileSlug: FirstName,
			ProfileName: FirstName,
			CreatedBy:   User.Id,
			TenantId:    User.TenantId,
		}

		merr := controllers.MemberConfig.CreateMemberProfile(Memberprofile)
		if merr != nil {
			fmt.Println("merr:", merr)
		}

		memberdetails, _ = UpdateMemberOTP(members, User.TenantId)
	} else {

		memberdetails, _ = UpdateMemberOTP(member, User.TenantId)
	}

	Otp := strconv.Itoa(memberdetails.Otp)

	expiry := memberdetails.OtpExpiry.In(controllers.TZONE).Format(controllers.Datelayout)

	linkedin := os.Getenv("LINKEDIN")
	facebook := os.Getenv("FACEBOOK")
	twitter := os.Getenv("TWITTER")
	youtube := os.Getenv("YOUTUBE")
	insta := os.Getenv("INSTAGRAM")

	if memberdetails.IsActive == 1 {

		var url_prefix = os.Getenv("BASE_URL")

		data := map[string]interface{}{

			"firstname":     memberdetails.FirstName,
			"otp":           Otp,
			"login_url":     url_prefix,
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
			"expiry":        expiry,
		}

		var wg sync.WaitGroup
		wg.Add(1)
		go controllers.SendUserOtp(&wg, data, memberdetails.Email, memberdetails.TenantId, "OTPGenerate")

	} else {
		controllers.WarnLog.Println(" Member email notification status not enabled")
	}

	newmenulist, err := GetMenuItemsListByTenantID(User.TenantId, website.Id)

	if err != nil {
		fmt.Println(err)
	}

	templatedetails, _ := controllers.MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	Template := GetTemplateName(c, templatedetails.TemplateName)

	template_name := strings.ToLower(templatedetails.TemplateName)
	template_name = strings.ReplaceAll(template_name, " ", "_")

	tmpl, err := template.ParseFiles("websites/themes/" + Template + "/layouts/_default/template_otpverification.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	controllers.RenderTemplate(c, tmpl, "template_otpverification.html", gin.H{"menulist": newmenulist, "template_name": template_name, "emailid": emailid})
}

func OtpVerification(c *gin.Context) {

	emailid := c.PostForm("emailid")

	otp, _ := strconv.Atoi(c.PostForm("otp"))

	User, _, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	member, _ := models.GetMemberByEmail(emailid, User.TenantId)

	currentTime := time.Now().UTC()

	templatedetails, _ := controllers.MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	Template := GetTemplateName(c, templatedetails.TemplateName)

	template_name := strings.ToLower(templatedetails.TemplateName)
	template_name = strings.ReplaceAll(template_name, " ", "_")

	tmpl, err := template.ParseFiles("websites/themes/" + Template + "/layouts/_default/template_otpverification.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	if member.Otp != otp {

		controllers.RenderTemplate(c, tmpl, "template_otpverification.html", gin.H{"template_name": template_name, "emailid": emailid, "invalid_otp": "Invalid Otp"})
		return
	}

	if currentTime.After(member.OtpExpiry) {

		controllers.RenderTemplate(c, tmpl, "template_otpverification.html", gin.H{"template_name": template_name, "emailid": emailid, "invalid_otp": "Otp Expied"})
		return
	}

	token, _ := CreateToken(member.Id, member.TenantId)

	Session, _ := controllers.Store.Get(c.Request, os.Getenv("TEMPLATE_SESSION_KEY"))
	Session.Values["token"] = token
	Session.Save(c.Request, c.Writer)

	c.Redirect(http.StatusFound, "/")

}

func CreateToken(id int, tenantid string) (string, error) {

	atClaims := jwt.MapClaims{}
	atClaims["member_id"] = id
	atClaims["tenant_id"] = tenantid
	atClaims["expiry_time"] = time.Now().UTC().Add(3 * 30 * 24 * time.Hour)
	atClaims["login_type"] = ""

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func UpdateMemberOTP(member models.TblMembers, tenantid string) (members models.TblMembers, err error) {

	otp := generateOTP()

	var loginmember models.TblMembers

	loginmember.Id = member.Id

	loginmember.Email = member.Email

	loginmember.ModifiedBy = member.Id

	loginmember.TenantId = member.TenantId

	loginmember.Otp, _ = strconv.Atoi(otp)

	ExpirationTime := time.Now().UTC().Add(5 * time.Minute)

	loginmember.OtpExpiry = ExpirationTime

	loginmember.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	merr := models.UpdateMemberOtp(loginmember, tenantid)

	if merr != nil {

		return models.TblMembers{}, merr
	}

	member.Otp, err = strconv.Atoi(otp)

	member.OtpExpiry = ExpirationTime

	if err != nil {

		return models.TblMembers{}, err
	}

	return member, nil
}

func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	otp := fmt.Sprintf("%6d", rand.Intn(900000)+100000)
	return otp
}

func MyProfiles(c *gin.Context) {

	User, website, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	newmenulist, err := GetMenuItemsListByTenantID(User.TenantId, website.Id)

	if err != nil {
		fmt.Println(err)
	}
	UserDetailsFunction(c)

	AllEntryList, _ := AllEntryList(User.TenantId, website.Id)

	templatedetails, _ := controllers.MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	Template := GetTemplateName(c, templatedetails.TemplateName)

	seodetail, err := controllers.MenuConfig.SeoDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	settingsdetail, err := controllers.MenuConfig.SettingsDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/themes/"+Template+"/layouts/_default/my_profiles.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	memberid := c.GetInt("member_id")

	var profile bool

	if memberid != 0 {

		profile = true

	} else {

		profile = false

	}

	var member models.TblMembers

	member, err = models.GetMemberById(memberid, User.TenantId)

	if err != nil {
		fmt.Println(err)
	}

	PageHeading := "My Profile"

	controllers.RenderTemplate(c, tmpl, "my_profiles.html", gin.H{"menulist": newmenulist, "searchlist": AllEntryList, "template_name": c.Query("template_name"), "seodetail": seodetail, "settingsdetail": settingsdetail, "member": member, "ProfileImagePath": member.ProfileImagePath, "profile": profile, "PageHeading": PageHeading})

}

func MyDownloads(c *gin.Context) {

	User, website, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	newmenulist, err := GetMenuItemsListByTenantID(User.TenantId, website.Id)

	if err != nil {
		fmt.Println(err)
	}
	UserDetailsFunction(c)

	AllEntryList, _ := AllEntryList(User.TenantId, website.Id)

	templatedetails, _ := controllers.MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	Template := GetTemplateName(c, templatedetails.TemplateName)

	seodetail, err := controllers.MenuConfig.SeoDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	settingsdetail, err := controllers.MenuConfig.SettingsDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/themes/"+Template+"/layouts/_default/my_downloads.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	memberid := c.GetInt("member_id")

	var profile bool

	if memberid != 0 {

		profile = true

	} else {

		profile = false

	}

	var member models.TblMembers

	member, err = models.GetMemberById(memberid, User.TenantId)

	if err != nil {
		fmt.Println(err)
	}

	PageHeading := "My Downloads"

	controllers.RenderTemplate(c, tmpl, "my_downloads.html", gin.H{"menulist": newmenulist, "searchlist": AllEntryList, "template_name": c.Query("template_name"), "seodetail": seodetail, "settingsdetail": settingsdetail, "ProfileImagePath": member.ProfileImagePath, "profile": profile, "PageHeading": PageHeading})

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

		imageName, imagePath, imageByte, err = controllers.ConvertBase64toByte(imagedata, "member")
		if err != nil {
			controllers.ErrorLog.Printf("convert base 64 to byte error : %s", err)
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

	err = controllers.MemberConfigWP.MemberFlexibleUpdate(Member, memberbid, User.Id, User.TenantId)
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

	templatedetails, _ := controllers.MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	Template := GetTemplateName(c, templatedetails.TemplateName)

	AllEntryList, _ := AllEntryList(User.TenantId, website.Id)

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/common/layouts/signup.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	controllers.RenderTemplate(c, tmpl, "signup.html", gin.H{"menulist": newmenulist, "searchlist": AllEntryList, "template_name": c.Query("template_name"), "userid": User.Id})
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

	templatedetails, err := controllers.MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	Template := GetTemplateName(c, templatedetails.TemplateName)

	AllEntryList, _ := AllEntryList(User.TenantId, webisite.Id)

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/common/layouts/signin.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	controllers.RenderTemplate(c, tmpl, "signin.html", gin.H{"menulist": newmenulist, "searchlist": AllEntryList, "template_name": c.Query("template_name")})
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

	templatedetails, err := controllers.MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	Template := GetTemplateName(c, templatedetails.TemplateName)

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/common/layouts/forgetPassword.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	controllers.RenderTemplate(c, tmpl, "forgetPassword.html", gin.H{"menulist": newmenulist, "searchlist": AllEntryList, "template_name": c.Query("template_name"), "userid": User.Id})

}

//Myprofile//

func MyProfile(c *gin.Context) {

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

	templatedetails, _ := controllers.MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	Template := GetTemplateName(c, templatedetails.TemplateName)

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/common/layouts/myprofile.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	controllers.RenderTemplate(c, tmpl, "myprofile.html", gin.H{"menulist": newmenulist, "searchlist": AllEntryList, "template_name": c.Query("template_name"), "memberdet": memberdet})

}

//signin functionlity//

func SignIn(c *gin.Context) {

	templatename := c.Query("template_name")

	User, _, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	emailid := c.PostForm("emailid")

	password := c.PostForm("password")

	token, pass, email, _ := model.Model.CheckMemberLogin(emailid, password, User.Id, User.TenantId)

	redirectURL := "/signin"
	if templatename != "" {
		redirectURL += "?template_name=" + (templatename)
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

	Session, _ := controllers.Store.Get(c.Request, "spurttemplate")
	Session.Values["token"] = token
	Session.Save(c.Request, c.Writer)
	var retunurl string
	if templatename != "" {
		retunurl += "?template_name=" + (templatename)
	}
	c.Redirect(301, "/"+retunurl)

}

func CheckNameInUser(c *gin.Context) {

	userid, _ := strconv.Atoi(c.PostForm("id"))

	name := c.PostForm("name")

	User, _, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	flg, err := controllers.MemberConfig.CheckNameInMember(userid, name, User.TenantId)

	if err != nil {
		fmt.Printf("User checkname error: %s", err)
		json.NewEncoder(c.Writer).Encode(false)
		return
	}

	json.NewEncoder(c.Writer).Encode(flg)

}

func CheckEmailInMember(c *gin.Context) {

	userid, _ := strconv.Atoi(c.PostForm("id"))

	email := c.PostForm("email")

	User, _, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	flg, err := controllers.MemberConfig.CheckEmailInMember(userid, email, User.TenantId)

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

	fmt.Println("name::", username)
	fmt.Println("email::", email)
	fmt.Println("password::", password)

	User, _, err := GetTenantByHost(c)

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

	userdata, Uerr := controllers.MemberConfig.CreateMember(Member)

	Memberprofile := mem.MemberprofilecreationUpdation{
		MemberId:    userdata.Id,
		ProfileSlug: username,
		ProfileName: username,
		CreatedBy:   User.Id,
		TenantId:    User.TenantId,
	}

	merr := controllers.MemberConfig.CreateMemberProfile(Memberprofile)
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

	templatename := c.Query("template_name")

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

		imageName, imagePath, imageByte, err = controllers.ConvertBase64toByte(imagedata, "member")
		if err != nil {
			controllers.ErrorLog.Printf("convert base 64 to byte error : %s", err)
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

	err = controllers.MemberConfigWP.MemberFlexibleUpdate(Member, memberbid, User.Id, User.TenantId)
	if err != nil {
		fmt.Println("update member error:", err)
	}

	// c.SetCookie("get-toast", "Member Updated Successfully", 3600, "", "", false, false)
	var retunurl string
	if templatename != "" {
		retunurl = "?template_name=" + (templatename)
	}
	c.Redirect(301, "/myprofile"+retunurl)

}

// check number in member
func CheckNumberInMember(c *gin.Context) {

	userid, _ := strconv.Atoi(c.PostForm("id"))

	number := c.PostForm("number")

	User, _, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	flg, err := controllers.MemberConfig.CheckNumberInMember(userid, number, User.TenantId)
	if err != nil {
		controllers.ErrorLog.Printf("member check number error: %s", err)
		json.NewEncoder(c.Writer).Encode(false)
		return
	}
	json.NewEncoder(c.Writer).Encode(flg)

}

//Change password page//

func ChangePassword(c *gin.Context) {

	token := c.Query("token")

	fmt.Println("tokendfdfdsffd", token)

	Claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, Claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			controllers.ErrorLog.Printf("Invalid token change password error: %s", err)
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

	templatedetails, err := controllers.MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)

	Template := GetTemplateName(c, templatedetails.TemplateName)

	tmpl, err := template.ParseFiles("websites/themes/"+Template+"/layouts/partials/header.html", "websites/themes/"+Template+"/layouts/partials/footer.html", "websites/themes/"+Template+"/layouts/partials/head.html", "websites/common/layouts/changepassword.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	controllers.RenderTemplate(c, tmpl, "changepassword.html", gin.H{"userid": userid, "menulist": newmenulist, "searchlist": AllEntryList, "template_name": c.Query("template_name")})

}

//New Password function//

func SetNewpassword(c *gin.Context) {

	templatename := c.Query("template_name")

	newPassword := c.PostForm("pass")

	confirmPassword := c.PostForm("cpass")

	userid, _ := strconv.Atoi(c.PostForm("userId"))

	User, _, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	if newPassword == confirmPassword {

		err := controllers.MemberConfigWP.MemberPasswordUpdate(newPassword, confirmPassword, "", userid, User.Id, User.TenantId)

		if err != nil {
			controllers.ErrorLog.Printf("set new password error: %s", err)
			c.SetCookie("Alert-msg", controllers.ErrInternalServerError, 3600, "", "", false, false)
			return
		}

		c.SetCookie("get-toast", "Password Updated Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		var retunurl string
		if templatename != "" {
			retunurl = "?template_name=" + (templatename)
		}
		c.Redirect(301, "/signin"+retunurl)
	}
}
func UserDetailsFunction(c *gin.Context) {

	session, _ := controllers.Store.Get(c.Request, "spurttemplate")

	tkn := session.Values["token"]
	if tkn != nil {
		token := tkn.(string)

		secret := os.Getenv("JWT_SECRET")

		Claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(token, Claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				fmt.Println(err)

			}

		}

		userIDFloat, ok := Claims["member_id"].(float64)
		if !ok {
			fmt.Println(ok, "err")

		}

		userid := int(userIDFloat)
		tenantid, ok := Claims["tenant_id"].(string)
		if !ok {
			fmt.Println(ok, "err")
			return
		}

		fmt.Printf("userid: %d, tenantid: %d\n", userid)

		member, _ := controllers.MemberConfig.GetMemberAndProfileData(userid, "", 0, "", tenantid)

		var firstn = strings.ToUpper(member.FirstName[:1])

		var lastn string
		if member.LastName != "" {
			lastn = strings.ToUpper(member.LastName[:1])
		}
		member.NameString = firstn + lastn

		member.NameString = firstn + lastn

		fmt.Println(member, userid, tenantid, "checkallvalues")

		if err == nil && member.Id != 0 {
			c.Set("userdetails", member)
		}
	}

}

func Logout(c *gin.Context) {

	templatename := c.Query("template_name")
	var retunurl string
	if templatename != "" {
		retunurl += "?template_name=" + (templatename)
	}
	session, err := controllers.Store.Get(c.Request, "spurttemplate")
	if err != nil {
		controllers.ErrorLog.Printf("Logout session get error: %s", err)
	}

	session.Values["token"] = ""
	session.Options.MaxAge = -1
	er := session.Save(c.Request, c.Writer)
	if er != nil {
		controllers.ErrorLog.Printf("Logout session save error: %s", er)
	}

	c.Writer.Header().Set("Pragma", "no-cache")

	c.SetCookie("spurttemplate", "", -1, "/", ".spurtcms.com", false, false)
	c.SetCookie("spurttemplate", "", -1, "/", ".lvh.me", false, false)

	c.Writer.Header().Set("Pragma", "no-cache")
	for _, cookie := range c.Request.Cookies() {
		c.SetCookie(cookie.Name, "", -1, "/", "", false, true)
	}

	c.Redirect(301, "/"+retunurl)
}

func SendLinkForForgotPass(c *gin.Context) {

	email := c.PostForm("emailid")

	templatename := c.PostForm("template_name")

	var retunurl string
	if templatename != "" {
		retunurl += "?template_name=" + (templatename)
	}

	User, _, err := GetTenantByHost(c)

	if err != nil {
		fmt.Println(err)
	}

	member, _ := controller.MemberInstance.GetMemberAndProfileData(0, email, 0, "", User.TenantId)

	id := member.Id
	token, _ := controllers.CreateTokenWithExpireTime(id)

	var url_prefix = os.Getenv("BASE_URL")
	linkedin := os.Getenv("LINKEDIN")
	facebook := os.Getenv("FACEBOOK")
	twitter := os.Getenv("TWITTER")
	youtube := os.Getenv("YOUTUBE")
	insta := os.Getenv("INSTAGRAM")

	var TemplateName string

	if templatename != "" {

		TemplateName = "&template_name=" + templatename

	}

	data := map[string]interface{}{
		"fname":         member.Username,
		"useremail":     member.Email,
		"resetpassword": "https://" + User.Subdomain + ".spurtcms.com/" + "change-password?token=" + token + TemplateName,
		"restpassurl":   "https://" + User.Subdomain + ".spurtcms.com/" + "change-password" + TemplateName,
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

	var wg sync.WaitGroup

	wg.Add(1)

	Chan := make(chan string, 1)

	go controllers.ForgetPasswordEmail(Chan, &wg, data, email, "Forgot Password", "")

	close(Chan)

	c.SetCookie("success", "Reset password email send successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(302, "/forget-password"+retunurl)

}
