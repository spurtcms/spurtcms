package controllers

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	memship "github.com/spurtcms/membership"
	csrf "github.com/utrack/gin-csrf"
)

// membership level

func Membership(c *gin.Context) {

	var (
		limt   int
		offset int
		filter memship.Filter
	)

	//get data from html url query
	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")
	filter.Level = strings.Trim(c.DefaultQuery("filter-level", ""), " ")
	filter.FromDate = strings.Trim(c.DefaultQuery("startdate", ""), " ")
	filter.ToDate = strings.Trim(c.DefaultQuery("enddate", ""), " ")

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
		ErrorLog.Printf("MembershipLevel authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipLevel authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		MembershiplevelList, TotalMembershipLevelCount, err := MembershipConfig.MembershipLevelsList(offset, limt, memship.Filter(filter), TenantId)

		// for _,Membershiplevel := range MembershiplevelList{
		// 	Intervel:=Membershiplevel.BillingCyclelimit *Membershiplevel.BillingfrequentType

		// 	switch {
		// 	case Intervel >= 28 && Intervel <= 30:
		// 		fmt.Println("1 MONTH")
		// 	case Intervel > 30 && Intervel <= 60:
		// 		fmt.Println("2 MONTH")
		// 	default:
		// 		fmt.Println("error")
		// 	}

		// }

		// fmt.Println("MembershiplevelList", MembershiplevelList)
		MembershiplevelLists, _, _ := MembershipConfig.MembershipLevelsList(0, Limit, memship.Filter{}, TenantId)

		if err != nil {
			log.Fatal("membership level list error", err)
			c.AbortWithError(500, err)
		}

		var paginationendcount = len(MembershiplevelList) + offset
		paginationstartcount := offset + 1
		Previous, Next, PageCount, Page := Pagination(pageno, int(TotalMembershipLevelCount), limt)

		translate, _ := TranslateHandler(c)

		menu := NewMenuController(c)

		_, TabName, _ := ModuleRouteName(c)

		// fmt.Println("menus::", menu.NameLength)

		c.HTML(200, "membershiplevel.html", gin.H{"csrf": csrf.GetToken(c), "Pagination": PaginationData{
			NextPage:     pageno + 1,
			PreviousPage: pageno - 1,
			TotalPages:   PageCount,
			TwoAfter:     pageno + 2,
			TwoBelow:     pageno - 2,
			ThreeAfter:   pageno + 3,
		}, "Menu": menu, "title": "Membership", "Searchtrue": filterflag, "Membershiplevellist": MembershiplevelLists, "Count": TotalMembershipLevelCount, "Previous": Previous, "Next": Next, "Page": Page, "Limit": limt, "PageCount": PageCount, "CurrentPage": pageno, "Paginationstartcount": paginationstartcount, "Filters": filter.Keyword, "Paginationendcount": paginationendcount, "linktitle": "Membership Level", "translate": translate, "Tabmenu": TabName, "Membershiplist": MembershiplevelList, "Filter": filter})

	}

}

func Newsubscriptionlevel(c *gin.Context) {

	var filter memship.Filter

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipLevel authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipLevel authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		DefaultTemplate := MembershipConfig.GetdefaultMembershiplevelTemplate()

		subscriptiongroup, _ := MembershipConfig.MembershipGroupList(0, 0, filter, TenantId, 1)
		translate, _ := TranslateHandler(c)

		menu := NewMenuController(c)

		_, TabName, _ := ModuleRouteName(c)

		c.HTML(200, "createsubscriptionlevel.html", gin.H{"csrf": csrf.GetToken(c), "endpoint": "/admin/membershiplevel/createmembershiplevel", "action": "Save", "Menu": menu, "title": "Membership", "linktitle": "Create Membership Level", "translate": translate, "Tabmenu": TabName, "Defaulmembershiplevel": DefaultTemplate, "SubscriptionGroupList": subscriptiongroup})

	}

}

func LoadMembershipTemplate(c *gin.Context) {

	membershiplevelId, _ := strconv.Atoi(c.DefaultQuery("membershipid", ""))

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipLevel authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipLevel authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		membershipLevel, err := MembershipConfig.MembershiplevelDetails(membershiplevelId)

		if err != nil {

			log.Fatal("Master default Membership level error", err)
			c.AbortWithStatusJSON(500, err)
		}

		c.JSON(200, gin.H{"SelectedMembership": membershipLevel})
	}

}

