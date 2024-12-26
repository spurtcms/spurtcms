package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"spurt-cms/models"
	storagecontroller "spurt-cms/storage-controller"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	mem "github.com/spurtcms/member"
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

	var memberlists []mem.Tblmember

	permisison, perr := NewAuth.IsGranted("Member", auth.Read, TenantId)

	if perr != nil {
		ErrorLog.Printf("MemberList authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("MemberList authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		MemberConfig.DataAccess = c.GetInt("dataaccess")
		MemberConfig.UserId = c.GetInt("userid")

		memberlist, Total_users, err := MemberConfig.ListMembers(offset, limt, mem.Filter(filter), flag, TenantId)
		if err != nil {
			ErrorLog.Printf("Get data from member list error :%s", err)
		}

		membergroup, _ := MemberConfig.GetGroupData(TenantId)

		for _, memberlistvalue := range memberlist {

			memberprof, _ := MemberConfig.GetMemberProfileByMemberId(memberlistvalue.Id, TenantId)
			memberlistvalue.CreatedDate = memberlistvalue.CreatedOn.In(TZONE).Format(Datelayout)

			if !memberlistvalue.ModifiedOn.IsZero() {
				memberlistvalue.ModifiedDate = memberlistvalue.ModifiedOn.In(TZONE).Format(Datelayout)
			} else {
				memberlistvalue.ModifiedDate = memberlistvalue.CreatedOn.In(TZONE).Format(Datelayout)
			}

			if memberlistvalue.ProfileImagePath != "" {
				if memberlistvalue.StorageType == "local" {
					memberlistvalue.ProfileImagePath = "/" + memberlistvalue.ProfileImagePath
				} else if memberlistvalue.StorageType == "aws" {
					memberlistvalue.ProfileImagePath = "/image-resize?name=" + memberlistvalue.ProfileImagePath
				}
			}
			memberlistvalue.Claimstatus = memberprof.ClaimStatus
			memberlistvalue.ProfileImage = memberprof.CompanyName
			memberlistvalue.Token, _ = NewAuth.GenerateMemberToken(memberlistvalue.Id, "admin", SecretKey, TenantId)
			memberlists = append(memberlists, memberlistvalue)
		}

		var paginationendcount = len(memberlist) + offset
		paginationstartcount := offset + 1
		Previous, Next, PageCount, Page := Pagination(pageno, int(Total_users), limt)

		translate, _ := TranslateHandler(c)

		menu := NewMenuController(c)

		ModuleName, TabName, _ := ModuleRouteName(c)

		storagetype, err := GetSelectedType()
		if err != nil {
			fmt.Printf("member list getting storagetype error: %s", err)
		}
		uper, _ := NewAuth.IsGranted("Member", auth.Update, TenantId)

		dper, _ := NewAuth.IsGranted("Member", auth.Delete, TenantId)

		fmt.Println("list", memberlist)

		c.HTML(200, "members.html", gin.H{"csrf": csrf.GetToken(c), "Pagination": PaginationData{
			NextPage:     pageno + 1,
			PreviousPage: pageno - 1,
			TotalPages:   PageCount,
			TwoAfter:     pageno + 2,
			TwoBelow:     pageno - 2,
			ThreeAfter:   pageno + 3,
		}, "Menu": menu, "Member": memberlists, "Searchtrue": filterflag, "Group": membergroup, "Count": Total_users, "Previous": Previous, "Next": Next, "Page": Page, "Limit": limt, "PageCount": PageCount, "CurrentPage": pageno, "Paginationstartcount": paginationstartcount, "Filters": filter, "Paginationendcount": paginationendcount, "title": ModuleName, "linktitle": ModuleName, "HeadTitle": translate.Memberss.Members, "translate": translate, "Filter": filter, "Membermenu": true, "membermenu": true, "Tabmenu": TabName, "Url": OwndeskUrl + CompanyProfilePath, "StorageType": storagetype.SelectedType, "permission": uper, "dpermission": dper})

	}
}

