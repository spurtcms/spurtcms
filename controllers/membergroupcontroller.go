package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	mem "github.com/spurtcms/member"
	"github.com/spurtcms/pkgcore/member"
	csrf "github.com/utrack/gin-csrf"
)

type memfilter struct {
	Keyword   string
	Category  string
	Status    string
	FromDate  string
	ToDate    string
	FirstName string
}

/* List Member Group*/
func MemberGroupList(c *gin.Context) {

	var (
		limt   int
		offset int
		filter memfilter
	)

	//get values from html from data
	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")
	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}
	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	permisison, _ := NewAuth.IsGranted("Members Group", auth.Read)
	if permisison {

		// get member group list
		var membergrouplist []mem.Tblmembergroup
		membergrouplists, _, err := MemberConfig.ListMemberGroup(mem.MemberGroupListReq{Limit: limt, Offset: offset, Keyword: filter.Keyword})
		if err != nil {
			ErrorLog.Printf("membergroup list details error: %s", err)
		}

		for _, membergroupvalue := range membergrouplists {
			if !membergroupvalue.ModifiedOn.IsZero() {
				membergroupvalue.DateString = membergroupvalue.ModifiedOn.In(TZONE).Format(Datelayout)
			} else {
				membergroupvalue.DateString = membergroupvalue.CreatedOn.In(TZONE).Format(Datelayout)
			}
			membergrouplist = append(membergrouplist, membergroupvalue)
		}

		_, Total_members, err := MemberConfig.ListMemberGroup(mem.MemberGroupListReq{Keyword: filter.Keyword})
		if err != nil {
			ErrorLog.Printf("membergroup list details error: %s", err)
		}

		var paginationendcount = len(membergrouplists) + offset
		paginationstartcount := offset + 1
		Previous, Next, PageCount, Page := Pagination(pageno, int(Total_members), limt)
		menu := NewMenuController(c)
		ModuleName, TabName, _ := ModuleRouteName(c)
		translate, _ := TranslateHandler(c)
		c.HTML(200, "groups.html", gin.H{"csrf": csrf.GetToken(c), "Pagination": PaginationData{
			NextPage:     pageno + 1,
			PreviousPage: pageno - 1,
			TotalPages:   PageCount,
			TwoAfter:     pageno + 2,
			TwoBelow:     pageno - 2,
			ThreeAfter:   pageno + 3,
		}, "title": ModuleName, "translate": translate, "Menu": menu, "Limit": limt, "MemberGroupList": membergrouplist, "Count": Total_members, "Previous": Previous, "Page": Page, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "HeadTitle": translate.Memberss.Members, "Filter": filter, "Membermenu": true, "Groupmenu": true, "Tabmenu": TabName})

	} else {

		c.Redirect(301, "/403-page")
	}
}

