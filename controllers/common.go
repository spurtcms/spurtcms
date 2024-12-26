package controllers

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"html/template"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"math"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"path"
	"sort"
	"spurt-cms/lang"
	"spurt-cms/models"
	storagecontroller "spurt-cms/storage-controller"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	"github.com/spurtcms/team"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RouteMod struct {
	Name      string
	RouteName string
}

type PaginationData struct {
	NextPage     int
	PreviousPage int
	CurrectPage  int
	TotalPages   int
	TwoAfter     int
	TwoBelow     int
	ThreeAfter   int
}

type Breadcrum struct {
	ModuleName  string
	RouteName   []RouteMod
	Arraylength int
}

func GenerateEmail(email, subject string, data map[string]interface{}, message string, wg *sync.WaitGroup) error {

	data1 := map[string]interface{}{
		"Body":          template.HTML(message),
		"AdminLogo":     data["admin_logo"].(string),
		"FbLogo":        data["fb_logo"].(string),
		"LinkedinLogo":  data["linkedin_logo"].(string),
		"TwitterLogo":   data["twitter_logo"].(string),
		"YoutubeLogo":   data["youtube_logo"].(string),
		"InstaLogo":     data["insta_log"].(string),
		"FacebookLink":  data["facebook"].(string),
		"InstagramLink": data["instagram"].(string),
		"YoutubeLink":   data["youtube"].(string),
		"LinkedinLink":  data["linkedin"].(string),
		"TwitterLink":   data["twitter"].(string),
		"FirstName":     data["fname"].(string),
	}

	t, err2 := template.ParseFiles("view/email/email-template.html")
	if err2 != nil {
		ErrorLog.Println(err2)
	}

	var tpl bytes.Buffer
	if err1 := t.Execute(&tpl, data1); err1 != nil {
		ErrorLog.Println(err1)
	}

	result := tpl.String()

	defer wg.Done()

	var (
		Mail     string
		Password string
		Host     string
		Port     string
		mail     models.TblEmailConfigurations
	)
	models.GetMail(&mail, TenantId)

	if mail.SmtpConfig != nil && mail.SelectedType == "smtp" {
		fmt.Println("")

		Mail = mail.SmtpConfig["Mail"].(string)
		Password = mail.SmtpConfig["Password"].(string)
		Host = mail.SmtpConfig["Host"].(string)
		Port = mail.SmtpConfig["Port"].(string)
	} else {

		Mail = os.Getenv("MAIL_USERNAME")
		Password = os.Getenv("MAIL_PASSWORD")
		Host = os.Getenv("MAIL_HOST")
		Port = os.Getenv("MAIL_PORT")
		fmt.Println("envvv", Mail, Password, Host, Port)
	}

	contentType := "text/html"
	// Set up the SMTP server configuration.
	// auth := smtp.PlainAuth("", os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), smtpHost)
	auth := smtp.PlainAuth("", Mail, Password, Host)

	// Compose the email.
	emailBody := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: %s; charset=UTF-8\r\n\r\n%s", Mail, email, subject, contentType, result)

	// Connect to the SMTP server and send the email.
	// err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, []byte(emailBody))
	err := smtp.SendMail(Host+":"+Port, auth, Mail, []string{email}, []byte(emailBody))

	fmt.Println("reach", err)

	if err != nil {
		ErrorLog.Printf("Failed to send email error : %s", err)
		return err
	}

	fmt.Println("Email sent successfully to:", email)
	return nil
}

func GenerateOwndeskEmail(email, subject, message string, wg *sync.WaitGroup) error {

	data1 := map[string]interface{}{
		"Body": template.HTML(message),
	}

	t, err2 := template.ParseFiles("view/email/owndesk-template.html")
	if err2 != nil {
		fmt.Println(err2)
	}

	var tpl bytes.Buffer
	if err1 := t.Execute(&tpl, data1); err1 != nil {
		log.Println(err1)
	}

	result := tpl.String()

	defer wg.Done()

	// from := os.Getenv("MAIL_USERNAME")
	var (
		Mail     string
		Password string
		Host     string
		Port     string
		mail     models.TblEmailConfigurations
	)

	models.GetMail(&mail, TenantId)

	if mail.SmtpConfig != nil && mail.SelectedType == "smtp" {

		Mail = mail.SmtpConfig["Mail"].(string)
		Password = mail.SmtpConfig["Password"].(string)
		Host = mail.SmtpConfig["Host"].(string)
		Port = mail.SmtpConfig["Port"].(string)
	} else {
		Mail = os.Getenv("MAIL_USERNAME")
		Password = os.Getenv("MAIL_PASSWORD")
		Host = os.Getenv("MAIL_HOST")
		Port = os.Getenv("MAIL_PORT")
	}

	contentType := "text/html"
	// Set up the SMTP server configuration.
	auth := smtp.PlainAuth("", Mail, Password, Host)

	// Compose the email.
	emailBody := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: %s; charset=UTF-8\r\n\r\n%s", Mail, email, subject, contentType, result)

	// Connect to the SMTP server and send the email.
	err := smtp.SendMail(Host+":"+Port, auth, Mail, []string{email}, []byte(emailBody))

	if err != nil {
		fmt.Println("Failed to send email:", err)
		return err
	} else {
		fmt.Println("Email sent successfully to:", email)
		return nil
	}

}

/*Translation*/
func TranslateHandler(c *gin.Context) (lang.Translation, error) {

	translation, ok := c.MustGet("translation").(lang.Translation)

	if !ok {
		ErrorLog.Println("Translate error tranlatehandler")
		return lang.Translation{}, errors.New("translate error")
	}

	return translation, nil
}

