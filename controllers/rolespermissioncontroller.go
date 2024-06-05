package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	nauth "github.com/spurtcms/auth"
	"github.com/spurtcms/pkgcore/auth"
	role "github.com/spurtcms/team-roles"
	csrf "github.com/utrack/gin-csrf"
)

var Role auth.Role

var Permission auth.PermissionAu

/*Roles list*/
func RoleView(c *gin.Context) {

	var (
		limt      int
		offset    int
		NewModule []TblModule
	)

	Role.Auth = AUTH

	//get values from html form data
	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	keyword := c.Query("keyword")

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	permisison, perr := NewAuth.IsGranted("roles", nauth.CRUD)
	if perr != nil {
		ErrorLog.Printf("roleslist authorization error: %s", perr)
	}

	if permisison {

		roles, rolecount, err := NewRole.RoleList(role.Rolelist{Limit: limt, Offset: offset, Filter: role.Filter{Keyword: keyword}})
		if err != nil {
			ErrorLog.Printf("roleslist error: %s", err)
		}

		var newroles []role.Tblrole
		for _, val := range roles {

			permissionid, err := NewRole.GetPermissionDetailsById(val.Id)
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
		}, "Menu": menu, "csrf": csrf.GetToken(c), "HeadTitle": translate.Roles + " & " + translate.Permission.Permissions, "translate": translate, "roles": newroles, "Count": rolecount, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "Previous": Previous, "Next": Next, "PageCount": PageCount, "Page": Page, "CurrentPage": pageno, "Module": NewModule, "Limit": limt, "keyword": keyword, "title": "Roles & Permissions", "SettingsHead": true, "Rolesmenu": true, "Tooltiptitle": translate.Setting.Rolestooltip})

		return

	}
	c.Redirect(301, "/403-page")

}

/*Get data using id*/
func GetRolePermissionData(c *gin.Context) {

	var roleid, _ = strconv.Atoi(c.Request.PostFormValue("id"))

	role, err := NewRole.GetRoleById(roleid)
	if err != nil {
		ErrorLog.Printf("get role details error: %s", err)
	}

	permissionid, perr := NewRole.GetPermissionDetailsById(roleid)
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

	permisison, perr := NewAuth.IsGranted("roles", nauth.CRUD)
	if perr != nil {
		ErrorLog.Printf("rolescreate authorization error: %s", perr)
	}

	if permisison {
		/*role created*/
		roledetail, rerr := NewRole.CreateRole(role.RoleCreation{Name: rolename, Description: description, CreatedBy: userid})
		if rerr != nil {
			ErrorLog.Printf("rolescreate error: %s", perr)
			json.NewEncoder(c.Writer).Encode(rerr)
			return
		}

		/*permission created*/
		perr := NewRole.CreatePermission(role.MultiPermissin{RoleId: roledetail.Id, Ids: idi, CreatedBy: userid})
		if perr != nil {
			ErrorLog.Printf("rolepermission create error: %s", perr)
		}

		c.JSON(200, gin.H{"role": "added"})
		return

	}

	c.Redirect(301, "/403-page")

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

	permisison, perr := NewAuth.IsGranted("roles", nauth.CRUD)
	if perr != nil {
		ErrorLog.Printf("rolesupdate authorization error: %s", perr)
	}

	if permisison {
		/*role update*/
		roledetail, rerr := NewRole.UpdateRole(role.RoleCreation{Name: rolename, Description: description, CreatedBy: userid}, roleid)
		if rerr != nil {
			ErrorLog.Printf("rolesupdate error: %s", perr)
			json.NewEncoder(c.Writer).Encode(rerr)
			return
		}

		/*permission created*/
		perr := NewRole.CreateUpdatePermission(role.MultiPermissin{RoleId: roledetail.Id, Ids: idi, CreatedBy: userid})
		if perr != nil {
			ErrorLog.Printf("rolepermission update error: %s", perr)
		}

		c.JSON(200, gin.H{"role": "updated"})
		return

	}

	c.Redirect(301, "/403-page")

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

	permisison, perr := NewAuth.IsGranted("roles", nauth.CRUD)
	if perr != nil {
		ErrorLog.Printf("rolesdelete authorization error: %s", perr)
	}

	if permisison {

		_, err := NewRole.DeleteRole([]int{}, id)

		if err != nil {
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			ErrorLog.Printf("rolesdelete authorization error: %s", perr)
		}

		c.SetCookie("get-toast", "Role Deleted Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, url)
		return

	}

	c.Redirect(301, "/403-page")

}

/*Check Role Already Exists*/
func CheckRoleAlreadyExists(c *gin.Context) {

	//get values from html form data
	roleid, _ := strconv.Atoi(c.PostForm("id"))
	Rolename := c.PostForm("name")

	_, err := NewRole.CheckRoleAlreadyExists(roleid, Rolename)

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

	err := NewRole.RoleStatus(id, val, userid)
	if err != nil {
		ErrorLog.Printf("roles status error: %s", err)
		json.NewEncoder(c.Writer).Encode(false)
	} else {
		json.NewEncoder(c.Writer).Encode(true)
	}

}

// multiple role delete
func MultiselectRoleDelete(c *gin.Context) {

	var roledata []map[string]string
	if err := json.Unmarshal([]byte(c.Request.PostFormValue("roleids")), &roledata); err != nil {
		fmt.Println(err)
	}

	var roleids []int

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

	permisison, perr := NewAuth.IsGranted("roles", nauth.CRUD)
	if perr != nil {
		ErrorLog.Printf("rolesmultiple delete authorization error: %s", perr)
	}

	if permisison {

		_, err := NewRole.DeleteRole(roleids, 0)
		if err != nil {
			ErrorLog.Printf("rolesmultiple delete error: %s", err)
			c.JSON(200, gin.H{"value": false})
			return
		}

		c.JSON(200, gin.H{"value": true, "url": url})
		return

	}

	c.Redirect(301, "/403-page")

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

	pageno := c.PostForm("page")
	var url string
	if pageno != "" {
		url = "/settings/roles?page=" + pageno
	} else {
		url = "/settings/roles/"

	}

	userid := c.GetInt("userid")

	permisison, perr := NewAuth.IsGranted("roles", nauth.CRUD)
	if perr != nil {
		ErrorLog.Printf("rolesmultiplestatus  authorization error: %s", perr)
	}

	if permisison {
		err := NewRole.MultiSelectRoleStatus(roleids, status, userid)
		if err != nil {
			ErrorLog.Printf("rolesmultiplestatus error: %s", err)
			c.JSON(200, gin.H{"value": false})
			return
		}

		c.JSON(200, gin.H{"value": true, "status": status, "url": url})
		return
	}

	c.Redirect(301, "/403-page")

}
