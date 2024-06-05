package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	cat "github.com/spurtcms/categories"
	"github.com/spurtcms/pkgcontent/categories"
	csrf "github.com/utrack/gin-csrf"
)

var category categories.Category

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

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	var FinalCategoriesList []cat.TblCategories

	permisison, perr := NewAuth.IsGranted("Categories Group", auth.Read)
	if perr != nil {
		ErrorLog.Printf("categorygrouplist authorization error: %s", perr)
	}

	if permisison {

		categorylist, Total_categories, err := CategoryConfig.CategoryGroupList(limt, offset, cat.Filter(filter))
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

		c.HTML(200, "categorygroup.html", gin.H{"Pagination": PaginationData{
			NextPage:     pageno + 1,
			PreviousPage: pageno - 1,
			TotalPages:   PageCount,
			TwoAfter:     pageno + 2,
			TwoBelow:     pageno - 2,
			ThreeAfter:   pageno + 3,
		}, "Menu": menu, "translate": translate, "lencategory": Total_categories, "csrf": csrf.GetToken(c), "categorylist": FinalCategoriesList, "Count": Total_categories, "Paginationendcount": paginationendcount, "Previous": Previous, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Page": Page, "Filter": filter, "Paginationstartcount": paginationstartcount, "HeadTitle": translate.Categories, "Categoriesmenu": true, "title": ModuleName, "Limit": limt, "Tabmenu": TabName})

		return

	}

	c.Redirect(301, "/403-page")

}

// Add Category Group
func CreateCategoryGroup(c *gin.Context) {

	userid := c.GetInt("userid")

	categorygroup := cat.CategoryCreate{
		CategoryName: c.PostForm("category_name"),
		Description:  c.PostForm("category_desc"),
		CreatedBy:    userid,
	}

	permisison, perr := NewAuth.IsGranted("Categories Group", auth.Create)
	if perr != nil {
		ErrorLog.Printf("createcategorygroup authorization error: %s", perr)
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
		return
	}

	c.Redirect(301, "/403-page")

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

	permisison, perr := NewAuth.IsGranted("Categories Group", auth.Update)
	if perr != nil {
		ErrorLog.Printf("updatecategorygroup authorization error: %s", perr)
	}

	if permisison {

		err := CategoryConfig.UpdateCategoryGroup(categorygroup)
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
		return
	}

	c.Redirect(301, "/403-page")

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

	permisison, perr := NewAuth.IsGranted("Categories Group", auth.Delete)
	if perr != nil {
		ErrorLog.Printf("deletecategorygroup authorization error: %s", perr)
	}

	if permisison {

		err := CategoryConfig.DeleteCategoryGroup(categoryId, userid)
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
		return

	}

	c.Redirect(301, "/403-page")

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

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	permisison, perr := NewAuth.IsGranted("Categories", auth.Read)
	if perr != nil {
		ErrorLog.Printf("Addcategory authorization error: %s", perr)
	}

	if permisison {

		_, _, _, Total_categories, err := CategoryConfig.ListCategory(0, 0, cat.Filter(filter), parent_id)
		categorylist, _, _, _, _ := CategoryConfig.ListCategory(offset, limt, cat.Filter(filter), parent_id)
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

		_, _, _, catcount, _ := CategoryConfig.ListCategory(offset, limt, cat.Filter{}, parent_id)
		_, _, categorys, _, _ := CategoryConfig.ListCategory(offset, limt, cat.Filter{}, parent_id)
		_, FinalModalCategoriesList, _, _, _ := CategoryConfig.ListCategory(offset, limt, cat.Filter{}, parent_id)

		//pagination calc
		paginationendcount := len(categorylist) + offset
		paginationstartcount := offset + 1
		Previous, Next, PageCount, Page := Pagination(pageno, int(Total_categories), limt)

		translate, _ := TranslateHandler(c)

		menu := NewMenuController(c)

		Folder, File, Media, err := GetMedia()
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
		}, "Menu": menu, "translate": translate, "catcount": catcount, "csrf": csrf.GetToken(c), "HeadTitle": translate.Categories, "Category": categorys, "Previous": Previous, "Next": Next, "PageCount": PageCount, "Page": Page, "CurrentPage": pageno, "categoryid": parent_id, "ParentCategoryName": categorys.CategoryName, "CategoryList": FinalCategoriesList, "Count": Total_categories, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "Filter": filter, "title": ModuleName, "FinalModalCategoriesList": FinalModalCategoriesList, "heading": categorys.CategoryName, "Back": "/categories", "Folder": Folder, "File": File, "Media": Media, "Categoriesmenu": true, "Limit": limt, "Tabmenu": TabName, "ModuleName": ModuleName, "StorageType": selectedtype.SelectedType})

		return

	}

	c.Redirect(301, "/403-page")

}

