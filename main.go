package main

import (
	"registry-auth-server/cmd"
	"registry-auth-server/config"
	"registry-auth-server/db"
)

func main() {
	config.Init()
	db.Init()
	cmd.Execute()
}
