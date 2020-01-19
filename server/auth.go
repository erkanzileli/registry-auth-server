package server

import (
	"registry-auth-server/auth"
	"registry-auth-server/config"

	"github.com/gin-gonic/gin"
)

func registerAuthRoutes(r *gin.Engine) {
	r.GET("/auth", func(c *gin.Context) {
		service := c.Query("service")

		if service != config.Global.TokenService {
			c.JSON(401, gin.H{})
			return
		}

		account := c.Query("account")
		authHeader, ok := c.Request.Header["Authorization"]
		// Login, authenticate
		if len(account) > 0 && ok {
			u, err := auth.ParseHeader(authHeader[0])
			if err != nil {
				c.JSON(400, "Auth header can not parsed!")
				return
			}
			if u.Authenticate() {
				response := map[string]string{
					"token": auth.CreateToken(u, config.Global.RegistryCertPath, config.Global.RegistryKeyPath),
				}
				c.JSON(200, response)
				return
			}
		}

		c.JSON(401, gin.H{})
	})
}
