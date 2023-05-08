package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/telegram-bot/internal/models"
	"github.com/telegram-bot/internal/storage"
)

type usersDataRepo struct {
	db *pgx.Conn
}

func NewUsersDataRepo(db *pgx.Conn) storage.UsersDataRepo {
	return &usersDataRepo{db: db}
}

func (r *usersDataRepo) Set(ctx context.Context, input *models.SetUsersDataInput) error {

	sql := `
INSERT INTO users_data (service_id, user_id, login, password, time) 
VALUES ($1, $2, $3, $4, $5);`

	_, err := r.db.Exec(ctx, sql, input.ServiceID, input.UserID, input.Login, input.Password, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (r *usersDataRepo) Get(ctx context.Context, input *models.GetUsersDataInput) ([]*models.UsersData, error) {

	sql := `
SELECT * FROM users_data
WHERE service_id = $1 AND user_id = $2;`

	rows, err := r.db.Query(ctx, sql, input.ServiceID, input.UserID)
	if err != nil {
		return nil, err
	}

	usersData := make([]*models.UsersData, 1)
	for rows.Next() {
		userData := &models.UsersData{}
		err = rows.Scan(&userData.ServiceID, &userData.ID, &userData.Login, &userData.Password, &userData.Time)
		if err != nil {
			return nil, err
		}
		usersData = append(usersData, userData)
	}

	return usersData, nil
}

func (r *usersDataRepo) Del(ctx context.Context, input *models.DelUsersDataInput) error {

	sql := `
DELETE FROM users_data
WHERE service_id = $1 AND user_id = $2 AND login = $3;`

	_, err := r.db.Exec(ctx, sql, input.ServiceID, input.UserID, input.Login)
	if err != nil {
		return err
	}

	return nil
}
