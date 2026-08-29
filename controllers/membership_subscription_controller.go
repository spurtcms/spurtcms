package controllers

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	membership "github.com/spurtcms/membership"
	memship "github.com/spurtcms/membership"
	csrf "github.com/utrack/gin-csrf"
)

func Subscription(c *gin.Context) {

	var (
		limt   int
		offset int
		filter memship.Filter
	)

	//get data from html url query
	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")
	filter.Gateway = c.Query("gateway")
	filter.Level = c.Query("filter-level")
	filter.TransactionId = c.Query("transactionid")

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

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipSubscription authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipSubscription authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		MembershiplevelList, _, _ := MembershipConfig.MembershipLevelsList(0, 0, membership.Filter{}, TenantId)

		SubscriptionList, SubscriptioCount, err := MembershipConfig.SubscriptionList(offset, limt, memship.Filter(filter), TenantId)

		var FinalSubscriptionList []memship.TblMembershipSubcriptions

		if err != nil {
			log.Fatal("Subscription List Error :", err)
			c.AbortWithError(500, err)
		}

		for _, val := range SubscriptionList {

			if !val.ModifiedOn.IsZero() {
				val.DateString = val.ModifiedOn.In(TZONE).Format(Datelayout)
			} else {
				val.DateString = val.CreatedOn.In(TZONE).Format(Datelayout)
			}

			FinalSubscriptionList = append(FinalSubscriptionList, val)

		}

		var paginationendcount = len(SubscriptionList) + offset
		paginationstartcount := offset + 1
		Previous, Next, PageCount, Page := Pagination(pageno, int(SubscriptioCount), limt)

		translate, _ := TranslateHandler(c)

		menu := NewMenuController(c)

		_, TabName, _ := ModuleRouteName(c)

		c.HTML(200, "subscription.html", gin.H{"csrf": csrf.GetToken(c), "Pagination": PaginationData{
			NextPage:     pageno + 1,
			PreviousPage: pageno - 1,
			TotalPages:   PageCount,
			TwoAfter:     pageno + 2,
			TwoBelow:     pageno - 2,
			ThreeAfter:   pageno + 3,
		}, "Menu": menu, "title": "Membership", "Searchtrue": filterflag, "Count": SubscriptioCount, "Previous": Previous, "Next": Next, "Page": Page, "Limit": limt, "PageCount": PageCount, "CurrentPage": pageno, "Paginationstartcount": paginationstartcount, "Filters": filter, "Paginationendcount": paginationendcount, "linktitle": "Subscription", "translate": translate, "Tabmenu": TabName, "SubscriptionList": FinalSubscriptionList, "MembershiplevelList": MembershiplevelList})

	}

}

func AddNewSubscription(c *gin.Context) {

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipSubscription authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipSubscription authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		MembershipMemberList, _ := MembershipConfig.MembershipListMembers(0, 0, membership.Filter{}, false, TenantId)

		MembershiplevelList, _, _ := MembershipConfig.MembershipLevelsList(0, 0, membership.Filter{}, TenantId)

		translate, _ := TranslateHandler(c)

		menu := NewMenuController(c)

		_, TabName, _ := ModuleRouteName(c)

		c.HTML(200, "createsubscription.html", gin.H{"csrf": csrf.GetToken(c), "action": "Save", "Menu": menu, "title": "Membership", "linktitle": "Create Subscription", "translate": translate, "Tabmenu": TabName, "Membershiplevellist": MembershiplevelList, "endpoint": "/admin/subscription/createmembershipsubscription", "MembershipMemberList": MembershipMemberList})

	}

}

