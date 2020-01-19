package main

import (
	"registry-auth-server/cmd"
	"registry-auth-server/config"
)

func main() {
	config.Init()
	cmd.Execute()
}
