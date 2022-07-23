package postgres

import (
	"context"
	"github.com/bekzod003/link-clean/internal/domain/entities"
	"github.com/bekzod003/link-clean/pkg/database/client/postgresql"
)

type userStorage struct {
	db postgresql.Client
}

func NewUserStorage(db postgresql.Client) *userStorage {
	return &userStorage{db: db}
}

func (u *userStorage) Create(ctx context.Context, user *entities.User) (int64, error) {
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

func (u *userStorage) Update(ctx context.Context, user *entities.User) error {
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

func (u *userStorage) Get(ctx context.Context, id int64) (*entities.User, error) {
	var user entities.User
	row := u.db.QueryRow(
		ctx,
		`SELECT
			id,
			username,
			first_name,
			last_name,
			created_at
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
