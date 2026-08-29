package routes

import (
	"fmt"
	"net/http"
	"spurt-cms/controllers"
	"spurt-cms/middleware"

	"strings"

	"github.com/gin-gonic/gin"
)

func TemplateRoutes(r *gin.Engine) {

	//templates images load..//

	r.Static("/websites/public", "./websites/public")

	r.Static("/websites/common/assets", "./websites/common/assets")

	r.Static("/websites/themes/content_verse/assets", "./websites/themes/content_verse/assets")

	r.Static("/support_sphere/assets", "./websites/themes/support_sphere/assets")

	TD := r.Group("")

	TD.Use(middleware.TemplateDashBoardAuth())

	// ==================================================================================

	// 🔹 Global Rewrite Middleware for routes
	r.Use(RewriteMiddleware())

	// Common Channel and Pages List Route
	r.GET("/:slug", RewriteSlugMiddleware())

	r.GET("/:slug/category/:categoryname", RewriteChannelMiddleware(controllers.CategoryBaseEntryList))
	r.GET("/:slug/:entryslug/:subpage", controllers.LoadPages)
	r.GET("/:slug/:entryslug", RewriteDetailSlugMiddleware())

	//Channel Routes

	r.GET("/channel/:cname", controllers.ChannelEntriesList)

	r.GET("/channel/:cname/:entryslug", controllers.EntryDetailsPage)

	//Categories Routes

	r.GET("/channel/:cname/category/:categoryname", controllers.CategoryBaseEntryList)

	// r.GET("/categories/:menuname", controllers.CategoryEntryList)

	// r.GET("/categories/:menuname/*dynamicname", controllers.CategoryDetailsPage)

	r.GET("/pages/:pagename", controllers.LoadPages)

	r.GET("/menuitems/:id", controllers.MenuIemsListForWebsite)

	// r.GET("/entries/:entryname", controllers.StaticEntryData)

	r.GET("/membership", controllers.MembershipList)

	r.GET("/membership/:id", controllers.MembershiDetail)

	r.GET("/signin", controllers.Signin)

	r.GET("/signup", controllers.SignUp)

	r.GET("/forget-password", controllers.ForgotPassword)

	r.POST("/forget", controllers.SendLinkForForgotPass)

	r.POST("/sign-in", controllers.SignIn)

	r.POST("/signup", controllers.NewUserSignUp)

	r.POST("/checknameinmember", controllers.CheckNameInUser)

	r.POST("/checkemailinmember", controllers.TemplateCheckEmailInMember)

	r.POST("/checknumberinmember", controllers.TemplateCheckNumberInMember)

	r.POST("/updateprofile", controllers.Updateprofile)

	r.GET("/change-password", controllers.TemplateChangePassword)

	r.POST("/set-newpassword", controllers.SetNewpassword)

	r.GET("/search", controllers.GlobalSearch)

	Tr := r.Group("")

	Tr.Use(middleware.TemplateJWTAuth())

	Tr.GET("/myprofile", controllers.TemplateMyProfile)

	Tr.POST("/update-profile", controllers.UpdateMember)

	Tr.GET("/logout", controllers.TemplateLogout)

	// Templateroutes end//

}

func RewriteSlugMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		slug := c.Param("slug")

		if slug == "" {
			c.Redirect(301, "/404-TemplatePage")
			return
		}

		result, err := controllers.GetRouteSlugFormDB(slug)

		fmt.Println("result", result)

		if err != nil {
			c.Redirect(301, "/404-TemplatePage")
			return
		}

		switch result.ModuleName {

		case "Channel":
			c.Request.URL.Path = "/channel/" + slug
			controllers.ChannelEntriesList(c)
			return

		case "Pages":
			c.Request.URL.Path = "/pages/" + slug
			controllers.LoadPages(c)
			return

		default:
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
	}
}
func RewriteDetailSlugMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		firstslug := c.Param("slug")

		slug := c.Param("entryslug")

		if slug == "" {

			c.Redirect(301, "/404-TemplatePage")
			return

		}

		result, err := controllers.GetRouteSlugFormDB(slug)

		if err != nil {
			c.Redirect(301, "/404-TemplatePage")
			return
		}

		switch result.ModuleName {

		case "Entries":
			c.Request.URL.Path = "/channel/" + firstslug + "/" + slug
			controllers.EntryDetailsPage(c)
			return

		case "Pages":
			c.Request.URL.Path = "/" + firstslug + "/" + slug
			controllers.LoadPages(c)
			return

		default:
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
	}
}
func RewriteMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		// --------------------
		// CHANNEL REDIRECT
		// --------------------
		if strings.HasPrefix(path, "/channel/") {
			cleanPath := strings.TrimPrefix(path, "/channel")
			if cleanPath == "" {
				cleanPath = "/"
			}
			c.Redirect(http.StatusMovedPermanently, cleanPath)
			c.Abort()
			return
		}

		// --------------------
		// PAGES REDIRECT (NEW)
		// --------------------
		if path == "/pages" || strings.HasPrefix(path, "/pages/") {
			cleanPath := strings.TrimPrefix(path, "/pages")
			if cleanPath == "" {
				cleanPath = "/"
			}
			c.Redirect(http.StatusMovedPermanently, cleanPath)
			c.Abort()
			return
		}

		c.Next()
	}
}

func RewriteChannelMiddleware(targetHandler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")

		if slug == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "slug required"})
			return
		}

		// Rewrite path to match original channel route
		originalPath := "/channel/" + slug + c.Request.URL.Path[len(slug)+1:]

		c.Request.URL.Path = originalPath

		targetHandler(c)
	}
}

func RewritePageMiddleware(targetHandler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {

		path := c.Request.URL.Path

		trimmed := strings.TrimPrefix(path, "/pages")

		if trimmed == "" {
			trimmed = "/"
		}

		originalPath := "/pages" + trimmed
		c.Request.URL.Path = originalPath
		targetHandler(c)
	}
}
