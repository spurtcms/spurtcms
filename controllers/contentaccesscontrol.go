package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	chn "github.com/spurtcms/channels"
	mem "github.com/spurtcms/member"
	memaccess "github.com/spurtcms/member-access"
	csrf "github.com/utrack/gin-csrf"
)


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

	var filterflag bool
	if filter.Keyword != "" {
		filterflag = true
	} else {
		filterflag = false
	}


	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if page != 0 {
		offset = (page - 1) * limt
	}

	permisison, perr := NewAuth.IsGranted("Member Restrict", auth.CRUD, TenantId)

	if perr != nil {
		ErrorLog.Printf("Member Restrict authorization error:")
	}

	if !permisison {
		ErrorLog.Printf("Member Restrict authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		var FinalContentAccessLists []memaccess.Tblaccesscontrol

		MemberaccessConfig.DataAccess = c.GetInt("dataaccess")
		MemberaccessConfig.UserId = c.GetInt("userid")

		//content access list
		contentAccessLists, content_access_count, err := MemberaccessConfig.ContentAccessList(limt, offset, filter, TenantId)
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
		}, "Menu": menu, "title": ModuleName, "Searchtrue": filterflag, "linktitle": "Member Restrict", "csrf": csrf.GetToken(c), "translate": translate, "Count": content_access_count, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": page, "Limit": limt, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "Pages": pages, "Filter": filter, "ContentAccess": FinalContentAccessLists, "Length": len(FinalContentAccessLists), "HeadTitle": translate.Memberss.Members, "Membermenu": true, "Restrictmenu": true, "Tabmenu": TabName})

	}

}

/*create content access control*/
func NewContentAccess(c *gin.Context) {

	var (
		spaceDefault   string
		ChannelDetails []chn.Tblchannel
	)

	filter := ""

	permisison, perr := NewAuth.IsGranted("Member Restrict", auth.CRUD, TenantId)

	if perr != nil {
		ErrorLog.Printf("Member Restrict authorization error:")
	}

	if !permisison {
		ErrorLog.Printf("Member Restrict authorization error")
		c.Redirect(301, "/403-page")
		return
	}
	if permisison {

		channelCount, err := ChannelConfig.GetChannelCount(TenantId) //get channel count
		if err != nil {
			ErrorLog.Printf("channel count error: %s", err)
		}

		_, memgrpcount, err := MemberConfig.ListMemberGroup(mem.MemberGroupListReq{ActiveGroupsOnly: true}, TenantId) // get active member group count
		if err != nil {
			ErrorLog.Printf("channel count error: %s", err)
		}

		membergroup, _, err := MemberConfig.ListMemberGroup(mem.MemberGroupListReq{Limit: int(memgrpcount), ActiveGroupsOnly: true}, TenantId) // get active member group
		if err != nil {
			ErrorLog.Printf("channel count error: %s", err)
		}

		menu := NewMenuController(c)
		translate, _ := TranslateHandler(c)


		fmt.Println("check _is", TenantId)
		ChannelDetails, err = ChannelConfig.GetChannelsWithEntries(TenantId) //get active channel&entries list
		if err != nil {
			ErrorLog.Printf("channel count error: %s", err)
		}
		spaceDefault = "channel"

		ModuleName, TabName, _ := ModuleRouteName(c)

		c.HTML(200, "addcontentaccess.html", gin.H{"Menu": menu, "linktitle": "Create Content Access", "csrf": csrf.GetToken(c), "translate": translate, "Filter": filter, "Membergroup": membergroup, "Button": "Save", "MembergroupLength": len(membergroup), "HeadTitle": translate.Memberaccess, "ChannelDetails": ChannelDetails, "ChannelLength": len(ChannelDetails), "ChannelCount": channelCount, "MemberCount": memgrpcount, "Membermenu": true, "Restrictmenu": true, "SpaceDefault": spaceDefault, "Tabmenu": TabName, "title": ModuleName, "Accesscontent": "Create Memberaccess"})
	}

}


