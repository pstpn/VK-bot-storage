package bot

import (
	"context"
	"encoding/base64"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/telegram-bot/config"
	"github.com/telegram-bot/internal/models"
)

func (h *Handler) startRoute() string {
	return StartMsg
}

func (h *Handler) setRoute(msg *tgbotapi.Message) string {

	fields := strings.Split(msg.CommandArguments(), " ")
	if len(fields) != 3 {
		return ParamsCountErrorMsg + "\n" + SetCommand
	}

	hash := base64.StdEncoding.EncodeToString([]byte(fields[2]))
	err := h.repo.U.Set(context.Background(), &models.SetUserDataInput{
		ServiceName: fields[0],
		ChatID:      models.ChatID(msg.Chat.ID),
		Login:       models.Login(fields[1]),
		Password:    models.Password(hash),
	})
	if err != nil {
		return SetDataErrorMsg + " \"" + fields[0] + "\""
	}
	go h.deleteMessageAfterDeadline(msg)

	return SetSuccessMsg + " \"" + fields[0] + "\""
}

func (h *Handler) getRoute(msg *tgbotapi.Message) string {

	fields := strings.Split(msg.CommandArguments(), " ")
	if len(fields) != 1 {
		return ParamsCountErrorMsg + "\n" + GetCommand
	}

	userData, err := h.repo.U.Get(context.Background(), &models.GetUserDataInput{
		ServiceName: fields[0],
		ChatID:      models.ChatID(msg.Chat.ID),
	})
	if err != nil {
		return GetDataErrorMsg + " \"" + fields[0] + "\""
	}

	password, err := base64.StdEncoding.DecodeString(string(userData.Password))
	if err != nil {
		return GetPasswordDecodeErrorMsg
	}

	return GetSuccessMsg + " \"" + fields[0] + "\":\n\n" +
		"Login: " + string(userData.Login) + "\n" +
		"Password: " + string(password)
}

func (h *Handler) delRoute(msg *tgbotapi.Message) string {

	fields := strings.Split(msg.CommandArguments(), " ")
	if len(fields) != 1 {
		return ParamsCountErrorMsg + "\n" + DelCommand
	}

	err := h.repo.U.Del(context.Background(), &models.DelUserDataInput{
		ServiceName: fields[0],
		ChatID:      models.ChatID(msg.Chat.ID),
	})
	if err != nil {
		return DelDataErrorMsg + " \"" + fields[0] + "\""
	}

	return DelSuccessMsg + " \"" + fields[0] + "\""
}

func (h *Handler) deleteMessageAfterDeadline(msg *tgbotapi.Message) {

	time.Sleep(config.DeletePasswordDeadline)
	h.bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
		ChatID:    msg.Chat.ID,
		MessageID: msg.MessageID,
	})
}
