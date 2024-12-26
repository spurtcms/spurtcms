package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"spurt-cms/config"
	"spurt-cms/controllers"
	"spurt-cms/lang"
	"spurt-cms/middleware"
	"spurt-cms/models"
	"spurt-cms/routes"
	"strings"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func TestSetupRouting(t *testing.T) {
	for _, router := range routes.SetupRoutes().Routes() {

		t.Run(router.Path, func(t *testing.T) {
			request, _ := http.NewRequest(router.Method, router.Path, nil)
			response := httptest.NewRecorder()
			routes.SetupRoutes().ServeHTTP(response, request)
			switch {
			case response.Code == http.StatusOK:
				fmt.Println("running the given status request:", t.Name())
			case response.Code == http.StatusMovedPermanently:
				fmt.Println("running the given status request:", t.Name())
			default:
				t.Fatalf("expected %d or %d, got %d", http.StatusOK, http.StatusMovedPermanently, response.Code)
			}
		})
	}
}

var c *gin.Context

var DB = config.SetupDB()

var Demotoken string

// Member group list
func TestMemberGroupList(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	Demotoken, _, _ = controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {
		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {
			t.Errorf("Failed to load translation: %v", err)
			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("SESSION_KEY")))
	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY"), store))

	r.Use(csrf.Middleware(csrf.Options{
		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	U := r.Group("/membersgroup") // Assuming your routes are under this group
	{
		U.GET("/", middleware.JWTAuth(), controllers.MemberGroupList)
	}

	req, err := http.NewRequest("GET", "/membersgroup/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", Demotoken)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = Demotoken
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedMemberGroup := "newtodaygrp"
	if !strings.Contains(w.Body.String(), expectedMemberGroup) {
		t.Errorf("Expected member group %s not found in response body", expectedMemberGroup)
	} else {
		t.Logf("Expected member group %v  found in response body", expectedMemberGroup)
	}
}


// Create member group

func TestCreateNewMemberGroup(t *testing.T) {
	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	MG := r.Group("/membersgroup") // Assuming your routes are under this group
	{
		MG.POST("/newgroup", middleware.JWTAuth(), controllers.CreateNewMemberGroup)
	}
	req, err := http.NewRequest("POST", "/membersgroup/newgroup", strings.NewReader("membergroup_name=newtodaygrp&membergroup_desc=success&1"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)

	}

}

// Update Member group
func TestUpdateMemberGroup(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	MG := r.Group("/membersgroup") // Assuming your routes are under this group
	{
		MG.POST("/updategroup", middleware.JWTAuth(), controllers.UpdateMemberGroup)
	}
	req, err := http.NewRequest("POST", "/membersgroup/updategroup", strings.NewReader("id=2&membergroup_name=newtodaygrp&membergroup_desc=success"))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// controllers.AUTH = spurtcore.NewInstance(&auth.Option{DB: DB, Token: token, Secret: os.Getenv("JWT_SECRET")})

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)

	}

}

// Delete member group

func TestDeleteMemberGroup(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			// c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	U := r.Group("/membersgroup") // Assuming your routes are under this group
	{
		U.GET("/deletegroup", middleware.JWTAuth(), controllers.DeleteMemberGroup)
	}
	var id = "13"
	req, err := http.NewRequest("GET", "/membersgroup/deletegroup?id="+id, nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)

	}

}

