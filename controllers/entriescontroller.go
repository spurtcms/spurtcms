package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	chn "github.com/spurtcms/channels"
	mem "github.com/spurtcms/member"
	"github.com/spurtcms/team"
	csrf "github.com/utrack/gin-csrf"
)

type Section struct {
	SectionId     int    `json:"SectionId"`
	SectionNewId  int    `json:"SectionNewId"`
	SectionName   string `json:"SectionName"`
	MasterFieldId int    `json:"MasterFieldId"`
	OrderIndex    int    `json:"OrderIndex"`
}

type Fiedlvalue struct {
	MasterFieldId    int            `json:"MasterFieldId"`
	FieldId          int            `json:"FieldId"`
	NewFieldId       int            `json:"NewFieldId"`
	SectionId        int            `json:"SectionId"`
	SectionNewId     int            `json:"SectionNewId"`
	FieldName        string         `json:"FieldName"`
	DateFormat       string         `json:"DateFormat"`
	TimeFormat       string         `json:"TimeFormat"`
	OptionValue      []OptionValues `json:"OptionValue"`
	CharacterAllowed int            `json:"CharacterAllowed"`
	IconPath         string         `json:"IconPath"`
	Url              string         `json:"Url"`
	OrderIndex       int            `json:"OrderIndex"`
	Mandatory        int            `json:"Mandatory"`
}

type OptionValues struct {
	Id         int    `json:"Id"`
	NewId      int    `json:"NewId"`
	FieldId    int    `json:"FieldId"`
	NewFieldId int    `json:"NewFieldId"`
	Value      string `json:"Value"`
	OrderIndex int    `json:"OrderIndex"`
}

type EntriesFilter struct {
	Keyword     string
	Title       string
	ChannelName string
	Status      string
	UserName    string
	CategoryId  string
}

/*entries list*/
func Entries(c *gin.Context) {

	var (
		limt           int
		offset         int
		filters        EntriesFilter
		unpublishroute string
		draftroute     string
		publishroute   string
	)

	id, _ := strconv.Atoi(c.Param("id"))
	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	if pageno == 0 {
		pageno = 1
	}

	filters.Keyword = strings.TrimSpace(c.Query("keyword"))
	filters.Status = strings.TrimSpace(c.Query("Status"))
	filters.ChannelName = strings.TrimSpace(c.Query("ChannelName"))
	filters.Title = strings.TrimSpace(c.Query("Title"))

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	chnanem, _, cerr := ChannelConfig.GetChannelsById(id, TenantId)
	if cerr != nil {
		ErrorLog.Printf("Entries get channel data error :%s", cerr)
	}

	flag := true

	channelist, _, clerr := ChannelConfig.ListChannel(chn.Channels{Limit: 100, Offset: 0, IsActive: flag, TenantId: TenantId})
	if clerr != nil {
		ErrorLog.Printf("channellist error :%s", clerr)
	}

	routeName := c.FullPath()

	var htmlname string

	if strings.Contains(routeName, "entrylist") {

		filters.Status = "Published"

		htmlname = "entries.html"

	}
	if strings.Contains(routeName, "unpublishentrieslist") {

		filters.Status = "Unpublished"

		htmlname = "entryunpublish.html"
	}
	if strings.Contains(routeName, "draftentrieslist") {

		filters.Status = "Draft"

		htmlname = "entrydraft.html"
	}

	var chnallist []chn.Tblchannel

	for _, val := range channelist {

		_, _, entrcount, _ := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: val.Id, Limit: 0, Offset: 0, Keyword: "", ChannelName: "", Title: "", Status: ""}, TenantId)

		val.EntriesCount = entrcount

		chnallist = append(chnallist, val)

	}
	/*list*/
	query.DataAccess = c.GetInt("dataaccess")
	query.UserId = c.GetInt("userid")

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("Entries authorization error: %s", perr)
	}

	// if permisison {

	ChannelConfig.DataAccess = c.GetInt("dataaccess")
	ChannelConfig.Userid = c.GetInt("userid")

	chnentry, filtercount, _, _ := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: id, Limit: limt, Offset: offset, Keyword: filters.Keyword, ChannelName: filters.ChannelName, Title: filters.Title, Status: filters.Status}, TenantId)

	// _, filtercount, _, _ := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: id, Limit: 0, Offset: 0, Keyword: filters.Keyword, ChannelName: filters.ChannelName, Title: filters.Title, Status: filters.Status})

	_, entrcount, _, _ := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: id, Limit: 0, Offset: 0}, TenantId)

	var allchnentry []chn.Tblchannelentries

	for index, entry := range chnentry {
		entry.Cno = strconv.Itoa((offset + index) + 1)
		entry.CreatedDate = entry.CreatedOn.In(TZONE).Format(Datelayout)
		if !entry.ModifiedOn.IsZero() {
			entry.ModifiedDate = entry.ModifiedOn.In(TZONE).Format(Datelayout)
		} else {
			entry.ModifiedDate = entry.CreatedOn.In(TZONE).Format(Datelayout)
		}

		if entry.ProfileImagePath != "" {
			entry.ProfileImagePath = "/image-resize?name=" + entry.ProfileImagePath
		}

		// entry.ChannelName = ""
		allchnentry = append(allchnentry, entry)

	}

	var filterflag bool
	if filters.Keyword == "" && filters.Status == "" {
		filterflag = true
	} else {
		filterflag = false
	}

	paginationendcount := len(chnentry) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(pageno, int(filtercount), limt)

	menu := NewMenuController(c)
	ModuleName, TabName, _ := ModuleRouteName(c)
	translate, _ := TranslateHandler(c)
	selectedtype, _ := GetSelectedType()

	viewurl := os.Getenv("VIEW_BASE_URL")

	if filters.Keyword == "" {

		unpublishroute = "/channel/unpublishentrieslist/" + c.Param("id")

		draftroute = "/channel/draftentrieslist/" + c.Param("id")

		publishroute = "/channel/entrylist/" + c.Param("id")

	} else {

		unpublishroute = "/channel/unpublishentrieslist/" + c.Param("id") + "?keyword=" + filters.Keyword

		draftroute = "/channel/draftentrieslist/" + c.Param("id") + "?keyword=" + filters.Keyword

		publishroute = "/channel/entrylist/" + c.Param("id") + "?keyword=" + filters.Keyword
	}

	c.HTML(200, htmlname, gin.H{"Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3,
	}, "Viewbaseurl": viewurl, "Menu": menu, "linktitle": ModuleName, "chcount1": entrcount, "csrf": csrf.GetToken(c), "channelname": chnanem.ChannelName, "chnid": id, "ChanEntrtlist": allchnentry, "entrycount": entrcount, "Next": Next, "Previous": Previous, "PageCount": PageCount, "CurrentPage": pageno, "limit": limt, "paginationendcount": paginationendcount, "paginationstartcount": paginationstartcount, "filter": filters, "title": ModuleName, "heading": chnanem.ChannelName, "translate": translate, "channellist": chnallist, "Page": Page, "HeadTitle": translate.Channell.Channels, "chentrycount": filtercount, "Cmsmenu": true, "filterflag": filterflag, "Entriestab": true, "Tabmenu": TabName, "StorageType": selectedtype.SelectedType, "unpublishroute": unpublishroute, "draftroute": draftroute, "publishroute": publishroute, "channelfilter": "true"})

	// }

	// c.Redirect(301, "/403-page")

}

