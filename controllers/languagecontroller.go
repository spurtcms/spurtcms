package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"spurt-cms/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func LanguageList(c *gin.Context) {

	User.Authority = &AUTH

	var (
		limit         int
		offset        int
		filter        models.Filter
		language1     []models.TblLanguage
		language      []models.TblLanguage
		conv_language []models.TblLanguage
	)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	if c.Query("limit") == "" {
		limit = Limit
	} else {
		limit, _ = strconv.Atoi(c.Query("limit"))
	}

	if page != 0 {
		offset = (page - 1) * limit
	}

	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")

	flag := false

	query.DataAccess = c.GetInt("dataaccess")
	query.UserId = c.GetInt("userid")

	_, Total_languages := query.GetLanguageList(&language1, 0, 0, filter, flag)

	query.GetLanguageList(&language, limit, offset, filter, flag)

	for _, language := range language {

		var default_lang models.TblLanguage
		models.GetLanguageById(&default_lang, CurrentLanugageId)

		if default_lang.Id == language.Id {
			language.DefaultLanguageId = default_lang.Id
		}

		if language.ModifiedOn.IsZero() {
			language.DateString = language.CreatedOn.In(TZONE).Format(Datelayout)
		} else {
			language.DateString = language.ModifiedOn.In(TZONE).Format(Datelayout)
		}

		conv_language = append(conv_language, language)
	}

	paginationendcount := len(language) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(page, int(Total_languages), limit)

	translate, _ := TranslateHandler(c)
	menu := NewMenuController(c)

	userDetails, err := User.GetUserDetails(c.GetInt("userid"))

	if err != nil {
		log.Println(err)
	}

	Folder, File, Media, _ := GetMedia()
	selectedtype, _ := GetSelectedType()

	c.HTML(200, "language.html", gin.H{"Pagination": PaginationData{
		NextPage:     page + 1,
		PreviousPage: page - 1,
		TotalPages:   PageCount,
		TwoAfter:     page + 2,
		TwoBelow:     page - 2,
		ThreeAfter:   page + 3,
	}, "Menu": menu, "Folder": Folder, "File": File, "Media": Media, "HeadTitle": translate.Language, "translate": translate, "csrf": csrf.GetToken(c), "Languages": conv_language, "Count": Total_languages, "Limit": limit, "Filter": filter, "UserDetails": userDetails, "Previous": Previous, "Next": Next, "Page": Page, "PageCount": PageCount, "CurrentPage": page, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "Langmenu": true, "SettingsHead": true, "Languagecodearr": Lang_codearr, "title": "Language", "Tooltiptitle": translate.Setting.Languagetooltip, "StorageType": selectedtype.SelectedType})

}

