package controllers

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"math"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"path/filepath"
	"spurt-cms/lang"
	"spurt-cms/models"
	storagecontroller "spurt-cms/storage-controller"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	senderDisplayName := "Spurtcms"
	fromHeader := fmt.Sprintf("%s <%s>", senderDisplayName, Mail)
	emailBody := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: %s; charset=UTF-8\r\n\r\n%s",
		fromHeader, email, subject, contentType, result)

	// Compose the email.
	// emailBody := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: %s; charset=UTF-8\r\n\r\n%s", Mail, email, subject, contentType, result)

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

func AudioConvertBase64toByte(Data string, storagepath, routename string) (audioname string, path string, audiobyte []byte, err error) {

	extEndIndex := strings.Index(Data, ";base64,")
	base64data := Data[strings.IndexByte(Data, ',')+1:]
	var ext = Data[11:extEndIndex]
	rand_num := strconv.Itoa(int(time.Now().Unix()))
	AudioName := "Audio-" + rand_num + "." + ext
	storagePath := storagepath + "/Audio-" + rand_num + "." + ext

	decode, err := base64.StdEncoding.DecodeString(base64data)

	if err != nil {
		fmt.Println(err)
	}

	return AudioName, storagePath, decode, nil
}
func DocumentConvertBase64toByte(file *multipart.FileHeader) (documentName string, path string, fileBytes []byte, err error) {

	fileContent, err := file.Open()
	if err != nil {
		return "", "", nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer fileContent.Close()

	fileBytes = make([]byte, file.Size)
	if _, err := fileContent.Read(fileBytes); err != nil {
		return "", "", nil, fmt.Errorf("failed to read file: %w", err)
	}

	ext := strings.ToLower(strings.Split(file.Filename, ".")[1])
	randNum := strconv.Itoa(int(time.Now().Unix()))
	documentName = "Document-" + randNum + "." + ext
	path = "documents/" + documentName

	return documentName, path, fileBytes, nil
}
func ConvertBase64toByte(imageData string, storagepath string) (imgname string, path string, imagebyte []byte, err error) {

	metaEnd := strings.Index(imageData, ";base64,")
	if metaEnd == -1 {
		err = errors.New("invalid base64 image data: missing ';base64,'")
		return "", "", nil, err
	}

	metaStart := strings.Index(imageData, "data:image/")
	if metaStart == -1 {
		err = errors.New("invalid base64 image data: missing 'data:image/'")
		return "", "", nil, err
	}

	ext := imageData[metaStart+11 : metaEnd]
	if ext == "" {
		err = errors.New("could not extract image extension")
		return "", "", nil, err
	}

	base64data := imageData[metaEnd+8:]
	if base64data == "" {
		err = errors.New("no base64 data found")
		return "", "", nil, err
	}

	randNum := strconv.FormatInt(time.Now().UnixNano(), 10)
	imageName := "IMG-" + randNum + "." + ext
	storagePath := strings.TrimRight(storagepath, "/") + "/" + imageName

	decode, err := base64.StdEncoding.DecodeString(base64data)
	if err != nil {
		return "", "", nil, fmt.Errorf("base64 decode error: %w", err)
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

	if templates.IsActive == 1 {

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

}

func SendUserOtp(wg *sync.WaitGroup, data map[string]interface{}, email string, tenantid string, action string) {

	var templates models.TblEmailTemplate
	models.GetTemplates(&templates, "OTPGenerate", tenantid)
	if templates.IsActive == 1 {
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

}

func Superadminnotificatin(wg *sync.WaitGroup, data map[string]interface{}, email string) {
	var templates models.TblEmailTemplate
	models.GetTemplates(&templates, "Superadminnotification", "")

	if templates.IsActive == 1 {
		sub := templates.TemplateSubject
		msg := templates.TemplateMessage

		replacer := strings.NewReplacer(
			"{FirstName}", data["firstname"].(string),
			"{Useremail}", data["useremail"].(string),
			"{Timestamp}", data["timestamp"].(string),
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
}
func Registereduserloginalert(wg *sync.WaitGroup, data map[string]interface{}, email string) {
	var templates models.TblEmailTemplate
	models.GetTemplates(&templates, "registereduserlogin", "")

	if templates.IsActive == 1 {
		sub := templates.TemplateSubject
		msg := templates.TemplateMessage

		replacer := strings.NewReplacer(
			"{FirstName}", data["firstname"].(string),
			"{Useremail}", data["useremail"].(string),
			"{Timestamp}", data["timestamp"].(string),
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

		fmt.Println(email, data, msg, "meassage")
		err := GenerateEmail(email, sub, data, msg, wg)
		if err != nil {
			ErrorLog.Printf("Cann't login your account Email %s error:", err)
		}
	}
}
func Loginedsuccessfully(wg *sync.WaitGroup, data map[string]interface{}, email string, tenant_id string, action string) {
	var templates models.TblEmailTemplate
	models.GetTemplates(&templates, "Logined successfully", tenant_id)
	if templates.IsActive == 1 {
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
}

func ChangePasswordEmail(Chan chan<- string, wg *sync.WaitGroup, data map[string]interface{}, email, action string) {

	var templates models.TblEmailTemplate
	models.GetTemplates(&templates, "Changepassword", TenantId)

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
	GenerateEmail(email, sub, nil, msg, wg)
}

func MemberCreateEmail(Chan chan<- string, wg *sync.WaitGroup, data map[string]interface{}, email string, tenantid string, action string) {

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

func SupportEmail(Chan chan<- string, wg *sync.WaitGroup, data map[string]interface{}, email string, tenantid string) {

	var templates models.TblEmailTemplate
	models.GetTemplates(&templates, "SupportEmail", "")

	if templates.IsActive == 1 {

		sub := templates.TemplateSubject
		msg := templates.TemplateMessage

		replacer := strings.NewReplacer(
			"{Service}", data["service"].(string),
			"{FirstName}", data["fname"].(string),
			"{Useremail}", data["email"].(string),
			"{Number}", data["number"].(string),
			"{Timestamp}", data["timezone"].(string),
			"{Country}", data["Country"].(string),
			"{Describe}", data["describe"].(string),
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

		Superadminemail := Superaddminmail()
		GenerateEmail(Superadminemail, sub, data, msg, wg)

		// // test
		// GenerateEmail("nithyar396@gmail.com", sub, data, msg, wg)
	}

}

func SupportuserEmail(Chan chan<- string, wg *sync.WaitGroup, data map[string]interface{}, email string, tenantid string) {

	var templates models.TblEmailTemplate
	models.GetTemplates(&templates, "SupportuserverificationEmail", "")

	if templates.IsActive == 1 {

		sub := templates.TemplateSubject
		msg := templates.TemplateMessage

		replacer := strings.NewReplacer(
			"{Service}", data["service"].(string),
			"{FirstName}", data["fname"].(string),
			"{Useremail}", data["email"].(string),
			"{Number}", data["number"].(string),
			"{Timestamp}", data["timezone"].(string),
			"{Country}", data["Country"].(string),
			"{Describe}", data["describe"].(string),
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
func ForgetPasswordEmail(Chan chan<- string, wg *sync.WaitGroup, data map[string]interface{}, email, action string, tenantid string) {

	var templates models.TblEmailTemplate

	models.GetTemplates(&templates, action, tenantid)

	sub := templates.TemplateSubject

	msg := templates.TemplateMessage

	replacer := strings.NewReplacer(

		"{FirstName}", data["fname"].(string),
		"{Passwordurl}", data["resetpassword"].(string),
		"{Passwordurlpath}", data["restpassurl"].(string),
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

	GenerateEmail(email, sub, data, msg, wg)
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

		routeName = "/channels/"
	}

	if strings.Contains(routeName, "/user/") {

		routeName = "/user/"
	}

	if strings.Contains(routeName, "/usergroup/") {

		routeName = "/usergroup/"
	}

	if strings.Contains(routeName, "/categories/") {

		routeName = "/categories/"
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

	if strings.Contains(routeName, "/entries") {

		routeName = "/channels/"
	}

	// if strings.HasPrefix(routeName, "/aiassistance") {

	// 	routeName = "/aiassistance"
	// }
	if strings.HasPrefix(routeName, "/blocks/defaultlist") {

		routeName = "/blocks"
	}
	if strings.Contains(routeName, "/defaultlist") {

		routeName = "/blocks"
	}
	if strings.HasPrefix(routeName, "/formbuilder") {

		routeName = "/cta"
	}

	if strings.Contains(routeName, "order") {

		routeName = "/members"

	}

	if strings.Contains(routeName, "subscription") {

		routeName = "/members"

	}

	if strings.Contains(routeName, "membershiplevel") {

		routeName = "/members"

	}

	if strings.Contains(routeName, "membershipgroup") {

		routeName = "/members"

	}

	if strings.Contains(routeName, "membership") {

		routeName = "/members"

	}

	if strings.Contains(routeName, "courses") {

		routeName = "/courses"

	}

	if strings.Contains(routeName, "listing") {

		routeName = "/listing"

	}

	if strings.Contains(routeName, "/website") {

		routeName = "/website"

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

	fmt.Println("checkdssss")

	routename := c.FullPath()

	var imageName, imagePath, audioName, audioPath, data, documentName, documentPath string

	var docdata *multipart.FileHeader

	if strings.Contains(routename, "uploadb64audio") {

		data = c.PostForm("audiodata")

	} else if strings.Contains(routename, "uploadb64image") {

		data = c.PostForm("imagedata")

	} else if strings.Contains(routename, "uploadb64document") {
		fileheader, err := c.FormFile("documentdata")
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid file upload: " + err.Error()})
			return
		}
		docdata = fileheader
	}

	baseurl := os.Getenv("BASE_URL")

	tenantDetails, _ := NewTeam.GetTenantDetails(TenantId)

	if data != "" || docdata.Filename != "" {

		var (
			imageByte    []byte
			audioByte    []byte
			documentByte []byte
			err          error
			// document multipart.File
		)

		if strings.Contains(routename, "uploadb64audio") {

			audioName, audioPath, audioByte, err = AudioConvertBase64toByte(data, "entries", routename)
			if err != nil {
				ErrorLog.Printf("convert base 64 to byte error : %s", err)
			}

			audioPath = tenantDetails.S3FolderName + audioPath

			uerr := storagecontroller.UploadCropImageS3(audioName, audioPath, audioByte)

			if uerr != nil {
				c.JSON(500, gin.H{"Error": "Upload failed: Check S3 credentials and try again."})

				return
			}

			audioPath = baseurl + "audio-resize?name=" + audioPath

			c.JSON(200, gin.H{"Audiopath": audioPath, "Audioname": audioName})

		} else if strings.Contains(routename, "uploadb64image") {

			imageName, imagePath, imageByte, err = ConvertBase64toByte(data, "entries")
			if err != nil {
				ErrorLog.Printf("convert base 64 to byte error : %s", err)
			}

			imagePath = tenantDetails.S3FolderName + imagePath

			uerr := storagecontroller.UploadCropImageS3(imageName, imagePath, imageByte)

			if uerr != nil {
				c.JSON(500, gin.H{"Error": "Upload failed: Check S3 credentials and try again."})

				return
			}

			imagePath = baseurl + "image-resize?name=" + imagePath

			c.JSON(200, gin.H{"imagepath": imagePath, "imgname": imageName})
		} else if strings.Contains(routename, "uploadb64document") {

			fmt.Println("checkdocumentsddff:")

			extension := filepath.Ext(docdata.Filename)

			documentName, documentPath, documentByte, err = DocumentConvertBase64toByte(docdata)
			if err != nil {
				ErrorLog.Printf("convert base 64 to byte error : %s", err)
			}

			documentPath = tenantDetails.S3FolderName + documentPath

			uerr := storagecontroller.UploadDocumentS3(documentName, documentPath, documentByte)
			if uerr != nil {
				c.JSON(500, gin.H{"Error": "Upload failed: Check S3 credentials and try again."})
				return
			}

			switch extension {
			case ".pdf", ".doc", ".docx", ".txt", ".zip":

				documentPath = baseurl + "document-access?name=" + documentPath

				c.JSON(200, gin.H{"documentpath": documentPath, "documentname": documentName})

			case ".jpg", ".jpeg", ".png", ".gif":

				documentPath = baseurl + "image-resize?name=" + documentPath

				c.JSON(200, gin.H{"documentpath": documentPath, "documentname": documentName})

			}

		}
	} else {
		c.JSON(400, gin.H{"Error": "No data provided"})
	}

}

type Field struct {
	FieldName   string
	FieldTypeId int
	OrderIndex  int
	ImagePath   string
	TenantID    string
	CreatedBy   int
	CreatedOn   time.Time
}
type FieldOption struct {
	ID          int
	FieldID     int
	OptionName  string
	OptionValue string
	CreatedOn   time.Time
	TenantID    string
	CreatedBy   int
}
type ChannelCategory struct {
	ChannelID  int    `gorm:"column:channel_id"`
	CategoryID string `gorm:"column:category_id"`
	CreatedAt  int    `gorm:"column:created_at"`
	TenantID   string `gorm:"column:tenant_id"`
	CreatedOn  time.Time
}

func CreateTenantDefaultData(userId int, tenantId string) error {

	var err error

	file, err := os.Open("tenant-defaults.sql")

	if err != nil {

		return err
	}

	scanner := bufio.NewScanner(file)

	var (
		defChildCatId, defParentCatId, tagId, defChanId, subparentid, parentKnowledgeBaseId, parentBlogId, parentjobsId, templateModuleId, channelId, moduleId, blogChanId, jobsChanId, coursecatid, courseId int
		defBlockIds                                                                                                                                                                                           []int
		blockDynamicRetrieve                                                                                                                                                                                  bool
	)

	for scanner.Scan() {

		query := scanner.Text()

		if len(query) > 0 && !strings.HasPrefix(query, "--") {

			currentTime, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

			parts := strings.Fields(fmt.Sprint(currentTime))

			timeStamp := fmt.Sprintf("%s %s", parts[0], parts[1])

			uuid := (uuid.New()).String()

			arr := strings.Split(uuid, "-")

			Uuid := arr[len(arr)-1]

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

				// Step 2: Get Child Category ID (Fashion & Clothing)
				if strings.Contains(lowerQuery, "sp_id") {
					err = db.Table("tbl_categories").Where("is_deleted=0 AND tenant_id=? AND parent_id=?", tenantId, defParentCatId).Pluck("id", &subparentid).Error
					if err != nil && err != gorm.ErrRecordNotFound {
						return err
					}
					query = strings.ReplaceAll(query, "sp_id", strconv.Itoa(defParentCatId))
				}

				err = db.Debug().Table("tbl_categories").Where("category_slug='fashion_clothing' and tenant_id=?", tenantId).Pluck("id", &subparentid).Error
				if err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

				fmt.Println(subparentid, "categorysubid")
				if strings.Contains(lowerQuery, "sub_id") {
					query = strings.ReplaceAll(query, "sub_id", strconv.Itoa(subparentid))
				}
				err = db.Debug().Table("tbl_categories").Where("category_slug='electronics' and tenant_id=?", tenantId).Pluck("id", &subparentid).Error
				if err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

				fmt.Println(subparentid, "categorysubid")
				if strings.Contains(lowerQuery, "e_id") {
					query = strings.ReplaceAll(query, "e_id", strconv.Itoa(subparentid))
				}
				err = db.Debug().Table("tbl_categories").Where("category_slug='home_furniture' and tenant_id=?", tenantId).Pluck("id", &subparentid).Error
				if err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

				if strings.Contains(lowerQuery, "h_id") {
					query = strings.ReplaceAll(query, "h_id", strconv.Itoa(subparentid))
				}
				err = db.Debug().Table("tbl_categories").Where("category_slug='beauty_personal_care' and tenant_id=?", tenantId).Pluck("id", &subparentid).Error
				if err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

				if strings.Contains(lowerQuery, "b_id") {
					query = strings.ReplaceAll(query, "b_id", strconv.Itoa(subparentid))
				}
				err = db.Debug().Table("tbl_categories").Where("category_slug='sports_outdoors' and tenant_id=?", tenantId).Pluck("id", &subparentid).Error
				if err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

				if strings.Contains(lowerQuery, "so_id") {
					query = strings.ReplaceAll(query, "so_id", strconv.Itoa(subparentid))
				}
				err = db.Debug().Table("tbl_categories").Where("category_slug='politics' and tenant_id=?", tenantId).Pluck("id", &subparentid).Error
				if err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

				if strings.Contains(lowerQuery, "po_id") {
					query = strings.ReplaceAll(query, "po_id", strconv.Itoa(subparentid))
				}
				err = db.Debug().Table("tbl_categories").Where("category_slug='finance' and tenant_id=?", tenantId).Pluck("id", &subparentid).Error
				if err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

				if strings.Contains(lowerQuery, "bc_id") {
					query = strings.ReplaceAll(query, "bc_id", strconv.Itoa(subparentid))
				}
				// err = db.Debug().Table("tbl_categories").Where("category_slug='technology' and tenant_id=?", tenantId).Pluck("id", &subparentid).Error
				// if err != nil && err != gorm.ErrRecordNotFound {
				// 	return err
				// }

				// if strings.Contains(lowerQuery, "tc_id") {
				// 	query = strings.ReplaceAll(query, "tc_id", strconv.Itoa(subparentid))
				// }

				err = db.Debug().Table("tbl_categories").Where("category_slug='tech_insights' and tenant_id=?", tenantId).Pluck("id", &subparentid).Error
				if err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

				if strings.Contains(lowerQuery, "tc_id") {
					query = strings.ReplaceAll(query, "tc_id", strconv.Itoa(subparentid))
				}
				err = db.Debug().Table("tbl_categories").Where("category_slug='health' and tenant_id=?", tenantId).Pluck("id", &subparentid).Error
				if err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

				if strings.Contains(lowerQuery, "hc_id") {
					query = strings.ReplaceAll(query, "hc_id", strconv.Itoa(subparentid))
				}
				err = db.Debug().Table("tbl_categories").Where("category_slug='sports' and tenant_id=?", tenantId).Pluck("id", &subparentid).Error
				if err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

				if strings.Contains(lowerQuery, "sc_id") {
					query = strings.ReplaceAll(query, "sc_id", strconv.Itoa(subparentid))
				}
				err = db.Debug().Table("tbl_categories").Where("category_slug='worldnews' and tenant_id=?", tenantId).Pluck("id", &subparentid).Error
				if err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

				if strings.Contains(lowerQuery, "w_id") {
					query = strings.ReplaceAll(query, "w_id", strconv.Itoa(subparentid))
				}
			case strings.Contains(lowerQuery, "tbl_module_permissions") && strings.Contains(lowerQuery, "insert into"):

				err = db.Table("tbl_channels").Where("is_deleted = 0 and is_active = 1").Pluck("id", &channelId).Error

				if err != nil && err != gorm.ErrRecordNotFound {

					return err
				}

				if strings.Contains(lowerQuery, "rname") {

					query = strings.ReplaceAll(lowerQuery, "rname", fmt.Sprintf("/entries/entrylist/%d ", channelId))
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

				if err = db.Debug().Table("tbl_categories").Where("is_deleted=0 and tenant_id=? and parent_id=0 and category_slug=?", tenantId, "default_category").Pluck("id", &defParentCatId).Error; err != nil {

					return err
				}

				if err = db.Debug().Table("tbl_categories").Where("is_deleted=0 and tenant_id=?  and category_slug=?", tenantId, "default1").Pluck("id", &defChildCatId).Error; err != nil {

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
			case (strings.Contains(lowerQuery, "tbl_blocks") && strings.Contains(lowerQuery, "insert into")):

				if err = db.Debug().Table("tbl_channels").Where("is_deleted=0 and tenant_id = ? and channel_name=?", tenantId, "Default Channel").Pluck("id", &blogChanId).Error; err != nil {

					return err
				}
				if err = db.Debug().Table("tbl_channels").Where("is_deleted=0 and tenant_id = ? and channel_name=?", tenantId, "Jobs").Pluck("id", &jobsChanId).Error; err != nil {

					return err
				}

				if strings.Contains(query, "dch_id") {
					query = strings.ReplaceAll(query, "dch_id", strconv.Itoa(blogChanId))

				}
				if strings.Contains(query, "jch_id") {
					query = strings.ReplaceAll(query, "jch_id", strconv.Itoa(jobsChanId))

				}
			case (strings.Contains(lowerQuery, "tbl_block_tags") || strings.Contains(lowerQuery, "tbl_block_collections")) && strings.Contains(lowerQuery, "insert into") && !blockDynamicRetrieve:

				if err = db.Table("tbl_block_mstr_tags").Select("id").Where("tenant_id=?", tenantId).Pluck("id", &tagId).Error; err != nil {

					return err
				}

				if err = db.Table("tbl_blocks").Select("id").Where("is_deleted=0 and tenant_id=?", tenantId).Scan(&defBlockIds).Error; err != nil {

					return err
				}

				blockDynamicRetrieve = true

			}

			if strings.Contains(query, "tbl_forms") {

				err = db.Table("tbl_channels").Where("channel_name= ? and tenant_id=?", "Default Channel", tenantId).Pluck("id", &channelId).Error
				if err != nil {
					return err
				}

				query = strings.ReplaceAll(query, "uu_id", Uuid)

				query = strings.ReplaceAll(query, "ch_id", strconv.Itoa(channelId))
			}

			if strings.Contains(lowerQuery, "tbl_courses") {

				err = db.Table("tbl_categories").Where("category_slug='engineering' and tenant_id=?", tenantId).Pluck("id", &coursecatid).Error
				if err != nil {
					return err
				}
				query = strings.ReplaceAll(query, "co_id", strconv.Itoa(coursecatid))
			}

			if strings.Contains(lowerQuery, "tbl_course_settings") {

				err = db.Table("tbl_courses").Where("title= ? and tenant_id=?", "Default Course", tenantId).Pluck("id", &courseId).Error
				if err != nil {
					return err
				}
				query = strings.ReplaceAll(query, "cos_id", strconv.Itoa(courseId))
			}

			replacer := strings.NewReplacer("u_id", strconv.Itoa(userId), "tid", tenantId, "current-time", timeStamp, "duid", Uuid)

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

			var channelIds []int
			err := db.Table("tbl_channels").Where("is_deleted = 0 AND tenant_id = ?", tenantId).Pluck("id", &channelIds).Error
			if err != nil {
				return err
			}

			channelFields := map[string][]Field{
				"blog": {
					{"Excerpt", 2, 1, "/public/img/text.svg", tenantId, userId, currentTime},
					{"Author Bio", 2, 2, "/public/img/text.svg", tenantId, userId, currentTime},
				},
				"ecommerce": {
					{"Price", 2, 1, "/public/img/text.svg", tenantId, userId, currentTime},
					{"Special Price", 2, 2, "/public/img/text.svg", tenantId, userId, currentTime},
					{"SKU", 2, 3, "/public/img/text.svg", tenantId, userId, currentTime},
					{"Quantity", 2, 4, "/public/img/text.svg", tenantId, userId, currentTime},
					{"Stock", 5, 5, "/public/img/select.svg", tenantId, userId, currentTime},
				},
				"appointment": {
					{"Date", 6, 1, "/public/img/date.svg", tenantId, userId, currentTime},
					{"Time", 4, 2, "/public/img/date-time.svg", tenantId, userId, currentTime},
					{"contact number", 2, 3, "/public/img/text.svg", tenantId, userId, currentTime},
					{"Email", 2, 4, "/public/img/text.svg", tenantId, userId, currentTime},
					{"DOB/Age", 4, 5, "/public/img/date-time.svg", tenantId, userId, currentTime},
				},

				"event": {
					{"Event Title", 2, 1, "/public/img/text.svg", tenantId, userId, currentTime},
					{"Event Date", 6, 2, "/public/img/date.svg", tenantId, userId, currentTime},
					{"Event Time", 4, 3, "/public/img/date-time.svg", tenantId, userId, currentTime},
					{"Event type", 5, 4, "/public/img/select.svg", tenantId, userId, currentTime},
					{"Event Description", 8, 5, "/public/img/text-area.svg", tenantId, userId, currentTime},
				},
				"jobs": {
					{"Key Responsibilities", 2, 1, "/public/img/text.svg", tenantId, userId, currentTime},
					{"Job Type", 5, 2, "/public/img/select.svg", tenantId, userId, currentTime},
					{"Salary Range", 2, 3, "/public/img/text.svg", tenantId, userId, currentTime},
					{"Experiance", 2, 5, "/public/img/text.svg", tenantId, userId, currentTime},
				},
				"news": {
					{"Frequency", 5, 1, "/public/img/select.svg", tenantId, userId, currentTime},
					{"Anchor/Host Details", 8, 2, "/public/img/text-area.svg", tenantId, userId, currentTime},
					{"Coverage Scope", 5, 3, "/public/img/select.svg", tenantId, userId, currentTime},
					{"Special Features", 5, 4, "/public/img/select.svg", tenantId, userId, currentTime},
					{"Multimedia Support", 5, 5, "/public/img/select.svg", tenantId, userId, currentTime},
				},
				"knowledge-base": {
					{"Version", 2, 1, "/public/img/text.svg", tenantId, userId, currentTime},
					{"Summary", 2, 1, "/public/img/text.svg", tenantId, userId, currentTime},
				},
			}
			fieldOptionsMap := map[string][]string{
				"Stock":            {"In Stock", "Out of Stock"},
				"Event type":       {"Webinar", "Conference", "Workshop"},
				"Job Type":         {"Full-Time", "Part-Time", "Contract", "Internship"},
				"Frequency":        {"24/7", "Daily", "Weekly"},
				"Coverage Scope":   {"Local", "International"},
				"Special Features": {"Debates", "Expert Panels", "Investigative Reports"},
			}
			var fieldIds []int

			for _, channelId := range channelIds {
				var channelSlug string

				err = db.Table("tbl_channels").Where("id = ? and tenant_id=?", channelId, tenantId).Pluck("slug_name", &channelSlug).Error
				if err != nil {
					return err
				}

				var categoryId int
				err = db.Table("tbl_categories").Where("tenant_id = ? AND category_slug = ?", tenantId, channelSlug).Pluck("id", &categoryId).Error
				if err != nil {
					return err
				}

				if categoryId != 0 {
					if err := insertCategoryHierarchy(db, tenantId, categoryId, channelId, userId, currentTime); err != nil {
						return err
					}
				}

				fields := channelFields[channelSlug]
				if len(fields) == 0 {

					continue
				}

				fieldIds = nil

				for _, field := range fields {

					var existingFieldId int
					err = db.Table("tbl_fields").Where("tenant_id = ? AND field_name = ?", tenantId, field.FieldName).
						Pluck("id", &existingFieldId).Error

					if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
						return err
					}

					if existingFieldId == 0 {

						if err = db.Table("tbl_fields").Create(&field).Error; err != nil {
							return err
						}

						err = db.Table("tbl_fields").Where("tenant_id = ?", tenantId).Order("id desc").Limit(1).Pluck("id", &existingFieldId).Error
						if err != nil {
							return err
						}

						if options, exists := fieldOptionsMap[field.FieldName]; exists {
							if err := InsertFieldOptions(db, existingFieldId, options, userId, tenantId, currentTime); err != nil {
								return (err)
							}
						}
					}

					fieldIds = append(fieldIds, existingFieldId)
				}

				for _, fieldId := range fieldIds {
					groupField := map[string]interface{}{
						"channel_id": channelId,
						"field_id":   fieldId,
						"tenant_id":  tenantId,
					}
					var existingGroupField int64
					err = db.Table("tbl_group_fields").Where(groupField).Count(&existingGroupField).Error
					if err != nil {
						return err
					}

					if existingGroupField == 0 {
						if err = db.Debug().Table("tbl_group_fields").Create(&groupField).Error; err != nil {
							return err
						}
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

//sitemap file//

func Sitemap(c *gin.Context) {
	data, err := ioutil.ReadFile("./sitemap.xml")
	if err != nil {
		c.String(http.StatusInternalServerError, "Unable to read sitemap.xml")
		return
	}

	// Serve the content as XML
	c.Data(http.StatusOK, "application/xml", data)
}
func InsertFieldOptions(db *gorm.DB, existingFieldId int, options []string, userid int, tenantid string, currentime time.Time) error {
	if len(options) == 0 {
		return errors.New("no options provided")
	}

	for _, option := range options {
		optionRecord := &FieldOption{
			FieldID:     existingFieldId,
			OptionValue: option,
			OptionName:  option,
			CreatedOn:   currentime,
			CreatedBy:   userid,
			TenantID:    tenantid,
		}

		if err := db.Table("tbl_field_options").Create(optionRecord).Error; err != nil {
			return err
		}
	}
	return nil
}
func insertCategoryHierarchy(db *gorm.DB, tenantId string, parentId int, channelId int, userId int, currentTime time.Time) error {

	var childCategoryIds []int
	err := db.Table("tbl_categories").Where("tenant_id = ? AND parent_id = ?", tenantId, parentId).Pluck("id", &childCategoryIds).Error
	if err != nil {
		return err
	}

	for _, childCategoryId := range childCategoryIds {
		if childCategoryId == 0 {
			continue
		}

		categoryIDString := fmt.Sprintf("%d,%d", parentId, childCategoryId)

		var existingRecord ChannelCategory
		err = db.Table("tbl_channel_categories").Where("channel_id = ? AND category_id = ?", channelId, categoryIDString).First(&existingRecord).Error

		if err == nil {
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		channelCategoryRecord := ChannelCategory{
			ChannelID:  channelId,
			CategoryID: categoryIDString,
			CreatedAt:  userId,
			TenantID:   tenantId,
			CreatedOn:  currentTime,
		}

		err = db.Table("tbl_channel_categories").Create(&channelCategoryRecord).Error
		if err != nil {
			return err
		}

		if err := insertSubCategories(db, tenantId, childCategoryId, channelId, userId, currentTime, parentId); err != nil {
			return err
		}
	}

	return nil
}

func insertSubCategories(db *gorm.DB, tenantId string, parentId int, channelId int, userId int, currentTime time.Time, grandParentId int) error {
	var subChildCategoryIds []int

	err := db.Table("tbl_categories").Where("tenant_id = ? AND parent_id = ?", tenantId, parentId).Pluck("id", &subChildCategoryIds).Error
	if err != nil {
		return err
	}

	for _, subChildCategoryId := range subChildCategoryIds {
		if subChildCategoryId == 0 {
			continue
		}

		categoryIDString := fmt.Sprintf("%d,%d,%d", grandParentId, parentId, subChildCategoryId)

		var existingRecord ChannelCategory
		err = db.Table("tbl_channel_categories").Where("channel_id = ? AND category_id = ?", channelId, categoryIDString).First(&existingRecord).Error

		if err == nil {
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		channelCategoryRecord := ChannelCategory{
			ChannelID:  channelId,
			CategoryID: categoryIDString,
			CreatedAt:  userId,
			TenantID:   tenantId,
			CreatedOn:  currentTime,
		}

		err = db.Table("tbl_channel_categories").Create(&channelCategoryRecord).Error
		if err != nil {
			return err
		}
	}

	return nil
}