/* Create Member Group*/
func CreateNewMemberGroup(c *gin.Context) {

	//get values from html from data
	userid := c.GetInt("userid")
	MemberGroup := mem.MemberGroupCreation{
		Name:        c.PostForm("membergroup_name"),
		Description: c.PostForm("membergroup_desc"),
		CreatedBy:   userid,
	}

	permisison, _ := NewAuth.IsGranted("Members Group", auth.Create)
	if permisison {

		err := MemberConfig.CreateMemberGroup(MemberGroup) //Create member group
		if err != nil {
			if err != nil {
				ErrorLog.Printf("create membergroup error: %s", err)
			}
		}
		
		if strings.Contains(fmt.Sprint(err), "given some values is empty") {
			c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
			c.Redirect(301, "/membersgroup/")
			return
		}

		if err != nil {
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
			c.Redirect(301, "/membersgroup/")
			return
		}

		c.SetCookie("get-toast", "Member Group Created Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, "/membersgroup/")

	} else {
		c.Redirect(301, "/403-page")
	}
}

/* Update Member Group*/
func UpdateMemberGroup(c *gin.Context) {

	var url string

	//get values from html from data
	id, _ := strconv.Atoi(c.PostForm("membergroup_id"))
	pageno := c.PostForm("memgrbpageno")
	if pageno != "" {
		url = "/membersgroup?page=" + pageno
	} else {
		url = "/membersgroup/"
	}
	userid := c.GetInt("userid")

	MemberGroup := mem.MemberGroupCreationUpdation{
		Name:        c.PostForm("membergroup_name"),
		Description: c.PostForm("membergroup_desc"),
		ModifiedBy:  userid,
	}

	permisison, _ := NewAuth.IsGranted("Members Group", auth.Update)
	if permisison {

		err := MemberConfig.UpdateMemberGroup(MemberGroup, id) //update member group
		if err != nil {
			ErrorLog.Printf("membergroup update error: %s", err)
		}
		if strings.Contains(fmt.Sprint(err), "given some values is empty") {
			c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
			c.Redirect(301, url)
			return
		}

		if err != nil {
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
			c.Redirect(301, url)
			return
		}
		c.SetCookie("get-toast", "Member Group Updated Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, url)

	} else {
		c.Redirect(301, "/403-page")
	}

}

/* Delete Member Group*/
func DeleteMemberGroup(c *gin.Context) {

	var url string

	//get values from html from data
	MemberGroupId, _ := strconv.Atoi(c.Query("id"))
	pageno := c.Query("page")
	if pageno != "" {
		url = "/membersgroup?page=" + pageno
	} else {
		url = "/membersgroup/"
	}
	userid := c.GetInt("userid")

	permisison, _ := NewAuth.IsGranted("Members Group", auth.Delete)
	if permisison {

		err := MemberConfig.DeleteMemberGroup(MemberGroupId, userid) //delete member group
		if err != nil {
			ErrorLog.Printf("membergroup delete error: %s", err)
		}

		if err != nil {
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
			c.Redirect(301, url)
			return
		}

		c.SetCookie("get-toast", "Member Group Deleted Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, url)

	} else {

		c.Redirect(301, "/403-page")
	}

}

/*Update Member Group status*/
func MemberIsActive(c *gin.Context) {

	//get values from html from data
	id, _ := strconv.Atoi(c.Request.PostFormValue("id"))
	val, _ := strconv.Atoi(c.Request.PostFormValue("isactive"))
	userid := c.GetInt("userid")

	permisison, _ := NewAuth.IsGranted("Members Group", auth.Update)
	if permisison {

		flg, err := MemberConfig.MemberGroupIsActive(id, val, userid) //update member group status
		if err != nil {
			ErrorLog.Printf("membergroup status update error: %s", err)
			json.NewEncoder(c.Writer).Encode(flg)

		} else {

			json.NewEncoder(c.Writer).Encode(flg)
		}

	} else {

		c.Redirect(301, "/403-page")
	}

}

func MemberGroupPopup(c *gin.Context) {

	auth := member.Memberauth{Authority: &AUTH}

	var id, _ = strconv.Atoi(c.PostForm("id"))

	membergroup, err := auth.MemberDeletePopup(id)

	if err != nil {

		log.Println(err)

		json.NewEncoder(c.Writer).Encode("false")

	} else {

		json.NewEncoder(c.Writer).Encode(membergroup)
	}

}

func CheckNameInMemberGroup(c *gin.Context) {

	id, _ := strconv.Atoi(c.PostForm("membergroup_id"))

	name := c.PostForm("membergroup_name")

	permisison, perr := NewAuth.IsGranted("Members Group", auth.Read)
	if perr != nil {
		ErrorLog.Printf("multi group delete authorization error :%s", perr)
	}

	if permisison {
		flg, err := MemberConfig.CheckNameInMemberGroup(id, name)
		if err != nil {
			ErrorLog.Printf("check name in member error :%s", perr)
			json.NewEncoder(c.Writer).Encode(false)
			return
		}
		json.NewEncoder(c.Writer).Encode(flg)

	} else {

		c.Redirect(301, "/403-page")
	}

}
func MultiSelectDeleteMembergroup(c *gin.Context) {

	var (
		memberdata []map[string]string
		Memberids  []int
		url        string
	)

	if err := json.Unmarshal([]byte(c.Request.PostFormValue("memberids")), &memberdata); err != nil {
		fmt.Println(err)
	}

	for _, val := range memberdata {
		memberid, _ := strconv.Atoi(val["memberid"])
		Memberids = append(Memberids, memberid)
	}

	pageno := c.PostForm("page")

	if pageno != "" {
		url = "/membersgroup?page=" + pageno
	} else {
		url = "/membersgroup/"
	}

	userid := c.GetInt("userid")

	permisison, perr := NewAuth.IsGranted("Members Group", auth.Delete)
	if perr != nil {
		ErrorLog.Printf("multi group delete authorization error :%s", perr)
	}

	if permisison {

		_, err := MemberConfig.MultiSelectedMemberDeletegroup(Memberids, userid)
		if err != nil {
			ErrorLog.Printf("multi group delete error :%s", perr)
			c.JSON(200, gin.H{"value": false})
			return
		}

		c.JSON(200, gin.H{"value": true, "url": url})

	} else {

		c.Redirect(301, "/403-page")
	}

}

func MultiSelectMembersgroupStatus(c *gin.Context) {

	var (
		memberdata []map[string]string
		memberids  []int
		statusint  int
		status     int
	)

	if err := json.Unmarshal([]byte(c.Request.PostFormValue("memberids")), &memberdata); err != nil {
		fmt.Println(err)
	}

	for _, val := range memberdata {

		memberid, _ := strconv.Atoi(val["memberid"])
		statusint, _ = strconv.Atoi(val["status"])

		if statusint == 0 {
			status = 1
		} else if statusint == 1 {
			status = 0

		}

		memberids = append(memberids, memberid)

	}

	pageno := c.PostForm("page")

	var url string
	if pageno != "" {
		url = "/membersgroup?page=" + pageno
	} else {
		url = "/membersgroup/"
	}

	userid := c.GetInt("userid")

	permisison, perr := NewAuth.IsGranted("Members Group", auth.Update)
	if perr != nil {
		ErrorLog.Printf("multi group status authorization error :%s", perr)
	}

	if permisison {

		_, err := MemberConfig.MultiSelectMembersgroupStatus(memberids, status, userid)

		if err != nil {
			ErrorLog.Printf("multi group delete authorization error :%s", err)
			c.JSON(200, gin.H{"value": false})
			return
		}

		c.JSON(200, gin.H{"value": true, "status": status, "url": url})

	} else {

		c.Redirect(301, "/403-page")
	}

}
