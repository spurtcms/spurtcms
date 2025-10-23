package middleware

import (
	"fmt"
	"os"
	"spurt-cms/controllers"
	"time"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func CSRFAuth() gin.HandlerFunc {

	return csrf.Middleware(csrf.Options{

		Secret:        fmt.Sprintf("%8d", time.Now().Unix()),
		IgnoreMethods: []string{"GET"},
		ErrorFunc: func(c *gin.Context) {

			csrftoken := c.PostForm("csrf")

			if csrftoken != csrf.GetToken(c) {

				sess, err := controllers.Store.Get(c.Request, os.Getenv("SESSION_KEY"))

				if err != nil {
					fmt.Println(err)
				}
				sess.Values["token"] = ""

				sess.Options.MaxAge = -1

				er := sess.Save(c.Request, c.Writer)
				if er != nil {
					fmt.Println(er)
				}
				c.Abort()

				c.Writer.Header().Set("Pragma", "no-cache")

				c.SetCookie("Alert-msg", "Internal Server Error", 3600, "", "", false, false)

				fmt.Println("csrf token mismatch")

				c.Redirect(301, "/")

			} else {

				c.Next()
			}

		},
	})
}

func CorsMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return
		}

		c.Next()
	}
}
