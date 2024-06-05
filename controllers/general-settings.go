package controllers

import (
	"fmt"
	"spurt-cms/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

// set default timezone in general settings
func SetInitialGeneralValues() {

	tblgeneralsetting, err := models.GetGeneralSettings()

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

	tblgeneralsetting, err := models.GetGeneralSettings()
	if err != nil {
		ErrorLog.Printf("Getting data from general settings Error:%s", err)
	}

	var language1 []models.TblLanguage

	query.DataAccess = c.GetInt("dataaccess")
	query.UserId = c.GetInt("userid")
	Language, _ := query.GetLanguageList(&language1, 0, 0, models.Filter{}, true)

	Folder, File, Media, _ := GetMedia()
	selectedtype, _ := GetSelectedType()

	menu := NewMenuController(c)
	translate, _ := TranslateHandler(c)

	timezones, err := models.GetTimeZones()
	if err != nil {
		ErrorLog.Printf("Getting data from general settings Error:%s", err)
	}
	

	c.HTML(200, "general-settings.html", gin.H{"csrf": csrf.GetToken(c), "HeadTitle": translate.Setting.Generalsettings, "Menu": menu, "translate": translate, "SettingsHead": true, "Generalmenu": true, "Tooltiptitle": translate.Setting.Teamstooltip, "title": "General Settings", "GeneralSetting": tblgeneralsetting, "Langugage": Language, "TimeZone": timezones, "Folder": Folder, "File": File, "Media": Media, "StorageType": selectedtype.SelectedType})
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

	//
	companyName := c.PostForm("companyname")
	imagePath1 := c.PostForm("logo_imgpath")
	expandimagePath := c.PostForm("expandlogo_imgpath")
	dateFormat := c.PostForm("dateformat")
	timeForamt := c.PostForm("timeformat")
	timezone := c.PostForm("timezon")
	languagedefault, _ := strconv.Atoi(c.PostForm("language"))

	var gensetting models.TblGeneralSetting

	gensetting.CompanyName = companyName
	gensetting.LogoPath = imagePath1
	gensetting.ExpandLogoPath = expandimagePath
	gensetting.DateFormat = dateFormat
	gensetting.TimeFormat = timeForamt
	gensetting.TimeZone = timezone
	gensetting.LanguageId = languagedefault
	gensetting.ModifiedBy = c.GetInt("userid")
	gensetting.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().In(TZONE).Format("2006-01-02 15:04:05"))

	err := models.UpdateGeneralSettings(gensetting)

	if err != nil {
		ErrorLog.Println(err)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.Redirect(301, "/settings/general-settings")
		return
	}

	SetTimeZone(timezone)
	Datelayout = DateFormater(dateFormat, timeForamt)
	CurrentLanugageId = languagedefault
	LogoPath = imagePath1
	ExpantLogoPath = expandimagePath

	c.SetCookie("get-toast", "General Settings Updated Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, "/settings/general-settings")

}