/*create entry*/
func CreateEntry(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	// flag := true

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("create Entry authorization error: %s", perr)
	}

	// if permisison {

	channelist, clerr := ChannelConfig.GetPermissionChannel(TenantId)
	if clerr != nil {
		ErrorLog.Printf("create Entry listchannel error: %s", clerr)
	}

	_, entrcount, _, cherr := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: id, Limit: 0, Offset: 0}, TenantId)
	if cherr != nil {
		ErrorLog.Printf("createEntry entrylist error: %s", cherr)
	}

	_, FinalSelectedCategories, cerr := ChannelConfig.GetChannelsById(id, TenantId)
	if cerr != nil {
		ErrorLog.Printf("createEntry channel data error: %s", cerr)
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

	field, err := ChannelConfig.GetAllMasterFieldType(TenantId)
	if err != nil {
		ErrorLog.Printf("createEntry get master fields error: %s", perr)
	}

	parsedTime, err := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	cdate := parsedTime.In(TZONE).Format("2006-01-02T15:04")
	translate, _ := TranslateHandler(c)
	menu := NewMenuController(c)

	selectedtype, _ := GetSelectedType()

	urlpath := map[string]interface{}{"path": os.Getenv("BASE_URL") + "uploadb64image", "payload": "imagedata", "success": map[string]interface{}{"imagepath": "imagepath", "imagename": "imagename"}}

	ubyte, _ := json.Marshal(urlpath)

	ModuleName, _, _ := ModuleRouteName(c)

	selectedchannel := channelist[len(channelist)-1]

	slchannelid := selectedchannel.Id

	c.HTML(200, "addentry.html", gin.H{"Menu": menu, "linktitle": "Create Entry", "title": ModuleName, "Fields": field, "translate": translate, "csrf": csrf.GetToken(c), "channellist": channelist, "CategoryName": FinalSelectedCategories, "entrycount": entrcount, "HeadTitle": translate.Channell.Channels, "Cmsmenu": true, "Entriestab": true, "currentdate": cdate, "StorageType": selectedtype.SelectedType, "Mode": "create", "Slchannelid": slchannelid, "Storagepath": string(ubyte)})

	return

	// }

	// c.Redirect(301, "/403-page")

}

