package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	nauth "github.com/spurtcms/auth"
	role "github.com/spurtcms/team-roles"
	csrf "github.com/utrack/gin-csrf"
)

/*Roles list*/
func RoleView(c *gin.Context) {

	var (
		limt      int
		offset    int
		NewModule []TblModule
	)

	//get values from html form data
	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	keyword := c.Query("keyword")


	var filterflag bool
	if keyword != "" {
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

	permisison, perr := NewAuth.IsGranted("Roles & Permissions", nauth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("roleslist authorization error: %s", perr)
	}

	if !permisison {
		c.Redirect(301, "/403-page")
		return
	}

	NewRole.DataAccess = c.GetInt("dataaccess")
	NewRole.UserId = c.GetInt("userid")

	roles, rolecount, err := NewRole.RoleList(role.Rolelist{Limit: limt, Offset: offset, Filter: role.Filter{Keyword: keyword}}, TenantId, false)
	if err != nil {
		ErrorLog.Printf("roleslist error: %s", err)
	}

	var newroles []role.Tblrole
	for _, val := range roles {

		permissionid, err := NewRole.GetPermissionDetailsById(val.Id, TenantId)
		if err != nil {
			ErrorLog.Printf("roleslist permission details error: %s", err)
		}

		if len(permissionid) > 0 {
			val.Slug = "Manage"
		} else {
			val.Slug = "Yettoconf"
		}

		if !val.ModifiedOn.IsZero() {
			val.CreatedDate = val.ModifiedOn.In(TZONE).Format(Datelayout)
		} else {
			val.CreatedDate = val.CreatedOn.In(TZONE).Format(Datelayout)
		}

		newroles = append(newroles, val)
	}

	/*Get module permission list*/
	// var NewModule []role.Tblmodule

	// Module, _, err := NewRole.PermissionListRoleId(100, 0, c.GetInt("role"), role.Filter{})
	// if err != nil {
	// 	ErrorLog.Printf("get permissionlist error: %s", err)
	// }
	Module := Permissions()

	for _, val := range Module {
		/*this is for display whitespcae remove */
		val.Description = strings.ReplaceAll(val.ModuleName, " ", "")
		NewModule = append(NewModule, val)
	}
	loguserid := c.GetInt("userid")

	roleid := NewRole.GetRoleids(loguserid)

	//pagination calc
	paginationendcount := len(roles) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(pageno, int(rolecount), limt)

	translate, _ := TranslateHandler(c)
	menu := NewMenuController(c)

	c.HTML(200, "rolesandpermission.html", gin.H{"Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3,
	}, "Menu": menu, "csrf": csrf.GetToken(c), "linktitle": "Roles & Permissions", "roleid": roleid, "Searchtrue": filterflag, "loguserid": loguserid, "HeadTitle": translate.Roles + " & " + translate.Permission.Permissions, "translate": translate, "roles": newroles, "Count": rolecount, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "Previous": Previous, "Next": Next, "PageCount": PageCount, "Page": Page, "CurrentPage": pageno, "Module": NewModule, "Limit": limt, "keyword": keyword, "title": "Roles & Permissions", "SettingsHead": true, "Rolesmenu": true, "Tooltiptitle": translate.Setting.Rolestooltip})

}

/*Get data using id*/
func GetRolePermissionData(c *gin.Context) {

	var roleid, _ = strconv.Atoi(c.Request.PostFormValue("id"))

	role, err := NewRole.GetRoleById(roleid, TenantId)
	if err != nil {
		ErrorLog.Printf("get role details error: %s", err)
	}

	permissionid, perr := NewRole.GetPermissionDetailsById(roleid, TenantId)
	if perr != nil {
		ErrorLog.Printf("get permission details error: %s", perr)
	}

	json.NewEncoder(c.Writer).Encode(gin.H{"role": role, "permissionid": permissionid})
}

