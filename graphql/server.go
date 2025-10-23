package graphql

import (
	"log"
	"os"
	"spurt-cms/graphql/routes"
	"sync"

	"github.com/gin-gonic/gin"
)

func RunGraphqlServer(wg *sync.WaitGroup) {

	defer wg.Done()

	port := os.Getenv("GRAPHQL_PORT")

	r := gin.Default()

	REngine := routes.GetEndpointHandlers(r)

	if err := REngine.Run(":" + port); err != nil{

		log.Fatal("")
	}

}
