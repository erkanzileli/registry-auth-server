package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"registry-auth/auth"
	"registry-auth/config"

	"github.com/gin-gonic/gin"
)

func registerRepositoryRoutes() {
	g := router.Group("repository")
	g.GET("read", read)
}

func read(c *gin.Context) {
	repos := getRepositories()
	fmt.Println("repos", repos)
	c.JSON(200, repos)
}

func getRepositories() map[string]interface{} {
	u := &auth.User{Username: "admin", Password: "123qweasd"}
	token := auth.CreateToken(u, config.Global.RegistryCertPath, config.Global.RegistryKeyPath)

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:5000/v2/_catalog", nil)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	resBody := make(map[string]interface{})

	err = json.Unmarshal(body, &resBody)
	if err != nil {
		panic(err)
	}

	return resBody
}
