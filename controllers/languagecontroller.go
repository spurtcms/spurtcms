package controllers

import (
	"fmt"
	"log"
	"spurt-cms/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func LanguageList(c *gin.Context) {

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

	_, Total_languages := query.GetLanguageList(&language1, 0, 0, filter, flag, TenantId)

	query.GetLanguageList(&language, limit, offset, filter, flag, TenantId)

	for _, language := range language {

		// var default_lang models.TblLanguage
		// models.GetLanguageById(&default_lang, CurrentLanugageId)

		// if default_lang.Id == language.Id {
		// 	language.DefaultLanguageId = default_lang.Id
		// }

		if language.ModifiedOn.IsZero() {
			language.DateString = language.CreatedOn.In(TZONE).Format(Datelayout)
		} else {
			language.DateString = language.ModifiedOn.In(TZONE).Format(Datelayout)
		}

		if language.ImagePath != "" {

			if !strings.Contains(language.ImagePath, "/public/img/") {
				if language.StorageType == "local" {
					language.ImagePath = "/" + language.ImagePath

				} else if language.StorageType == "aws" {
					language.ImagePath = "/image-resize?name=" + language.ImagePath
				}

			}

		}

		conv_language = append(conv_language, language)
	}

	paginationendcount := len(language) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(page, int(Total_languages), limit)

	translate, _ := TranslateHandler(c)
	menu := NewMenuController(c)

	userDetails, _, err := NewTeamWP.GetUserById(c.GetInt("userid"), []int{})

	if err != nil {
		log.Println(err)
	}

	selectedtype, _ := GetSelectedType()

	c.HTML(200, "language.html", gin.H{"Pagination": PaginationData{
		NextPage:     page + 1,
		PreviousPage: page - 1,
		TotalPages:   PageCount,
		TwoAfter:     page + 2,
		TwoBelow:     page - 2,
		ThreeAfter:   page + 3,
	}, "Menu": menu, "HeadTitle": translate.Language, "translate": translate, "csrf": csrf.GetToken(c), "Languages": conv_language, "Count": Total_languages, "Limit": limit, "Filter": filter, "UserDetails": userDetails, "Previous": Previous, "Next": Next, "Page": Page, "PageCount": PageCount, "CurrentPage": page, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "Langmenu": true, "SettingsHead": true, "Languagecodearr": Lang_codearr, "title": "Languages", "linktitle": "Language", "Tooltiptitle": translate.Setting.Languagetooltip, "StorageType": selectedtype.SelectedType})

}

func SetDefaultLanguage(c *gin.Context) {

	isDefault, err := strconv.Atoi(c.PostForm("isDefault"))
	if err != nil {
		c.SetCookie("lang", "", 3600, "", "", false, false)
	}

	langId, err := strconv.Atoi(c.PostForm("langId"))
	if err != nil {
		c.SetCookie("lang", "", 3600, "", "", false, false)
	}

	fmt.Println("isDefault", isDefault)

	editedLang, err := models.SetDefaultLang(langId, isDefault, TenantId)
	if err != nil {
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		c.Redirect(301, "/settings/languages/")
		return
	}

	if editedLang.IsDefault == 1 {
		models.UpdateLanuguageInGeneral(editedLang.Id, TenantId)
		CurrentLanugageId = editedLang.Id
	}

	var defaultLang models.TblLanguage
	err = models.GetDefaultLanguage(&defaultLang, c.GetInt("userid"), TenantId)
	if err != nil {
		c.SetCookie("lang", "", 3600, "", "", false, false)

	} else {

		conv_langId := strconv.Itoa(defaultLang.Id)
		c.SetCookie("lang", conv_langId, 3600, "", "", false, false)

		// if c.PostForm("lang_default") == conv_langId {
		// 	c.SetCookie("lang", conv_langId, 3600, "", "", false, false)
		// } else {
		// 	c.SetCookie("lang", "", 3600, "", "", false, false)
		// }
	}

	c.JSON(200, gin.H{"value": true})

}

// func AddLanguage(c *gin.Context) {

// 	if c.PostForm("flag_imgpath") == "" || c.PostForm("lang_name") == "" || c.PostForm("lang_code") == "" {
// 		c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
// 		c.Redirect(301, "/settings/languages/")
// 		return
// 	}

// 	var newlanguagedata models.TblLanguage

