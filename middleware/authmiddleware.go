package middleware

import (
	"fmt"
	"os"
	"regexp"
	"spurt-cms/config"
	"spurt-cms/controllers"
	"spurt-cms/models"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spurtcms/pkgcontent/channels"
	spurtcore "github.com/spurtcms/pkgcore"
	"github.com/spurtcms/pkgcore/auth"
)

// check the jwt token with authorized
func JWTAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		session, _ := controllers.Store.Get(c.Request, os.Getenv("SESSION_KEY"))

		tkn := session.Values["token"]

		if tkn == nil {
			c.Abort()

			c.Writer.Header().Set("Pragma", "no-cache")

			c.Redirect(301, "/")

		} else {

			session.Values["token"] = tkn

			session.Options.MaxAge = 60 * 60 * 2 //2hours

			session.Save(c.Request, c.Writer)

			token := tkn.(string)

			userid, err := controllers.NewAuth.VerifyToken(token, os.Getenv("JWT_SECRET"))

			if err != nil {
				controllers.ErrorLog.Println("err", err)
			}

			Claims := jwt.MapClaims{}
			tkn, err := jwt.ParseWithClaims(token, Claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("ACCESS_SECRET")), nil
			})
			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					fmt.Println(err)
					return
				}

				fmt.Println(err)
				c.Abort()

				c.Writer.Header().Set("Pragma", "no-cache")

				session.Options.MaxAge = -1

				session.Save(c.Request, c.Writer)

				c.Redirect(301, "/")
				return
			}
			if !tkn.Valid {
				fmt.Println(tkn)
				return
			}

			// userid := Claims["user_id"]

			// var user models.TblUser
			controllers.AUTH = spurtcore.NewInstance(&auth.Option{DB: controllers.DB, Token: token, Secret: os.Getenv("JWT_SECRET")})

			controllers.Token = token

			controllers.User.Authority = &controllers.AUTH

			user, _ := controllers.User.GetUserDetails(userid)

			if user.Id == 0 {

				c.Abort()

				c.Writer.Header().Set("Pragma", "no-cache")

				session.Options.MaxAge = -1

				session.Save(c.Request, c.Writer)

				c.Redirect(301, "/")

			}

			//set middleware user datas
			session.Values["userName"] = user.FirstName

			session.Save(c.Request, c.Writer)

			c.Set("username", user.FirstName)

			c.Set("userid", user.Id)

			c.Set("role", user.RoleId)

			c.Set("dataaccess", user.DataAccess)

			controllers.AUTH = spurtcore.NewInstance(&auth.Option{DB: controllers.DB, Token: token, Secret: os.Getenv("JWT_SECRET")})

			controllers.NewAuth.RoleId = user.RoleId

			c.Header("Cache-Control", "no-cache, no-store, must-revalidate")

			c.Header("Pragma", "no-cache")

			c.Header("Expires", "0")

			c.Next()

		}

	}
}

func ModulePermissionByUserId() gin.HandlerFunc {

	return func(c *gin.Context) {

		if c.Request.Method == "GET" || c.Request.Method == "POST" {

			c.Writer.Header().Set("Pragma", "no-cache")

			userid, _ := c.Get("userid")

			roleid, _ := c.Get("role")

			if roleid == 1 {

				c.Next()

				return
			}

			var route []controllers.TblModulePermission

			InitialPermissionsRoute(&route, roleid.(int))

			controllers.PermissionModRoute = []string{}

			for _, val := range route {

				var s = strings.Split(val.RouteName, "/")

				controllers.PermissionModRoute = append(controllers.PermissionModRoute, s[1])

			}

			var role []models.TblUser

			GetModulePermissionByUserId(&role, userid.(int))

			var fullaccess []models.TblUser

			var fa []string

			GetModuleFullAccessPermissionByUserId(&fullaccess, userid.(int))

			for _, val := range fullaccess {

				fa = append(fa, val.RouteName)
				fa = append(fa, val.DisplayName)
			}

			defaultroute := []string{"channel/getcontent", "channel/linktool", "channel/imageupload",
				"/membersgroup/memberisactive", "/membersgroup/memberpopup", "/settings/myprofile",
				"/categories/deletepopup", "/settings/changepassword", "/member/checkemailinmember", "/member/checknumberinmember", "/member/checknameinmember", "/membersgroup/checkmemgroup", "channel/articles", "/media/createfolder", "/media/singlefolder", "/categories/filtercategory", "/categories/checkcatgroup", "/categories/checksubcategoryname", "/membersgroup/checknameinmembergrp", "channel/newentry", "channel/entrylist/", "/settings/data/entrieslist", "/settings/data/export", "/settings/data/importdata"}

			fa = append(fa, defaultroute...)

			RequestURL := strings.Split(c.Request.URL.String(), "?")[0] /*split the request url query params*/

			if strings.Contains(RequestURL, "channels") {

				if !FullAccessPermissionChannelRoute(strings.TrimPrefix(RequestURL, "/"), fa) {

					c.Abort()

					c.Redirect(301, "/403-page")

					return
				}

			} else if strings.Contains(RequestURL, "channel") && !strings.Contains(RequestURL, "settings") { /*entry fullaccess permissions*/

				if !FullAccessPermissionEntryRoute(RequestURL, fa) {

					c.Abort()

					c.Redirect(301, "/403-page")

					return
				}

			} else if strings.Contains(RequestURL, "settings") || strings.Contains(RequestURL, "spaces") { /*settings fullaccess permissions*/


				if !FullAccessPermissionSettingsRoute(strings.TrimPrefix(RequestURL, "/"), fa) {

					c.Abort()

					c.Redirect(301, "/403-page")

					return
				}

			} else if strings.Contains(RequestURL, "memberaccess") {

				if !FullAccessPermissionSettingsRoute(c.Request.URL.String(), fa) {

					c.Abort()

					c.Redirect(301, "/403-page")

					return
				}

			} else {

				var in []string

				var individualaccess []models.TblUser

				GetModuleIndividualAccessPermissionByUserId(&individualaccess, userid.(int))

				for _, val := range individualaccess {

					in = append(in, strings.TrimSuffix(val.RouteName, "/"))
				}

				in = append(in, defaultroute...)

				if !IndividualAccessPermissions(strings.TrimSuffix(RequestURL, "/"), in) {

					c.Abort()

					c.Redirect(301, "/403-page")

					return
				}

			}

			c.Next()

		}
		c.Next()

	}
}

