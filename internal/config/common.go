package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/J-Rivard/pso2-filter/internal/clients/db"

	"github.com/J-Rivard/pso2-filter/internal/clients/bot"
)

type Config struct {
	BotParams *bot.Parameters
	DBParams  *db.Parameters
}

const (
	token     = "BotToken"
	db_pw     = "db_pw"
	db_user   = "db_user"
	db_host   = "db_host"
	db_schema = "db_schema"
	db_name   = "db_name"
)

func New() (*Config, error) {
	if v := validateEnvironment(); v != nil {
		return nil, v
	}

	return &Config{
		BotParams: &bot.Parameters{
			Token: os.Getenv(token),
		},
		DBParams: &db.Parameters{
			Host:     os.Getenv(db_host),
			Username: os.Getenv(db_user),
			Password: os.Getenv(db_pw),
			Schema:   os.Getenv(db_schema),
			DBName:   os.Getenv(db_name),
		},
	}, nil
}

func validateEnvironment() error {
	requiredEnvVars := []string{token, db_user, db_pw, db_host, db_schema, db_name}

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
