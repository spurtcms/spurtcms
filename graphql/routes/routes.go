package routes

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"spurt-cms/graphql/controller"
	"spurt-cms/graphql/graph"
	"spurt-cms/graphql/middleware"
	"spurt-cms/graphql/resolvers"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func GetEndpointHandlers(r *gin.Engine) *gin.Engine {

	r.Static("/public", "./graphql/public")

	var htmlFilePath []string

	filepath.Walk("./graphql", func(path string, _ fs.FileInfo, _ error) error {

		if strings.HasSuffix(path, ".html") {

			htmlFilePath = append(htmlFilePath, path)
		}

		return nil
	})

	r.LoadHTMLFiles(htmlFilePath...)

	// r.Use(middleware.LimitRequestBodySize(20 << 20))

	r.Use(middleware.CorsMiddleware())

	r.GET("/image-resize", controller.ImageResize)

	r.GET("/", func(c *gin.Context) {

		h := playground.Handler("GraphQL playground", os.Getenv("GRAPHQL_API_URL"))

		h.ServeHTTP(c.Writer, c.Request)

	})

	r.POST("/query", func(c *gin.Context) {

		execSchema := graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{}, Directives: graph.DirectiveRoot{Auth: middleware.AuthMiddleware}})

		srv := handler.NewDefaultServer(execSchema)

		ctx := context.WithValue(c.Request.Context(), controller.GinContext, c)

		srv.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	})

	r.GET("/apidocs", func(c *gin.Context) {

		c.HTML(200, "spectaqldocs.html", nil)
	})

	return r
}
