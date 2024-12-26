package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	storagecontroller "spurt-cms/storage-controller"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	cat "github.com/spurtcms/categories"
	csrf "github.com/utrack/gin-csrf"
)

var CategoriesImagepath string

type CategoriesFilter struct {
	Keyword  string
	Category string
	Status   string
	FromDate string
	ToDate   string
}

// Category Parent List
func CategoryGroupList(c *gin.Context) {

	var (
		limt   int
		offset int
		filter cat.Filter
	)

	//get values from url query
	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")

	var filterflag bool
	if filter.Keyword != "" {
		filterflag = true
	} else {
		filterflag = false
	}

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	var FinalCategoriesList []cat.TblCategories

	permisison, perr := NewAuth.IsGranted("Categories Group", auth.Read, TenantId)
	if perr != nil {
		ErrorLog.Printf("categorygrouplist authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("categorygroup authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		CategoryConfig.DataAccess = c.GetInt("dataaccess")
		CategoryConfig.UserId = c.GetInt("userid")

		categorylist, Total_categories, err := CategoryConfig.CategoryGroupList(limt, offset, cat.Filter(filter), TenantId)
		if err != nil {
			ErrorLog.Printf("category group list  error: %s", err)
		}
		for _, val := range categorylist {

			if !val.ModifiedOn.IsZero() {
				val.DateString = val.ModifiedOn.In(TZONE).Format(Datelayout)
			} else {
				val.DateString = val.CreatedOn.In(TZONE).Format(Datelayout)
			}
			FinalCategoriesList = append(FinalCategoriesList, val)

		}

		// _, chcount, _ := category.CategoryGroupList(limt, offset, categories.Filter{})

		//pagination calc
		paginationendcount := len(categorylist) + offset
		paginationstartcount := offset + 1
		Previous, Next, PageCount, Page := Pagination(pageno, int(Total_categories), limt)

		menu := NewMenuController(c)

		translate, _ := TranslateHandler(c)

		ModuleName, TabName, _ := ModuleRouteName(c)

		uper, _ := NewAuth.IsGranted("Categories Group", auth.Update, TenantId)

		dper, _ := NewAuth.IsGranted("Categories Group", auth.Delete, TenantId)

		c.HTML(200, "categorygroup.html", gin.H{"Pagination": PaginationData{
			NextPage:     pageno + 1,
			PreviousPage: pageno - 1,
			TotalPages:   PageCount,
			TwoAfter:     pageno + 2,
			TwoBelow:     pageno - 2,
			ThreeAfter:   pageno + 3,
		}, "Menu": menu, "translate": translate, "Searchtrue": filterflag, "lencategory": Total_categories, "csrf": csrf.GetToken(c), "categorylist": FinalCategoriesList, "Count": Total_categories, "Paginationendcount": paginationendcount, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Filter": filter, "Paginationstartcount": paginationstartcount, "HeadTitle": translate.Categories, "Categoriesmenu": true, "title": ModuleName, "linktitle": ModuleName, "Limit": limt, "Tabmenu": TabName, "permission": uper, "dpermission": dper})

	}

}

// Add Category Group
func CreateCategoryGroup(c *gin.Context) {

	userid := c.GetInt("userid")

	categorygroup := cat.CategoryCreate{
		CategoryName: c.PostForm("category_name"),
		Description:  c.PostForm("category_desc"),
		CreatedBy:    userid,
		TenantId:     TenantId,
	}

	permisison, perr := NewAuth.IsGranted("Categories Group", auth.Create, TenantId)
	if perr != nil {
		ErrorLog.Printf("createcategorygroup authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("categorygroup authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		err := CategoryConfig.CreateCategoryGroup(categorygroup)
		if strings.Contains(fmt.Sprint(err), "given some values is empty") {
			ErrorLog.Printf("createcategorygroup mandatory field error: %s", err)
			c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
			c.Redirect(301, "/categories/")
			return
		}

		if err != nil {
			ErrorLog.Printf("createcategorygroup error: %s", perr)
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			c.Redirect(http.StatusMovedPermanently, "/categories/")
			return
		}

		c.SetCookie("get-toast", "Category Group Created Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(http.StatusMovedPermanently, "/categories/")

	}
}

// Category Group Update
func UpdateCategoryGroup(c *gin.Context) {

	//get data from html form data
	pageno := c.Request.PostFormValue("catpageno")
	id, _ := strconv.Atoi(c.PostForm("category_id"))
	userid := c.GetInt("userid")

	var url string
	if pageno != "" {
		url = "/categories?page=" + pageno
	} else {
		url = "/categories/"

	}

	categorygroup := cat.CategoryCreate{
		Id:           id,
		CategoryName: c.PostForm("category_name"),
		Description:  c.PostForm("category_desc"),
		ModifiedBy:   userid,
	}

	permisison, perr := NewAuth.IsGranted("Categories Group", auth.Update, TenantId)
	if perr != nil {
		ErrorLog.Printf("updatecategorygroup authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("categorygroup authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		err := CategoryConfig.UpdateCategoryGroup(categorygroup, TenantId)
		if strings.Contains(fmt.Sprint(err), "given some values is empty") {
			ErrorLog.Printf("updatecategorygroup mandatory field error: %s", perr)
			c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
			c.Redirect(301, url)
			return
		}

		if err != nil {
			ErrorLog.Printf("updatecategorygroup error: %s", perr)
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			c.Redirect(http.StatusMovedPermanently, url)
			return
		}

		c.SetCookie("get-toast", "Category Group Updated Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(http.StatusMovedPermanently, url)

	}
}

// Delete Category Group
func DeleteCategoryGroup(c *gin.Context) {

	categoryId, _ := strconv.Atoi(c.Param("id"))
	pageno := c.Query("page")
	userid := c.GetInt("userid")

	var url string
	if pageno != "" {
		url = "/categories?page=" + pageno
	} else {
		url = "/categories/"

	}

	permisison, perr := NewAuth.IsGranted("Categories Group", auth.Delete, TenantId)
	if perr != nil {
		ErrorLog.Printf("deletecategorygroup authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("categorygroup authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		err := CategoryConfig.DeleteCategoryGroup(categoryId, userid, TenantId)

		if strings.Contains(fmt.Sprint(err), "given some values is empty") {
			ErrorLog.Printf("deletecategorygroup mandatory field error: %s", perr)
			c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
			c.Redirect(301, url)
			return
		}

		if err != nil {
			ErrorLog.Printf("deletecategorygroup error: %s", perr)
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			c.Redirect(301, url)
			return
		}

		c.SetCookie("get-toast", "Category Group Deleted Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, url)

	}

}

/*Add sub Category list*/
func AddCategory(c *gin.Context) {

	var (
		limt                int
		offset              int
		filter              CategoriesFilter
		FinalCategoriesList []cat.TblCategories
	)

	//get data from html form data
	parent_id, _ := strconv.Atoi(c.Param("id"))
	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")

	var filterflag bool
	if filter.Keyword != "" {
		filterflag = true
	} else {
		filterflag = false
	}

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	permisison, perr := NewAuth.IsGranted("Categories", auth.Read, TenantId)
	if perr != nil {
		ErrorLog.Printf("Addcategory authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("categorygroup authorization error")
		c.Redirect(301, "/403-page")
		return
	}
	uper, _ := NewAuth.IsGranted("Categories", auth.Update, TenantId)

	dper, _ := NewAuth.IsGranted("Categories", auth.Delete, TenantId)
	_, _, _, Total_categories, err := CategoryConfig.ListCategory(0, 0, cat.Filter(filter), parent_id, TenantId)
	categorylist, _, _, _, _ := CategoryConfig.ListCategory(offset, limt, cat.Filter(filter), parent_id, TenantId)
	if err != nil {
		ErrorLog.Printf("categorylist error: %s", err)
	}

	for _, val := range categorylist {
		if !val.ModifiedOn.IsZero() {
			val.DateString = val.ModifiedOn.In(TZONE).Format(Datelayout)
		} else {
			val.DateString = val.CreatedOn.In(TZONE).Format(Datelayout)
		}
		FinalCategoriesList = append(FinalCategoriesList, val)
	}

	// _, _, _, catcount, _ := CategoryConfig.ListCategory(offset, limt, cat.Filter{}, parent_id)
	_, FinalModalCategoriesList, categorys, catcount, _ := CategoryConfig.ListCategory(offset, limt, cat.Filter{}, parent_id, TenantId)
	// _, FinalModalCategoriesList, _, _, _ := CategoryConfig.ListCategory(offset, limt, cat.Filter{}, parent_id)

	//pagination calc
	paginationendcount := len(categorylist) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(pageno, int(Total_categories), limt)

	translate, _ := TranslateHandler(c)

	menu := NewMenuController(c)

	if err != nil {
		ErrorLog.Printf("Getmedia error in addcategory error: %s", err)
	}

	ModuleName, TabName, _ := ModuleRouteName(c)

	if TabName == "Categories" {
		TabName = "Categories Group"
	}

	if c.Request.Method == "POST" {
		c.JSON(200, gin.H{"CategoryList": FinalCategoriesList, "Count": Total_categories, "translate": translate})
		return
	}

	selectedtype, _ := GetSelectedType()
	c.HTML(200, "category.html", gin.H{"Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3,
	}, "Menu": menu, "translate": translate, "Searchtrue": filterflag, "catcount": catcount, "csrf": csrf.GetToken(c), "HeadTitle": translate.Categories, "Category": categorys, "Previous": Previous, "Next": Next, "PageCount": PageCount, "Page": Page, "CurrentPage": pageno, "categoryid": parent_id, "ParentCategoryName": categorys.CategoryName, "CategoryList": FinalCategoriesList, "Count": Total_categories, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "Filter": filter, "title": "Categories", "linktitle": "Categories", "FinalModalCategoriesList": FinalModalCategoriesList, "heading": categorys.CategoryName, "Back": "/categories", "Categoriesmenu": true, "Limit": limt, "Tabmenu": TabName, "ModuleName": ModuleName, "StorageType": selectedtype.SelectedType, "permission": uper, "dpermission": dper})

}

// Add Children Category
func AddSubCategory(c *gin.Context) {

	//get value from html form data
	id := (c.PostForm("categoryid"))
	imagedata := c.PostForm("categoryimages")
	userid := c.GetInt("userid")
	var ParentId, _ = strconv.Atoi(c.PostForm("subcategoryid"))

	var imageName, imagePath string

	storagetype, err := GetSelectedType()

	if err != nil {
		ErrorLog.Printf("error get storage type error: %s", err)
	}

	if storagetype.SelectedType == "local" {

		if imagedata != "" {

			_, imagePath, err = ConvertBase64(imagedata, strings.TrimPrefix(storagetype.Local+"/categories/addcategory", "/"))

			if err != nil {
				ErrorLog.Printf("error get storage type error: %s", err)
			}
		}

	} else if storagetype.SelectedType == "aws" {

		tenantDetails, err := NewTeam.GetTenantDetails(TenantId)
		if err != nil {
			ErrorLog.Printf("error get storage type error: %s", err)
		}

		var imageByte []byte

		if imagedata != "" {
			imageName, imagePath, imageByte, err = ConvertBase64toByte(imagedata, "categories/addcategory")
			if err != nil {
				fmt.Println(err)
			}

			imagePath = tenantDetails.S3FolderName + imagePath

			uerr := storagecontroller.UploadCropImageS3(imageName, imagePath, imageByte)
			if uerr != nil {
				c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
				c.Redirect(301, "/categories/")
				return
			}
		}
	}

	subcategory := cat.CategoryCreate{
		CategoryName: c.PostForm("cname"),
		Description:  c.PostForm("cdesc"),
		ImagePath:    imagePath,
		ParentId:     ParentId,
		CreatedBy:    userid,
		TenantId:     TenantId,
	}

	permisison, perr := NewAuth.IsGranted("Categories", auth.Create, TenantId)
	if perr != nil {
		ErrorLog.Printf("Addsubcategory authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("categorygroup authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		err := CategoryConfig.AddCategory(subcategory)
		if strings.Contains(fmt.Sprint(err), "given some values is empty") {
			ErrorLog.Printf("Addsubcategory mandatory field error: %s", perr)
			c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
			c.Redirect(301, "/categories/")
			return
		}

		if err != nil {
			ErrorLog.Printf("Addsubcategory error: %s", perr)
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			c.Redirect(301, "/categories/addcategory/"+id)
			return
		}

		c.SetCookie("get-toast", "Category Created Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.Redirect(301, "/categories/addcategory/"+id)

	}
}

// Edit Sub Category
func EditSubCategory(c *gin.Context) {

	var id, _ = strconv.Atoi(c.Query("id"))

	permisison, perr := NewAuth.IsGranted("Categories", auth.Read, TenantId)
	if perr != nil {
		ErrorLog.Printf("editsubcategory authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("categorygroup authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		categorys, err := CategoryConfig.GetSubCategoryDetails(id, TenantId)
		if err != nil {
			ErrorLog.Printf("editsubcategory error :%s", err)
		}
		c.JSON(200, categorys)
		CategoriesImagepath = categorys.ImagePath

	}

}

// Update Childern Category
func UpdateSubCategory(c *gin.Context) {

	//get value from html form data
	ParentId, _ := strconv.Atoi(c.PostForm("pcategoryid"))
	imagedata := c.PostForm("image")
	userid := c.GetInt("userid")
	Categoryid, _ := strconv.Atoi(c.PostForm("id"))

	var imageName, imagePath string

	storagetype, err := GetSelectedType()
	if err != nil {
		ErrorLog.Printf("error get storage type error: %s", err)
	}

	if storagetype.SelectedType == "local" {
		if CategoriesImagepath != imagedata {
			_, imagePath, err = ConvertBase64(imagedata, strings.TrimPrefix(storagetype.Local+"/categories/addcategory", "/"))

			if err != nil {
				ErrorLog.Printf("error get storage type error: %s", err)
			}
		} else {
			imagePath = imagedata
		}

	} else if storagetype.SelectedType == "aws" {

		tenantDetails, err := NewTeam.GetTenantDetails(TenantId)
		if err != nil {
			ErrorLog.Printf("error get storage type error: %s", err)
		}

		var imageByte []byte

		if CategoriesImagepath != imagedata {
			imageName, imagePath, imageByte, err = ConvertBase64toByte(imagedata, "categories/addcategory")
			if err != nil {
				fmt.Println(err)
			}

			imagePath = tenantDetails.S3FolderName + imagePath

			uerr := storagecontroller.UploadCropImageS3(imageName, imagePath, imageByte)
			if uerr != nil {
				c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
				c.Redirect(301, "/categories/")
				return
			}
		} else {
			imagePath = imagedata
		}
	}

	subcategory := cat.CategoryCreate{
		Id:           Categoryid,
		CategoryName: c.PostForm("cname"),
		Description:  c.PostForm("cdesc"),
		ImagePath:    imagePath,
		ParentId:     ParentId,
		ModifiedBy:   userid,
	}

	permisison, perr := NewAuth.IsGranted("Categories", auth.Update, TenantId)
	if perr != nil {
		ErrorLog.Printf("updatesubcategory authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("categorygroup authorization error")
		c.JSON(200, gin.H{"value": false, "url": "/403-page"})
		return
	}

	if permisison {

		err := CategoryConfig.UpdateSubCategory(subcategory, TenantId) // update category
		if strings.Contains(fmt.Sprint(err), "given some values is empty") {
			ErrorLog.Printf("updatesubcategory mandatory field error: %s", perr)
			c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
			c.Redirect(301, "/categories/")
			return
		}

		if err != nil {
			ErrorLog.Printf("updatesubcategory error: %s", perr)
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			c.JSON(200, gin.H{"value": false})
			return
		}

		c.SetCookie("get-toast", "Category Updated Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.JSON(200, gin.H{"value": true})

	}
}

// Delete Children Category
func DeleteSubCategory(c *gin.Context) {

	var url string

	//get value from html form data
	parent_id := (c.Query("categoryid"))
	pageno := c.Query("page")
	var categoryid, _ = strconv.Atoi(c.Query("id"))
	userid := c.GetInt("userid")
	if pageno != "" {
		url = "/categories/addcategory/" + parent_id + "?page=" + pageno
	} else {
		url = "/categories/addcategory/" + parent_id

	}

	permisison, perr := NewAuth.IsGranted("Categories", auth.Delete, TenantId)
	if perr != nil {
		ErrorLog.Printf("deletesubcategory authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("categorygroup authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		err := CategoryConfig.DeleteSubCategory(categoryid, userid, TenantId) // delete subcategory
		if strings.Contains(fmt.Sprint(err), "given some values is empty") {
			ErrorLog.Printf("deletesubcategory mandatory error:")
			c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
			c.Redirect(301, "/categories/")
			return
		}

		if err != nil {
			ErrorLog.Printf("deletesubcategory error: %s", perr)
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		} else {
			c.SetCookie("get-toast", "Category Deleted Successfully", 3600, "", "", false, false)
			c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		}

		c.Redirect(301, url)
	}

}

// To Check Group Name id Already Exists
func CheckCategoryGroupName(c *gin.Context) {

	//get value from html form data
	category_id, _ := strconv.Atoi(c.PostForm("category_id"))
	category_name := c.PostForm("category_name")

	permisison, perr := NewAuth.IsGranted("Categories", auth.Read, TenantId)
	if perr != nil {
		ErrorLog.Printf("checkcategorygroupname authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("categorygroup authorization error")
		c.Redirect(301, "/403-page")
		return
	}
	if permisison {

		flg, err := CategoryConfig.CheckCategroyGroupName(category_id, category_name, TenantId)
		if err != nil {
			ErrorLog.Printf("checkcategorygroupname  error: %s", err)
			json.NewEncoder(c.Writer).Encode(false)
			return
		}

		json.NewEncoder(c.Writer).Encode(flg)
	}
}

// SubCategory name check
func CheckSubCategoryName(c *gin.Context) {

	//get value from html form data
	category_id, _ := strconv.Atoi(c.PostForm("parentid"))
	permisison, perr := NewAuth.IsGranted("Categories", auth.Read, TenantId)
	if perr != nil {
		ErrorLog.Printf("CheckSubCategoryName authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("categorygroup authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		categorylistinchk, _, _, _, err := CategoryConfig.ListCategory(0, Limit, cat.Filter{}, category_id, TenantId)
		if err != nil {
			ErrorLog.Printf("checkcategorygroupname authorization error: %s", err)
		}
		var catid []int
		if len(categorylistinchk) != 0 {
			for _, val := range categorylistinchk {
				catid = append(catid, val.Id)
			}
		}

		category_name := c.PostForm("cname")
		categoryid, _ := strconv.Atoi(c.PostForm("categoryid"))
		flg, err := CategoryConfig.CheckSubCategroyName(catid, categoryid, category_name, TenantId)

		if err != nil {
			ErrorLog.Printf("CheckSubCategoryName error: %s", err)
			json.NewEncoder(c.Writer).Encode(false)
			return
		}

		json.NewEncoder(c.Writer).Encode(flg)
	}
}

func MultiSelectCategoryGroupDelete(c *gin.Context) {

	//get value from html form data
	categoryId := c.PostFormArray("categorygrbids[]")
	pageno := c.PostForm("page")
	userid := c.GetInt("userid")
	var url string

	permisison, perr := NewAuth.IsGranted("Categories Group", auth.Delete, TenantId)
	if perr != nil {
		ErrorLog.Printf("CheckSubCategoryName authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("categorygroup authorization error")
		c.JSON(200, gin.H{"value": false, "url": "/403-page"})
		return
	}

	if permisison {
		categoryIntIds := make([]int, len(categoryId))
		for i, id := range categoryId {
			intId, _ := strconv.Atoi(id)
			categoryIntIds[i] = intId
		}

		err := CategoryConfig.MultiSelectDeleteCategoryGroup(categoryIntIds, userid, TenantId) //delete selected category group
		if err != nil {
			ErrorLog.Printf("MultiSelectCategoryGroupDelete error: %s", err)
			c.JSON(200, gin.H{"value": false})
			return
		}

		_, Total_categories, _ := CategoryConfig.CategoryGroupList(0, 0, cat.Filter{}, TenantId)

		if pageno != "" {
			page, _ := strconv.Atoi(pageno)
			page = page - 1
			multi := page * 10
			multiInt64 := int64(multi)
			if Total_categories > multiInt64 {
				url = "/categories?page=" + pageno
			} else {
				pagee, _ := strconv.Atoi(pageno)
				_page := pagee - 1
				pages := strconv.Itoa(_page)
				url = "/categories?page=" + pages
			}
		} else {
			url = "/categories/"
		}
		c.JSON(200, gin.H{"value": true, "url": url})
	}

}

func MultiSelectCategoryDelete(c *gin.Context) {

	//get value from html form data
	categoryId := c.PostFormArray("categoryids[]")
	parent_id := (c.PostForm("parentid"))
	Parentid, _ := strconv.Atoi(parent_id)
	pageno := c.PostForm("page")
	userid := c.GetInt("userid")
	var url string

	permisison, perr := NewAuth.IsGranted("Categories", auth.Delete, TenantId)
	if perr != nil {
		ErrorLog.Printf("CheckSubCategoryName authorization error: %s", perr)
	}
	if !permisison {
		ErrorLog.Printf("category authorization error")
		c.JSON(200, gin.H{"value": false, "url": "/403-page"})
		return
	}
	if permisison {
		categoryIntIds := make([]int, len(categoryId))
		for i, id := range categoryId {
			intId, _ := strconv.Atoi(id)
			categoryIntIds[i] = intId
		}

		err := CategoryConfig.MultiselectSubCategoryDelete(categoryIntIds, userid, TenantId) //delete selected categories
		if err != nil {
			ErrorLog.Printf("MultiSelectCategoriesDelete error: %s", err)
			c.JSON(200, gin.H{"value": false})
			return
		}

		_, _, _, Total_categories, err := CategoryConfig.ListCategory(0, 0, cat.Filter{}, Parentid, TenantId)

		if pageno != "" {
			abc, _ := strconv.Atoi(pageno)
			abc = abc - 1
			multi := abc * 10
			multiInt64 := int64(multi)
			if Total_categories > multiInt64 {
				url = "/categories/addcategory/" + parent_id + "?page=" + pageno
			} else {
				abcd, _ := strconv.Atoi(pageno)
				_page := abcd - 1
				pages := strconv.Itoa(_page)
				url = "/categories/addcategory/" + parent_id + "?page=" + pages
			}
		} else {
			url = "/categories/addcategory/" + parent_id

		}

		c.JSON(200, gin.H{"value": true, "url": url})
	}

}
