package bot

import (
	"context"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/telegram-bot/internal/storage"
)

func RunBot(token string, r *storage.Repo) error {

	b, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}
	b.Debug = true
	b.Buffer = 0
	log.Printf("The bot with name \"%s\" has been successfully launched\n", b.Self.UserName)

	err = handleBot(b, r)
	if err != nil {
		return err
	}

	return nil
}

func handleBot(b *tgbotapi.BotAPI, r *storage.Repo) error {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates, _ := b.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Text {
		case "/vadno":
			go func() {
				ok, _ := b.GetUpdatesChan(tgbotapi.UpdateConfig{
					Offset:  0,
					Limit:   0,
					Timeout: 20,
				})
				b.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Enter service"))
				r.GetService(context.Background(), (<-ok).Message.Text)
				//go deleteMessageAfterDeadline(b, update.Message)
			}()
		default:
			b.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Incorrect action! Try again!"))
		}

	}

	return nil
}

func deleteMessageAfterDeadline(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	time.Sleep(10 * time.Second)
	bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
		ChatID:    msg.Chat.ID,
		MessageID: msg.MessageID,
	})
}
