package controllers

import "github.com/gin-gonic/gin"

func CategoriesSettings(c *gin.Context) {

	menu := NewMenuController(c)

	ModuleName, TabName,_ := ModuleRouteName(c)

	translate, _ := TranslateHandler(c)

	c.HTML(200, "category_settings.html", gin.H{"Menu": menu, "HeadTitle": translate.Categories, "Categoriesmenu": true, "title": ModuleName, "Tabmenu": TabName})

}
