package controllers

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"spurt-cms/models"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func Getsupport(c *gin.Context) {

	CountryList, _ := models.GetCountryLIst()

	menu := NewMenuController(c)

	translate, _ := TranslateHandler(c)

	timezones, err := models.GetTimeZones(TenantId)
	if err != nil {
		ErrorLog.Printf("Getting data from general settings Error:%s", err)
	}

	// ModuleName, TabName, _ := ModuleRouteName(c)

	c.HTML(200, "support.html", gin.H{"csrf": csrf.GetToken(c), "Menu": menu, "translate": translate, "TimeZone": timezones, "CountryList": CountryList})
}

func Supportsubmited(c *gin.Context) {

	status := c.Query("status")
	var pament string
	if status != "cancelled" {
		pament = "Success"
	} else {
		pament = "Cancel"
	}

	menu := NewMenuController(c)

	translate, _ := TranslateHandler(c)

	// ModuleName, TabName, _ := ModuleRouteName(c)

	c.HTML(200, "submited.html", gin.H{"csrf": csrf.GetToken(c), "Payment": pament, "Menu": menu, "translate": translate})
}

// func Sendsupportemail(c *gin.Context) {

// 	companyname := c.PostForm("companyname")
// 	contectemail := c.PostForm("contectemail")
// 	contectnumber := c.PostForm("contectnumber")
// 	timezone := c.PostForm("timezone")
// 	Country := c.PostForm("Country")
// 	describe := c.PostForm("describe")
// 	service := c.PostForm("service")

// 	var url_prefix = os.Getenv("BASE_URL")
// 	linkedin := os.Getenv("LINKEDIN")
// 	facebook := os.Getenv("FACEBOOK")
// 	twitter := os.Getenv("TWITTER")
// 	youtube := os.Getenv("YOUTUBE")
// 	insta := os.Getenv("INSTAGRAM")

// 	data := map[string]interface{}{
// 		"service":       service,
// 		"fname":         companyname,
// 		"email":         contectemail,
// 		"number":        contectnumber,
// 		"timezone":      timezone,
// 		"Country":       Country,
// 		"describe":      describe,
// 		"admin_logo":    url_prefix + "public/img/SpurtCMSlogo.png",
// 		"fb_logo":       url_prefix + "public/img/email-icons/facebook.png",
// 		"linkedin_logo": url_prefix + "public/img/email-icons/linkedin.png",
// 		"twitter_logo":  url_prefix + "public/img/email-icons/x.png",
// 		"youtube_logo":  url_prefix + "public/img/email-icons/youtube.png",
// 		"insta_log":     url_prefix + "public/img/email-icons/instagram.png",
// 		"facebook":      facebook,
// 		"instagram":     insta,
// 		"youtube":       youtube,
// 		"linkedin":      linkedin,
// 		"twitter":       twitter,
// 	}

// 	Chan := make(chan string, 1)

// 	var wg1 sync.WaitGroup

// 	wg1.Add(1)

// 	go SupportEmail(Chan, &wg1, data, contectemail, TenantId)

// 	var wg sync.WaitGroup

// 	wg.Add(1)

// 	go SupportuserEmail(Chan, &wg, data, contectemail, TenantId)

// 	close(Chan)

// 	// c.Redirect(301, "/getsupport/")
// 	c.Redirect(301, "/getsupport/supportsubmited")
// }

func MakeStripePayment(c *gin.Context) {

	companyname := c.PostForm("companyname")
	contectemail := c.PostForm("contectemail")
	contectnumber := c.PostForm("contectnumber")
	Country := c.PostForm("Country")
	// Interval := c.PostForm("interval")

	var asset_code string

	var PaymentURL = os.Getenv("GETSUPPORTAPI")

	var Return_URL = os.Getenv("return_url")

	var product_type_key = os.Getenv("product_type_key")
	var product_id = os.Getenv("product_id")

	// var asset_code_monthly = os.Getenv("asset_code_monthly")

	var asset_code_Yearly = os.Getenv("asset_code_Yearly")

	asset_code = asset_code_Yearly


	// switch Interval {
	// case "year":
	// 	asset_code = asset_code_Yearly
	// case "month":
	// 	asset_code = asset_code_monthly
	// default:
	// 	fmt.Println("Invalid interval")
	// }

	var business_id = os.Getenv("business_id")

	// Create the form data
	formData := url.Values{
		"return_url":       {Return_URL},
		"cancel_url":       {"http://dev.spurtcms.com/admin/get-spurtcmspro"},
		"product_type_key": {product_type_key},
		"email":            {contectemail},
		"product_id":       {product_id},
		"asset_code":       {asset_code},
		"country":          {Country},
		// "interval":         {Interval},
		"firstname":        {companyname},
		"phone":            {contectnumber},
		"business_id":      {business_id},
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", PaymentURL, bytes.NewBufferString(formData.Encode()))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	c.Data(http.StatusOK, "text/html", body)

}
