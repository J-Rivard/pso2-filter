package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/J-Rivard/pso2-filter/internal/clients/bot"
)

type Config struct {
	BotParams *bot.Parameters
}

const (
	token = "BotToken"
)

func New() (*Config, error) {
	if v := validateEnvironment(); v != nil {
		return nil, v
	}

	return &Config{
		BotParams: &bot.Parameters{
			Token: os.Getenv(token),
		},
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
