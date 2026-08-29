package controllers

import (
	"encoding/json"
	"fmt"
	storagecontroller "spurt-cms/storage-controller"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	membership "github.com/spurtcms/membership"
	memship "github.com/spurtcms/membership"
	csrf "github.com/utrack/gin-csrf"
)

func MembershipMemberList(c *gin.Context) {

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

	flag := false

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("Membershipmembers authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Membershipmembers authorization error")
		c.Redirect(301, "/403-page")
		return
	}
	if permisison {

		var FinalMembershipMemberList []memship.TblMembershipMembers

		MembershipMemberList, TolatmemberCount := MembershipConfig.MembershipListMembers(offset, limt, memship.Filter(filter), flag, TenantId)

		for _, val := range MembershipMemberList {

			ExpiryDays := val.PlanDuration * val.MultiplyDuration

			paymenton := val.EndDate

			if !paymenton.IsZero() {

				val.EndDate = paymenton.Add(time.Duration(ExpiryDays) * 24 * time.Hour)

				val.DateStringend = val.EndDate.In(TZONE).Format(Datelayout)

				val.DateString = paymenton.In(TZONE).Format(Datelayout)
			}

			FinalMembershipMemberList = append(FinalMembershipMemberList, val)

		}
		MembershiplevelList, _, _ := MembershipConfig.MembershipLevelsList(0, Limit, membership.Filter{}, TenantId)
		var paginationendcount = len(MembershipMemberList) + offset
		paginationstartcount := offset + 1
		Previous, Next, PageCount, Page := Pagination(pageno, int(TolatmemberCount), limt)

		translate, _ := TranslateHandler(c)

		menu := NewMenuController(c)

		_, TabName, _ := ModuleRouteName(c)

		storagetype, err := GetSelectedType()
		if err != nil {
			fmt.Printf("member list getting storagetype error: %s", err)
		}
		uper, _ := NewAuth.IsGranted("Member", auth.Update, TenantId)

		dper, _ := NewAuth.IsGranted("Member", auth.Delete, TenantId)

		c.HTML(200, "memberss.html", gin.H{"csrf": csrf.GetToken(c), "Pagination": PaginationData{
			NextPage:     pageno + 1,
			PreviousPage: pageno - 1,
			TotalPages:   PageCount,
			TwoAfter:     pageno + 2,
			TwoBelow:     pageno - 2,
			ThreeAfter:   pageno + 3,
		}, "Menu": menu, "endpoint": "/admin/membership/newmember", "Count": TolatmemberCount, "Previous": Previous, "Next": Next, "Page": Page, "Limit": limt, "Filter": filter, "Paginationstartcount": paginationstartcount, "Paginationendcount": paginationendcount, "MemberList": FinalMembershipMemberList, "Searchtrue": filterflag, "PageCount": PageCount, "CurrentPage": pageno, "Filters": filter.Keyword, "title": "Membership", "linktitle": "Membership", "HeadTitle": translate.Memberss.Members, "translate": translate, "Membermenu": true, "membermenu": true, "Tabmenu": TabName, "StorageType": storagetype.SelectedType, "permission": uper, "dpermission": dper, "Membershiplevellist": MembershiplevelList})

	}
}

func MembershipCreateMember(c *gin.Context) {

	userid := c.GetInt("userid")

	IsActive, _ := strconv.Atoi(c.PostForm("mem_activestatus"))
	imagedata := c.PostForm("memcrop_base64")
	var imageName, imagePath string

	storagetype, err := GetSelectedType()

	if err != nil {
		ErrorLog.Printf("error get storage type error: %s", err)
	}

	if storagetype.SelectedType == "local" {
		if imagedata != "" {
			imageName, imagePath, err = ConvertBase64(imagedata, strings.TrimPrefix(storagetype.Local+"/member", "/"))

			if err != nil {
				ErrorLog.Printf("error get storage type error: %s", err)
			}
		}

	} else if storagetype.SelectedType == "aws" {

		tenantDetails, err := NewTeam.GetTenantDetails(TenantId)
		if err != nil {
			fmt.Println(err)
		}

		var imageByte []byte

		if imagedata != "" {
			imageName, imagePath, imageByte, err = ConvertBase64toByte(imagedata, "member")
			if err != nil {
				fmt.Println(err)
			}

			imagePath = tenantDetails.S3FolderName + imagePath

			uerr := storagecontroller.UploadCropImageS3(imageName, imagePath, imageByte)
			if uerr != nil {
				c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
				c.Redirect(301, "/admin/membership/")
				return
			}
		}
	}

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("Membershipmembers authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Membershipmembers authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		MembershipMember := memship.TblMembershipMembers{

			FirstName:        c.PostForm("mem_name"),
			LastName:         c.PostForm("mem_lname"),
			Email:            c.PostForm("mem_email"),
			MobileNo:         c.PostForm("mem_mobile"),
			ProfileImage:     imageName,
			ProfileImagePath: imagePath,
			IsActive:         IsActive,
			Username:         c.PostForm("mem_usrname"),
			Password:         c.PostForm("mem_pass"),
			CreatedBy:        userid,
			StorageType:      storagetype.SelectedType,
			TenantId:         TenantId,
		}

		MembershipConfig.CreateMembershipMembers(MembershipMember)

		c.SetCookie("get-toast", "MembershipMemberCreatedSuccessfully", 3600, "", "", false, false)

	}

	c.Redirect(301, "/admin/membership/")

}