func Editmembershiplevel(c *gin.Context) {
	var filter memship.Filter

	Membershipis, _ := strconv.Atoi(c.Query("Id"))

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipLevel authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipLevel authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		GetEditMembership, err := MembershipConfig.MembershiplevelEdit(Membershipis, TenantId)

		if err != nil {
			log.Fatal("Edit membership level error", err)
			// c.AbortWithError(500, err)
		}
		subscriptiongroup, _ := MembershipConfig.MembershipGroupList(0, 0, filter, TenantId, 1)

		for _,Group:=range subscriptiongroup{

			if GetEditMembership.MembergroupLevelId==Group.Id{
                GetEditMembership.GroupId=Group.Id
				GetEditMembership.GroupName=Group.GroupName
			}

		}

		translate, _ := TranslateHandler(c)

		menu := NewMenuController(c)

		_, TabName, _ := ModuleRouteName(c)

		c.HTML(200, "createsubscriptionlevel.html", gin.H{"csrf": csrf.GetToken(c), "endpoint": "/admin/membershiplevel/updatembershiplevel", "action": "Update", "Tabmenu": TabName, "linktitle": "Edit Membership Level", "translate": translate, "Menu": menu, "title": "Membership", "SelectedMembership": GetEditMembership, "SubscriptionGroupList": subscriptiongroup})

	}

}

func CreateMembershipLevels(c *gin.Context) {

	SubscriptionGroupid, _ := strconv.Atoi(c.PostForm("subscriptiongroupid"))
	isRecurringInt, _ := strconv.Atoi(c.PostForm("isrecurring"))

	isdiscount, _ := strconv.Atoi(c.PostForm("Discount"))

	discountpercentage, _ := strconv.Atoi(c.PostForm("discountpercentage"))

	discountedprice := c.PostForm("discountedprice")
	DiscountedPrice, _ := strconv.ParseFloat(discountedprice, 64)

	initialPayment := c.PostForm("initialpayment")
	InitialPayment, _ := strconv.ParseFloat(initialPayment, 64)

	billingAmount := c.PostForm("billingamount")
	BillingAmount, _ := strconv.ParseFloat(billingAmount, 64)

	billingfrequentvalue, _ := strconv.Atoi(c.PostForm("billingfrequentvalue"))

	billingcyclelimit, _ := strconv.Atoi(c.PostForm("billingcyclelimit"))

	customtrial, _ := strconv.Atoi(c.PostForm("customtrial"))

	customtriallimit, _ := strconv.Atoi(c.PostForm("trialbillinglimit"))

	trirbillingamount := c.PostForm("trirbillingamount")
	Trirbillingamount, _ := strconv.ParseFloat(trirbillingamount, 64)

	billingfrequenttype, _ := strconv.Atoi(c.PostForm("billingfrequenttype"))

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipLevel authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipLevel authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		// isactive, _ := strconv.Atoi(c.PostForm("subscriptionactive"))

		if isRecurringInt != 1 {
			BillingAmount = 0.00
			billingfrequentvalue = 0
			billingcyclelimit = 0
			billingfrequenttype = 0
		}

		if customtrial != 1 {
			Trirbillingamount = 0.00
			customtriallimit = 0
		}

		subscriptionlevel := memship.TblMstrMembershiplevel{
			SubscriptionName:       c.PostForm("subscriptionname"),
			Description:            c.PostForm("subscriptiondescription"),
			MembershiplevelDetails: c.PostForm("subscriptiondesc"),
			MembergroupLevelId:     SubscriptionGroupid,
			InitialPayment:         InitialPayment,
			IsDiscount:             isdiscount,
			DiscountPercentage:     discountpercentage,
			DiscountedAmount:       DiscountedPrice,
			RecurrentSubscription:  isRecurringInt,

			BillingAmount:        BillingAmount,
			BillingfrequentValue: billingfrequentvalue,
			BillingfrequentType:  billingfrequenttype,
			BillingCyclelimit:    billingcyclelimit,

			CustomTrial:        customtrial,
			TrialBillingAmount: Trirbillingamount,
			TrialBillingLimit:  customtriallimit,
			IsActive:           1,
		}

		err := MembershipConfig.MembershipLevelsCreate(subscriptionlevel, TenantId)
		if err != nil {

			log.Fatal("Membership level create error", err)

			c.AbortWithError(500, err)

		}

		c.SetCookie("get-toast", "MembershiplevelCreatedSuccessfully", 3600, "", "", false, false)

		c.Redirect(301, "/admin/membershiplevel")
	}

}

