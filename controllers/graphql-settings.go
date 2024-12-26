package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"spurt-cms/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func GraphqlTokenApi(c *gin.Context) {

	var limt, offset int

	keyword := strings.TrimSpace(c.Request.URL.Query().Get("keyword"))

	//get values from html form data
	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	var lists []models.TblGraphqlSettings
	list, count, err := models.GetListOfTokens(limt, offset, keyword, TenantId)
	if err != nil {
		ErrorLog.Printf("get the list token error: %s", err)
	}

	for _, val := range list {

		if !val.ModifiedOn.IsZero() {
			val.DateString = val.ModifiedOn.In(TZONE).Format(Datelayout)
		} else {
			val.DateString = val.CreatedOn.In(TZONE).Format(Datelayout)
		}

		lists = append(lists, val)

	}

	var paginationendcount = len(lists) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(pageno, int(count), limt)

	menu := NewMenuController(c)
	translate, _ := TranslateHandler(c)

	c.HTML(200, "graphql.html", gin.H{"csrf": csrf.GetToken(c), "HeadTitle": "Graphql Settings", "linktitle": "Graphql API", "Menu": menu, "translate": translate, "SettingsHead": true, "Graphqlmenu": true, "title": "Graphql Api", "Tokens": lists, "totalcount": count, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Limit": limt, "filter": keyword, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3,
	}})

}

func CreateTokenGrapqhl(c *gin.Context) {

	menu := NewMenuController(c)
	translate, _ := TranslateHandler(c)

	c.HTML(200, "graphql-manage.html", gin.H{"csrf": csrf.GetToken(c), "HeadTitle": "Graphql Settings", "Menu": menu, "translate": translate, "SettingsHead": true, "Graphqlmenu": true, "title": "Graphql Api", "CreateForm": true})

}

func CreateToken(c *gin.Context) {

	var graphql models.TblGraphqlSettings
	response := make(gin.H)
	graphql.TokenName = c.PostForm("tokenName")
	apiToken, err := GenerateApiToken(64)
	if err != nil {
		ErrorLog.Printf("genereate token error: %v", err)
		response["token"] = ""
		response["status"] = 0
		c.AbortWithStatusJSON(500, response)
		return
	}
	graphql.Token = apiToken
	graphql.Description = c.PostForm("tokenDesc")
	graphql.Duration = c.PostForm("duration")
	graphql.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	graphql.CreatedBy = c.GetInt("userid")
	graphql.TenantId = TenantId

	if graphql.Duration == "7 Days" {

		graphql.ExpiryTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().AddDate(0, 0, 7).Format("2006-01-02 15:04:05"))

	} else if graphql.Duration == "15 Days" {

		graphql.ExpiryTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().AddDate(0, 0, 15).Format("2006-01-02 15:04:05"))

	} else if graphql.Duration == "30 Days" {

		graphql.ExpiryTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().AddDate(0, 0, 30).Format("2006-01-02 15:04:05"))

	}

	err = models.CreateApiToken(graphql)

	if err != nil {
		ErrorLog.Printf("create token error: %s", err)
		response["token"] = ""
		response["status"] = 0
		c.AbortWithStatusJSON(500, response)
		return
	}

	response["token"] = apiToken
	response["status"] = 1
	c.JSON(200, response)

}

func UpdateGraphqlApi(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	token, err := models.GetTokenDetailsById(id, TenantId)
	if err != nil {
		ErrorLog.Printf("get token by id error: %s", err)
	}
	c.JSON(http.StatusOK, gin.H{"data": token})
}