// 	//adding new  lang json file to the code
// 	jsonfile, _ := c.FormFile("lang_json")

// 	ext := strings.Split(jsonfile.Filename, ".")[len(strings.Split(jsonfile.Filename, "."))-1]

// 	if ext == "json" {

// 		newlanguagedata.LanguageCode = c.PostForm("lang_code")
// 		json_dst := fmt.Sprintf("locales/%s", c.PostForm("lang_code")+"."+ext)
// 		newlanguagedata.JsonPath = json_dst
// 		c.SaveUploadedFile(jsonfile, json_dst)
// 		localjsondata, _ := os.Open(json_dst)

// 		var jsonData map[string]interface{}
// 		json.NewDecoder(localjsondata).Decode(&jsonData)

// 	}

// 	// get the selected type for image upload from the db
// 	storageType, err := GetSelectedType()
// 	if err != nil {
// 		ErrorLog.Println(err)

// 		ErrorLog.Printf("Add Language error: %s", err)
// 		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
// 		c.Redirect(301, "/settings/languages/")

// 		return
// 	}

// 	var langDefault = c.PostForm("lang_default")

// 	if langDefault != "" {
// 		newlanguagedata.IsDefault, err = strconv.Atoi(langDefault)
// 		if err != nil {
// 			ErrorLog.Println(err)

// 			ErrorLog.Printf("Add Language error: %s", err)
// 			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
// 			c.Redirect(301, "/settings/languages/")

// 			return
// 		}

// 	} else {

// 		newlanguagedata.IsDefault = 0
// 	}
// 	//getting the default lang value from the request

// 	//image processing according to the selected type
// 	langData := c.PostForm("flag_imgpath")

// 	if strings.Contains(langData, "data:image/jpeg;base64") || strings.Contains(langData, "data:image/png;base64") || strings.Contains(langData, "data:image/svg+xml;base64") {

// 		// image processing if the selected type is aws
// 		if storageType.SelectedType == "aws" {

// 			tenantDetails, err := NewTeam.GetTenantDetails(TenantId)
// 			if err != nil {
// 				ErrorLog.Println(err)

// 				ErrorLog.Printf("Add Language error: %s", err)
// 				c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
// 				c.Redirect(301, "/settings/languages/")

// 				return
// 			}

// 			var (
// 				tempString, imageName string
// 				imageByte             []byte
// 			)

// 			imageName, tempString, imageByte, err = ConvertBase64toByte(langData, "lang")
// 			if err != nil {
// 				ErrorLog.Println(err)

// 				ErrorLog.Printf("Add Language error: %s", err)
// 				c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
// 				c.Redirect(301, "/settings/languages/")

// 				return
// 			}

// 			tempString = tenantDetails.S3FolderName + tempString

// 			err = storagecontroller.UploadCropImageS3(imageName, tempString, imageByte)
// 			if err != nil {
// 				ErrorLog.Println(err)

// 				ErrorLog.Printf("Add Language error: %s", err)
// 				c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
// 				c.Redirect(301, "/settings/languages/")

// 				return
// 			}

// 			newlanguagedata.ImagePath = tempString

// 		} else if storageType.SelectedType == "local" {
// 			// image processing if the selected type is local
// 			var tempImgPath string

// 			_, tempImgPath, err = ConvertBase64(langData, "storage/lang")
// 			if err != nil {
// 				ErrorLog.Println(err)

// 				ErrorLog.Printf("Add Language error: %s", err)
// 				c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
// 				c.Redirect(301, "/settings/languages/")

// 				return
// 			}

// 			newlanguagedata.ImagePath = "/" + tempImgPath

// 		}

// 	} else {
// 		newlanguagedata.ImagePath = langData
// 	}

// 	newlanguagedata.LanguageName = c.PostForm("lang_name")
// 	newlanguagedata.IsStatus = 1
// 	newlanguagedata.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
// 	newlanguagedata.CreatedBy = c.GetInt("userid")
// 	newlanguagedata.StorageType = storageType.SelectedType
// 	newlanguagedata.TenantId = TenantId

// 	err = models.AddNewLanguage(&newlanguagedata, c.GetInt("userid"))
// 	if err != nil {
// 		ErrorLog.Printf("Add Language error: %s", err)
// 		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
// 		c.Redirect(301, "/settings/languages/")
// 		return
// 	}

