package config

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
	BotToken string
}

const (
	token = "BotToken"
)

func New() (*Config, error) {
	if v := validateEnvironment(); v != nil {
		return nil, v
	}

	return &Config{
		BotToken: os.Getenv(token),
	}, nil
}

func validateEnvironment() error {
	requiredEnvVars := []string{token}

	missingEnvVars := ""

	for _, v := range requiredEnvVars {
		if os.Getenv(v) == "" {
			missingEnvVars += v + ","
		}
	}

	if missingEnvVars != "" {
		return errors.New(fmt.Sprintf("Missing env vars: %s", missingEnvVars))
	}

	return nil
}