func UpdateSubscriptionLevel(c *gin.Context) {

	userid := c.GetInt("userid")

	SubscriptionId, _ := strconv.Atoi(c.PostForm("subscriptionid"))
	SubscriptionGroupid, _ := strconv.Atoi(c.PostForm("subscriptiongroupid"))
	isRecurringInt, _ := strconv.Atoi(c.PostForm("isrecurring"))

	isdiscount, _ := strconv.Atoi(c.PostForm("Discount"))

	discountpercentage, _ := strconv.Atoi(c.PostForm("discountpercentage"))

	discountedprice := c.PostForm("discountedprice")
	DiscountedPrice, _ := strconv.ParseFloat(discountedprice, 64)

	initialPayment := c.PostForm("initialpayment")
	InitialPayment, _ := strconv.ParseFloat(initialPayment, 64)

	billingAmount := c.PostForm("billingamount")
	BillingAmount, _ := strconv.ParseFloat(billingAmount, 64)

	billingfrequentvalue, _ := strconv.Atoi(c.PostForm("billingfrequentvalue"))

	billingcyclelimit, _ := strconv.Atoi(c.PostForm("billingcyclelimit"))

	customtrial, _ := strconv.Atoi(c.PostForm("customtrial"))

	customtriallimit, _ := strconv.Atoi(c.PostForm("trialbillinglimit"))

	trirbillingamount := c.PostForm("trirbillingamount")
	Trirbillingamount, _ := strconv.ParseFloat(trirbillingamount, 64)

	billingfrequenttype, _ := strconv.Atoi(c.PostForm("billingfrequenttype"))

	isactive, _ := strconv.Atoi(c.PostForm("subscriptionactive"))

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipLevel authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipLevel authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		if isRecurringInt != 1 {
			BillingAmount = 0.00
			billingfrequentvalue = 0
			billingcyclelimit = 0
			billingfrequenttype = 0
		}

		if customtrial != 1 {
			Trirbillingamount = 0.00
			customtriallimit = 0
		}

		subscriptionupdate := memship.TblMstrMembershiplevel{
			Id:                     SubscriptionId,
			SubscriptionName:       c.PostForm("subscriptionname"),
			Description:            c.PostForm("subscriptiondescription"),
			MembershiplevelDetails: c.PostForm("subscriptiondesc"),
			MembergroupLevelId:     SubscriptionGroupid,
			InitialPayment:         InitialPayment,
			IsDiscount:             isdiscount,
			DiscountPercentage:     discountpercentage,
			DiscountedAmount:       DiscountedPrice,
			RecurrentSubscription:  isRecurringInt,
			BillingAmount:        BillingAmount,
			BillingfrequentValue: billingfrequentvalue,
			BillingfrequentType:  billingfrequenttype,
			BillingCyclelimit:    billingcyclelimit,
			CustomTrial:        customtrial,
			TrialBillingAmount: Trirbillingamount,
			TrialBillingLimit:  customtriallimit,
			IsActive:           isactive,
			ModifiedBy:         userid,
		}

		err := MembershipConfig.UpdateSubscription(subscriptionupdate, TenantId)
		if err != nil {

			log.Fatal("Membership level update error", err)

			c.AbortWithError(500, err)

		}

		c.SetCookie("get-toast", "MembershiplevelUpdatedSuccessfully", 3600, "", "", false, false)

		c.Redirect(301, "/admin/membershiplevel")
	}

}

func DeleteSubscriptionLevel(c *gin.Context) {
	userid := c.GetInt("userid")

	subscriptionid, _ := strconv.Atoi(c.Query("Id"))

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipLevel authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipLevel authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		err := MembershipConfig.SubscriptionDelete(TenantId, subscriptionid, userid)

		if err != nil {

			log.Fatal("Membership level Delete error", err)

			c.AbortWithError(500, err)

		}

		c.SetCookie("get-toast", "MembershiplevelDeletedSuccessfully", 3600, "", "", false, false)

		c.Redirect(301, "/admin/membershiplevel")
	}

}

func MultiselectDeletemembershipLevel(c *gin.Context) {
	userid := c.GetInt("userid")

	pageno := c.PostForm("page")
	MembershipLevelids := c.PostFormArray("membershiplevelids[]")

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipLevel authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipLevel authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		MembershipLevelIntIds := make([]int, len(MembershipLevelids))
		for i, id := range MembershipLevelids {
			intId, _ := strconv.Atoi(id)
			MembershipLevelIntIds[i] = intId
		}

		err := MembershipConfig.DeleteMultiselectMembershipLevel(MembershipLevelIntIds, userid)

		if err != nil {

			log.Fatal("Membership level multiple Delete error", err)

			c.AbortWithError(500, err)

		}

		var filter memship.Filter

		_, TotalMembershipLevelCount, err := MembershipConfig.MembershipLevelsList(0, 0, memship.Filter(filter), TenantId)

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
			if TotalMembershipLevelCount > multiInt64 {
				url = "/admin/membershiplevel?page=" + pageno
			} else {
				pagee, _ := strconv.Atoi(pageno)
				_page := pagee - 1
				pages := strconv.Itoa(_page)
				url = "/admin/membershiplevel?page=" + pages
			}
		} else {
			url = "/admin/membershiplevel/"
		}

		c.JSON(200, gin.H{"value": true, "url": url})

	}

}

func MembershipLevelIsactive(c *gin.Context) {
	userid := c.GetInt("userid")

	id, _ := strconv.Atoi(c.Request.PostFormValue("id"))
	val, _ := strconv.Atoi(c.Request.PostFormValue("isactive"))

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipLevel authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipLevel authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		flg, err := MembershipConfig.ChangesMembershipLevelIsactive(id, val, userid, TenantId)

		if err != nil {
			ErrorLog.Printf("membership level status change error: %s", err)
			json.NewEncoder(c.Writer).Encode(flg)

		} else {
			json.NewEncoder(c.Writer).Encode(flg)
		}

		c.JSON(200, gin.H{"value": true})
	}

}
