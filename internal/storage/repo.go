package storage

import (
	"context"

	"github.com/telegram-bot/internal/models"
)

type Repo struct {
	ud UsersDataRepo
	sv ServicesRepo
}

func NewRepo(ud UsersDataRepo, sv ServicesRepo) *Repo {
	return &Repo{
		ud: ud,
		sv: sv,
	}
}

func (r *Repo) GetService(ctx context.Context, serviceName string) (*models.Service, error) {
	return r.sv.Get(ctx, &models.GetServiceInput{ServiceName: serviceName})
}