// 	c.SetCookie("get-toast", "Language Added Successfully", 3600, "", "", false, false)
// 	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
// 	c.Redirect(301, "/settings/languages/")
// }

// func EditLanguage(c *gin.Context) {

// 	languageId, _ := strconv.Atoi(c.Param("id"))
// 	userid := c.GetInt("userid")
// 	var language models.TblLanguage

// 	err := models.GetLanguageDetails(&language, languageId, userid, TenantId)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	if !strings.Contains(language.ImagePath, "/public/img/") {
// 		if language.ImagePath != "" {
// 			if language.StorageType == "local" {
// 				language.ImagePath = "/" + language.ImagePath

// 			} else if language.StorageType == "aws" {
// 				language.ImagePath = "/image-resize?name=" + language.ImagePath
// 			}
// 		}

// 	}

// 	language.DefaultLanguageId = CurrentLanugageId

// 	c.JSON(200, gin.H{"csrf": csrf.GetToken(c), "Language": language, "UploadUrl": "/settings/languages/updatelanguage", "Button": "Update"})
// }

// func UpdateLanguage(c *gin.Context) {

// 	if c.PostForm("lang_id") == "" || c.PostForm("lang_name") == "" {

// 		c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
// 		c.Redirect(301, "/settings/languages/")
// 		return
// 	}

// 	var editLanguage models.TblLanguage

// 	editLanguage.Id, _ = strconv.Atoi(c.PostForm("lang_id"))
// 	editLanguage.LanguageCode = c.PostForm("lang_code")
// 	editLanguage.IsDefault, _ = strconv.Atoi(c.PostForm("lang_default"))
// 	editLanguage.IsStatus, _ = strconv.Atoi(c.PostForm("lang_status"))
// 	editLanguage.LanguageName = c.PostForm("lang_name")
// 	editLanguage.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
// 	editLanguage.ModifiedBy = c.GetInt("userid")

// 	// get the selected type for image upload from the db
// 	storageType, err := GetSelectedType()
// 	if err != nil {
// 		ErrorLog.Println(err)

// 		ErrorLog.Printf("Add Language error: %s", err)
// 		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
// 		c.Redirect(301, "/settings/languages/")

// 		return
// 	}

// 	var tempImgPath string

// 	langData := c.PostForm("flag_imgpath")

// 	if strings.Contains(langData, "data:image/jpeg;base64") || strings.Contains(langData, "data:image/png;base64") || strings.Contains(langData, "data:image/svg+xml;base64") {

// 		if storageType.SelectedType == "aws" {

// 			tenantDetails, err := NewTeam.GetTenantDetails(TenantId)
// 			if err != nil {
// 				ErrorLog.Println(err)

// 				ErrorLog.Printf("Add Language error: %s", err)
// 				c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
// 				c.Redirect(301, "/settings/languages/")

// 				return
// 			}

// 			var (
// 				tempString, imageName string
// 				imageByte             []byte
// 			)

// 			imageName, tempString, imageByte, err = ConvertBase64toByte(langData, "lang")
// 			if err != nil {
// 				ErrorLog.Println(err)

// 				ErrorLog.Printf("Add Language error: %s", err)
// 				c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
// 				c.Redirect(301, "/settings/languages/")

// 				return
// 			}

// 			tempString = tenantDetails.S3FolderName + tempString

// 			err = storagecontroller.UploadCropImageS3(imageName, tempString, imageByte)
// 			if err != nil {
// 				ErrorLog.Println(err)

// 				ErrorLog.Printf("Add Language error: %s", err)
// 				c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
// 				c.Redirect(301, "/settings/languages/")

// 				return
// 			}

// 			editLanguage.ImagePath = tempString

// 		} else if storageType.SelectedType == "local" {

// 			_, tempImgPath, err = ConvertBase64(langData, "storage/lang")
// 			if err != nil {
// 				ErrorLog.Println(err)

// 				ErrorLog.Printf("Add Language error: %s", err)
// 				c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
// 				c.Redirect(301, "/settings/languages/")

// 				return
// 			}

// 			editLanguage.ImagePath = "/" + tempImgPath

// 		}

// 	} else {

// 		prefixEndIndex := strings.Index(langData, "=")
// 		fmt.Println("prefix", prefixEndIndex)