func CreateMember(c *gin.Context) {

	userid := c.GetInt("userid")

	var MemberGroupId, _ = strconv.Atoi(c.PostForm("membergroupvalue"))
	IsActive, _ := strconv.Atoi(c.PostForm("mem_activestat"))
	imagedata := c.PostForm("memcrop_base64")
	var imageName, imagePath, profileslug, profilename string

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
				c.Redirect(301, "/member/")
				return
			}
		}
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
		StorageType:      storagetype.SelectedType,
		TenantId:         TenantId,
	}

	permisison, perr := NewAuth.IsGranted("Member", auth.Create, TenantId)
	if perr != nil {
		ErrorLog.Printf("Member create authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Member authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		memberdata, err := MemberConfig.CreateMember(Member)

		Memberprofile := mem.MemberprofilecreationUpdation{
			MemberId:    memberdata.Id,
			ProfileSlug: profileslug,
			ProfileName: profilename,
			CreatedBy:   userid,
			TenantId:    TenantId,
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

		// var email models.TblEmailTemplate
		// err = models.GetTemplates(&email, "Createmember", TenantId)
		// if err != nil {
		// 	ErrorLog.Printf("emailtemplate get error: %s", merr)
		// }

		if IsActive == 1 {

			var web_url = os.Getenv("BASE_URL")
			var url_prefix = os.Getenv("BASE_URL")
			fname := c.PostForm("mem_name")
			uname := c.PostForm("mem_email")
			pas := c.PostForm("mem_pass")
			email := c.PostForm("mem_email")
			var owndeskmail = "hello@owndesk.in"
			linkedin := os.Getenv("LINKEDIN")
			facebook := os.Getenv("FACEBOOK")
			twitter := os.Getenv("TWITTER")
			youtube := os.Getenv("YOUTUBE")
			insta := os.Getenv("INSTAGRAM")

			data := map[string]interface{}{

				"fname":         fname,
				"email":         uname,
				"Pass":          pas,
				"login_url":     web_url,
				"owndesk_logo":  url_prefix + "public/img/email-icons/logo.png",
				"admin_logo":    url_prefix + "public/img/SpurtCMSlogo.png",
				"fb_logo":       url_prefix + "public/img/email-icons/facebook.png",
				"linkedin_logo": url_prefix + "public/img/email-icons/linkedin.png",
				"twitter_logo":  url_prefix + "public/img/email-icons/x.png",
				"youtube_logo":  url_prefix + "public/img/email-icons/youtube.png",
				"insta_log":     url_prefix + "public/img/email-icons/instagram.png",
				"contactemail":  owndeskmail,
				"facebook":      facebook,
				"instagram":     insta,
				"youtube":       youtube,
				"linkedin":      linkedin,
				"twitter":       twitter,
			}

			// var wg sync.WaitGroup

			// wg.Add(1)

			Chan := make(chan string, 1)

			// go MemberActivationEmail(Chan, &wg, data, email, "memberactivation")

			var wg1 sync.WaitGroup

			wg1.Add(1)

			go MemberCreateEmail(Chan, &wg1, data, email, TenantId, "Createmember")

			close(Chan)
		}

	}

}

// edit member
func EditMember(c *gin.Context) {

	var (
		lastn      string
		NameString string
		firstn     string
	)
	permisison, perr := NewAuth.IsGranted("Member", auth.Update, TenantId)
	if perr != nil {
		ErrorLog.Printf("edit member authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Member authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

		var id, _ = strconv.Atoi(c.Param("id"))

		group, _ := MemberConfig.GetGroupData(TenantId)
		member, _ := MemberConfig.GetMemberDetails(id, TenantId)
		memberprof, _ := MemberConfig.GetMemberProfileByMemberId(id, TenantId)

		if member.FirstName != "" {
			lastn = strings.ToUpper(member.FirstName[:1])
		}
		if member.LastName != "" {
			lastn += strings.ToUpper(member.LastName[:1])
		}

		member.NameString = firstn + lastn
		translate, _ := TranslateHandler(c)

		menu := NewMenuController(c)

		if memberprof.CompanyName != "" {
			NameString = memberprof.CompanyName[:1]
		}

		if member.ProfileImagePath != "" {
			if member.StorageType == "local" {
				member.ProfileImagePath = "/" + member.ProfileImagePath
			} else if member.StorageType == "aws" {
				member.ProfileImagePath = "/image-resize?name=" + member.ProfileImagePath
			}
		}

		if memberprof.CompanyLogo != "" {
			if memberprof.StorageType == "local" {
				memberprof.CompanyLogo = "/" + memberprof.CompanyLogo
			} else if memberprof.StorageType == "aws" {
				memberprof.CompanyLogo = "/image-resize?name=" + memberprof.CompanyLogo
			}
		}

		ModuleName, TabName, _ := ModuleRouteName(c)

		// fmt.Println(ModuleName, "modulename")

		storagetype, err := GetSelectedType()
		if err != nil {
			fmt.Printf("member list getting storagetype error: %s", err)
		}

		c.HTML(200, "memberprofile.html", gin.H{"csrf": csrf.GetToken(c), "Menu": menu, "Member": member, "Group": group, "title": ModuleName, "HeadTitle": translate.Memberss.Members, "translate": translate, "Membermenu": true, "membermenu": true, "MemberProfile": memberprof, "NameString": NameString, "Tabmenu": TabName, "CurrentPage": pageno, "StorageType": storagetype.SelectedType})

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
	fmt.Println("mem grp id", MemberGroupId)
	IsActive, _ := strconv.Atoi(c.PostForm("mem_activestat"))
	// mailstatus, _ := strconv.Atoi(c.PostForm("mem_emailactive"))
	imagedata := c.PostForm("crop_data")

	var imageName, imagePath string

	companylogo := c.PostForm("crop_data2")
	var companystoragepath string

	storagetype, err := GetSelectedType()
	if err != nil {
		ErrorLog.Printf("error get storage type error: %s", err)
	}

	if storagetype.SelectedType == "local" {

		if imagedata != "" {
			imageName, imagePath, err = ConvertBase64(imagedata, strings.TrimPrefix(storagetype.Local+"/member", "/"))
			if err != nil {
				ErrorLog.Printf("update member convertbase64 error: %s", err)
			}
		}

		if companylogo != "" {

			_, companystoragepath, _ = ConvertBase64(companylogo, strings.TrimPrefix(storagetype.Local+"/member/company", "/"))

		}

	} else if storagetype.SelectedType == "aws" {
		tenantDetails, err := NewTeam.GetTenantDetails(TenantId)
		if err != nil {
			fmt.Println(err)
		}

		if imagedata != "" {

			var (
				imageByte []byte
				err       error
			)

			imageName, imagePath, imageByte, err = ConvertBase64toByte(imagedata, "member")
			if err != nil {
				ErrorLog.Printf("convert base 64 to byte error : %s", err)
			}

			imagePath = tenantDetails.S3FolderName + imagePath

			uerr := storagecontroller.UploadCropImageS3(imageName, imagePath, imageByte)
			if uerr != nil {

				c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
				c.Redirect(301, "/member/")
				return

			}
		}

		if companylogo != "" {
			tenantDetails, err := NewTeam.GetTenantDetails(TenantId)
			if err != nil {
				fmt.Println(err)
			}

			var imageByte []byte

			_, companystoragepath, imageByte, err = ConvertBase64toByte(companylogo, "member/company")
			if err != nil {
				ErrorLog.Printf("convert base 64 to byte error : %s", err)
			}

			companystoragepath = tenantDetails.S3FolderName + companystoragepath

			uerr := storagecontroller.UploadCropImageS3(imageName, companystoragepath, imageByte)
			if uerr != nil {

				c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
				c.Redirect(301, "/member/")
				return

			}

		}

	}

	profileid, _ := strconv.Atoi(c.PostForm("profileid"))
	// claimstatus, _ := strconv.Atoi(c.PostForm("com_activestat"))

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
		StorageType:      storagetype.SelectedType,
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
		SeoTitle:        c.PostForm("metaTitle"),
		SeoKeyword:      c.PostForm("metaKeyword"),
		SeoDescription:  c.PostForm("metaDescription"),
		ModifiedBy:      userid,
		StorageType:     storagetype.SelectedType,
	}

	permisison, perr := NewAuth.IsGranted("Member", auth.Update, TenantId)
	if perr != nil {
		ErrorLog.Printf("Update Member authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Member authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	err = MemberConfig.UpdateMember(Member, member_id, TenantId)
	if err != nil {
		ErrorLog.Printf("update member error: %s", perr)
	}

	merr := MemberConfig.UpdateMemberProfile(Memberprofile, TenantId)
	if merr != nil {
		ErrorLog.Printf("update member profile error: %s", perr)
	}

	// var email models.TblEmailTemplate
	// err = models.GetTemplates(&email, "Owndesk", TenantId)
	// if err != nil {
	// 	ErrorLog.Printf("emailtemplate get error: %s", err)
	// }

	// if mailstatus == 1 && claimstatus == 1 && email.IsActive == 1 {

	// 	email := c.PostForm("mem_email")
	// 	var url_prefix = os.Getenv("BASE_URL")

	// 	data := map[string]interface{}{
	// 		"client_name":    c.PostForm("mem_name"),
	// 		"company_name":   c.PostForm("companyname"),
	// 		"logo":           url_prefix + "public/img/email-icons/logo.png",
	// 		"fb_logo":        url_prefix + "public/img/email-icons/facebook.png",
	// 		"linkedin_logo":  url_prefix + "public/img/email-icons/linkedin.png",
	// 		"spacex_logo":    url_prefix + "public/img/email-icons/x.png",
	// 		"youtube_logo":   url_prefix + "public/img/email-icons/youtube.png",
	// 		"instagram_logo": url_prefix + "public/img/email-icons/instagram.png",
	// 	}

	// 	var wg sync.WaitGroup

	// 	wg.Add(1)

	// 	Chan := make(chan string, 1)

	// 	go OwndeskmemberEmail(Chan, &wg, data, email, "Owndesk")

	// 	close(Chan)

	// }

	if strings.Contains(fmt.Sprint(err), "given some values is empty") {
		ErrorLog.Printf("member update mandatory fields error: %s", err)
		c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
		c.Redirect(301, url)
		return
	}

	if err != nil {
		fmt.Println("member errorsss")
		ErrorLog.Printf("member update error: %s", err)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
		c.Redirect(301, url)
		return
	}

	c.SetCookie("get-toast", "Member Updated Successfully", 3600, "", "", false, false)
	fmt.Println("member updayte")
	// c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, url)

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

	permission, perr := NewAuth.IsGranted("Member", auth.Delete, TenantId)
	if perr != nil {
		ErrorLog.Printf("rolescreate authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if !permission {
		ErrorLog.Printf("Member authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	err := MemberConfig.DeleteMember(id, userid, TenantId)
	if err != nil {
		ErrorLog.Printf("member delete error: %s", err)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.Redirect(301, url)
		return
	}

	_, totalRecords, _ := MemberConfig.ListMembers(0, 100, mem.Filter{}, false, TenantId)

	recordsPerPage := Limit
	totalPages := (int(totalRecords) + recordsPerPage - 1) / recordsPerPage

	currentPage := 1
	if pageno != "" {
		currentPage, _ = strconv.Atoi(pageno)
	}

	if currentPage > totalPages && totalPages > 0 {

		url = "/member?page=" + strconv.Itoa(totalPages)
	}

	c.SetCookie("get-toast", "Member Deleted Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)

	if totalRecords == 0 {
		url = "/member?page=1"
	}

	c.Redirect(301, url)
}

// check email
func CheckEmailInMember(c *gin.Context) {

	userid, _ := strconv.Atoi(c.PostForm("id"))

	email := c.PostForm("email")

	permisison, perr := NewAuth.IsGranted("Member", auth.Read, TenantId)
	if perr != nil {
		ErrorLog.Printf("member check email authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Member authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		flg, err := MemberConfig.CheckEmailInMember(userid, email, TenantId)

		if err != nil {
			ErrorLog.Printf("member check email error: %s", err)
			json.NewEncoder(c.Writer).Encode(false)
			return
		}

		json.NewEncoder(c.Writer).Encode(flg)

	}
}

// check number in member
func CheckNumberInMember(c *gin.Context) {

	userid, _ := strconv.Atoi(c.PostForm("id"))

	number := c.PostForm("number")

	permisison, perr := NewAuth.IsGranted("Member", auth.Read, TenantId)
	if perr != nil {
		ErrorLog.Printf("member check number authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Member authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		flg, err := MemberConfig.CheckNumberInMember(userid, number, TenantId)
		if err != nil {
			ErrorLog.Printf("member check number error: %s", err)
			json.NewEncoder(c.Writer).Encode(false)
			return
		}
		json.NewEncoder(c.Writer).Encode(flg)

	}
}

// check profile name in member
func CheckProfileNameInMember(c *gin.Context) {

	userid, _ := strconv.Atoi(c.PostForm("id"))

	Profilename := c.PostForm("name")

	flg := false

	if Profilename != "" {

		flg2, err := MemberConfig.CheckProfileSlugInMember(userid, Profilename, TenantId)

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

	permisison, perr := NewAuth.IsGranted("Member", auth.Read, TenantId)
	if perr != nil {
		ErrorLog.Printf("member check name authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Member authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		flg, err := MemberConfig.CheckNameInMember(userid, name, TenantId)

		if err != nil {
			ErrorLog.Printf("member checkname error: %s", err)
			json.NewEncoder(c.Writer).Encode(false)
			return
		}

		json.NewEncoder(c.Writer).Encode(flg)

	}

}

/*permission Member status*/
func MemberStatus(c *gin.Context) {

	userid := c.GetInt("userid")

	id, _ := strconv.Atoi(c.Request.PostFormValue("id"))
	val, _ := strconv.Atoi(c.Request.PostFormValue("isactive"))
	fmt.Println("datas", userid, id, val)

	permisison, perr := NewAuth.IsGranted("Member", auth.Update, TenantId)
	if perr != nil {
		ErrorLog.Printf("member status authorization error: %s", perr)
	}
	if !permisison {
		ErrorLog.Printf("Member authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		flg, err := MemberConfig.MemberStatus(id, val, userid, TenantId)

		if err != nil {
			ErrorLog.Printf("member status change error: %s", perr)
			json.NewEncoder(c.Writer).Encode(flg)

		} else {
			json.NewEncoder(c.Writer).Encode(flg)
		}
	}

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

	permisison, perr := NewAuth.IsGranted("Member", auth.Delete, TenantId)
	if perr != nil {
		ErrorLog.Printf("member multiple delete authorization error: %s", perr)
	}
	if !permisison {
		ErrorLog.Printf("Member authorization error: %s", perr)
		c.JSON(200, gin.H{"value": false, "url": "/403-page"})
		return
	}

	if permisison {

		_, err := MemberConfig.MultiSelectedMemberDelete(Memberids, userid, TenantId)
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{"value": false})
			return
		}
		_, totalRecords, _ := MemberConfig.ListMembers(0, 100, mem.Filter{}, false, TenantId)

		recordsPerPage := Limit
		totalPages := (int(totalRecords) + recordsPerPage - 1) / recordsPerPage

		fmt.Println(totalPages, "totalpagesss")

		currentPage := 1
		if pageno != "" {
			currentPage, _ = strconv.Atoi(pageno)
		}

		if currentPage > totalPages && totalPages > 0 {

			url = "/member?page=" + strconv.Itoa(totalPages)
		}

		if totalRecords == 0 {
			url = "/member?page=1"
		}
		c.JSON(200, gin.H{"value": true, "url": url})
	}

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
	permisison, perr := NewAuth.IsGranted("Member", auth.Update, TenantId)
	if perr != nil {
		ErrorLog.Printf("member multiple status authorization error: %s", perr)
	}
	if !permisison {
		ErrorLog.Printf("Member authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		_, err := MemberConfig.MultiSelectMembersStatus(memberids, status, userid, TenantId)
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{"value": false})
			return
		}

		c.JSON(200, gin.H{"value": true, "status": status, "url": url})
	}

}

func CheckProfileSlugInMember(c *gin.Context) {

	//get values from html from data
	userid, _ := strconv.Atoi(c.PostForm("id"))
	ProfileSlug := c.PostForm("name")
	flg := false

	permisison, perr := NewAuth.IsGranted("Member", auth.Read, TenantId)
	if perr != nil {
		ErrorLog.Printf("member check profile slug authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Member authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if ProfileSlug != "" {
		flg2, err := MemberConfig.CheckProfileSlugInMember(userid, ProfileSlug, TenantId) //check memberprofile slug
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
}

// member claim status if status enable then only claim email send
func ActivateClaimStatus(c *gin.Context) {

	memberid, err1 := strconv.Atoi(c.Request.PostFormValue("memberid"))
	if err1 != nil {
		ErrorLog.Printf("cannot parse memberid into integer error: %s", err1)
	}

	claimstatus := c.Request.PostFormValue("claimb1")

	var claimint int
	if claimstatus == "on" {
		claimint = 1
	} else {
		claimint = 0
	}

	permisison, perr := NewAuth.IsGranted("Member", auth.Read, TenantId)
	if perr != nil {
		ErrorLog.Printf("member check profile slug authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Member authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	member, err := MemberConfig.GetMemberDetails(memberid, TenantId)
	if err != nil {
		ErrorLog.Printf("get member details error: %s", err)
	}

	memberprof, err := MemberConfig.GetMemberProfileByMemberId(memberid, TenantId)
	if err != nil {
		ErrorLog.Printf("get member profiles details error: %s", err)
	}

	currenttime, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))

	if err := DB.Debug().Table("tbl_member_profiles").Where("member_id=?", memberid).UpdateColumns(map[string]interface{}{
		"claim_status": claimint, "modified_by": c.GetInt("userid"), "claim_date": currenttime}).Error; err != nil {
		ErrorLog.Printf("Activate claim query status error: %s", err)
	}

	if member.IsActive == 0 {
		ErrorLog.Println("member status not enable can't send mail")
		c.SetCookie("get-toast", "Member Claim Updated Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, "/member")
		return
	}

	var email models.TblEmailTemplate
	err = models.GetTemplates(&email, "Owndesk", TenantId)
	if err != nil {
		ErrorLog.Printf("emailtemplate get error: %s", err)
	}

	if email.IsActive == 1 && claimint == 1 {
		Token, _ := NewAuth.GenerateMemberToken(member.Id, "admin", SecretKey, TenantId)
		SendMemberClaimEmail(member.Email, member.FirstName+" "+member.LastName, memberprof.CompanyName, Token, memberprof.ProfileSlug)
	} else {
		WarnLog.Println("email claim status not enabled")
	}

	c.SetCookie("get-toast", "Member Claim Updated Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, "/member")
}

// Get member profile
func GetMemberProfileByMemberId(c *gin.Context) {

	memberid, _ := strconv.Atoi(c.PostForm("memberid"))

	permisison, perr := NewAuth.IsGranted("Member", auth.Read, TenantId)
	if perr != nil {
		ErrorLog.Printf("member check profile slug authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Member authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	member, err := MemberConfig.GetMemberDetails(memberid, TenantId)
	if err != nil {
		ErrorLog.Printf("get member details error: %s", err)
		json.NewEncoder(c.Writer).Encode(gin.H{"flg": false})
		return
	}

	memberprof, err := MemberConfig.GetMemberProfileByMemberId(memberid, TenantId)
	if err != nil {
		json.NewEncoder(c.Writer).Encode(gin.H{"flg": false})
		ErrorLog.Printf("get member profiles details error: %s", err)
	}

	if member.ProfileImagePath != "" {
		if member.StorageType == "local" {
			member.ProfileImagePath = "/" + member.ProfileImagePath
		} else if member.StorageType == "aws" {
			member.ProfileImagePath = "/image-resize?name=" + member.ProfileImagePath
		}
	}

	if memberprof.CompanyLogo != "" {
		if memberprof.StorageType == "local" {
			memberprof.CompanyLogo = "/" + memberprof.CompanyLogo
		} else if memberprof.StorageType == "aws" {
			memberprof.CompanyLogo = "/image-resize?name=" + memberprof.CompanyLogo
		}
	}

	if !memberprof.ClaimDate.IsZero() {
		memberprof.Website = memberprof.ClaimDate.Format("02 Jan 2006")
	} else {
		memberprof.Website = time.Now().Format("02 Jan 2006")
	}

	currenttime := time.Now().In(TZONE).Format("02 Jan 2006")
	fmt.Println("dsfdsfs", member)

	fmt.Println("final")

	json.NewEncoder(c.Writer).Encode(gin.H{"member": member, "memberprofile": memberprof, "flg": true, "currenttime": currenttime})

}

// send mail claim
func SendMemberClaimEmail(email, memberName, companyName string, Token string, profileslug string) {

	var url_prefix = os.Getenv("BASE_URL")

	data := map[string]interface{}{
		"client_name":    memberName,
		"company_name":   companyName,
		"logo":           url_prefix + "public/img/email-icons/logo.png",
		"fb_logo":        url_prefix + "public/img/email-icons/facebook.png",
		"linkedin_logo":  url_prefix + "public/img/email-icons/linkedin.png",
		"spacex_logo":    url_prefix + "public/img/email-icons/x.png",
		"youtube_logo":   url_prefix + "public/img/email-icons/youtube.png",
		"instagram_logo": url_prefix + "public/img/email-icons/instagram.png",
		"Link":           os.Getenv("OWNDESK_URL") + "company/" + profileslug + "/" + Token,
		"Slug":           profileslug,
	}

	var wg sync.WaitGroup

	wg.Add(1)

	Chan := make(chan string, 1)

	go OwndeskmemberEmail(Chan, &wg, data, email, "Owndesk")

	close(Chan)

}
