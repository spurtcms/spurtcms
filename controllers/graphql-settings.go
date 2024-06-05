package controllers

import (
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func GraphqlTokenApi(c *gin.Context) {

	menu := NewMenuController(c)

	translate, _ := TranslateHandler(c)

	c.HTML(200, "graphql.html", gin.H{"csrf": csrf.GetToken(c), "HeadTitle": "Graphql Settings", "Menu": menu, "translate": translate, "SettingsHead": true, "Graphqlmenu": true, "title": "Graphql Api"})

}

func CreateTokenGrapqhl(c *gin.Context) {

	menu := NewMenuController(c)

	translate, _ := TranslateHandler(c)

	c.HTML(200, "graphql-manage.html", gin.H{"csrf": csrf.GetToken(c), "HeadTitle": "Graphql Settings", "Menu": menu, "translate": translate, "SettingsHead": true, "Graphqlmenu": true, "title": "Graphql Api"})

}