/*give pagination by current page and count of record*/
func Pagination(PageNo, RecordCount, Limit int) (Previous int, Next int, Pagecount int, Page []int) {

	Previous = PageNo - 1
	Next = PageNo + 1

	if RecordCount > Limit {

		Pagecount = int(math.Ceil(float64(RecordCount) / float64(Limit)))

	}
	for i := 1; i <= Pagecount; i++ {
		Page = append(Page, i)
	}
	return Previous, Next, Pagecount, Page
}

func ConvertBase64(imageData string, storagepath string) (imgname string, path string, err error) {

	extEndIndex := strings.Index(imageData, ";base64,")
	base64data := imageData[strings.IndexByte(imageData, ',')+1:]
	var ext = imageData[11:extEndIndex]
	rand_num := strconv.Itoa(int(time.Now().Unix()))
	imageName := "IMG-" + rand_num + "." + ext
	err = os.MkdirAll(storagepath, 0777)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	storagePath := storagepath + "/IMG-" + rand_num + "." + ext
	decode, err := base64.StdEncoding.DecodeString(base64data)

	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	file, err := os.Create(storagePath)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	if _, err := file.Write(decode); err != nil {
		fmt.Println(err)
		return "", "", err
	}

	return imageName, storagePath, err
}

func ConvertBase64toByte(imageData string, storagepath string) (imgname string, path string, imagebyte []byte, err error) {

	extEndIndex := strings.Index(imageData, ";base64,")
	base64data := imageData[strings.IndexByte(imageData, ',')+1:]
	var ext = imageData[11:extEndIndex]
	rand_num := strconv.Itoa(int(time.Now().Unix()))
	imageName := "IMG-" + rand_num + "." + ext
	storagePath := storagepath + "/IMG-" + rand_num + "." + ext
	decode, err := base64.StdEncoding.DecodeString(base64data)

	if err != nil {
		fmt.Println(err)
	}

	return imageName, storagePath, decode, nil
}

func ConvertBase64WithName(imageData string, storagepath string, imagename string) (imgname string, path string, err error) {

	// extEndIndex := strings.Index(imageData, ";base64,")
	base64data := imageData[strings.IndexByte(imageData, ',')+1:]
	// var ext = imageData[11:extEndIndex]
	// rand_num := strconv.Itoa(int(time.Now().Unix()))
	// imageName := "IMG-" + rand_num + "." + ext
	imageName := imagename
	os.MkdirAll(storagepath, 0755)
	storagePath := storagepath + imagename
	decode, err := base64.StdEncoding.DecodeString(base64data)

	if err != nil {
		fmt.Println(err)
	}
	file, err := os.Create(storagePath)
	if err != nil {
		fmt.Println(err)
	}
	if _, err := file.Write(decode); err != nil {
		fmt.Println(err)
	}

	return imageName, storagePath, err
}

func ConvertBase64WithName1(imageData string, imagename string) (imgname string, path string, err error) {

	// extEndIndex := strings.Index(imageData, ";base64,")
	base64data := imageData[strings.IndexByte(imageData, ',')+1:]
	// var ext = imageData[11:extEndIndex]
	// rand_num := strconv.Itoa(int(time.Now().Unix()))
	// imageName := "IMG-" + rand_num + "." + ext
	imageName := imagename
	// os.MkdirAll(storagepath, 0755)
	storagePath := imagename
	decode, err := base64.StdEncoding.DecodeString(base64data)

	if err != nil {
		fmt.Println(err)
	}
	file, err := os.Create(storagePath)
	if err != nil {
		fmt.Println(err)
	}
	if _, err := file.Write(decode); err != nil {
		fmt.Println(err)
	}

	return imageName, storagePath, err
}

func UserCreateMail(wg *sync.WaitGroup, data map[string]interface{}, email, action string) {

	var templates models.TblEmailTemplate
	models.GetTemplates(&templates, "createuser", TenantId)

	sub := templates.TemplateSubject
	msg := templates.TemplateMessage

	replacer := strings.NewReplacer(
		"{FirstName}", data["fname"].(string),
		"{UserName}", data["uname"].(string),
		"{Password}", data["Pass"].(string),
		"{Loginurl}", data["login_url"].(string),
		"{AdminLogo}", data["admin_logo"].(string),
		"{FbLogo}", data["fb_logo"].(string),
		"{LinkedinLogo}", data["linkedin_logo"].(string),
		"{TwitterLogo}", data["twitter_logo"].(string),
		"{InstaLogo}", data["insta_log"].(string),
		"{YoutubeLogo}", data["youtube_logo"].(string),
		"{FacebookLink}", data["facebook"].(string),
		"{InstagramLink}", data["instagram"].(string),
		"{YoutubeLink}", data["youtube"].(string),
		"{LinkedinLink}", data["linkedin"].(string),
		"{TwitterLink}", data["twitter"].(string),
	)

	msg = replacer.Replace(msg)
	err := GenerateEmail(email, sub, data, msg, wg)
	if err != nil {
		ErrorLog.Printf("Cann't send UserCreate Email to %s error: %s", data["fname"].(string), err)
	}

}

