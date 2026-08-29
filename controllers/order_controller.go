package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	membership "github.com/spurtcms/membership"
	csrf "github.com/utrack/gin-csrf"
)

func OrderList(c *gin.Context) {

	var filter membership.Filter
	var limt int
	var offset int

	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")
	filter.OrderId, _ = strconv.Atoi(strings.Trim(c.DefaultQuery("orderid", ""), " "))
	filter.TransactionId = strings.Trim(c.DefaultQuery("transactionid", ""), " ")
	filter.Level = strings.Trim(c.DefaultQuery("filter-level", ""), " ")

	//Search--datas
	var searchflag bool
	if filter.Keyword != "" {
		searchflag = true
	} else {
		searchflag = false
	}

	//Filter--datat
	var filterflag bool
	if filter.Level != "" || filter.OrderId != 0 || filter.TransactionId != "" {
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

	permisison, perr := NewAuth.IsGranted("Membership", auth.Create, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipOrder authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipOrder authorization error")
		c.Redirect(301, "/403-page")
		return
	}
	var FinalOrderList []membership.TblMembershipOrders
	var TotalCount int64

	MembershiplevelList, _, _ := MembershipConfig.MembershipLevelsList(0, 0, membership.Filter{}, TenantId)

	orderlist, count, err := MembershipConfig.OrderList(limt, offset, filter, TenantId)
	if err != nil {
		fmt.Println(err)
	}
	TotalCount = count

	if permisison {

		for _, val := range orderlist {

			if !val.ModifiedOn.IsZero() {
				val.DateString = val.ModifiedOn.In(TZONE).Format(Datelayout)
			} else {
				val.DateString = val.CreatedOn.In(TZONE).Format(Datelayout)
			}
			FinalOrderList = append(FinalOrderList, val)

		}

	}

	//pagination calc
	paginationendcount := len(orderlist) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(pageno, int(TotalCount), limt)

	translate, _ := TranslateHandler(c)

	menu := NewMenuController(c)

	_, TabName, _ := ModuleRouteName(c)

	c.HTML(200, "orders.html", gin.H{"Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3,
	}, "csrf": csrf.GetToken(c), "Menu": menu, "Searchtrue": searchflag, "Filtertrue": filterflag, "title": "Membership", "linktitle": "Orders", "translate": translate, "Tabmenu": TabName, "Orderlist": FinalOrderList, "Count": TotalCount, "Paginationendcount": paginationendcount, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Filter": filter, "Paginationstartcount": paginationstartcount, "Limit": limt, "Membershiplevellist": MembershiplevelList})

}

func Createorder(c *gin.Context) {

	MembershiplevelList, _, _ := MembershipConfig.MembershipLevelsList(0, 0, membership.Filter{}, TenantId)

	SubscriptionList, _, _ := MembershipConfig.SubscriptionList(0, 0, membership.Filter{}, TenantId)

	translate, _ := TranslateHandler(c)

	menu := NewMenuController(c)

	_, TabName, _ := ModuleRouteName(c)

	c.HTML(200, "createorder.html", gin.H{"csrf": csrf.GetToken(c), "Menu": menu, "title": "Membership", "linktitle": "Create Order", "translate": translate, "Tabmenu": TabName, "Membershiplevellist": MembershiplevelList, "SubscriptionList": SubscriptionList})
}

func AddOrder(c *gin.Context) {

	userid := c.GetInt("userid")

	memberuserid, _ := strconv.Atoi(c.PostForm("userid"))

	membershiplevelid, _ := strconv.Atoi(c.PostForm("membershiplevelid"))

	billingname := c.PostForm("billingname")

	billingstreet := c.PostForm("billingstreet")

	billingstreet2 := c.PostForm("billingstreet2")

	billingcity := c.PostForm("billingcity")

	billingstate := c.PostForm("billingstate")

	billingpostalcode := c.PostForm("billingpostalcode")

	billingcountry := c.PostForm("billingcountry")

	billingphone := c.PostForm("billingphone")

	subtotal, _ := strconv.Atoi(c.PostForm("subtotal"))

	tax, _ := strconv.Atoi(c.PostForm("tax"))

	total := c.PostForm("total")

	paymenttype := c.PostForm("paymenttype")

	status := c.PostForm("status")

	gateway := c.PostForm("gateway")

	gatewayenvironment := c.PostForm("gatewayenvironment")

	paymenttransactionid, _ := strconv.Atoi(c.PostForm("paymenttransactionid"))

	subscriptiontransactionid, _ := strconv.Atoi(c.PostForm("subscriptiontransactionid"))

	permisison, perr := NewAuth.IsGranted("Membership", auth.Create, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipOrder authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipOrder authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	Order := membership.TblMembershipOrder{
		UserId:                    memberuserid,
		MembershiplevelId:         membershiplevelid,
		BillingName:               billingname,
		BillingStreet:             billingstreet,
		BillingStreet2:            billingstreet2,
		BillingCity:               billingcity,
		BillingState:              billingstate,
		BillingPostalcode:         billingpostalcode,
		BillingCountry:            billingcountry,
		BillingPhone:              billingphone,
		SubTotal:                  subtotal,
		Tax:                       tax,
		Total:                     total,
		PaymentType:               paymenttype,
		Status:                    status,
		Gateway:                   gateway,
		GatewayEnvironment:        gatewayenvironment,
		PaymenttransactionId:      paymenttransactionid,
		SubscriptiontransactionId: subscriptiontransactionid,
		CreatedBy:                 userid,
		TenantId:                  TenantId,
	}

	if permisison {
		err := MembershipConfig.CreateOrder(Order)
		if err != nil {
			fmt.Println(err)
			c.Redirect(301, "/admin/order/")
			return
		}
	}

	c.SetCookie("get-toast", "Order Created Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, "/admin/order/")

}

func EditOrder(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipOrder authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipOrder authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		SubscriptionList, _, _ := MembershipConfig.SubscriptionList(0, 0, membership.Filter{}, TenantId)

		MembershiplevelList, _, _ := MembershipConfig.MembershipLevelsList(0, 0, membership.Filter{}, TenantId)

		orderlist, err := MembershipConfig.EditMembershipOrder(id, TenantId)
		if err != nil {
			fmt.Println(err)
		}

		translate, _ := TranslateHandler(c)

		menu := NewMenuController(c)

		_, TabName, _ := ModuleRouteName(c)

		c.HTML(200, "createorder.html", gin.H{"csrf": csrf.GetToken(c), "Menu": menu, "title": "Membership", "linktitle": "Edit Order", "translate": translate, "Tabmenu": TabName, "orderlist": orderlist, "Membershiplevellist": MembershiplevelList, "SubscriptionList": SubscriptionList})
	}

	c.Redirect(301, "/403-page")

}

func UpdateOrder(c *gin.Context) {

	userid := c.GetInt("userid")

	id, _ := strconv.Atoi(c.PostForm("orderid"))

	memberuserid, _ := strconv.Atoi(c.PostForm("userid"))

	membershiplevelid, _ := strconv.Atoi(c.PostForm("membershiplevelid"))

	billingname := c.PostForm("billingname")

	billingstreet := c.PostForm("billingstreet")

	billingstreet2 := c.PostForm("billingstreet2")

	billingcity := c.PostForm("billingcity")

	billingstate := c.PostForm("billingstate")

	billingpostalcode := c.PostForm("billingpostalcode")

	billingcountry := c.PostForm("billingcountry")

	billingphone := c.PostForm("billingphone")

	subtotal, _ := strconv.Atoi(c.PostForm("subtotal"))

	tax, _ := strconv.Atoi(c.PostForm("tax"))

	total := c.PostForm("total")

	paymenttype := c.PostForm("paymenttype")

	status := c.PostForm("status")

	gateway := c.PostForm("gateway")

	gatewayenvironment := c.PostForm("gatewayenvironment")

	paymenttransactionid, _ := strconv.Atoi(c.PostForm("paymenttransactionid"))

	subscriptiontransactionid, _ := strconv.Atoi(c.PostForm("subscriptiontransactionid"))

	permisison, perr := NewAuth.IsGranted("Membership", auth.Update, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipOrder authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipOrder authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	Order := membership.TblMembershipOrder{
		UserId:                    memberuserid,
		MembershiplevelId:         membershiplevelid,
		BillingName:               billingname,
		BillingStreet:             billingstreet,
		BillingStreet2:            billingstreet2,
		BillingCity:               billingcity,
		BillingState:              billingstate,
		BillingPostalcode:         billingpostalcode,
		BillingCountry:            billingcountry,
		BillingPhone:              billingphone,
		SubTotal:                  subtotal,
		Tax:                       tax,
		Total:                     total,
		PaymentType:               paymenttype,
		Status:                    status,
		Gateway:                   gateway,
		GatewayEnvironment:        gatewayenvironment,
		PaymenttransactionId:      paymenttransactionid,
		SubscriptiontransactionId: subscriptiontransactionid,
		ModifiedBy:                userid,
	}

	if permisison {

		err := MembershipConfig.UpdateMembershipOrder(Order, id, TenantId)
		if err != nil {
			fmt.Println(err)
			c.Redirect(301, "/admin/order/")
			return
		}

	}

	c.SetCookie("get-toast", "Order Updated Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, "/admin/order/")
}

func DeleteOrder(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	pageno := c.Query("page")

	userid := c.GetInt("userid")

	var url string

	permisison, perr := NewAuth.IsGranted("Membership", auth.Update, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipOrder authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipOrder authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		err := MembershipConfig.DeleteMembershipOrder(id, userid, TenantId)
		if err != nil {
			fmt.Println(err)
			c.Redirect(301, "/403-page")
			return
		}
	}

	_, count, _ := MembershipConfig.OrderList(0, 0, membership.Filter{}, TenantId)

	if pageno != "" {
		page, _ := strconv.Atoi(pageno)
		page = page - 1
		multi := page * 10
		multiInt64 := int64(multi)
		if count > multiInt64 {
			url = "/admin/order?page=" + pageno
		} else {
			pagee, _ := strconv.Atoi(pageno)
			_page := pagee - 1
			pages := strconv.Itoa(_page)
			url = "/admin/order?page=" + pages
		}
	} else {
		url = "/admin/order/"
	}
	c.SetCookie("get-toast", "Order Deleted Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, url)
}

func MultiSelectDeleteOrder(c *gin.Context) {

	orderids := c.PostFormArray("orderids[]")
	pageno := c.PostForm("page")
	userid := c.GetInt("userid")
	var url string

	permisison, perr := NewAuth.IsGranted("Membership", auth.Update, TenantId)
	if perr != nil {
		ErrorLog.Printf("MembershipOrder authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MembershipOrder authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		orderIntIds := make([]int, len(orderids))
		for i, id := range orderids {
			intId, _ := strconv.Atoi(id)
			orderIntIds[i] = intId
		}

		err := MembershipConfig.MultiSelectDeleteOrder(orderIntIds, userid, TenantId) 
		if err != nil {
			ErrorLog.Printf("MultiSelectOrderDelete error: %s", err)
			c.JSON(200, gin.H{"value": false})
			return
		}

		_, count, _ := MembershipConfig.OrderList(0, 0, membership.Filter{}, TenantId)

		if pageno != "" {
			page, _ := strconv.Atoi(pageno)
			page = page - 1
			multi := page * 10
			multiInt64 := int64(multi)
			if count > multiInt64 {
				url = "/admin/order?page=" + pageno
			} else {
				pagee, _ := strconv.Atoi(pageno)
				_page := pagee - 1
				pages := strconv.Itoa(_page)
				url = "/admin/order?page=" + pages
			}
		} else {
			url = "/admin/order/"
		}
		c.JSON(200, gin.H{"value": true, "url": url})
	}
}