/*Full Access Permission Entry*/
func FullAccessPermissionEntryRoute(url string, fullacc []string) bool {

	var channelAuth channels.Channel

	var pass = false

	re := regexp.MustCompile("[0-9]+")

	if strings.Contains(url, "/channel/editentry") {

		chanellname := strings.Split(url, "/")

		if len(chanellname) != 0 {

			channelAuth.Authority = &controllers.AUTH

			ch, _ := channelAuth.GetchannelByName(strings.ReplaceAll(chanellname[2], "%20", " "))

			for _, val := range fullacc {

				s := re.FindAllString(val, -1)

				if len(s) != 0 {

					n, _ := strconv.Atoi(s[0])

					if ch.Id == n {

						pass = true

						break
					}

				}

			}

		}

	}

	for _, val := range fullacc {

		s := re.FindAllString(val, -1)

		if len(s) != 0 {

			if strings.Contains(strings.ToLower(url), s[0]) {

				pass = true

				break

			}
		} else if strings.EqualFold(val, url) {

			// log.Println("----")

			pass = true

			break

		}

	}

	return pass
}

/*Full Access Permissions Settings*/
func FullAccessPermissionSettingsRoute(url string, fullacc []string) bool {

	var pass = false

	for _, val := range fullacc {

		if strings.Contains(strings.ReplaceAll(strings.ToLower(val), " ", ""), strings.ToLower(url)) || strings.Contains(strings.ToLower(url), strings.ReplaceAll(strings.ToLower(val), " ", "")) && val != "" {

			pass = true

			break

		}
	}

	return pass
}

/*Full Access Permissions channel*/
func FullAccessPermissionChannelRoute(url string, fullacc []string) bool {

	var pass = false

	for _, val := range fullacc {

		if strings.Contains(strings.ReplaceAll(strings.ToLower(val), " ", ""), strings.ToLower(url)) || strings.Contains(strings.ToLower(url), strings.ReplaceAll(strings.ToLower(val), " ", "")) {

			pass = true

			break

		}
	}

	return pass
}

/*Individual */
func IndividualAccessPermissions(url string, route []string) bool {

	var pass = false

	for _, val := range route {

		var furl string

		if strings.Contains(url, "categories") {

			re := regexp.MustCompile(`\d`)

			furl = strings.TrimSuffix(re.ReplaceAllString(url, ""), "/")

		} else {

			furl = url
		}

		if strings.EqualFold(val, furl) {

			pass = true

			break

		}
	}

	return pass

}

var db = config.SetupDB()

/*initial route call*/
func InitialPermissionsRoute(tblmodper *[]controllers.TblModulePermission, roleid int) error {

	if err := db.Table("tbl_module_permissions").Select("route_name,display_name").Joins("inner join tbl_modules on tbl_modules.id=tbl_module_permissions.module_id").Joins("inner join tbl_role_permissions on tbl_role_permissions.permission_id=tbl_module_permissions.id").Where("tbl_role_permissions.role_id=? and tbl_module_permissions.display_name='View' or tbl_module_permissions.display_name='Member Access'", roleid).Find(&tblmodper).Error; err != nil {

		return err
	}

	return nil
}

func GetModulePermissionByUserId(role *[]models.TblUser, id int) error {

	if err := db.Table("tbl_users").Select("mpermission.display_name", "mpermission.module_id", "mpermission.route_name").Joins("inner join tbl_role_permissions on tbl_role_permissions.role_id = tbl_users.role_id").Joins("inner join tbl_module_permissions as mpermission on mpermission.id = tbl_role_permissions.permission_id").Where("tbl_users.id=?", id).Find(&role).Error; err != nil {

		return err

	}
	return nil
}

func GetModuleIndividualAccessPermissionByUserId(role *[]models.TblUser, id int) error {

	if err := db.Debug().Table("tbl_users").Select("mpermission.display_name", "mpermission.module_id", "mpermission.route_name", "mpermission.full_access_permission").Joins("inner join tbl_role_permissions on tbl_role_permissions.role_id = tbl_users.role_id").Joins("inner join tbl_module_permissions  as mpermission on mpermission.id = tbl_role_permissions.permission_id").Where("tbl_users.id=?", id).Find(&role).Error; err != nil {

		return err

	}
	return nil
}

func GetModuleFullAccessPermissionByUserId(role *[]models.TblUser, id int) error {

	if err := db.Table("tbl_users").Select("mpermission.slug_name as display_name", "mpermission.module_id", "mpermission.route_name", "mpermission.full_access_permission").Joins("inner join tbl_role_permissions on tbl_role_permissions.role_id = tbl_users.role_id").Joins("inner join tbl_module_permissions  as mpermission on mpermission.id = tbl_role_permissions.permission_id").Where("mpermission.full_access_permission = 1 and tbl_users.id=?", id).Find(&role).Error; err != nil {

		return err

	}
	return nil
}