func SendUserOtp(wg *sync.WaitGroup, data map[string]interface{}, email, action string) {

	var templates models.TblEmailTemplate
	models.GetTemplates(&templates, "OTPGenerate", -1)

	sub := templates.TemplateSubject
	msg := templates.TemplateMessage

	replacer := strings.NewReplacer(
		"{FirstName}", data["firstname"].(string),
		"{OTP}", data["otp"].(string),
		"{Loginurl}", data["login_url"].(string),
		"{AdminLogo}", data["admin_logo"].(string),
		"{FbLogo}", data["fb_logo"].(string),
		"{LinkedinLogo}", data["linkedin_logo"].(string),
		"{TwitterLogo}", data["twitter_logo"].(string),
		"{YoutubeLogo}", data["youtube_logo"].(string),
		"{InstaLogo}", data["insta_log"].(string),
		"{FacebookLink}", data["facebook"].(string),
		"{InstagramLink}", data["instagram"].(string),
		"{YoutubeLink}", data["youtube"].(string),
		"{LinkedinLink}", data["linkedin"].(string),
		"{TwitterLink}", data["twitter"].(string),
		"{Expirytime}", data["expiry"].(string),
	)

	msg = replacer.Replace(msg)
	err := GenerateEmail(email, sub, data, msg, wg)
	if err != nil {
		ErrorLog.Printf("Cann't send OtpCreate Email to %s error:", err)
	}

}

func Loginedsuccessfully(wg *sync.WaitGroup, data map[string]interface{}, email, action string) {
	var templates models.TblEmailTemplate
	models.GetTemplates(&templates, "Logined successfully", -1)

	sub := templates.TemplateSubject
	msg := templates.TemplateMessage

	replacer := strings.NewReplacer(
		"{FirstName}", data["firstname"].(string),
		"{AdminLogo}", data["admin_logo"].(string),
		"{FbLogo}", data["fb_logo"].(string),
		"{LinkedinLogo}", data["linkedin_logo"].(string),
		"{TwitterLogo}", data["twitter_logo"].(string),
		"{YoutubeLogo}", data["youtube_logo"].(string),
		"{InstaLogo}", data["insta_log"].(string),
		"{FacebookLink}", data["facebook"].(string),
		"{InstagramLink}", data["instagram"].(string),
		"{YoutubeLink}", data["youtube"].(string),
		"{LinkedinLink}", data["linkedin"].(string),
		"{TwitterLink}", data["twitter"].(string),
	)

	msg = replacer.Replace(msg)
	err := GenerateEmail(email, sub, data, msg, wg)
	if err != nil {
		ErrorLog.Printf("Cann't login your account Email %s error:", err)
	}
}

func ChangePasswordEmail(Chan chan<- string, wg *sync.WaitGroup, data map[string]interface{}, email, action string) {

	var templates models.TblEmailTemplate
	models.GetTemplates(&templates, "changepassword", TenantId)

	sub := templates.TemplateSubject
	msg := templates.TemplateMessage

	replacer := strings.NewReplacer(
		"{FirstName}", data["fname"].(string),
		"{AdminLogo}", data["admin_logo"].(string),
		"{FbLogo}", data["fb_logo"].(string),
		"{LinkedinLogo}", data["linkedin_logo"].(string),
		"{TwitterLogo}", data["twitter_logo"].(string),
		"{InstaLogo}", data["insta_log"].(string),
		"{YoutubeLogo}", data["youtube_logo"].(string),
		"{FacebookLink}", data["facebook"].(string),
		"{InstagramLink}", data["instagram"].(string),
		"{YoutubeLink}", data["youtube"].(string),
		"{LinkedinLink}", data["linkedin"].(string),
		"{TwitterLink}", data["twitter"].(string),
		"{Password}", data["password"].(string),
	)

	fmt.Println(email, sub, msg, replacer, "alldatamail")
	msg = replacer.Replace(msg)
	GenerateEmail(email, sub, data, msg, wg)
}

func MemberCreateEmail(Chan chan<- string, wg *sync.WaitGroup, data map[string]interface{}, email string, tenantid int, action string) {

	var templates models.TblEmailTemplate
	models.GetTemplates(&templates, "Createmember", tenantid)

	if templates.IsActive == 1 {

		sub := templates.TemplateSubject
		msg := templates.TemplateMessage

		replacer := strings.NewReplacer(
			"{FirstName}", data["fname"].(string),
			"{EmailId}", data["email"].(string),
			"{Password}", data["Pass"].(string),
			"{Loginurl}", data["login_url"].(string),
			"{AdminLogo}", data["admin_logo"].(string),
			"{FbLogo}", data["fb_logo"].(string),
			"{LinkedinLogo}", data["linkedin_logo"].(string),
			"{TwitterLogo}", data["twitter_logo"].(string),
			"{YoutubeLogo}", data["youtube_logo"].(string),
			"{InstaLogo}", data["insta_log"].(string),
			"{ContactEmail}", data["contactemail"].(string),
			"{FacebookLink}", data["facebook"].(string),
			"{InstagramLink}", data["instagram"].(string),
			"{YoutubeLink}", data["youtube"].(string),
			"{LinkedinLink}", data["linkedin"].(string),
			"{TwitterLink}", data["twitter"].(string),
		)

		msg = replacer.Replace(msg)
		GenerateEmail(email, sub, data, msg, wg)
	}

}

