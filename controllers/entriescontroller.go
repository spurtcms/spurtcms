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
	"github.com/spurtcms/pkgcontent/channels"
	"github.com/spurtcms/pkgcore/member"
	"github.com/spurtcms/team"
	csrf "github.com/utrack/gin-csrf"
)

type PaginationData struct {
	NextPage     int
	PreviousPage int
	CurrectPage  int
	TotalPages   int
	TwoAfter     int
	TwoBelow     int
	ThreeAfter   int
}

var Channel channels.Channel

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
		limt    int
		offset  int
		filters EntriesFilter
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

	chnanem, _, cerr := ChannelConfig.GetChannelsById(id)
	if cerr != nil {
		ErrorLog.Printf("Entries get channel data error :%s", cerr)
	}

	flag := true

	channelist, _, clerr := ChannelConfig.ListChannel(100, 0, chn.Filter{}, flag, true)
	if clerr != nil {
		ErrorLog.Printf("channellist error :%s", clerr)
	}

	/*list*/
	query.DataAccess = c.GetInt("dataaccess")
	query.UserId = c.GetInt("userid")

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD)
	if perr != nil {
		ErrorLog.Printf("Entries authorization error: %s", perr)
	}

	// if permisison {

	chnentry, _, _, _ := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: id, Limit: limt, Offset: offset, Keyword: filters.Keyword, ChannelName: filters.ChannelName, Title: filters.Title, Status: filters.Status})

	_, filtercount, _, _ := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: id, Limit: 0, Offset: 0, Keyword: filters.Keyword, ChannelName: filters.ChannelName, Title: filters.Title, Status: filters.Status})

	_, entrcount, _, _ := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: id, Limit: 0, Offset: 0})

	var allchnentry []chn.Tblchannelentries

	for index, entry := range chnentry {
		entry.Cno = strconv.Itoa((offset + index) + 1)
		entry.CreatedDate = entry.CreatedOn.In(TZONE).Format(Datelayout)
		if !entry.ModifiedOn.IsZero() {
			entry.ModifiedDate = entry.ModifiedOn.In(TZONE).Format(Datelayout)
		} else {
			entry.ModifiedDate = entry.CreatedOn.In(TZONE).Format(Datelayout)
		}
		allchnentry = append(allchnentry, entry)
	}

	var filterflag bool
	if filters.Keyword == "" && filters.ChannelName == "" && filters.Title == "" && id == 0 {
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

	c.HTML(200, "entries.html", gin.H{"Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3,
	}, "Menu": menu, "chcount1": entrcount, "csrf": csrf.GetToken(c), "channelname": chnanem.ChannelName, "chnid": id, "ChanEntrtlist": allchnentry, "entrycount": entrcount, "Next": Next, "Previous": Previous, "PageCount": PageCount, "CurrentPage": pageno, "limit": limt, "paginationendcount": paginationendcount, "paginationstartcount": paginationstartcount, "filter": filters, "title": ModuleName, "heading": chnanem.ChannelName, "translate": translate, "channellist": channelist, "Page": Page, "HeadTitle": translate.Channell.Channels, "chentrycount": filtercount, "Cmsmenu": true, "filterflag": filterflag, "Entriestab": true, "Tabmenu": TabName, "StorageType": selectedtype.SelectedType})

	return

	// }

	// c.Redirect(301, "/403-page")

}

