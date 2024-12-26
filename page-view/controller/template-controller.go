package controller

import (
	"html/template"
	"spurt-cms/controllers"
	"strings"

	"github.com/gin-gonic/gin"
)

func PageView(c *gin.Context) {

	uuid := c.Param("id")
	arr := strings.Split(uuid, "-")
	entries, _, err := controllers.ChannelConfigWP.EntryPreview(arr[len(arr)-1])
	if err != nil {
		controllers.ErrorLog.Printf("preview  details api error: %s", err)
	}
	content := template.HTML(entries.Description)
	c.HTML(200, "page-view.html", gin.H{"Entries": content, "Entry": entries})
}

func FileNotFound(c *gin.Context) {

	c.HTML(200, "404pageview.html", gin.H{"page not found": true})
}
