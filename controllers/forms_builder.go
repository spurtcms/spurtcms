package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"spurt-cms/models"
	storagecontroller "spurt-cms/storage-controller"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	"github.com/spurtcms/channels"
	forms "github.com/spurtcms/forms-builders"
	csrf "github.com/utrack/gin-csrf"
)

func FormbuilderList(c *gin.Context) {

	var (
		limt       int
		offset     int
		filter     forms.Filter
		NameString string
	)

	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")
	filter.FromDate = c.Query("fromdate")
	filter.ToDate = c.Query("todate")
	filter.ChannelSlug = c.Query("channel")

	// fmt.Println(filter.ChannelSlug, "channeliduuu")

	viewurl := os.Getenv("VIEW_BASE_URL")

	var filterflag string
	if filter.Keyword != "" || filter.FromDate != "" || filter.ToDate != "" {
		filterflag = "true"
	} else if filter.ChannelSlug != "" {

		filterflag = "ftrue"
	} else {
		filterflag = "false"
	}

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	routename := c.FullPath()

	var htmlname string

	var Status int

	var defaultlist int

	if strings.Contains(routename, "defaultctalist") {

		htmlname = "defaultcta.html"

		Status = 1

		defaultlist = 1
	} else {

		htmlname = "cta.html"

		Status = 1

		defaultlist = 0

	}

	var FormsBuidlersList []forms.TblForms

	permisison, perr := NewAuth.IsGranted("Forms Builder", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("formbuilders authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("formbuilders authorization error")
		c.Redirect(301, "/403-page")
		return
	}
	// channelist, clerr := ChannelConfig.GetPermissionChannel(TenantId)
	channelist, _, err := ChannelConfig.ListChannel(channels.Channels{Limit: 200, Offset: 0, IsActive: true, TenantId: TenantId})
	if err != nil {
		ErrorLog.Printf("%v: %v", ErrGettingTemplateModuleList, err)

	}

	channelSlug := c.Query("channel")

	chnanem, err := models.GetChannelDetailsWithTemplateCount(channelSlug, TenantId)
	if err != nil {
		ErrorLog.Printf("%v: %v", ErrGettingTemplateModuleList, err)

	}

	var finalChannelList []channels.Tblchannel

	for _, chnList := range channelist {

		_, ctacount, _, _ := FormConfig.FormBuildersList(0, 0, forms.Filter{}, TenantId, Status, 0, chnList.SlugName, defaultlist)
		chnList.EntriesCount = int(ctacount)
		finalChannelList = append(finalChannelList, chnList)
	}

	if permisison {

		FormConfig.DataAccess = c.GetInt("dataaccess")
		FormConfig.UserId = c.GetInt("userid")

		Formlist, TotalFormsCount, responseCount, err := FormConfig.FormBuildersList(limt, offset, forms.Filter(filter), TenantId, Status, 0, "", defaultlist)

		if err != nil {
			fmt.Println(err)
		}

		for _, val := range Formlist {

			val.DateString = val.ModifiedOn.In(TZONE).Format(Datelayout)

			val.CreatedDate = val.CreatedOn.In(TZONE).Format(Datelayout)

			val.Channelnamearr = strings.Split(val.ChannelName, ",")

			var first = val.FirstName
			var last = val.LastName
			var firstn = strings.ToUpper(first[:1])
			var lastn string
			if val.LastName != "" {
				lastn = strings.ToUpper(last[:1])
			}
			var Name = firstn + lastn
			val.NameString = Name

			imgcontain := "/image-resize?name="

			if val.ProfileImagePath != "" {
				userimg := val.ProfileImagePath
				imgflag := strings.Contains(userimg, imgcontain)
				if !imgflag {
					val.ProfileImagePath = "/image-resize?name=" + val.ProfileImagePath
				}
			}

			FormsBuidlersList = append(FormsBuidlersList, val)

		}

		//pagination calc
		paginationendcount := len(Formlist) + offset
		paginationstartcount := offset + 1
		Previous, Next, PageCount, Page := Pagination(pageno, int(TotalFormsCount), limt)

		menu := NewMenuController(c)

		// ModuleName, TabName, _ := ModuleRouteName(c)

		ModuleName := "Forms Builder"

		translate, _ := TranslateHandler(c)

		ctabannerval, _ := c.Cookie("ctabanner")

		if ctabannerval == "" {

			ctabannerval = "true"
		}
		ctadesc, _ := c.Cookie("ctadesc")

		if ctadesc == "" {

			ctadesc = "true"
		}

		c.HTML(200, htmlname, gin.H{"Pagination": PaginationData{
			NextPage:     pageno + 1,
			PreviousPage: pageno - 1,
			TotalPages:   PageCount,
			TwoAfter:     pageno + 2,
			TwoBelow:     pageno - 2,
			ThreeAfter:   pageno + 3,
		}, "Menu": menu, "ctabanner": ctabannerval, "chnid": channelSlug, "ctadesc": ctadesc, "channelname": chnanem.ChannelName, "entrycount": TotalFormsCount, "Viewbaseurl": viewurl, "Formlist": FormsBuidlersList, "ResponseCount": responseCount, "SearchTrues": filterflag, "csrf": csrf.GetToken(c), "Count": TotalFormsCount, "Limit": limt, "Paginationendcount": paginationendcount, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Filter": filter, "Paginationstartcount": paginationstartcount, "translate": translate, "title": ModuleName, "linktitle": "Forms", "Tabmenu": "CTA", "NameString": NameString, "channelist": finalChannelList})
	}

}

func FormResponses(c *gin.Context) {

	var (
		limt   int
		offset int
		filter forms.Filter
	)

	//get values from url query
	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")
	filter.TicketNo = c.Query("ticket")
	filter.FromDate = c.Query("fromdate")
	filter.ToDate = c.Query("todate")

	var filterflag string
	if filter.Keyword != "" || filter.TicketNo != "" || filter.FromDate != "" || filter.ToDate != "" {
		filterflag = "true"
	} else {
		filterflag = "false"
	}

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	permisison, perr := NewAuth.IsGranted("Forms Builder", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("cta status authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("formbuilders authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	menu := NewMenuController(c)

	// ModuleName, TabName, _ := ModuleRouteName(c)

	ModuleName := "Forms Builder"

	translate, _ := TranslateHandler(c)

	ctabannerval, _ := c.Cookie("ctabanner")

	if ctabannerval == "" {

		ctabannerval = "true"
	}
	ctadesc, _ := c.Cookie("ctadesc")

	if ctadesc == "" {

		ctadesc = "true"
	}

	var Formsresponselist []forms.TblFormResponses

	responselist, count, err := FormConfig.OverAllFormResponses(limt, offset, filter, TenantId)
	if err != nil {
		fmt.Println("overall responses not fetched:", err)
	}

	for _, val := range responselist {

		val.DateString = val.CreatedOn.In(TZONE).Format(Datelayout)

		Formsresponselist = append(Formsresponselist, val)

	}

	//pagination calc
	paginationendcount := len(responselist) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(pageno, int(count), limt)

	c.HTML(200, "form-responses.html", gin.H{"Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3,
	}, "csrf": csrf.GetToken(c), "Menu": menu, "ctabanner": ctabannerval, "ctadesc": ctadesc, "translate": translate, "title": ModuleName, "linktitle": "Forms", "Tabmenu": "CTA", "Responses": Formsresponselist, "Count": count, "Paginationendcount": paginationendcount, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Paginationstartcount": paginationstartcount, "Limit": limt, "SearchTrues": filterflag, "Filter": filter})
}

func ResponseDetail(c *gin.Context) {

	ticket := c.Param("ticket")

	response, err := FormConfig.ResponseDetail(ticket, TenantId)
	if err != nil {
		fmt.Println("response not fetched:", err)
	}

	var result []map[string]interface{}

	cerr := json.Unmarshal([]byte(response.FormResponse), &result)
	if cerr != nil {
		fmt.Printf("string to json conversion error %v", cerr)
	}

	var name string
	var email string
	var message string

	for _, item := range result {
		if item["question"] == "Name" {
			name = item["value"].(string)
		}
		if item["question"] == "Email" {
			email = item["value"].(string)
		}
		if item["question"] == "Message" {
			message = item["value"].(string)
		}
	}

	DateString := response.CreatedOn.In(TZONE).Format(Datelayout)

	messages := template.HTML(message)

	replyforresponse, err := FormConfig.ReplyForResponseList(ticket, TenantId)
	if err != nil {
		fmt.Printf("no data found %v", err)
	}

	for i, val := range replyforresponse {

		replyforresponse[i].DateString =
			replyforresponse[i].CreatedOn.In(TZONE).Format(Datelayout)

		replyforresponse[i].Replycontent = template.HTML(val.Reply)

	}

	MAIL_USERNAME := os.Getenv("MAIL_USERNAME")

	menu := NewMenuController(c)

	ModuleName := "Forms Builder"

	translate, _ := TranslateHandler(c)

	c.HTML(200, "form-responses-detail.html", gin.H{"csrf": csrf.GetToken(c), "Menu": menu, "translate": translate, "title": ModuleName, "linktitle": "Forms", "Tabmenu": "CTA", "response": response, "name": name, "email": email, "message": messages, "DateString": DateString, "replyforresponse": replyforresponse, "MAIL_USERNAME": MAIL_USERNAME})
}

func CloseTicket(c *gin.Context) {

	ticket := c.Param("ticket")

	_, err := FormConfig.Closeticket(ticket, TenantId)
	if err != nil {
		fmt.Println("ticket not yet closed:", err)
		return
	}

	c.SetCookie("get-toast", "Ticket Closed Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(303, "/admin/cta/form-responses")

}

func ReplyForResponses(c *gin.Context) {

	htmlContent := c.PostForm("htmlContent")
	email := c.PostForm("email")
	ticket := c.PostForm("ticket")
	username := c.PostForm("username")

	// fmt.Println("htmlContent::", htmlContent, "email::", email, "ticket::", ticket)
	permission, perr := NewAuth.IsGranted("Forms Builder", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("formbuilders authorization error:%s", perr)
	}

	if !permission {
		ErrorLog.Printf("formbuilders authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	createdon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	userid := c.GetInt("userid")

	reply := forms.TblReplyForResponse{
		Ticket:    ticket,
		Reply:     htmlContent,
		CreatedBy: userid,
		CreatedOn: createdon,
		TenantId:  TenantId,
	}

	status, err := FormConfig.ReplyForResponses(reply)
	if err != nil {
		ErrorLog.Printf("formbuilders error: %s", perr)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.Redirect(http.StatusMovedPermanently, "/admin/cta/form-responses")
	}

	linkedin := os.Getenv("LINKEDIN")
	facebook := os.Getenv("FACEBOOK")
	twitter := os.Getenv("TWITTER")
	youtube := os.Getenv("YOUTUBE")
	insta := os.Getenv("INSTAGRAM")
	var url_prefix = os.Getenv("BASE_URL")

	data := map[string]interface{}{
		"firstname":     username,
		"useremail":     email,
		"message":       htmlContent,
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
	go RepliedSuccessfully(&wg, data, email, TenantId, "Replied Successfully")

	c.JSON(200, gin.H{"status": status})

}

// Default Blocks list

func DefaultCtaList(c *gin.Context) {

	var (
		filter models.Filter
		limt   int
		offset int
	)

	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")
	filter.ChannelName = strings.Trim(c.DefaultQuery("channel", ""), "")

	var filterflag bool
	var keywordtrue bool
	if filter.Keyword != "" || filter.ChannelName != "" {
		filterflag = true

	} else {
		filterflag = false
	}

	if filter.Keyword != "" {
		keywordtrue = true
	}
	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}
	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	endurl := os.Getenv("MASTER_CTA_ENDPOINTURL")

	req, err := http.NewRequest("GET", endurl, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request: " + err.Error()})
		return
	}
	query := req.URL.Query()
	if filterflag {
		query.Add("keyword", filter.Keyword)
		query.Add("channel", filter.ChannelName)
	}
	query.Add("limit", strconv.Itoa(limt))
	query.Add("offset", strconv.Itoa(offset))
	query.Add("page", strconv.Itoa(pageno))
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	masterconnect := true

	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("Error connecting to master server:", err)
		masterconnect = false
	} else {
		defer resp.Body.Close()
	}

	var responseData ResponseData
	if masterconnect {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err == nil {
			fmt.Println("Error response:", err)
			resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			err = json.NewDecoder(resp.Body).Decode(&responseData)
			if err != nil {
				masterconnect = false
			}
		} else {
			masterconnect = false
		}
	}

	if !masterconnect {
		responseData = ResponseData{
			DefaultctaList: []models.TblMstrForms{},
			AllctaList:     []models.TblMstrForms{},
			FinalctaList:   []models.TblMstrForms{},
			BlockCount:     0,
		}
	}

	permisison, perr := NewAuth.IsGranted("Blocks", auth.CRUD, TenantId)

	if perr != nil {
		ErrorLog.Printf("block collection list authorization error: %s", perr)
	}

	if perr != nil {
		ErrorLog.Printf("block authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		ctabannerval, _ := c.Cookie("ctabanner")

		if ctabannerval == "" {

			ctabannerval = "true"
		}
		channelist, _, err := ChannelConfig.ListChannel(channels.Channels{Limit: 200, Offset: 0, IsActive: true, TenantId: TenantId})
		if err != nil {
			ErrorLog.Printf("%v: %v", ErrGettingTemplateModuleList, err)

		}
		var finalChannelList []channels.Tblchannel

		for _, chnList := range channelist {
			blockCount := 0

			for _, blockInfo := range responseData.AllctaList {
				if blockInfo.ChannelName == chnList.ChannelName {

					blockCount++
				}
			}

			chnList.EntriesCount = blockCount

			finalChannelList = append(finalChannelList, chnList)
		}
		var paginationendcount = len(responseData.DefaultctaList) + offset
		paginationstartcount := offset + 1
		Previous, Next, PageCount, Page := Pagination(pageno, int(responseData.BlockCount), limt)
		menu := NewMenuController(c)
		// ModuleName, TabName, _ := ModuleRouteName(c)
		ModuleName := "Forms Builder"

		translate, _ := TranslateHandler(c)
		var active string
		storagetype, err := GetSelectedType()

		if err != nil {
			fmt.Printf("blocklist getting storagetype error: %s", err)
		}

		c.HTML(200, "defaultcta.html", gin.H{"Menu": menu, "Pagination": PaginationData{
			NextPage:     pageno + 1,
			PreviousPage: pageno - 1,
			TotalPages:   PageCount,
			TwoAfter:     pageno + 2,
			TwoBelow:     pageno - 2,
			ThreeAfter:   pageno + 3,
		}, "ctabanner": ctabannerval, "chnid": filter.ChannelName, "keywordtrue": keywordtrue, "Count": responseData.BlockCount, "Searchtrue": filterflag, "linktitle": "Forms", "Limit": limt, "UserId": c.GetInt("userid"), "Formlist": responseData.FinalctaList, "Previous": Previous, "Page": Page, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "StorageType": storagetype.SelectedType, "Filter": filter, "title": ModuleName, "csrf": csrf.GetToken(c), "Active": active, "translate": translate, "channelist": finalChannelList})

	}

}
func AddForms(c *gin.Context) {

	menu := NewMenuController(c)

	_, TabName, _ := ModuleRouteName(c)

	ModuleName := "Forms Builder"

	translate, _ := TranslateHandler(c)

	channelist, _, err := ChannelConfig.ListChannel(channels.Channels{Limit: 200, Offset: 0, IsActive: true, TenantId: TenantId})
	if err != nil {
		ErrorLog.Printf("%v: %v", ErrGettingTemplateModuleList, err)

	}

	newpath := os.Getenv("BASE_URL")

	c.HTML(200, "create-event.html", gin.H{"Menu": menu, "uploadurl": newpath, "mode": "create", "translate": translate, "csrf": csrf.GetToken(c), "title": ModuleName, "linktitle": "Create Form", "Tabmenu": TabName, "channellist": channelist})

}

func CreateForms(c *gin.Context) {
	userid := c.GetInt("userid")
	button := c.PostForm("button")
	form := c.PostForm("form")
	Title := c.PostForm("title")
	description := c.PostForm("description")
	formimages := c.PostForm("formimages")
	channelids := c.PostFormArray("channelid[]")
	channelnames := c.PostFormArray("channelname[]")
	channelidString := strings.Join(channelids, ",")
	channelnameString := strings.Join(channelnames, ",")
	image := c.PostForm("image")
	imagename := c.PostForm("imagename")

	if Title == "" {

		Title = "Form"
	}

	var status string

	switch button {
	case "publish-form":
		status = "1"
	case "save-form":
		status = "0"
	default:
		status = "2"
	}

	Stats, err := strconv.Atoi(status)
	if err != nil {
		ErrorLog.Printf("invalid status: %s", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	var (
		imageByte []byte
		errs      error
	)

	var imageName, imagePath string

	tenantDetails, tenanterr := NewTeam.GetTenantDetails(TenantId)

	if tenanterr != nil {
		fmt.Println(tenanterr)
	}

	// fmt.Println("checkformcreate:", formimages)

	if formimages != "" {
		imageName, imagePath, imageByte, errs = ConvertBase64toByte(formimages, "cta")
		if errs != nil {
			fmt.Println(err)
		}

		imagePath = tenantDetails.S3FolderName + imagePath

		uerr := storagecontroller.UploadCropImageS3(imageName, imagePath, imageByte)
		if uerr != nil {
			c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
			c.Redirect(301, "/admin/cta")
		}
	}

	tblForm := forms.TblForm{
		FormTitle:            Title,
		FormData:             form,
		Status:               Stats,
		IsActive:             1,
		CreatedBy:            userid,
		TenantId:             TenantId,
		ChannelId:            channelidString,
		ChannelName:          channelnameString,
		FormPreviewImagepath: image,
		FormPreviewImagename: imagename,
		FormDescription:      description,
		ImageName:            imageName,
		ImagePath:            imagePath,
	}

	permission, perr := NewAuth.IsGranted("Forms Builder", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("formbuilders authorization error:%s", perr)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if !permission {
		ErrorLog.Printf("formbuilders authorization error")
		c.Redirect(http.StatusMovedPermanently, "/403-page")
		return
	}

	_, err1 := FormConfig.CreateForms(tblForm)
	if err1 != nil {
		fmt.Println("checkformtitle")
		ErrorLog.Printf("form creation error: %s", err)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.JSON(200, gin.H{"data": "/admin/cta"})
		return
	}

	c.SetCookie("get-toast", "Form Created Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.JSON(200, gin.H{"data": "/admin/cta"})
}

func Status(c *gin.Context) {

	Id := c.Query("id")
	Status := c.Query("status")
	pageno := c.Query("page")
	modifiedby := c.GetInt("userid")

	permission, perr := NewAuth.IsGranted("Forms Builder", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("formbuilders authorization error:%s", perr)
	}

	if !permission {
		ErrorLog.Printf("formbuilders authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	id, _ := strconv.Atoi(Id)
	var status int
	var homeurl string
	var toast string

	if Status == "1" {

		status = 2

		toast = "Form Unpublished Successfully"

		if pageno != "" {

			homeurl = "/admin/cta?page=" + pageno

		} else {

			homeurl = "/admin/cta"

		}

	} else if Status == "2" {

		status = 1

		toast = "Form Published Successfully"

		if pageno != "" {

			homeurl = "/admin/cta/unpublished?page=" + pageno

		} else {

			homeurl = "/admin/cta/unpublished"

		}

	} else if Status == "0" {

		status = 1

		toast = "Form Published Successfully"

		if pageno != "" {

			homeurl = "/admin/cta/draft?page=" + pageno

		} else {

			homeurl = "/admin/cta/draft"

		}

	}

	err := FormConfig.StatusChange(id, status, modifiedby, TenantId)
	if err != nil {
		ErrorLog.Printf("formbuilders error: %s", perr)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.Redirect(http.StatusMovedPermanently, homeurl)
	}
	c.SetCookie("get-toast", toast, 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(http.StatusMovedPermanently, homeurl)
}

func DeleteForm(c *gin.Context) {

	Id := c.Query("id")
	Status := c.Query("status")
	pageno := c.Query("page")
	deletedby := c.GetInt("userid")

	permission, perr := NewAuth.IsGranted("Forms Builder", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("formbuilders authorization error:%s", perr)
	}

	if !permission {
		ErrorLog.Printf("formbuilders authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	id, _ := strconv.Atoi(Id)
	var homeurl string

	if Status == "1" {

		if pageno != "" {

			homeurl = "/admin/cta?page=" + pageno

		} else {

			homeurl = "/admin/cta"

		}

	} else if Status == "2" {

		if pageno != "" {

			homeurl = "/admin/cta/unpublished?page=" + pageno

		} else {

			homeurl = "/admin/cta/unpublished"

		}

	} else if Status == "0" {

		if pageno != "" {

			homeurl = "/admin/cta/draft?page=" + pageno

		} else {

			homeurl = "/admin/cta/draft"

		}

	}

	err := FormConfig.Formdelete(id, deletedby, TenantId)
	if err != nil {
		ErrorLog.Printf("formbuilders error: %s", perr)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.Redirect(http.StatusMovedPermanently, homeurl)
	}
	c.SetCookie("get-toast", "Form Deleted Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(http.StatusMovedPermanently, homeurl)

}

func FormEdit(c *gin.Context) {

	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	permission, perr := NewAuth.IsGranted("Forms Builder", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("formbuilders authorization error:%s", perr)
	}

	if !permission {
		ErrorLog.Printf("formbuilders authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	forms, err := FormConfig.FormsEdit(Id, TenantId)
	if err != nil {
		ErrorLog.Printf("formbuilders error: %s", perr)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.Redirect(http.StatusMovedPermanently, "/admin/cta")
	}

	menu := NewMenuController(c)

	_, TabName, _ := ModuleRouteName(c)

	ModuleName := "Forms Builder"

	translate, _ := TranslateHandler(c)
	channelist, _, err := ChannelConfig.ListChannel(channels.Channels{Limit: 200, Offset: 0, IsActive: true, TenantId: TenantId})
	if err != nil {
		ErrorLog.Printf("%v: %v", ErrGettingTemplateModuleList, err)

	}
	newpath := os.Getenv("BASE_URL")

	c.HTML(200, "create-event.html", gin.H{"Menu": menu, "uploadurl": newpath, "formdetails": forms, "formid": forms.Id, "formdata": forms.FormData, "mode": "edit", "translate": translate, "csrf": csrf.GetToken(c), "title": ModuleName, "linktitle": "Edit Form", "Tabmenu": TabName, "channellist": channelist})

}

func FormDuplicate(c *gin.Context) {

	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	permission, perr := NewAuth.IsGranted("Forms Builder", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("formbuilders authorization error:%s", perr)
	}

	if !permission {
		ErrorLog.Printf("formbuilders authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	forms, err := FormConfig.FormsEdit(Id, TenantId)
	if err != nil {
		ErrorLog.Printf("formbuilders error: %s", perr)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.Redirect(http.StatusMovedPermanently, "/admin/cta")
	}

	menu := NewMenuController(c)

	ModuleName, TabName, _ := ModuleRouteName(c)

	translate, _ := TranslateHandler(c)

	c.HTML(200, "create-event.html", gin.H{"Menu": menu, "formdata": forms.FormData, "mode": "edit", "translate": translate, "csrf": csrf.GetToken(c), "title": ModuleName, "linktitle": ModuleName, "Tabmenu": TabName})

}

func FormUpdate(c *gin.Context) {
	id := c.PostForm("id")
	button := c.PostForm("button")
	form := c.PostForm("form")
	Title := c.PostForm("title")
	description := c.PostForm("description")
	formimages := c.PostForm("formimages")
	userid := c.GetInt("userid")
	channelids := (c.PostFormArray("channelid[]"))
	channelnames := c.PostFormArray("channelname[]")
	channelidString := strings.Join(channelids, ",")
	channelnameString := strings.Join(channelnames, ",")
	image := c.PostForm("image")
	imagename := c.PostForm("imagename")
	var status string

	if button == "publish-form" {
		status = "1"
	} else if button == "save-form" {
		status = "0"
	} else {
		status = "2"
	}

	Id, err := strconv.Atoi(id)
	if err != nil {
		ErrorLog.Printf("invalid form id: %s", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid form id"})
		return
	}
	Stats, err := strconv.Atoi(status)
	if err != nil {
		ErrorLog.Printf("invalid status: %s", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	formdata, _ := FormConfig.FormsEdit(Id, TenantId)

	var (
		imageByte []byte
		errs      error
	)

	var imageName, imagePath string

	if formdata.ImagePath != formimages {

		tenantDetails, tenanterr := NewTeam.GetTenantDetails(TenantId)

		if tenanterr != nil {
			fmt.Println(tenanterr)
		}

		if formimages != "" {
			imageName, imagePath, imageByte, errs = ConvertBase64toByte(formimages, "cta")
			if errs != nil {
				fmt.Println(err)
			}

			imagePath = tenantDetails.S3FolderName + imagePath

			uerr := storagecontroller.UploadCropImageS3(imageName, imagePath, imageByte)
			if uerr != nil {
				c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
				c.Redirect(301, "/admin/cta")
			}
		}

	}

	tblForm := forms.TblForm{
		Id:                   Id,
		FormTitle:            Title,
		FormData:             form,
		Status:               Stats,
		ModifiedBy:           userid,
		ChannelId:            channelidString,
		ChannelName:          channelnameString,
		FormPreviewImagepath: image,
		FormPreviewImagename: imagename,
		FormDescription:      description,
		ImageName:            imageName,
		ImagePath:            imagePath,
	}

	permission, perr := NewAuth.IsGranted("Forms Builder", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("formbuilders authorization error:%s", perr)
		c.Redirect(301, "/403-page")
		return
	}
	if !permission {
		ErrorLog.Printf("formbuilders authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	errr := FormConfig.UpdateForms(tblForm, TenantId)
	if errr != nil {
		ErrorLog.Printf("formbuilders error: %s", errr) // Corrected: Use errr
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.JSON(200, gin.H{"data": "/admin/cta"})
		return
	}

	c.SetCookie("get-toast", "Form Updated Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.JSON(200, gin.H{"data": "/admin/cta"})
}

func MultiDelete(c *gin.Context) {

	formids := c.PostFormArray("formids[]")
	pageno := c.PostForm("page")
	url := c.PostForm("urlvalue")
	userid := c.GetInt("userid")

	var homeurl string

	permisison, perr := NewAuth.IsGranted("Forms Builder", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("formbuilders authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("formbuilders authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	formIntIds := make([]int, len(formids))
	for i, id := range formids {
		intId, _ := strconv.Atoi(id)
		formIntIds[i] = intId
	}

	err := FormConfig.MultiSelectDeleteForm(formIntIds, userid, TenantId)
	if err != nil {
		ErrorLog.Printf("MultiSelectformDelete error: %s", err)
		c.JSON(200, gin.H{"value": false})
		return
	}

	if strings.Contains(url, "unpublished") {
		// fmt.Println("unpublished:")

		if pageno == "" {

			homeurl = "/admin/cta/unpublished"

		} else {

			_, TotalFormsCount, _, _ := FormConfig.FormBuildersList(10, 0, forms.Filter{}, TenantId, 2, 0, "", 0)

			page, _ := strconv.Atoi(pageno)
			page = page - 1
			multi := page * 10
			multiInt64 := int64(multi)
			if TotalFormsCount > multiInt64 {
				homeurl = "/admin/cta/unpublished/?page=" + pageno
			} else {
				pagee, _ := strconv.Atoi(pageno)
				_page := pagee - 1
				pages := strconv.Itoa(_page)
				homeurl = "/admin/cta/unpublished/?page=" + pages
			}

		}

	} else if strings.Contains(url, "draft") {
		// fmt.Println("draft:")
		if pageno == "" {

			homeurl = "/admin/cta/draft"

		} else {

			_, TotalFormsCount, _, _ := FormConfig.FormBuildersList(10, 0, forms.Filter{}, TenantId, 0, 0, "", 0)

			page, _ := strconv.Atoi(pageno)
			page = page - 1
			multi := page * 10
			multiInt64 := int64(multi)
			if TotalFormsCount > multiInt64 {
				homeurl = "/admin/cta/draft/?page=" + pageno
			} else {
				pagee, _ := strconv.Atoi(pageno)
				_page := pagee - 1
				pages := strconv.Itoa(_page)
				homeurl = "/admin/cta/draft/?page=" + pages
			}

		}

	} else {
		// fmt.Println("published:")

		if pageno == "" {

			homeurl = "/admin/cta"

		} else {

			_, TotalFormsCount, _, _ := FormConfig.FormBuildersList(10, 0, forms.Filter{}, TenantId, 1, 0, "", 0)

			page, _ := strconv.Atoi(pageno)
			page = page - 1
			multi := page * 10
			multiInt64 := int64(multi)
			if TotalFormsCount > multiInt64 {
				homeurl = "/admin/cta/?page=" + pageno
			} else {
				pagee, _ := strconv.Atoi(pageno)
				_page := pagee - 1
				pages := strconv.Itoa(_page)
				homeurl = "/admin/cta/?page=" + pages
			}

		}

	}

	c.JSON(200, gin.H{"value": true, "url": homeurl})
}

func MultiSelectStatusChange(c *gin.Context) {

	formids := c.PostFormArray("formids[]")
	page := c.PostForm("page")
	url := c.PostForm("urlvalue")
	userid := c.GetInt("userid")

	var homeurl string
	var status int
	var toast string

	if url == "unpublished" {

		status = 1
		toast = "Form Published Successfully"
		if page == "" {

			homeurl = "/admin/cta/unpublished"

		} else {

			homeurl = "/admin/cta/unpublished/?page=" + page

		}

	} else if url == "draft" {

		status = 1
		toast = "Form Published Successfully"
		if page == "" {

			homeurl = "/admin/cta/draft"

		} else {

			homeurl = "/admin/cta/draft/?page=" + page

		}

	} else {

		status = 2
		toast = "Form Unpublished Successfully"
		if page == "" {

			homeurl = "/admin/cta/"

		} else {

			homeurl = "/admin/cta/?page=" + page

		}

	}

	permisison, perr := NewAuth.IsGranted("Forms Builder", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("formbuilders authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("formbuilders authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	formIntIds := make([]int, len(formids))
	for i, id := range formids {
		intId, _ := strconv.Atoi(id)
		formIntIds[i] = intId
	}

	err := FormConfig.MultiSelectStatus(formIntIds, status, userid, TenantId)
	if err != nil {
		ErrorLog.Printf("MultiSelectformDelete error: %s", err)
		c.JSON(200, gin.H{"value": false})
		return
	}

	c.JSON(200, gin.H{"value": true, "url": homeurl, "toast": toast})

}

func Formdetail(c *gin.Context) {

	var (
		limt   int
		offset int
		filter forms.Filter
	)

	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")
	filter.FromDate = c.Query("fromdate")
	filter.ToDate = c.Query("todate")
	entryid := c.Query("entryid")
	// fmt.Println("entryid:", entryid)
	EntryId, _ := strconv.Atoi(entryid)

	var filterflag bool
	if filter.Keyword != "" || filter.FromDate != "" || filter.ToDate != "" {
		filterflag = true
	} else {
		filterflag = false
	}

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	id := c.Param("id")
	userid := c.GetInt("userid")

	FormId, _ := strconv.Atoi(id)

	var Formsresponselist []forms.TblFormResponses

	responselist, TotalResponseCount, formname, _ := FormConfig.FormDetailLists(limt, offset, forms.Filter(filter), FormId, userid, EntryId, TenantId)

	// fmt.Println("Formdetail::", responselist)

	var FormData []string

	for _, val := range responselist {

		val.DateString = val.CreatedOn.In(TZONE).Format(Datelayout)

		Formsresponselist = append(Formsresponselist, val)

		responses := val.FormResponse

		FormData = append(FormData, responses)

	}
	var input [][]map[string]interface{}

	for _, dataStr := range FormData {
		var temp []map[string]interface{}
		if err := json.Unmarshal([]byte(dataStr), &temp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format: " + err.Error()})
			return
		}
		input = append(input, temp)
	}

	//pagination calc
	paginationendcount := len(responselist) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(pageno, int(TotalResponseCount), limt)

	menu := NewMenuController(c)

	_, TabName, _ := ModuleRouteName(c)

	ModuleName := "Forms Builder"

	translate, _ := TranslateHandler(c)

	c.HTML(200, "form-detail.html", gin.H{"Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3,
	}, "Menu": menu, "FormId": FormId, "FormName": formname, "Responses": Formsresponselist, "Responselist": input, "SearchTrues": filterflag, "Count": TotalResponseCount, "Limit": limt, "Paginationendcount": paginationendcount, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Filter": filter, "Paginationstartcount": paginationstartcount, "translate": translate, "title": ModuleName, "linktitle": ModuleName, "Tabmenu": TabName})

}

func FormIsactive(c *gin.Context) {

	//get data from html
	id, _ := strconv.Atoi(c.Request.PostFormValue("id"))
	val, _ := strconv.Atoi(c.Request.PostFormValue("isactive"))
	userid := c.GetInt("userid")

	permisison, perr := NewAuth.IsGranted("Forms Builder", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("cta status authorization error: %s", perr)
	}

	if permisison {
		flg, err := FormConfig.ChangeFormStatus(id, val, userid, TenantId)
		if err != nil {
			ErrorLog.Printf("cta status error: %s", perr)
			json.NewEncoder(c.Writer).Encode(flg)
			return
		}
		json.NewEncoder(c.Writer).Encode(flg)
		return

	}

	c.Redirect(301, "/403-page")
}

func AddCollection(c *gin.Context) {

	title := c.PostForm("title")
	cimg := c.PostForm("cimg")
	desc := c.PostForm("desc")
	smallimg := c.PostForm("smallimg")
	fdata := c.PostForm("fdata")
	channelname := c.PostForm("channelname")

	permisison, perr := NewAuth.IsGranted("Forms Builder", auth.CRUD, TenantId)

	if perr != nil {
		ErrorLog.Printf("Forms Builder authorization error :%s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Forms Builder authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}
	channeldetails, err := ChannelConfig.GetchannelByName(channelname, TenantId)
	if err != nil {
		ErrorLog.Printf("%v: %v", ErrGettingTemplateModuleList, err)

	}

	if channeldetails.ChannelName != channelname {

		c.JSON(200, gin.H{"data": "Pleaseaddtherequiredchannelbeforeproceeding"})
		return
	}

	channelid := strconv.Itoa(channeldetails.Id)
	if permisison {

		userid := c.GetInt("userid")

		tblForm := forms.TblForm{
			FormTitle:            title,
			FormData:             fdata,
			Status:               1,
			IsActive:             1,
			CreatedBy:            userid,
			TenantId:             TenantId,
			ChannelId:            channelid,
			ChannelName:          channelname,
			FormDescription:      desc,
			FormImagePath:        smallimg,
			FormPreviewImagepath: cimg,
		}

		_, err := FormConfig.CreateForms(tblForm)

		if err != nil {

			c.JSON(200, gin.H{"data": false})
			return

		}

		c.JSON(200, gin.H{"data": true})

	}
}

func RemoveCta(c *gin.Context) {

	id := c.Param("id")

	// fmt.Println("form not working", id)

	permisison, perr := NewAuth.IsGranted("Forms Builder", auth.CRUD, TenantId)

	if perr != nil {
		ErrorLog.Printf("Forms Builder authorization error :%s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Forms Builder authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		userid := c.GetInt("userid")

		_, err := FormConfig.Removectatomycollecton(id, TenantId, userid)

		if err != nil {
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
			c.JSON(200, false)
			return
		}

		c.SetCookie("get-toast", "Collection Remove Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, "/admin/cta")

	}
}

func CtaPreview(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	// fmt.Println("form not working", id)

	permisison, perr := NewAuth.IsGranted("Forms Builder", auth.CRUD, TenantId)

	if perr != nil {
		ErrorLog.Printf("Forms Builder authorization error :%s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Forms Builder authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	formdata, err := FormConfig.GetCtaById(id)

	if err != nil {
		ErrorLog.Printf("cta get details error: %s", err)
		c.JSON(200, gin.H{"value": false})
		return
	}

	c.JSON(200, gin.H{"value": true, "Formdata": formdata})

}
