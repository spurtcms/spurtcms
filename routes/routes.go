package routes

import (
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

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.GET("/image-resize", controllers.ResizeImage)

	r.Use(middleware.CSRFAuth())

	r.MaxMultipartMemory = 8 << 20

	r.Use(middleware.CorsMiddleware())

	D := r.Group("")

	D.Use(middleware.DashBoardAuth())

	D.GET("/", controllers.Login)

	r.POST("/", controllers.CheckLogin)

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

	/*channels module*/
	CH := C.Group("/channels", middleware.ModulePermissionByUserId())

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

	CS.GET("/deletepopup", controllers.CategoryPopup)

	CS.POST("/checksubcategoryname", controllers.CheckSubCategoryName)

	CS.POST("/multiselectcategorydelete", controllers.MultiSelectCategoryDelete)

	/*content access control module*/
	CA := C.Group("/memberaccess", middleware.ModulePermissionByUserId())

	CA.GET("/", controllers.ContentAccessControlList)

	CA.GET("/newcontentaccesscontrol", controllers.NewContentAccess)

	CA.GET("/getpages", controllers.GetPages)

	CA.POST("/grant-accesscontrol", controllers.GrantContentAccessControl)

	CA.GET("/edit-accesscontrol/:accessId", controllers.EditContentAcess)

	CA.GET("/getaccess-pages", controllers.GetContentAccessPages)

	CA.POST("/update-accesscontrol", controllers.UpdateContentAccessControl)

	CA.GET("/delete-accesscontrol/:accessId", controllers.DeleteAccessControl)

	CA.GET("/copy-accesscontrol/:accessId", controllers.EditContentAcess)

	CA.GET("/get-channels", controllers.GetChannelsAndItsEntries)

	/*Member Module*/
	M := C.Group("/member")

	M.GET("/", middleware.ModulePermissionByUserId(), controllers.MemberList)

	M.POST("/newmember", middleware.ModulePermissionByUserId(), controllers.CreateMember)

	M.GET("/updatemember/:id", controllers.EditMember)

	M.POST("/updatemember", middleware.ModulePermissionByUserId(), controllers.UpdateMember)

	M.GET("/deletemember", middleware.ModulePermissionByUserId(), controllers.DeleteMember)

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

	/*Member group Module*/
	MG := C.Group("/membersgroup")

	MG.GET("/", middleware.ModulePermissionByUserId(), controllers.MemberGroupList)

	MG.POST("/newgroup", middleware.ModulePermissionByUserId(), controllers.CreateNewMemberGroup)

	MG.POST("/updategroup", middleware.ModulePermissionByUserId(), controllers.UpdateMemberGroup)

	MG.GET("/deletegroup", middleware.ModulePermissionByUserId(), controllers.DeleteMemberGroup)

	MG.POST("/groupisactive", controllers.MemberIsActive)

	MG.POST("/checknameinmembergrp", controllers.CheckNameInMemberGroup)

	MG.POST("/deleteselectedmembergroup", controllers.MultiSelectDeleteMembergroup)

	MG.POST("/multiselectmembergroup", controllers.MultiSelectMembersgroupStatus)

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

	R.GET("/",  controllers.RoleView)

	R.POST("/checkrole", controllers.CheckRoleAlreadyExists)

	R.POST("/createrole",  controllers.RoleCreate)

	R.POST("/updaterole",  controllers.RoleUpdate)

	R.POST("/getroledetail", controllers.GetRolePermissionData)

	R.GET("/deleterole", controllers.DeleteRole)

	R.POST("/roleisactive", controllers.RoleIsActive)

	R.POST("/multiselectroledelete", controllers.MultiselectRoleDelete)

	R.POST("/multiselectrolestatus", controllers.MultiSelectRoleStatus)

	/*Language Module*/
	L := S.Group("/languages")

	L.GET("/", middleware.ModulePermissionByUserId(), controllers.LanguageList)

	L.POST("/addlanguage", middleware.ModulePermissionByUserId(), controllers.AddLanguage)

	L.GET("/editlanguage/:id", controllers.EditLanguage)

	L.POST("/updatelanguage", middleware.ModulePermissionByUserId(), controllers.UpdateLanguage)

	L.GET("/deletelanguage/:id", middleware.ModulePermissionByUserId(), controllers.DeleteLanguage)

	L.GET("/downloadlanguage/:id", controllers.DownloadLanguage)

	L.POST("/languageisactive", controllers.LanguageStatus)

	L.POST("/checklanguagename", controllers.CheckLanguageName)

	L.POST("/multiselectlanguagedelete", controllers.MultiselectLanguageDelete)

	L.POST("/multiselectlanguagestatus", controllers.MultiSelectLanguageStatus)

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
	ET := S.Group("/emails", middleware.ModulePermissionByUserId())

	ET.GET("/", controllers.EmailTemplate)

	ET.GET("/edit-template", controllers.EditTemplate)

	ET.POST("/update-temp", controllers.UpdateTemplate)

	ET.POST("/templateisactive", controllers.TempIsActive)

	ET.GET("/emailconfig/", controllers.EmailConfig)

	ET.POST("/update", controllers.CreateEmailConfig)

	/*General Settings*/
	GS := S.Group("/general-settings", middleware.ModulePermissionByUserId())

	GS.GET("/", controllers.GeneralSettings)

	GS.POST("/update", controllers.UpdateGeneralSettings)

	/*Graphql settings*/

	GR := S.Group("/graphql", middleware.ModulePermissionByUserId())

	GR.GET("/", controllers.GraphqlTokenApi)

	GR.GET("/create", controllers.CreateTokenGrapqhl)

	/*Media Library*/
	ME := r.Group("/media")

	ME.GET("/", controllers.MediaList)

	ME.POST("/createfolder", controllers.AddFolder)

	ME.POST("/uploadimage", controllers.UploadImageMedia)

	ME.POST("/deletefolfile", controllers.DeleteFolderFile)

	ME.POST("/singlefolder", controllers.FolderDetails)

	ME.POST("/upload", controllers.UploadImage)

	ME.GET("/settings", controllers.MediaSettings)

	ME.POST("/settings/update", controllers.MediaStorageUpdate)

	/*System Module*/
	SS := S.Group("/personalize")

	SS.GET("/", controllers.Personalization)

	SS.POST("/update", controllers.PersonalizationUpdate)

	/*Data Module*/
	DT := S.Group("/data", middleware.ModulePermissionByUserId())

	DT.GET("/", controllers.Data)

	DT.GET("/entrieslist/:id", controllers.Entrieslist)

	DT.GET("/export", controllers.Export)

	DT.POST("/importdata", controllers.ImportData)

	DT.GET("/err-xlsx-download", controllers.DownloaderrXlsx)

	// ECO := S.Group("/email-config")

	// ECO.GET("/", controllers.EmailConfig)

	// ECO.POST("/update", controllers.CreateEmailConfig)

	return r

}
