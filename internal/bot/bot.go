package bot

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/telegram-bot/internal/storage"
)

type Handler struct {
	bot  *tgbotapi.BotAPI
	repo *storage.Repo
}

func NewHandler(token string, r *storage.Repo) (*Handler, error) {

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	bot.Buffer = 0

	return &Handler{
		bot:  bot,
		repo: r,
	}, nil
}

func (h *Handler) Handle() error {

	updates, _ := h.bot.GetUpdatesChan(tgbotapi.UpdateConfig{
		Offset:  0,
		Timeout: 30,
	})

	for update := range updates {
		if update.Message == nil {
			continue
		}
		var answer string
		currentMsg := update.Message

		switch currentMsg.Command() {
		case "start":
			answer = h.startRoute()
		case "set":
			answer = h.setRoute(currentMsg)
		case "get":
			answer = h.getRoute(currentMsg)
		case "del":
			answer = h.delRoute(currentMsg)
		}
		replyMsg, err := h.bot.Send(tgbotapi.NewMessage(currentMsg.Chat.ID, answer))
		if err != nil {
			return err
		}

		if strings.HasPrefix(answer, GetSuccessMsg) {
			go h.deleteMessageAfterDeadline(&replyMsg)
		}
	}

	return nil
}
