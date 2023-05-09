package storage

import (
	"context"

	"github.com/telegram-bot/internal/models"
)

type UsersDataRepo interface {
	Set(ctx context.Context, input *models.SetUserDataInput) error
	Get(ctx context.Context, input *models.GetUserDataInput) (*models.UserData, error)
	Del(ctx context.Context, input *models.DelUserDataInput) error
}
