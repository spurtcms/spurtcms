package middleware

import "github.com/gin-gonic/gin"

func TemplateViewAuth() gin.HandlerFunc{

	return func(c *gin.Context) {

		c.Next()
	}
}