// 		prefixRemovedPath := langData[prefixEndIndex+1:]
// 		fmt.Println("prefixRemoved", prefixRemovedPath)

// 		editLanguage.ImagePath = prefixRemovedPath
// 	}

// 	jsonFile, _ := c.FormFile("lang_json")

// 	if jsonFile != nil {

// 		ext := strings.Split(jsonFile.Filename, ".")[len(strings.Split(jsonFile.Filename, "."))-1]

// 		if ext == "json" {
// 			json_dst := fmt.Sprintf("locales/%s", c.PostForm("lang_code")+"."+ext)
// 			editLanguage.JsonPath = json_dst
// 			c.SaveUploadedFile(jsonFile, json_dst)
// 			localjsondata, _ := os.Open(json_dst)

// 			var jsonData map[string]interface{}
// 			json.NewDecoder(localjsondata).Decode(&jsonData)
// 		}
// 	} else {

// 		var languageData models.TblLanguage
// 		userid := c.GetInt("userid")
// 		err = models.GetLanguageDetails(&languageData, editLanguage.Id, userid, TenantId)
// 		if err != nil {
// 			log.Println(err)
// 		}

// 		editLanguage.JsonPath = languageData.JsonPath
// 	}

// 	editLanguage.StorageType = storageType.SelectedType

// 	err = models.UpdateLanguage(&editLanguage, c.GetInt("userid"), TenantId)
// 	if err != nil {
// 		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
// 		c.Redirect(301, "/settings/languages/")
// 		return
// 	}

// 	if editLanguage.IsDefault == 1 {
// 		models.UpdateLanuguageInGeneral(editLanguage.Id, TenantId)
// 		CurrentLanugageId = editLanguage.Id
// 	}

// 	var default_lang models.TblLanguage
// 	err = models.GetDefaultLanguage(&default_lang, c.GetInt("userid"), TenantId)

// 	if err != nil {
// 		c.SetCookie("lang", "", 3600, "", "", false, false)

// 	} else {

// 		conv_langId := strconv.Itoa(default_lang.Id)

// 		if c.PostForm("lang_default") == conv_langId {
// 			c.SetCookie("lang", conv_langId, 3600, "", "", false, false)
// 		} else {
// 			c.SetCookie("lang", "", 3600, "", "", false, false)
// 		}
// 	}

// 	c.SetCookie("get-toast", "Language Updated Successfully", 3600, "", "", false, false)
// 	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
// 	c.Redirect(301, "/settings/languages/")

// }

// func DeleteLanguage(c *gin.Context) {

// 	languageId, _ := strconv.Atoi(c.Param("id"))

// 	userid := c.GetInt("userid")

// 	var (
// 		default_lang models.TblLanguage
// 		language     models.TblLanguage
// 	)

// 	err := models.GetDefaultLanguage(&default_lang, userid, TenantId)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	if languageId == default_lang.Id || languageId == 1 {

// 		c.SetCookie("Alert-msg", "DefaultLanguageNotAllowedToDelete", 3600, "", "", false, false)
// 		c.Redirect(301, "/settings/languages/")
// 		return
// 	}

// 	/*uncomment the below code to use for language json file delete purpose*/

// 	// var langData models.TblLanguage

// 	// models.GetLanguageById(&langData,languageId)

// 	// log.Println("language json delete path",langData.JsonPath)

// 	// err = os.Remove(langData.JsonPath)

// 	// if err != nil {

// 	//     log.Println("Error deleting file:", err)

// 	//     return
// 	// }

// 	// fmt.Println("language json deleted successfully!")

// 	language.Id = languageId
// 	language.IsDeleted = 1
// 	language.DeletedBy = c.GetInt("userid")
// 	language.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

// 	err = models.DeleteLanguage(&language, TenantId)

// 	if err != nil {
// 		ErrorLog.Printf("DeleteMultipleUser  error: %s", err)
// 		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
// 		return
// 	}

// 	c.SetCookie("get-toast", "Language Deleted Successfully", 3600, "", "", false, false)
// 	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
// 	c.Redirect(301, "/settings/languages/")
// }

// func DownloadLanguage(c *gin.Context) {

// 	langId, _ := strconv.Atoi(c.Param("id"))

// 	var langData models.TblLanguage

