package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"registry-auth-server/config"
	"registry-auth-server/server"
)

// Command definition
var runServerCmd = &cobra.Command{
	Use:   "runserver",
	Short: "Start the server",
	Run:   RunServerCommand,
}

// Flags
var (
	UserAPI bool
	// RegistrySSLKeyPath  string
	// RegistrySSLCertPath string
)

// RunServerCommand func
func RunServerCommand(cmd *cobra.Command, args []string) {
	checkFile(config.Global.RegistryCertPath)
	checkFile(config.Global.RegistryKeyPath)

	server.RunServer()
}

func checkFile(path string) {
	filePath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	ok, err := isDir(filePath)
	if err != nil {
		// Permission
		log.Fatal(err)
	}
	if ok {
		log.Fatal("Please specify only a file. Not a directory.")
	}
}

func isDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}
