package templateroutes

import (
	"os"
	"path/filepath"
	templatecontroller "spurt-cms/websites/controller"
	"spurt-cms/websites/middleware"
	"strings"

	"github.com/gin-gonic/gin"
)

func TemplateRoutes(r *gin.Engine) {

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})

	r.LoadHTMLFiles(htmlfiles...)

	r.Static("/websites/public", "./websites/public")

	r.Static("/websites/common/assets", "./websites/common/assets")

	r.Static("/websites/themes/content_verse/assets", "./websites/themes/content_verse/assets")

	r.Static("/ai_templates/assets", "./websites/themes/ai_templates/assets")

	r.Static("/support_sphere/assets", "./websites/themes/support_sphere/assets")

	r.Static("/courses/assets", "./websites/themes/courses/assets")

	r.Static("/jobs/assets", "./websites/themes/jobs/assets")

	// Digital products Templates APIs Start

	r.GET("/login", templatecontroller.Login)

	r.POST("/emailverification", templatecontroller.EmailVerification)

	r.POST("/otpverification", templatecontroller.OtpVerification)

	r.Use(middleware.TemplateJWTAuth())

	r.GET("/listings/:menuname", templatecontroller.ListingsList)

	r.GET("/listings/:menuname/:listingname", templatecontroller.ListingsDetailsPage)

	r.GET("/profile", templatecontroller.MyProfiles)

	r.POST("/profile", templatecontroller.UpdateprofileData)

	r.GET("/downloads", templatecontroller.MyDownloads)

	// end

	r.GET("/channel/:cname", templatecontroller.ChannelEntriesList)

	r.GET("/channel/:cname/*dynamicname", templatecontroller.EntryDetailsPage)

	r.GET("/pages/:pagename", templatecontroller.StaticPageData)

	r.GET("/entries/:entryname", templatecontroller.StaticEntryData)

	r.GET("/courses/:coursename/:uuid", templatecontroller.CoursesDetails)

	r.GET("/courses/:coursename/:uuid/:lessonid", templatecontroller.CoursesDetails)

	r.GET("/membership", templatecontroller.MembershipList)

	r.GET("/membership/:id", templatecontroller.MembershiDetail)

	r.GET("/signin", templatecontroller.Signin)

	r.GET("/signup", templatecontroller.SignUp)

	r.GET("/forget-password", templatecontroller.ForgotPassword)

	r.POST("/forget", templatecontroller.SendLinkForForgotPass)

	r.GET("/myprofile", templatecontroller.MyProfile)

	r.POST("/singin", templatecontroller.SignIn)

	r.POST("/signup", templatecontroller.NewUserSignUp)

	r.POST("/checknameinmember", templatecontroller.CheckNameInUser)

	r.POST("/checkemailinmember", templatecontroller.CheckEmailInMember)

	r.POST("/checknumberinmember", templatecontroller.CheckNumberInMember)

	r.POST("/updateprofile", templatecontroller.Updateprofile)

	r.GET("/change-password", templatecontroller.ChangePassword)

	r.POST("/set-newpassword", templatecontroller.SetNewpassword)

	r.GET("/logout", templatecontroller.Logout)

	r.NoRoute(func(ctx *gin.Context) {

		ctx.Redirect(301, "nodata.html")
	})

}
