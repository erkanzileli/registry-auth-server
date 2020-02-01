package server

import (
	"fmt"
	"log"
	"registry-auth/config"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Endpoint string
	Method   string
	Handler  func(c *gin.Context)
}

var router *gin.Engine

func registerRoutes() {
	registerAuthRoutes(router)
}

// RunServer creates a gin instance and starts to listen http calls
func RunServer() {
	router = gin.Default()

	registerRoutes()

	err := router.Run(fmt.Sprintf("%s:%s", config.Global.Host, config.Global.Port))
	if err != nil {
		log.Fatal(err)
	}
}