func AllEntries(c *gin.Context) {

	var (
		limt   int
		offset int
	)

	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	var filters EntriesFilter
	filters.Keyword = strings.TrimSpace(c.Query("keyword"))
	filters.Status = strings.TrimSpace(c.Query("Status"))
	filters.ChannelName = strings.TrimSpace(c.Query("ChannelName"))
	filters.Title = strings.TrimSpace(c.Query("Title"))

	var filterflag bool
	if filters.Keyword == "" {
		filterflag = true
	} else {
		filterflag = false
	}

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("all entries authorization error: %s", perr)
	}

	// if permisison {

	ChannelConfig.DataAccess = c.GetInt("dataaccess")
	ChannelConfig.Userid = c.GetInt("userid")

	routeName := c.FullPath()

	var entrystatus string

	var htmlname string

	if strings.Contains(routeName, "entrylist") {

		entrystatus = "Published"

		htmlname = "entries.html"

	}
	if strings.Contains(routeName, "unpublishentries") {

		entrystatus = "Unpublished"

		htmlname = "entryunpublish.html"
	}
	if strings.Contains(routeName, "draftentries") {

		entrystatus = "Draft"

		htmlname = "entrydraft.html"
	}

	_, Totalentris, _, _ := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: 0, Limit: 0, Offset: 0, Keyword: filters.Keyword, ChannelName: filters.ChannelName, Title: filters.Title, Status: entrystatus}, TenantId)

	_, Totalentris1, _, _ := ChannelConfig.ChannelEntriesList(chn.Entries{}, TenantId)

	chlist, _, _, _ := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: 0, Limit: limt, Offset: offset, Keyword: filters.Keyword, ChannelName: filters.ChannelName, Title: filters.Title, Status: entrystatus}, TenantId)

	var chlist1 []chn.Tblchannelentries

	var newentrylist []chn.Tblchannelentries

	var childrenlist1 []chn.Tblchannelentries

	for index, entry := range chlist {

		newentrylist, _ = ChannelConfig.EntrylistByParentId(entry.Id, TenantId)

		// fmt.Println(newentrylist, "newentrylist")

		fmt.Println(entry.ParentId, "channelentrilist")
		entry.Cno = strconv.Itoa((offset + index) + 1)
		entry.CreatedDate = entry.CreatedOn.In(TZONE).Format(Datelayout)
		if !entry.ModifiedOn.IsZero() {
			entry.ModifiedDate = entry.ModifiedOn.In(TZONE).Format(Datelayout)
		} else {
			entry.ModifiedDate = entry.CreatedOn.In(TZONE).Format(Datelayout)
		}

		if entry.ProfileImagePath != "" {
			entry.ProfileImagePath = "/image-resize?name=" + entry.ProfileImagePath
		}

		for _, val := range newentrylist {

			val.CreatedDate = val.CreatedOn.In(TZONE).Format(Datelayout)
			if !val.ModifiedOn.IsZero() {
				val.ModifiedDate = val.ModifiedOn.In(TZONE).Format(Datelayout)
			} else {
				val.ModifiedDate = val.CreatedOn.In(TZONE).Format(Datelayout)
			}

			if val.ProfileImagePath != "" {
				val.ProfileImagePath = "/image-resize?name=" + val.ProfileImagePath
			}

			childrenlist1 = append(childrenlist1, val)
		}
		entry.ChildrenList = childrenlist1

		chlist1 = append(chlist1, entry)

	}

	channelist, clerr := ChannelConfig.GetPermissionChannel(TenantId)
	if clerr != nil {
		ErrorLog.Printf(" Entry listchannel error: %s", clerr)
	}
	//membergroup list
	_, memgrpcount, err := MemberConfig.ListMemberGroup(mem.MemberGroupListReq{ActiveGroupsOnly: true}, TenantId) // get active member group count
	if err != nil {
		ErrorLog.Printf("channel count error: %s", err)
	}
	membergroup, _, err := MemberConfig.ListMemberGroup(mem.MemberGroupListReq{Limit: int(memgrpcount), ActiveGroupsOnly: true}, TenantId) // get active member group
	if err != nil {
		ErrorLog.Printf("channel count error: %s", err)
	}

	//pagination calc
	paginationendcount := len(chlist) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(pageno, int(Totalentris), limt)

	menu := NewMenuController(c)
	translate, _ := TranslateHandler(c)
	ModuleName, tabname, _ := ModuleRouteName(c)
	selectedtype, _ := GetSelectedType()

	viewurl := os.Getenv("VIEW_BASE_URL")

	c.HTML(200, htmlname, gin.H{"Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3,
	}, "Viewbaseurl": viewurl, "Menu": menu, "translate": translate, "Page": Page, "chentrycount": Totalentris, "linktitle": ModuleName, "entrycount": Totalentris1, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "limit": limt, "paginationendcount": paginationendcount, "paginationstartcount": paginationstartcount, "filter": filters, "ChanEntrtlist": chlist1, "csrf": csrf.GetToken(c), "channellist": channelist, "title": ModuleName, "HeadTitle": translate.Channell.Channels, "Cmsmenu": true, "filterflag": filterflag, "Entriestab": true, "Tabmenu": tabname, "StorageType": selectedtype.SelectedType, "chnid": -1, "Membergroup": membergroup})

	// }

	// c.Redirect(301, "/403-page")

}

func DeleteEntries(c *gin.Context) {

	entryId, _ := strconv.Atoi(c.Query("id"))
	channame := c.Query("cname")
	pageno := c.Query("page")
	pagename := c.Query("pname")

	_, _, err := ChannelConfig.FetchChannelEntryDetail(chn.EntriesInputs{TenantId: TenantId}, []int{})
	if err != nil {
		ErrorLog.Printf("fetch entry detail error: %s", err)
		return
	}

	var url string

	var route string

	if pagename == "publish" {

		route = "/channel/entrylist/"

	}
	if pagename == "unpublish" {

		route = "/channel/unpublishentries/"

	}
	if pagename == "draft" {

		route = "/channel/draftentries/"

	}

	if pageno != "" {
		url = route + "?page=" + pageno
	} else {
		url = route
	}

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("delete entries authorization error: %s", perr)
	}

	// if permisison {

	userid := c.GetInt("userid")

	_, err = ChannelConfig.DeleteEntry(channame, userid, entryId, TenantId)
	if err != nil {
		ErrorLog.Printf("delete entries error: %s", perr)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.Redirect(301, url)
		return
	}

	c.SetCookie("get-toast", "Entry Deleted Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)

	c.Redirect(301, url)

	// }

	// c.Redirect(301, "/403-page")

}

func EntryStatus(c *gin.Context) {

	channelname := c.Request.PostFormValue("cname")
	id, _ := strconv.Atoi(c.Request.PostFormValue("entryid"))
	status, _ := strconv.Atoi(c.Request.PostFormValue("status"))
	userid := c.GetInt("userid")

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("entry status authorization error: %s", perr)
	}
	// if permisison {

	_, err := ChannelConfig.EntryStatus(channelname, id, status, userid, TenantId)
	if err != nil {
		ErrorLog.Printf("entry status change error: %s", perr)
		json.NewEncoder(c.Writer).Encode(false)
		return
	}

	json.NewEncoder(c.Writer).Encode(true)

	// }

	// c.Redirect(301, "/403-page")

}

