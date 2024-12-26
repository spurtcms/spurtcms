package controllers

import (
	"fmt"
	"spurt-cms/models"
	storagecontroller "spurt-cms/storage-controller"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

// set default timezone in general settings
func SetInitialGeneralValues() {

	tblgeneralsetting, err := models.GetGeneralSettings(TenantId)

	fmt.Println("defaultValues", tblgeneralsetting)

	if err != nil {
		zone, _ := time.LoadLocation("Asia/Kolkata")
		ErrorLog.Printf("Getting data from general settings Error: %s", err)
		TZONE = zone
		CurrentLanugageId = 1
		Datelayout = "02 Jan 2006 3:04 pm"
		return
	}

	CurrentLanugageId = tblgeneralsetting.LanguageId
	LogoPath = tblgeneralsetting.LogoPath
	ExpantLogoPath = tblgeneralsetting.ExpandLogoPath
	Datelayout = DateFormater(tblgeneralsetting.DateFormat, tblgeneralsetting.TimeFormat)
	zone, zerr := time.LoadLocation(tblgeneralsetting.TimeZone)

	if zerr != nil {
		fmt.Println(zerr)
		ErrorLog.Printf("Invalid Timezone Error: %s", err)
		TZONE = zone
		return
	}

	TZONE = zone

}

func GeneralSettings(c *gin.Context) {

	roleId := c.GetInt("role")

	fmt.Println("roleId", roleId)

	fmt.Println("tenantIdGetSettings", TenantId)

	tblgeneralsetting, err := models.GetGeneralSettings(TenantId)
	if err != nil {
		ErrorLog.Printf("Getting data from general settings Error:%s", err)
	}

	fmt.Println("dateFormattttt", tblgeneralsetting.DateFormat)

	var language1 []models.TblLanguage

	query.DataAccess = c.GetInt("dataaccess")
	query.UserId = c.GetInt("userid")
	Language, _ := query.GetLanguageList(&language1, 0, 0, models.Filter{}, true, TenantId)

	selectedtype, err := GetSelectedType()
	if err != nil {
		ErrorLog.Printf("Getting data from selected type Error:%s", err)
	}

	if !strings.Contains(tblgeneralsetting.LogoPath, "/public/img/") {
		if tblgeneralsetting.LogoPath != "" {

			if tblgeneralsetting.StorageType == "local" {
				tblgeneralsetting.LogoPath = "/" + tblgeneralsetting.LogoPath

			} else if tblgeneralsetting.StorageType == "aws" {
				tblgeneralsetting.LogoPath = "/image-resize?name=" + tblgeneralsetting.LogoPath
			}
		}

	}

	menu := NewMenuController(c)
	translate, _ := TranslateHandler(c)

	timezones, err := models.GetTimeZones(TenantId)
	if err != nil {
		ErrorLog.Printf("Getting data from general settings Error:%s", err)
	}

	c.HTML(200, "general-settings.html", gin.H{"csrf": csrf.GetToken(c), "HeadTitle": translate.Setting.Generalsettings, "Menu": menu, "translate": translate, "SettingsHead": true, "Generalmenu": true, "Tooltiptitle": translate.Setting.Teamstooltip, "title": "General Settings", "linktitle": "General Settings", "GeneralSetting": tblgeneralsetting, "Langugage": Language, "TimeZone": timezones, "StorageType": selectedtype.SelectedType, "RoleId": roleId})
}

func DateFormater(date string, time string) string {

	var Layout string

	switch date {

	case "dd/mm/yyyy":
		if time == "12" {
			Layout = "02/01/2006 3:04 pm"
		} else {
			Layout = "02/01/2006 15:04"
		}

	case "mm/dd/yyyy":
		if time == "12" {
			Layout = "01/02/2006 3:04 pm"
		} else {
			Layout = "01/02/2006 15:04"
		}

	case "dd-mm-yyyy":
		if time == "12" {
			Layout = "01-02-2006 3:04 pm"
		} else {
			Layout = "01-02-2006 15:04"
		}

	case "yyyy-mm-dd":
		if time == "12" {
			Layout = "2006-01-02 3:04 pm"
		} else {
			Layout = "2006-01-02 15:04"
		}

	case "dd mmm yyyy":
		if time == "12" {
			Layout = "02 Jan 2006 3:04 pm"
		} else {
			Layout = "02 Jan 2006 15:04"
		}

	default:
		if time == "12" {
			Layout = "02 Jan 2006 3:04 pm"
		} else {
			Layout = "02 Jan 2006 15:04"
		}
	}

	return Layout
}

func SetTimeZone(timezone string) {

	zone, err := time.LoadLocation(timezone)

	if err != nil {
		fmt.Println(ErrLoadLocation)
		fmt.Println("Set Default UTC")
	}

	TZONE = zone

}

func UpdateGeneralSettings(c *gin.Context) {

	companyName := c.PostForm("companyname")
	// imagePath1 := c.PostForm("logo_imgpath")
	expandimagePath := c.PostForm("expandlogo_imgpath")
	dateFormat := c.PostForm("dateformat")
	timeForamt := c.PostForm("timeformat")
	timezone := c.PostForm("timezon")
	languagedefault, _ := strconv.Atoi(c.PostForm("language"))
	fmt.Println("dateFOrmat", dateFormat)

	var (
		imgPath string
		err     error
	)

	response := make(gin.H)

	storageType, err := GetSelectedType()
	if err != nil {
		ErrorLog.Println(err)

		response["status"] = 0
		c.AbortWithStatusJSON(500, response)

		return
	}

	if strings.Contains(expandimagePath, "data:image/jpeg;base64") || strings.Contains(expandimagePath, "data:image/png;base64") || strings.Contains(expandimagePath, "data:image/svg+xml;base64") {

		tenantDetails, err := NewTeam.GetTenantDetails(TenantId)
		if err != nil {
			ErrorLog.Println(err)

			response["status"] = 0
			c.AbortWithStatusJSON(500, response)

			return
		}

		if storageType.SelectedType == "aws" {
			var (
				tempString, imageName string
				imageByte             []byte
			)

			imageName, tempString, imageByte, err = ConvertBase64toByte(expandimagePath, "user")
			if err != nil {
				ErrorLog.Println(err)

				response["status"] = 0
				c.AbortWithStatusJSON(500, response)

				return
			}

			tempString = tenantDetails.S3FolderName + tempString

			err = storagecontroller.UploadCropImageS3(imageName, tempString, imageByte)
			if err != nil {
				ErrorLog.Println(err)

				response["status"] = 0
				c.AbortWithStatusJSON(500, response)

				return
			}

			imgPath = tempString

		} else if storageType.SelectedType == "local" {
			var tempImgPath string

			_, tempImgPath, err = ConvertBase64(expandimagePath, "storage/user")
			if err != nil {
				ErrorLog.Println(err)

				response["status"] = 0
				c.AbortWithStatusJSON(500, response)

				return
			}

			imgPath = "/" + tempImgPath
		}

	} else {

		prefixEndIndex := strings.Index(expandimagePath, "=")
		fmt.Println("prefix", prefixEndIndex)

		prefixRemovedPath := expandimagePath[prefixEndIndex+1:]
		fmt.Println("prefixRemoved", prefixRemovedPath)

		imgPath = prefixRemovedPath
	}

	var gensetting models.TblGeneralSetting

	gensetting.CompanyName = companyName
	gensetting.LogoPath = imgPath
	// gensetting.ExpandLogoPath = imgPath
	gensetting.DateFormat = dateFormat
	fmt.Println("dateFormat", gensetting.DateFormat)
	gensetting.TimeFormat = timeForamt
	gensetting.TimeZone = timezone

	gensetting.LanguageId = languagedefault
	models.SetDefaultLang(gensetting.LanguageId, 0, TenantId)

	gensetting.ModifiedBy = c.GetInt("userid")
	gensetting.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().In(TZONE).Format("2006-01-02 15:04:05"))
	gensetting.StorageType = storageType.SelectedType

	err = models.UpdateGeneralSettings(gensetting, TenantId)
	if err != nil {
		ErrorLog.Println(err)
		// c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		// c.Redirect(301, "/settings/general-settings/")
		response["status"] = 0
		c.AbortWithStatusJSON(500, response)
		return
	}

	SetTimeZone(timezone)
	Datelayout = DateFormater(dateFormat, timeForamt)
	CurrentLanugageId = languagedefault
	c.SetCookie("lang", string(languagedefault), 3600, "", "", false, false)
	// LogoPath = imagePath1
	ExpantLogoPath = expandimagePath

	// c.SetCookie("get-toast", "General Settings Updated Successfully", 3600, "", "", false, false)
	// c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	// c.Redirect(301, "/settings/general-settings/")

	response["status"] = 1
	c.JSON(200, response)

}
