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

var envPrefix = "AUTH_"

var Global = &config{}

func Init() {
	setDefaults()
	readFromEnv()
}

func readFromEnv() {
	// todo: env parser
	Global.RegistryCertPath = os.Getenv(envPrefix + "REGISTRY_CERT_PATH")
	Global.RegistryKeyPath = os.Getenv(envPrefix + "REGISTRY_KEY_PATH")
	Global.TokenIssuer = os.Getenv(envPrefix + "REGISTRY_TOKEN_ISSUER")
	Global.TokenService = os.Getenv(envPrefix + "REGISTRY_TOKEN_SERVICE")
	Global.Host = os.Getenv(envPrefix + "HOST")
	Global.Port = os.Getenv(envPrefix + "PORT")
}

func setDefaults() {
	Global.Port = "8000"
	Global.Host = "0.0.0.0"
}
