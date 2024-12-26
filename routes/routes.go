package routes

import (
	"net/http"
	"os"
	"path/filepath"
	"spurt-cms/controllers"
	"spurt-cms/lang"
	"spurt-cms/middleware"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func limitRequestBodySize(limit int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, limit)
		if err := c.Request.ParseForm(); err != nil {
			c.Abort()
			return
		}
		c.Next()
	}
}

func SetupRoutes() *gin.Engine {

	r := gin.Default()

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	r.Static("/public", "./public")

	r.Static("/storage", "./storage")

	r.Static("/locales", "./locales")

	r.Use(middleware.CorsMiddleware())

	r.POST("/gqlSaveLocal", controllers.SaveFileInLocal)

	r.POST("/uploadb64image", controllers.EditroImageSave)

	r.POST("/s3upload", controllers.UploadFilesToS3)

	// r.POST("gqlSaveB64Local", controllers.SaveB64InLocal)

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(limitRequestBodySize(20 << 20)) // 20 MB

	r.GET("/image-resize", controllers.ResizeImage)

	r.Use(middleware.CSRFAuth())

	r.MaxMultipartMemory = 8 << 20

	D := r.Group("")

	D.Use(middleware.DashBoardAuth())

	D.GET("/", controllers.WebsiteLogin)

	r.POST("/adminlogin", controllers.CheckLogin)

	D.GET("/forgot", controllers.ForgetPassword)

	D.POST("/reset-password", controllers.SendLinkForForgotPass)

	D.GET("/change-password", controllers.NewPassword)

	D.POST("/set-newpassword", controllers.SetNewPassword)

	r.Use(middleware.JWTAuth())

	r.GET("/lastactive", controllers.LastActive)

	r.GET("/logout", controllers.Logout)

	r.Use(lang.TranslateHandler)

	r.GET("/403-page", controllers.Unauthorized)

	r.GET("/404-page", controllers.FileNotFound)

	r.NoRoute(func(ctx *gin.Context) {

		ctx.Redirect(301, "/404-page")
	})

	/*Cms*/
	C := r.Group("")

	/* Dashboard */

	DH := C.Group("/dashboard")

	DH.GET("/", controllers.DashboardView)

	/*entries module*/
	CE := C.Group("/channel")

	CE.GET("/entrylist/:id", controllers.Entries)

	CE.GET("/newentry", controllers.CreateEntry)

	CE.GET("/entrylist", controllers.AllEntries)

	CE.GET("/deleteentries/", controllers.DeleteEntries)

	CE.POST("/changestatus/:id", controllers.EntryStatus)

	CE.POST("/draftentry/:id", controllers.PublishEntry)

	CE.POST("/publishentry/:id", controllers.PublishEntry)

	CE.GET("/editentry/:channelname/:id", controllers.EditEntry)

	CE.GET("/editentrydetails/:channelname/:id", controllers.EditDetails)

	CE.GET("/editsentry/:id", controllers.EditEntry)

	CE.GET("/editentrydetail/:id", controllers.EditDetails)

	CE.GET("/copyentry/:id", controllers.EditEntry)

	CE.GET("/channelfields/:id", controllers.ChannelFields)

	CE.POST("/imageupload", controllers.ImageUpload)

	CE.POST("/feature", controllers.MakeFeature)

	CE.GET("/memberdetails/", controllers.MemberDetails)

	CE.GET("/userdetails/", controllers.UserDetails)

	CE.POST("/checkentriesorder", controllers.CheckEnriesOrder)

	CE.POST("/deleteselectedentry", controllers.DeleteSelectedEntry)

	CE.POST("/unpublishselectedentry", controllers.UnpublishSelectedEntry)

	CE.GET("/settings", controllers.ChannelSettingView)

	CE.POST("/settings/update", controllers.ChannelSettingUpdate)

	CE.GET("/checkmandatoryfields/:id", controllers.CheckMandatoryFields)

	CE.GET("/previewdetails/:id", controllers.CheckMandatoryFields)

	CE.GET("/unpublishentries", controllers.AllEntries)

	CE.GET("/draftentries", controllers.AllEntries)

	CE.GET("/unpublishentrieslist/:id", controllers.Entries)

	CE.GET("/draftentrieslist/:id", controllers.Entries)

	CE.POST("/entryisactive", controllers.EntryIsActive)

	CE.POST("/entryparentidupdate/:id", controllers.EntryParentIdUpdate)

	CE.POST("/reorder", controllers.EntryReorder)

	CE.POST("/updatepermissionmembergroupid", controllers.UpdateAccPermissionMembergroupId)

	/*channels module*/
	CH := C.Group("/channels")

	CH.GET("/", controllers.ChannelList)

	CH.POST("/newcategory", controllers.CreateCategoryGroup)

	CH.GET("/newchannel", controllers.CreateChannelPage)

	CH.POST("/newchannel", controllers.CreateChannel)

	CH.GET("/editchannel/:id", controllers.EditChannel)

	CH.POST("/updatechannel", controllers.UpdateChannel)

	CH.POST("/getfields", controllers.GetFieldDataByChannelId)

	CH.GET("/deletechannel", controllers.DeleteChannel)

	CH.POST("/status", controllers.ChannelIsActiveStatus)

	CH.POST("/pagination", controllers.PaginationList)

	CH.POST("/Updatechannelfields", controllers.Updatechannelfields)

	CH.GET("/channeltype", controllers.ChannelType)

	/* Category Module*/
	CS := r.Group("/categories")

	CS.GET("/", controllers.CategoryGroupList)

	CS.POST("/newcategory", controllers.CreateCategoryGroup)

	CS.POST("/updatecategory", controllers.UpdateCategoryGroup)

	CS.GET("/deletecategory/:id", controllers.DeleteCategoryGroup)

	CS.POST("/checkcategoryname", controllers.CheckCategoryGroupName)

	CS.GET("/settings", controllers.CategoriesSettings)

	CS.POST("/multiselectcategorygrbdelete", controllers.MultiSelectCategoryGroupDelete)

	// Subcategory List

	CS.GET("/addcategory/:id", controllers.AddCategory)

	CS.POST("/addsubcategory/:categoryid", controllers.AddSubCategory)

	CS.GET("/editsubcategory", controllers.EditSubCategory)

	CS.POST("/editsubcategory", controllers.UpdateSubCategory)

	CS.GET("/removecategory", controllers.DeleteSubCategory)

	// CS.GET("/deletepopup", controllers.CategoryPopup)

	CS.POST("/checksubcategoryname", controllers.CheckSubCategoryName)

	CS.POST("/multiselectcategorydelete", controllers.MultiSelectCategoryDelete)

	/*content access control module*/
	CA := C.Group("/memberaccess")

	CA.GET("/", controllers.ContentAccessControlList)

	CA.GET("/newcontentaccesscontrol", controllers.NewContentAccess)

	// CA.GET("/getpages", controllers.GetPages)

	CA.POST("/grant-accesscontrol", controllers.GrantContentAccessControl)

	CA.GET("/edit-accesscontrol/:accessId", controllers.EditContentAcess)

	CA.GET("/getaccess-pages", controllers.GetContentAccessPages)

	CA.POST("/update-accesscontrol", controllers.UpdateContentAccessControl)

	CA.GET("/delete-accesscontrol/:accessId", controllers.DeleteAccessControl)

	CA.GET("/copy-accesscontrol/:accessId", controllers.EditContentAcess)

	CA.GET("/get-channels", controllers.GetChannelsAndItsEntries)

	/*Member Module*/
	M := C.Group("/member")

	M.GET("/", controllers.MemberList)

	M.POST("/newmember", controllers.CreateMember)

	M.GET("/updatemember/:id", controllers.EditMember)

	M.POST("/updatemember", controllers.UpdateMember)

	M.GET("/deletemember", controllers.DeleteMember)

	M.POST("/checkemailinmember", controllers.CheckEmailInMember)

	M.POST("/checknameinmember", controllers.CheckNameInMember)

	M.POST("/checknumberinmember", controllers.CheckNumberInMember)

	M.POST("/checkprofilesluginmember", controllers.CheckProfileSlugInMember)

	M.GET("/settings", controllers.MemberSettings)

	M.POST("/memberisactive", controllers.MemberStatus)

	M.POST("/deleteselectedmember", controllers.MultiSelectDeleteMember)

	M.POST("/multiselectmemberstatus", controllers.MultiSelectMembersStatus)

	M.POST("/settings/update", controllers.MemberSettingUpdate)

	M.GET("/settings/edittemplate", controllers.EditMemberSettingTemplate)

	M.POST("/settings/updatetemplate", controllers.UpdateMemberSettingTemplate)

	M.POST("/updateactivateclaim", controllers.ActivateClaimStatus)

	M.POST("/getmemberdetails", controllers.GetMemberProfileByMemberId)

	/*Member group Module*/
	MG := C.Group("/membersgroup")

	MG.GET("/", controllers.MemberGroupList)

	MG.POST("/newgroup", controllers.CreateNewMemberGroup)

	MG.POST("/updategroup", controllers.UpdateMemberGroup)

	MG.GET("/deletegroup", controllers.DeleteMemberGroup)

	MG.POST("/groupisactive", controllers.MemberIsActive)

	MG.POST("/checknameinmembergrp", controllers.CheckNameInMemberGroup)

	MG.POST("/deleteselectedmembergroup", controllers.MultiSelectDeleteMembergroup)

	MG.POST("/multiselectmembergroup", controllers.MultiSelectMembersgroupStatus)

	MG.POST("/chkmemgrphavemember", controllers.Chkmemgrphavemember)

	/*Settings Module*/
	S := C.Group("/settings")

	S.GET("/", controllers.SettingView)

	S.GET("/myprofile", controllers.MyProfile)

	S.GET("/changepassword", controllers.ChangePassword)

	S.POST("/updateprofile", controllers.UpdateProfile)

	S.POST("/checkemail", controllers.CheckEmail)

	S.POST("/checknumber", controllers.CheckNumber)

	S.POST("/checkusername", controllers.CheckUsername)

	S.POST("/updatepassword", controllers.UptPassword)

	S.POST("/checkpassword", controllers.CheckPassword)

	/*Roles Module*/
	R := S.Group("/roles")

	R.GET("/", controllers.RoleView)

	R.POST("/checkrole", controllers.CheckRoleAlreadyExists)

	R.POST("/createrole", controllers.RoleCreate)

	R.POST("/updaterole", controllers.RoleUpdate)

	R.POST("/getroledetail", controllers.GetRolePermissionData)

	R.GET("/deleterole", controllers.DeleteRole)

	R.POST("/roleisactive", controllers.RoleIsActive)

	R.POST("/multiselectroledelete", controllers.MultiselectRoleDelete)

	R.POST("/multiselectrolestatus", controllers.MultiSelectRoleStatus)

	R.POST("/chkroleshaveuser", controllers.Chkroleshaveuser)

	/*Language Module*/
	L := S.Group("/languages")

	L.GET("/", controllers.LanguageList)

	// L.POST("/addlanguage", controllers.AddLanguage)

	// L.GET("/editlanguage/:id", controllers.EditLanguage)

	// L.POST("/updatelanguage", controllers.UpdateLanguage)

	// L.GET("/deletelanguage/:id", controllers.DeleteLanguage)

	// L.GET("/downloadlanguage/:id", controllers.DownloadLanguage)

	// L.POST("/languageisactive", controllers.LanguageStatus)

	// L.POST("/checklanguagename", controllers.CheckLanguageName)

	// L.POST("/multiselectlanguagedelete", controllers.MultiselectLanguageDelete)

	// L.POST("/multiselectlanguagestatus", controllers.MultiSelectLanguageStatus)

	L.POST("/setdefaultlanguage", controllers.SetDefaultLanguage)

	/*User Module*/
	U := S.Group("/users")

	U.GET("/", controllers.Userlist)

	U.POST("/createuser", controllers.CreateUser)

	U.GET("/delete-user/:id", controllers.DeleteUser)

	U.GET("/edit-user", controllers.EditUser)

	U.POST("/update-user", controllers.UpdateUser)

	U.POST("/checkuserdata", controllers.CheckUserData)

	U.POST("/checkemail", controllers.CheckEmail)

	U.POST("/checknumber", controllers.CheckNumber)

	U.POST("/checkusername", controllers.CheckUsername)

	U.POST("/deleteselectedusers", controllers.DeleteMultipleUser)

	U.POST("/selectedusersaccesschange", controllers.SelectedUsersAccessChange)

	U.POST("/changeActiveStatus", controllers.ChangeActiveStatus)

	U.POST("/selectedUserStatusChange", controllers.SelectedUsersStatusChange)

	/*Email Templates*/
	ET := S.Group("/emails")

	ET.GET("/", controllers.EmailTemplate)

	ET.GET("/edit-template", controllers.EditTemplate)

	ET.POST("/update-temp", controllers.UpdateTemplate)

	ET.POST("/templateisactive", controllers.TempIsActive)

	ET.GET("/emailconfig/", controllers.EmailConfig)

	ET.POST("/update", controllers.CreateEmailConfig)

	/*General Settings*/
	GS := S.Group("/general-settings")

	GS.GET("/", controllers.GeneralSettings)

	GS.POST("/update", controllers.UpdateGeneralSettings)

	/*System Module*/
	SS := S.Group("/personalize")

	SS.GET("/", controllers.Personalization)

	SS.POST("/update", controllers.PersonalizationUpdate)

	/*Data Module*/
	DT := S.Group("/data")

	DT.GET("/", controllers.Data)

	DT.GET("/entrieslist/:id", controllers.Entrieslist)

	DT.GET("/export", controllers.Export)

	DT.POST("/importdata", controllers.ImportData)

	DT.GET("/err-xlsx-download", controllers.DownloaderrXlsx)

	DT.GET("/secondimportpage", controllers.ImportPageTwo)

	DT.GET("/thirdimportpage", controllers.ImportPageThree)

	DT.GET("/exportpage", controllers.ExportPage)

	// GR.GET("/", controllers.GraphqlTokenApi)

	// GR.GET("/create", controllers.CreateTokenGrapqhl)

	// GR.POST("/create", controllers.CreateToken)

	// GR.GET("/edit/:id", controllers.UpdateGraphqlApi)

	// GR.POST("/update", controllers.UpdateGraphqlToken)

	// GR.GET("/delete/:id", controllers.DeleteToken)

	// GR.POST("/multideletegraphtoken", controllers.MultiDeleteGraphqlToken)

	G := C.Group("/graphql")

	G.GET("/", controllers.GraphqlTokenApi)

	G.POST("/create", controllers.CreateToken)

	G.GET("/delete/:id", controllers.DeleteToken)

	G.GET("/edit/:id", controllers.UpdateGraphqlApi)

	G.POST("/update", controllers.UpdateGraphqlToken)

	G.POST("/multideletetokens", controllers.MultiDeleteGraphqlToken)

	G.GET("/playground", controllers.PlaygroundView)

	T := C.Group("/templates")

	T.GET("/", controllers.ListTemplates)

	return r

}
