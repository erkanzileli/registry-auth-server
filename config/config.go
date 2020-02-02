package config

import (
	"os"
)

type config struct {
	// HTTP server
	Host string
	Port string

	// Registry TLS
	RegistryKeyPath  string
	RegistryCertPath string

	TokenIssuer string

	// Audience
	TokenService string
}

var envPrefix = "REGISTRY_AUTH_"

var Global = &config{}

func Init() {
	setDefaults()
	readFromEnv()
}

func readFromEnv() {
	// todo: env parser
	Global.RegistryCertPath = os.Getenv(envPrefix + "CERT_PATH")
	Global.RegistryKeyPath = os.Getenv(envPrefix + "KEY_PATH")
	Global.TokenIssuer = os.Getenv(envPrefix + "TOKEN_ISSUER")
	Global.TokenService = os.Getenv(envPrefix + "TOKEN_SERVICE")

	host := os.Getenv(envPrefix + "HOST")
	if len(host) > 0 {
		Global.Host = host
	}

	port := os.Getenv(envPrefix + "PORT")
	if len(port) > 0 {
		Global.Port = port
	}
}

func setDefaults() {
	Global.Port = "8000"
	Global.Host = "0.0.0.0"
}
