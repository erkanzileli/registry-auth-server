package main

import (
	"registry-auth/cmd"
	"registry-auth/config"
	"registry-auth/db"
)

func main() {
	config.Init()
	db.Init()
	cmd.Execute()
}
