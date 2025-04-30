package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"spurt-cms/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spurtcms/auth"
	chn "github.com/spurtcms/channels"
	csrf "github.com/utrack/gin-csrf"
)

var CurrentPage int

/*Channel List*/
func ChannelList(c *gin.Context) {

	var (
		limt   int
		offset int
	)

	// get data from html form data
	limit := c.Query("limit")
	keyword := strings.TrimSpace(c.DefaultQuery("keyword", ""))
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	var filterflag bool
	if keyword != "" {
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

	permisison, perr := NewAuth.IsGranted("Channels", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("channellist authorization error: %s", perr)
	}

	if !permisison { //check permission if permission false redirect to 403 page
		c.Redirect(301, "/403-page")
		return
	}
	channelBannervalue, _ := c.Cookie("channelbanner")

	if channelBannervalue == "" {

		channelBannervalue = "true"
	}
	ChannelConfig.DataAccess = c.GetInt("dataaccess")
	ChannelConfig.Userid = c.GetInt("userid")

	routename := c.FullPath()

	var htmlname string

	if strings.Contains(routename, "defaultchannels") {

		htmlname = "defaultchannel.html"

		TenantId = ""

	} else {

		htmlname = "channels.html"

	}

	channelslist, channelcount, err := ChannelConfig.ListChannel(chn.Channels{Limit: Limit, Offset: offset, Keyword: keyword, IsActive: false, CreateOnly: true, TenantId: TenantId, Count: true, AuthorDetail: true})
	var Finalchannelslist []chn.Tblchannel

	for _, val := range channelslist {

		_, entrcount, _, _ := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: val.Id, Limit: 0, Offset: 0}, TenantId)
		val.EntriesCount = int(entrcount)

		list := val
		if val.ModifiedOn.IsZero() {
			list.DateString = val.CreatedOn.In(TZONE).Format(Datelayout)
		} else {
			list.DateString = val.ModifiedOn.In(TZONE).Format(Datelayout)
		}

		var first = val.AuthorDetails.FirstName
		var last = val.AuthorDetails.LastName
		var firstn = strings.ToUpper(first[:1])
		var lastn string
		if val.LastName != "" {
			lastn = strings.ToUpper(last[:1])
		}
		var Name = firstn + lastn
		list.NameString = Name
		list.Username = val.AuthorDetails.Username

		// firstname= val.AuthorDetails.FirstName

		imgcontain := "/image-resize?name="

		if val.AuthorDetails.ProfileImagePath != "" {
			userimg := val.AuthorDetails.ProfileImagePath
			imgflag := strings.Contains(userimg, imgcontain)
			if !imgflag {
				list.ProfileImagePath = "/image-resize?name=" + val.AuthorDetails.ProfileImagePath
			}
		}
		Finalchannelslist = append(Finalchannelslist, list)
	}

	if err != nil {
		ErrorLog.Printf("channellist error :%s", err)
	}

	//pagination calc
	paginationendcount := len(channelslist) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(pageno, int(channelcount), limt)

	menu := NewMenuController(c)
	translate, _ := TranslateHandler(c)
	ModuleName, TabName, _ := ModuleRouteName(c)
	CurrentPage = pageno

	c.HTML(200, htmlname, gin.H{"Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3,
	}, "channelbanner": channelBannervalue, "Menu": menu, "csrf": csrf.GetToken(c), "channel": Finalchannelslist, "Searchtrue": filterflag, "count": channelcount, "HeadTitle": translate.Channell.Channels, "filter": keyword, "translate": translate, "Channelsmenu": true, "Cmsmenu": true, "title": ModuleName, "Tabmenu": TabName, "linktitle": ModuleName, "Paginationendcount": paginationendcount, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Paginationstartcount": paginationstartcount, "Limit": limt})

}
func DefaultChannelList(c *gin.Context) {

	var (
		limt   int
		offset int
	)

	// get data from html form data
	limit := c.Query("limit")
	keyword := strings.TrimSpace(c.DefaultQuery("keyword", ""))
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	var filterflag bool
	if keyword != "" {
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

	permisison, perr := NewAuth.IsGranted("Channels", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("channellist authorization error: %s", perr)
	}

	if !permisison { //check permission if permission false redirect to 403 page
		c.Redirect(301, "/403-page")
		return
	}
	channelBannervalue, _ := c.Cookie("channelbanner")

	if channelBannervalue == "" {

		channelBannervalue = "true"
	}
	ChannelConfig.DataAccess = c.GetInt("dataaccess")
	ChannelConfig.Userid = c.GetInt("userid")

	endurl := os.Getenv("MASTER_CHANNELS_ENDPOINTURL")

	responseData, err := ChannelConfig.DefaultChannelList(endurl, Limit, offset, chn.Filter{Keyword: keyword})

	fmt.Println(responseData)
	var paginationendcount = responseData.BlockCount + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(pageno, int(responseData.BlockCount), limt)
	menu := NewMenuController(c)
	ModuleName, TabName, _ := ModuleRouteName(c)
	translate, _ := TranslateHandler(c)

	if err != nil {
		fmt.Printf("blocklist getting storagetype error: %s", err)
	}

	c.HTML(200, "defaultchannel.html", gin.H{"Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3,
	}, "channelbanner": channelBannervalue, "Menu": menu, "csrf": csrf.GetToken(c), "channel": template.HTML(responseData.Channelliststring), "Searchtrue": filterflag, "count": responseData.BlockCount, "HeadTitle": translate.Channell.Channels, "filter": keyword, "translate": translate, "Channelsmenu": true, "Cmsmenu": true, "title": ModuleName, "Tabmenu": TabName, "linktitle": ModuleName, "Paginationendcount": paginationendcount, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Paginationstartcount": paginationstartcount, "Limit": limt})

}

/*Create Channel*/
func CreateChannelPage(c *gin.Context) {

	permisison, perr := NewAuth.IsGranted("Channels", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("createchannel authorization error: %s", perr)
	}

	if permisison {

		field, err := ChannelConfig.GetAllMasterFieldType(TenantId)
		if err != nil {
			ErrorLog.Printf("get allmasterfield error: %s", err)
		}

		AllCategorieswithSubCategories, _ := CategoryConfig.AllCategoriesWithSubList(TenantId)

		menu := NewMenuController(c)
		ModuleName, _, _ := ModuleRouteName(c)
		translate, _ := TranslateHandler(c)

		uuid := (uuid.New()).String()

		arr := strings.Split(uuid, "-")

		UniqueId := arr[len(arr)-1]

		c.HTML(200, "addchannel.html", gin.H{"Menu": menu, "linktitle": "Create Channel", "title": ModuleName, "csrf": csrf.GetToken(c), "Fields": field, "Button": "Save", "AllCategories": AllCategorieswithSubCategories, "Title": "Create Channel", "Back": "/settings/channels/channellist", "HeadTitle": "Create Channel", "translate": translate, "Channelsmenu": true, "Cmsmenu": true, "UniqueId": UniqueId})

		return

	}

	c.Redirect(301, "/403-page")

}

func CreateChannel(c *gin.Context) {

	//get data from html
	channelname := c.Request.PostFormValue("channelname")
	channeluniqueid := c.Request.PostFormValue("channeluniqueid")
	channeldesc := c.Request.PostFormValue("channeldesc")
	sectionvalue := c.Request.FormValue("sections")
	fval := c.Request.PostFormValue("fiedlvalue")
	categoryvalue := c.PostFormArray("categoryvalue[]")
	userid := c.GetInt("userid")

	type sections struct {
		Fiedlvalue []Section `json:"sections"`
	}

	type field struct {
		Fiedlvalue []Fiedlvalue `json:"fiedlvalue"`
	}

	var (
		Sections sections
		fieldval field
	)

	if err1 := json.Unmarshal([]byte(sectionvalue), &Sections); err1 != nil {
		ErrorLog.Printf("section unmarshal create channel error: %s", err1)
	}

	if err2 := json.Unmarshal([]byte(fval), &fieldval); err2 != nil {
		ErrorLog.Printf("fieldval unmarshal create channel error: %s", err2)
	}

	if channelname == "" || channeldesc == "" {
		ErrorLog.Printf("create channel error mandatory field ")
		json.NewEncoder(c.Writer).Encode(false)
		return
	}

	permisison, perr := NewAuth.IsGranted("Channels", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("channelcreate authorization error: %s", perr)
	}

	if permisison {

		var sectionss []chn.Section
		for _, val := range Sections.Fiedlvalue {
			var sect chn.Section
			sect.MasterFieldId = val.MasterFieldId
			sect.OrderIndex = val.OrderIndex
			sect.SectionId = val.SectionId
			sect.SectionName = val.SectionName
			sect.SectionNewId = val.SectionNewId
			sectionss = append(sectionss, sect)
		}

		var FieldValuesss []chn.Fiedlvalue
		for _, val := range fieldval.Fiedlvalue {
			var fieldval chn.Fiedlvalue
			fieldval.CharacterAllowed = val.CharacterAllowed
			fieldval.DateFormat = val.DateFormat
			fieldval.FieldId = val.FieldId
			fieldval.FieldName = val.FieldName
			fieldval.IconPath = val.IconPath
			fieldval.MasterFieldId = val.MasterFieldId
			fieldval.NewFieldId = val.NewFieldId
			fieldval.OrderIndex = val.OrderIndex
			fieldval.SectionId = val.SectionId
			fieldval.SectionNewId = val.SectionNewId
			fieldval.TimeFormat = val.TimeFormat
			fieldval.Url = val.Url
			fieldval.Mandatory = val.Mandatory

			var choptions []chn.OptionValues
			for _, opt := range val.OptionValue {
				var optss chn.OptionValues
				optss.FieldId = opt.FieldId
				optss.Id = opt.Id
				optss.NewFieldId = opt.NewFieldId
				optss.NewId = opt.NewId
				optss.Value = opt.Value
				optss.OrderIndex = opt.OrderIndex
				choptions = append(choptions, optss)
			}

			fieldval.OptionValue = choptions
			FieldValuesss = append(FieldValuesss, fieldval)

		}

		Entry := "Entries"
		moduleid, err := models.Entryid(Entry, TenantId)
		if err != nil {
			fmt.Println(err)
		}

		newchannel, cerr := ChannelConfig.CreateChannel(chn.ChannelCreate{ChannelName: channelname, ChannelUniqueId: channeluniqueid, ChannelDescription: channeldesc, CategoryIds: categoryvalue, CreatedBy: userid}, moduleid, TenantId)

		if cerr != nil {
			ErrorLog.Printf("channelcreate error: %s", cerr)
		}

		ferr := ChannelConfig.CreateAdditionalFields(chn.ChannelAddtionalField{Sections: sectionss, FieldValues: FieldValuesss, CreatedBy: userid}, newchannel.Id, TenantId)

		if ferr != nil {
			ErrorLog.Printf("channelcreate additional field error: %s", ferr)
		}

		c.SetCookie("get-toast", "Channel Created Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		json.NewEncoder(c.Writer).Encode(true)

		return

	}

	c.Redirect(301, "/403-page")

}

/*Edit Channel*/
func EditChannel(c *gin.Context) {

	//get data from url query
	page := c.Query("page")
	id, _ := strconv.Atoi(c.Param("id"))

	permisison, perr := NewAuth.IsGranted("Channels", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("editchannel authorization error: %s", perr)
	}

	pageno := CurrentPage

	if permisison {

		chndata, FinalSelectedCategories, cerr := ChannelConfig.GetChannelsById(id, TenantId)
		if cerr != nil {
			ErrorLog.Printf("editchannel getchanneldata error: %s", perr)
		}

		channelname := chndata.ChannelName

		field, err := ChannelConfig.GetAllMasterFieldType(TenantId)
		if err != nil {
			ErrorLog.Printf("editchannel get all master field  error: %s", perr)
		}

		AllCategorieswithSubCategories, _ := CategoryConfig.AllCategoriesWithSubList(TenantId)
		fmt.Println("AllCategorieswithSubCategories:", AllCategorieswithSubCategories)
		fmt.Println("FinalSelectedCategories:", FinalSelectedCategories)

		menu := NewMenuController(c)
		translate, _ := TranslateHandler(c)
		ModuleName, _, _ := ModuleRouteName(c)

		c.HTML(200, "addchannel.html", gin.H{"Menu": menu, "Channelname": channelname, "translate": translate, "csrf": csrf.GetToken(c), "Fields": field, "Button": "Update", "channel": chndata, "AllCategories": AllCategorieswithSubCategories, "title": ModuleName, "linktitle": "Edit Channel", "Back": "/settings/channels/channellist", "Page": page, "HeadTitle": translate.Channell.Channels, "ChannelId": id, "SelectedCategories": FinalSelectedCategories, "Channelsmenu": true, "Cmsmenu": true, "Title": "Edit Channel", "CurrentPage": pageno})

		return

	}
	c.Redirect(301, "/403-page")

}

/*Get Field Data By channelId*/
func GetFieldDataByChannelId(c *gin.Context) {

	var id, _ = strconv.Atoi(c.Request.PostFormValue("id"))

	permisison, perr := NewAuth.IsGranted("Channels", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("getfieldata authorization error: %s", perr)
	}

	if permisison {

		Section, Fieldvalue, ferr := ChannelConfig.GetChannelsFieldsById(id, TenantId)

		if ferr != nil {
			ErrorLog.Printf("getfieldata error: %s", perr)
		}

		_, FinalSelectedCategories, cerr := ChannelConfig.GetChannelsById(id, TenantId)

		if cerr != nil {
			ErrorLog.Printf("getchannel error: %s", perr)
		}

		var idstr []string

		for _, val := range FinalSelectedCategories {

			var str string

			for index, cat := range val.Categories {
				if index+1 != len(val.Categories) {
					str += strconv.Itoa(cat.Id) + ","
				} else {
					str += strconv.Itoa(cat.Id)
				}
			}

			idstr = append(idstr, str)
		}

		json.NewEncoder(c.Writer).Encode(gin.H{"Section": Section, "FieldValue": Fieldvalue, "SelectedCategory": idstr})
		return

	}

	c.Redirect(301, "/403-page")

}

/*Update Channel*/
func UpdateChannel(c *gin.Context) {

	//get data from html form data
	channelid, _ := strconv.Atoi(c.Request.PostFormValue("id"))
	channelname := c.Request.PostFormValue("channelname")
	channeluniqueid := c.Request.PostFormValue("channeluniqueid")
	fmt.Println("channeluniqueid:", channeluniqueid)
	channeldesc := c.Request.PostFormValue("channeldesc")
	sectionvalue := c.Request.FormValue("sections")
	deletesection := c.Request.FormValue("deletesections")
	deleteoption := c.Request.PostFormValue("deleteoptions")
	fval := c.Request.PostFormValue("fiedlvalue")
	deletefval := c.Request.PostFormValue("deletefields")
	categoryvalue := c.PostFormArray("categoryvalue[]")
	userid := c.GetInt("userid")

	type sections struct {
		Fiedlvalue []Section `json:"sections"`
	}

	type field struct {
		Fiedlvalue []Fiedlvalue `json:"fiedlvalue"`
	}

	type delsections struct {
		Fiedlvalue []Section `json:"deletesecion"`
	}

	type delfield struct {
		Fiedlvalue []Fiedlvalue `json:"deletefields"`
	}
	//delete Option
	type deleteopt struct {
		Fiedlvalue []OptionValues `json:"deleteoption"`
	}

	//create and update
	var (
		Sections       sections
		fieldval       field
		deleteSections delsections
		deleteFields   delfield
		DeleteOpt      deleteopt
	)

	if secerr := json.Unmarshal([]byte(sectionvalue), &Sections); secerr != nil {
		ErrorLog.Printf("Updatechannel unmarshal section error: %s", secerr)
	}

	if fieerr := json.Unmarshal([]byte(fval), &fieldval); fieerr != nil {
		ErrorLog.Printf("Updatechannel unmarshal field error: %s", fieerr)
	}

	if delsecerr := json.Unmarshal([]byte(deletesection), &deleteSections); delsecerr != nil {
		ErrorLog.Printf("Updatechannel unmarshal delete section error: %s", delsecerr)
	}

	if delfieerr := json.Unmarshal([]byte(deletefval), &deleteFields); delfieerr != nil {
		ErrorLog.Printf("Updatechannel unmarshal delete field error: %s", delfieerr)
	}

	if deloper := json.Unmarshal([]byte(deleteoption), &DeleteOpt); deloper != nil {
		ErrorLog.Printf("Updatechannel unmarshal delete option error: %s", deloper)
	}

	permisison, perr := NewAuth.IsGranted("Channels", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("updatechannel authorization error: %s", perr)
	}

	if permisison {

		var Sectionss []chn.Section
		for _, val := range Sections.Fiedlvalue {
			var sec chn.Section
			sec.MasterFieldId = val.MasterFieldId
			sec.OrderIndex = val.OrderIndex
			sec.SectionId = val.SectionId
			sec.SectionName = val.SectionName
			sec.SectionNewId = val.SectionNewId
			Sectionss = append(Sectionss, sec)
		}

		var FieldValuess []chn.Fiedlvalue

		for _, val := range fieldval.Fiedlvalue {
			var fieldval chn.Fiedlvalue
			fieldval.CharacterAllowed = val.CharacterAllowed
			fieldval.DateFormat = val.DateFormat
			fieldval.FieldId = val.FieldId
			fieldval.FieldName = val.FieldName
			fieldval.IconPath = val.IconPath
			fieldval.MasterFieldId = val.MasterFieldId
			fieldval.NewFieldId = val.NewFieldId
			fieldval.OrderIndex = val.OrderIndex
			fieldval.SectionId = val.SectionId
			fieldval.SectionNewId = val.SectionNewId
			fieldval.TimeFormat = val.TimeFormat
			fieldval.Url = val.Url
			fieldval.Mandatory = val.Mandatory

			var choptions []chn.OptionValues
			for _, opt := range val.OptionValue {
				var optss chn.OptionValues
				optss.FieldId = opt.FieldId
				optss.Id = opt.Id
				optss.NewFieldId = opt.NewFieldId
				optss.NewId = opt.NewId
				optss.Value = opt.Value
				optss.OrderIndex = opt.OrderIndex
				choptions = append(choptions, optss)
			}

			fieldval.OptionValue = choptions
			FieldValuess = append(FieldValuess, fieldval)
		}

		var DeleteSectionss []chn.Section
		for _, val := range deleteSections.Fiedlvalue {
			var sec chn.Section
			sec.MasterFieldId = val.MasterFieldId
			sec.OrderIndex = val.OrderIndex
			sec.SectionId = val.SectionId
			sec.SectionName = val.SectionName
			sec.SectionNewId = val.SectionNewId
			DeleteSectionss = append(DeleteSectionss, sec)
		}

		var DeleteFieldValuess []chn.Fiedlvalue
		for _, val := range deleteFields.Fiedlvalue {
			var fieldval chn.Fiedlvalue
			fieldval.CharacterAllowed = val.CharacterAllowed
			fieldval.DateFormat = val.DateFormat
			fieldval.FieldId = val.FieldId
			fieldval.FieldName = val.FieldName
			fieldval.IconPath = val.IconPath
			fieldval.MasterFieldId = val.MasterFieldId
			fieldval.NewFieldId = val.NewFieldId
			fieldval.OrderIndex = val.OrderIndex
			fieldval.SectionId = val.SectionId
			fieldval.SectionNewId = val.SectionNewId
			fieldval.TimeFormat = val.TimeFormat
			fieldval.Url = val.Url

			var choptions []chn.OptionValues
			for _, opt := range val.OptionValue {
				var optss chn.OptionValues
				optss.FieldId = opt.FieldId
				optss.Id = opt.Id
				optss.NewFieldId = opt.NewFieldId
				optss.NewId = opt.NewId
				optss.Value = opt.Value
				choptions = append(choptions, optss)
			}
			fieldval.OptionValue = choptions
			DeleteFieldValuess = append(DeleteFieldValuess, fieldval)
		}

		var DeleteOptionsvalues []chn.OptionValues
		for _, val := range DeleteOpt.Fiedlvalue {

			var DeleteOptionsvalue chn.OptionValues
			DeleteOptionsvalue.FieldId = val.FieldId
			DeleteOptionsvalue.Id = val.Id
			DeleteOptionsvalue.NewFieldId = val.NewFieldId
			DeleteOptionsvalue.NewId = val.NewId
			DeleteOptionsvalue.Value = val.Value
			DeleteOptionsvalues = append(DeleteOptionsvalues, DeleteOptionsvalue)

		}
		if cerr := ChannelConfig.EditChannel(channelname, channeluniqueid, channeldesc, userid, channelid, categoryvalue, TenantId); cerr != nil {
			ErrorLog.Printf("edit channel error: %s", cerr)
		}

		ferr := ChannelConfig.UpdateChannelField(chn.ChannelUpdate{Sections: Sectionss, FieldValues: FieldValuess, Deletesections: DeleteSectionss, DeleteFields: DeleteFieldValuess, DeleteOptionsvalue: DeleteOptionsvalues, CategoryIds: categoryvalue, ModifiedBy: userid}, channelid, TenantId)

		if ferr != nil {
			ErrorLog.Printf("edit channel additional field error: %s", ferr)
		}
		c.SetCookie("get-toast", "Channel Updated Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		json.NewEncoder(c.Writer).Encode(true)
		return

	}

	c.Redirect(301, "/403-page")

}

/*Delete channel*/
func DeleteChannel(c *gin.Context) {

	channelid, _ := strconv.Atoi(c.Query("id"))
	userid := c.GetInt("userid")
	pageno := c.Query("page")

	var url string
	if pageno != "" {
		url = "/channels/?page=" + pageno
	} else {
		url = "/channels/"

	}

	// if channelid == 1 {
	// 	ErrorLog.Println("trying to delete default channel error")
	// 	c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
	// 	c.Redirect(301, "/channels/")
	// 	return
	// }

	permisison, perr := NewAuth.IsGranted("Channels", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("delete channel authorization error: %s", perr)
	}

	if permisison {

		ModuleName := strconv.Itoa(channelid)

		routename := "/entries/entrylist/" + ModuleName

		err := ChannelConfig.DeleteChannel(channelid, userid, routename, TenantId)

		if err != nil {
			ErrorLog.Printf("delete channel error: %s", perr)
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			c.Redirect(301, url)
			return
		}

		c.SetCookie("get-toast", "Channel Deleted Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, url)
		return

	}

	c.Redirect(301, "/403-page")

}

/*Channel status*/
func ChannelIsActiveStatus(c *gin.Context) {

	//get data from html
	id, _ := strconv.Atoi(c.Request.PostFormValue("id"))
	val, _ := strconv.Atoi(c.Request.PostFormValue("isactive"))
	userid := c.GetInt("userid")

	permisison, perr := NewAuth.IsGranted("Channels", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("channel status authorization error: %s", perr)
	}

	if permisison {
		flg, err := ChannelConfig.ChangeChannelStatus(id, val, userid, TenantId)
		if err != nil {
			ErrorLog.Printf("channel status error: %s", perr)
			json.NewEncoder(c.Writer).Encode(flg)
			return
		}
		json.NewEncoder(c.Writer).Encode(flg)
		return

	}

	c.Redirect(301, "/403-page")

}

// Channel Pagination list //
func PaginationList(c *gin.Context) {

	var Finalchannelslist []chn.Tblchannel

	offset, _ := strconv.Atoi(c.PostForm("offset"))
	loffset := offset * Limit

	flag := false

	channelslist, _, err := ChannelConfig.ListChannel(chn.Channels{Limit: Limit, Offset: loffset, IsActive: flag, TenantId: TenantId})

	for _, val := range channelslist {
		list := val
		if val.ModifiedOn.IsZero() {
			list.DateString = val.CreatedOn.In(TZONE).Format(Datelayout)
		} else {
			list.DateString = val.ModifiedOn.In(TZONE).Format(Datelayout)
		}
		Finalchannelslist = append(Finalchannelslist, list)
	}

	if err != nil {
		ErrorLog.Printf("channel pagination error: %s", err)
		json.NewEncoder(c.Writer).Encode(Finalchannelslist)
		return
	}

	json.NewEncoder(c.Writer).Encode(Finalchannelslist)

}

func ChannelType(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	ChannelType := c.Query("channelType")
	userid := c.GetInt("userid")

	var Channels chn.Tblchannel
	if ChannelType == "mychannels" {
		Channels = chn.Tblchannel{
			Id:          Id,
			ChannelType: "others",
			ModifiedBy:  userid,
		}
	} else {
		Channels = chn.Tblchannel{
			Id:          Id,
			ChannelType: "mychannels",
			ModifiedBy:  userid,
		}
	}
	err := ChannelConfig.ChannelType(Channels)

	if err != nil {

		ErrorLog.Printf("channel Type error: %s", err)
	}
	c.SetCookie("get-toast", "ChannelType Changed Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, "/channels/")
}

func AddChannelTomyCollection(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	permisison, perr := NewAuth.IsGranted("Channels", auth.CRUD, TenantId)

	if perr != nil {
		ErrorLog.Printf("Channels authorization error :%s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Channels authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	endurl := os.Getenv("MASTER_CHANNELS_ENDPOINTURL")

	req, err := http.NewRequest("GET", endurl, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request: " + err.Error()})
		return
	}
	query := req.URL.Query()

	query.Add("id", strconv.Itoa(id))
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

	type channeldetailsresponse struct {
		Channeldetail models.TblMstrchannel `json:"channeldetail"`
	}
	var responseData channeldetailsresponse
	var bodyBytes []byte
	if masterconnect {
		bodyBytes, err = io.ReadAll(resp.Body)
		if err != nil {
			masterconnect = false
			fmt.Println("Error reading response body:", err)
		} else {

			err = json.Unmarshal(bodyBytes, &responseData)
			if err != nil {
				masterconnect = false
				fmt.Println("Error decoding JSON:", err)
			}
		}
	}

	if !masterconnect {
		responseData = channeldetailsresponse{
			Channeldetail: models.TblMstrchannel{},
		}
	}
	Entry := "Entries"
	moduleid, err := models.Entryid(Entry, TenantId)
	if err != nil {
		fmt.Println(err)
	}
	var fieldValuesMap = map[string][]chn.Fiedlvalue{
		"blog": {
			{FieldName: "Excerpt", FieldId: 1, IconPath: "/public/img/text.svg", OrderIndex: 1, Mandatory: 0, MasterFieldId: 2},
			{FieldName: "Author Bio", FieldId: 2, IconPath: "/public/img/text.svg", OrderIndex: 2, Mandatory: 0, MasterFieldId: 2},
		},
		"ecommerce": {
			{FieldName: "Price", FieldId: 1, IconPath: "/public/img/text.svg", OrderIndex: 1, Mandatory: 0, MasterFieldId: 2},
			{FieldName: "Special Price", FieldId: 2, IconPath: "/public/img/text.svg", OrderIndex: 2, Mandatory: 0, MasterFieldId: 2},
			{FieldName: "SKU", FieldId: 3, IconPath: "/public/img/text.svg", OrderIndex: 3, Mandatory: 0, MasterFieldId: 2},
			{FieldName: "Quantity", FieldId: 4, IconPath: "/public/img/text.svg", OrderIndex: 4, Mandatory: 0, MasterFieldId: 2},
			{FieldName: "Stock", FieldId: 5, IconPath: "/public/img/select.svg", OrderIndex: 5, Mandatory: 0, MasterFieldId: 5},
		},
		"appointment": {
			{FieldName: "Date", FieldId: 1, IconPath: "/public/img/date.svg", OrderIndex: 1, Mandatory: 0, MasterFieldId: 6},
			{FieldName: "Time", FieldId: 2, IconPath: "/public/img/date-time.svg", OrderIndex: 2, Mandatory: 0, MasterFieldId: 5},
			{FieldName: "Contact Number", FieldId: 3, IconPath: "/public/img/text.svg", OrderIndex: 3, Mandatory: 0, MasterFieldId: 2},
			{FieldName: "Email", FieldId: 4, IconPath: "/public/img/text.svg", OrderIndex: 4, Mandatory: 0, MasterFieldId: 2},
			{FieldName: "DOB/Age", FieldId: 5, IconPath: "/public/img/date-time.svg", OrderIndex: 5, Mandatory: 0, MasterFieldId: 2},
		},
		"event": {
			{FieldName: "Event Title", FieldId: 1, IconPath: "/public/img/text.svg", OrderIndex: 1, Mandatory: 0, MasterFieldId: 2},
			{FieldName: "Event Date", FieldId: 2, IconPath: "/public/img/date.svg", OrderIndex: 2, Mandatory: 0, MasterFieldId: 6},
			{FieldName: "Event Time", FieldId: 3, IconPath: "/public/img/date-time.svg", OrderIndex: 3, Mandatory: 0, MasterFieldId: 5},
			{FieldName: "Event Type", FieldId: 4, IconPath: "/public/img/select.svg", OrderIndex: 4, Mandatory: 0, MasterFieldId: 2},
			{FieldName: "Event Description", FieldId: 5, IconPath: "/public/img/text-area.svg", OrderIndex: 5, Mandatory: 0, MasterFieldId: 8},
		},
		"jobs": {
			{FieldName: "Key Responsibilities", FieldId: 1, IconPath: "/public/img/text.svg", OrderIndex: 1, Mandatory: 0, MasterFieldId: 2},
			{FieldName: "Job Type", FieldId: 2, IconPath: "/public/img/select.svg", OrderIndex: 2, Mandatory: 0, MasterFieldId: 5},
			{FieldName: "Salary Range", FieldId: 3, IconPath: "/public/img/text.svg", OrderIndex: 3, Mandatory: 0, MasterFieldId: 2},
			{FieldName: "Experience", FieldId: 5, IconPath: "/public/img/text.svg", OrderIndex: 5, Mandatory: 0, MasterFieldId: 2},
		},
		"news": {
			{FieldName: "Frequency", FieldId: 1, IconPath: "/public/img/select.svg", OrderIndex: 1, Mandatory: 0, MasterFieldId: 5},
			{FieldName: "Anchor/Host Details", FieldId: 2, IconPath: "/public/img/text-area.svg", OrderIndex: 2, Mandatory: 0, MasterFieldId: 2},
			{FieldName: "Coverage Scope", FieldId: 3, IconPath: "/public/img/select.svg", OrderIndex: 3, Mandatory: 0, MasterFieldId: 5},
			{FieldName: "Special Features", FieldId: 4, IconPath: "/public/img/select.svg", OrderIndex: 4, Mandatory: 0, MasterFieldId: 5},
			{FieldName: "Multimedia Support", FieldId: 5, IconPath: "/public/img/select.svg", OrderIndex: 5, Mandatory: 0, MasterFieldId: 2},
		},
		"knowledge-base": {
			{FieldName: "Version", FieldId: 1, IconPath: "/public/img/text.svg", OrderIndex: 1, Mandatory: 0, MasterFieldId: 2},
			{FieldName: "Summary", FieldId: 1, IconPath: "/public/img/text.svg", OrderIndex: 1, Mandatory: 0, MasterFieldId: 2},
		},
		"1:1call": {
			{FieldName: "Duration In(mins)", FieldId: 1, IconPath: "/public/img/text.svg", OrderIndex: 1, Mandatory: 0, MasterFieldId: 17},
			{FieldName: "Short Description", FieldId: 5, IconPath: "/public/img/text-area.svg", OrderIndex: 1, Mandatory: 0, MasterFieldId: 8},
		},
	}

	// Define a map to store option values for select fields
	var fieldOptionsMap = map[string][]string{
		"Stock":            {"In Stock", "Out of Stock"},
		"Event Type":       {"Webinar", "Conference", "Workshop"},
		"Job Type":         {"Full-Time", "Part-Time", "Contract", "Internship"},
		"Frequency":        {"24/7", "Daily", "Weekly"},
		"Coverage Scope":   {"Local", "International"},
		"Special Features": {"Debates", "Expert Panels", "Investigative Reports"},
	}

	type field struct {
		Fiedlvalue []Fiedlvalue `json:"fiedlvalue"`
	}

	// channeldetails, _, err := ChannelConfig.GetChannelsById(id, -1)
	channeldetails := responseData.Channeldetail
	channellist, _ := ChannelConfig.GetchannelByName(channeldetails.ChannelName, TenantId)
	var CollectionCount int

	if channeldetails.ChannelName == channellist.ChannelName {

		CollectionCount = channellist.CollectionCount + 1
		CollectionCountstring := strconv.Itoa(CollectionCount)
		channeldetails.ChannelName = "Copy of " + channellist.ChannelName + " " + "(" + CollectionCountstring + ")"

		Channels := chn.Tblchannel{
			Id:              channellist.Id,
			CollectionCount: CollectionCount,
		}
		err := ChannelConfig.ChannelType(Channels)

		if err != nil {

			ErrorLog.Printf("channel Type error: %s", err)
		}
	} else {

		CollectionCount = 1

	}

	var FieldValuesss []chn.Fiedlvalue

	// Check if the channel slug exists in the map
	if fieldValues, ok := fieldValuesMap[channeldetails.SlugName]; ok {
		for _, val := range fieldValues {
			var fieldval chn.Fiedlvalue
			fieldval.CharacterAllowed = 0 // Assign default if needed
			fieldval.DateFormat = ""      // Assign default if needed
			fieldval.FieldId = val.FieldId
			fieldval.FieldName = val.FieldName
			fieldval.IconPath = val.IconPath
			fieldval.MasterFieldId = val.MasterFieldId // Assign default if needed
			fieldval.NewFieldId = 0                    // Assign default if needed
			fieldval.OrderIndex = val.OrderIndex
			fieldval.SectionId = 0    // Assign default if needed
			fieldval.SectionNewId = 0 // Assign default if needed
			fieldval.TimeFormat = ""  // Assign default if needed
			fieldval.Url = ""         // Assign default if needed
			fieldval.Mandatory = val.Mandatory

			// Assign option values if applicable
			if options, ok := fieldOptionsMap[fieldval.FieldName]; ok {

				fmt.Println("checkoptionssd")
				var choptions []chn.OptionValues
				for _, opt := range options {
					var optss chn.OptionValues
					optss.Value = opt
					choptions = append(choptions, optss)
				}
				fieldval.OptionValue = choptions
			}

			FieldValuesss = append(FieldValuesss, fieldval)
		}
	}

	categoryvalue, _ := getCategoryIds(channeldetails.SlugName, TenantId)

	if permisison {

		userid := c.GetInt("userid")

		uuid := (uuid.New()).String()

		arr := strings.Split(uuid, "-")

		UniqueId := arr[len(arr)-1]

		newchannel, _ := ChannelConfig.CreateChannel(chn.ChannelCreate{ChannelName: channeldetails.ChannelName, ChannelUniqueId: UniqueId, ChannelDescription: channeldetails.ChannelDescription, CategoryIds: categoryvalue, CreatedBy: userid, CollectionCount: CollectionCount, ImagePath: channeldetails.ImagePath}, moduleid, TenantId)

		ferr := ChannelConfig.CreateAdditionalFields(chn.ChannelAddtionalField{FieldValues: FieldValuesss, CreatedBy: userid}, newchannel.Id, TenantId)

		if ferr != nil {
			ErrorLog.Printf("channelcreate additional field error: %s", ferr)
		}

		if err != nil {
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
			c.JSON(200, false)
			return
		}

		c.SetCookie("get-toast", "Collection Added Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, "/channels")

	}
}

func CloneChannels(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	permisison, perr := NewAuth.IsGranted("Channels", auth.CRUD, TenantId)

	if perr != nil {
		ErrorLog.Printf("Channels authorization error :%s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Channels authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	Entry := "Entries"
	moduleid, err := models.Entryid(Entry, TenantId)
	if err != nil {
		fmt.Println(err)
	}

	_, Fieldvalue, ferr := ChannelConfig.GetChannelsFieldsById(id, TenantId)

	if ferr != nil {
		ErrorLog.Printf("getfieldata error: %s", perr)
	}

	channeldetails, _, _ := ChannelConfig.GetChannelsById(id, TenantId)
	fmt.Println("channeldetails:", channeldetails.CloneCount)
	CloneCount := channeldetails.CloneCount + 1

	Channels := chn.Tblchannel{
		Id:         channeldetails.Id,
		CloneCount: CloneCount,
	}
	cerr := ChannelConfig.ChannelType(Channels)

	if cerr != nil {

		ErrorLog.Printf("channel Type error: %s", err)
	}

	var Count string

	if channeldetails.CloneCount != 0 {

		Count = "(" + strconv.Itoa(channeldetails.CloneCount) + ")"

	}

	var FieldValuesss []chn.Fiedlvalue

	// Check if the channel slug exists in the map

	for _, val := range Fieldvalue {
		var fieldval chn.Fiedlvalue
		fieldval.CharacterAllowed = 0 // Assign default if needed
		fieldval.DateFormat = ""      // Assign default if needed
		fieldval.FieldId = val.FieldId
		fieldval.FieldName = val.FieldName
		fieldval.IconPath = val.IconPath
		fieldval.MasterFieldId = val.MasterFieldId // Assign default if needed
		fieldval.NewFieldId = 0                    // Assign default if needed
		fieldval.OrderIndex = val.OrderIndex
		fieldval.SectionId = 0    // Assign default if needed
		fieldval.SectionNewId = 0 // Assign default if needed
		fieldval.TimeFormat = ""  // Assign default if needed
		fieldval.Url = ""         // Assign default if needed
		fieldval.Mandatory = val.Mandatory

		// Assign option values if applicable

		fmt.Println("checkoptionssd")
		var choptions []chn.OptionValues
		for _, opt := range val.OptionValue {
			var optss chn.OptionValues
			optss.Value = opt.Value
			choptions = append(choptions, optss)
		}
		fieldval.OptionValue = choptions

		FieldValuesss = append(FieldValuesss, fieldval)
	}

	fmt.Println("FieldValuesss:", FieldValuesss)

	_, FinalSelectedCategories, cerr := ChannelConfig.GetChannelsById(id, TenantId)

	if cerr != nil {
		ErrorLog.Printf("getchannel error: %s", perr)
	}

	var categoryvalue []string

	for _, val := range FinalSelectedCategories {

		var str string

		for index, cat := range val.Categories {
			if index+1 != len(val.Categories) {
				str += strconv.Itoa(cat.Id) + ","
			} else {
				str += strconv.Itoa(cat.Id)
			}
		}

		categoryvalue = append(categoryvalue, str)
	}

	if permisison {

		userid := c.GetInt("userid")

		uuid := (uuid.New()).String()

		arr := strings.Split(uuid, "-")

		UniqueId := arr[len(arr)-1]

		newchannel, _ := ChannelConfig.CreateChannel(chn.ChannelCreate{ChannelName: "Clone of " + channeldetails.ChannelName + Count, ChannelUniqueId: UniqueId, ChannelDescription: channeldetails.ChannelDescription, CategoryIds: categoryvalue, CreatedBy: userid, ImagePath: channeldetails.ImagePath}, moduleid, TenantId)

		ferr := ChannelConfig.CreateAdditionalFields(chn.ChannelAddtionalField{FieldValues: FieldValuesss, CreatedBy: userid}, newchannel.Id, TenantId)

		if ferr != nil {
			ErrorLog.Printf("channelcreate additional field error: %s", ferr)
		}

		if err != nil {
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
			c.JSON(200, false)
			return
		}

		c.SetCookie("get-toast", "Collection Cloned Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, "/channels")

	}
}

func CheckChannelName(c *gin.Context) {

	channelid, _ := strconv.Atoi(c.PostForm("channel_id"))

	name := c.PostForm("channel_title")

	permisison, perr := NewAuth.IsGranted("Channels", auth.Read, TenantId)
	if perr != nil {
		ErrorLog.Printf("Channels check name authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Channels authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		flg, err := ChannelConfig.CheckNameInChannel(channelid, name, TenantId)

		if err != nil {
			ErrorLog.Printf("Channels checkname error: %s", err)
			json.NewEncoder(c.Writer).Encode(false)
			return
		}

		json.NewEncoder(c.Writer).Encode(flg)

	}

}

// Get category during add to collection function//
func getCategoryIds(channelSlug string, tenantId string) ([]string, error) {
	var categoryId int
	err := db.Table("tbl_categories").Where("tenant_id = ? AND category_slug = ?", tenantId, channelSlug).Pluck("id", &categoryId).Error
	if err != nil {
		return nil, err
	}

	if categoryId == 0 {
		return []string{}, nil
	}

	var categoryValue []string

	var childCategoryIds []int
	err = db.Table("tbl_categories").Where("tenant_id = ? AND parent_id = ?", tenantId, categoryId).Pluck("id", &childCategoryIds).Error
	if err != nil {
		return nil, err
	}

	for _, childId := range childCategoryIds {
		if childId != 0 {
			categoryValue = append(categoryValue, fmt.Sprintf("%d,%d", categoryId, childId))

			var subCategoryIds []int
			err = db.Table("tbl_categories").Where("tenant_id = ? AND parent_id = ?", tenantId, childId).Pluck("id", &subCategoryIds).Error
			if err != nil {
				log.Println(err)
			} else {
				for _, subCategoryId := range subCategoryIds {
					if subCategoryId != 0 {
						categoryValue = append(categoryValue, fmt.Sprintf("%d,%d,%d", categoryId, childId, subCategoryId))
					}
				}
			}
		}
	}

	return categoryValue, nil
}