// Add Children Category
func AddSubCategory(c *gin.Context) {

	//get value from html form data
	id := (c.PostForm("categoryid"))
	imagedata := c.PostForm("categoryimage")
	userid := c.GetInt("userid")
	var ParentId, _ = strconv.Atoi(c.PostForm("subcategoryid"))

	subcategory := cat.CategoryCreate{
		CategoryName: c.PostForm("cname"),
		Description:  c.PostForm("cdesc"),
		ImagePath:    imagedata,
		ParentId:     ParentId,
		CreatedBy:    userid,
	}

	permisison, perr := NewAuth.IsGranted("Categories", auth.Create)
	if perr != nil {
		ErrorLog.Printf("Addsubcategory authorization error: %s", perr)
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
		return

	}

	c.Redirect(301, "/403-page")

}

// Edit Sub Category
func EditSubCategory(c *gin.Context) {

	var id, _ = strconv.Atoi(c.Query("id"))

	permisison, perr := NewAuth.IsGranted("Categories", auth.Update)
	if perr != nil {
		ErrorLog.Printf("editsubcategory authorization error: %s", perr)
	}

	if permisison {
		categorys, err := CategoryConfig.GetSubCategoryDetails(id)
		if err != nil {
			ErrorLog.Printf("editsubcategory error :%s", err)
		}
		c.JSON(200, categorys)
		return
	}

	c.Redirect(301, "/403-page")

}

// Update Childern Category
func UpdateSubCategory(c *gin.Context) {

	//get value from html form data
	ParentId, _ := strconv.Atoi(c.PostForm("pcategoryid"))
	imagedata := c.PostForm("image")
	userid := c.GetInt("userid")
	Categoryid, _ := strconv.Atoi(c.PostForm("id"))

	subcategory := cat.CategoryCreate{
		Id:           Categoryid,
		CategoryName: c.PostForm("cname"),
		Description:  c.PostForm("cdesc"),
		ImagePath:    imagedata,
		ParentId:     ParentId,
		ModifiedBy:   userid,
	}

	permisison, perr := NewAuth.IsGranted("Categories", auth.Update)
	if perr != nil {
		ErrorLog.Printf("updatesubcategory authorization error: %s", perr)
	}

	if permisison {

		err := CategoryConfig.UpdateSubCategory(subcategory) // update category
		if strings.Contains(fmt.Sprint(err), "given some values is empty") {
			ErrorLog.Printf("updatesubcategory mandatory field error: %s", perr)
			c.SetCookie("Alert-msg", "Pleaseenterthemandatoryfields", 3600, "", "", false, false)
			c.Redirect(301, "/categories/")
			return
		}

		if err != nil {
			ErrorLog.Printf("updatesubcategory error: %s", perr)
			c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			c.JSON(200, false)
			return
		}

		c.SetCookie("get-toast", "Category Updated Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.JSON(200, true)
		return
	}

	c.Redirect(301, "/403-page")

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

	permisison, perr := NewAuth.IsGranted("Categories", auth.Delete)
	if perr != nil {
		ErrorLog.Printf("deletesubcategory authorization error: %s", perr)
	}

	if permisison {

		err := CategoryConfig.DeleteSubCategory(categoryid, userid) // delete subcategory
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
		return
	}

	c.Redirect(301, "/403-page")

}

// Children Category Delete Id Passing
func CategoryPopup(c *gin.Context) {

	category = categories.Category{Authority: &AUTH}

	var id, _ = strconv.Atoi(c.Query("id"))

	category, err := category.DeletePopup(id)

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, category)
}