// user list
func TestUserList(t *testing.T) {

	controllers.PackageInitialize()

	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		// c.Set("csrf", "csrftoken")
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	U := r.Group("/settings/users") // Assuming your routes are under this group
	{
		U.GET("/", middleware.JWTAuth(), controllers.Userlist)
	}

	req, err := http.NewRequest("GET", "/settings/users/", nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedUser := "nithya menon"
	if !strings.Contains(w.Body.String(), expectedUser) {
		t.Errorf("Expected user %s not found in response body", expectedUser)
	} else {

		t.Logf("Expected user %v  found in response body", expectedUser)
	}
}

// delete user

func TestDeleteUser(t *testing.T) {

	controllers.PackageInitialize()

	// Create a new Gin router
	r := gin.Default()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	U := r.Group("/settings/users") // Assuming your routes are under this group
	{
		U.GET("/delete-user/:id", middleware.JWTAuth(), controllers.DeleteUser)
	}

	id := "44"
	req, err := http.NewRequest("GET", "/settings/users/delete-user/"+id, nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("myspurtcms", token)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently { // Change the expected status code to StatusMovedPermanently
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	}

}

// Create user
func TestCreateUser(t *testing.T) {

	controllers.PackageInitialize()

	// Create a new Gin router
	r := gin.Default()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	// Define a handler function for the CreateUser route
	U := r.Group("/settings/users")
	{
		U.POST("/createuser", middleware.JWTAuth(), controllers.CreateUser)
	}

	// Create a new HTTP request with the form data encoded
	req, err := http.NewRequest("POST", "/settings/users/createuser", strings.NewReader("user_fname=roja&user_email=&user_mob=123456789&user_name=johndoe&user_pass=password&user_role=1&mem_activestat=1&mem_data_access=1"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("myspurtcms", token)

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently {
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	}

}

// Update User
func TestUpdateUser(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	// Define a handler function for the CreateUser route
	U := r.Group("/settings/users")
	{
		U.POST("/update-user", middleware.JWTAuth(), controllers.UpdateUser)
	}

	// Create a new HTTP request with the form data encoded
	req, err := http.NewRequest("POST", "/settings/users/update-user", strings.NewReader("userid=43&user_fname=rojasri&user_email=john@gmail.com&user_mob=123456789&user_name=johndoe&user_pass=password&user_role=1&mem_activestat=1&mem_data_access=1"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("myspurtcms", token)

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently {
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	}

}

// Member list
func TestMemberList(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			// c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	U := r.Group("/member") // Assuming your routes are under this group
	{
		U.GET("/", middleware.JWTAuth(), controllers.MemberList)
	}

	req, err := http.NewRequest("GET", "/member/", nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)

	}

	expectedMember := "IPL"
	if !strings.Contains(w.Body.String(), expectedMember) {
		t.Errorf("Expected member %s not found in response body", expectedMember)
	} else {

		t.Logf("Expected member %v  found in response body", expectedMember)
	}

}

// Create Member
func TestCreateMember(t *testing.T) {

	controllers.PackageInitialize()

	// Create a new Gin router
	r := gin.Default()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	// Define a handler function for the CreateUser route
	U := r.Group("/member")
	{
		U.POST("/newmember", middleware.JWTAuth(), controllers.CreateMember)
	}

	// Create a new HTTP request with the form data encoded
	req, err := http.NewRequest("POST", "/member/newmember", strings.NewReader("mem_name=IPL&mem_email=ipl@gmail.com&mem_mobile=8824003212&mem_usrname=welcome&mem_pass=mempassword123&mem_activestat=1&membergroupvalue=1"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("myspurtcms", token)

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently {
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	}

}

// Update Member
func TestUpdateMember(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	// Define a handler function for the CreateUser route
	U := r.Group("/member")
	{
		U.POST("/updatemember", middleware.JWTAuth(), controllers.UpdateMember)
	}

	// Create a new HTTP request with the form data encoded
	req, err := http.NewRequest("POST", "/member/updatemember", strings.NewReader("mem_name=IPL 2024&mem_email=ipl@gmail.com&mem_mobile=8824003212&mem_usrname=welcome&mem_pass=mempassword123&mem_activestat=1&membergroupvalue=1"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("myspurtcms", token)

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently {
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	}

}

// Delete Member
func TestDeleteMember(t *testing.T) {

	controllers.PackageInitialize()

	// Create a new Gin router
	r := gin.Default()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	fmt.Println(token, "token")

	U := r.Group("/member") // Assuming your routes are under this group
	{
		U.GET("/deletemember", middleware.JWTAuth(), controllers.DeleteMember)
	}

	id := "44"
	req, err := http.NewRequest("GET", "/member/deletemember?id="+id, nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("myspurtcms", token)

	// controllers.AUTH = spurtcore.NewInstance(&auth.Option{DB: DB, Token: token, Secret: os.Getenv("JWT_SECRET")})

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently { // Change the expected status code to StatusMovedPermanently
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	}

}

//Testing function LanguageList

func TestLanguageList(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)
	fmt.Println(token, "token101")
	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		// c.Set("csrf", "csrftoken")
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	U := r.Group("/settings/languages") // Assuming your routes are under this group
	{
		U.GET("/", middleware.JWTAuth(), controllers.LanguageList)
	}

	req, err := http.NewRequest("GET", "/settings/languages/", nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedUser := "English"
	if !strings.Contains(w.Body.String(), expectedUser) {
		t.Errorf("Expected language %s not found in response body", expectedUser)
	} else {

		t.Logf("Expected language %v  found in response body", expectedUser)
	}

}

// Testing Function of Roleslist//
func TestRolesList(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		// c.Set("csrf", "csrftoken")
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	U := r.Group("/settings/roles") // Assuming your routes are under this group
	{
		U.GET("/", middleware.JWTAuth(), controllers.RoleView)
	}

	req, err := http.NewRequest("GET", "/settings/roles/", nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedUser := "Manager"
	if !strings.Contains(w.Body.String(), expectedUser) {
		t.Errorf("Expected Role %s not found in response body", expectedUser)
	} else {

		t.Logf("Expected Role %v  found in response body", expectedUser)
	}

}

//Testing function of Roles create//

func TestRolesCreate(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	// Define a handler function for the CreateUser route
	U := r.Group("/settings/roles")
	{
		U.POST("/createrole", middleware.JWTAuth(), controllers.RoleCreate)
	}

	// Create a new HTTP request with the form data encoded
	req, err := http.NewRequest("POST", "/settings/roles/createrole", strings.NewReader("rolename=&roledesc=To develope web applications"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("myspurtcms", token)

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

}

// Testing function of Rolesupdate//
func TestRoleUpdate(t *testing.T) {

	controllers.PackageInitialize()
	r := gin.Default()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	// Define a handler function for the CreateUser route
	U := r.Group("/settings/roles")
	{
		U.POST("/updaterole", middleware.JWTAuth(), controllers.RoleUpdate)
	}

	// Create a new HTTP request with the form data encoded
	req, err := http.NewRequest("POST", "/settings/roles/updaterole", strings.NewReader("roleid=47&rolename=Senior Developer&roledesc=rojasri"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("myspurtcms", token)

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

//Testing function of roledelete //

func TestDeleteRole(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	fmt.Println(token, "token")

	U := r.Group("/settings/roles") // Assuming your routes are under this group
	{
		U.GET("/deleterole", middleware.JWTAuth(), controllers.DeleteRole)
	}
	id := "48"
	req, err := http.NewRequest("GET", "/settings/roles/deleterole?id="+id, nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("myspurtcms", token)

	// controllers.AUTH = spurtcore.NewInstance(&auth.Option{DB: DB, Token: token, Secret: os.Getenv("JWT_SECRET")})

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently { // Change the expected status code to StatusMovedPermanently
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	}

}

//Test function of spacelist //

func TestSpaceList(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		// c.Set("csrf", "csrftoken")
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	req, err := http.NewRequest("GET", "/spaces/", nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	w := httptest.NewRecorder()

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedUser := "demospace"
	if !strings.Contains(w.Body.String(), expectedUser) {
		t.Errorf("Expected spacename %s not found in response body", expectedUser)
	} else {

		t.Logf("Expected spacename %v  found in response body", expectedUser)
	}
}

// Test function for Createspace
func TestCreateSpace(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	// Create a new HTTP request with the form data encoded
	req, err := http.NewRequest("POST", "/spaces/createspace", strings.NewReader("spacename=&spacedescription=To develope web applications&catiddd=15"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("myspurtcms", token)

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently {
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	}

}

//Test function of delete space//

func TestDeleteSpace(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	// U := r.Group("/spaces") // Assuming your routes are under this group
	// {
	// 	U.GET("/deletespace", controllers.DeleteSpace)
	// }

	id := "6"
	req, err := http.NewRequest("GET", "/spaces/deletespace?id="+id, nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("myspurtcms", token)

	// controllers.AUTH = spurtcore.NewInstance(&auth.Option{DB: DB, Token: token, Secret: os.Getenv("JWT_SECRET")})

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently { // Change the expected status code to StatusMovedPermanently
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	}

}

//Test function for UpdateSpace//

func TestUpdateSpace(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	// Define a handler function for the CreateUser route
	// U := r.Group("/spaces")
	// {
	// 	U.POST("/updatespace", controllers.UpdateSpace)
	// }

	// Create a new HTTP request with the form data encoded
	req, err := http.NewRequest("POST", "/spaces/updatespace", strings.NewReader("id=5&spacename=galaxyspace&spacedescription=To develope web applications&catiddd=15"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("myspurtcms", token)

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently {
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	}

}

//Test Function for my profile//

func TestMyProfile(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)

		c.Set("userid", 2)

		// c.Set("csrf", "csrftoken")
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	U := r.Group("/settings") // Assuming your routes are under this group
	{
		U.GET("/myprofile", middleware.JWTAuth(), controllers.MyProfile)
	}

	req, err := http.NewRequest("GET", "/settings/myprofile", nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedUser := "ParveshIbrahim"
	if !strings.Contains(w.Body.String(), expectedUser) {
		t.Errorf("Expected username %s not found in response body", expectedUser)
	} else {

		t.Logf("Expected username %v  found in response body", expectedUser)
	}

}

//Test function for Update myprofile//

func TestUpdateProfile(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	// Define a handler function for the CreateUser route
	U := r.Group("/settings")
	{
		U.POST("/updateprofile", middleware.JWTAuth(), controllers.UpdateProfile)
	}

	// Create a new HTTP request with the form data encoded
	req, err := http.NewRequest("POST", "/settings/updateprofile", strings.NewReader("user_fname=&user_email=john@gmail.com&user_mob=123456789&user_name=johndoe&user_pass=password&user_role=1&mem_activestat=1&mem_data_access=1"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("myspurtcms", token)

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently {
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	}

}

// Test function for security change password//

func TestUptPassword(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	// Define a handler function for the CreateUser route
	U := r.Group("/settings")
	{
		U.POST("/updatepassword", middleware.JWTAuth(), controllers.UptPassword)
	}

	// Create a new HTTP request with the form data encoded
	req, err := http.NewRequest("POST", "/settings/updatepassword", strings.NewReader("pass=Admin@123&cpass=Admin@123"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("myspurtcms", token)

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently {
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	}

}

// Content Access list
func TestContentAccessControlList(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)
	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	U := r.Group("/memberaccess") // Assuming your routes are under this group
	{
		U.GET("/", middleware.JWTAuth(), controllers.ContentAccessControlList)
	}

	req, err := http.NewRequest("GET", "/memberaccess/", nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)

	}

	expectedMemberAccess := "ccc"
	if !strings.Contains(w.Body.String(), expectedMemberAccess) {
		t.Errorf("Expected member access %s not found in response body", expectedMemberAccess)
	} else {

		t.Logf("Expected member access %v  found in response body", expectedMemberAccess)
	}

}

// Create Content Access

func TestGrantContentAccessControl(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	MG := r.Group("/memberaccess") // Assuming your routes are under this group
	{
		MG.POST("/grant-accesscontrol", middleware.JWTAuth(), controllers.GrantContentAccessControl)
	}
	formDate := url.Values{}

	formDate.Set("title", "Today Access Permission")
	formDate.Set("spaces", `{"spaces":["1"]}`)
	formDate.Set("pages", `{"pages":[]}`)
	formDate.Set("subpages", `{"subpages":[]}`)
	formDate.Set("pagegroups", `{"pagegroups":[]}`)
	formDate.Set("channels", `{"channels":[]}`)
	formDate.Set("entries", `{"channelEntries":[]}`)
	formDate.Set("memgrps", `{"access_granted_memgrps":["1"]}`)

	req, err := http.NewRequest("POST", "/memberaccess/grant-accesscontrol", strings.NewReader(formDate.Encode()))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// controllers.AUTH = spurtcore.NewInstance(&auth.Option{DB: DB, Token: token, Secret: os.Getenv("JWT_SECRET")})

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)

	}

}

// Update Content Access
func TestUpdateContentAccessControl(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)
	MG := r.Group("/memberaccess") // Assuming your routes are under this group
	{
		MG.POST("/update-accesscontrol", middleware.JWTAuth(), controllers.UpdateContentAccessControl)
	}

	formDate := url.Values{}

	formDate.Set("content_access_id", "23")
	formDate.Set("title", "Today Access Permission success")
	formDate.Set("spaces", `{"spaces":["1"]}`)
	formDate.Set("pages", `{"pages":[]}`)
	formDate.Set("subpages", `{"subpages":[]}`)
	formDate.Set("pagegroups", `{"pagegroups":[]}`)
	formDate.Set("channels", `{"channels":[]}`)
	formDate.Set("entries", `{"channelEntries":[]}`)
	formDate.Set("memgrps", `{"access_granted_memgrps":["1","33"]}`)

	req, err := http.NewRequest("POST", "/memberaccess/update-accesscontrol", strings.NewReader(formDate.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// controllers.AUTH = spurtcore.NewInstance(&auth.Option{DB: DB, Token: token, Secret: os.Getenv("JWT_SECRET")})

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)

	}

}

// Delete Content Access

func TestDeleteAccessControl(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			// c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	U := r.Group("/memberaccess") // Assuming your routes are under this group
	{
		U.GET("/delete-accesscontrol/:accessId", middleware.JWTAuth(), controllers.DeleteAccessControl)
	}
	var id = "23"
	req, err := http.NewRequest("GET", "/memberaccess/delete-accesscontrol/"+id, nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)

	}

}

// Channel list
func TestChannelList(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	U := r.Group("/channels") // Assuming your routes are under this group
	{
		U.GET("/", middleware.JWTAuth(), controllers.ChannelList)
	}

	req, err := http.NewRequest("GET", "/channels/", nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)

	}

	expectedChannel := "Welcome Today 133"
	if !strings.Contains(w.Body.String(), expectedChannel) {
		t.Errorf("Expected Channel %s not found in response body", expectedChannel)
	} else {

		t.Logf("Expected Channel %v  found in response body", expectedChannel)
	}

	// if !strings.Contains(w.Body.String(), expectedChannel) {
	// 	t.Logf("Expected Channel %v not  found in response body", expectedChannel)

	// } else {
	// 	t.Errorf("Expected Channel %s found in response body", expectedChannel)
	// }

}

// Create Channel

func TestCreateChannel(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	MG := r.Group("/channels") // Assuming your routes are under this group
	{
		MG.POST("/newchannel", middleware.JWTAuth(), controllers.CreateChannel)
	}
	formData := url.Values{}

	formData.Set("channelname", "Welcome")
	formData.Set("channeldesc", "Welcome All")
	formData.Set("sectionvalue", `{"sections":[{"SectionId":"13","SectionNewId":"34","SectionName":"New Section","MasterFieldId":"0","OrderIndex":"0"}]}`)
	// formData.Set("fval", `{"fiedlvalue":[{"MasterFieldId":"0","FieldId":"14","NewFieldId":"34","MasterFieldId":"0","FieldName":"Text"}]}`)
	formData.Set("categoryvalue", "[12]")
	req, err := http.NewRequest("POST", "/channels/newchannel", strings.NewReader(formData.Encode()))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// controllers.AUTH = spurtcore.NewInstance(&auth.Option{DB: DB, Token: token, Secret: os.Getenv("JWT_SECRET")})

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)

	}

}

// Update Channel
func TestUpdateChannel(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)
	MG := r.Group("/channels") // Assuming your routes are under this group
	{
		MG.POST("/updatechannel", middleware.JWTAuth(), controllers.UpdateChannel)
	}

	formData := url.Values{}

	formData.Set("id", "111")
	formData.Set("channelname", "Welcome Today Test")
	formData.Set("channeldesc", "Welcome")
	formData.Set("sectionvalue", `{"sections":[{"SectionId":"13","SectionNewId":"34","SectionName":"New Section","MasterFieldId":"0","OrderIndex":"0"}]}`)
	// formData.Set("fval", `{"fiedlvalue":[{"MasterFieldId":"0","FieldId":"14","NewFieldId":"34","MasterFieldId":"0","FieldName":"Text"}]}`)
	formData.Set("categoryvalue", "12")

	req, err := http.NewRequest("POST", "/channels/updatechannel", strings.NewReader(formData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// controllers.AUTH = spurtcore.NewInstance(&auth.Option{DB: DB, Token: token, Secret: os.Getenv("JWT_SECRET")})

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)

	}

}

// Delete Channel

func TestDeleteChannel(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			// c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	U := r.Group("/channels") // Assuming your routes are under this group
	{
		U.GET("/deletechannel", middleware.JWTAuth(), controllers.DeleteChannel)
	}
	var id = "105"
	req, err := http.NewRequest("GET", "/channels/deletechannel?id="+id, nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)

	}

}

//Test Function for PagesList//

func TestPagesList(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		// c.Set("csrf", "csrftoken")
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	// U := r.Group("/spaces")
	// {
	// 	U.GET("/:id", controllers.PagesList)
	// }
	id := "4"
	req, err := http.NewRequest("GET", "/spaces/"+id, nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedUser := "Protecting Medium’s mission and authentic writers means doing more to protect quality"
	if !strings.Contains(w.Body.String(), expectedUser) {
		t.Errorf("Expected pagename %s not found in response body", expectedUser)
	} else {

		t.Logf("Expected pagename %v  found in response body", expectedUser)
	}

}

//Test function for insertpage//

func TestInsertPage(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	// Define a handler function for the CreateUser route
	// U := r.Group("/spaces")
	// {
	// 	U.POST("/pagecreate", controllers.InsertPage)
	// }
	formData := url.Values{}

	formData.Set("spaceid", "123")
	formData.Set("save", "save")
	formData.Set("createpages", `{"newpagesarr":[{"Name":"Page unittest","Content":"<p>Content 1</p>"}]}`)

	req, err := http.NewRequest("POST", "/spaces/pagecreate", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("myspurtcms", token)

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

}

// Channel Entry list
func TestAllEntries(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	U := r.Group("/channel") // Assuming your routes are under this group
	{
		U.GET("/entrylist/", middleware.JWTAuth(), controllers.AllEntries)
	}

	req, err := http.NewRequest("GET", "/channel/entrylist/", nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)

	}

	expectedChannel := "Today"
	if !strings.Contains(w.Body.String(), expectedChannel) {
		t.Errorf("Expected Channel Entries %s not found in response body", expectedChannel)
	} else {

		t.Logf("Expected Channel Entries %v  found in response body", expectedChannel)
	}

}

// Channel Entry list
func TestEntries(t *testing.T) {

	controllers.PackageInitialize()
	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	U := r.Group("/channel") // Assuming your routes are under this group
	{
		U.GET("/entrylist/:id", middleware.JWTAuth(), controllers.Entries)
	}

	id := "97"

	req, err := http.NewRequest("GET", "/channel/entrylist/"+id, nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)

	}

	expectedChannel := "sdfasdvvv"
	if !strings.Contains(w.Body.String(), expectedChannel) {
		t.Errorf("Expected Channel Entries %s not found in response body", expectedChannel)
	} else {

		t.Logf("Expected Channel Entries %v  found in response body", expectedChannel)
	}

}

// Create Channel Entry
func TestCreateEntry(t *testing.T) {

	controllers.PackageInitialize()

	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)
	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			// c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	MG := r.Group("/channel") // Assuming your routes are under this group
	{
		MG.GET("/newentry", middleware.JWTAuth(), controllers.EditEntry)
	}

	req, err := http.NewRequest("GET", "/channel/newentry", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)

	}

}

// Change Entry status

func TestEntryStatus(t *testing.T) {

	// Create a new Gin router
	r := gin.Default()

	controllers.PackageInitialize()
	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)
	MG := r.Group("/channel") // Assuming your routes are under this group
	{
		MG.POST("/changestatus/:id", middleware.JWTAuth(), controllers.EntryStatus)
	}

	// status publish functionality
	id := "102"

	req, err := http.NewRequest("POST", "/channel/changestatus/"+id, strings.NewReader("entryid=111&cname=Blog&status=0"))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// controllers.AUTH = spurtcore.NewInstance(&auth.Option{DB: DB, Token: token, Secret: os.Getenv("JWT_SECRET")})

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code tohttp.StatusMovedPermanently
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)

	}

}

// Edit Channel Entry
func TestEditDetails(t *testing.T) {

	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	controllers.PackageInitialize()
	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			// c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	MG := r.Group("/channel") // Assuming your routes are under this group
	{
		MG.GET("/editentrydetails/:channelname/:id", middleware.JWTAuth(), controllers.EditDetails)
	}

	id := "110"
	channelname := "Blog"

	req, err := http.NewRequest("GET", "/channel/editentrydetails/"+channelname+"/"+id, nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)

	}

	expectedChannel := "4 ways to create enums in Go"
	if !strings.Contains(w.Body.String(), expectedChannel) {
		t.Errorf("Expected Image %s not found in response body", expectedChannel)
	} else {

		t.Logf("Expected Image %v  found in response body", expectedChannel)
	}

}

// Publish Channel Entry

func TestPublishEntry(t *testing.T) {

	// Create a new Gin router
	r := gin.Default()

	controllers.PackageInitialize()
	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)
	MG := r.Group("/channel") // Assuming your routes are under this group
	{
		MG.POST("/publishentry/:id", controllers.PublishEntry)
	}

	// status publish functionality
	id := "108"
	// req, err := http.NewRequest("POST", "/channel/publishentry/"+id, strings.NewReader("id=105&cname=Welcome&title=Today&text=today test case pass&image=/storage/media/deer(680*340)(2024-03-27T11:08:49.783Z).jpg&status=1"))

	req, err := http.NewRequest("POST", "/channel/draftentry/"+id, strings.NewReader("id=105&cname=Welcome&title=Today&text=today test case pass&image=/storage/media/deer(680*340)(2024-03-27T11:08:49.783Z).jpg&status=1"))

	//  create entry
	// id := "undefined"
	// req, err := http.NewRequest("POST", "/channel/publishentry/"+id, strings.NewReader("id=98&cname=Welcome&title=Today&text=today test case pass&image=/storage/media/deer(680*340)(2024-03-27T11:08:49.783Z).jpg&status=1"))

	// draft functionality
	// req, err := http.NewRequest("POST", "/channel/publishentry/"+id, strings.NewReader("id=98&cname=Welcome&title=Today&text=today test case pass&image=/storage/media/deer(680*340)(2024-03-27T11:08:49.783Z).jpg&status=0"))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// controllers.AUTH = spurtcore.NewInstance(&auth.Option{DB: DB, Token: token, Secret: os.Getenv("JWT_SECRET")})

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code tohttp.StatusMovedPermanently
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)

	}

}

// Edit Channel Entry
func TestEditEntry(t *testing.T) {

	controllers.PackageInitialize()
	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			// c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	MG := r.Group("/channel") // Assuming your routes are under this group
	{
		MG.GET("/copyentry/:id", middleware.JWTAuth(), controllers.EditEntry)
	}

	id := "108"

	req, err := http.NewRequest("GET", "/channel/copyentry/"+id, nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)

	}

}

// Delete Channel Entry

func TestDeleteEntries(t *testing.T) {

	controllers.PackageInitialize()
	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			// c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	U := r.Group("/channel") // Assuming your routes are under this group
	{
		U.GET("/deleteentries/", middleware.JWTAuth(), controllers.DeleteEntries)
	}
	var id = "105"
	var cname = "NEW"
	var page = "1"
	req, err := http.NewRequest("GET", "/channel/deleteentries/?id="+id+"&cname="+cname+"&page"+page, nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)

	}

}

// Category group list
func TestCategoryGroupList(t *testing.T) {

	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	controllers.PackageInitialize()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	U := r.Group("/categories") // Assuming your routes are under this group
	{
		U.GET("/", middleware.JWTAuth(), controllers.CategoryGroupList)
	}

	req, err := http.NewRequest("GET", "/categories/", nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)

	}

	expectedCategoryGroup := "default"
	if !strings.Contains(w.Body.String(), expectedCategoryGroup) {
		t.Errorf("Expected Category group name %s not found in response body", expectedCategoryGroup)
	} else {

		t.Logf("Expected Category group name %v  found in response body", expectedCategoryGroup)
	}

}

// Create Category group

func TestCreateCategoryGroup(t *testing.T) {

	// Create a new Gin router
	r := gin.Default()

	controllers.PackageInitialize()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	MG := r.Group("/categories") // Assuming your routes are under this group
	{
		MG.POST("/newcategory", middleware.JWTAuth(), controllers.CreateCategoryGroup)
	}
	req, err := http.NewRequest("POST", "/categories/newcategory", strings.NewReader("category_name=IPL&category_desc=success"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// controllers.AUTH = spurtcore.NewInstance(&auth.Option{DB: DB, Token: token, Secret: os.Getenv("JWT_SECRET")})

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)

	}

}

// Update Category group
func TestUpdateCategoryGroup(t *testing.T) {

	// Create a new Gin router
	r := gin.Default()

	controllers.PackageInitialize()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)
	MG := r.Group("/categories") // Assuming your routes are under this group
	{
		MG.POST("/updatecategory", middleware.JWTAuth(), controllers.UpdateCategoryGroup)
	}
	req, err := http.NewRequest("POST", "/categories/updatecategory", strings.NewReader("category_id=33&category_name=Today IPL&category_desc=success"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// controllers.AUTH = spurtcore.NewInstance(&auth.Option{DB: DB, Token: token, Secret: os.Getenv("JWT_SECRET")})

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)

	}

}

// Delete Category group

func TestDeleteCategoryGroup(t *testing.T) {

	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	controllers.PackageInitialize()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)
	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			// c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	U := r.Group("/categories") // Assuming your routes are under this group
	{
		U.GET("/deletecategory/:id", middleware.JWTAuth(), controllers.DeleteCategoryGroup)
	}
	var id = "33"
	req, err := http.NewRequest("GET", "/categories/deletecategory/"+id, nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)

	}

}

// Check Category group name is already exists
func TestCheckCategoryGroupName(t *testing.T) {

	// Create a new Gin router
	r := gin.Default()

	controllers.PackageInitialize()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)
	MG := r.Group("/categories") // Assuming your routes are under this group
	{
		MG.POST("/checkcategoryname", middleware.JWTAuth(), controllers.CheckCategoryGroupName)
	}
	req, err := http.NewRequest("POST", "/categories/checkcategoryname", strings.NewReader("category_id=20&category_name=default"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// controllers.AUTH = spurtcore.NewInstance(&auth.Option{DB: DB, Token: token, Secret: os.Getenv("JWT_SECRET")})

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)

	}

}