func PublishEntry(c *gin.Context) {

	//get data from html form data
	eid, _ := strconv.Atoi(c.Param("id"))

	cid, _ := strconv.Atoi(c.Request.PostFormValue("id"))
	cname := c.Request.PostFormValue("cname")
	title := c.Request.PostFormValue("title")
	text := c.Request.PostFormValue("text")
	img := c.Request.PostFormValue("image")
	authorname := c.Request.PostFormValue("author")
	createdate := c.Request.PostFormValue("createtime")
	publishtime := c.Request.PostFormValue("publishtime")
	readingtime, _ := strconv.Atoi(c.Request.PostFormValue("readingtime"))
	sortorder, _ := strconv.Atoi(c.Request.PostFormValue("sortorder"))
	tagname := c.Request.PostFormValue("tagname")
	extxt := c.Request.PostFormValue("extxt")
	orderindex, _ := strconv.Atoi(c.Request.PostFormValue("orderindex"))
	status, _ := strconv.Atoi(c.Request.PostFormValue("status"))
	categoryids := c.PostFormArray("categoryids[]")
	userid := c.GetInt("userid")

	layout := "2006-01-02T15:04"

	// Parse string to time.Time
	ctTime, _ := time.Parse(layout, createdate)
	pbTime, _ := time.Parse(layout, publishtime)

	fmt.Println("ctpt", ctTime, pbTime)

	categoryidsStr := strings.Join(categoryids, ",")

	var seodetails map[string]string

	if err := json.Unmarshal([]byte(c.Request.PostFormValue("seodetails")), &seodetails); err != nil {
		ErrorLog.Printf("publishentry unmarshall seodetails error: %s", err)
	}

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("publish entry authorization error: %s", perr)
	}

	// if permisison {

	entries := chn.EntriesRequired{

		Title:       title,
		Content:     text,
		ChannelId:   cid,
		CoverImage:  img,
		CategoryIds: categoryidsStr,
		Status:      status,
		Author:      authorname,
		CreateTime:  ctTime,
		PublishTime: pbTime,
		ReadingTime: readingtime,
		SortOrder:   sortorder,
		Tag:         tagname,
		Excerpt:     extxt,
		OrderIndex:  orderindex,

		SEODetails: chn.SEODetails{

			MetaTitle:       seodetails["title"],
			MetaDescription: seodetails["desc"],
			MetaKeywords:    seodetails["keyword"],
			MetaSlug:        seodetails["slug"],
			ImageAltTag:     seodetails["imgtag"],
		},
	}

	var (
		AdditionalFields []chn.AdditionalFields
		channeldata      []map[string]string
	)

	if err := json.Unmarshal([]byte(c.Request.PostFormValue("channeldata")), &channeldata); err != nil {
		ErrorLog.Printf("publish entry unmarshll channeldata error: %s", perr)
	}

	for _, channel := range channeldata {

		Fid, err := strconv.Atoi(channel["fieldid"])
		if err != nil {
			ErrorLog.Printf("publishentry fieldid convert error: %s", err)
		}

		FieldidInt, err := strconv.Atoi(channel["fid"])
		if err != nil {
			ErrorLog.Printf("publishentry fieldid convert error: %s", err)
		}

		if channel["value"] == "" && Fid != 0 {

			additionalField := chn.AdditionalFields{
				FieldName:  channel["name"],
				FieldValue: channel["value"],
				FieldId:    FieldidInt,
				Id:         Fid,
			}
			AdditionalFields = append(AdditionalFields, additionalField)
		}

		if channel["value"] != "" {

			additionalField := chn.AdditionalFields{
				FieldName:  channel["name"],
				FieldValue: channel["value"],
				FieldId:    FieldidInt,
				Id:         Fid,
			}
			AdditionalFields = append(AdditionalFields, additionalField)
		}

	}

	if eid != 0 {

		entries.ModifiedBy = userid
		_, err := ChannelConfig.UpdateEntry(entries, cname, eid, TenantId)
		ChannelConfig.UpdateAdditionalField(AdditionalFields, eid, TenantId)

		if err != nil {
			ErrorLog.Printf("publishentry error: %s", err)
			json.NewEncoder(c.Writer).Encode(0)
			return
		}

		if status == 1 {

			c.SetCookie("get-toast", "Entry Published Successfully", 3600, "", "", false, false)
			c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		}

		c.JSON(200, gin.H{"Channelname": cname, "id": eid})

	} else {
		if status == 1 {
			entries.IsActive = 1
		}
		Totalentries, totalentriescount, _, _ := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: 0, Limit: 0, Offset: 0}, TenantId)
		fmt.Println("Totalentries", totalentriescount)
		for index, value := range Totalentries {
			value.OrderIndex = 2 + index

			_, err := ChannelConfig.UpdateEntryOrderIndex(value.OrderIndex, value.Id, userid, TenantId)
			fmt.Println("err", err)

		}
		entries.CreatedBy = userid
		chenid, _, err1 := ChannelConfig.CreateEntry(entries, TenantId)
		ChannelConfig.CreateChannelEntryFields(chenid.Id, userid, AdditionalFields, TenantId)
		channeldet, _, _ := ChannelConfig.GetChannelsById(chenid.ChannelId, TenantId)

		if err1 != nil {
			ErrorLog.Printf("publishentry error: %s", err1)
			json.NewEncoder(c.Writer).Encode(0)
			return
		}

		if status == 1 {
			c.SetCookie("get-toast", "Entry Published Successfully", 3600, "", "", false, false)
			c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		}

		c.JSON(200, gin.H{"Channelname": channeldet.ChannelName, "id": chenid.Id})

	}

	// }

	// c.Redirect(301, "/403-page")

}