func CreateMembershipSubscription(c *gin.Context) {

	userid := c.GetInt("userid")

	SubscriptionTransactionId := c.PostForm("subscriptiontransactionid")

	Gateway := c.PostForm("gateway")
	GatewayEnvironment := c.PostForm("gatewayenvironment")
	Userid, _ := strconv.Atoi(c.PostForm("userid"))
	MembershipLevelId, _ := strconv.Atoi(c.PostForm("membershiplevelid"))

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipSubscription authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipSubscription authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		var CreateSubscription = membership.TblMembershipSubcriptions{
			MemberId:                  Userid,
			MembershipLevelId:         MembershipLevelId,
			Gateway:                   Gateway,
			GatewayEnvironment:        GatewayEnvironment,
			SubscriptionTransactionId: SubscriptionTransactionId,
			IsActive:                  1,
		}

		err := MembershipConfig.MembershipCreateSubscription(CreateSubscription, TenantId, userid)

		if err != nil {
			log.Fatal("Create Subscription Error :", err)
			c.AbortWithError(500, err)
		}

		c.SetCookie("get-toast", "SubscriptionCreatedSuccessfully", 3600, "", "", false, false)

		c.Redirect(301, "/admin/subscription")
	}

}

func EditSubscription(c *gin.Context) {
	SubscriptionId, _ := strconv.Atoi(c.DefaultQuery("Id", ""))

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipSubscription authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipSubscription authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		EditSubscription, err := MembershipConfig.SubscriptionEdit(SubscriptionId, TenantId)

		if err != nil {
			log.Fatal("SubscriptionEdit Error:", err)
			c.AbortWithError(500, err)
		}

		MembershipMemberList, _ := MembershipConfig.MembershipListMembers(0, Limit, membership.Filter{}, false, TenantId)

		MembershiplevelList, _, _ := MembershipConfig.MembershipLevelsList(0, Limit, membership.Filter{}, TenantId)

		translate, _ := TranslateHandler(c)

		menu := NewMenuController(c)

		_, TabName, _ := ModuleRouteName(c)

		c.HTML(200, "createsubscription.html", gin.H{"csrf": csrf.GetToken(c), "action": "Update", "Menu": menu, "title": "Membership", "linktitle": "Edit Subscription", "translate": translate, "Tabmenu": TabName, "EditSubscription": EditSubscription, "endpoint": "/admin/subscription/updatesubscription", "Membershiplevellist": MembershiplevelList, "MembershipMemberList": MembershipMemberList})

	}

}

func UpdateSubscription(c *gin.Context) {
	pageno := c.Query("page")

	userid := c.GetInt("userid")

	subscriptionid, _ := strconv.Atoi(c.PostForm("subscriptionid"))

	SubscriptionTransactionId := c.PostForm("subscriptiontransactionid")

	Gateway := c.PostForm("gateway")
	GatewayEnvironment := c.PostForm("gatewayenvironment")
	Userid, _ := strconv.Atoi(c.PostForm("userid"))
	MembershipLevelId, _ := strconv.Atoi(c.PostForm("membershiplevelid"))

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipSubscription authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipSubscription authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		var UpdateSubscription = membership.TblMembershipSubcriptions{
			Id:                        subscriptionid,
			MemberId:                  Userid,
			MembershipLevelId:         MembershipLevelId,
			Gateway:                   Gateway,
			GatewayEnvironment:        GatewayEnvironment,
			SubscriptionTransactionId: SubscriptionTransactionId,
		}

		err := MembershipConfig.SubscriptionUpdate(UpdateSubscription, userid, TenantId)

		if err != nil {
			log.Fatal("Subscription Update error:", err)
			c.AbortWithError(500, err)
		}

		var filter memship.Filter

		_, SubscriptioCount, err := MembershipConfig.SubscriptionList(0, 0, memship.Filter(filter), TenantId)

		var url string
		if pageno != "" {
			page, _ := strconv.Atoi(pageno)
			page = page - 1
			multi := page * 10
			multiInt64 := int64(multi)
			if SubscriptioCount > multiInt64 {
				url = "/admin/subscription?page=" + pageno
			} else {
				pagee, _ := strconv.Atoi(pageno)
				_page := pagee - 1
				pages := strconv.Itoa(_page)
				url = "/admin/subscription?page=" + pages
			}
		} else {
			url = "/admin/subscription/"
		}

		c.SetCookie("get-toast", "SubscriptionUpdaredSuccessfully", 3600, "", "", false, false)

		c.Redirect(301, url)
	}

}

