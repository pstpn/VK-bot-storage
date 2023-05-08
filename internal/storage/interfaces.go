package storage

import (
	"context"

	"github.com/telegram-bot/internal/models"
)

type UsersDataRepo interface {
	Set(ctx context.Context, input *models.SetUsersDataInput) error
	Get(ctx context.Context, input *models.GetUsersDataInput) ([]*models.UsersData, error)
	Del(ctx context.Context, input *models.DelUsersDataInput) error
}

type ServicesRepo interface {
	Get(ctx context.Context, input *models.GetServiceInput) (*models.Service, error)
}