func MemberActivationEmail(Chan chan<- string, wg *sync.WaitGroup, data map[string]interface{}, email, action string) {

	fmt.Println("activationmailll")
	var templates models.TblEmailTemplate

	models.GetTemplates(&templates, "memberactivation", TenantId)

	sub := templates.TemplateSubject
	msg := templates.TemplateMessage

	replacer := strings.NewReplacer(
		"{AdminLogo}", data["owndesk_logo"].(string),
		"{FirstName}", data["fname"].(string),
		"{FbLogo}", data["fb_logo"].(string),
		"{LinkedinLogo}", data["linkedin_logo"].(string),
		"{TwitterLogo}", data["twitter_logo"].(string),
		"{YoutubeLogo}", data["youtube_logo"].(string),
		"{InstaLogo}", data["insta_log"].(string),
		"{FacebookLink}", data["facebook"].(string),
		"{InstagramLink}", data["instagram"].(string),
		"{YoutubeLink}", data["youtube"].(string),
		"{LinkedinLink}", data["linkedin"].(string),
		"{TwitterLink}", data["twitter"].(string),
	)

	msg = replacer.Replace(msg)
	GenerateEmail(email, sub, nil, msg, wg)

}
func ForgetPasswordEmail(Chan chan<- string, wg *sync.WaitGroup, data map[string]interface{}, email, action string) {

	var templates models.TblEmailTemplate

	models.GetTemplates(&templates, action, TenantId)

	sub := templates.TemplateSubject

	msg := templates.TemplateMessage

	replacer := strings.NewReplacer(
		"{FirstName}", safeGetString(data, "fname"),
		"{Passwordurl}", safeGetString(data, "resetpassword"),
		"{Passwordurlpath}", safeGetString(data, "restpassurl"),
		"{AdminLogo}", safeGetString(data, "admin_logo"),
		"{FbLogo}", safeGetString(data, "fb_logo"),
		"{LinkedinLogo}", safeGetString(data, "linkedin_logo"),
		"{TwitterLogo}", safeGetString(data, "twitter_logo"),
		"{YoutubeLogo}", safeGetString(data, "youtube_logo"),
		"{InstaLogo}", safeGetString(data, "insta_log"),
		"{FacebookLink}", safeGetString(data, "facebook"),
		"{InstagramLink}", safeGetString(data, "instagram"),
		"{YoutubeLink}", safeGetString(data, "youtube"),
		"{LinkedinLink}", safeGetString(data, "linkedin"),
		"{TwitterLink}", safeGetString(data, "twitter"),
	)
	msg = replacer.Replace(msg)

	fmt.Println(email, sub, data, "emaildetails")
	GenerateEmail(email, sub, data, msg, wg)
}
func safeGetString(data map[string]interface{}, key string) string {

	fmt.Println("safestring")
	if value, exists := data[key]; exists {
		if strValue, ok := value.(string); ok {
			return strValue
		}
		fmt.Printf("Warning: Value for key '%s' is not a string\n", key)
	} else {
		fmt.Printf("Warning: Key '%s' does not exist in data\n", key)
	}
	return "" // Return an empty string if the key doesn't exist or is not a string
}
func OwndeskmemberEmail(Chan chan<- string, wg *sync.WaitGroup, data map[string]interface{}, email, action string) error {

	var templates models.TblEmailTemplate

	models.GetTemplates(&templates, action, TenantId)

	sub := templates.TemplateSubject

	msg := templates.TemplateMessage

	replacer := strings.NewReplacer(
		"{ClientName}", data["client_name"].(string),
		"{CompanyName}", data["company_name"].(string),
		"{Logo}", data["logo"].(string),
		"{Facebook}", data["fb_logo"].(string),
		"{Linkedin}", data["linkedin_logo"].(string),
		"{Spacex}", data["spacex_logo"].(string),
		"{Youtube}", data["youtube_logo"].(string),
		"{Instagram}", data["instagram_logo"].(string),
		"{Link}", data["Link"].(string),
		"{Slug}", data["Slug"].(string),
		"<figure", "<div",
		"</figure", "</div",
		"&nbsp;", "",
	)

	sub = replacer.Replace(sub)
	msg = replacer.Replace(msg)
	return GenerateOwndeskEmail(email, sub, msg, wg)
}

func HashingPassword(pass string) string {

	passbyte, err := bcrypt.GenerateFromPassword([]byte(pass), 14)

	if err != nil {

		panic(err)

	}

	return string(passbyte)
}

/*Function BreadCrumbs*/
// func BreadCrumbs(url string) Breadcrum {

// 	re := regexp.MustCompile(`\d`)

// 	Output := strings.TrimRight(re.ReplaceAllString(url, ""), "/")

// 	var tblmodper models.TblModulePermission

// 	models.GetBreadCrumbsbyURL(&tblmodper, Output)

// 	var parent_id int

// 	parent_id = tblmodper.ParentId

// 	var bc Breadcrum

// 	var rotmod RouteMod

// 	rotmod.Name = tblmodper.BreadcrumbName

// 	rotmod.RouteName = tblmodper.RouteName

// 	bc.RouteName = append(bc.RouteName, rotmod)

// 	for {

// 		if parent_id != 0 {

// 			var parentid models.TblModulePermission

// 			models.GetParentUrl(&parentid, parent_id)

// 			parent_id = parentid.ParentId

// 			bc.ModuleName = tblmodper.ModuleName

// 			var rotmod RouteMod

// 			rotmod.Name = parentid.BreadcrumbName

// 			rotmod.RouteName = parentid.RouteName

// 			bc.RouteName = append(bc.RouteName, rotmod)

// 		} else {

// 			break
// 		}
// 	}

// 	var ArrangeBreadcrumbs Breadcrum

// 	ArrangeBreadcrumbs.Arraylength = len(bc.RouteName) - 1