// Sub Category list
func TestAddCategory(t *testing.T) {

	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	controllers.PackageInitialize()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	U := r.Group("/categories") // Assuming your routes are under this group
	{
		U.GET("/addcategory/:id", middleware.JWTAuth(), controllers.AddCategory)
	}
	var pcid = "20"
	req, err := http.NewRequest("GET", "/categories/addcategory/"+pcid, nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)

	}

	expectedSubcategoryname := "Tech"
	if !strings.Contains(w.Body.String(), expectedSubcategoryname) {
		t.Errorf("Expected Sub Category name %s not found in response body", expectedSubcategoryname)
	} else {

		t.Logf("Expected Sub Category name %v  found in response body", expectedSubcategoryname)
	}

}

// Create Sub Category

func TestAddSubCategory(t *testing.T) {

	// Create a new Gin router
	r := gin.Default()

	controllers.PackageInitialize()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	MG := r.Group("/categories") // Assuming your routes are under this group
	{
		MG.POST("/addsubcategory/:categoryid", middleware.JWTAuth(), controllers.AddSubCategory)
	}
	var pcid = "20"
	req, err := http.NewRequest("POST", "/categories/addsubcategory/"+pcid, strings.NewReader("subcategoryid=20&cname=T20&cdesc=Enjoyment"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// controllers.AUTH = spurtcore.NewInstance(&auth.Option{DB: DB, Token: token, Secret: os.Getenv("JWT_SECRET")})

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)

	}

}