func GrantContentAccessControl(c *gin.Context) {

	var (
		
		entriez      entry
		memgrp_parse memgrp
		conv_memgrps []int
	)

	//get values from html form data
	title := c.PostForm("title")
	entries := c.PostForm("entries")
	memgrps := c.PostForm("memgrps")
	userid := c.GetInt("userid")
	fmt.Println("title", title, entries, memgrps)


	json.Unmarshal([]byte(memgrps), &memgrp_parse)
	json.Unmarshal([]byte(entries), &entriez)


	memgrp_array := memgrp_parse.Memgrp
	entries_array := entriez.Entries

	if title == "" || len(memgrp_array) == 0 || len(entries_array) == 0 {

		c.SetCookie("Alert-msg", "Please grant atleast one page access to MemberGroups", 3600, "", "", false, false)

		return
	}


	for _, memgrp_id := range memgrp_array {
		conv_memgrp_id, _ := strconv.Atoi(memgrp_id)
		conv_memgrps = append(conv_memgrps, conv_memgrp_id)
	}

	permisison, perr := NewAuth.IsGranted("Member Restrict", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("Member Restrict authorization error:")
	}

	if !permisison {
		ErrorLog.Printf("Member Restrict authorization error")
		c.Redirect(301, "/403-page")
		return
	}
	if permisison {
		fmt.Println("idcheck", TenantId)

		acc, cerr := MemberaccessConfig.CreateAccessControl(title, userid, TenantId) //create contentaccess
		if cerr != nil {
			ErrorLog.Printf("memberaccess create error: %s", cerr)
		}

		cherr := MemberaccessConfig.CreateRestrictEntries(acc.Id, conv_memgrps, entries_array, userid, TenantId) //create accessentries
		if cherr != nil {
			ErrorLog.Printf("createaccesscontrolpage error: %s", cherr)
		}


		if cerr != nil {

			log.Println(cerr)

			json.NewEncoder(c.Writer).Encode(false)
		}

		c.SetCookie("get-toast", "Content Access Rights Granted Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		json.NewEncoder(c.Writer).Encode(true)

	}

}

func EditContentAcess(c *gin.Context) {

	var (
		spaceDefault   string
		ChannelDetails []chn.Tblchannel
	)

	content_access_id, _ := strconv.Atoi(c.Param("accessId"))

	permisison, perr := NewAuth.IsGranted("Member Restrict", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("Member Restrict authorization error:")
	}

	if !permisison {
		ErrorLog.Printf("Member Restrict authorization error")
		c.Redirect(301, "/403-page")
		return
	}
	if permisison {

		accessControl, err := MemberaccessConfig.GetControlAccessById(content_access_id, TenantId)
		if err != nil {
			ErrorLog.Printf("memberaccess list details error: %s", err)
		}

		channelCount, err := ChannelConfig.GetChannelCount(TenantId)
		if err != nil {
			ErrorLog.Printf("getchannelcount error: %s", err)
		}
		_, memgrpcount, err := MemberConfig.ListMemberGroup(mem.MemberGroupListReq{ActiveGroupsOnly: true}, TenantId) //get membergroup count
		if err != nil {
			ErrorLog.Printf("membergroupcount error: %s", err)
		}

		membergroup, _, err := MemberConfig.ListMemberGroup(mem.MemberGroupListReq{Limit: int(memgrpcount), ActiveGroupsOnly: true}, TenantId) //get active member group
		if err != nil {
			ErrorLog.Printf("membergroup error: %s", err)
		}

		translate, _ := TranslateHandler(c)

		menu := NewMenuController(c)


		ChannelDetails, err = ChannelConfig.GetChannelsWithEntries(TenantId) //get active channel&entries list
		if err != nil {
			ErrorLog.Printf("getchannelentries error: %s", err)
		}
		spaceDefault = "channel"

		ModuleName, TabName, _ := ModuleRouteName(c)
		if strings.Contains(c.Request.URL.Path, "copy-accesscontrol") {
			c.HTML(200, "addcontentaccess.html", gin.H{"Menu": menu, "csrf": csrf.GetToken(c), "translate": translate, "Membergroup": membergroup, "MembergroupLength": len(membergroup), "HeadTitle": translate.Memberaccess, "ChannelDetails": ChannelDetails, "ChannelLength": len(ChannelDetails), "ChannelCount": channelCount, "MemberCount": memgrpcount, "AccessControl": accessControl, "Button": "Copy", "Membermenu": true, "Restrictmenu": true, "SpaceDefault": spaceDefault, "Tabmenu": TabName, "title": ModuleName, "Accesscontent": "Edit Memberaccess"})
			return
		}

		c.HTML(200, "addcontentaccess.html", gin.H{"Menu": menu, "linktitle": "Edit Content Access", "csrf": csrf.GetToken(c), "translate": translate, "Membergroup": membergroup, "MembergroupLength": len(membergroup), "HeadTitle": translate.Memberaccess, "ChannelDetails": ChannelDetails, "ChannelLength": len(ChannelDetails), "ChannelCount": channelCount, "MemberCount": memgrpcount, "AccessControl": accessControl, "Button": "Update", "Membermenu": true, "Restrictmenu": true, "SpaceDefault": spaceDefault, "Tabmenu": TabName, "title": ModuleName, "Accesscontent": "Edit Memberaccess"})

	}

}

func GetContentAccessPages(c *gin.Context) {

	var (
		conv_memgrpIds  []string
		conv_spaceIds   []string
		conv_channelIds []string
	)

	accessId, _ := strconv.Atoi(c.Query("content_access_id"))

	permisison, perr := NewAuth.IsGranted("Member Restrict", auth.CRUD, TenantId)

	if perr != nil {
		ErrorLog.Printf("Member Restrict authorization error:")
	}

	if !permisison {
		ErrorLog.Printf("Member Restrict authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		accessControl, aerr := MemberaccessConfig.GetControlAccessById(accessId, TenantId) //content access
		if aerr != nil {
			ErrorLog.Printf("getmemberaccess error: %s", aerr)
		}


		membergroupIds, merr := MemberaccessConfig.GetaccessMemberGroup(accessId, TenantId) //access membergroup
		if merr != nil {
			ErrorLog.Printf("getaccessmembergroup error: %s", merr)
		}

		channelIds, channelEntries, cerr := MemberaccessConfig.GetselectedEntiresByAccessControlId(accessId, TenantId) //access entries
		if cerr != nil {
			ErrorLog.Printf("getaccessentries list details error: %s", cerr)
		}

		for _, memgrpId := range membergroupIds {
			conv_memgrpId := strconv.Itoa(memgrpId)
			conv_memgrpIds = append(conv_memgrpIds, conv_memgrpId)
		}

		
		for _, channelId := range channelIds {
			conv_channelid := strconv.Itoa(channelId)
			conv_channelIds = append(conv_channelIds, conv_channelid)
		}

		c.JSON(200, gin.H{"AccessId": accessControl, "MemberGroups": conv_memgrpIds, "Spaces": conv_spaceIds, "Channels": conv_channelIds, "ChannelEntries": channelEntries})

	}

}

func UpdateContentAccessControl(c *gin.Context) {

	var (
		
		memgrp_parse memgrp
		entriez      entry
		conv_memgrps []int
	)

	//get values from hmtl form data
	accessId, _ := strconv.Atoi(c.PostForm("content_access_id"))
	title := c.PostForm("title")

	entries := c.PostForm("entries")
	memgrps := c.PostForm("memgrps")


	json.Unmarshal([]byte(memgrps), &memgrp_parse)
	json.Unmarshal([]byte(entries), &entriez)



	memgrp_array := memgrp_parse.Memgrp
	
	entries_array := entriez.Entries

	if title == "" || len(memgrp_array) == 0 || len(entries_array) == 0 {

		c.SetCookie("Alert-msg", "Please grant atleast one page access to MemberGroups", 3600, "", "", false, false)

		return
	}


	for _, memgrp_id := range memgrp_array {
		conv_memgrp_id, _ := strconv.Atoi(memgrp_id)
		conv_memgrps = append(conv_memgrps, conv_memgrp_id)
	}

	userid := c.GetInt("userid")

	permisison, perr := NewAuth.IsGranted("Member Restrict", auth.CRUD, TenantId)

	if perr != nil {
		ErrorLog.Printf("Member Restrict authorization error:")
	}

	if !permisison {
		ErrorLog.Printf("Member Restrict authorization error")
		c.Redirect(301, "/403-page")
		return
	}
	if permisison {

		err := MemberaccessConfig.UpdateAccessControl(accessId, title, userid, TenantId) //update access control
		if err != nil {
			ErrorLog.Printf("updateaccesscontrol error: %s", err)
		}

		uerr := MemberaccessConfig.UpdateRestrictEntries(accessId, conv_memgrps, entries_array, userid, TenantId) //update accesscontrol page
		if uerr != nil {
			ErrorLog.Printf("updateaccesscontrolpage error: %s", uerr)
		}

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

	permisison, perr := NewAuth.IsGranted("Member Restrict", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("Member Restrict authorization error:")
	}

	if !permisison {
		ErrorLog.Printf("Member Restrict authorization error")
		c.Redirect(301, "/403-page")
		return
	}
	if permisison {

		err := MemberaccessConfig.DeleteMemberAccessControl(accessId, userid, TenantId) //deleteaccesscontrol
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

	}

}

func GetChannelsAndItsEntries(c *gin.Context) {

	Channellist, err := ChannelConfig.GetChannelsWithEntries(TenantId) //get active chennal&entries
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, gin.H{"Channels": Channellist})
}
