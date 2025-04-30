package middleware

import (
	"os"
	"spurt-cms/controllers"

	"github.com/gin-gonic/gin"
)

func DashBoardAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		session, _ := controllers.Store.Get(c.Request, os.Getenv("SESSION_KEY"))

		token := session.Values["token"]

		if token != nil {
			
			c.Writer.Header().Set("Pragma", "no-cache")

			c.Redirect(301, "/dashboard")

			return
		}

	}
}
