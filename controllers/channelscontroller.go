package controllers

import (
	"encoding/json"
	"fmt"
	"spurt-cms/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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

	ChannelConfig.DataAccess = c.GetInt("dataaccess")
	ChannelConfig.Userid = c.GetInt("userid")

	channelslist, channelcount, err := ChannelConfig.ListChannel(chn.Channels{Limit: Limit, Offset: offset, Keyword: keyword, IsActive: false, CreateOnly: true, TenantId: TenantId, Count: true})
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

	c.HTML(200, "channels.html", gin.H{"Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3,
	}, "Menu": menu, "csrf": csrf.GetToken(c), "channel": Finalchannelslist, "Searchtrue": filterflag, "count": channelcount, "HeadTitle": translate.Channell.Channels, "filter": keyword, "translate": translate, "Channelsmenu": true, "Cmsmenu": true, "title": ModuleName, "Tabmenu": TabName, "linktitle": ModuleName, "Paginationendcount": paginationendcount, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Paginationstartcount": paginationstartcount, "Limit": limt})

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

		c.HTML(200, "addchannel.html", gin.H{"Menu": menu, "linktitle": "Create Channel", "title": ModuleName, "csrf": csrf.GetToken(c), "Fields": field, "Button": "Save", "AllCategories": AllCategorieswithSubCategories, "Title": "Create Channel", "Back": "/settings/channels/channellist", "HeadTitle": "Create Channel", "translate": translate, "Channelsmenu": true, "Cmsmenu": true})

		return

	}

	c.Redirect(301, "/403-page")

}

func CreateChannel(c *gin.Context) {

	//get data from html
	channelname := c.Request.PostFormValue("channelname")
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
			fieldval.Mandatory=val.Mandatory

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

		newchannel, cerr := ChannelConfig.CreateChannel(chn.ChannelCreate{ChannelName: channelname, ChannelDescription: channeldesc, CategoryIds: categoryvalue, CreatedBy: userid}, moduleid, TenantId)

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
		if cerr := ChannelConfig.EditChannel(channelname, channeldesc, userid, channelid, categoryvalue, TenantId); cerr != nil {
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

		routename := "/channel/entrylist/" + ModuleName

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