func EditEntry(c *gin.Context) {

	flag := true
	id, _ := strconv.Atoi(c.Param("id"))
	entries, _ := ChannelConfig.EntryDetailsById(chn.IndivEntriesReq{EntryId: id}, TenantId)
	_, perr := NewAuth.IsGranted("Entries", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("edit entry page authorization error: %s", perr)
	}

	// if permisison {

	channelist, _, cherr := ChannelConfig.ListChannel(chn.Channels{Limit: 100, Offset: 0, IsActive: flag, TenantId: TenantId})
	if cherr != nil {
		ErrorLog.Printf("edit entry get channel list error: %s", cherr)
	}

	field, err := ChannelConfig.GetAllMasterFieldType(TenantId)
	if err != nil {
		ErrorLog.Printf("edit entry get master fields error: %s", err)
	}

	AllCategorieswithSubCategories, _ := CategoryConfig.AllCategoriesWithSubList(TenantId)

	menu := NewMenuController(c)

	translate, _ := TranslateHandler(c)
	selectedtype, _ := GetSelectedType()

	urlpath := map[string]interface{}{"path": os.Getenv("BASE_URL") + "uploadb64image", "payload": "imagedata", "success": map[string]interface{}{"imagepath": "imagepath", "imagename": "imagename"}}

	ubyte, _ := json.Marshal(urlpath)

	ModuleName, _, _ := ModuleRouteName(c)

	var slchannelid int

	if entries.ChannelId != 0 {

		slchannelid = entries.ChannelId

	} else {
		selectedchannel := channelist[len(channelist)-1]

		slchannelid = selectedchannel.Id
	}
	viewurl := os.Getenv("VIEW_BASE_URL")

	c.HTML(200, "addentry.html", gin.H{"Menu": menu, "Viewbaseurl": viewurl, "linktitle": "Edit Entry", "title": ModuleName, "Entries": entries, "Fields": field, "translate": translate, "csrf": csrf.GetToken(c), "channellist": channelist, "AllCategories": AllCategorieswithSubCategories, "HeadTitle": translate.Channell.Channels, "Cmsmenu": true, "Entriestab": true, "StorageType": selectedtype.SelectedType, "editbread": "Edit Entry", "Mode": "edit", "Slchannelid": slchannelid, "Storagepath": string(ubyte)})

	return

	// }

	// c.Redirect(301, "/403-page")

}

func EditDetails(c *gin.Context) {

	var Createtime, Publishedtime string

	id, _ := strconv.Atoi(c.Param("id"))
	cname := c.Param("channelname")

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("edit entry authorization error: %s", perr)
	}

	// if permisison {

	entries, err := ChannelConfig.EntryDetailsById(chn.IndivEntriesReq{ChannelName: cname, EntryId: id}, TenantId)

	layout := "2006-01-02T15:04"

	if !entries.CreateTime.IsZero() {
		Createtime = entries.CreateTime.Format(layout)
	}

	if !entries.PublishedTime.IsZero() {
		Publishedtime = entries.PublishedTime.Format(layout)
	}

	if err != nil {
		ErrorLog.Printf("edit entry authorization error: %s", perr)
	}

	channelid := entries.ChannelId
	membername := entries.TblChannelEntryField

	var memid int

	for _, val := range membername {
		if val.FieldTypeId == 14 {
			memid, _ = strconv.Atoi(val.FieldValue)
		}
	}

	memberfirstname, err := MemberConfigWP.GetMemberDetails(memid, TenantId)
	if err != nil {
		ErrorLog.Printf("edit entry get member details error: %s", err)
	}

	_, FinalSelectedCategories, cerr := ChannelConfig.GetChannelsById(channelid, TenantId)

	if cerr != nil {
		ErrorLog.Printf("edit entry get channeldetails error: %s", cerr)
	}

	Section, Fieldvalue, ferr := ChannelConfig.GetChannelsFieldsById(channelid, TenantId)
	if ferr != nil {
		ErrorLog.Printf("edit entry channel field data error: %s", ferr)
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

	json.NewEncoder(c.Writer).Encode(gin.H{"Entries": entries, "Createtime": Createtime, "Publishedtime": Publishedtime, "Section": Section, "FieldValue": Fieldvalue, "SelectedCategory": idstr, "CategoryName": FinalSelectedCategories, "Memberlist": memberfirstname.FirstName + " " + memberfirstname.LastName})

	return

	// }
	// c.Redirect(301, "/403-page")

}

func ChannelFields(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("channelfield  error: %s", perr)
	}

	// if permisison {

	_, FinalSelectedCategories, cerr := ChannelConfig.GetChannelsById(id, TenantId)
	if cerr != nil {
		ErrorLog.Printf("channel  error: %s", cerr)
	}

	Section, Fieldvalue, ferr := ChannelConfig.GetChannelsFieldsById(id, TenantId)
	if ferr != nil {
		ErrorLog.Printf("channel  error: %s", ferr)
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

	json.NewEncoder(c.Writer).Encode(gin.H{"Section": Section, "FieldValue": Fieldvalue, "SelectedCategory": idstr, "CategoryName": FinalSelectedCategories, "Categorycount": len(FinalSelectedCategories)})

	return

	// }

	// c.Redirect(301, "/403-page")

}

