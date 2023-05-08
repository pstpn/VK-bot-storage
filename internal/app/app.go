package app

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/telegram-bot/config"
	"github.com/telegram-bot/internal/bot"
	"github.com/telegram-bot/internal/storage"
	"github.com/telegram-bot/internal/storage/postgres"
)

func Run(cfg *config.Config) {

	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - pgx.Connect: %w", err))
	}
	defer func() {
		err = db.Close(context.Background())
		if err != nil {
			log.Fatal(fmt.Errorf("app - Run - db.Close: %w", err))
		}
	}()

	repo := storage.NewRepo(
		postgres.NewUsersDataRepo(db),
		postgres.NewServicesRepo(db),
	)

	err = bot.RunBot(cfg.Bot.Token, repo)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - bot.RunBot: %w", err))
	}
}
