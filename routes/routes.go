package routes

import (
	"fmt"
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

	fmt.Println("htmlfiles::", htmlfiles)
	r.LoadHTMLFiles(htmlfiles...)

	r.Static("/public", "./public")

	r.Static("/storage", "./storage")

	r.Static("/locales", "./locales")

	r.POST("/webhook/stripe", controllers.HandleStripeWebhook)

	r.Use(middleware.CorsMiddleware())

	r.POST("/gqlSaveLocal", controllers.SaveFileInLocal)

	r.POST("/uploadb64image", controllers.EditroImageSave)

	r.POST("/uploadb64audio", controllers.EditroImageSave)

	r.POST("/uploadb64document", controllers.EditroImageSave)

	r.POST("/uploadform64image", controllers.EditroImageSave)

	r.POST("/s3upload", controllers.UploadFilesToS3)

	// r.POST("gqlSaveB64Local", controllers.SaveB64InLocal)

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(limitRequestBodySize(20 << 20)) // 20 MB

	r.GET("/image-resize", controllers.ResizeImage)

	r.GET("/audio-resize", controllers.AudioResize)

	r.GET("/document-access", controllers.DocumentResize)

	r.GET("/sitemap.xml", controllers.Sitemap)

	r.Use(middleware.CSRFAuth())

	TemplateRoutes(r)

	r.MaxMultipartMemory = 8 << 20

	D := r.Group("")

	// D.Use(middleware.DashBoardAuth())
	D.GET("/", controllers.NewWebsitePage) // website landding route

	// D.GET("/s/*dynamicname",controllers.StaticPage)

	// if (os.Getenv("TENANTID")) == "1" {

	D.GET("/admin", controllers.Login)

	r.POST("/adminlogin", controllers.CheckLogin)

	D.GET("/admin/forgot", controllers.ForgetPassword)

	D.POST("/admin/reset-password", controllers.SendLinkForForgotPass)

	D.GET("/admin/change-password", controllers.NewPassword)

	D.POST("/admin/set-newpassword", controllers.SetNewPassword)

	r.Use(middleware.JWTAuth())

	r.GET("/lastactive", controllers.LastActive)

	r.GET("/admin/logout", controllers.Logout)

	r.Use(lang.TranslateHandler)

	r.GET("/403-page", controllers.Unauthorized)

	r.GET("/404-page", controllers.FileNotFound)

	r.GET("/404-TemplatePage", controllers.FileNotFoundForTemplates)

	r.NoRoute(func(ctx *gin.Context) {

		urlPath := ctx.Request.URL.Path

		if strings.Contains(urlPath, "admin") {

			ctx.Redirect(301, "/404-page")

		} else if !strings.Contains(urlPath, "channel") || !strings.Contains(urlPath, "listings") || !strings.Contains(urlPath, "pages") || !strings.Contains(urlPath, "entries") || !strings.Contains(urlPath, "courses") {

			ctx.Redirect(301, "/404-TemplatePage")
		}
	})

	/*Cms*/
	C := r.Group("/admin")

	/* Dashboard */

	C.POST("/checkduplicateroute", controllers.CheckDupliateRoute)

	DH := C.Group("/dashboard")

	DH.GET("/", controllers.DashboardView)

	/*entries module*/
	CE := C.Group("/entries")

	CE.GET("/entrylist/:id", controllers.Entries)

	CE.GET("/create/:id", controllers.CreateEntry)

	CE.GET("/entrylist", controllers.AllEntries)

	CE.GET("/form-detail", controllers.Formdetail)

	CE.GET("/deleteentries/", controllers.DeleteEntries)

	CE.POST("/changestatus/:id", controllers.EntryStatus)

	CE.POST("/draftentry/:id", controllers.PublishEntry)

	CE.POST("/publishentry/:id", controllers.PublishEntry)

	CE.GET("/edit/:channelname/:id", controllers.EditEntry)

	CE.GET("/editentrydetails/:channelname/:id", controllers.EditDetails)

	CE.GET("/edits/:id", controllers.EditEntry)

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

	CE.GET("/ctadetails", controllers.FormDetails)

	Bt := C.Group("/browsetheme")

	Bt.GET("", controllers.BrowseTheme)

	Bt.GET("/configure/:template", controllers.Configuration)

	Bt.GET("/templatesetting/:template", controllers.Templatesetting)

	// Bt.GET("/datasetting/:template", controllers.DataSettinng)

	// AI assistance

	AI := C.Group("/aiassistance")

	AI.GET("/", controllers.Aicontents)

	AI.POST("/receivetopicdata", controllers.Receivetopicdata)

	AI.POST("/receiveproductdata", controllers.Receiveproductdata)

	AI.POST("/generatetopics", controllers.Generatetopics)

	AI.POST("/generatearticle", controllers.Generatearticle)

	AI.POST("/generate-ecom-article", controllers.GenerateEcommercearticle)
	// Form Builder

	// FB := C.Group("/cta")

	// FB.GET("/", controllers.FormbuilderList)

	// FB.GET("/form-responses", controllers.FormResponses)

	// FB.GET("/response-detail/:ticket", controllers.ResponseDetail)

	// FB.GET("/close-ticket/:ticket", controllers.CloseTicket)

	// FB.POST("/replyforresponse", controllers.ReplyForResponses)

	// FB.GET("/form-detail/:id", controllers.Formdetail)

	// FB.GET("/create", controllers.AddForms)

	// FB.GET("/unpublished", controllers.FormbuilderList)

	// FB.GET("/draft", controllers.FormbuilderList)

	// FB.POST("/createforms", controllers.CreateForms)

	// FB.GET("/formstatus", controllers.Status)

	// FB.GET("/deleteform", controllers.DeleteForm)

	// FB.GET("/edit/:id", controllers.FormEdit)

	// FB.POST("/updateforms", controllers.FormUpdate)

	// FB.POST("/multiselectformdelete", controllers.MultiDelete)

	// FB.POST("/multiselectstatuschange", controllers.MultiSelectStatusChange)

	// FB.GET("/formduplicate/:id", controllers.FormDuplicate)

	// FB.POST("/isactive", controllers.FormIsactive)

	// FB.GET("/defaultctalist", controllers.DefaultCtaList)

	// FB.POST("/addtomycollection", controllers.AddCollection)

	// FB.GET("/removecta/:id", controllers.RemoveCta)

	// FB.GET("/ctapreview/:id", controllers.CtaPreview)

	/* Get Support Module*/
	// SP := C.Group("/get-spurtcmspro")

	// SP.GET("", controllers.Getsupport)

	// SP.POST("/payment", controllers.MakeStripePayment)

	// SP.GET("/supportsubmited", controllers.Supportsubmited)

	/*channels module*/
	CH := C.Group("/channels")

	CH.GET("/", controllers.ChannelList)

	CH.POST("/newcategory", controllers.CreateCategoryGroup)

	CH.GET("/create", controllers.CreateChannelPage)

	CH.POST("/create", controllers.CreateChannel)

	CH.GET("/edit/:id", controllers.EditChannel)

	CH.POST("/updatechannel", controllers.UpdateChannel)

	CH.POST("/getfields", controllers.GetFieldDataByChannelId)

	CH.GET("/deletechannel", controllers.DeleteChannel)

	CH.POST("/status", controllers.ChannelIsActiveStatus)

	CH.POST("/pagination", controllers.PaginationList)

	CH.POST("/Updatechannelfields", controllers.Updatechannelfields)

	CH.GET("/channeltype", controllers.ChannelType)

	CH.GET("/defaultchannels", controllers.DefaultChannelList)

	CH.GET("/addtomycollection/:id", controllers.AddChannelTomyCollection)

	CH.GET("/clone/:id", controllers.CloneChannels)

	CH.POST("/checktitle", controllers.CheckChannelName)

	/* Category Module*/
	CS := C.Group("/categories")

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
	M := C.Group("/user")

	M.GET("/", controllers.MemberList)

	M.POST("/newmember", controllers.CreateMember)

	M.GET("/edit/:id", controllers.EditMember)

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
	MG := C.Group("/usergroup")

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

	// /*Graphql settings*/

	// GR := S.Group("/graphql")

	// GR.GET("/", controllers.GraphqlTokenApi)

	// GR.GET("/create", controllers.CreateTokenGrapqhl)

	// GR.POST("/create", controllers.CreateToken)

	// GR.GET("/edit/:id", controllers.UpdateGraphqlApi)

	// GR.POST("/update", controllers.UpdateGraphqlToken)

	// GR.GET("/delete/:id", controllers.DeleteToken)

	// GR.POST("/multideletegraphtoken", controllers.MultiDeleteGraphqlToken)

	/*Media Library*/
	ME := C.Group("/media")

	ME.GET("/", controllers.MediaList)

	ME.POST("/createfolder", controllers.AddFolder)

	ME.POST("/uploadimage", controllers.UploadImageMedia)

	ME.POST("/deletefolfile", controllers.DeleteFolderFile)

	ME.POST("/singlefolder", controllers.FolderDetails)

	ME.POST("/upload", controllers.UploadImage)

	ME.POST("/loadmore", controllers.LoadMore)

	ME.GET("/settings", controllers.MediaSettings)

	ME.POST("/settings/update", controllers.MediaStorageUpdate)

	ME.POST("/renamemediapath", controllers.RenameMediaPath)

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

	// AI Settings

	AIS := S.Group("/aisettings")

	AIS.GET("/", controllers.AiSettings)

	AIS.POST("/create", controllers.CreateAIApiKeySetting)

	AIS.POST("/update", controllers.UpdateAIApiKeySetting)

	AIS.GET("/delete/:id", controllers.DeleteModule)

	AIS.POST("/multiselectdelete", controllers.MultiselectDeleteModule)

	AIS.POST("/status", controllers.StatusAction)

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

	// T.GET("/:id", controllers.Templates)

	B := C.Group("/blocks")

	B.GET("", controllers.BlockList)

	B.GET("/collection", controllers.CollectionList)

	B.GET("/defaultlist", controllers.DefaultBlockList)

	B.POST("/create", controllers.CreateBlock)

	B.GET("/blockdetails", controllers.EditBlock)

	B.POST("/update", controllers.UpdateBlock)

	B.GET("/deleteblock", controllers.DeleteBlock)

	B.POST("/checktitle", controllers.CheckTitleInBlock)

	B.POST("/blockisactive", controllers.BlockIsActive)

	B.POST("/createcollection", controllers.CreateCollection)

	B.GET("/removecollection", controllers.CollectionRemove)

	B.POST("/addtomycollection", controllers.Addtomycollection)

	// paid membership pro routes

	PM := C.Group("/membership")

	PM.GET("/", controllers.MembershipMemberList)

	PM.POST("/newmember", controllers.MembershipCreateMember)

	PM.GET("/EditMember", controllers.MembershipEditMember)

	PM.POST("/updatemember", controllers.MembershipUpdateMember)

	PM.GET("/deletemembershipmember", controllers.MembershipDeleteMember)

	PM.POST("/multiselectdeletemember", controllers.MultiselectDeleteMember)

	PM.POST("/membershipisactive", controllers.MembershipIsactive)

	PML := C.Group("/membershiplevel")

	PML.GET("/", controllers.Membership)

	PML.GET("/newsubscriptionlevel", controllers.Newsubscriptionlevel)

	PML.GET("/loadmembershiptemplate", controllers.LoadMembershipTemplate)

	PML.POST("/createmembershiplevel", controllers.CreateMembershipLevels)

	PML.GET("/editmembershiplevel", controllers.Editmembershiplevel)

	PML.POST("/updatembershiplevel", controllers.UpdateSubscriptionLevel)

	PML.GET("/deletemembershiplevel", controllers.DeleteSubscriptionLevel)

	PML.POST("/multiselectdeletemembershiplevel", controllers.MultiselectDeletemembershipLevel)

	PML.POST("/membershiplevelisactive", controllers.MembershipLevelIsactive)

	//

	PMS := C.Group("/subscription")

	PMS.GET("/", controllers.Subscription)

	// PMS.GET("/createsubscription", controllers.Createsubscription)

	PMS.GET("/addnewsubscription", controllers.AddNewSubscription)

	PMS.POST("/createmembershipsubscription", controllers.CreateMembershipSubscription)

	PMS.GET("/editsubscription", controllers.EditSubscription)

	PMS.POST("/updatesubscription", controllers.UpdateSubscription)

	PMS.GET("/deletesubscription", controllers.Deletesubscription)

	PMS.POST("/subscriptionisactive", controllers.SubscriptionIsactive)

	PMS.POST("/multiselectdeletesubscription", controllers.MultiSelectDeleteSubscription)

	//MemberShip Order

	PMO := C.Group("/order")

	PMO.GET("/", controllers.OrderList)

	PMO.GET("/createorder", controllers.Createorder)

	PMO.POST("/addorder", controllers.AddOrder)

	PMO.GET("/editorder/:id", controllers.EditOrder)

	PMO.POST("/updateorder", controllers.UpdateOrder)

	PMO.GET("/deleteorder/:id", controllers.DeleteOrder)

	PMO.POST("/multiselectdeleteorder", controllers.MultiSelectDeleteOrder)

	PMG := C.Group("/membershipgroup")

	PMG.GET("/", controllers.ListMembershipgroup)

	PMG.POST("/createmembershiplevelgroup", controllers.CreateMembershipGroupLevel)

	// PMG.GET("/editmembershiplevelgroup", controllers.EditMembershipGroupLevel)

	PMG.POST("/updatemembershiplevelgroup", controllers.UpdateMembershipGroups)

	PMG.GET("/deletemembershiplevelgroup", controllers.DeleteMembershipGroup)

	PMG.POST("/multiselectdeletemembershipgroup", controllers.MultiselectDeleteMembershipGroup)

	PMG.POST("/membershipgroupisactive", controllers.MembershipGroupIsactive)

	PI := C.Group("/integration")

	PI.GET("/", controllers.PaymentIntgrationList)

	PI.GET("/master", controllers.MasterIntegrationList)

	PI.GET("/addcollection", controllers.AddIntegrationToCollection)

	PI.GET("/edit", controllers.IntegrationEdit)

	PI.POST("/manage", controllers.IntegrationManage)

	//Courses APIs

	//Listing APIs

	//Menu Routes

	We := C.Group("/website")

	We.GET("/", controllers.WebsiteList)

	We.POST("/createwebsite", controllers.CreateWebsite)

	We.POST("/updatewebsite", controllers.UpdateWebsite)

	We.GET("/deletewebsite/:id", controllers.DeleteWebsite)

	We.POST("/checksitename", controllers.CheckSiteName)

	We.GET("/template/:id", controllers.GoTemplateUpdate)

	We.GET("/seo", controllers.Seo)

	We.POST("/seopage", controllers.SeoPage)

	We.GET("/setting", controllers.Settings)

	We.POST("/settingpage/:template", controllers.SettingPage)

	We.POST("/duplicatedomainname", controllers.CheckDomainName)

	Me := We.Group("/menu")

	Me.GET("/", controllers.MenuList)

	Me.POST("/createmenu", controllers.CreateMenu)

	Me.POST("/updatemenu", controllers.UpdateMenu)

	Me.GET("/deletemenu/:id", controllers.DeleteMenu)

	Me.POST("/checkmenuname", controllers.CheckMenuName)

	Me.POST("/menustatuschange", controllers.MenuStatusChange)

	Me.GET("/publish/:id", controllers.MenuPublish)

	Me.GET("/menuitems/:id", controllers.MenuIemsList)

	Me.POST("/createmenuitems", controllers.CreateMenuItem)

	Me.POST("/updatemenuitems", controllers.UpdateMenuItem)

	Me.GET("/deletemenuitem/:id", controllers.DeleteMenuItem)

	Me.GET("/editmenuitem/:id", controllers.EditMenuItem)

	Me.POST("/updatemenuitemsorder", controllers.UpdateMenuitemOrder)

	Wp := We.Group("/pages")

	// =============

	Wp.GET("/", controllers.TemplateEditPage)

	Wp.GET("/createpage", controllers.AddPageInWebsite)

	Wp.GET("/editpage/:id", controllers.AddPageInWebsite)

	Wp.POST("/savepagedata", controllers.SavePage)

	Wp.GET("/deletepage/:id", controllers.DeletePage)

	Wp.POST("/pagestatuschange", controllers.PageStatusChange)

	Wp.POST("/updatepagesorder", controllers.UpdatePagesOrder)

	Wp.GET("/clone/:id", controllers.ClonePage)

	Ww := We.Group("/widgets")

	Ww.GET("/", controllers.WidgetsList)

	Ww.GET("/createwidget", controllers.CreateWidget)

	Ww.POST("/savewidget", controllers.SaveWidget)

	Ww.GET("/editwidget/:id", controllers.CreateWidget)

	Ww.POST("/updatewidget", controllers.SaveWidget)

	Ww.GET("/deletewidget/:id", controllers.DeleteWidget)

	Ww.POST("/widgetstatuschange", controllers.WidgetStatusChange)

	// Br.GET("/brandsitename", controllers.BrandSiteName)

	// Br.GET("/brandsitelogo", controllers.BrandSiteLogo)

	Gs := C.Group("getting-started")

	Gs.GET("/", controllers.GettingStarted)

	Gs.GET("/selectchannel", controllers.SelectChannel)

	Gs.GET("/selectlanguage", controllers.SelectLanguage)

	Gs.GET("/selectcountry", controllers.SelectCountry)

	Gs.GET("/createdomain", controllers.CreateDomain)

	Gs.POST("/usagemodeupdate", controllers.UsageModeUpdate)

	Gs.POST("/updateselectchannel", controllers.UpdateChannelInfo)

	Gs.POST("/selectedlanguage", controllers.SelectedLanguage)

	Gs.POST("/selectedcountry", controllers.SelectedCountry)

	Gs.POST("/subdomain", controllers.SubDomain)

	return r

}