/*Roles Create or update*/
func RoleCreate(c *gin.Context) {

	//get values from html form data
	rolename := c.Request.PostFormValue("rolename")
	description := c.Request.PostFormValue("roledesc")
	permissionsid := c.PostFormArray("permissionid[]")
	roleisactive:= c.Request.PostFormValue("roleisactive")
	fmt.Println("print",roleisactive)
	isactive, err := strconv.Atoi(roleisactive)
	if err != nil {
		panic(err)
	}


	fmt.Println("role data", rolename, permissionsid)

	var idi = []int{}
	if len(permissionsid) > 0 {

		//convert stringarray to intarray
		for _, i := range permissionsid {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			idi = append(idi, j)
		}
	}

	idi = append(idi, 1)

	userid := c.GetInt("userid")
	fmt.Println("dtetet1")

	permisison, perr := NewAuth.IsGranted("Roles & Permissions", nauth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("rolescreate authorization error: %s", perr)
	}

	if !permisison {
		c.Redirect(301, "/403-page")
		return
	}
	/*role created*/

	fmt.Println("tetetett2")
	roledetail, rerr := NewRole.CreateRole(role.RoleCreation{Name: rolename, Description: description, CreatedBy: userid, TenantId: TenantId},isactive)
	if rerr != nil {
		ErrorLog.Printf("rolescreate error: %s", rerr)
		json.NewEncoder(c.Writer).Encode(rerr)
		return
	}

	/*permission created*/
	prerr := NewRole.CreatePermission(role.MultiPermissin{RoleId: roledetail.Id, Ids: idi, CreatedBy: userid, TenantId: TenantId})
	if perr != nil {
		ErrorLog.Printf("rolepermission create error: %s", prerr)
	}

	c.JSON(200, gin.H{"role": "added"})

}

/*Roles Create or update*/
func RoleUpdate(c *gin.Context) {

	//get values from html form data
	roleid, _ := strconv.Atoi(c.Request.PostFormValue("roleid"))
	rolename := c.Request.PostFormValue("rolename")
	description := c.Request.PostFormValue("roledesc")
	permissionsid := c.PostFormArray("permissionid[]")

	var idi = []int{}
	if len(permissionsid) > 0 {

		//convert stringarray to intarray
		for _, i := range permissionsid {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			idi = append(idi, j)
		}
	}

	idi = append(idi, 1, 21)

	userid := c.GetInt("userid")

	permisison, perr := NewAuth.IsGranted("Roles & Permissions", nauth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("rolesupdate authorization error: %s", perr)
	}

	if !permisison {
		c.Redirect(301, "/403-page")
		return
	}
	/*role update*/
	roledetail, rerr := NewRole.UpdateRole(role.RoleCreation{Name: rolename, Description: description, CreatedBy: userid}, roleid, TenantId)
	if rerr != nil {
		ErrorLog.Printf("rolesupdate error: %s", perr)
		json.NewEncoder(c.Writer).Encode(rerr)
		return
	}

	/*permission created*/
	prerr := NewRole.CreateUpdatePermission(role.MultiPermissin{RoleId: roledetail.Id, Ids: idi, CreatedBy: userid, TenantId: TenantId})
	if perr != nil {
		ErrorLog.Printf("rolepermission update error: %s", prerr)
	}

	c.JSON(200, gin.H{"role": "updated"})

}