// To Check Group Name id Already Exists
func CheckCategoryGroupName(c *gin.Context) {

	//get value from html form data
	category_id, _ := strconv.Atoi(c.PostForm("category_id"))
	category_name := c.PostForm("category_name")

	permisison, perr := NewAuth.IsGranted("Categories", auth.Read)
	if perr != nil {
		ErrorLog.Printf("checkcategorygroupname authorization error: %s", perr)
	}
	if permisison {

		flg, err := CategoryConfig.CheckCategroyGroupName(category_id, category_name)
		if err != nil {
			ErrorLog.Printf("checkcategorygroupname  error: %s", err)
			json.NewEncoder(c.Writer).Encode(false)
			return
		}

		json.NewEncoder(c.Writer).Encode(flg)
		return

	}

	c.Redirect(301, "/403-page")

}

// SubCategory name check
func CheckSubCategoryName(c *gin.Context) {

	//get value from html form data
	category_id, _ := strconv.Atoi(c.PostForm("parentid"))
	permisison, perr := NewAuth.IsGranted("Categories", auth.Read)
	if perr != nil {
		ErrorLog.Printf("CheckSubCategoryName authorization error: %s", perr)
	}
	if permisison {
		categorylistinchk, _, _, _, err := CategoryConfig.ListCategory(0, Limit, cat.Filter{}, category_id)
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
		flg, err := CategoryConfig.CheckSubCategroyName(catid, categoryid, category_name)

		if err != nil {
			ErrorLog.Printf("CheckSubCategoryName error: %s", err)
			json.NewEncoder(c.Writer).Encode(false)
			return
		}

		json.NewEncoder(c.Writer).Encode(flg)
		return

	}

	c.Redirect(301, "/403-page")

}

func MultiSelectCategoryGroupDelete(c *gin.Context) {

	//get value from html form data
	categoryId := c.PostFormArray("categorygrbids[]")
	pageno := c.PostForm("page")
	userid := c.GetInt("userid")
	var url string
	if pageno != "" {
		url = "/categories?page=" + pageno
	} else {
		url = "/categories/"

	}
	permisison, perr := NewAuth.IsGranted("Categories", auth.Delete)
	if perr != nil {
		ErrorLog.Printf("CheckSubCategoryName authorization error: %s", perr)
	}
	if permisison {
		categoryIntIds := make([]int, len(categoryId))
		for i, id := range categoryId {
			intId, _ := strconv.Atoi(id)
			categoryIntIds[i] = intId
		}

		err := CategoryConfig.MultiSelectDeleteCategoryGroup(categoryIntIds, userid) //delete selected category group
		if err != nil {
			ErrorLog.Printf("MultiSelectCategoryGroupDelete error: %s", err)
			c.JSON(200, gin.H{"value": false})
			return
		}

		c.JSON(200, gin.H{"value": true, "url": url})
		return
	}
	c.Redirect(301, "/403-page")
}

func MultiSelectCategoryDelete(c *gin.Context) {

	//get value from html form data
	categoryId := c.PostFormArray("categoryids[]")
	parent_id := (c.PostForm("parentid"))
	pageno := c.PostForm("page")
	userid := c.GetInt("userid")
	var url string
	if pageno != "" {
		url = "/categories/addcategory/" + parent_id + "?page=" + pageno
	} else {
		url = "/categories/addcategory/" + parent_id

	}

	permisison, perr := NewAuth.IsGranted("Categories", auth.Delete)
	if perr != nil {
		ErrorLog.Printf("CheckSubCategoryName authorization error: %s", perr)
	}
	if permisison {
		categoryIntIds := make([]int, len(categoryId))
		for i, id := range categoryId {
			intId, _ := strconv.Atoi(id)
			categoryIntIds[i] = intId
		}

		err := CategoryConfig.MultiselectSubCategoryDelete(categoryIntIds, userid) //delete selected categories
		if err != nil {
			ErrorLog.Printf("MultiSelectCategoriesDelete error: %s", err)
			c.JSON(200, gin.H{"value": false})
			return
		}

		c.JSON(200, gin.H{"value": true, "url": url})
		return
	}
	c.Redirect(301, "/403-page")
}