// 	for i := len(bc.RouteName) - 1; i >= 0; i-- {

// 		ArrangeBreadcrumbs.RouteName = append(ArrangeBreadcrumbs.RouteName, bc.RouteName[i])
// 	}

// 	return ArrangeBreadcrumbs
// }

// Unauthorized
func Unauthorized(c *gin.Context) {

	menu := NewMenuController(c)

	translate, _ := TranslateHandler(c)

	c.HTML(200, "403page.html", gin.H{"Menu": menu, "translate": translate})

}

// 404 page
func FileNotFound(c *gin.Context) {

	menu := NewMenuController(c)

	translate, _ := TranslateHandler(c)

	c.HTML(200, "404page.html", gin.H{"Menu": menu, "translate": translate})

}

func ModuleRouteName(c *gin.Context) (modulename string, tabname string, moduleid int) {

	menu := NewMenuController(c)

	routeName := c.FullPath()

	lastIndex := strings.LastIndex(routeName, ":id")

	if lastIndex != -1 {

		routeName = routeName[:lastIndex]
	}

	if strings.HasPrefix(routeName, "/memberaccess/") {

		routeName = "/memberaccess/"
	}

	if strings.Contains(routeName, "entry") {

		routeName = "/channel/entrylist"
	}

	if strings.Contains(routeName, "/updatemember/") {

		routeName = "/member/"
	}

	if strings.Contains(routeName, "/channels/") {

		routeName = "/channels/"
	}

	if strings.Contains(routeName, "applicants") {

		routeName = "/jobs/applicants/"
	}
	if strings.Contains(routeName, "jobinfo") {

		routeName = "/jobs/"
	}
	if strings.Contains(routeName, "addjob") {

		routeName = "/jobs/"
	}
	if strings.Contains(routeName, "editjob") {

		routeName = "/jobs/"
	}
	if strings.Contains(routeName, "/products") {

		routeName = "/ecommerce/products"
	}

	if strings.Contains(routeName, "/customer") {

		routeName = "/ecommerce/customer"
	}
	if strings.Contains(routeName, "/order") {

		routeName = "/ecommerce/order"
	}

	if strings.Contains(routeName, "/entrylist") {

		routeName = "/channel/entrylist"
	}
	if strings.Contains(routeName, "/unpublishentries") {

		routeName = "/channel/entrylist"
	}
	if strings.Contains(routeName, "/draftentries") {

		routeName = "/channel/entrylist"
	}

	if strings.HasPrefix(routeName, "/aiassistance") {

		routeName = "/aiassistance"
	}

	if strings.HasPrefix(routeName, "/formsbuilder") {

		routeName = "/formsbuilder"
	}

	for _, val := range menu.TblModule {

		for _, val1 := range val.SubModule {

			for _, val2 := range val1.Routes {

				if val2.RouteName == routeName {

					tabname = val1.ModuleName

					modulename = val.ModuleName

					moduleid = val.Id

					break
				}
			}

		}
	}

	return modulename, tabname, moduleid
}

func SaveFileInLocal(c *gin.Context) {

	bytes, err := io.ReadAll(c.Request.Body)

	if err != nil {

		fmt.Println("err :", err)

		c.AbortWithStatus(400)

		return

	}

	data, err := url.ParseQuery(string(bytes))

	if err != nil {

		fmt.Println("err:", err)

		c.AbortWithStatus(400)

		return
	}

	imgFile := data.Get("imgFile")

	imgFileName := data.Get("imgFileName")

	base64data := imgFile[strings.IndexByte(imgFile, ',')+1:]

	path := data.Get("imgPath")

	err = os.MkdirAll(path, os.ModePerm)

	if err != nil {

		fmt.Println("err:", err)

		c.AbortWithStatus(400)

		return
	}

	storagePath := fmt.Sprintf("%s%s", path, imgFileName)

	file, err := os.Create(storagePath)

	if err != nil {

		fmt.Println("err:", err)

		c.AbortWithStatus(400)

		return
	}

	decode, err := base64.StdEncoding.DecodeString(base64data)

	if err != nil {

		fmt.Println("err:", err)

		c.AbortWithStatus(400)

		return
	}

	if _, err := file.Write(decode); err != nil {

		fmt.Println("err:", err)

		c.AbortWithStatus(400)

		return
	}

	c.JSON(200, gin.H{"StoragePath": storagePath})
}

func EditroImageSave(c *gin.Context) {

	var imageName, imagePath string

	imagedata := c.PostForm("imagedata")

	baseurl := os.Getenv("BASE_URL")

	tenantDetails, _ := NewTeam.GetTenantDetails(TenantId)

	if imagedata != "" {

		var (
			imageByte []byte
			err       error
		)

		imageName, imagePath, imageByte, err = ConvertBase64toByte(imagedata, "entries")
		if err != nil {
			ErrorLog.Printf("convert base 64 to byte error : %s", err)
		}

		imagePath = tenantDetails.S3FolderName + imagePath

		uerr := storagecontroller.UploadCropImageS3(imageName, imagePath, imageByte)

		if uerr != nil {

			c.JSON(200, gin.H{"Error": "Invalid or missing S3 credentials. Please verify your configuration and try again."})

			return
		}
		imagePath = baseurl + "image-resize?name=" + imagePath
	}

	c.JSON(200, gin.H{"imagepath": imagePath, "imgname": imageName})

}