func AddLanguage(c *gin.Context) {

	if c.PostForm("flag_imgpath") == "" || c.PostForm("lang_name") == "" || c.PostForm("lang_code") == "" {
		c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
		c.Redirect(301, "/settings/languages/")
		return
	}

	var newlanguagedata models.TblLanguage

	jsonfile, _ := c.FormFile("lang_json")

	ext := strings.Split(jsonfile.Filename, ".")[len(strings.Split(jsonfile.Filename, "."))-1]

	if ext == "json" {

		newlanguagedata.LanguageCode = c.PostForm("lang_code")
		json_dst := fmt.Sprintf("locales/%s", c.PostForm("lang_code")+"."+ext)
		newlanguagedata.JsonPath = json_dst
		c.SaveUploadedFile(jsonfile, json_dst)
		localjsondata, _ := os.Open(json_dst)

		var jsonData map[string]interface{}
		json.NewDecoder(localjsondata).Decode(&jsonData)

	}

	newlanguagedata.IsDefault, _ = strconv.Atoi(c.PostForm("lang_default"))
	newlanguagedata.ImagePath = c.PostForm("flag_imgpath")
	newlanguagedata.LanguageName = c.PostForm("lang_name")
	newlanguagedata.IsStatus, _ = strconv.Atoi(c.PostForm("lang_status"))
	newlanguagedata.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	newlanguagedata.CreatedBy = c.GetInt("userid")

	err := models.AddNewLanguage(&newlanguagedata, c.GetInt("userid"))

	if err != nil {
		ErrorLog.Printf("Add Language error: %s", err)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.Redirect(301, "/settings/languages/")
		return
	}

	c.SetCookie("get-toast", "Language Added Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, "/settings/languages/")
}

func EditLanguage(c *gin.Context) {

	languageId, _ := strconv.Atoi(c.Param("id"))
	userid := c.GetInt("userid")
	var language models.TblLanguage

	err := models.GetLanguageDetails(&language, languageId, userid)
	if err != nil {
		log.Println(err)
	}

	language.DefaultLanguageId = CurrentLanugageId

	c.JSON(200, gin.H{"csrf": csrf.GetToken(c), "Language": language, "UploadUrl": "/settings/languages/updatelanguage", "Button": "Update"})
}

func UpdateLanguage(c *gin.Context) {

	if c.PostForm("lang_id") == "" || c.PostForm("flag_imgpath") == "" || c.PostForm("lang_name") == "" {

		c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
		c.Redirect(301, "/settings/languages/")
		return
	}

	var editLanguage models.TblLanguage

	editLanguage.Id, _ = strconv.Atoi(c.PostForm("lang_id"))
	editLanguage.ImagePath = c.PostForm("flag_imgpath")
	editLanguage.LanguageCode = c.PostForm("lang_code")
	editLanguage.IsDefault, _ = strconv.Atoi(c.PostForm("lang_default"))
	editLanguage.IsStatus, _ = strconv.Atoi(c.PostForm("lang_status"))
	editLanguage.LanguageName = c.PostForm("lang_name")
	editLanguage.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	editLanguage.ModifiedBy = c.GetInt("userid")

	// if c.PostForm("lang_json") == "" {

	// 	jsonfile, _ := c.FormFile("lang_json")

	// 	ext := strings.Split(jsonfile.Filename, ".")[len(strings.Split(jsonfile.Filename, "."))-1]

	// 	if ext == "json" {

	// 		json_dst := fmt.Sprintf("locales/%s", c.PostForm("lang_code")+"."+ext)

	// 		editLanguage.JsonPath = json_dst

	// 		c.SaveUploadedFile(jsonfile, json_dst)

	// 		localjsondata, _ := os.Open(json_dst)

	// 		var jsonData map[string]interface{}

	// 		json.NewDecoder(localjsondata).Decode(&jsonData)

	// 	}

	// } else {

	// 	editLanguage.JsonPath = c.PostForm("lang_json")
	// }

	err := models.UpdateLanguage(&editLanguage, c.GetInt("userid"))

	if editLanguage.IsDefault == 1 {
		models.UpdateLanuguageInGeneral(editLanguage.Id)
		CurrentLanugageId = editLanguage.Id
	}

	if err != nil {
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.Redirect(301, "/settings/languages/")
		return
	}

	var default_lang models.TblLanguage
	err = models.GetDefaultLanguage(&default_lang, c.GetInt("userid"))

	if err != nil {
		c.SetCookie("lang", "", 3600, "", "", false, false)

	} else {

		conv_langId := strconv.Itoa(default_lang.Id)

		if c.PostForm("lang_default") == conv_langId {
			c.SetCookie("lang", conv_langId, 3600, "", "", false, false)
		} else {
			c.SetCookie("lang", "", 3600, "", "", false, false)
		}
	}

	c.SetCookie("get-toast", "Language Updated Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, "/settings/languages/")

}

func DeleteLanguage(c *gin.Context) {

	languageId, _ := strconv.Atoi(c.Param("id"))

	userid := c.GetInt("userid")

	var (
		default_lang models.TblLanguage
		language     models.TblLanguage
	)

	err := models.GetDefaultLanguage(&default_lang, userid)

	if err != nil {
		log.Println(err)
	}

	if languageId == default_lang.Id || languageId == 1 {

		c.SetCookie("Alert-msg", "DefaultLanguageNotAllowedToDelete", 3600, "", "", false, false)
		c.Redirect(301, "/settings/languages/")
		return
	}

	/*uncomment the below code to use for language json file delete purpose*/

	// var langData models.TblLanguage

	// models.GetLanguageById(&langData,languageId)

	// log.Println("language json delete path",langData.JsonPath)

	// err = os.Remove(langData.JsonPath)

	// if err != nil {

	//     log.Println("Error deleting file:", err)

	//     return
	// }

	// fmt.Println("language json deleted successfully!")

	language.Id = languageId
	language.IsDeleted = 1
	language.DeletedBy = c.GetInt("userid")
	language.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err = models.DeleteLanguage(&language)

	if err != nil {
		ErrorLog.Printf("DeleteMultipleUser  error: %s", err)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		return
	}

	c.SetCookie("get-toast", "Language Deleted Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	c.Redirect(301, "/settings/languages/")
}

func DownloadLanguage(c *gin.Context) {

	langId, _ := strconv.Atoi(c.Param("id"))

	var langData models.TblLanguage

	err := models.GetLanguageById(&langData, langId)
	if err != nil {
		log.Println(err)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		return
	}

	jsonData, err := ioutil.ReadFile(langData.JsonPath)
	if err != nil {
		log.Println(err)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+langData.LanguageName+"."+strings.Split(langData.JsonPath, ".")[1])
	c.Data(200, "application/json", jsonData)
	c.SetCookie("get-toast", "Language Json file Download Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)

}

func LanguageStatus(c *gin.Context) {

	var langData models.TblLanguage

	id, _ := strconv.Atoi(c.Request.PostFormValue("id"))

	langData.IsStatus, _ = strconv.Atoi(c.Request.PostFormValue("isactive"))
	langData.ModifiedBy = c.GetInt("userid")
	langData.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	models.Languageisactive(&langData, id)

	var language models.TblLanguage
	models.GetLanguageDetails(&language, id, c.GetInt("userid"))

	c.JSON(200, gin.H{"language": language})

}

// language name check
func CheckLanguageName(c *gin.Context) {

	var chekclang models.TblLanguage

	name := c.PostForm("langname")
	id, _ := strconv.Atoi(c.PostForm("id"))

	err := models.Checklang(&chekclang, name, id)

	if err != nil {
		ErrorLog.Printf("check language name error: %s", err)
		json.NewEncoder(c.Writer).Encode(false)
		return
	}

	json.NewEncoder(c.Writer).Encode(true)

}

func MultiselectLanguageDelete(c *gin.Context) {

	var languagedata []map[string]string

	if err := json.Unmarshal([]byte(c.Request.PostFormValue("languageids")), &languagedata); err != nil {
		ErrorLog.Printf("Multiple language delete unmarshall error: %s", err)
	}

	var languageids []int

	for _, val := range languagedata {
		languageid, _ := strconv.Atoi(val["languageid"])
		languageids = append(languageids, languageid)
	}

	var language models.TblLanguage

	language.IsDeleted = 1
	language.DeletedBy = c.GetInt("userid")
	language.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	err := models.MultiSelectDeleteLanguage(&language, languageids)

	if err != nil {
		ErrorLog.Printf("Multiple language delete error: %s", err)
		c.JSON(200, gin.H{"value": false})
		return
	}

	c.JSON(200, gin.H{"value": true, "url": "/settings/languages/"})
}

func MultiSelectLanguageStatus(c *gin.Context) {

	var (
		languagedata []map[string]string
		languageids  []int
		statusint    int
		status       int
	)

	if err := json.Unmarshal([]byte(c.Request.PostFormValue("languageids")), &languagedata); err != nil {
		ErrorLog.Printf("Multiple language status change unmarshall error: %s", err)
	}

	for _, val := range languagedata {

		languageid, _ := strconv.Atoi(val["languageid"])
		statusint, _ = strconv.Atoi(val["status"])
		if statusint == 0 {
			status = 1
		} else if statusint == 1 {
			status = 0
		}

		languageids = append(languageids, languageid)

	}

	var langData models.TblLanguage
	langData.IsStatus = status
	langData.ModifiedBy = c.GetInt("userid")
	langData.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := models.MultiSelectLanguageisactive(&langData, languageids)

	if err != nil {
		ErrorLog.Printf("Multiple language status change  error: %s", err)
		c.JSON(200, gin.H{"value": false})
		return
	}

	c.JSON(200, gin.H{"value": true, "status": status, "url": "/settings/languages/"})

}
