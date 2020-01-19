package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "registry-auth-server",
	Short:   "Authentication & Authorization server for Private Docker Registries.",
	Version: "1.0.0",
}

// Execute root command as help command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(runServerCmd)

	// Add flag to startCmd
	// runServerCmd.Flags().StringVarP(&RegistrySSLKeyPath, "registry-key", "", "", "Registry SSL key path")
	// runServerCmd.Flags().StringVarP(&RegistrySSLCertPath, "registry-crt", "", "", "Registry SSL certificate path")

	// Required Flags
	// runServerCmd.MarkFlagRequired("registry-key")
	// runServerCmd.MarkFlagRequired("registry-crt")
}
