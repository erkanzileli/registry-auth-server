package server

import (
	"registry-auth/common"
	"registry-auth/db/model"
	"registry-auth/server"

	"github.com/gin-gonic/gin"
)

func registerActionRoutes() {
	g := server.Router.Group("action")
	g.GET("read", read)
}

func read(c *gin.Context) {
	var actions = model.Action{}
	common.DB.Find(&actions)
}