func Deletesubscription(c *gin.Context) {
	pageno := c.Query("page")

	userid := c.GetInt("userid")

	SubscriptionId, _ := strconv.Atoi(c.Query("Id"))

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipSubscription authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipSubscription authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		err := MembershipConfig.SubscriptionsDelete(SubscriptionId, userid, TenantId)

		if err != nil {
			log.Fatal("Subscription Delete error:", err)
			c.AbortWithError(500, err)
		}

		var filter memship.Filter

		_, SubscriptioCount, err := MembershipConfig.SubscriptionList(0, 0, memship.Filter(filter), TenantId)

		var url string
		if pageno != "" {
			page, _ := strconv.Atoi(pageno)
			page = page - 1
			multi := page * 10
			multiInt64 := int64(multi)
			if SubscriptioCount > multiInt64 {
				url = "/admin/subscription?page=" + pageno
			} else {
				pagee, _ := strconv.Atoi(pageno)
				_page := pagee - 1
				pages := strconv.Itoa(_page)
				url = "/admin/subscription?page=" + pages
			}
		} else {
			url = "/admin/subscription/"
		}
		c.SetCookie("get-toast", "SubscriptionDeletedSuccessfully", 3600, "", "", false, false)

		c.Redirect(301, url)
	}

}

func SubscriptionIsactive(c *gin.Context) {
	userid := c.GetInt("userid")

	id, _ := strconv.Atoi(c.Request.PostFormValue("id"))
	val, _ := strconv.Atoi(c.Request.PostFormValue("isactive"))

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipSubscription authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipSubscription authorization error")
		c.Redirect(301, "/403-page")
		return
	}
	if permisison {
		_, err := MembershipConfig.ChangesSubscriptionIsactive(id, val, userid, TenantId)

		if err != nil {
			log.Fatal("Subscription Update Issue :", err)
			c.AbortWithStatusJSON(500, err)

		}

		fmt.Println("reachh::")

		c.JSON(200, gin.H{"value": true})
	}

}

func MultiSelectDeleteSubscription(c *gin.Context) {

	fmt.Println("multi deleted::")

	userid := c.GetInt("userid")

	pageno := c.PostForm("page")
	SubsriptionIds := c.PostFormArray("subsriptionids[]")

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipSubscription authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipSubscription authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		SubscriptionIntIds := make([]int, len(SubsriptionIds))
		for i, id := range SubsriptionIds {
			intId, _ := strconv.Atoi(id)
			SubscriptionIntIds[i] = intId
		}

		err := MembershipConfig.DeleteMultiSelectSubscription(SubscriptionIntIds, userid)

		if err != nil {

			log.Fatal("Membership Subscription Delete error", err)

			// c.AbortWithError(500, err)

		}

		var filter memship.Filter

		_, SubscriptioCount, err := MembershipConfig.SubscriptionList(0, 0, memship.Filter(filter), TenantId)

		if err != nil {

			log.Fatal("Membership level count error", err)

			c.AbortWithError(500, err)

		}

		var url string
		if pageno != "" {
			page, _ := strconv.Atoi(pageno)
			page = page - 1
			multi := page * 10
			multiInt64 := int64(multi)
			if SubscriptioCount > multiInt64 {
				url = "/admin/subscription?page=" + pageno
			} else {
				pagee, _ := strconv.Atoi(pageno)
				_page := pagee - 1
				pages := strconv.Itoa(_page)
				url = "/admin/subscription?page=" + pages
			}
		} else {
			url = "/admin/subscription/"
		}

		c.JSON(200, gin.H{"value": true, "url": url})
	}

}
