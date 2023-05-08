package main

import (
	"log"

	"github.com/telegram-bot/config"
	"github.com/telegram-bot/internal/app"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
