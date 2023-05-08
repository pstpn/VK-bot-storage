package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/telegram-bot/internal/models"
	"github.com/telegram-bot/internal/storage"
)

type servicesRepo struct {
	db *pgx.Conn
}

func NewServicesRepo(db *pgx.Conn) storage.ServicesRepo {
	return &servicesRepo{db: db}
}

func (s *servicesRepo) Get(ctx context.Context, input *models.GetServiceInput) (*models.Service, error) {

	sql := `
SELECT * FROM services
WHERE name = $1;`

	rows, err := s.db.Query(ctx, sql, input.ServiceName)
	if err != nil {
		return nil, err
	}

	var service *models.Service
	err = rows.Scan(&service.ID, &service.Name)
	if err != nil {
		return nil, err
	}

	return service, nil
}
