package pageview

import (
	"os"
	viewcontroller "spurt-cms/page-view/controller"
	"spurt-cms/page-view/middleware"
	"sync"

	"github.com/gin-gonic/gin"
)

func RunTemplateView(wg *sync.WaitGroup) {

	defer wg.Done()

	r := gin.Default()

	r.Use(middleware.TemplateViewAuth())

	r.Static("/public", "./public")

	r.LoadHTMLGlob("view/**/*.html")

	r.GET("/:id", viewcontroller.PageView)

	r.GET("/404-pageview", viewcontroller.FileNotFound)

    r.NoRoute(func(ctx *gin.Context) {

		ctx.Redirect(301, os.Getenv("VIEW_BASE_URL")+"/404-pageview")
	})
	err := r.Run(":" + os.Getenv("VIEW_PORT"))



	if err != nil {

		panic(err)
	}

}