func CreateTenantDefaultData(userId, tenantId int) error {

	var err error

	file, err := os.Open("defaults.sql")

	if err != nil {

		return err
	}

	scanner := bufio.NewScanner(file)

	var (
		defChildCatId, defParentCatId, tagId, defChanId, parentKnowledgeBaseId, parentBlogId, parentjobsId, templateModuleId, channelId, moduleId int
		defBlockIds                                                                                                                               []int
		blockDynamicRetrieve                                                                                                                      bool
	)

	for scanner.Scan() {

		query := scanner.Text()

		if len(query) > 0 && !strings.HasPrefix(query, "--") {

			currentTime, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

			parts := strings.Fields(fmt.Sprint(currentTime))

			timeStamp := fmt.Sprintf("%s %s", parts[0], parts[1])

			lowerQuery := strings.ToLower(query)

			switch {

			case strings.Contains(lowerQuery, "tbl_categories") && strings.Contains(lowerQuery, "insert into"):

				err = db.Table("tbl_categories").Where("is_deleted=0 and tenant_id=? and parent_id=0", tenantId).Pluck("id", &defParentCatId).Error

				if err != nil && err != gorm.ErrRecordNotFound {

					return err
				}

				if strings.Contains(lowerQuery, "pid") {

					query = strings.ReplaceAll(query, "pid", strconv.Itoa(defParentCatId))

				}
			case strings.Contains(lowerQuery, "tbl_module_permissions") && strings.Contains(lowerQuery, "insert into"):

				err = db.Table("tbl_channels").Where("is_deleted = 0 and is_active = 1").Pluck("id", &channelId).Error

				if err != nil && err != gorm.ErrRecordNotFound {

					return err
				}

				if strings.Contains(lowerQuery, "rname") {

					query = strings.ReplaceAll(lowerQuery, "rname", fmt.Sprintf("/channel/entrylist/%d ", channelId))
				}

				err = db.Table("tbl_modules").Where(" is_active = 1 and parent_id != 0 and module_name = ?", "Entries").Pluck("id", &moduleId).Error

				if err != nil && err != gorm.ErrRecordNotFound {

					return err
				}

				fmt.Println("mid", moduleId)

				if strings.Contains(query, "mid") {

					query = strings.ReplaceAll(query, "mid", strconv.Itoa(moduleId))
				}

			case strings.Contains(lowerQuery, "tbl_channel_categories") && strings.Contains(lowerQuery, "insert into"):

				if err = db.Table("tbl_categories").Where("is_deleted=0 and tenant_id=? and parent_id=?", tenantId, defParentCatId).Pluck("id", &defChildCatId).Error; err != nil {

					return err
				}

				if err = db.Table("tbl_channels").Where("is_deleted=0 and tenant_id = ?", tenantId).Pluck("id", &defChanId).Error; err != nil {

					return err
				}

				if strings.Contains(lowerQuery, "mapcat") {

					mapCategories := fmt.Sprintf("%v,%v", defParentCatId, defChildCatId)

					query = strings.ReplaceAll(query, "mapcat", mapCategories)

				}

				if strings.Contains(lowerQuery, "chid") {

					query = strings.ReplaceAll(query, "chid", strconv.Itoa(defChanId))
				}
			case strings.Contains(lowerQuery, "tbl_template_modules") && strings.Contains(lowerQuery, "insert into"):

				if err = db.Table("tbl_channels").Where("is_deleted=0 and tenant_id = ?", tenantId).Pluck("id", &defChanId).Error; err != nil {

					return err
				}

				if strings.Contains(lowerQuery, "chid") {
					query = strings.ReplaceAll(lowerQuery, "chid", strconv.Itoa(defChanId))

				}
			case strings.Contains(lowerQuery, "tbl_templates") && strings.Contains(lowerQuery, "insert into"):

				fmt.Println("Enter", lowerQuery)
				if err = db.Debug().Table("tbl_template_modules").Where("is_deleted = 0 and tenant_id = ?", tenantId).Pluck("id", &templateModuleId).Error; err != nil {
					return err

				}

				if strings.Contains(lowerQuery, "tempmoduleid") {
					query = strings.ReplaceAll(lowerQuery, "tempmoduleid", strconv.Itoa(templateModuleId))
				}

				if err = db.Table("tbl_channels").Where("is_deleted=0 and tenant_id = ?", tenantId).Pluck("id", &defChanId).Error; err != nil {

					return err
				}

				if strings.Contains(query, "chid") {
					query = strings.ReplaceAll(query, "chid", strconv.Itoa(defChanId))

				}
			case (strings.Contains(lowerQuery, "tbl_block_tags") || strings.Contains(lowerQuery, "tbl_block_collections")) && strings.Contains(lowerQuery, "insert into") && !blockDynamicRetrieve:

				if err = db.Table("tbl_block_mstr_tags").Select("id").Where("tenant_id=?", tenantId).Pluck("id", &tagId).Error; err != nil {

					return err
				}

				if err = db.Table("tbl_blocks").Select("id").Where("is_deleted=0 and tenant_id=?", tenantId).Scan(&defBlockIds).Error; err != nil {

					return err
				}

				blockDynamicRetrieve = true

			case strings.Contains(lowerQuery, "tbl_graphql_settings") && strings.Contains(lowerQuery, "insert into"):

				var apiKey string

				apiKey, err = GenerateTenantApiToken(64)

				if err != nil {

					return err
				}

				query = strings.ReplaceAll(query, "apikey", apiKey)

			case strings.Contains(lowerQuery, "tbl_categories") && strings.Contains(lowerQuery, "insert into") && strings.Contains(lowerQuery, "Knowledge Base"):

				err = db.Table("tbl_categories").Where("is_deleted = 0 and tenant_id = ? and catgeory_slug = ?", tenantId, "knowledge-base").Pluck("id", &parentKnowledgeBaseId).Error

				if err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

				if strings.Contains(lowerQuery, "pid") {
					query = strings.ReplaceAll(query, "pid", strconv.Itoa(parentKnowledgeBaseId))
				}

			case strings.Contains(lowerQuery, "tbl_categories") && strings.Contains(lowerQuery, "insert into") && (strings.Contains(lowerQuery, "nextjs-starter-theme") || strings.Contains(lowerQuery, "nextjs-blog-theme") || strings.Contains(lowerQuery, "nextjs-blogs2-theme") || strings.Contains(lowerQuery, "nextjs-blogs4-theme")):

				err = db.Table("tbl_categories").Where("is_deleted = 0 and tenant_id = ? and catgeory_slug = ?", tenantId, "blog").Pluck("id", &parentBlogId).Error

				if err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

				if strings.Contains(lowerQuery, "pid") {
					query = strings.ReplaceAll(query, "pid", strconv.Itoa(parentBlogId))
				}

			case strings.Contains(lowerQuery, "tbl_categories") && strings.Contains(lowerQuery, "insert into") && strings.Contains(lowerQuery, "jobs"):

				err = db.Table("tbl_categories").Where("is_deleted = 0 and tenant_id = ? and catgeory_slug = ?", tenantId, "jobs").Pluck("id", &parentjobsId).Error

				if err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

				if strings.Contains(lowerQuery, "pid") {
					query = strings.ReplaceAll(query, "pid", strconv.Itoa(parentjobsId))
				}

			}

			replacer := strings.NewReplacer("uid", strconv.Itoa(userId), "tid", strconv.Itoa(tenantId), "current-time", timeStamp)

			finalQuery := replacer.Replace(query)

			switch {

			default:

				if err = db.Debug().Exec(finalQuery).Error; err != nil {

					fmt.Println("enter err", err)

					return err
				}

			case strings.Contains(lowerQuery, "tbl_block_tags") && strings.Contains(lowerQuery, "insert into"):

				for _, v := range defBlockIds {

					replacer := strings.NewReplacer("blid", strconv.Itoa(v), "tagid", strconv.Itoa(tagId))

					modQuery := replacer.Replace(finalQuery)

					if err = db.Exec(modQuery).Error; err != nil {

						return err
					}
				}

			case strings.Contains(lowerQuery, "tbl_block_collections") && strings.Contains(lowerQuery, "insert into"):

				for _, v := range defBlockIds {

					replacer := strings.NewReplacer("blid", strconv.Itoa(v), "tagid", strconv.Itoa(tagId))

					modQuery := replacer.Replace(finalQuery)

					if err = db.Exec(modQuery).Error; err != nil {

						return err
					}
				}
			}

		}

	}

	return err
}