// 	err := models.GetLanguageById(&langData, langId)
// 	if err != nil {
// 		log.Println(err)
// 		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
// 		return
// 	}

// 	jsonData, err := ioutil.ReadFile(langData.JsonPath)
// 	if err != nil {
// 		log.Println(err)
// 		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
// 		return
// 	}

// 	c.Header("Content-Disposition", "attachment; filename="+langData.LanguageName+"."+strings.Split(langData.JsonPath, ".")[1])
// 	c.Data(200, "application/json", jsonData)
// 	c.SetCookie("get-toast", "Language Json file Download Successfully", 3600, "", "", false, false)
// 	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)

// }

// func LanguageStatus(c *gin.Context) {

// 	var langData models.TblLanguage

// 	id, err := strconv.Atoi(c.Request.PostFormValue("id"))
// 	if err != nil {
// 		c.JSON(500, gin.H{"language": "", "status": false})
// 	}

// 	langData.IsStatus, _ = strconv.Atoi(c.Request.PostFormValue("isactive"))
// 	langData.ModifiedBy = c.GetInt("userid")
// 	langData.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
// 	models.Languageisactive(&langData, id, TenantId)

// 	var language models.TblLanguage
// 	models.GetLanguageDetails(&language, id, c.GetInt("userid"), TenantId)

// 	c.JSON(200, gin.H{"language": language, "status": true})

// }

// // language name check
// func CheckLanguageName(c *gin.Context) {

// 	var chekclang models.TblLanguage

// 	name := c.PostForm("langname")
// 	id, _ := strconv.Atoi(c.PostForm("id"))

// 	err := models.Checklang(&chekclang, name, id, TenantId)

// 	if err != nil {
// 		ErrorLog.Printf("check language name error: %s", err)
// 		json.NewEncoder(c.Writer).Encode(false)
// 		return
// 	}

// 	json.NewEncoder(c.Writer).Encode(true)

// }

// func MultiselectLanguageDelete(c *gin.Context) {

// 	var languagedata []map[string]string

// 	if err := json.Unmarshal([]byte(c.Request.PostFormValue("languageids")), &languagedata); err != nil {
// 		ErrorLog.Printf("Multiple language delete unmarshall error: %s", err)
// 	}

// 	fmt.Println("languageData", languagedata)

// 	var languageids []int

// 	for _, val := range languagedata {
// 		languageid, _ := strconv.Atoi(val["languageId"])
// 		languageids = append(languageids, languageid)
// 	}

// 	fmt.Println("langaugeIds", languageids)

// 	var language models.TblLanguage

// 	language.IsDeleted = 1
// 	language.DeletedBy = c.GetInt("userid")
// 	language.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
// 	err := models.MultiSelectDeleteLanguage(&language, languageids, TenantId)

// 	if err != nil {
// 		ErrorLog.Printf("Multiple language delete error: %s", err)
// 		c.JSON(200, gin.H{"value": false})
// 		return
// 	}

// 	c.JSON(200, gin.H{"value": true, "url": "/settings/languages/"})
// }

// func MultiSelectLanguageStatus(c *gin.Context) {

// 	var (
// 		languagedata []map[string]string
// 		languageids  []int
// 		statusint    int
// 		status       int
// 	)

// 	if err := json.Unmarshal([]byte(c.Request.PostFormValue("languageids")), &languagedata); err != nil {
// 		ErrorLog.Printf("Multiple language status change unmarshall error: %s", err)
// 	}

// 	for _, val := range languagedata {

// 		languageid, _ := strconv.Atoi(val["languageId"])
// 		statusint, _ = strconv.Atoi(val["status"])
// 		if statusint == 0 {
// 			status = 1
// 		} else if statusint == 1 {
// 			status = 0
// 		}

// 		languageids = append(languageids, languageid)

// 	}

// 	var langData models.TblLanguage
// 	langData.IsStatus = status
// 	langData.ModifiedBy = c.GetInt("userid")
// 	langData.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

// 	err := models.MultiSelectLanguageisactive(&langData, languageids, TenantId)

// 	if err != nil {
// 		ErrorLog.Printf("Multiple language status change  error: %s", err)
// 		c.JSON(200, gin.H{"value": false})
// 		return
// 	}

// 	c.JSON(200, gin.H{"value": true, "status": status, "url": "/settings/languages/"})

// }