/*create entry*/
func CreateEntry(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	flag := true

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD)
	if perr != nil {
		ErrorLog.Printf("create Entry authorization error: %s", perr)
	}

	// if permisison {

		channelist, _, clerr := ChannelConfig.ListChannel(100, 0, chn.Filter{}, flag, true)
		if clerr != nil {
			ErrorLog.Printf("create Entry listchannel error: %s", clerr)
		}

		_, entrcount, _, cherr := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: id, Limit: 0, Offset: 0})
		if cherr != nil {
			ErrorLog.Printf("createEntry entrylist error: %s", cherr)
		}

		_, FinalSelectedCategories, cerr := ChannelConfig.GetChannelsById(id)
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

		field, err := ChannelConfig.GetAllMasterFieldType()
		if err != nil {
			ErrorLog.Printf("createEntry get master fields error: %s", perr)
		}

		currentTime := time.Now()
		cdate := currentTime.Format(Datelayout)
		ModuleName, _ := c.Cookie("modulename")
		translate, _ := TranslateHandler(c)
		menu := NewMenuController(c)

		Folder, File, Media, merr := GetMedia()
		if merr != nil {
			ErrorLog.Printf("createEntry media error: %s", perr)
		}
		selectedtype, _ := GetSelectedType()

		c.HTML(200, "addentry.html", gin.H{"Menu": menu, "title": ModuleName, "Fields": field, "translate": translate, "Folder": Folder, "File": File, "Media": Media, "csrf": csrf.GetToken(c), "channellist": channelist, "CategoryName": FinalSelectedCategories, "entrycount": entrcount, "HeadTitle": translate.Channell.Channels, "Cmsmenu": true, "Entriestab": true, "currentdate": cdate, "StorageType": selectedtype.SelectedType})

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
	if filters.Status == "" && filters.ChannelName == "" && filters.Title == "" {
		filterflag = true
	} else {
		filterflag = false
	}

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD)
	if perr != nil {
		ErrorLog.Printf("all entries authorization error: %s", perr)
	}

	// if permisison {

		_, Totalentris, _, _ := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: 0, Limit: 0, Offset: 0, Keyword: filters.Keyword, ChannelName: filters.ChannelName, Title: filters.Title, Status: filters.Status})

		_, Totalentris1, _, _ := ChannelConfig.ChannelEntriesList(chn.Entries{})

		chlist, _, _, _ := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: 0, Limit: limt, Offset: offset, Keyword: filters.Keyword, ChannelName: filters.ChannelName, Title: filters.Title, Status: filters.Status})

		var chlist1 []chn.Tblchannelentries

		for index, entry := range chlist {

			entry.Cno = strconv.Itoa((offset + index) + 1)
			entry.CreatedDate = entry.CreatedOn.In(TZONE).Format(Datelayout)
			if !entry.ModifiedOn.IsZero() {
				entry.ModifiedDate = entry.ModifiedOn.In(TZONE).Format(Datelayout)
			} else {
				entry.ModifiedDate = entry.CreatedOn.In(TZONE).Format(Datelayout)
			}
			chlist1 = append(chlist1, entry)
		}

		flag1 := true

		channelist, _, err := ChannelConfig.ListChannel(100, 0, chn.Filter{}, flag1, true)
		if err != nil {
			ErrorLog.Printf("all entries listchannel error: %s", perr)
		}

		//pagination calc
		paginationendcount := len(chlist) + offset
		paginationstartcount := offset + 1
		Previous, Next, PageCount, Page := Pagination(pageno, int(Totalentris), limt)

		menu := NewMenuController(c)
		translate, _ := TranslateHandler(c)
		ModuleName, tabname, _ := ModuleRouteName(c)
		selectedtype, _ := GetSelectedType()

		c.HTML(200, "entries.html", gin.H{"Pagination": PaginationData{
			NextPage:     pageno + 1,
			PreviousPage: pageno - 1,
			TotalPages:   PageCount,
			TwoAfter:     pageno + 2,
			TwoBelow:     pageno - 2,
			ThreeAfter:   pageno + 3,
		}, "Menu": menu, "translate": translate, "Page": Page, "chentrycount": Totalentris, "entrycount": Totalentris1, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "limit": limt, "paginationendcount": paginationendcount, "paginationstartcount": paginationstartcount, "filter": filters, "ChanEntrtlist": chlist1, "csrf": csrf.GetToken(c), "channellist": channelist, "title": ModuleName, "HeadTitle": translate.Channell.Channels, "Cmsmenu": true, "filterflag": filterflag, "Entriestab": true, "Tabmenu": tabname, "StorageType": selectedtype.SelectedType})

		return

	// }

	// c.Redirect(301, "/403-page")

}

