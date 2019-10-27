package server

import (
	"registry-auth-server/auth"

	"github.com/gin-gonic/gin"
)

// RunServer creates a gin instance and starts to listen http calls
func RunServer() {

	r := gin.Default()
	r.GET("/auth", func(c *gin.Context) {
		service := c.Query("service")

		if service != service {
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
					"token": auth.CreateToken(u),
				}
				c.JSON(200, response)
				return
			}
		}

		c.JSON(401, gin.H{})
	})

	r.Run(":8000")

}
