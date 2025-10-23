package middleware

import (
	"fmt"
	"os"
	"spurt-cms/controllers"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func TemplateJWTAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Header("Cache-Control", "no-store, no-cache, must-revalidate, private")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")

		session, _ := controllers.Store.Get(c.Request, os.Getenv("TEMPLATE_SESSION_KEY"))

		tkn := session.Values["token"]

		if tkn != nil {

			token := tkn.(string)

			member_id, err := VerifyToken(token, os.Getenv("JWT_SECRET"))

			if err != nil {
				controllers.ErrorLog.Println("err", err)
			}

			c.Set("member_id", member_id)
		}
	}
}

func VerifyToken(token string, secret string) (memberid int, err error) {

	Claims := jwt.MapClaims{}

	tkn, err := jwt.ParseWithClaims(token, Claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println(err)
			return 0, err
		}

		return 0, err
	}

	if !tkn.Valid {
		fmt.Println(tkn)
		return 0, err
	}

	member_id := Claims["member_id"]

	return int(member_id.(float64)), nil
}