func DeleteEntries(c *gin.Context) {

	entryId, _ := strconv.Atoi(c.Query("id"))
	channame := c.Query("cname")
	pageno := c.Query("page")

	var url string
	if pageno != "" {
		url = "/channel/entrylist?page=" + pageno
	} else {
		url = "/channel/entrylist/"
	}

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD)
	if perr != nil {
		ErrorLog.Printf("delete entries authorization error: %s", perr)
	}

	// if permisison {

		userid := c.GetInt("userid")

		_, err := ChannelConfig.DeleteEntry(channame, userid, entryId)
		if err != nil {
			ErrorLog.Printf("delete entries error: %s", perr)
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			return
		}

		c.SetCookie("get-toast", "Entry Deleted Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, url)
		return

	// }

	// c.Redirect(301, "/403-page")

}

func EntryStatus(c *gin.Context) {

	channelname := c.Request.PostFormValue("cname")
	id, _ := strconv.Atoi(c.Request.PostFormValue("entryid"))
	status, _ := strconv.Atoi(c.Request.PostFormValue("status"))
	userid := c.GetInt("userid")

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD)
	if perr != nil {
		ErrorLog.Printf("entry status authorization error: %s", perr)
	}
	// if permisison {

		_, err := ChannelConfig.EntryStatus(channelname, id, status, userid)
		if err != nil {
			ErrorLog.Printf("entry status change error: %s", perr)
			json.NewEncoder(c.Writer).Encode(false)
			return
		}

		json.NewEncoder(c.Writer).Encode(true)
		return

	// }

	// c.Redirect(301, "/403-page")

}

func PublishEntry(c *gin.Context) {

	//get data from html form data
	eid, _ := strconv.Atoi(c.Param("id"))
	cid, _ := strconv.Atoi(c.Request.PostFormValue("id"))
	cname := (c.Request.PostFormValue("cname"))
	title := (c.Request.PostFormValue("title"))
	text := (c.Request.PostFormValue("text"))
	img := (c.Request.PostFormValue("image"))
	authorname := (c.Request.PostFormValue("author"))
	createdate := (c.Request.PostFormValue("createtime"))
	publishtime := (c.Request.PostFormValue("publishtime"))
	readingtime, _ := strconv.Atoi(c.Request.PostFormValue("readingtime"))
	sortorder, _ := strconv.Atoi(c.Request.PostFormValue("sortorder"))
	tagname := (c.Request.PostFormValue("tagname"))
	extxt := (c.Request.PostFormValue("extxt"))
	status, _ := strconv.Atoi(c.Request.PostFormValue("status"))
	categoryids := c.PostFormArray("categoryids[]")
	userid := c.GetInt("userid")

	layout := "2006-01-02T15:04"

	// Parse string to time.Time
	ctTime, _ := time.Parse(layout, createdate)
	pbTime, _ := time.Parse(layout, publishtime)

	categoryidsStr := strings.Join(categoryids, ",")

	var seodetails map[string]string

	if err := json.Unmarshal([]byte(c.Request.PostFormValue("seodetails")), &seodetails); err != nil {
		ErrorLog.Printf("publishentry unmarshall seodetails error: %s", err)
	}

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD)
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

			additionalField := chn.AdditionalFields{
				FieldName:  channel["name"],
				FieldValue: channel["value"],
				FieldId:    FieldidInt,
				Id:         Fid,
			}

			AdditionalFields = append(AdditionalFields, additionalField)
		}

		if eid != 0 {

			entries.ModifiedBy = userid
			_, err := ChannelConfig.UpdateEntry(entries, cname, eid)
			ChannelConfig.UpdateAdditionalField(AdditionalFields, eid)

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

			entries.CreatedBy = userid
			chenid, _, err1 := ChannelConfig.CreateEntry(entries)
			ChannelConfig.CreateChannelEntryFields(chenid.Id, userid, AdditionalFields)
			channeldet, _, _ := ChannelConfig.GetChannelsById(chenid.ChannelId)

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
			return
		}

		return

	// }

	// c.Redirect(301, "/403-page")

}

