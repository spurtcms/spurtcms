package controllers

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	chn "github.com/spurtcms/channels"
	mem "github.com/spurtcms/member"
	memaccess "github.com/spurtcms/member-access"
	lms "github.com/spurtcms/pkgcontent/spaces"
	csrf "github.com/utrack/gin-csrf"
)

type subpage1 struct {
	Subpage []memaccess.SubPage `json:"subpages"`
}

type page1 struct {
	Page []memaccess.Page `json:"pages"`
}

type space struct {
	Space []string `json:"spaces"`
}

type memgrp struct {
	Memgrp []string `json:"access_granted_memgrps"`
}

type entry struct {
	Entries []memaccess.Entry `json:"channelEntries"`
}

/*content access list*/
func ContentAccessControlList(c *gin.Context) {

	var (
		limt   int
		offset int
		filter memaccess.Filter
	)

	//get values from html from data
	limit := strconv.Itoa(Limit)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if page != 0 {
		offset = (page - 1) * limt
	}

	permisison, _ := NewAuth.IsGranted("Member Restrict", auth.CRUD)
	if permisison {
		
		var FinalContentAccessLists []memaccess.Tblaccesscontrol
		
		//content access list
		contentAccessLists, content_access_count, err := MemberaccessConfig.ContentAccessList(limt, offset, filter)
		if err != nil {
			ErrorLog.Printf("memberaccess list details error: %s", err)
		}
		
		for _, val := range contentAccessLists {
			
			if !val.ModifiedOn.IsZero() {
				val.DateString = val.ModifiedOn.In(TZONE).Format("02 Jan 2006 03:04 PM")
			} else {
				val.DateString = val.CreatedOn.In(TZONE).Format("02 Jan 2006 03:04 PM")
			}

			FinalContentAccessLists = append(FinalContentAccessLists, val)
		}

		if err != nil {
			log.Println(err)
		}

		paginationendcount := len(contentAccessLists) + offset
		paginationstartcount := offset + 1
		Previous, Next, PageCount, pages := Pagination(page, int(content_access_count), limt)
		
		translate, _ := TranslateHandler(c)
		menu := NewMenuController(c)
		ModuleName, TabName, _ := ModuleRouteName(c)
		
		c.HTML(200, "contentaccesscontrollist.html", gin.H{"Pagination": PaginationData{
			NextPage:     page + 1,
			PreviousPage: page - 1,
			TotalPages:   PageCount,
			TwoAfter:     page + 2,
			TwoBelow:     page - 2,
			ThreeAfter:   page + 3,
		}, "Menu": menu, "title": ModuleName, "csrf": csrf.GetToken(c), "translate": translate, "Count": content_access_count, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": page, "Limit": limt, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "Pages": pages, "Filter": filter, "ContentAccess": FinalContentAccessLists, "Length": len(FinalContentAccessLists), "HeadTitle": translate.Memberss.Members, "Membermenu": true, "Restrictmenu": true, "Tabmenu": TabName})
		return
	}
	c.Redirect(301, "/403-page")

}

/*create content access control*/
func NewContentAccess(c *gin.Context) {

	var (
		spaceDefault   string
		ChannelDetails []chn.Tblchannel
	)

	spaceAuth := lms.MemberSpace{MemAuth: &AUTH}
	filter := lms.Filter{Keyword: strings.Trim(c.DefaultQuery("keyword", ""), " ")}

	permisison, _ := NewAuth.IsGranted("Member Restrict", auth.CRUD)
	if permisison {

		_, space_count, err := spaceAuth.MemberSpaceList(0, 0, filter)
		if err != nil {
			log.Println(err)
		}

		spaceDetails, _, err := spaceAuth.MemberSpaceList(int(space_count), 0, filter)
		if err != nil {
			log.Println(err)
		}

		channelCount, err := ChannelConfig.GetChannelCount() //get channel count
		if err != nil {
			ErrorLog.Printf("channel count error: %s", err)
		}

		_, memgrpcount, err := MemberConfig.ListMemberGroup(mem.MemberGroupListReq{ActiveGroupsOnly: true}) // get active member group count
		if err != nil {
			ErrorLog.Printf("channel count error: %s", err)
		}

		membergroup, _, err := MemberConfig.ListMemberGroup(mem.MemberGroupListReq{Limit: int(memgrpcount), ActiveGroupsOnly: true}) // get active member group
		if err != nil {
			ErrorLog.Printf("channel count error: %s", err)
		}

		menu := NewMenuController(c)
		translate, _ := TranslateHandler(c)

		if os.Getenv("SPACE_DEFAULT") == "yes" {
			spaceDefault = "space"

		} else {
			
			ChannelDetails, err = ChannelConfig.GetChannelsWithEntries() //get active channel&entries list
			if err != nil {
				ErrorLog.Printf("channel count error: %s", err)
			}
			spaceDefault = "channel"
		}

		ModuleName, TabName, _ := ModuleRouteName(c)
		
		c.HTML(200, "addcontentaccess.html", gin.H{"Menu": menu, "csrf": csrf.GetToken(c), "translate": translate, "SpaceDetails": spaceDetails, "Filter": filter, "Membergroup": membergroup, "Button": "Save", "MembergroupLength": len(membergroup), "SpaceLength": len(spaceDetails), "HeadTitle": translate.Memberaccess, "SpaceCount": space_count, "ChannelDetails": ChannelDetails, "ChannelLength": len(ChannelDetails), "ChannelCount": channelCount, "MemberCount": memgrpcount, "Membermenu": true, "Restrictmenu": true, "SpaceDefault": spaceDefault, "Tabmenu": TabName, "title": ModuleName, "Accesscontent": "Create Memberaccess"})
		return
	}
	c.Redirect(301, "/403-page")
}

func GetPages(c *gin.Context) {

	spaceid, _ := strconv.Atoi(c.Query("sid"))
	
	var Pg lms.Page
	Pg.Authority = &AUTH
	
	page_groups, pages, subpages, err := Pg.PageList(spaceid)
	if err != nil {
		log.Println(err)
		return
	}
	
	c.JSON(200, gin.H{"Pagegroups": page_groups, "Pages": pages, "Subpages": subpages})
}

func GrantContentAccessControl(c *gin.Context) {

	var (
		page_parse       page1
		subpage_parse    subpage1
		space_parse      space
		entriez          entry
		memgrp_parse     memgrp
		conv_space_array []int
		conv_memgrps     []int
	)

	//get values from html form data
	title := c.PostForm("title")
	spaces := c.PostForm("spaces")
	pagez := c.PostForm("pages")
	subpages := c.PostForm("subpages")
	entries := c.PostForm("entries")
	memgrps := c.PostForm("memgrps")
	userid := c.GetInt("userid")

	json.Unmarshal([]byte(pagez), &page_parse)
	json.Unmarshal([]byte(subpages), &subpage_parse)
	json.Unmarshal([]byte(spaces), &space_parse)
	json.Unmarshal([]byte(memgrps), &memgrp_parse)
	json.Unmarshal([]byte(entries), &entriez)
	
	subpage_array := subpage_parse.Subpage
	page_array := page_parse.Page
	space_array := space_parse.Space
	memgrp_array := memgrp_parse.Memgrp
	entries_array := entriez.Entries

	if title == "" || len(memgrp_array) == 0 || (len(space_array) == 0 && len(page_array) == 0 && len(subpage_array) == 0 && len(entries_array) == 0) {

		c.SetCookie("Alert-msg", "Please grant atleast one page access to MemberGroups", 3600, "", "", false, false)

		return
	}
	
	for _, space_id := range space_array {
		conv_space_id, _ := strconv.Atoi(space_id)
		conv_space_array = append(conv_space_array, conv_space_id)
	}

	for _, memgrp_id := range memgrp_array {
		conv_memgrp_id, _ := strconv.Atoi(memgrp_id)
		conv_memgrps = append(conv_memgrps, conv_memgrp_id)
	}

	permisison, _ := NewAuth.IsGranted("Member Restrict", auth.CRUD)
	if permisison {
		
		acc, cerr := MemberaccessConfig.CreateAccessControl(title, userid) //create contentaccess
		if cerr != nil {
			ErrorLog.Printf("memberaccess create error: %s", cerr)
		}
		
		cherr := MemberaccessConfig.CreateRestrictEntries(acc.Id, conv_memgrps, entries_array, userid) //create accessentries
		if cherr != nil {
			ErrorLog.Printf("createaccesscontrolpage error: %s", cherr)
		}

		// isAccessCreated, err := auth.CreateMemberAccessControl(access_creation_input)

		if cerr != nil {

			log.Println(cerr)

			json.NewEncoder(c.Writer).Encode(false)
		}

		c.SetCookie("get-toast", "Content Access Rights Granted Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		json.NewEncoder(c.Writer).Encode(true)
		return
	}
	c.Redirect(301, "/403-page")
}

func EditContentAcess(c *gin.Context) {

	var (
		spaceDefault   string
		ChannelDetails []chn.Tblchannel
	)

	content_access_id, _ := strconv.Atoi(c.Param("accessId"))

	permisison, _ := NewAuth.IsGranted("Member Restrict", auth.CRUD)
	if permisison {
		
		accessControl, err := MemberaccessConfig.GetControlAccessById(content_access_id)
		if err != nil {
			ErrorLog.Printf("memberaccess list details error: %s", err)
		}
		
		spaceAuth := lms.MemberSpace{MemAuth: &AUTH}
		filter := lms.Filter{Keyword: strings.Trim(c.DefaultQuery("keyword", ""), " ")}
		_, space_count, err := spaceAuth.MemberSpaceList(0, 0, filter)
		if err != nil {
			log.Println(err)
		}
		
		spaceDetails, _, err := spaceAuth.MemberSpaceList(int(space_count), 0, filter)
		if err != nil {
			log.Println(err)
		}

		channelCount, err := ChannelConfig.GetChannelCount()
		if err != nil {
			ErrorLog.Printf("getchannelcount error: %s", err)
		}
		_, memgrpcount, err := MemberConfig.ListMemberGroup(mem.MemberGroupListReq{ActiveGroupsOnly: true}) //get membergroup count
		if err != nil {
			ErrorLog.Printf("membergroupcount error: %s", err)
		}

		membergroup, _, err := MemberConfig.ListMemberGroup(mem.MemberGroupListReq{Limit: int(memgrpcount), ActiveGroupsOnly: true}) //get active member group
		if err != nil {
			ErrorLog.Printf("membergroup error: %s", err)
		}

		translate, _ := TranslateHandler(c)

		menu := NewMenuController(c)

		if os.Getenv("SPACE_DEFAULT") == "yes" {

			spaceDefault = "space"

		} else {

			ChannelDetails, err = ChannelConfig.GetChannelsWithEntries() //get active channel&entries list
			if err != nil {
				ErrorLog.Printf("getchannelentries error: %s", err)
			}
			spaceDefault = "channel"
		}

		ModuleName, TabName, _ := ModuleRouteName(c)
		if strings.Contains(c.Request.URL.Path, "copy-accesscontrol") {
			c.HTML(200, "addcontentaccess.html", gin.H{"Menu": menu, "csrf": csrf.GetToken(c), "translate": translate, "SpaceDetails": spaceDetails, "Filter": filter, "Membergroup": membergroup, "MembergroupLength": len(membergroup), "SpaceLength": len(spaceDetails), "HeadTitle": translate.Memberaccess, "SpaceCount": space_count, "ChannelDetails": ChannelDetails, "ChannelLength": len(ChannelDetails), "ChannelCount": channelCount, "MemberCount": memgrpcount, "AccessControl": accessControl, "Button": "Copy", "Membermenu": true, "Restrictmenu": true, "SpaceDefault": spaceDefault, "Tabmenu": TabName, "title": ModuleName, "Accesscontent": "Edit Memberaccess"})
			return
		}
		
		c.HTML(200, "addcontentaccess.html", gin.H{"Menu": menu, "csrf": csrf.GetToken(c), "translate": translate, "SpaceDetails": spaceDetails, "Filter": filter, "Membergroup": membergroup, "MembergroupLength": len(membergroup), "SpaceLength": len(spaceDetails), "HeadTitle": translate.Memberaccess, "SpaceCount": space_count, "ChannelDetails": ChannelDetails, "ChannelLength": len(ChannelDetails), "ChannelCount": channelCount, "MemberCount": memgrpcount, "AccessControl": accessControl, "Button": "Update", "Membermenu": true, "Restrictmenu": true, "SpaceDefault": spaceDefault, "Tabmenu": TabName, "title": ModuleName, "Accesscontent": "Edit Memberaccess"})
		return
	}
	c.Redirect(301, "/403-page")
}

func GetContentAccessPages(c *gin.Context) {

	var (
		conv_memgrpIds  []string
		conv_spaceIds   []string
		conv_channelIds []string
	)

	accessId, _ := strconv.Atoi(c.Query("content_access_id"))
	
	permisison, _ := NewAuth.IsGranted("Member Restrict", auth.CRUD)
	if permisison {
		
		accessControl, aerr := MemberaccessConfig.GetControlAccessById(accessId) //content access
		if aerr != nil {
			ErrorLog.Printf("getmemberaccess error: %s", aerr)
		}
		
		spaceIds, serr := MemberaccessConfig.GetselectedSpacesByAccessControlId(accessId) //access spaces
		if serr != nil {
			ErrorLog.Printf("getspaces error: %s", serr)
		}
		
		pagegroups, gerr := MemberaccessConfig.GetselectedGroupByAccessControlId(accessId) //access pagegroup
		if gerr != nil {
			ErrorLog.Printf("getpagegroup error: %s", gerr)
		}
		
		pages, perr := MemberaccessConfig.GetselectedPageByAccessControlId(accessId) //access page
		if perr != nil {
			ErrorLog.Printf("getpages error: %s", perr)
		}
		
		subpages, sperr := MemberaccessConfig.GetselectedSubPageByAccessControlId(accessId) //access subpage
		if sperr != nil {
			ErrorLog.Printf("getsubpages error: %s", sperr)
		}
		
		membergroupIds, merr := MemberaccessConfig.GetaccessMemberGroup(accessId) //access membergroup
		if merr != nil {
			ErrorLog.Printf("getaccessmembergroup error: %s", merr)
		}
		
		channelIds, channelEntries, cerr := MemberaccessConfig.GetselectedEntiresByAccessControlId(accessId) //access entries
		if cerr != nil {
			ErrorLog.Printf("getaccessentries list details error: %s", cerr)
		}

		for _, memgrpId := range membergroupIds {
			conv_memgrpId := strconv.Itoa(memgrpId)
			conv_memgrpIds = append(conv_memgrpIds, conv_memgrpId)
		}

		for _, spaceId := range spaceIds {
			conv_spaceid := spaceId
			conv_spaceIds = append(conv_spaceIds, conv_spaceid)

		}
		
		for _, channelId := range channelIds {
			conv_channelid := strconv.Itoa(channelId)
			conv_channelIds = append(conv_channelIds, conv_channelid)
		}

		c.JSON(200, gin.H{"AccessId": accessControl, "Pages": pages, "Subpages": subpages, "Pagegroups": pagegroups, "MemberGroups": conv_memgrpIds, "Spaces": conv_spaceIds, "Channels": conv_channelIds, "ChannelEntries": channelEntries})
		return
	}
	c.Redirect(301, "/403-page")
}

func UpdateContentAccessControl(c *gin.Context) {

	var (
		page_parse       page1
		subpage_parse    subpage1
		space_parse      space
		memgrp_parse     memgrp
		entriez          entry
		conv_space_array []int
		conv_memgrps     []int
	)

	//get values from hmtl form data
	accessId, _ := strconv.Atoi(c.PostForm("content_access_id"))
	title := c.PostForm("title")
	spaces := c.PostForm("spaces")
	pagez := c.PostForm("pages")
	subpages := c.PostForm("subpages")
	entries := c.PostForm("entries")
	memgrps := c.PostForm("memgrps")

	json.Unmarshal([]byte(pagez), &page_parse)
	json.Unmarshal([]byte(subpages), &subpage_parse)
	json.Unmarshal([]byte(spaces), &space_parse)
	json.Unmarshal([]byte(memgrps), &memgrp_parse)
	json.Unmarshal([]byte(entries), &entriez)

	subpage_array := subpage_parse.Subpage
	page_array := page_parse.Page
	space_array := space_parse.Space
	memgrp_array := memgrp_parse.Memgrp
	entries_array := entriez.Entries

	if title == "" || len(memgrp_array) == 0 || (len(space_array) == 0 && len(page_array) == 0 && len(subpage_array) == 0 && len(entries_array) == 0) {

		c.SetCookie("Alert-msg", "Please grant atleast one page access to MemberGroups", 3600, "", "", false, false)

		return
	}

	for _, space_id := range space_array {
		conv_space_id, _ := strconv.Atoi(space_id)
		conv_space_array = append(conv_space_array, conv_space_id)
	}

	for _, memgrp_id := range memgrp_array {
		conv_memgrp_id, _ := strconv.Atoi(memgrp_id)
		conv_memgrps = append(conv_memgrps, conv_memgrp_id)
	}

	userid := c.GetInt("userid")

	permisison, _ := NewAuth.IsGranted("Member Restrict", auth.CRUD)
	if permisison {
		
		err := MemberaccessConfig.UpdateAccessControl(accessId, title, userid) //update access control
		if err != nil {
			ErrorLog.Printf("updateaccesscontrol error: %s", err)
		}
		
		uerr := MemberaccessConfig.UpdateRestrictEntries(accessId, conv_memgrps, entries_array, userid) //update accesscontrol page
		if uerr != nil {
			ErrorLog.Printf("updateaccesscontrolpage error: %s", uerr)
		}
		// isAccessUpdated, err := auth.UpdateMemberAccessControl(access_update_input, accessId)

		if err != nil {
			log.Println(err)
			json.NewEncoder(c.Writer).Encode(false)
		}
		
		c.SetCookie("get-toast", "Content Access Rights Updated Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		json.NewEncoder(c.Writer).Encode(true)
		return
	}
	c.Redirect(301, "/403-page")
}

func DeleteAccessControl(c *gin.Context) {
	
	var url string
	accessId, _ := strconv.Atoi(c.Param("accessId"))
	userid := c.GetInt("userid")
	
	permisison, _ := NewAuth.IsGranted("Member Restrict", auth.CRUD)
	if permisison {
		
		err := MemberaccessConfig.DeleteMemberAccessControl(accessId, userid) //deleteaccesscontrol
		if err != nil {
			ErrorLog.Printf("deleteaccescontrol error: %s", err)
		}
		
		pageno := c.Query("page")
		if pageno != "" {
			url = "/memberaccess?page=" + pageno
		} else {
			url = "/memberaccess/"
		}
		
		c.SetCookie("get-toast", "Content Access Rights Deleted Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, url)
		return
	}

	c.Redirect(301, "/403-page")
}

func GetChannelsAndItsEntries(c *gin.Context) {
	
	Channellist, err := ChannelConfig.GetChannelsWithEntries() //get active chennal&entries
	if err != nil {
		log.Println(err)
	}
	
	c.JSON(200, gin.H{"Channels": Channellist})
}
