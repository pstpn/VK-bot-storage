package postgres

import (
	"context"

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

func (r *usersDataRepo) Set(ctx context.Context, input *models.SetUserDataInput) error {

	sql := `
INSERT INTO users_data (chat_id, service_name, login, password) 
VALUES ($1, $2, $3, $4);`

	_, err := r.db.Exec(ctx, sql, input.ChatID, input.ServiceName, input.Login, input.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *usersDataRepo) Get(ctx context.Context, input *models.GetUserDataInput) (*models.UserData, error) {

	sql := `
SELECT chat_id, service_name, login, password FROM users_data
WHERE chat_id = $1 AND service_name = $2;`

	userData := &models.UserData{}
	err := r.db.QueryRow(ctx, sql, input.ChatID, input.ServiceName).
		Scan(&userData.ChatID, &userData.ServiceName, &userData.Login, &userData.Password)
	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (r *usersDataRepo) Del(ctx context.Context, input *models.DelUserDataInput) error {

	sql := `
DELETE FROM users_data
WHERE chat_id = $1 AND service_name = $2;`

	_, err := r.db.Exec(ctx, sql, input.ChatID, input.ServiceName)
	if err != nil {
		return err
	}

	return nil
}