// Update Sub Category
func TestUpdateSubCategory(t *testing.T) {

	// Create a new Gin router
	r := gin.Default()

	controllers.PackageInitialize()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	MG := r.Group("/categories") // Assuming your routes are under this group
	{
		MG.POST("/editsubcategory", middleware.JWTAuth(), controllers.UpdateSubCategory)
	}
	req, err := http.NewRequest("POST", "/categories/editsubcategory", strings.NewReader("id=36&pcategoryid=22&cname=Today IPL&cdesc=success"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// controllers.AUTH = spurtcore.NewInstance(&auth.Option{DB: DB, Token: token, Secret: os.Getenv("JWT_SECRET")})

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)

	}

}

// Delete Sub Category

func TestDeleteSubCategory(t *testing.T) {

	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	controllers.PackageInitialize()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)
	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			// c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	U := r.Group("/categories") // Assuming your routes are under this group
	{
		U.GET("/removecategory", middleware.JWTAuth(), controllers.DeleteSubCategory)
	}
	var id = "36"
	var pcid = "20"
	req, err := http.NewRequest("GET", "/categories/removecategory?id="+id+"&&categoryid="+pcid, nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)

	}

}

// Dashboard list

func TestDashboardView(t *testing.T) {

	// Create a new Gin router
	r := gin.Default()

	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	controllers.PackageInitialize()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)
	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	U := r.Group("/dashboard") // Assuming your routes are under this group
	{
		U.GET("/", middleware.JWTAuth(), controllers.DashboardView)
	}

	req, err := http.NewRequest("GET", "/dashboard/", nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d ", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)

	}

	expectedList := "nithya"
	if strings.Contains(w.Body.String(), expectedList) {
		t.Errorf("Expected Image %s not found in response body", expectedList)
	} else {

		t.Logf("Expected Image %v  found in response body", expectedList)
	}

}

