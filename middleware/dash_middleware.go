package middleware

import (
	"fmt"
	"net/http"
	"os"
	"spurt-cms/controllers"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func DashBoardAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		session, _ := controllers.Store.Get(c.Request, os.Getenv("SESSION_KEY"))

		tkn := session.Values["token"]

		if tkn != "" && tkn != nil {

			tkn := session.Values["token"].(string)

			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(tkn, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil || !token.Valid {
				c.Redirect(http.StatusFound, "/")
				return
			}

			userIDFloat, ok := claims["user_id"].(float64)
			if !ok {

				c.Redirect(http.StatusFound, "/")
				return
			}
			userID := int(userIDFloat)

			user, _, _ := controllers.NewTeamWP.GetUserById(userID, []int{})

			host := c.Request.Host

			var redirectURL string

			if host == "spurtcms.com" {

				redirectURL = fmt.Sprintf("https://%s.spurtcms.com/admin/dashboard", user.FirstName+strconv.Itoa(user.Id))

			} else {
				redirectURL = fmt.Sprintf("http://%s.lvh.me:8080/admin/dashboard", user.Subdomain)
			}

			c.Writer.Header().Set("Pragma", "no-cache")

			c.Redirect(301, redirectURL)

			return
		}

	}
}