func GetRequestScopedTenantDetails(c *gin.Context) (team.TblUser, error) {

	userData, exists := c.Get("userDetails")

	if !exists {

		return team.TblUser{}, errors.New("failed to fetch tenant details")
	}

	return userData.(team.TblUser), nil

}

func GenerateTenantApiToken(length int) (string, error) {

	b := make([]byte, length) // Create a slice to hold 32 bytes of random data

	if _, err := rand.Read(b); err != nil { // Fill the slice with random data and handle any errors

		return "", err // Return an empty string and the error if something went wrong
	}

	return base64.URLEncoding.EncodeToString(b), nil // Encode the random bytes to a URL-safe base64 string
}

func UploadFilesToS3(c *gin.Context) {

	path := strings.TrimSuffix(c.PostForm("path"), "/")

	fmt.Println("path", path)

	multipartHeader, err := c.FormFile("file")

	if err != nil {

		fmt.Println("failed to get file")

		c.JSON(200, gin.H{"messgae": "failed to get file", "status": 0, "path": ""})

		return
	}

	file, err := multipartHeader.Open()

	if err != nil {

		fmt.Println("failed to open file")

		c.JSON(200, gin.H{"messgae": "failed to open file", "status": 0, "path": ""})

		return
	}

	path = path + "/"

	err = storagecontroller.UploadFileS3(file, multipartHeader, path)

	if err != nil {

		fmt.Println("failed to upload file to the s3")

		c.JSON(200, gin.H{"messgae": "failed to upload file to the s3", "status": 0, "path": ""})

		return
	}

	c.JSON(200, gin.H{"messgae": "file uploaded sucessfully", "status": 1, "path": path + multipartHeader.Filename})
}

