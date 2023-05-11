package configs

import (
	"log"
	"mgmt/internal/env"
)

type Config struct {
	TemplateDir string
}

func returnString(envString string, envErr error) string {
	if envErr != nil {
		log.Fatalf("func FetchVars failed with error: %v", envErr)
	}

	return envString
}

func FetchVars() *Config {
	return &Config{
		TemplateDir: returnString((env.GetDefault("TEMPLATE_DIR", "/var/www/html/"))),
	}
}
