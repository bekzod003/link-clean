package postgres

import (
	"context"
	"github.com/bekzod003/link-clean/pkg/database/client/postgresql"
	"time"

	"github.com/bekzod003/link-clean/internal/domain/entities"
)

type userStorage struct {
	db postgresql.Client
}

func NewUserStorage(db postgresql.Client) *userStorage {
	return &userStorage{db: db}
}

func (u *userStorage) Create(user *entities.User) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := u.db.Exec(
		ctx,
		`INSERT INTO "users"
			(id, username, first_name, last_name)
		VALUES
			($1, $2, $3, $4)`,
		user.ID,
		user.Username,
		user.FirstName,
		user.LastName,
	)
	return user.ID, err
}

func (u *userStorage) Update(user *entities.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := u.db.Exec(
		ctx,
		`UPDATE "users"
		SET
			username = $2,
			first_name = $3,
			last_name = $4
		WHERE id = $1`,
		user.ID,
		user.Username,
		user.FirstName,
		user.LastName,
	)
	return err
}

func (u *userStorage) Get(id int64) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var user entities.User
	row := u.db.QueryRow(
		ctx,
		`SELECT
			id,
			username,
			first_name,
			last_name,
			created_at,
			updated_at,
			deleted_at
		FROM users WHERE id = $1`,
		id,
	)
	if err := row.Scan(
		&user.ID,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &user, nil
}
