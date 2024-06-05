package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	mem "github.com/spurtcms/member"
	"github.com/spurtcms/pkgcore/member"
	csrf "github.com/utrack/gin-csrf"
)

func MemberList(c *gin.Context) {

	var (
		limt   int
		offset int
		filter mem.Filter
	)

	//get data from html url query
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

	flag := false

	var memberlists []mem.Tblmember

	permisison, perr := NewAuth.IsGranted("Member", auth.Read)
	if perr != nil {
		ErrorLog.Printf("MemberList authorization error: %s", perr)
	}

	if permisison {

		memberlist, Total_users, err := MemberConfig.ListMembers(offset, limt, mem.Filter(filter), flag)
		if err != nil {
			ErrorLog.Printf("Get data from member list error :%s", err)
		}

		membergroup, _ := MemberConfig.GetGroupData()

		for _, memberlistvalue := range memberlist {

			memberprof, _ := MemberConfig.GetMemberProfileByMemberId(memberlistvalue.Id)
			memberlistvalue.CreatedDate = memberlistvalue.CreatedOn.In(TZONE).Format(Datelayout)

			if !memberlistvalue.ModifiedOn.IsZero() {
				memberlistvalue.ModifiedDate = memberlistvalue.ModifiedOn.In(TZONE).Format(Datelayout)
			} else {
				memberlistvalue.ModifiedDate = memberlistvalue.CreatedOn.In(TZONE).Format(Datelayout)
			}

			memberlistvalue.ProfileImage = memberprof.CompanyName
			memberlistvalue.Token, _ = MemberConfig.GenerateMemberToken(memberlistvalue.Id, SecretKey)
			memberlists = append(memberlists, memberlistvalue)
		}

		var paginationendcount = len(memberlist) + offset
		paginationstartcount := offset + 1
		Previous, Next, PageCount, Page := Pagination(pageno, int(Total_users), limt)

		translate, _ := TranslateHandler(c)

		menu := NewMenuController(c)

		ModuleName, TabName, _ := ModuleRouteName(c)

		c.HTML(200, "members.html", gin.H{"csrf": csrf.GetToken(c), "Pagination": PaginationData{
			NextPage:     pageno + 1,
			PreviousPage: pageno - 1,
			TotalPages:   PageCount,
			TwoAfter:     pageno + 2,
			TwoBelow:     pageno - 2,
			ThreeAfter:   pageno + 3,
		}, "Menu": menu, "Member": memberlists, "Group": membergroup, "Count": Total_users, "Previous": Previous, "Next": Next, "Page": Page, "Limit": limt, "PageCount": PageCount, "CurrentPage": pageno, "Paginationstartcount": paginationstartcount, "Filters": filter, "Paginationendcount": paginationendcount, "title": ModuleName, "HeadTitle": translate.Memberss.Members, "translate": translate, "Filter": filter, "Membermenu": true, "membermenu": true, "Tabmenu": TabName, "Url": OwndeskUrl + CompanyProfilePath})

	} else {

		c.Redirect(301, "/403-page")

	}
}
func CreateMember(c *gin.Context) {

	userid := c.GetInt("userid")

	var MemberGroupId, _ = strconv.Atoi(c.PostForm("membergroupvalue"))
	IsActive, _ := strconv.Atoi(c.PostForm("mem_activestat"))
	imagedata := c.PostForm("crop_data")
	var imageName, imagePath, profileslug, profilename string

	if imagedata != "" {
		imageName, imagePath, _ = ConvertBase64(imagedata, "storage/member")
	}

	frname := strings.ToLower(c.PostForm("mem_name"))
	lrname := strings.ToLower(c.PostForm("mem_lname"))

	if c.PostForm("mem_lname") != "" {
		slug := frname + "-" + lrname
		profileslug = strings.ReplaceAll(slug, " ", "-")
	} else {
		slug := frname
		profileslug = strings.ReplaceAll(slug, " ", "-")
	}

	if c.PostForm("mem_lname") != "" {
		profilename = frname + " " + lrname
	} else {
		profilename = frname
	}

	Member := mem.MemberCreationUpdation{

		FirstName:        c.PostForm("mem_name"),
		LastName:         c.PostForm("mem_lname"),
		Email:            c.PostForm("mem_email"),
		MobileNo:         c.PostForm("mem_mobile"),
		ProfileImage:     imageName,
		ProfileImagePath: imagePath,
		GroupId:          MemberGroupId,
		IsActive:         IsActive,
		Username:         c.PostForm("mem_usrname"),
		Password:         c.PostForm("mem_pass"),
		CreatedBy:        userid,
	}

	permisison, perr := NewAuth.IsGranted("Member", auth.Create)
	if perr != nil {
		ErrorLog.Printf("rolescreate authorization error: %s", perr)
	}

	if permisison {

		memberdata, err := MemberConfig.CreateMember(Member)

		Memberprofile := mem.MemberprofilecreationUpdation{
			MemberId:    memberdata.Id,
			ProfileSlug: profileslug,
			ProfileName: profilename,
			CreatedBy:   userid,
		}

		merr := MemberConfig.CreateMemberProfile(Memberprofile)
		if merr != nil {
			ErrorLog.Printf("memberprofile create error: %s", merr)
		}

		if strings.Contains(fmt.Sprint(err), "given some values is empty") {
			c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
			c.Redirect(301, "/member/")
			return
		}

		if err != nil {
			ErrorLog.Printf("memberprofile create error: %s", merr)
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
			c.Redirect(301, "/member/")
			return
		}

		c.SetCookie("get-toast", "Member Created Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, "/member/")

		if IsActive == 1 {

			var web_url = os.Getenv("WEB_URL")
			var url_prefix = os.Getenv("BASE_URL")
			fname := c.PostForm("mem_name")
			uname := c.PostForm("mem_email")
			pas := c.PostForm("mem_pass")
			email := c.PostForm("mem_email")

			data := map[string]interface{}{

				"fname":         fname,
				"email":         uname,
				"Pass":          pas,
				"login_url":     web_url,
				"admin_logo":    url_prefix + "public/img/spurtcms.png",
				"fb_logo":       url_prefix + "public/img/facebook.png",
				"linkedin_logo": url_prefix + "public/img/linkedin.png",
				"twitter_logo":  url_prefix + "public/img/twitter.png",
			}

			var wg sync.WaitGroup

			wg.Add(1)

			Chan := make(chan string, 1)

			go MemberCreateEmail(Chan, &wg, data, email, "Createmember")

			close(Chan)
		}

	} else {

		c.Redirect(301, "/403-page")

	}

}

// edit member
func EditMember(c *gin.Context) {

	permisison, perr := NewAuth.IsGranted("Member", auth.Update)
	if perr != nil {
		ErrorLog.Printf("edit member authorization error: %s", perr)
	}

	if permisison {

		pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

		var id, _ = strconv.Atoi(c.Param("id"))

		group, _ := MemberConfig.GetGroupData()
		member, _ := MemberConfig.GetMemberDetails(id)
		memberprof, _ := MemberConfig.GetMemberProfileByMemberId(id)

		var firstn = strings.ToUpper(member.FirstName[:1])
		var lastn, NameString string
		if member.LastName != "" {
			lastn = strings.ToUpper(member.LastName[:1])
		}

		member.NameString = firstn + lastn
		translate, _ := TranslateHandler(c)

		menu := NewMenuController(c)

		if memberprof.CompanyName != "" {
			NameString = memberprof.CompanyName[:1]
		}

		ModuleName, TabName, _ := ModuleRouteName(c)

		c.HTML(200, "memberprofile.html", gin.H{"csrf": csrf.GetToken(c), "Menu": menu, "Member": member, "Group": group, "title": ModuleName, "HeadTitle": translate.Memberss.Members, "translate": translate, "Membermenu": true, "membermenu": true, "MemberProfile": memberprof, "NameString": NameString, "Tabmenu": TabName, "CurrentPage": pageno})

	} else {

		c.Redirect(301, "/403-page")

	}
}

// update member
func UpdateMember(c *gin.Context) {

	userid := c.GetInt("userid")
	member_id, _ := strconv.Atoi(c.PostForm("mem_id"))
	pageno := c.PostForm("memgrbpageno")

	var url string
	if pageno != "" {
		url = "/member?page=" + pageno
	} else {
		url = "/member/"
	}

	var MemberGroupId, _ = strconv.Atoi(c.PostForm("membergroupvalue"))
	IsActive, _ := strconv.Atoi(c.PostForm("mem_activestat"))
	mailstatus, _ := strconv.Atoi(c.PostForm("mem_emailactive"))
	imagedata := c.PostForm("crop_data")

	var imageName, imagePath string
	if imagedata != "" {
		imageName, imagePath, _ = ConvertBase64(imagedata, "storage/member")
	}

	companylogo := c.PostForm("crop_data2")
	var companystoragepath string
	if companylogo != "" {
		_, companystoragepath, _ = ConvertBase64(companylogo, "storage/member/company")
	}

	profileid, _ := strconv.Atoi(c.PostForm("profileid"))
	claimstatus, _ := strconv.Atoi(c.PostForm("com_activestat"))

	Member := mem.MemberCreationUpdation{
		FirstName:        c.PostForm("mem_name"),
		LastName:         c.PostForm("mem_lname"),
		Email:            c.PostForm("mem_email"),
		MobileNo:         c.PostForm("mem_mobile"),
		ProfileImage:     imageName,
		ProfileImagePath: imagePath,
		GroupId:          MemberGroupId,
		IsActive:         IsActive,
		ModifiedBy:       userid,
		Username:         c.PostForm("mem_usrname"),
		Password:         c.PostForm("mem_pass"),
	}

	Memberprofile := mem.MemberprofilecreationUpdation{
		MemberId:        member_id,
		CompanyName:     c.PostForm("companyname"),
		CompanyLocation: c.PostForm("companylocation"),
		ProfileName:     c.PostForm("profilename"),
		ProfileSlug:     c.PostForm("profilepage"),
		About:           c.PostForm("aboutcompany"),
		ProfileId:       profileid,
		CompanyLogo:     companystoragepath,
		LinkedIn:        c.PostForm("linkedin"),
		Website:         c.PostForm("website"),
		Twitter:         c.PostForm("twitter"),
		ClaimStatus:     claimstatus,
		SeoTitle:        c.PostForm("metaTitle"),
		SeoKeyword:      c.PostForm("metaKeyword"),
		SeoDescription:  c.PostForm("metaDescription"),
		ModifiedBy:      userid,
	}

	permisison, perr := NewAuth.IsGranted("Member", auth.Update)
	if perr != nil {
		ErrorLog.Printf(" authorization error: %s", perr)
	}

	if permisison {

		err := MemberConfig.UpdateMember(Member, member_id)
		if err != nil {
			log.Println(err)
		}

		merr := MemberConfig.UpdateMemberProfile(Memberprofile)
		if merr != nil {
			log.Println(merr)
		}

		if mailstatus == 1 && claimstatus == 1 {

			email := c.PostForm("mem_email")
			var url_prefix = os.Getenv("BASE_URL")

			data := map[string]interface{}{
				"client_name":    c.PostForm("mem_name"),
				"company_name":   c.PostForm("companyname"),
				"logo":           url_prefix + "public/img/email-icons/logo.png",
				"fb_logo":        url_prefix + "public/img/email-icons/facebook.png",
				"linkedin_logo":  url_prefix + "public/img/email-icons/linkedin.png",
				"spacex_logo":    url_prefix + "public/img/email-icons/x.png",
				"youtube_logo":   url_prefix + "public/img/email-icons/youtube.png",
				"instagram_logo": url_prefix + "public/img/email-icons/instagram.png",
			}

			var wg sync.WaitGroup

			wg.Add(1)

			Chan := make(chan string, 1)

			go OwndeskmemberEmail(Chan, &wg, data, email, "Owndesk")

			close(Chan)

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

		c.SetCookie("get-toast", "Member Updated Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, url)

	} else {

		c.Redirect(301, "/403-page")

	}
}

// delete member
func DeleteMember(c *gin.Context) {

	var id, _ = strconv.Atoi(c.Query("id"))
	pageno := c.Query("page")
	var url string

	if pageno != "" {
		url = "/member?page=" + pageno
	} else {
		url = "/member/"
	}

	userid := c.GetInt("userid")

	permisison, perr := NewAuth.IsGranted("Member", auth.Delete)
	if perr != nil {
		ErrorLog.Printf("rolescreate authorization error: %s", perr)
	}

	if permisison {

		err := MemberConfig.DeleteMember(id, userid)

		if err != nil {
			ErrorLog.Printf("member delete error: %s", perr)
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			c.Redirect(301, url)
			return
		}

		c.SetCookie("get-toast", "Member Deleted Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, url)

	} else {

		c.Redirect(301, "/403-page")

	}
}

func MemberPopup(c *gin.Context) {

	auth := member.Memberauth{Authority: &AUTH}

	var member member.TblMember

	var id, _ = strconv.Atoi(c.PostForm("id"))

	member, err := auth.MemberDeletePopup(id)

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, member)
}

// check email
func CheckEmailInMember(c *gin.Context) {

	userid, _ := strconv.Atoi(c.PostForm("id"))

	email := c.PostForm("email")

	permisison, perr := NewAuth.IsGranted("Member", auth.Read)
	if perr != nil {
		ErrorLog.Printf("member check email authorization error: %s", perr)
	}

	if permisison {

		flg, err := MemberConfig.CheckEmailInMember(userid, email)

		if err != nil {
			ErrorLog.Printf("member check email error: %s", err)
			json.NewEncoder(c.Writer).Encode(false)
			return
		}

		json.NewEncoder(c.Writer).Encode(flg)

	} else {

		c.Redirect(301, "/403-page")

	}
}

// check number in member
func CheckNumberInMember(c *gin.Context) {

	userid, _ := strconv.Atoi(c.PostForm("id"))

	number := c.PostForm("number")

	permisison, perr := NewAuth.IsGranted("Member", auth.Read)
	if perr != nil {
		ErrorLog.Printf("member check number authorization error: %s", perr)
	}

	if permisison {
		flg, err := MemberConfig.CheckNumberInMember(userid, number)
		if err != nil {
			ErrorLog.Printf("member check number error: %s", perr)
			json.NewEncoder(c.Writer).Encode(false)
			return
		}
		json.NewEncoder(c.Writer).Encode(flg)

	} else {

		c.Redirect(301, "/403-page")

	}
}

// check profile name in member
func CheckProfileNameInMember(c *gin.Context) {

	userid, _ := strconv.Atoi(c.PostForm("id"))

	Profilename := c.PostForm("name")

	flg := false

	if Profilename != "" {

		flg2, err := MemberConfig.CheckProfileSlugInMember(userid, Profilename)

		flg = flg2
		if err != nil {
			ErrorLog.Printf("member profile name error: %s", err)
			json.NewEncoder(c.Writer).Encode(false)
			return
		}

		json.NewEncoder(c.Writer).Encode(flg)

		return
	}

	json.NewEncoder(c.Writer).Encode(flg)

}

func CheckNameInMember(c *gin.Context) {

	userid, _ := strconv.Atoi(c.PostForm("id"))

	name := c.PostForm("name")

	permisison, perr := NewAuth.IsGranted("Member", auth.Read)
	if perr != nil {
		ErrorLog.Printf("member check name authorization error: %s", perr)
	}

	if permisison {

		flg, err := MemberConfig.CheckNameInMember(userid, name)

		if err != nil {
			ErrorLog.Printf("member checkname error: %s", err)
			json.NewEncoder(c.Writer).Encode(false)
			return
		}

		json.NewEncoder(c.Writer).Encode(flg)

	}
	c.Redirect(301, "/403-page")

}

func GenerateMemberToken(c *gin.Context) {

	memberid, err := strconv.Atoi(c.PostForm("memberid"))

	if err != nil {

		c.AbortWithStatus(422)

		return
	}

	secretKey := os.Getenv("JWT_SECRET")

	permisison, perr := NewAuth.IsGranted("Member", auth.Create)
	if perr != nil {
		ErrorLog.Printf("member generate token authorization error: %s", perr)
	}

	if permisison {
		token, err := MemberConfig.GenerateMemberToken(memberid, secretKey)

		if err != nil {

			c.AbortWithStatus(400)

			return
		}

		url := os.Getenv("OWNDESK_URL")

		c.JSON(200, gin.H{"token": token, "url": url})
	}
	c.Redirect(301, "/403-page")
}

/*permission Member status*/
func MemberStatus(c *gin.Context) {

	userid := c.GetInt("userid")

	id, _ := strconv.Atoi(c.Request.PostFormValue("id"))
	val, _ := strconv.Atoi(c.Request.PostFormValue("isactive"))

	permisison, perr := NewAuth.IsGranted("Member", auth.Update)
	if perr != nil {
		ErrorLog.Printf("member status authorization error: %s", perr)
	}

	if permisison {

		flg, err := MemberConfig.MemberStatus(id, val, userid)

		if err != nil {
			ErrorLog.Printf("member status change error: %s", perr)
			json.NewEncoder(c.Writer).Encode(flg)

		} else {
			json.NewEncoder(c.Writer).Encode(flg)
		}
	}
	c.Redirect(301, "/403-page")
}

func MultiSelectDeleteMember(c *gin.Context) {

	var memberdata []map[string]string

	if err := json.Unmarshal([]byte(c.Request.PostFormValue("memberids")), &memberdata); err != nil {
		fmt.Println(err)
	}

	var Memberids []int

	for _, val := range memberdata {
		memberid, _ := strconv.Atoi(val["memberid"])
		Memberids = append(Memberids, memberid)
	}

	userid := c.GetInt("userid")
	pageno := c.PostForm("page")

	var url string
	if pageno != "" {
		url = "/member?page=" + pageno
	} else {
		url = "/member/"
	}

	permisison, perr := NewAuth.IsGranted("Member", auth.Delete)
	if perr != nil {
		ErrorLog.Printf("member multiple delete authorization error: %s", perr)
	}

	if permisison {

		_, err := MemberConfig.MultiSelectedMemberDelete(Memberids, userid)
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{"value": false})
			return
		}

		c.JSON(200, gin.H{"value": true, "url": url})
	}
	c.Redirect(301, "/403-page")
}

// multi select member status
func MultiSelectMembersStatus(c *gin.Context) {

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
		url = "/member?page=" + pageno
	} else {
		url = "/member/"
	}
	userid := c.GetInt("userid")
	permisison, perr := NewAuth.IsGranted("Member", auth.Update)
	if perr != nil {
		ErrorLog.Printf("member multiple status authorization error: %s", perr)
	}

	if permisison {
		_, err := MemberConfig.MultiSelectMembersStatus(memberids, status, userid)
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{"value": false})
			return
		}

		c.JSON(200, gin.H{"value": true, "status": status, "url": url})
	}
	c.Redirect(301, "/403-page")
}

func CheckProfileSlugInMember(c *gin.Context) {

	//get values from html from data
	userid, _ := strconv.Atoi(c.PostForm("id"))
	ProfileSlug := c.PostForm("name")
	flg := false

	permisison, perr := NewAuth.IsGranted("Member", auth.Read)
	if perr != nil {
		ErrorLog.Printf("member check profile slug authorization error: %s", perr)
	}

	if permisison {

		if ProfileSlug != "" {
			flg2, err := MemberConfig.CheckProfileSlugInMember(userid, ProfileSlug) //check memberprofile slug
			flg = flg2
			if err != nil {
				ErrorLog.Printf("memberprofileslug error: %s", err)
				json.NewEncoder(c.Writer).Encode(false)
				return
			}
			json.NewEncoder(c.Writer).Encode(flg)
			return
		}
		json.NewEncoder(c.Writer).Encode(flg)
		return
	}
	c.Redirect(301, "/403-page")
}