/*upload image*/
func ImageUpload(c *gin.Context) {

	file, _, err := c.Request.FormFile("file")

	if err != nil {

		fmt.Println("Error Retrieving the File")

		fmt.Println(err)

		return
	}

	defer file.Close()

	os.MkdirAll("storage/entry", 0755)

	tempFile, err := ioutil.TempFile("storage/entry", "IMG-*.png")

	Filename := tempFile.Name()
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	_, err1 := tempFile.Write(fileBytes)

	if err1 != nil {

		json.NewEncoder(c.Writer).Encode(gin.H{"status": false})

	}

	var url string

	if os.Getenv("BASE_URL") == "" {

		url = os.Getenv("DOMAIN_URL")

	} else {

		url = os.Getenv("BASE_URL")

	}

	data := map[string]interface{}{
		"success": true,
		"url":     url + Filename,
	}

	json.NewEncoder(c.Writer).Encode(data)

}

func Updatechannelfields(c *gin.Context) {

	channelid, _ := strconv.Atoi(c.Request.PostFormValue("id"))

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("update channelfield authorization  error: %s", perr)
	}

	// if permisison {

	channeldetails, _, cerr := ChannelConfig.GetChannelsById(channelid, TenantId)
	if cerr != nil {
		ErrorLog.Printf("update channelfield error: %s", cerr)
	}

	channelname := channeldetails.ChannelName
	channeldesc := channeldetails.ChannelDescription
	sectionvalue := c.Request.FormValue("sections")
	deletesection := c.Request.FormValue("deletesections")
	deleteoption := c.Request.PostFormValue("deleteoptions")
	fval := c.Request.PostFormValue("fiedlvalue")
	deletefval := c.Request.PostFormValue("deletefields")
	categoryvalue := c.PostFormArray("categoryvalue[]")

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
	type deleteopt struct {
		Fiedlvalue []OptionValues `json:"deleteoption"`
	}

	var (
		Sections            sections
		fieldval            field
		deleteSections      delsections
		deleteFields        delfield
		DeleteOpt           deleteopt
		Sectionss           []chn.Section
		FieldValuess        []chn.Fiedlvalue
		DeleteSectionss     []chn.Section
		DeleteFieldValuess  []chn.Fiedlvalue
		DeleteOptionsvalues []chn.OptionValues
	)

	if secerr := json.Unmarshal([]byte(sectionvalue), &Sections); secerr != nil {
		ErrorLog.Printf("update channelfield unmarshall section error: %s", secerr)
	}

	if fielerr := json.Unmarshal([]byte(fval), &fieldval); fielerr != nil {
		ErrorLog.Printf("update channelfield unmarshall field error: %s", fielerr)
	}

	if delesecerr := json.Unmarshal([]byte(deletesection), &deleteSections); delesecerr != nil {
		ErrorLog.Printf("update channelfield unmarshall deletesection error: %s", delesecerr)
	}

	if delfieerr := json.Unmarshal([]byte(deletefval), &deleteFields); delfieerr != nil {
		ErrorLog.Printf("update channelfield unmarshall deletefield error: %s", delfieerr)
	}

	if delopterr := json.Unmarshal([]byte(deleteoption), &DeleteOpt); delopterr != nil {
		ErrorLog.Printf("update channelfield unmarshall delteoption error: %s", delopterr)
	}

	for _, val := range Sections.Fiedlvalue {

		var sec chn.Section
		sec.MasterFieldId = val.MasterFieldId
		sec.OrderIndex = val.OrderIndex
		sec.SectionId = val.SectionId
		sec.SectionName = val.SectionName
		sec.SectionNewId = val.SectionNewId
		Sectionss = append(Sectionss, sec)

	}

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

	for _, val := range deleteSections.Fiedlvalue {

		var sec chn.Section
		sec.MasterFieldId = val.MasterFieldId
		sec.OrderIndex = val.OrderIndex
		sec.SectionId = val.SectionId
		sec.SectionName = val.SectionName
		sec.SectionNewId = val.SectionNewId
		DeleteSectionss = append(DeleteSectionss, sec)

	}

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
			optss.OrderIndex = opt.OrderIndex
			choptions = append(choptions, optss)
		}

		fieldval.OptionValue = choptions
		DeleteFieldValuess = append(DeleteFieldValuess, fieldval)

	}

	for _, val := range DeleteOpt.Fiedlvalue {

		var DeleteOptionsvalue chn.OptionValues
		DeleteOptionsvalue.FieldId = val.FieldId
		DeleteOptionsvalue.Id = val.Id
		DeleteOptionsvalue.NewFieldId = val.NewFieldId
		DeleteOptionsvalue.NewId = val.NewId
		DeleteOptionsvalue.Value = val.Value
		DeleteOptionsvalues = append(DeleteOptionsvalues, DeleteOptionsvalue)
	}
	fmt.Println("efcfe", FieldValuess)
	userid := c.GetInt("userid")

	if ederr := ChannelConfig.EditChannel(channelname, channeldesc, userid, channelid, categoryvalue, TenantId); ederr != nil {
		ErrorLog.Printf("update channel error: %s", ederr)
	}

	if updatechn := ChannelConfig.UpdateChannelField(chn.ChannelUpdate{Sections: Sectionss, FieldValues: FieldValuess, Deletesections: DeleteSectionss, DeleteFields: DeleteFieldValuess, DeleteOptionsvalue: DeleteOptionsvalues, CategoryIds: categoryvalue, ModifiedBy: userid}, channelid, TenantId); updatechn != nil {
		ErrorLog.Printf("update channelfield error: %s", updatechn)
	}

	json.NewEncoder(c.Writer).Encode(true)

	return

	// }

	// c.Redirect(301, "/403-page")

}

