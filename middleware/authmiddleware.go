package middleware

import (
	"fmt"
	"os"
	"spurt-cms/config"
	"spurt-cms/controllers"
	"spurt-cms/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

			session.Options.MaxAge = 60 * 60 * 24 * 90 //90 days

			session.Save(c.Request, c.Writer)

			token := tkn.(string)

			userid, _, err := controllers.NewAuth.VerifyToken(token, os.Getenv("JWT_SECRET"))

			if err != nil {
				controllers.ErrorLog.Println("err", err)
			}

			Claims := jwt.MapClaims{}
			tkn, err := jwt.ParseWithClaims(token, Claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
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

			controllers.Token = token

			controllers.NewTeam.Auth.UserId = userid

			user, _, _ := controllers.NewTeamWP.GetUserById(userid, []int{})

			if user.Id == 0 {

				c.Abort()

				c.Writer.Header().Set("Pragma", "no-cache")

				session.Options.MaxAge = -1

				session.Save(c.Request, c.Writer)

				c.Redirect(301, "/")

			}

			c.Set("tenant_id", user.TenantId)

			c.Set("userDetails", user)

			controllers.Tenant_Id(c)

			//set middleware user datas
			session.Values["userName"] = user.FirstName

			session.Save(c.Request, c.Writer)

			c.Set("username", user.FirstName)

			c.Set("userid", user.Id)

			c.Set("role", user.RoleId)

			c.Set("dataaccess", user.DataAccess)

			controllers.NewAuth.RoleId = user.RoleId

			c.Header("Cache-Control", "no-cache, no-store, must-revalidate")

			c.Header("Pragma", "no-cache")

			c.Header("Expires", "0")

			c.Next()

		}

	}
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