func EditEntry(c *gin.Context) {

	flag1 := true

	permisison, perr := NewAuth.IsGranted("Entries", auth.CRUD)
	if perr != nil {
		ErrorLog.Printf("edit entry page authorization error: %s", perr)
	}

	if permisison {

		channelist, _, cherr := ChannelConfig.ListChannel(100, 0, chn.Filter{}, flag1, true)
		if cherr != nil {
			ErrorLog.Printf("edit entry get channel list error: %s", cherr)
		}

		field, err := ChannelConfig.GetAllMasterFieldType()
		if err != nil {
			ErrorLog.Printf("edit entry get master fields error: %s", err)
		}

		AllCategorieswithSubCategories, _ := CategoryConfig.AllCategoriesWithSubList()

		menu := NewMenuController(c)

		_, _, Media, merr := GetMedia()
		if merr != nil {
			ErrorLog.Printf("edit entry media error: %s", merr)
		}

		translate, _ := TranslateHandler(c)
		selectedtype, _ := GetSelectedType()

		c.HTML(200, "addentry.html", gin.H{"Menu": menu, "Fields": field, "translate": translate, "csrf": csrf.GetToken(c), "channellist": channelist, "AllCategories": AllCategorieswithSubCategories, "Media": Media, "HeadTitle": translate.Channell.Channels, "Cmsmenu": true, "Entriestab": true, "StorageType": selectedtype.SelectedType})

		return

	}

	c.Redirect(301, "/403-page")

}

func EditDetails(c *gin.Context) {

	mauth := member.Memberauth{Authority: &AUTH}
	var Createtime, Publishedtime string

	id, _ := strconv.Atoi(c.Param("id"))
	cname := c.Param("channelname")

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD)
	if perr != nil {
		ErrorLog.Printf("edit entry authorization error: %s", perr)
	}

	// if permisison {

		entries, err := ChannelConfig.EntryDetailsById(chn.IndivEntriesReq{ChannelName: cname, EntryId: id})

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

		memberfirstname, err := mauth.GetMemberDetails(memid)
		if err != nil {
			ErrorLog.Printf("edit entry get member details error: %s", err)
		}

		_, FinalSelectedCategories, cerr := ChannelConfig.GetChannelsById(channelid)

		if cerr != nil {
			ErrorLog.Printf("edit entry get channeldetails error: %s", perr)
		}

		Section, Fieldvalue, ferr := ChannelConfig.GetChannelsFieldsById(channelid)
		if ferr != nil {
			ErrorLog.Printf("edit entry channel field data error: %s", perr)
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

	channelAuth.Authority = &AUTH

	id, _ := strconv.Atoi(c.Param("id"))

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD)
	if perr != nil {
		ErrorLog.Printf("channelfield  error: %s", perr)
	}

	// if permisison {

		_, FinalSelectedCategories, cerr := ChannelConfig.GetChannelsById(id)
		if cerr != nil {
			ErrorLog.Printf("channel  error: %s", cerr)
		}

		Section, Fieldvalue, ferr := ChannelConfig.GetChannelsFieldsById(id)
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

		json.NewEncoder(c.Writer).Encode(gin.H{"Section": Section, "FieldValue": Fieldvalue, "SelectedCategory": idstr, "CategoryName": FinalSelectedCategories})

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

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD)
	if perr != nil {
		ErrorLog.Printf("update channelfield authorization  error: %s", perr)
	}

	// if permisison {

		channeldetails, _, cerr := ChannelConfig.GetChannelsById(channelid)
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

		userid := c.GetInt("userid")

		if ederr := ChannelConfig.EditChannel(channelname, channeldesc, userid, channelid, categoryvalue); ederr != nil {
			ErrorLog.Printf("update channel error: %s", ederr)
		}

		if updatechn := ChannelConfig.UpdateChannelField(chn.ChannelUpdate{Sections: Sectionss, FieldValues: FieldValuess, Deletesections: DeleteSectionss, DeleteFields: DeleteFieldValuess, DeleteOptionsvalue: DeleteOptionsvalues, CategoryIds: categoryvalue, ModifiedBy: userid}, channelid); updatechn != nil {
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

	channelAuth.Authority = &AUTH

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD)
	if perr != nil {
		ErrorLog.Printf("make feature error: %s", perr)
	}

	// if permisison {

		chn, _ := ChannelConfig.GetchannelByName(channelname)

		_, err := ChannelConfig.MakeFeature(chn.Id, entryid, feature)
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

	auth := member.Memberauth{Authority: &AUTH}

	var filter member.Filter

	filter.FirstName = c.Query("keyword")

	flag1 := false

	if filter.FirstName != "" {

		memberlist, _, err := auth.ListMembers(0, 0, filter, flag1)

		if err != nil {

			fmt.Println(err)
		}
		c.JSON(200, memberlist)
	} else {

		c.JSON(200, "")
	}

}