func MembershipEditMember(c *gin.Context) {

	MemberId, _ := strconv.Atoi(c.DefaultQuery("id", ""))

	EditMember := MembershipConfig.EditMembershipMember(MemberId)

	c.JSON(200, gin.H{"EditMember": EditMember, "endpoint": "/admin/membership/updatemember"})
}

func MembershipUpdateMember(c *gin.Context) {
	userid := c.GetInt("userid")
	pageno := c.Query("page")

	IsActive, _ := strconv.Atoi(c.PostForm("mem_activestatus"))

	MemberId, _ := strconv.Atoi(c.PostForm("Update_memberid"))

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("Membershipmembers authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Membershipmembers authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		MembershipMember := memship.TblMembershipMembers{

			Id:        MemberId,
			FirstName: c.PostForm("mem_name"),
			LastName:  c.PostForm("mem_lname"),
			Email:     c.PostForm("mem_email"),
			MobileNo:  c.PostForm("mem_mobile"),
			// ProfileImage:     imageName,
			// ProfileImagePath: imagePath,
			IsActive:  IsActive,
			Username:  c.PostForm("mem_usrname"),
			Password:  c.PostForm("mem_pass"),
			CreatedBy: userid,
			TenantId:  TenantId,
		}

		MembershipConfig.UpdateMembershipMember(MembershipMember)

		var filter memship.Filter

		_, TolatmemberCount := MembershipConfig.MembershipListMembers(0, 0, memship.Filter(filter), false, TenantId)

		var url string
		if pageno != "" {
			page, _ := strconv.Atoi(pageno)
			page = page - 1
			multi := page * 10
			multiInt64 := int64(multi)
			if TolatmemberCount > multiInt64 {
				url = "/admin/membership?page=" + pageno
			} else {
				pagee, _ := strconv.Atoi(pageno)
				_page := pagee - 1
				pages := strconv.Itoa(_page)
				url = "/admin/membership?page=" + pages
			}
		} else {
			url = "/admin/membership/"
		}

		c.SetCookie("get-toast", "MembershipMemberUpdaredSuccessfully", 3600, "", "", false, false)

		c.Redirect(301, url)
	}

}

func MembershipDeleteMember(c *gin.Context) {
	pageno := c.Query("page")

	userid := c.GetInt("userid")

	MemberId, _ := strconv.Atoi(c.Query("Id"))

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("Membershipmembers authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Membershipmembers authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		MembershipConfig.DeleteMembershipMember(MemberId, userid)

		var filter memship.Filter

		_, TolatmemberCount := MembershipConfig.MembershipListMembers(0, 0, memship.Filter(filter), false, TenantId)

		var url string
		if pageno != "" {
			page, _ := strconv.Atoi(pageno)
			page = page - 1
			multi := page * 10
			multiInt64 := int64(multi)
			if TolatmemberCount > multiInt64 {
				url = "/admin/membership?page=" + pageno
			} else {
				pagee, _ := strconv.Atoi(pageno)
				_page := pagee - 1
				pages := strconv.Itoa(_page)
				url = "/admin/membership?page=" + pages
			}
		} else {
			url = "/admin/membership/"
		}

		c.SetCookie("get-toast", "MembershipMemberDeletedSuccessfully", 3600, "", "", false, false)

		c.Redirect(301, url)
	}

}

func MultiselectDeleteMember(c *gin.Context) {
	userid := c.GetInt("userid")

	pageno := c.PostForm("page")
	MemberIds := c.PostFormArray("membersids[]")

	MemberIntIds := make([]int, len(MemberIds))
	for i, id := range MemberIds {
		intId, _ := strconv.Atoi(id)
		MemberIntIds[i] = intId
	}

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("Membershipmembers authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Membershipmembers authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		MembershipConfig.DeleteMultiselectMember(MemberIntIds, userid)

		var filter memship.Filter

		_, TolatmemberCount := MembershipConfig.MembershipListMembers(0, 0, memship.Filter(filter), false, TenantId)

		var url string
		if pageno != "" {
			page, _ := strconv.Atoi(pageno)
			page = page - 1
			multi := page * 10
			multiInt64 := int64(multi)
			if TolatmemberCount > multiInt64 {
				url = "/admin/membership?page=" + pageno
			} else {
				pagee, _ := strconv.Atoi(pageno)
				_page := pagee - 1
				pages := strconv.Itoa(_page)
				url = "/admin/membership?page=" + pages
			}
		} else {
			url = "/admin/membership/"
		}

		c.JSON(200, gin.H{"value": true, "url": url})
	}
}

func MembershipIsactive(c *gin.Context) {
	userid := c.GetInt("userid")

	id, _ := strconv.Atoi(c.Request.PostFormValue("id"))
	val, _ := strconv.Atoi(c.Request.PostFormValue("isactive"))

	permisison, perr := NewAuth.IsGranted("Membership", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("Membershipmembers authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Membershipmembers authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		flg, err := MembershipConfig.ChangeMembershipStatus(id, val, userid, TenantId)

		if err != nil {
			ErrorLog.Printf("membership status change error: %s", err)
			json.NewEncoder(c.Writer).Encode(flg)

		} else {
			json.NewEncoder(c.Writer).Encode(flg)
		}
	}

}