// Test Personalize function

func TestPersonalization(t *testing.T) {

	// Create a new Gin router
	r := gin.Default()

	r.Use(func(c *gin.Context) {

		c.Set("userid", 2)
		c.Next()
	})
	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	controllers.PackageInitialize()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		// c.Set("userid", 2)
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	U := r.Group("/settings") // Assuming your routes are under this group
	{
		U.GET("/personalize/", middleware.JWTAuth(), controllers.Personalization)
	}

	req, err := http.NewRequest("GET", "/settings/personalize/", nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)

	}

}

// test function of personalizationupdate
func TestPersonalizationUpdate(t *testing.T) {

	// Create a new Gin router
	r := gin.Default()
	r.Use(func(c *gin.Context) {

		c.Set("userid", 2)
		c.Next()
	})

	controllers.PackageInitialize()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)
	// Define a handler function for the CreateUser route
	U := r.Group("/settings/personalize")
	{
		U.POST("/update", middleware.JWTAuth(), controllers.PersonalizationUpdate)
	}

	// Create a new HTTP request with the form data encoded
	req, err := http.NewRequest("POST", "/settings/personalize/update", strings.NewReader("userid=2&color-change=rgb(57, 129, 234)&logo_imgpath=/public/img/logo.svg&expandlogo_imgpath=/public/img/logo-bg.svg"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("myspurtcms", token)

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusMovedPermanently {
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	}

}

// Test function of data
func TestData(t *testing.T) {

	r := gin.Default()

	r.Use(func(c *gin.Context) {

		c.Set("userid", 1)
		c.Next()
	})
	r.Use(lang.TranslateHandler)

	var htmlfiles []string

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlfiles = append(htmlfiles, path)
		}
		return nil
	})
	r.LoadHTMLFiles(htmlfiles...)

	controllers.PackageInitialize()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	json_folder := "locales/"

	var default_lang models.TblLanguage

	translation, err := lang.LoadTranslation(default_lang.JsonPath)

	if err != nil {

		translation, err = lang.LoadTranslation(json_folder + "en.json")

		if err != nil {

			c.AbortWithError(http.StatusInternalServerError, err)

			return
		}
	}

	r.Use(func(c *gin.Context) {
		c.Set("translation", translation)
		// c.Set("userid", 2)
		c.Next()
	})

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	U := r.Group("/settings/data")
	{
		U.GET("/", middleware.JWTAuth(), controllers.Data)
	}

	req, err := http.NewRequest("GET", "/settings/data/", nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK { // Change the expected status code to http.StatusOK
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)

	}

}