/*Delete Role*/
func DeleteRole(c *gin.Context) {

	//get values from html form data
	var id, _ = strconv.Atoi(c.Query("id"))
	pageno := c.Query("page")

	var url string
	if pageno != "" {
		url = "/settings/roles?page=" + pageno
	} else {
		url = "/settings/roles/"

	}

	permisison, perr := NewAuth.IsGranted("Roles & Permissions", nauth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("rolesdelete authorization error: %s", perr)
	}

	if !permisison {
		c.Redirect(301, "/403-page")
		return
	}

	_, err := NewRole.DeleteRole([]int{}, id, TenantId)

	if err != nil {
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		ErrorLog.Printf("rolesdelete authorization error: %s", err)
	}
	_, rolecount, err := NewRole.RoleList(role.Rolelist{Limit: 0, Offset: 0, Filter: role.Filter{Keyword: ""}}, TenantId, false)
	if err != nil {
		ErrorLog.Printf("roleslist error: %s", err)
	}

	
	recordsPerPage := Limit
	totalPages := (int(rolecount) + recordsPerPage - 1) / recordsPerPage

	currentPage := 1
	if pageno != "" {
		currentPage, _ = strconv.Atoi(pageno)
	}

	if currentPage > totalPages && totalPages > 0 {
		url = "/settings/roles?page=" + strconv.Itoa(totalPages)
	}

	c.SetCookie("get-toast", "Role Deleted Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, url)

}

/*Check Role Already Exists*/
func CheckRoleAlreadyExists(c *gin.Context) {

	//get values from html form data
	roleid, _ := strconv.Atoi(c.PostForm("id"))
	Rolename := c.PostForm("name")

	_, err := NewRole.CheckRoleAlreadyExists(roleid, Rolename, TenantId)

	var role bool
	if err != nil {
		ErrorLog.Printf("check roles error: %s", err)
		role = false
	} else {
		role = true
	}

	c.JSON(200, gin.H{"role": role})
}

/*permission role status*/
func RoleIsActive(c *gin.Context) {

	//get values from html form data
	id, _ := strconv.Atoi(c.Request.PostFormValue("id"))
	val, _ := strconv.Atoi(c.Request.PostFormValue("isactive"))
	userid := c.GetInt("userid")

	err := NewRole.RoleStatus(id, val, userid, TenantId)
	if err != nil {
		ErrorLog.Printf("roles status error: %s", err)
		json.NewEncoder(c.Writer).Encode(false)
	} else {
		json.NewEncoder(c.Writer).Encode(true)
	}

}

// multiple role delete
func MultiselectRoleDelete(c *gin.Context) {

	var (
		roledata []map[string]string
		roleids  []int
	)

	if err := json.Unmarshal([]byte(c.Request.PostFormValue("roleids")), &roledata); err != nil {
		ErrorLog.Printf("roles multiple delete unmarshall error: %s", err)
	}

	for _, val := range roledata {
		roleid, _ := strconv.Atoi(val["roleid"])
		roleids = append(roleids, roleid)
	}

	pageno := c.PostForm("page")
	var url string
	if pageno != "" {
		url = "/settings/roles?page=" + pageno
	} else {
		url = "/settings/roles/"

	}

	permisison, perr := NewAuth.IsGranted("Roles & Permissions", nauth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("rolesmultiple delete authorization error: %s", perr)
	}

	if !permisison {
		c.Redirect(301, "/403-page")
		return
	}

	_, err := NewRole.DeleteRole(roleids, 0, TenantId)
	if err != nil {
		ErrorLog.Printf("rolesmultiple delete error: %s", err)
		c.JSON(200, gin.H{"value": false})
		return
	}

	_, rolecount, err := NewRole.RoleList(role.Rolelist{Limit: 0, Offset: 0, Filter: role.Filter{Keyword: ""}}, TenantId, false)
	if err != nil {
		ErrorLog.Printf("roleslist error: %s", err)
	}

	
	recordsPerPage := Limit
	totalPages := (int(rolecount) + recordsPerPage - 1) / recordsPerPage

	currentPage := 1
	if pageno != "" {
		currentPage, _ = strconv.Atoi(pageno)
	}

	if currentPage > totalPages && totalPages > 0 {
		url = "/settings/roles?page=" + strconv.Itoa(totalPages)
	}

	c.JSON(200, gin.H{"value": true, "url": url})

}

func MultiSelectRoleStatus(c *gin.Context) {

	var roledata []map[string]string
	if err := json.Unmarshal([]byte(c.Request.PostFormValue("roleids")), &roledata); err != nil {
		fmt.Println(err)
	}

	var (
		roleids   []int
		statusint int
		status    int
	)

	for _, val := range roledata {

		roleid, _ := strconv.Atoi(val["roleid"])
		statusint, _ = strconv.Atoi(val["status"])
		if statusint == 0 {
			status = 1
		} else if statusint == 1 {
			status = 0
		}

		roleids = append(roleids, roleid)

	}

	fmt.Println("state", status)

	pageno := c.PostForm("page")
	var url string
	if pageno != "" {
		url = "/settings/roles?page=" + pageno
	} else {
		url = "/settings/roles/"

	}

	userid := c.GetInt("userid")

	permisison, perr := NewAuth.IsGranted("Roles & Permissions", nauth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("rolesmultiplestatus  authorization error: %s", perr)
	}

	if !permisison {
		c.Redirect(301, "/403-page")
		return
	}
	err := NewRole.MultiSelectRoleStatus(roleids, status, userid, TenantId)
	if err != nil {
		ErrorLog.Printf("rolesmultiplestatus error: %s", err)
		c.JSON(200, gin.H{"value": false})
		return
	}
	c.SetCookie("get-toast", "Role Updated Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.JSON(200, gin.H{"value": true, "status": status, "url": url})
}


func Chkroleshaveuser(c *gin.Context){
	var (
		rolesdata []map[string]string
		rolesids  []int
	)
	roles_id, _ := strconv.Atoi(c.PostForm("rolesid"))


	if err := json.Unmarshal([]byte(c.Request.PostFormValue("roleids")), &rolesdata); err != nil {
		fmt.Println(err)
	}
	for _, val := range rolesdata {
		roleid, _ := strconv.Atoi(val["roleid"])
		rolesids = append(rolesids, roleid)
	}

	_, err := NewRole.Rolescheckusers(roles_id,rolesids, TenantId)

	if err != nil {
		ErrorLog.Printf("Roles and permission check hane Error :%s", err)
		c.JSON(200, gin.H{"value": false})
		return
	}
	c.JSON(200, gin.H{"value": true})
}