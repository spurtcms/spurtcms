package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"spurt-cms/models"
	storagecontroller "spurt-cms/storage-controller"
	"strconv"
	"strings"
	"sync"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	chn "github.com/spurtcms/channels"
	"github.com/spurtcms/courses"
	"github.com/spurtcms/menu"
	"github.com/spurtcms/team"
	csrf "github.com/utrack/gin-csrf"
	"golang.org/x/net/html"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/linkedin"
	"gorm.io/gorm"
)

var (
	ErrorOtpExpiry = errors.New("otp expired")

	ErrorInvalidOTP = errors.New("invalid OTP")
)

type UserInfo struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	GivenName string `json:"given_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type WebAuth struct {
	UserId        int
	ExpiryTime    int
	ExpiryFlg     bool
	SecretKey     string
	DB            *gorm.DB
	AuthFlg       bool
	PermissionFlg bool
	RoleId        int
	RoleName      string
	DataAccess    int
}

// Goolgle credentials
var (
	googleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	oauthStateString = os.Getenv("OAUTH_STATE")
)

// Linkedin credentials

var (
	linkedinConf = &oauth2.Config{
		ClientID:     os.Getenv("LINKEDIN_CLIENT_ID"),
		ClientSecret: os.Getenv("LINKEDIN_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("LINKEDIN_REDIRECT_URL"),
		Scopes:       []string{"r_liteprofile", "r_emailaddress"},
		Endpoint:     linkedin.Endpoint,
	}
)

//Star Repo

var (
	githubConfig = &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GITHUB_REDIRECT_URL"),
		Scopes:       []string{"user:email", "repo"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}
)

func NewWebsitePage(c *gin.Context) {

	TENANTID := os.Getenv("TENANTID")

	host := c.Request.Host
	subdomain := strings.Split(host, ".")[0]
	port := os.Getenv("PORT")

	if subdomain == "spurtcms" || host == "www.spurtcms.com" || subdomain == "lvh" || subdomain == "localhost:"+port || subdomain == "dev" || host == "103.189.89.34" {

		session, _ := Store.Get(c.Request, os.Getenv("SESSION_KEY"))
		tkn := session.Values["token"]

		if tkn != "" && tkn != nil && TENANTID != "1" {

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

		} else if TENANTID == "1" {

			WebsiteLoad(c, TENANTID, "")

			return

		} else {
			c.HTML(200, "index.html", gin.H{"linktitle": "Open Source Golang Based CMS Solution | spurtCMS", "description": "A CMS Solution that can be tailor-made for your content management needs. Customize it for any unique Delivery purpose.", "keywords": "Open source cms solutions, Golang open source cms, Golang CMS Solution, Golang Content Management System, Custom fields, Golang open source content management system, GOlang open source cms systems, content management software open source with Golang.", "OGImage": "/public/img/index.png"})
		}
	} else if subdomain != "" && subdomain != "lvh" && subdomain != "localhost:8000" && subdomain != "spurtcms" {

		WebsiteLoad(c, "", subdomain)

	}

}

func WebsiteLoad(c *gin.Context, Tenantid, subdomain string) {

	// tag := c.Query("tag")

	var website menu.TblWebsite

	if Tenantid == "1" {

		website, _ = MenuConfig.GetWebsiteById(1, "1")
	} else {

		website, _ = MenuConfig.GetWebsiteByName(subdomain)
	}
	if website.Id == 0 {

		c.HTML(404, "nodata.html", nil)
		c.Abort()
		return
	}

	fmt.Println(website, "websitdetailsdfd")

	User, _, _ := NewTeamWP.GetUserById(website.CreatedBy, []int{})

	templatedet, _ := MenuConfig.GetTemplateById(website.TemplateId, website.TenantId)

	menulist, _, _ := MenuConfig.MenuList(100, 0, menu.Filter{}, User.TenantId, website.Id)

	var newmenulist []menu.TblMenus
	var parentid int
	var menuitemslist []menu.TblMenus

	for _, val := range menulist {

		if val.Status == 1 {
			parentid = val.Id
		}

	}

	menuitemslist, err := MenuConfig.GetMenusByParentid(parentid, User.TenantId)
	if err != nil {
		fmt.Println(err)
	}

	for _, val := range menuitemslist {
		if val.ParentId != 0 {
			newmenulist = append(newmenulist, val)
		}
	}

	var (
		limt, offset, Count int
	)

	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	Location := strings.TrimSpace(c.Query("Location"))
	Title := strings.TrimSpace(c.Query("Title"))
	template_name := strings.ToLower(templatedet.TemplateName)
	template_name = strings.ReplaceAll(template_name, " ", "_")

	switch template_name {

	case "content_chronicle":
		limt = 12
	case "content_Verse":
		limt = 15
	case "support_sphere":
		limt = 0
	default:
		limt = Limit
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	var chnentry []chn.Tblchannelentries

	c.Set("menulist", newmenulist)

	if template_name == "jobs" {

		entries, _, count, _ := ChannelConfigWP.FlexibleChannelEntriesList(
			chn.EntriesInputs{
				Status:              "Published",
				ChannelName:         "Jobs",
				Limit:               limt,
				Offset:              offset,
				GetAdditionalFields: true,

				Title:    Title,
				Location: Location,
				TenantId: User.TenantId,
			},
		)
		Count = count

		for i := range entries {
			entries[i].CreatedDate = entries[i].CreatedOn.In(TZONE).Format("Jan 2 2006")
			if entries[i].ProfileImagePath != "" {
				entries[i].ProfileImagePath = "/image-resize?name=" + entries[i].ProfileImagePath
			}
			entries[i].Description = StripHTML(entries[i].Description)

			categoryid, _ := strconv.Atoi(entries[i].CategoriesId)
			categorydetails, _ := CategoryConfig.GetSubCategoryDetails(categoryid, User.TenantId)
			entries[i].CategoryGroup = categorydetails.CategoryName
			channelinfo, _, _ := ChannelConfigWP.GetChannelsById(entries[i].ChannelId, Tenantid)

			entries[i].ChannelName = channelinfo.ChannelName

		}
		chnentry = entries
	} else if template_name != "support_sphere" {

		for _, val := range newmenulist {
			if val.TypeId != 0 {
				var entries []chn.Tblchannelentries
				var count int

				entries, _, count, _ = ChannelConfigWP.ChannelEntriesList(
					chn.Entries{
						ChannelId: val.TypeId,
						Status:    "Published",
						Limit:     limt - len(chnentry),
						Offset:    offset,
					},
					User.TenantId,
				)
				Count = count
				for i := range entries {
					entries[i].CreatedDate = entries[i].CreatedOn.In(TZONE).Format("Jan 2 2006")
					if entries[i].ProfileImagePath != "" {
						entries[i].ProfileImagePath = "/image-resize?name=" + entries[i].ProfileImagePath
					}
					entries[i].Description = StripHTML(entries[i].Description)

					categoryid, _ := strconv.Atoi(entries[i].CategoriesId)
					categorydetails, _ := CategoryConfig.GetSubCategoryDetails(categoryid, User.TenantId)
					entries[i].CategoryGroup = categorydetails.CategoryName
				}

				chnentry = append(chnentry, entries...)

				if len(chnentry) >= limt {
					chnentry = chnentry[:limt]
					break
				}
			}
		}
	} else {
		for i := range newmenulist {
			val := &newmenulist[i]
			if val.TypeId != 0 {
				entries, _, _, _ := ChannelConfigWP.ChannelEntriesList(
					chn.Entries{
						ChannelId: val.TypeId,
						Status:    "Published",
					},
					User.TenantId,
				)

				val.Count = len(entries)

				if len(entries) > 3 {
					entries = entries[:3]
				}
				chnentry = append(chnentry, entries...)
			}
		}
	}

	batchSize := 4
	var batches [][]chn.Tblchannelentries
	for i := 0; i < len(chnentry); i += batchSize {
		end := i + batchSize
		if end > len(chnentry) {
			end = len(chnentry)
		}
		batch := chnentry[i:end]
		batches = append(batches, batch)
	}

	var login bool
	session, _ := Store.Get(c.Request, os.Getenv("SESSION_KEY"))
	tkn, ok := session.Values["token"].(string)
	if !ok || tkn == "" {
		login = true

	}

	seodetail, err := MenuConfig.SeoDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	settingsdetail, err := MenuConfig.SettingsDetail(User.TenantId, website.Id)
	if err != nil {
		fmt.Println(err)
	}

	var paginationendcount = len(chnentry) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(pageno, int(Count), limt)
	UserDetailsFunction(c)

	memberdet, _ := c.Get("userdetails")
	// Parse templates
	tmpl, err := template.ParseFiles("websites/themes/"+template_name+"/layouts/partials/header.html", "websites/themes/"+template_name+"/layouts/partials/footer.html", "websites/themes/"+template_name+"/layouts/partials/head.html", "websites/themes/"+template_name+"/layouts/_default/content_list.html")

	if err != nil {

		fmt.Println(err, "templateerr")
	}

	data := gin.H{"Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3}, "memberdet": memberdet, "menulist": newmenulist, "searchlist": chnentry, "login": login, "template_name": c.Query("template_name"), "seodetail": seodetail, "settingsdetail": settingsdetail, "Count": Count, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Limit": limt, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "Home": "Home"}

	CourseList, err := TemplateCourseList(limt, offset, User.TenantId)

	if err != nil {
		fmt.Println(err)
	}

	switch template_name {
	case "content_chronicle":
		data["Entries"] = batches
	case "support_sphere":
		data["Entries"] = chnentry
	case "courses":

		for _, val := range newmenulist {

			if val.SlugName == "all-courses" {

				data["course"] = CourseList

			}
		}

	case "ai_templates":

		// var ids []string

		// for _, val := range newmenulist {

		// 	ids = append(ids, val.ListingsIds)
		// }

		// allIds := strings.Join(ids, ",")

		// listingIds := strings.Split(allIds, ",")

		// listings, _ := ListingConfig.GetListingsByIds(listingIds, tag, User.TenantId)

		// data["listingslist"] = listings

		var homeurl string

		for index, val := range newmenulist {

			if index == 0 {
				homeurl = val.UrlPath
			}
		}
		data["homeurl"] = homeurl

	default:
		data["entries"] = chnentry
	}

	RenderTemplate(c, tmpl, "content_list.html", data)
}
func TemplateCourseList(limit int, offset int, tenantid string) ([]courses.TblCourses, error) {
	var CourseList []courses.TblCourses

	var filteredCoursesList []courses.TblCourses

	courseslist, _, err := CoursesConfig.CoursesList(limit, offset, courses.Filter{}, tenantid)

	for _, val := range courseslist {

		if val.Status == 1 {

			filteredCoursesList = append(filteredCoursesList, val)
		}
	}
	if err != nil {
		fmt.Println(err)
		return []courses.TblCourses{}, err
	}

	for _, value := range filteredCoursesList {

		if value.ImagePath != "" {

			value.ImagePath = "/image-resize?name=" + value.ImagePath

		}

		if value.ProfileImagePath != "" {

			value.ProfileImagePath = "/image-resize?name=" + value.ProfileImagePath

		}

		var first = value.FirstName
		var last = value.LastName
		var firstn string
		var lastn string

		if value.FirstName != "" {
			firstn = strings.ToUpper(first[:1])
		}
		if value.LastName != "" {
			lastn = strings.ToUpper(last[:1])
		}

		value.NameString = firstn + lastn

		if !value.ModifiedOn.IsZero() {
			value.DateString = value.ModifiedOn.In(TZONE).Format(Datelayout)
		} else {
			value.DateString = value.CreatedOn.In(TZONE).Format(Datelayout)
		}

		value.CreatedTime = value.CreatedOn.In(TZONE).Format(Datelayout)

		CourseList = append(CourseList, value)
	}

	return CourseList, nil
}

func GetTemplateName(c *gin.Context) string {
	templateName := c.Query("template_name")
	var baseName string
	if subdomain := strings.Split(c.Request.Host, ".")[0]; subdomain != "" &&
		subdomain != "lvh" && subdomain != "localhost:8000" && subdomain != "spurtcms" {
		User, err := models.GetuserBySubdomain(subdomain)
		if err == nil {
			templatedetails, err := MenuConfig.GetTemplateById(User.GoTemplateDefault, User.TenantId)
			if err == nil && templateName == "" {
				baseName = templatedetails.TemplateName
			}
		}
	}
	if templateName != "" {
		baseName = templateName
	}
	baseName = strings.ToLower(baseName)
	baseName = strings.ReplaceAll(baseName, " ", "_")

	return baseName
}
func RenderTemplate(c *gin.Context, tmpl *template.Template, templateName string, data interface{}) {

	err := tmpl.ExecuteTemplate(c.Writer, templateName, data)

	if err != nil {

		c.String(http.StatusInternalServerError, err.Error())
	}
}
func StripHTML(input string) string {
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		return input
	}
	var buf bytes.Buffer
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			buf.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return buf.String()
}
func WebsiteLogin(c *gin.Context) {

	c.HTML(200, "weblogin.html", gin.H{"csrf": csrf.GetToken(c), "Menu": NewMenuController(c), "linktitle": "Login"})
}

func GoogleLogin(c *gin.Context) {

	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// Doc/umentation
func Documentation(c *gin.Context) {

	routeName := c.FullPath()

	if routeName == "/documentation/" {

		c.HTML(200, "document.html", gin.H{"linktitle": "Documentation in SpurtCMS | CMS Documentation", "description": "Access SpurtCMS Documentation – Your guide to comprehensive CMS documentation. Explore Golang-powered insights for efficient and seamless content management system development.", "keywords": "CMS Developer Documentation, CMS Documentation in SpurtCMS, Documentation for SpurtCMS, Getting Started SpurtCMS Documentation, cms website documentation, Documentation Spurtcms.", "Document": false, "OGImage": "/public/img/documentation.png"})

	} else {
		c.HTML(200, "document.html", gin.H{"linktitle": "Prerequisite Documentation | SpurtCMS", "description": "Prepare with SpurtCMS Prerequisite Documentation – Your essential guide to setup. Ensure a smooth start to Golang-powered CMS development with comprehensive prerequisites.", "keywords": "CMS Prerequisite Documentation, Prerequisite Documentation in SpurtCMS, Prerequisite Documentation for SpurtCMS, Getting Started Prerequisite Documentation, cms website Prerequisite documentation, Documentation Prerequisite Spurtcms.", "Document": true, "OGImage": "/public/img/documentation.png"})
	}

}

func CmsAdmin(c *gin.Context) {

	c.HTML(200, "cms-admin.html", gin.H{"linktitle": "CMS Admin Setup | SpurtCMS", "description": "Effortlessly set up CMS Admin with SpurtCMS – Your guide to seamless Golang-powered administration. Build a robust CMS system with our user-friendly admin setup.", "keywords": "Setup cms admin, CMS Admin setup in SpurtCMS, CMS admin for SpurtCMS, Getting Started to setup CMS admin, cms Admin,Admin setup for CMS, Build a CMS admin, The CMS admin setup."})
}

func Graphqlapi(c *gin.Context) {

	c.HTML(200, "graphqlapi.html", gin.H{"linktitle": "GraphQL API Setup"})
}

func TemplateDoc(c *gin.Context) {

	c.HTML(200, "templatedoc.html", gin.H{"linktitle": "Template for SpurtCMS | SpurtCMS Open source code for CMS Website", "description": "With spurtCMS templates create content effortlessly for different content genres. Streamline your admin interface for tailored content delivery.", "keywords": "Golang open source template, Open source cms with Golang, Spurtcms template open source cms, Build a cms website with Golang, Golang open source platform, Best open source cms for Golang, Golang open source cms for website. Golang dynamic content management platform, open-source CMS built on Golang."})
}

func Blog(c *gin.Context) {

	c.HTML(200, "blogs.html", gin.H{"linktitle": "Blogs"})
}

func CLI(c *gin.Context) {

	c.HTML(200, "cli.html", gin.H{"linktitle": "CLI Document", "Document": true})
}

func Prebuild(c *gin.Context) {

	c.HTML(200, "prebuild.html", gin.H{"linktitle": "Prebuild Document", "Document": true})
}
func AboutUs(c *gin.Context) {

	c.HTML(200, "aboutus.html", gin.H{"linktitle": "About SpurtCMS | CMS Website Development Company", "description": "Discover SpurtCMS, your trusted CMS Website Development Company. We specialize in Golang, setting the standard for dynamic and innovative content management solutions.", "keywords": "Dedicated Golang developers for cms, Golang Developers for cms website, Spurtcms Golang team, Build a cms website with Golang Developers, Golang cms open source platform, Best open source cms for Golang team."})
}

//Opensource Package

func Packages(c *gin.Context) {

	c.HTML(200, "package.html", gin.H{"linktitle": "Packages SpurtCMS | Understanding Golang Packages in CMS", "description": "Dive into Packages SpurtCMS, delving into Golang Packages for CMS. Uncover the foundation of our comprehensive Golang package offerings for enhanced CMS development.", "keywords": "Goalng packages, spurtcms golang packages, Go packages, Spurtcms golang main packages, Go custom packages, Find golang cms packages."})
}

//OpenSource Cms-admin

func OpenSourceCmsAdmin(c *gin.Context) {

	c.HTML(200, "opencms-admin.html", gin.H{"linktitle": "Admin SpurtCMS | Golang Admin Content Management System", "keywords": "Golang admin cms, admin tool for content management system, Dynamic website admin, Golang dynamic admin control panel, build admin panel with golang, Best golang admin panel for cms, Build a website with Golang admin.", "description": "Elevate your CMS experience with spurtCMS Admin – a Golang-powered Admin Content Management System. Unlock powerful features for seamless content management."})
}

//Opensource Content-Template

func ContentTemplate(c *gin.Context) {

	c.HTML(200, "content-template.html", gin.H{"linktitle": "Content-Template"})
}

// GraphqlApi Document
func GraphqlApi(c *gin.Context) {

	c.HTML(200, "graphqldoc.html", gin.H{"linktitle": "GraphQL APIs"})
}

func CommercialLicense(c *gin.Context) {

	c.HTML(200, "commercial-license.html", gin.H{"linktitle": "SpurtCms Commercial License | CMS Source Code Commercial License", "description": "Secure your project with SpurtCMS Commercial License – Your key to CMS Source Code Commercial License. Leverage the power of Golang for a reliable and customizable CMS solution.", "keywords": "Golang headless cms, Golang cms open source, Golang cms framework, Golang cms solution, Find a cms source code, best content management system for websites."})
}

func TermsandPolicies(c *gin.Context) {

	c.HTML(200, "terms-policy.html", gin.H{"linktitle": "SpurtCMS Terms and Policy | Open source website for CMS", "description": "Review SpurtCMS Terms and Policy – Navigating the open-source website for CMS. Understand the guidelines and principles governing our dynamic content management system.", "keywords": "Best Golang cms website, Golang website cms, headless cms website, headless cms features, Golang open source cms, content management system terms, and policy, terms, and policy cms system."})
}

func Download(c *gin.Context) {

	c.HTML(200, "download.html", gin.H{"linktitle": "SpurtCMS Download | Golang cms open source code", "description": "Get started with SpurtCMS – Download Golang CMS open source code. Access the foundation for your content management system and kickstart your development journey.", "keywords": "CMS source code, Download Spurtcms source code, Golang with CMS, content management system source code, open source content management system, Develop cms website with open source, Golang cms open source code."})
}

//Our Story

func OurStory(c *gin.Context) {

	c.HTML(200, "ourstory.html", gin.H{"linktitle": "About SpurtCMS | Golang CMS Development Company", "description": "Explore SpurtCMS, your trusted Golang CMS Development Company. We specialize in crafting dynamic solutions, setting the standard for CMS development with Golang expertise.", "keywords": "Spurtcms story, Golang CMS development, Best Golang-based CMS Development, open-source CMS project with Golang, Golang content management solution, build CMS website with Golang, Golang open source CMS."})

}

//How it Helps for CMS

func HowItHelpls(c *gin.Context) {

	c.HTML(200, "help.html", gin.H{"linktitle": "How It Helps with Spurt Cms"})
}

//License Key

func LicenseKey(c *gin.Context) {

	c.HTML(200, "licensekey.html", gin.H{"linktitle": "SpurtCMS License Key"})
}

func Superaddminmail() string {

	userdetails, err := NewTeamWP.UserDetails(team.Team{Id: 1, TenantId: ""})
	if err != nil {
		log.Fatal("", err)
	}
	return userdetails.Email

}

func GoogleCallback(c *gin.Context) {

	state := c.Query("state")
	if state != oauthStateString {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	code := c.Query("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Get user info
	client := googleOauthConfig.Client(oauth2.NoContext, token)

	userInfo, err := getUserInfo(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	firstName, lastName := separateNames(userInfo.Name)

	var googleuser auth.SocialLogin

	googleuser.FirstName = firstName

	googleuser.LastName = lastName

	googleuser.GivenName = userInfo.GivenName

	googleuser.Email = userInfo.Email

	token1, userdetails, isNewUser, err := NewAuth.CheckWebAuth(&googleuser)

	if err != nil && err.Error() == "user disabled please contact admin" {

		c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "userisactive", 3600, "", "", false, false)
		c.Redirect(301, "/admin")
		return
	}
	linkedin := os.Getenv("LINKEDIN")
	facebook := os.Getenv("FACEBOOK")
	twitter := os.Getenv("TWITTER")
	youtube := os.Getenv("YOUTUBE")
	insta := os.Getenv("INSTAGRAM")
	var url_prefix = os.Getenv("BASE_URL")

	Superadminemail := Superaddminmail()

	timestamp := userdetails.CreatedOn.In(TZONE).Format(Datelayout)
	if isNewUser {

		var s3FolderName = userdetails.Username + "_" + userdetails.TenantId

		s3Path, err := storagecontroller.CreateFolderToS3(s3FolderName, "/")

		if err != nil {

			fmt.Println(err)

			c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
			c.Redirect(301, "/admin")

			return
		}

		err = NewAuth.UpdateS3FolderName(userdetails.TenantId, userdetails.Id, s3Path)

		if err != nil {

			fmt.Println(err)

			return
		}

		err = CreateTenantDefaultData(userdetails.Id, userdetails.TenantId)

		if err != nil {
			ErrorLog.Printf("CheckLogin error: %v", err)
			c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
			c.SetCookie("Alert-msg", "Internal Server Error", 3600, "", "", false, false)
			c.Redirect(301, "/admin")
			return
		}

		data := map[string]interface{}{

			"firstname":     userdetails.FirstName,
			"useremail":     userdetails.Email,
			"timestamp":     timestamp,
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
		go Superadminnotificatin(&wg, data, Superadminemail)

	}

	if userdetails.RoleId == 1 {

		c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "You are not registered with us!", 3600, "", "", false, false)
		c.Redirect(301, "/admin")
		return
	}
	if !userdetails.LastLogin.IsZero() {
		Lastactive := userdetails.LastLogin.In(TZONE).Format(Datelayout)
		Lastlogin[userdetails.Id] = Lastactive
	} else {
		Lastlogin[userdetails.Id] = "--"
	}
	// host := c.Request.Host
	// if strings.Contains(host, "spurtcms.com") {
	// 	Store.Options = &sessions.Options{
	// 		Path:     "/",
	// 		Domain:   ".spurtcms.com", // Use your real domain!
	// 		MaxAge:   86400 * 90,
	// 		HttpOnly: true,
	// 		Secure:   true, // Set to true in production
	// 		SameSite: http.SameSiteLaxMode,
	// 	}
	// }

	Session, _ := Store.Get(c.Request, os.Getenv("SESSION_KEY"))
	Session.Values["token"] = token1
	Session.Save(c.Request, c.Writer)

	if userdetails.IsActive == 1 {

		data := map[string]interface{}{

			"firstname":     userdetails.FirstName,
			"useremail":     userdetails.Email,
			"timestamp":     timestamp,
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
		wg.Add(2)
		go Loginedsuccessfully(&wg, data, userdetails.Email, userdetails.TenantId, "Logined successfully")

		if !isNewUser {
			go Registereduserloginalert(&wg, data, Superadminemail)
		}

	} else {
		WarnLog.Println(" Team email notification status not enabled")
	}

	// var baseURL string

	// switch {
	// case strings.Contains(host, "spurtcms.com"):
	// 	baseURL = fmt.Sprintf("https://%s.spurtcms.com", userdetails.Subdomain)

	// default:
	// 	c.Redirect(http.StatusFound, "/admin/dashboard")
	// 	return
	// }

	path := "/admin/dashboard"
	if isNewUser {
		path = "/admin/getting-started"
	}

	redirectURL := path
	c.Redirect(http.StatusFound, redirectURL)
}

func getUserInfo(client *http.Client) (UserInfo, error) {
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return UserInfo{}, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return UserInfo{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var userInfo UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return UserInfo{}, fmt.Errorf("failed to decode user info: %w", err)
	}

	return userInfo, nil
}

// regirstered login function//
func RegisteredUserLogin(c *gin.Context) {

	emailid := c.PostForm("emailid")

	user, err := NewAuth.CheckOtpLogin(emailid)

	if err != nil && err.Error() == "user disabled please contact admin" {

		c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "userisactive", 3600, "", "", false, false)
		c.Redirect(301, "/admin")
		return
	}

	if user.RoleId == 1 {

		c.SetCookie("log-toast", "You are not registered with us", 3600, "", "", false, false)
		c.Redirect(301, "/admin")
		return

	}

	Otp := strconv.Itoa(user.Otp)

	expiry := user.OtpExpiry.In(TZONE).Format(Datelayout)

	linkedin := os.Getenv("LINKEDIN")
	facebook := os.Getenv("FACEBOOK")
	twitter := os.Getenv("TWITTER")
	youtube := os.Getenv("YOUTUBE")
	insta := os.Getenv("INSTAGRAM")

	if user.IsActive == 1 {

		var url_prefix = os.Getenv("BASE_URL")

		data := map[string]interface{}{

			"firstname":     user.FirstName,
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
		go SendUserOtp(&wg, data, user.Email, user.TenantId, "OTPGenerate")

	} else {
		WarnLog.Println(" Team email notification status not enabled")
	}

	c.HTML(200, "weblogin.html", gin.H{"userinfo": user, "csrf": csrf.GetToken(c), "linktitle": "Login"})
}

func OtpVerification(c *gin.Context) {

	otp, _ := strconv.Atoi(c.PostForm("otpno"))

	email := c.PostForm("emailid")

	tenantid := c.PostForm("tenantid")

	userdetails, err := NewTeamWP.UserDetails(team.Team{EmailId: email, TenantId: ""})

	if err != nil {

		log.Println(err)
	}

	userinfo, token, isNewUser, err := NewAuth.OtpLoginVerification(otp, email, tenantid)

	if err != nil && err.Error() == "invalid OTP" {

		fmt.Println("invalidotp", err)
		c.SetCookie("log-toast", "Invalid Otp", 3600, "", "", false, false)
		c.HTML(200, "weblogin.html", gin.H{"userinfo": userdetails, "csrf": csrf.GetToken(c), "linktitle": "Login"})
		return
	}
	if err != nil && err.Error() == "otp expired" {
		c.SetCookie("log-toast", "Otp Expired", 3600, "", "", false, false)
		c.HTML(200, "weblogin.html", gin.H{"userinfo": userdetails, "csrf": csrf.GetToken(c), "linktitle": "Login"})
		return

	}

	linkedin := os.Getenv("LINKEDIN")
	facebook := os.Getenv("FACEBOOK")
	twitter := os.Getenv("TWITTER")
	youtube := os.Getenv("YOUTUBE")
	insta := os.Getenv("INSTAGRAM")
	var url_prefix = os.Getenv("BASE_URL")
	timestamp := userdetails.CreatedOn.In(TZONE).Format(Datelayout)
	Superadminemail := Superaddminmail()

	if isNewUser {

		var s3FolderName = userinfo.Username + "_" + userinfo.TenantId

		s3Path, err := storagecontroller.CreateFolderToS3(s3FolderName, "/")

		if err != nil {

			fmt.Println(err)
			c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
			c.Redirect(301, "/admin")

			return
		}

		err = NewAuth.UpdateS3FolderName(TenantId, userinfo.Id, s3Path)

		if err != nil {

			fmt.Println(err)

			return
		}

		err = CreateTenantDefaultData(userinfo.Id, userinfo.TenantId)

		if err != nil {
			ErrorLog.Printf("CheckLogin error: %v", err)
			c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
			c.SetCookie("Alert-msg", "Internal Server Error", 3600, "", "", false, false)
			c.Redirect(301, "/admin")
			return
		}

		data := map[string]interface{}{

			"firstname":     userdetails.FirstName,
			"useremail":     userdetails.Email,
			"timestamp":     timestamp,
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
		go Superadminnotificatin(&wg, data, Superadminemail)

	} else {
		WarnLog.Println("Tenant email notification status not enabled")
	}

	LogoutTime(userinfo.Id)
	// host := c.Request.Host
	// if strings.Contains(host, "spurtcms.com") {
	// 	Store.Options = &sessions.Options{
	// 		Path:     "/",
	// 		Domain:   ".spurtcms.com", // Use your real domain!
	// 		MaxAge:   86400 * 90,
	// 		HttpOnly: true,
	// 		Secure:   true, // Set to true in production
	// 		SameSite: http.SameSiteLaxMode,
	// 	}
	// } else if strings.Contains(host, "lvh.me") {

	// 	Store.Options = &sessions.Options{
	// 		Path:     "/",
	// 		Domain:   ".lvh.me", // Use your real domain!
	// 		MaxAge:   86400 * 90,
	// 		HttpOnly: true,
	// 		Secure:   false, // Set to true in production
	// 		SameSite: http.SameSiteLaxMode,
	// 	}
	// }

	Session, _ := Store.Get(c.Request, os.Getenv("SESSION_KEY"))
	Session.Values["token"] = token
	Session.Save(c.Request, c.Writer)

	if userdetails.IsActive == 1 {

		data := map[string]interface{}{
			"firstname":     userdetails.FirstName,
			"useremail":     userdetails.Email,
			"timestamp":     timestamp,
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
		wg.Add(2)
		go Loginedsuccessfully(&wg, data, userdetails.Email, userinfo.TenantId, "Logined successfully")

		if !isNewUser {
			go Registereduserloginalert(&wg, data, Superadminemail)
		}

	} else {
		WarnLog.Println(" Team email notification status not enabled")
	}

	// var baseURL string

	// switch {
	// case strings.Contains(host, "spurtcms.com"):
	// 	baseURL = fmt.Sprintf("https://%s.spurtcms.com", userinfo.Subdomain)

	// case strings.Contains(host, "lvh.me"):
	// 	baseURL = fmt.Sprintf("http://%s.lvh.me:8080", userinfo.Subdomain)

	// default:
	// 	c.Redirect(http.StatusFound, "/admin/dashboard")
	// 	return
	// }

	path := "/admin/dashboard"
	if isNewUser {
		path = "/admin/getting-started"
	}

	redirectURL := path
	c.Redirect(http.StatusFound, redirectURL)

}

// Template view page
func TemplatePage(c *gin.Context) {

	c.HTML(200, "template-view.html", gin.H{"linktitle": "Golang, Next.js CMS Website templates | Demo templates for CMS websites", "description": "Discover Golang and Next.js CMS templates for Ecommerce, LMS, blogging, online booking, job portals, portfolios, and memberships. Elevate your website now.", "keywords": "Ecommerce website templates, LMS website templates, ecommerce CMS for your online store template, Template for blogging websites, website templates for online booking systems, templates for job portal websites, Template for online portfolio websites, Templates for membership websites, spurtCMS website templates, content management system templates.", "Active": "Templates", "OGImage": "/public/img/Blog1.jpg"})
}

// Block view page

func BlockPage(c *gin.Context) {

	c.HTML(200, "blocks-view.html", gin.H{"linktitle": "Open source CMS with Blocks | SpurtCMS Blocks", "description": "Create stunning layouts in spurtCMS using customizable blocks. Add galleries, testimonials, or timelines for visually rich Golang CMS pages.", "keywords": "Best Golang cms website, blocks for cms, blocks in cms website, golang cms features, Golang open source cms blocks, content management system blocks.", "Active": "Blocks", "OGImage": "/public/img/Blocks1.jpg"})
}

// Entries view page

func EntriesPage(c *gin.Context) {

	c.HTML(200, "entries-view.html", gin.H{"linktitle": "Creating Entries for CMS Website | SpurtCMS Craft CMS Entries", "description": "Create, edit, and enrich content in spurtCMS with an HTML editor. Add media files for dynamic entries in this Golang-powered content management system.", "keywords": "Craft CMS website, open source Cms entries, Create CMS entries, CMS in Entries, Craft CMS Entries, Best open source CMS website.", "Active": "Entries", "OGImage": "/public/img/Work%20at%20Your%20Pace.png"})
}

// Member view page
func MembersPage(c *gin.Context) {

	c.HTML(200, "members-view.html", gin.H{"linktitle": "Open source cms Personalized Member Dashboards | SpurtCMS", "description": "Control access and roles with spurtCMS member management. A Golang-based CMS ensuring secure collaboration with defined user permissions.", "keywords": "Create a member for cms website, open source cms website Personalized members, Members list for the cms website, Open source content management system members, Create an open source cms members page, best open source cms platform.", "Active": "Members",
		"OGImage": "/public/img/Member.jpg"})
}

// Form view page
func FormPage(c *gin.Context) {

	c.HTML(200, "forms-view.html", gin.H{"linktitle": "Golang CMS website Forms | SpurtCMS Forms", "description": "Build dynamic forms in spurtCMS, a Golang-based CMS. Add free text, dropdowns, or file uploads to customize forms for every requirement.", "keywords": "Forms for cms, Event forms for cms website, Best open source cms website, Forms in cms, 	Content management system Form, Best CMS for form, Forms in CMS website, CMS Form.", "Active": "Forms", "OGImage": "/public/img/create-form-banner.svg"})

}

// Graphql Api Keys view page

func ApiKeyPage(c *gin.Context) {

	c.HTML(200, "apikeys-view.html", gin.H{"linktitle": "Graphql API", "Active": "Graphql API"})
}

// Graphql Playground view page

func GraphqlPage(c *gin.Context) {

	c.HTML(200, "graphql-view.html", gin.H{"linktitle": "Open source CMS With GraphQL | SpurtCMS GraphQL", "description": "Fetch precise data with GraphQL APIs in spurtCMS. Optimize performance for scalable, modern web and mobile applications powered by Golang.", "keywords": "CMS With Graphql, open source CMS Graphql, Golang content management systems Graphql, CMS Golang Graphql, open source CMS Graphql, CMS open source Graphql.", "Active": "Graphql Playground", "OGImage": "/public/img/graph-ql.png"})
}

//  Webhook view page

func WebhookPage(c *gin.Context) {

	c.HTML(200, "webhook-view.html", gin.H{"linktitle": "Create Webhooks for CMS Website | Open source CMS Webhooks", "description": "Automate workflows and send real-time updates with spurtCMS webhooks. Integrate with third-party apps seamlessly using this Golang-powered CMS.", "keywords": "CMS Webhooks, Best open source CMS webhooks, golang cms webhooks, Webhooks in CMS, Webhooks using CMS Website", "Active": "Webhook", "OGImage": "/public/img/webhookcreate.jpg"})
}

// category view page
func CategoryPage(c *gin.Context) {

	c.HTML(200, "category-view.html", gin.H{"linktitle": "Create a Category for open source CMS website | Open source CMS Category", "description": "Organize content effectively in spurtCMS, a Golang-based CMS. Use categories and subcategories to streamline structure and enhance searchability.", "keywords": "Creating a cms category, open source cms website category, category list for cms website, Open source content management system category, Create an open source cms category, open source cms platform.", "Active": "Categories", "OGImage": "/public/img/Category.jpg"})

}

// channel view page
func ChannelPage(c *gin.Context) {

	c.HTML(200, "channels-view.html", gin.H{"linktitle": "Create your channel for CMS website | Open source CMS Channel", "description": "Manage diverse content types in spurtCMS, a powerful Golang CMS. Create and organize blogs, eCommerce, or events with flexibility and ease.", "keywords": "Creating a cms channel, open source cms website channel, channel for cms website, Open source content management system channel, Create an open source cms channel, golang open source cms platform.", "Active": "Channels", "OGImage": "/public/img/channellistpage.png"})

}

// aiwriting page view
func AiWritingPage(c *gin.Context) {

	c.HTML(200, "aiwriting-view.html", gin.H{"linktitle": "Open source CMS with AI | SpurtCMS AI Content Writing", "Active": "AI Writing", "description": "Generate SEO-friendly content effortlessly with AI in spurtCMS. Specify tone, keywords, and word count to create high-quality entries in minutes.", "keywords": "Open source cms ai, Spurtcms in ai, AI cms open source platform, Ai cms open source, AI open source content management system, Publish content open source cms with AI, Best open source AI cms platform.", "OGImage": "/public/img/AI2.jpg"})
}

//Contact us Page view

func ContactusPage(c *gin.Context) {

	c.HTML(200, "contactus.html", gin.H{"linktitle": "SpurtCMS Contact Us | Golang Open Source CMS Solutions", "description": "Connect with SpurtCMS – Your source for Golang Open Source CMS Solutions. Reach out through our Contact Us page for tailored content management system expertise.", "keywords": "Golang open source solution, Golang cms solutions, built cms website with Golang, CMS Golang Tech Stack, Golang headless cms, golang cms open source, Golang cms framework, Golang Customize CMS Solution."})
}

//linked in login//

func LinkedInLogin(c *gin.Context) {

	url := linkedinConf.AuthCodeURL("state")
	c.Redirect(http.StatusFound, url)
}

func LinkedInCallBack(c *gin.Context) {

	code := c.Query("code")

	token, err := linkedinConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	client := linkedinConf.Client(oauth2.NoContext, token)

	resp, err := client.Get("https://api.linkedin.com/v2/me")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	var profile struct {
		ID                 string `json:"id"`
		LocalizedFirstName string `json:"localizedFirstName"`
		LocalizedLastName  string `json:"localizedLastName"`
	}
	err = json.NewDecoder(resp.Body).Decode(&profile)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	emailResp, err := client.Get("https://api.linkedin.com/v2/emailAddress?q=members&projection=(elements*(handle~))")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	defer emailResp.Body.Close()

	var emailResponse struct {
		Elements []struct {
			Handle struct {
				EmailAddress string `json:"emailAddress"`
			} `json:"handle~"`
		} `json:"elements"`
	}
	err = json.NewDecoder(emailResp.Body).Decode(&emailResponse)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	var email string

	if len(emailResponse.Elements) > 0 {
		email = emailResponse.Elements[0].Handle.EmailAddress

	} else {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var linkuserinfo auth.SocialLogin

	linkuserinfo.Email = email

	linkuserinfo.FirstName = profile.LocalizedFirstName

	linkuserinfo.LastName = profile.LocalizedLastName

	linkuserinfo.GivenName = profile.LocalizedFirstName

	token1, userdetails, isNewUser, err := NewAuth.CheckWebAuth(&linkuserinfo)

	if err != nil && err.Error() == "user disabled please contact admin" {

		c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "This account is inactive please contact the admin", 3600, "", "", false, false)
		c.Redirect(301, "/admin")
		return
	}
	linkedin := os.Getenv("LINKEDIN")
	facebook := os.Getenv("FACEBOOK")
	twitter := os.Getenv("TWITTER")
	youtube := os.Getenv("YOUTUBE")
	insta := os.Getenv("INSTAGRAM")
	var url_prefix = os.Getenv("BASE_URL")

	Superadminemail := Superaddminmail()

	timestamp := userdetails.CreatedOn.In(TZONE).Format(Datelayout)

	if isNewUser {

		var s3FolderName = userdetails.Username + "_" + userdetails.TenantId

		s3Path, err := storagecontroller.CreateFolderToS3(s3FolderName, "/")

		if err != nil {

			fmt.Println(err)

			c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
			c.Redirect(301, "/admin")
			return
		}

		err = NewAuth.UpdateS3FolderName(userdetails.TenantId, userdetails.Id, s3Path)

		if err != nil {

			fmt.Println(err)

			return
		}

		err = CreateTenantDefaultData(userdetails.Id, userdetails.TenantId)

		if err != nil {
			ErrorLog.Printf("CheckLogin error: %v", err)
			c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
			c.SetCookie("Alert-msg", "Internal Server Error", 3600, "", "", false, false)
			c.Redirect(301, "/admin")
			return
		}

		data := map[string]interface{}{

			"firstname":     userdetails.FirstName,
			"useremail":     userdetails.Email,
			"timestamp":     timestamp,
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
		go Superadminnotificatin(&wg, data, Superadminemail)
	}

	if userdetails.RoleId == 1 {

		c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "You are not registered with us!", 3600, "", "", false, false)
		c.Redirect(301, "/admin")
		return
	}
	if !userdetails.LastLogin.IsZero() {
		Lastactive := userdetails.LastLogin.In(TZONE).Format(Datelayout)
		Lastlogin[userdetails.Id] = Lastactive
	} else {
		Lastlogin[userdetails.Id] = "--"
	}
	// host := c.Request.Host
	// if strings.Contains(host, "spurtcms.com") {
	// 	Store.Options = &sessions.Options{
	// 		Path:     "/",
	// 		Domain:   ".spurtcms.com", // Use your real domain!
	// 		MaxAge:   86400 * 90,
	// 		HttpOnly: true,
	// 		Secure:   true, // Set to true in production
	// 		SameSite: http.SameSiteLaxMode,
	// 	}
	// }
	Session, _ := Store.Get(c.Request, os.Getenv("SESSION_KEY"))
	Session.Values["token"] = token1
	Session.Save(c.Request, c.Writer)

	if userdetails.IsActive == 1 {

		data := map[string]interface{}{

			"firstname":     userdetails.FirstName,
			"useremail":     userdetails.Email,
			"timestamp":     timestamp,
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
		wg.Add(2)
		go Loginedsuccessfully(&wg, data, userdetails.Email, userdetails.TenantId, "Logined successfully")

		if !isNewUser {
			go Registereduserloginalert(&wg, data, Superadminemail)
		}

	} else {
		WarnLog.Println(" Team email notification status not enabled")
	}

	// var baseURL string

	// switch {
	// case strings.Contains(host, "spurtcms.com"):
	// 	baseURL = fmt.Sprintf("https://%s.spurtcms.com", userdetails.Subdomain)

	// default:
	// 	c.Redirect(http.StatusFound, "/admin/dashboard")
	// 	return
	// }

	path := "/admin/dashboard"
	if isNewUser {
		path = "/admin/getting-started"
	}

	redirectURL := path
	c.Redirect(http.StatusFound, redirectURL)
}

func separateNames(fullName string) (string, string) {
	parts := strings.Fields(fullName)
	if len(parts) < 2 {

		return fullName, ""
	}
	firstName := parts[0]
	lastName := parts[1]

	return firstName, lastName
}

func GithubLogin(c *gin.Context) {
	authURL := githubConfig.AuthCodeURL("state")
	c.Redirect(http.StatusFound, authURL)
}

func GithubCallBack(c *gin.Context) {

	code := c.Query("code")
	if code == "" {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	token, err := githubConfig.Exchange(c, code)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	repoOwner := "spurtcms"
	repoName := "spurtcms"
	err = StarRepository(token.AccessToken, repoOwner, repoName)
	if err != nil {
		log.Println("error in staring the repo")
		return
	}

	c.Redirect(302, "https://github.com/spurtcms/spurtcms")
}

//star repo function//

func StarRepository(accessToken, owner, repo string) error {
	url := fmt.Sprintf("https://api.github.com/user/starred/%s/%s", owner, repo)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "token "+accessToken)
	req.Header.Set("Accept", "application/vnd.github+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to star repository")
	}

	return nil
}

func LogoutTime(userid int) {

	userdata, _, _ := NewTeam.GetUserById(userid, []int{})

	if !userdata.LastLogin.IsZero() {
		Lastactive := userdata.LastLogin.In(TZONE).Format(Datelayout)
		Lastlogin[userdata.Id] = Lastactive
	} else {
		Lastlogin[userdata.Id] = "--"
	}
}

func UserDetailsFunction(c *gin.Context) {

	session, _ := Store.Get(c.Request, "spurttemplate")

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

		member, _ := MemberConfig.GetMemberAndProfileData(userid, "", 0, "", tenantid)

		var firstn = strings.ToUpper(member.FirstName[:1])

		var lastn string
		if member.LastName != "" {
			lastn = strings.ToUpper(member.LastName[:1])
		}
		member.NameString = firstn + lastn

		member.NameString = firstn + lastn

		if err == nil && member.Id != 0 {
			c.Set("userdetails", member)
		}
	}
}