func UpdateGraphqlToken(c *gin.Context) {

	var id, _ = strconv.Atoi(c.PostForm("id"))

	var graphql = make(map[string]interface{})
	graphql["token_name"] = c.Request.FormValue("tokenName")
	graphql["description"] = c.Request.FormValue("tokenDesc")
	graphql["duration"] = c.Request.FormValue("duration")
	graphql["modified_on"], _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	graphql["modified_by"] = c.GetInt("userid")

	var dur = c.Request.FormValue("duration")

	if dur == "7 Days" {

		graphql["expiry_time"], _ = time.Parse("2006-01-02 15:04:05", time.Now().AddDate(0, 0, 7).Format("2006-01-02 15:04:05"))

	} else if dur == "15 Days" {

		graphql["expiry_time"], _ = time.Parse("2006-01-02 15:04:05", time.Now().AddDate(0, 0, 15).Format("2006-01-02 15:04:05"))

	} else if dur == "30 Days" {

		graphql["expiry_time"], _ = time.Parse("2006-01-02 15:04:05", time.Now().AddDate(0, 0, 30).Format("2006-01-02 15:04:05"))

	}

	err := models.UpdateApiToken(graphql, id, TenantId)
	if err != nil {
		ErrorLog.Printf("create token error: %s", err)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
	} else {

		c.SetCookie("get-toast", "Graphql Settings Updated Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	}

}

func DeleteToken(c *gin.Context) {

	var id, _ = strconv.Atoi(c.Param("id"))
	fmt.Println("id", id)
	var deletedon, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	if err := models.DeleteApiToken(id, c.GetInt("userid"), deletedon, TenantId); err != nil {
		ErrorLog.Printf("delete token error: %s", err)
		c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
	} else {

		c.SetCookie("get-toast", "Graphql Settings Deleted Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
	}

	c.Redirect(301, "/graphql/")
}

func GenerateApiToken(length int) (string, error) {
	b := make([]byte, length)               // Create a slice to hold 32 bytes of random data
	if _, err := rand.Read(b); err != nil { // Fill the slice with random data and handle any errors
		return "", err // Return an empty string and the error if something went wrong
	}
	return base64.URLEncoding.EncodeToString(b), nil // Encode the random bytes to a URL-safe base64 string
}

//Multi Delete graphql api//

func MultiDeleteGraphqlToken(c *gin.Context) {

	var tokenIds []string
	var tokenIntIds []int

	if err := json.Unmarshal([]byte(c.Request.PostFormValue("tokenids")), &tokenIds); err != nil {
		ErrorLog.Printf("delete multiple entry error: %s", err)
	}

	fmt.Println(tokenIds, "tokenidesss")

	for _, ids := range tokenIds {
		intIds, _ := strconv.Atoi(ids)

		tokenIntIds = append(tokenIntIds, intIds)
	}

	pageno := c.PostForm("page")
	fmt.Println("pageNonnnn", pageno)

	var url string

	if pageno != "" {
		url = "/graphql/?page=" + pageno
	} else {
		url = "/graphql/"

	}

	var graphql models.TblGraphqlSettings
	graphql.IsDeleted = 1
	graphql.DeletedBy = c.GetInt("userid")
	graphql.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	err := models.MultiDeleteGraphqltoken(&graphql, tokenIntIds, TenantId)

	if err != nil {
		ErrorLog.Printf("Multiple graphqlapi delete error: %s", err)
		c.JSON(200, gin.H{"value": false})
		return
	}

	_, count, err := models.GetListOfTokens(0, 0, "", TenantId)
	if err != nil {
		ErrorLog.Printf("get the list token error: %s", err)
	}

	recordsPerPage := Limit
	totalPages := (int(count) + recordsPerPage - 1) / recordsPerPage

	currentPage := 1
	if pageno != "" {
		currentPage, _ = strconv.Atoi(pageno)
	}

	if currentPage > totalPages && totalPages > 0 {

		url = "/graphql?page=" + strconv.Itoa(totalPages)
	}

	c.JSON(200, gin.H{"value": true, "url": url})
	c.SetCookie("get-toast", "Graphql Settings Deleted Successfully", 3600, "", "", false, false)
	c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)

}

func PlaygroundView(c *gin.Context) {

	url := os.Getenv("GRAPHQL_URL")

	menu := NewMenuController(c)
	translate, _ := TranslateHandler(c)

	c.HTML(http.StatusOK, "playground.html", gin.H{"url": url, "Menu": menu, "translate": translate, "title": "Graphql Playground", "linktitle": "Graphql Playground"})
}
