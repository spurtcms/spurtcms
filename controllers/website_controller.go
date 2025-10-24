package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"spurt-cms/models"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	chn "github.com/spurtcms/channels"
	"github.com/spurtcms/courses"
	"github.com/spurtcms/menu"
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