// view image
func ResizeImage(c *gin.Context) {

	fileName := c.Query("name")
	filePath := c.Query("path")
	width, _ := strconv.ParseUint(c.Query("width"), 10, 64)
	height, _ := strconv.ParseUint(c.Query("height"), 10, 64)
	extention := path.Ext(fileName)

	rawObject, err := storagecontroller.GetObjectFromS3(filePath + fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	var buf bytes.Buffer

	buf.ReadFrom(rawObject.Body)

	if extention == ".svg" {
		svgData := buf.String()
		c.Data(http.StatusOK, "image/svg+xml", []byte(svgData))
		return
	}

	Image, _, erri := image.Decode(bytes.NewReader(buf.Bytes()))
	if erri != nil {
		fmt.Println(erri)
		return
	}

	newImage := resize.Resize(uint(width), uint(height), Image, resize.Lanczos3)
	if extention == ".png" {
		png.Encode(c.Writer, newImage)
		return
	}

	if extention == ".jpeg" || extention == ".jpg" {
		_ = jpeg.Encode(c.Writer, newImage, nil)
	}

}

// Media selected type
func GetSelectedType() (storagetyp models.TblStorageType, err error) {

	storagetype, err := models.GetStorageValue(TenantId)

	if err != nil {

		fmt.Println(err)

		return models.TblStorageType{}, err
	}

	return storagetype, nil

}

func GetMedia() (Folders []storagecontroller.Medias, Files []storagecontroller.Medias, TotalMedia []storagecontroller.Medias, Nextcontin string, err error) {

	var (
		MediaErr error
		Media    []storagecontroller.Medias
		Folder   []storagecontroller.Medias
		File     []storagecontroller.Medias
		nextcont string
		search   string
	)

	selectedtype, _ := GetSelectedType()

	if selectedtype.SelectedType == "local" {

		Folder, File, MediaErr = storagecontroller.MediaLocalList(search, "")

		if MediaErr != nil {

			ErrorLog.Println(MediaErr)
		}

	} else if selectedtype.SelectedType == "aws" {

		// resp, _ := storagecontroller.ListS3Bucket()

		resp1, err := storagecontroller.ListS3BucketWithPath("media/")

		nextcont = *resp1.NextContinuationToken

		if err != nil {

			fmt.Println("please check the aws fields")

			return []storagecontroller.Medias{}, []storagecontroller.Medias{}, []storagecontroller.Medias{}, nextcont, err
		}

		Folder, File = MakeFolderandFileArr(resp1, "media/", "")

	}

	sort.SliceStable(Folder, func(i, j int) bool {

		return Folder[i].ModTime.Unix() > Folder[j].ModTime.Unix()
	})

	sort.SliceStable(File, func(i, j int) bool {

		return File[i].ModTime.Unix() > File[j].ModTime.Unix()
	})

	Media = append(Media, Folder...)

	Media = append(Media, File...)

	return Folder, File, Media, nextcont, nil

}

// make folder and file array
func MakeFolderandFileArr(resp1 *s3.ListObjectsV2Output, parentfoldername string, searchfilter string) (folders []storagecontroller.Medias, files []storagecontroller.Medias) {

	var (
		Folder []storagecontroller.Medias
		File   []storagecontroller.Medias
	)

	if searchfilter != "" {

		for _, val := range resp1.CommonPrefixes {

			if strings.Contains(strings.ToLower(RemoveSpecialCharacter(strings.Replace(*val.Prefix, parentfoldername, "", 1))), strings.ToLower(searchfilter)) {

				// obj, _ := storagecontroller.GetObjectFromS3(*val.Prefix)

				var Folde storagecontroller.Medias
				filename := RemoveSpecialCharacter(strings.Replace(*val.Prefix, parentfoldername, "", 1))
				Folde.File = true
				Folde.Path = filename
				Folde.AliaseName = *val.Prefix
				Folde.Name = filename
				// Folde.ModTime = *obj.LastModified
				Folder = append(Folder, Folde)
			}

		}

		for _, val := range resp1.Contents {

			if strings.Contains(strings.ToLower(*val.Key), strings.ToLower(searchfilter)) {

				if strings.Contains(*val.Key, ".jpeg") || strings.Contains(*val.Key, ".png") || strings.Contains(*val.Key, ".jpg") || strings.Contains(*val.Key, ".svg") {

					var file storagecontroller.Medias
					filename := strings.Replace(*val.Key, parentfoldername, "", 1)
					Lastupdate := *val.LastModified
					file.File = false
					file.Name = filename
					file.AliaseName = *val.Key
					file.Path = parentfoldername
					file.ModTime = Lastupdate
					File = append(File, file)

				}
			}

		}

	} else {

		for _, val := range resp1.CommonPrefixes {

			folderPrefix := *val.Prefix
			var folderCount, imageCount int
			// List objects in the current folder
			resp2, err := storagecontroller.ListS3BucketWithPath(folderPrefix)
			if err != nil {
				// Handle error
				continue
			}

			for _, obj := range resp2.Contents {

				if strings.HasSuffix(*obj.Key, "/") {
					// Increment folder count
					folderCount++
				} else {
					// Increment image count
					imageCount++
				}
			}
			folderCount = len(resp2.CommonPrefixes)

			// obj, _ := storagecontroller.GetObjectFromS3(*val.Prefix)

			var Folde storagecontroller.Medias
			filename := RemoveSpecialCharacter(strings.Replace(*val.Prefix, parentfoldername, "", 1))
			Folde.File = true
			Folde.Path = filename
			Folde.AliaseName = *val.Prefix
			Folde.Name = filename
			Folde.TotalSubMedia = folderCount + imageCount

			// Folde.ModTime = *obj.LastModified
			Folder = append(Folder, Folde)

		}

		for _, val := range resp1.Contents {

			if strings.Contains(*val.Key, ".jpeg") || strings.Contains(*val.Key, ".png") || strings.Contains(*val.Key, ".jpg") || strings.Contains(*val.Key, ".svg") {

				var file storagecontroller.Medias
				filename := strings.Replace(*val.Key, parentfoldername, "", 1)
				Lastupdate := *val.LastModified
				file.File = false
				file.Name = filename
				file.AliaseName = *val.Key
				file.Path = parentfoldername
				file.ModTime = Lastupdate
				File = append(File, file)

			}

		}

	}

	return Folder, File
}