func TestImportData(t *testing.T) {
	// Create a new Gin context for testing
	r := gin.Default()

	r.Static("/storage", "./storage")

	r.Use(func(c *gin.Context) {

		c.Set("userid", 1)
		c.Next()
	})

	controllers.PackageInitialize()

	token, _, _ := controllers.NewAuth.Checklogin("Admin", "Admin@123", controllers.TenantId)

	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))

	r.Use(sessions.Sessions(os.Getenv("SESSION_KEY3"), store))

	r.Use(csrf.Middleware(csrf.Options{

		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()

		},
	}))

	U := r.Group("/settings/data")
	{
		U.POST("/importdata", middleware.JWTAuth(), controllers.ImportData)
	}

	req, err := http.NewRequest("POST", "/settings/data/importdata", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	// Create a mock image file
	imageFile, err := zipWriter.Create("image.jpg")
	if err != nil {

		log.Println(err)
	}

	// Open the image file
	imgFile, err := os.Open("storage/media/dog.jpg")
	if err != nil {
		log.Println(err)
	}
	defer imgFile.Close()

	// Write mock image data from the image file
	_, err = io.Copy(imageFile, imgFile)
	if err != nil {
		log.Println(err)
	}
	err = zipWriter.Close()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(buf, "zipfile")

	formFile := make(map[string][]*multipart.FileHeader)
	formFile["image"] = []*multipart.FileHeader{
		{
			Filename: "mock.zip",
			Size:     int64(buf.Len()),
		},
	}

	req.MultipartForm = &multipart.Form{
		File: formFile,
	}

	req.PostForm = url.Values{
		"id":        []string{"123"},
		"fielddata": []string{`[{"field":"name","type":"string"},{"field":"age","type":"int"}]`},
	}
	req.Header.Add("myspurtcms", token)

	// Create a new session
	session, err := store.Get(req, os.Getenv("SESSION_KEY"))
	if err != nil {
		t.Errorf("Failed to get session: %v", err)
		return
	}

	// Set the token in the session
	session.Values["token"] = token
	err = session.Save(req, httptest.NewRecorder())
	if err != nil {
		t.Errorf("Failed to save session: %v", err)
		return
	}
	// Call the ImportData function
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the HTTP status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

}
