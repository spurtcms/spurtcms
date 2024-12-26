package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"spurt-cms/models"
	storagecontroller "spurt-cms/storage-controller"

	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	"github.com/spurtcms/team"
	teamroles "github.com/spurtcms/team-roles"
	csrf "github.com/utrack/gin-csrf"
)

var query models.QUERY

// user list
func Userlist(c *gin.Context) {

	var (
		limt, offset int
		filter       team.Filters
	)

	//get values from html form data
	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")

	var filterflag bool
	if filter.Keyword != "" {
		filterflag = true
	} else {
		filterflag = false
	}

	permisison, perr := NewAuth.IsGranted("Team", auth.CRUD, TenantId) //check role based permissions
	if perr != nil {
		ErrorLog.Printf("Userlist authorization error: %s", perr)
	}

	if !permisison {
		c.Redirect(301, "/403-page")
		return
	}

	NewTeam.Dataaccess = c.GetInt("dataaccess")
	NewTeam.Userid = c.GetInt("userid")

	userlist, count, err := NewTeam.ListUser(limt, offset, team.Filters{Keyword: filter.Keyword}, TenantId)
	if err != nil {
		ErrorLog.Printf("Getting userlist data error : %s", err)
	}

	var userlists []team.TblUser
	for _, val := range userlist {
		if !val.ModifiedOn.IsZero() { //change date format to display on list
			val.ModuleName = val.ModifiedOn.In(TZONE).Format(Datelayout)
		} else {
			val.ModuleName = val.CreatedOn.In(TZONE).Format(Datelayout)
		}

		if val.ProfileImagePath != "" {
			if val.StorageType == "local" {
				val.ProfileImagePath = "/" + val.ProfileImagePath
			} else if val.StorageType == "aws" {
				val.ProfileImagePath = "/image-resize?name=" + val.ProfileImagePath
			}
		}

		userlists = append(userlists, val)
	}

	var paginationendcount = len(userlist) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(pageno, int(count), limt)

	menu := NewMenuController(c)

	translate, _ := TranslateHandler(c)

	roles, _, rerr := NewRole.RoleList(teamroles.Rolelist{GetAllData: true}, TenantId, true)
	if rerr != nil {
		ErrorLog.Printf("Getting all rolelist data error : %s", rerr)
	}

	fmt.Println("role", roles)

	loguserid, _ := c.Get("userid")

	storagetype, err := GetSelectedType()
	if err != nil {
		fmt.Printf("member list getting storagetype error: %s", err)
	}

	roleidlog := c.GetInt("role")

	c.HTML(200, "teams.html", gin.H{"csrf": csrf.GetToken(c), "Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3,
	}, "HeadTitle": translate.Setting.Teams, "linktitle": "Teams", "Searchtrue": filterflag, "Roles": roles, "user": userlists, "Count": count, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Limit": limt, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "Menu": menu, "Filter": filter, "translate": translate, "chcount": count, "SettingsHead": true, "Teamsmenu": true, "Tooltiptitle": translate.Setting.Teamstooltip, "title": "Teams", "loggedinuser": loguserid, "StorageType": storagetype.SelectedType, "roleidlog": roleidlog})

}

// create user
func CreateUser(c *gin.Context) {

	//check validation
	if c.PostForm("user_fname") == "" || c.PostForm("user_email") == "" || c.PostForm("user_mob") == "" || c.PostForm("user_name") == "" || c.PostForm("user_pass") == "" || c.PostForm("user_role") == "" {
		ErrorLog.Printf("create user error mandatory field error")
		c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
		c.Redirect(301, "/settings/users/")
		return
	}

	//get values from html form data
	Roleid, _ := strconv.Atoi(c.PostForm("user_role"))
	IsActive, _ := strconv.Atoi(c.PostForm("mem_activestat"))
	DataAccess, _ := strconv.Atoi(c.PostForm("mem_data_access"))
	imagedata := c.PostForm("crop_data")

	// tenantDetails, err := NewTeam.GetTenantDetails(TenantId)
	// if err != nil {
	// 	ErrorLog.Printf("unable to get the tenant details: %s", err)
	// }

	// if (tenantDetails.Id == 1) && (tenantDetails.S3FolderName == "") {
	// 	var s3FolderName = "SuperAdminFolder"
	// 	s3Path, err := storagecontroller.CreateFolderToS3(s3FolderName, "")
	// 	if err != nil {
	// 		fmt.Println("error creating a folder for super user in s3 bucket:", err)
	// 		ErrorLog.Printf("error creating a folder for super user in s3 bucket: %v", err)
	// 		return
	// 	}

	// 	err = NewTeam.UpdateS3FolderName(0, tenantDetails.Id, s3Path)
	// 	if err != nil {
	// 		fmt.Println("error updating folder path for super admin in Database:", err)
	// 		ErrorLog.Printf("error updating folder path for super admin in Database: %v", err)
	// 		return
	// 	}

	// 	tenantDetails.S3FolderName = s3Path

	// }

	userDetails, err := GetRequestScopedTenantDetails(c)

	if err != nil {
		ErrorLog.Printf("%s", err)
		return
	}

	var imageName, imagePath string

	storagetype, err := GetSelectedType()

	if err != nil {
		ErrorLog.Printf("error get storage type error: %s", err)
	}

	if storagetype.SelectedType == "local" {

		if imagedata != "" {
			imageName, imagePath, _ = ConvertBase64(imagedata, "storage/user")
		}

	} else if storagetype.SelectedType == "aws" {

		if imagedata != "" {

			var (
				imageByte []byte
				err       error
			)

			imageName, imagePath, imageByte, err = ConvertBase64toByte(imagedata, "user")
			if err != nil {
				ErrorLog.Printf("convert base 64 to byte error : %s", err)
			}

			imagePath = userDetails.S3FolderName + imagePath

			uerr := storagecontroller.UploadCropImageS3(imageName, imagePath, imageByte)
			if uerr != nil {

				c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
				c.Redirect(301, "/settings/users/")
				return
	
			}
		}

	}

	//pass values to the team struct
	Newuser := team.TeamCreate{
		FirstName:         c.PostForm("user_fname"),
		LastName:          c.PostForm("user_lname"),
		Email:             c.PostForm("user_email"),
		MobileNo:          c.PostForm("user_mob"),
		Username:          c.PostForm("user_name"),
		Password:          c.PostForm("user_pass"),
		ProfileImage:      imageName,
		ProfileImagePath:  imagePath,
		RoleId:            Roleid,
		IsActive:          IsActive,
		DataAccess:        DataAccess,
		CreatedBy:         userDetails.Id,
		StorageType:       storagetype.SelectedType,
		TenantId:          TenantId,
		S3FolderPath:      userDetails.S3FolderName,
		DefaultLanguageId: userDetails.DefaultLanguageId,
	}

	permisison, perr := NewAuth.IsGranted("Team", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("CreateUser authorization error: %s", perr)
	}

	if !permisison {
		c.Redirect(301, "/403-page")
		return
	}

	_, _, terr := NewTeam.CreateUser(Newuser)

	if strings.Contains(fmt.Sprint(terr), "given some values is empty") {
		c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
		c.Redirect(301, "/settings/users/")
		return
	}

	if err != nil {
		ErrorLog.Printf("Creating User error: %s", err)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
		c.Redirect(301, "/settings/users/")
		return
	}

	// var email models.TblEmailTemplate
	// err = models.GetTemplates(&email, "createuser", TenantId)
	// if err != nil {
	// 	ErrorLog.Printf("emailtemplate get error: %s", err)
	// }

	if IsActive == 1 {
		fname := c.PostForm("user_fname")
		uname := c.PostForm("user_name")
		pas := c.PostForm("user_pass")
		email := c.PostForm("user_email")

		var url_prefix = os.Getenv("BASE_URL")
		linkedin := os.Getenv("LINKEDIN")
		facebook := os.Getenv("FACEBOOK")
		twitter := os.Getenv("TWITTER")
		youtube := os.Getenv("YOUTUBE")
		insta := os.Getenv("INSTAGRAM")

		data := map[string]interface{}{

			"fname":         fname,
			"uname":         uname,
			"Pass":          pas,
			"login_url":     url_prefix,
			"admin_logo":    url_prefix + "public/img/SpurtCMSlogo.png",
			"fb_logo":       url_prefix + "public/img/email-icons/facebook.png",
			"linkedin_logo": url_prefix + "public/img/email-icons/linkedin.png",
			"twitter_logo":  url_prefix + "public/img/email-icons/x.png",
			"youtube_logo":  url_prefix + "public/img/email-icons/youtube.png",
			"insta_log":     url_prefix + "public/img/email-icons/instagram.png",
			"facebook":      facebook,
			"instagram":     insta,
			"youtube":       youtube,
			"linkedin":      linkedin,
			"twitter":       twitter,
		}

		var wg sync.WaitGroup
		wg.Add(1)
		go UserCreateMail(&wg, data, email, "createuser")

	} else {
		WarnLog.Println("Create Team email notification status not enabled")
	}

	c.SetCookie("get-toast", "User Created Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, "/settings/users/")
}

// check email   already exists
func CheckEmail(c *gin.Context) {

	//get values from html form data
	userid, _ := strconv.Atoi(c.PostForm("id"))
	email := c.PostForm("email")

	_, _, err := NewTeam.CheckEmail(email, userid, TenantId)

	if err != nil {
		ErrorLog.Printf("check email error: %s", err)
		json.NewEncoder(c.Writer).Encode(false)
		return
	}

	json.NewEncoder(c.Writer).Encode(true)

}

// check number already exists
func CheckNumber(c *gin.Context) {

	//get values from html form data
	userid, _ := strconv.Atoi(c.PostForm("id"))
	number := c.PostForm("number")

	_, err := NewTeam.CheckNumber(number, userid, TenantId)

	if err != nil {
		ErrorLog.Printf("check number error: %s", err)
		json.NewEncoder(c.Writer).Encode(false)
		return
	}

	json.NewEncoder(c.Writer).Encode(true)

}

// check username already exists
func CheckUsername(c *gin.Context) {

	userid, _ := strconv.Atoi(c.PostForm("id"))
	username := c.PostForm("username")
	fmt.Println("username", username)

	_, err := NewTeam.CheckUsername(username, userid, TenantId)
	if err != nil {
		ErrorLog.Printf("check Username error: %s", err)
		json.NewEncoder(c.Writer).Encode(false)
		return
	}

	json.NewEncoder(c.Writer).Encode(true)

}

// get user details
func EditUser(c *gin.Context) {

	//get values from html form data
	var id, _ = strconv.Atoi(c.Query("id"))
	userdet, _, _ := NewTeam.GetUserById(id, []int{})

	var firstn = strings.ToUpper(userdet.FirstName[:1])

	var lastn string
	if userdet.LastName != "" {
		lastn = strings.ToUpper(userdet.LastName[:1])
	}
	userdet.NameString = firstn + lastn

	if userdet.ProfileImagePath != "" {
		if userdet.StorageType == "local" {
			userdet.ProfileImagePath = "/" + userdet.ProfileImagePath
		} else if userdet.StorageType == "aws" {
			userdet.ProfileImagePath = "/image-resize?name=" + userdet.ProfileImagePath
		}
	}

	c.JSON(200, userdet)
}

// Update User
func UpdateUser(c *gin.Context) {

	//get values from html form data
	Roleid, _ := strconv.Atoi(c.PostForm("user_role"))
	IsActive, _ := strconv.Atoi(c.PostForm("mem_activestat"))
	DataAccess, _ := strconv.Atoi(c.PostForm("mem_data_access"))
	imagedata := c.PostForm("crop_data")
	pageno := c.PostForm("pageno")

	var (
		url                  string
		imageName, imagePath string
	)

	if pageno != "" {
		url = "/settings/users?page=" + pageno
	} else {
		url = "/settings/users/"

	}

	//check validation
	if c.PostForm("user_fname") == "" || c.PostForm("user_email") == "" || c.PostForm("user_mob") == "" || c.PostForm("user_name") == "" || c.PostForm("user_role") == "" {
		c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
		c.Redirect(301, "/settings/users/")
		return
	}

	userid, _ := strconv.Atoi(c.PostForm("userid"))

	fmt.Println("userId", userid)

	storagetype, err := GetSelectedType()
	if err != nil {
		ErrorLog.Printf("error get storage type error: %s", err)
	}

	if storagetype.SelectedType == "local" {

		if imagedata != "" {
			imageName, imagePath, _ = ConvertBase64(imagedata, storagetype.Local+"/user")
		}

	} else if storagetype.SelectedType == "aws" {

		userDetails, _, err := NewTeam.GetUserById(userid, []int{})
		if err != nil {
			ErrorLog.Printf("error get storage type error: %s", err)
		}

		fmt.Println("userDetails", userDetails)

		if imagedata != "" {

			var (
				imageByte []byte
				err       error
			)

			imageName, imagePath, imageByte, err = ConvertBase64toByte(imagedata, "user")
			if err != nil {
				ErrorLog.Printf("convert base 64 to byte error : %s", err)
			}

			imagePath = userDetails.S3FolderName + imagePath

			uerr := storagecontroller.UploadCropImageS3(imageName, imagePath, imageByte)
			if uerr != nil {

				c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
				c.Redirect(301, "/settings/users/")
				return
	
			}
		}

	}

	//pass values to struct
	Newuser := team.TeamCreate{
		FirstName:        c.PostForm("user_fname"),
		LastName:         c.PostForm("user_lname"),
		Email:            c.PostForm("user_email"),
		MobileNo:         c.PostForm("user_mob"),
		Username:         c.PostForm("user_name"),
		Password:         c.PostForm("user_pass"),
		ProfileImage:     imageName,
		ProfileImagePath: imagePath,
		RoleId:           Roleid,
		IsActive:         IsActive,
		DataAccess:       DataAccess,
		StorageType:      storagetype.SelectedType,
	}

	permisison, perr := NewAuth.IsGranted("Team", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("UpdateUser authorization error: %s", perr)
	}

	if !permisison {
		c.Redirect(301, "/403-page")
		return
	}
	NewTeam.Dataaccess = c.GetInt("dataaccess")
	NewTeam.Userid = c.GetInt("userid")

	_, terr := NewTeam.UpdateUser(Newuser, userid, TenantId)

	if strings.Contains(fmt.Sprint(terr), "given some values is empty") {
		c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
		c.Redirect(301, url)
		return
	}

	if err != nil {
		ErrorLog.Printf("Update User error: %s", err)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
		c.Redirect(301, url)
		return
	}
	c.SetCookie("get-toast", "User Updated Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, url)

}

// DeleteUser
func DeleteUser(c *gin.Context) {

	//get values from html form data
	deluserId, _ := strconv.Atoi(c.Param("id"))
	pageno := c.Query("page")

	var url string

	if pageno != "" {
		url = "/settings/users?page=" + pageno
	} else {
		url = "/settings/users/"
	}

	userid := c.GetInt("userid")

	permisison, perr := NewAuth.IsGranted("Team", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("DeleteUser authorization error: %s", perr)
	}

	if !permisison {
		c.Redirect(301, "/403-page")
		return
	}

	err := NewTeam.DeleteUser([]int{}, deluserId, userid, TenantId)
	if err != nil {
		ErrorLog.Printf("DeleteUser error: %s", perr)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		return
	}

	_, count, err := NewTeam.ListUser(0, 0, team.Filters{Keyword: ""}, TenantId)
	if err != nil {
		ErrorLog.Printf("Getting userlist data error : %s", err)
	}

	recordsPerPage := Limit
	totalPages := (int(count) + recordsPerPage - 1) / recordsPerPage

	currentPage := 1
	if pageno != "" {
		currentPage, _ = strconv.Atoi(pageno)
	}

	if currentPage > totalPages && totalPages > 0 {
		url = "/settings/users?page=" + strconv.Itoa(totalPages)
	}

	c.SetCookie("get-toast", "User Deleted Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(http.StatusMovedPermanently, url)
}

// delete multiple user
func DeleteMultipleUser(c *gin.Context) {

	var userData []map[string]string
	err := json.Unmarshal([]byte(c.Request.PostFormValue("userIds")), &userData)

	if err != nil {
		ErrorLog.Printf("Unmarshal error in deletemultipleuser error: %s", err)
		c.JSON(200, gin.H{"value": false})
		return
	}

	pageNo := c.Request.PostFormValue("page")
	var url string
	if pageNo != "" {
		url = "/settings/users?page=" + pageNo
	} else {
		url = "/settings/users/"
	}

	var userIds []int
	for _, val := range userData {
		userId, _ := strconv.Atoi(val["entryid"])
		userIds = append(userIds, userId)
	}

	userid := c.GetInt("userid")

	permisison, perr := NewAuth.IsGranted("Team", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("DeleteMultipleUser authorization error: %s", perr)
	}

	if !permisison {
		c.Redirect(301, "/403-page")
		return
	}

	err = NewTeam.DeleteUser(userIds, 0, userid, TenantId)
	if err != nil {
		ErrorLog.Printf("DeleteMultipleUser  error: %s", err)
		c.JSON(200, gin.H{"value": false})
		return
	}

	_, count, err := NewTeam.ListUser(0, 0, team.Filters{Keyword: ""}, TenantId)
	if err != nil {
		ErrorLog.Printf("Getting userlist data error : %s", err)
	}


	recordsPerPage := Limit
	totalPages := (int(count) + recordsPerPage - 1) / recordsPerPage

	currentPage := 1
	if pageNo != "" {
		currentPage, _ = strconv.Atoi(pageNo)
	}

	if currentPage > totalPages && totalPages > 0 {
		url = "/settings/users?page=" + strconv.Itoa(totalPages)
	}

	c.JSON(200, gin.H{"value": true, "url": url})

}

/* Check overall validation in single function */
func CheckUserData(c *gin.Context) {

	//get values from html form data
	userid, _ := strconv.Atoi(c.PostForm("id"))
	username := c.PostForm("username")
	email := c.PostForm("email")
	number := c.PostForm("mobile")

	fmt.Println("reapt", userid, username, email, number)

	_, err1 := NewTeam.CheckUsername(username, userid, TenantId)
	_, _, err2 := NewTeam.CheckEmail(email, userid, TenantId)
	_, err3 := NewTeam.CheckNumber(number, userid, TenantId)

	var Number, User, Email bool

	if err1 != nil {
		ErrorLog.Printf("DeleteMultipleUser  error: %s", err1)
		User = false
	} else {
		User = true
	}

	if err2 != nil {
		ErrorLog.Printf("DeleteMultipleUser  error: %s", err2)
		Email = false
	} else {
		Email = true
	}

	if err3 != nil {
		Number = false
	} else {
		Number = true
	}

	c.JSON(200, gin.H{"user": User, "number": Number, "email": Email})
}

// Multi Select Change Access
func SelectedUsersAccessChange(c *gin.Context) {

	var (
		entryIds  []int
		statusInt int
		status    string
		entryData []map[string]string
		url       string
	)

	err := json.Unmarshal([]byte(c.Request.PostFormValue("userIds")), &entryData)
	if err != nil {
		c.JSON(200, gin.H{"value": false})
		ErrorLog.Printf("Unmarshall error multiple access change error : %s", err)
		return
	}

	for _, val := range entryData {
		entryId, _ := strconv.Atoi(val["entryid"])
		status = val["status"]

		if status == "All User" {
			statusInt = 1
		} else if status == "User's Only" {
			statusInt = 0
		}

		entryIds = append(entryIds, entryId)

	}

	pageno := c.PostForm("page")
	if pageno != "" {
		url = "/settings/users/?page=" + pageno
	} else {
		url = "/settings/users/"
	}

	userid := c.GetInt("userid")

	permisison, perr := NewAuth.IsGranted("Team", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("multiple access change authorization error: %s", perr)
	}

	if !permisison {
		c.Redirect(301, "/403-page")
		return
	}

	err = NewTeam.ChangeAccess(entryIds, userid, statusInt, TenantId)
	if err != nil {
		ErrorLog.Printf("while changing the access in multiple user error: %s", err)
		c.JSON(200, gin.H{"value": false})
		return
	}
	c.JSON(200, gin.H{"value": true, "status": statusInt, "url": url})

}

func ChangeActiveStatus(c *gin.Context) {

	//get values from html form data
	id, _ := strconv.Atoi(c.Request.PostFormValue("id"))
	activeStatus, _ := strconv.Atoi(c.Request.PostFormValue("isActive"))
	userid := c.GetInt("userid")

	permisison, perr := NewAuth.IsGranted("Team", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("DeleteUser authorization error: %s", perr)
	}

	if !permisison {
		c.Redirect(301, "/403-page")
		return
	}

	flg, err := NewTeam.ChangeActiveStatus(id, activeStatus, userid, TenantId)
	if err != nil {
		ErrorLog.Printf("change active status in single user error : %s ", err)
		json.NewEncoder(c.Writer).Encode(flg)
	}

	json.NewEncoder(c.Writer).Encode(flg)

}

func SelectedUsersStatusChange(c *gin.Context) {

	var (
		userdIds        []int
		activeStatusInt int
		activeStatus    int
		userData        []map[string]string
	)

	if err := json.Unmarshal([]byte(c.Request.PostFormValue("userIds")), &userData); err != nil {
		ErrorLog.Printf("unmarshall userid multiple status change error: %s", err)
	}

	for _, val := range userData {

		userId, _ := strconv.Atoi(val["entryid"])
		activeStatusInt, _ = strconv.Atoi(val["activeStatus"])

		if activeStatusInt == 0 {
			activeStatus = 1
		} else if activeStatusInt == 1 {
			activeStatus = 0
		}

		userdIds = append(userdIds, userId)
	}

	pageno := c.PostForm("page")
	var url string

	if pageno != "" {
		url = "/settings/users/?page=" + pageno
	} else {
		url = "/settings/users/"
	}

	userid := c.GetInt("userid")

	permisison, perr := NewAuth.IsGranted("Team", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("multiple statuc change authorization error: %s", perr)
	}

	if !permisison {
		c.Redirect(301, "/403-page")
		return
	}

	err := NewTeam.SelectedUserStatusChange(userdIds, activeStatus, userid, TenantId)
	if err != nil {
		ErrorLog.Printf("change active status in multiple user error : %s ", err)
		c.JSON(200, gin.H{"value": false})
		return
	}

	c.JSON(200, gin.H{"value": true, "status": activeStatus, "url": url})

}
