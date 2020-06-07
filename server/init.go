package server

import (
	"fmt"
	"log"
	"registry-auth/config"

	"github.com/gin-gonic/gin"
)

// Router variable represents entire HTTP router
var Router *gin.Engine

func registerRoutes() {
	// registerAccessTypeRoutes()
	// registerAccessRoutes()
	registerAuthRoutes()
	registerRepositoryRoutes()
	// registerTagRoutes()
}

// RunServer creates a gin instance and starts to listen http calls
func RunServer() {
	Router = gin.Default()

	registerRoutes()

	err := Router.Run(fmt.Sprintf("%s:%s", config.Global.Host, config.Global.Port))
	if err != nil {
		log.Fatal(err)
	}
}
