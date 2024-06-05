package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"spurt-cms/models"

	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	"github.com/spurtcms/pkgcore/teams"
	"github.com/spurtcms/team"
	teamroles "github.com/spurtcms/team-roles"
	csrf "github.com/utrack/gin-csrf"
)

var query models.QUERY

var User teams.TeamAuth

// user list
func Userlist(c *gin.Context) {

	var (
		limt, offset int
		filter       teams.Filters
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

	permisison, perr := NewAuth.IsGranted("Team", auth.CRUD) //check role based permissions
	if perr != nil {
		ErrorLog.Printf("Userlist authorization error: %s", perr)
	}

	if permisison {

		userlist, count, err := NewTeam.ListUser(limt, offset, team.Filters{Keyword: filter.Keyword})
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
			userlists = append(userlists, val)
		}

		var paginationendcount = len(userlist) + offset
		paginationstartcount := offset + 1
		Previous, Next, PageCount, Page := Pagination(pageno, int(count), limt)

		menu := NewMenuController(c)

		translate, _ := TranslateHandler(c)

		roles, _, rerr := NewRole.RoleList(teamroles.Rolelist{GetAllData: true})
		if rerr != nil {
			ErrorLog.Printf("Getting all rolelist data error : %s", rerr)
		}

		c.HTML(200, "teams.html", gin.H{"csrf": csrf.GetToken(c), "Pagination": PaginationData{
			NextPage:     pageno + 1,
			PreviousPage: pageno - 1,
			TotalPages:   PageCount,
			TwoAfter:     pageno + 2,
			TwoBelow:     pageno - 2,
			ThreeAfter:   pageno + 3,
		}, "HeadTitle": translate.Setting.Teams, "Roles": roles, "user": userlists, "Count": count, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Limit": limt, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "Menu": menu, "Filter": filter, "translate": translate, "chcount": count, "SettingsHead": true, "Teamsmenu": true, "Tooltiptitle": translate.Setting.Teamstooltip, "title": "Teams"})

		return

	}

	c.Redirect(301, "/403-page")

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

	var imageName, imagePath string

	if imagedata != "" {
		imageName, imagePath, _ = ConvertBase64(imagedata, "storage/user")
	}

	//pass values to the team struct
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
	}

	permisison, perr := NewAuth.IsGranted("Team", auth.CRUD)
	if perr != nil {
		ErrorLog.Printf("CreateUser authorization error: %s", perr)
	}

	if permisison {
		_, err := NewTeam.CreateUser(Newuser)

		if strings.Contains(fmt.Sprint(err), "given some values is empty") {
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

		if IsActive == 1 {
			fname := c.PostForm("user_fname")
			uname := c.PostForm("user_name")
			pas := c.PostForm("user_pass")
			email := c.PostForm("user_email")

			var url_prefix = os.Getenv("BASE_URL")

			data := map[string]interface{}{

				"fname":         fname,
				"uname":         uname,
				"Pass":          pas,
				"login_url":     url_prefix,
				"admin_logo":    url_prefix + "public/img/spurtcms.png",
				"fb_logo":       url_prefix + "public/img/facebook.png",
				"linkedin_logo": url_prefix + "public/img/linkedin.png",
				"twitter_logo":  url_prefix + "public/img/twitter.png",
			}

			var wg sync.WaitGroup
			wg.Add(1)
			go UserCreateMail(&wg, data, email, "Createuser")

		}

		c.SetCookie("get-toast", "User Created Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, "/settings/users/")
		return

	}

	c.Redirect(301, "/403-page")

}

// check email   already exists
func CheckEmail(c *gin.Context) {

	//get values from html form data
	userid, _ := strconv.Atoi(c.PostForm("id"))
	email := c.PostForm("email")

	_, _, err := NewTeam.CheckEmail(email, userid)

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

	_, err := NewTeam.CheckNumber(number, userid)

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

	_, err := NewTeam.CheckUsername(username, userid)
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
	userdet, _ := NewTeam.GetUserById(id)

	var firstn = strings.ToUpper(userdet.FirstName[:1])

	var lastn string
	if userdet.LastName != "" {
		lastn = strings.ToUpper(userdet.LastName[:1])
	}

	userdet.NameString = firstn + lastn
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

	if imagedata != "" {
		imageName, imagePath, _ = ConvertBase64(imagedata, "storage/user")
	}

	userid, _ := strconv.Atoi(c.PostForm("userid"))

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
	}

	permisison, perr := NewAuth.IsGranted("Team", auth.CRUD)
	if perr != nil {
		ErrorLog.Printf("UpdateUser authorization error: %s", perr)
	}

	if permisison {

		_, err := NewTeam.UpdateUser(Newuser, userid)

		if strings.Contains(fmt.Sprint(err), "given some values is empty") {
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

		return

	}

	c.Redirect(301, "/403-page")

}

// DeleteUser
func DeleteUser(c *gin.Context) {

	//get values from html form data
	userId, _ := strconv.Atoi(c.Param("id"))
	pageno := c.Query("page")

	var url string

	if pageno != "" {
		url = "/settings/users?page=" + pageno
	} else {
		url = "/settings/users/"
	}

	userid := c.GetInt("userid")

	permisison, perr := NewAuth.IsGranted("Team", auth.CRUD)
	if perr != nil {
		ErrorLog.Printf("DeleteUser authorization error: %s", perr)
	}

	if permisison {

		err := NewTeam.DeleteUser([]int{}, userId, userid)
		if err != nil {
			ErrorLog.Printf("DeleteUser error: %s", perr)
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			return
		}

		c.SetCookie("get-toast", "User Deleted Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(http.StatusMovedPermanently, url)
		return
	}

	c.Redirect(301, "/403-page")

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

	permisison, perr := NewAuth.IsGranted("Team", auth.CRUD)
	if perr != nil {
		ErrorLog.Printf("DeleteMultipleUser authorization error: %s", perr)
	}

	if permisison {

		err = NewTeam.DeleteUser(userIds, 0, userid)
		if err != nil {
			ErrorLog.Printf("DeleteMultipleUser  error: %s", err)
			c.JSON(200, gin.H{"value": false})
			return
		}

		c.JSON(200, gin.H{"value": true, "url": url})
		return

	}

	c.Redirect(301, "/403-page")

}

/* Check overall validation in single function */
func CheckUserData(c *gin.Context) {

	//get values from html form data
	userid, _ := strconv.Atoi(c.PostForm("id"))
	username := c.PostForm("username")
	email := c.PostForm("email")
	number := c.PostForm("mobile")

	_, err1 := NewTeam.CheckUsername(username, userid)
	_, _, err2 := NewTeam.CheckEmail(email, userid)
	_, err3 := NewTeam.CheckNumber(number, userid)

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

		if status == "All Users" {
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

	permisison, perr := NewAuth.IsGranted("Team", auth.CRUD)
	if perr != nil {
		ErrorLog.Printf("multiple access change authorization error: %s", perr)
	}

	if permisison {

		err = NewTeam.ChangeAccess(entryIds, userid, statusInt)
		if err != nil {
			ErrorLog.Printf("while changing the access in multiple user error: %s", err)
			c.JSON(200, gin.H{"value": false})
			return
		}
		c.JSON(200, gin.H{"value": true, "status": statusInt, "url": url})
		return
	}

	c.Redirect(301, "/403-page")

}

func ChangeActiveStatus(c *gin.Context) {

	//get values from html form data
	id, _ := strconv.Atoi(c.Request.PostFormValue("id"))
	activeStatus, _ := strconv.Atoi(c.Request.PostFormValue("isActive"))
	userid := c.GetInt("userid")

	permisison, perr := NewAuth.IsGranted("Team", auth.CRUD)
	if perr != nil {
		ErrorLog.Printf("DeleteUser authorization error: %s", perr)
	}

	if permisison {

		flg, err := NewTeam.ChangeActiveStatus(id, activeStatus, userid)
		if err != nil {
			ErrorLog.Printf("change active status in single user error : %s ", err)
			json.NewEncoder(c.Writer).Encode(flg)
		}

		json.NewEncoder(c.Writer).Encode(flg)
		return

	}

	c.Redirect(301, "/403-page")

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

	permisison, perr := NewAuth.IsGranted("Team", auth.CRUD)
	if perr != nil {
		ErrorLog.Printf("multiple statuc change authorization error: %s", perr)
	}

	if permisison {

		err := NewTeam.SelectedUserStatusChange(userdIds, activeStatus, userid)
		if err != nil {
			ErrorLog.Printf("change active status in multiple user error : %s ", err)
			c.JSON(200, gin.H{"value": false})
			return
		}

		c.JSON(200, gin.H{"value": true, "status": activeStatus, "url": url})
		return
	}

	c.Redirect(301, "/403-page")

}
