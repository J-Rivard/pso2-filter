package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/J-Rivard/pso2-filter/internal/clients/db"

	"github.com/J-Rivard/pso2-filter/internal/clients/bot"
	"github.com/J-Rivard/pso2-filter/internal/config"
	"github.com/J-Rivard/pso2-filter/internal/logging"
	"github.com/rs/zerolog"
)

func main() {

	logger, err := logging.New(zerolog.ConsoleWriter{Out: os.Stderr}, logging.Debug)
	if err != nil {
		panic(err)
	}

	cfg, err := config.New()
	if err != nil {
		logger.LogFatal(logging.FormattedLog{
			"action": "startup",
			"error":  err.Error(),
		})
	}

	database, err := db.New(cfg.DBParams)
	if err != nil {
		logger.LogFatal(logging.FormattedLog{
			"action": "startup",
			"error":  err.Error(),
		})
	}

	bot, err := bot.New(cfg.BotParams, database, logger)
	if err != nil {
		logger.LogFatal(logging.FormattedLog{
			"action": "startup",
			"error":  err.Error(),
		})
	}

	err = bot.Start()
	if err != nil {
		logger.LogFatal(logging.FormattedLog{
			"action": "startup",
			"error":  err.Error(),
		})
	}

	logger.LogApplication(logging.FormattedLog{
		"action":   "startup",
		"metadata": "pso2-filter now running",
	})

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	err = bot.Stop()
	if err != nil {
		logger.LogFatal(logging.FormattedLog{
			"action": "shutdown",
			"error":  err.Error(),
		})
	}
}
