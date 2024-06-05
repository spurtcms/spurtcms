package controllers

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/smtp"
	"os"
	"spurt-cms/lang"
	"spurt-cms/models"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RouteMod struct {
	Name      string
	RouteName string
}

type Breadcrum struct {
	ModuleName  string
	RouteName   []RouteMod
	Arraylength int
}

func GenerateEmail(email, subject, message string, wg *sync.WaitGroup) error {

	data1 := map[string]interface{}{
		"Body": template.HTML(message),
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

	models.GetMail(&mail)

	if mail.StmpConfig != nil && mail.SelectedType == "smtp" {

		Mail = mail.StmpConfig["Mail"].(string)
		Password = mail.StmpConfig["Password"].(string)
		Host = mail.StmpConfig["Host"].(string)
		Port = mail.StmpConfig["Port"].(string)
	} else {
		Mail = os.Getenv("MAIL_USERNAME")
		Password = os.Getenv("MAIL_PASSWORD")
		Host = os.Getenv("MAIL_HOST")
		Port = os.Getenv("MAIL_PORT")
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

	if err != nil {
		ErrorLog.Printf("Failed to send email error : %s", err)
		return err
	}

	fmt.Printf("Email sent successfully to:", email)
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
	from := os.Getenv("MAIL_USERNAME")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	contentType := "text/html"
	// Set up the SMTP server configuration.
	auth := smtp.PlainAuth("", os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), smtpHost)

	// Compose the email.
	emailBody := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: %s; charset=UTF-8\r\n\r\n%s", from, email, subject, contentType, result)

	// Connect to the SMTP server and send the email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, []byte(emailBody))
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
	os.MkdirAll(storagepath, 0755)
	storagePath := storagepath + "/IMG-" + rand_num + "." + ext
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
	models.GetTemplates(&templates, "Createuser")

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
	)

	msg = replacer.Replace(msg)
	err := GenerateEmail(email, sub, msg, wg)
	if err != nil {
		ErrorLog.Printf("Cann't send UserCreate Email to %s error: %s", data["fname"].(string), err)
	}

}

func ChangePasswordEmail(Chan chan<- string, wg *sync.WaitGroup, data map[string]interface{}, email, action string) {

	var templates models.TblEmailTemplate
	models.GetTemplates(&templates, "Changepassword")

	sub := templates.TemplateSubject
	msg := templates.TemplateMessage

	replacer := strings.NewReplacer(
		"{FirstName}", data["fname"].(string),
		"{AdminLogo}", data["admin_logo"].(string),
		"{FbLogo}", data["fb_logo"].(string),
		"{LinkedinLogo}", data["linkedin_logo"].(string),
		"{TwitterLogo}", data["twitter_logo"].(string),
	)

	msg = replacer.Replace(msg)
	GenerateEmail(email, sub, msg, wg)
}

func MemberCreateEmail(Chan chan<- string, wg *sync.WaitGroup, data map[string]interface{}, email, action string) {

	var templates models.TblEmailTemplate
	models.GetTemplates(&templates, "Createmember")

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
	)

	msg = replacer.Replace(msg)
	GenerateEmail(email, sub, msg, wg)
}

func ForgetPasswordEmail(Chan chan<- string, wg *sync.WaitGroup, data map[string]interface{}, email, action string) {

	var templates models.TblEmailTemplate

	models.GetTemplates(&templates, action)

	sub := templates.TemplateSubject

	msg := templates.TemplateMessage

	replacer := strings.NewReplacer(
		"{FirstName}", data["uname"].(string),
		"{Passwordurl}", data["resetpassword"].(string),
		"{Passwordurlpath}", data["restpassurl"].(string),
		"{AdminLogo}", data["admin_logo"].(string),
		"{FbLogo}", data["fb_logo"].(string),
		"{LinkedinLogo}", data["linkedin_logo"].(string),
		"{TwitterLogo}", data["twitter_logo"].(string),
	)

	msg = replacer.Replace(msg)

	GenerateEmail(email, sub, msg, wg)
}

func OwndeskmemberEmail(Chan chan<- string, wg *sync.WaitGroup, data map[string]interface{}, email, action string) {

	var templates models.TblEmailTemplate

	models.GetTemplates(&templates, action)

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
		"<figure", "<div",
		"</figure", "</div",
		"&nbsp;", "",
	)

	sub = replacer.Replace(sub)
	msg = replacer.Replace(msg)
	GenerateOwndeskEmail(email, sub, msg, wg)
}

func hashingPassword(pass string) string {

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