func CheckEnriesOrder(c *gin.Context) {

	Channel.Authority = &AUTH
	chid, _ := strconv.Atoi(c.PostForm("chid"))
	eid, _ := strconv.Atoi(c.PostForm("eid"))
	order, _ := strconv.Atoi(c.PostForm("order"))

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD)
	if perr != nil {
		ErrorLog.Printf("check entreis order error: %s", perr)
	}

	// if permisison {

		chnentry, _, _, err := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: chid, Limit: 0, Offset: 0})
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

	permisison, perr := NewAuth.IsGranted("Team", auth.CRUD)
	if perr != nil {
		ErrorLog.Printf("user detaisl authorization error: %s", perr)
	}

	if permisison {

		if fname != "" {

			userlist, _, err := NewTeam.ListUser(10, 0, team.Filters{FirstName: fname})
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

	if err := json.Unmarshal([]byte(c.Request.PostFormValue("entryids")), &entrydata); err != nil {
		ErrorLog.Printf("delete multiple entry error: %s", err)
	}

	for _, val := range entrydata {
		entryid, _ := strconv.Atoi(val["entryid"])
		entryids = append(entryids, entryid)
	}

	if pageno != "" {
		url = "/channel/entrylist?page=" + pageno
	} else {
		url = "/channel/entrylist/"
	}

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD)
	if perr != nil {
		ErrorLog.Printf("delete multiplle entries authorization error: %s", perr)
	}

	// if permisison {

		_, err := ChannelConfig.DeleteSelectedEntry(entryids, userid)
		if err != nil {
			ErrorLog.Printf("delete multiplle entries authorization error: %s", err)
			c.JSON(200, gin.H{"value": false})
			return
		}

		c.JSON(200, gin.H{"value": true, "url": url})
		return

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

	if pageno != "" {
		url = "/channel/entrylist?page=" + pageno
	} else {
		url = "/channel/entrylist/"
	}

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD)
	if perr != nil {
		ErrorLog.Printf("unpublished multiplle entries authorization error: %s", perr)
	}

	// if permisison {

		_, err := Channel.UnpublishSelectedEntry(entryids, statusint)
		if err != nil {
			ErrorLog.Printf("unpublished multiplle entries error: %s", err)
			c.JSON(200, gin.H{"value": false})
			return
		}

		c.JSON(200, gin.H{"value": true, "status": statusint, "url": url})
		return

	// }

	// c.Redirect(301, "/403-page")

}