// make feature
func MakeFeature(c *gin.Context) {

	channelname := c.PostForm("cname")
	entryid, _ := strconv.Atoi(c.PostForm("entryid"))
	feature, _ := strconv.Atoi(c.PostForm("feature"))

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("make feature error: %s", perr)
	}

	// if permisison {

	chn, _ := ChannelConfig.GetchannelByName(channelname, TenantId)

	_, err := ChannelConfig.MakeFeature(chn.Id, entryid, feature, TenantId)
	if err != nil {
		ErrorLog.Printf("make feature error: %s", err)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
		return
	}

	c.SetCookie("get-toast", "Entry Updated Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	return

	// }

	// c.Redirect(301, "/403-page")

}

func MemberDetails(c *gin.Context) {

	var filter mem.Filter

	filter.FirstName = c.Query("keyword")

	flag1 := false

	if filter.FirstName != "" {

		memberlist, _, err := MemberConfigWP.ListMembers(0, 0, filter, flag1, TenantId)

		if err != nil {

			fmt.Println(err)
		}
		c.JSON(200, memberlist)
	} else {

		c.JSON(200, "")
	}

}

func CheckEnriesOrder(c *gin.Context) {

	chid, _ := strconv.Atoi(c.PostForm("chid"))
	eid, _ := strconv.Atoi(c.PostForm("eid"))
	order, _ := strconv.Atoi(c.PostForm("order"))

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("check entreis order error: %s", perr)
	}

	// if permisison {

	chnentry, _, _, err := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: chid, Limit: 0, Offset: 0}, TenantId)
	if err != nil {
		ErrorLog.Printf("check entreis order ,entrieslist error: %s", err)
	}

	for _, val := range chnentry {
		if (order == val.SortOrder) && (eid != val.Id) {
			c.JSON(200, gin.H{"value": true})
			return
		}
	}

	c.JSON(200, gin.H{"value": false})
	return

	// }

	// c.Redirect(301, "/403-page")

}

func UserDetails(c *gin.Context) {

	fname := c.Query("keyword")
	NewTeam.Dataaccess = c.GetInt("dataaccess")
	NewTeam.Userid = c.GetInt("userid")
	permisison, perr := NewAuth.IsGranted("Team", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("user detaisl authorization error: %s", perr)
	}

	if permisison {

		if fname != "" {

			userlist, _, err := NewTeam.ListUser(10, 0, team.Filters{FirstName: fname}, TenantId)
			if err != nil {
				ErrorLog.Printf("get user details error: %s", perr)
			}

			c.JSON(200, userlist)
			return
		}

		c.JSON(200, "")
		return

	}

	c.Redirect(301, "/403-page")

}

// MULTI SELECT ENTRY DELETE FUNCTION//
func DeleteSelectedEntry(c *gin.Context) {

	var (
		entrydata []map[string]string
		entryids  []int
		url       string
	)

	userid := c.GetInt("userid")
	pageno := c.PostForm("page")
	pname := c.PostForm("url")

	if err := json.Unmarshal([]byte(c.Request.PostFormValue("entryids")), &entrydata); err != nil {
		ErrorLog.Printf("delete multiple entry error: %s", err)
	}

	for _, val := range entrydata {
		entryid, _ := strconv.Atoi(val["entryid"])
		entryids = append(entryids, entryid)
	}

	_, _, err := ChannelConfig.FetchChannelEntryDetail(chn.EntriesInputs{TenantId: TenantId}, entryids)

	if err != nil {
		ErrorLog.Printf("fetching multiple entries data: %v", err)
		return
	}

	if pageno != "" {
		url = pname + "?page=" + pageno
	} else {
		url = pname
	}

	permisison, perr := NewAuth.IsGranted("Entries", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("delete multiplle entries authorization error: %s", perr)
	}
	if !permisison {
		ErrorLog.Printf("Entries authorization error")
		c.JSON(200, gin.H{"value": false})
		return
	}
	// if permisison {

	_, err = ChannelConfig.DeleteSelectedEntry(entryids, userid, TenantId)
	if err != nil {
		ErrorLog.Printf("delete multiplle entries  error: %s", err)
		c.JSON(200, gin.H{"value": false})
		return
	}

	c.JSON(200, gin.H{"value": true, "url": url})

	// }

	// c.Redirect(301, "/403-page")

}

// Multi selected entry Unpublish //
func UnpublishSelectedEntry(c *gin.Context) {

	var (
		entrydata []map[string]string
		entryids  []int
		statusint int
		status    string
		url       string
	)

	pageno := c.PostForm("page")
	userid := c.GetInt("userid")
	geturl := c.PostForm("url")
	if err := json.Unmarshal([]byte(c.Request.PostFormValue("entryids")), &entrydata); err != nil {
		fmt.Println(err)
	}

	for _, val := range entrydata {

		entryid, _ := strconv.Atoi(val["entryid"])
		status = val["status"]

		if status == "Draft" {
			statusint = 1
		} else if status == "Published" {
			statusint = 2
		} else if status == "Unpublished" {
			statusint = 1
		}

		entryids = append(entryids, entryid)
	}

	_, _, err := ChannelConfig.FetchChannelEntryDetail(chn.EntriesInputs{TenantId: TenantId}, entryids)

	if err != nil {
		ErrorLog.Printf("fetching multiple entries data: %v", err)
		return
	}

	if pageno != "" {
		url = geturl + "?page=" + pageno
	} else {
		url = geturl
	}

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("unpublished multiplle entries authorization error: %s", perr)
	}

	// if permisison {

	_, err = ChannelConfig.UnpublishSelectedEntry(entryids, statusint, userid, TenantId)
	if err != nil {
		ErrorLog.Printf("unpublished multiplle entries error: %s", err)
		c.JSON(200, gin.H{"value": false})
		return
	}

	c.JSON(200, gin.H{"value": true, "status": statusint, "url": url})

	// }

	// c.Redirect(301, "/403-page")

}

