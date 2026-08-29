package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	// "strings"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	memship "github.com/spurtcms/membership"
	csrf "github.com/utrack/gin-csrf"
)

func ListMembershipgroup(c *gin.Context) {
	var (
		limt   int
		offset int
		filter memship.Filter
	)

	//get data from html url query
	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")
	filter.Status = strings.Trim(c.DefaultQuery("status", ""), " ")
	filter.ToDate = strings.Trim(c.DefaultQuery("lastupdated", ""), " ")

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

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	// flag := false

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipLevelGroup authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipLevelGroup authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		var Finalsubscriptiongroup []memship.TblMstrMembergrouplevel

		subscriptiongroup, TotalGroupcount := MembershipConfig.MembershipGroupList(offset, limt, memship.Filter(filter), TenantId, 0)
		fmt.Println("")

		for _, val := range subscriptiongroup {

			if !val.ModifiedOn.IsZero() {
				val.DateString = val.ModifiedOn.In(TZONE).Format(Datelayout)
			} else {
				val.DateString = val.CreatedOn.In(TZONE).Format(Datelayout)
			}
			Finalsubscriptiongroup = append(Finalsubscriptiongroup, val)

		}

		var paginationendcount = len(subscriptiongroup) + offset
		paginationstartcount := offset + 1
		Previous, Next, PageCount, Page := Pagination(pageno, int(TotalGroupcount), limt)

		translate, _ := TranslateHandler(c)

		menu := NewMenuController(c)

		_, TabName, _ := ModuleRouteName(c)

		c.HTML(200, "membershipgroup.html", gin.H{"csrf": csrf.GetToken(c), "Pagination": PaginationData{
			NextPage:     pageno + 1,
			PreviousPage: pageno - 1,
			TotalPages:   PageCount,
			TwoAfter:     pageno + 2,
			TwoBelow:     pageno - 2,
			ThreeAfter:   pageno + 3,
		}, "endpoint": "/admin/membershipgroup/createmembershiplevelgroup", "SubscriptionGroupList": Finalsubscriptiongroup, "Menu": menu, "title": "Membership", "linktitle": "Membership Group", "translate": translate, "Tabmenu": TabName, "Count": TotalGroupcount, "Previous": Previous, "Next": Next, "Page": Page, "Limit": limt, "Filter": filter, "Paginationstartcount": paginationstartcount, "Paginationendcount": paginationendcount, "Searchtrue": filterflag, "PageCount": PageCount, "CurrentPage": pageno, "Filters": filter.Keyword, "HeadTitle": translate.Memberss.Members, "Membermenu": true, "membermenu": true})

	}

}

func CreateMembershipGroupLevel(c *gin.Context) {
	userid := c.GetInt("userid")

	GroupName := c.PostForm("plangroupname")
	Description := c.PostForm("plangroupdesc")
	// IsActive, _ := strconv.Atoi(c.PostForm("groupactive"))
	IsActive := 1

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipLevelGroup authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipLevelGroup authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		MembershipConfig.MembershipGroupLevelCreate(GroupName, Description, IsActive, TenantId, userid)

		c.SetCookie("get-toast", "MembershiplevelGroupCreatedSuccessfully", 3600, "", "", false, false)
		c.Redirect(301, "/admin/membershipgroup")
	}

}

func EditMembershipGroupLevel(c *gin.Context) {

	membershiplecelgroupID, _ := strconv.Atoi(c.Query("Id"))

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipLevelGroup authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipLevelGroup authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		Membershipgroupdata := MembershipConfig.MembershipGroupLevelEdit(membershiplecelgroupID)
		translate, _ := TranslateHandler(c)

		menu := NewMenuController(c)

		_, TabName, _ := ModuleRouteName(c)

		c.HTML(200, "membershipgroup.html", gin.H{"csrf": csrf.GetToken(c), "endpoint": "/admin/membershipgroup/updatemembershiplevelgroup", "MembershipgroupUpdatedata": Membershipgroupdata, "Menu": menu, "title": "Membership", "linktitle": "Create MembershipGroup", "translate": translate, "Tabmenu": TabName})

	}

}

func UpdateMembershipGroups(c *gin.Context) {
	userid := c.GetInt("userid")

	SubscriptionId, _ := strconv.Atoi(c.PostForm("membershipgroupid"))

	GroupName := c.PostForm("plangroupname")
	Description := c.PostForm("plangroupdesc")
	IsActive, _ := strconv.Atoi(c.PostForm("groupactive"))

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipLevelGroup authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipLevelGroup authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		MembershipConfig.MembershipGrupUpdate(GroupName, Description, IsActive, TenantId, userid, SubscriptionId)
		c.SetCookie("get-toast", "MembershiplevelGroupUpdatedSuccessfully", 3600, "", "", false, false)

		c.Redirect(301, "/admin/membershipgroup")
	}

}

func DeleteMembershipGroup(c *gin.Context) {
	userid := c.GetInt("userid")

	Membershipgroupid, _ := strconv.Atoi(c.Query("Id"))

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipLevelGroup authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipLevelGroup authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		MembershipConfig.MembershipGroupDelete(Membershipgroupid, userid, TenantId)

		c.SetCookie("get-toast", "MembershiplevelGroupDeletedSuccessfully", 3600, "", "", false, false)

		c.Redirect(301, "/admin/membershipgroup")
	}

}

func MultiselectDeleteMembershipGroup(c *gin.Context) {
	userid := c.GetInt("userid")

	pageno := c.PostForm("page")
	MembershipGroupIds := c.PostFormArray("membershipgroupids[]")

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipLevelGroup authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipLevelGroup authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		MembershipGroupIntIds := make([]int, len(MembershipGroupIds))
		for i, id := range MembershipGroupIds {
			intId, _ := strconv.Atoi(id)
			MembershipGroupIntIds[i] = intId
		}

		MembershipConfig.DeleteMultiselectMembershipGroup(MembershipGroupIntIds, userid)

		var filter memship.Filter

		_, TolatmembershipGroupCount := MembershipConfig.MembershipGroupList(0, 0, memship.Filter(filter), TenantId, 0)

		var url string
		if pageno != "" {
			page, _ := strconv.Atoi(pageno)
			page = page - 1
			multi := page * 10
			multiInt64 := int64(multi)
			if TolatmembershipGroupCount > multiInt64 {
				url = "/admin/membershipgroup?page=" + pageno
			} else {
				pagee, _ := strconv.Atoi(pageno)
				_page := pagee - 1
				pages := strconv.Itoa(_page)
				url = "/admin/membershipgroup?page=" + pages
			}
		} else {
			url = "/admin/membershipgroup/"
		}

		c.JSON(200, gin.H{"value": true, "url": url})
	}

}

func MembershipGroupIsactive(c *gin.Context) {
	userid := c.GetInt("userid")

	id, _ := strconv.Atoi(c.Request.PostFormValue("id"))
	val, _ := strconv.Atoi(c.Request.PostFormValue("isactive"))

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipLevelGroup authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipLevelGroup authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		flg, err := MembershipConfig.ChangeMembershipGroupStatus(id, val, userid, TenantId)

		if err != nil {
			ErrorLog.Printf("membership group level status change error: %s", err)
			json.NewEncoder(c.Writer).Encode(flg)

		} else {
			json.NewEncoder(c.Writer).Encode(flg)
		}
	}

}