func CheckMandatoryFields(c *gin.Context) {

	entryid, _ := strconv.Atoi(c.Param("id"))

	entrydetails, err := ChannelConfig.EntryDetailsById(chn.IndivEntriesReq{EntryId: entryid}, TenantId)

	if err != nil {

		ErrorLog.Printf("Getting Entry details error: %s", err)
		c.JSON(200, err)
		return
	}

	c.JSON(200, entrydetails)
}

/*permission Member status*/
func EntryIsActive(c *gin.Context) {

	userid := c.GetInt("userid")

	id, _ := strconv.Atoi(c.Request.PostFormValue("id"))
	val, _ := strconv.Atoi(c.Request.PostFormValue("isactive"))
	fmt.Println("datas", userid, id, val)

	permisison, perr := NewAuth.IsGranted("Entries", auth.Update, TenantId)
	if perr != nil {
		ErrorLog.Printf("Entry status authorization error: %s", perr)
	}
	if !permisison {
		ErrorLog.Printf("Member authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		flg, err := ChannelConfig.EntryIsActive(id, val, userid, TenantId)

		if err != nil {
			ErrorLog.Printf("entry status change error: %s", perr)
			json.NewEncoder(c.Writer).Encode(flg)

		} else {
			json.NewEncoder(c.Writer).Encode(flg)
		}
	}

}

func EntryParentIdUpdate(c *gin.Context) {

	pid, _ := strconv.Atoi(c.PostForm("parententryid"))

	entryid, _ := strconv.Atoi(c.PostForm("entryid"))

	userid := c.GetInt("userid")

	fmt.Println(pid, entryid, "parentid")

	permisison, perr := NewAuth.IsGranted("Entries", auth.Update, TenantId)
	if perr != nil {
		ErrorLog.Printf("Entry parent update authorization error: %s", perr)
	}
	if !permisison {
		ErrorLog.Printf("Entries authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		if pid == -1 {

			pid = 0
		}

		err := ChannelConfig.EntryParentIdUpdate(pid, entryid, TenantId, userid)

		entry, _ := ChannelConfig.EntryDetailsById(chn.IndivEntriesReq{EntryId: entryid}, TenantId)
		entry.CreatedDate = entry.CreatedOn.In(TZONE).Format(Datelayout)
		if !entry.ModifiedOn.IsZero() {
			entry.ModifiedDate = entry.ModifiedOn.In(TZONE).Format(Datelayout)
		} else {
			entry.ModifiedDate = entry.CreatedOn.In(TZONE).Format(Datelayout)
		}

		if entry.ProfileImagePath != "" {
			entry.ProfileImagePath = "/image-resize?name=" + entry.ProfileImagePath
		}
		user, _, _ := NewTeam.GetUserById(entry.CreatedBy, []int{})

		var first = user.FirstName
		var last = user.LastName
		var firstn string
		var lastn string

		if user.FirstName != "" {
			firstn = strings.ToUpper(first[:1])
		}
		if user.LastName != "" {
			lastn = strings.ToUpper(last[:1])
		}

		user.NameString = firstn + lastn

		entry.AuthorDetail.NameString = user.NameString

		fmt.Println(entry, "entrydetails")
		if err != nil {
			ErrorLog.Printf("entry parentid change error: %s", perr)
			json.NewEncoder(c.Writer).Encode(err)

		} else {
			json.NewEncoder(c.Writer).Encode(entry)
		}
	}

}

func EntryReorder(c *gin.Context) {

	var (
		limt   int
		offset int
	)

	limit := c.Query("limit")

	pageno, _ := strconv.Atoi(c.PostForm("pageno"))

	fmt.Println(pageno, "pagenooo")

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if pageno != 1 && pageno != 0 {
		offset = (pageno - 1) * limt
	} else {

		offset = 1
	}

	fmt.Println(offset, "offsetdetails")

	neworder := c.PostFormArray("neworder[]")

	fmt.Println(neworder, pageno, "entrydata")
	userid := c.GetInt("userid")

	neworderint := make([]int, len(neworder))
	for i, str := range neworder {
		id, _ := strconv.Atoi(str)
		neworderint[i] = id
	}

	err := ChannelConfig.UpdateEntryOrder(neworderint, TenantId, userid, offset)

	if err != nil {

		json.NewEncoder(c.Writer).Encode(err)
	} else {

		json.NewEncoder(c.Writer).Encode(nil)
	}
}

func UpdateAccPermissionMembergroupId(c *gin.Context) {

	membergroupids := c.PostFormArray("memgrpids[]")

	entryid, _ := strconv.Atoi(c.PostForm("entryid"))

	userid := c.GetInt("userid")

	membergrbids := strings.Join(membergroupids, ",")

	err := ChannelConfig.UpdateMemberGroupIds(membergrbids, entryid, TenantId, userid)

	if err != nil {

		json.NewEncoder(c.Writer).Encode(false)
	} else {

		json.NewEncoder(c.Writer).Encode(true)
	}
}
