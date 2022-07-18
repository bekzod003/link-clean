package postgres

import (
	"context"
	"time"

	"github.com/bekzod003/link-clean/internal/domain/entities"
	"github.com/jackc/pgx/v4/pgxpool"
)

type tagStorage struct {
	db *pgxpool.Pool
}

func NewTagStorage(db *pgxpool.Pool) *tagStorage {
	return &tagStorage{db: db}
}

func (t *tagStorage) Create(tag *entities.CreateTag) (id int64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = t.db.QueryRow(
		ctx,
		"INSERT INTO tags (title, user_id) VALUES ($1, $2) RETURNING id",
		tag.Title,
		tag.UserID,
	).Scan(&id)
	return id, err
}

func (t *tagStorage) Get(id int64) (*entities.Tag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var tag entities.Tag
	row := t.db.QueryRow(
		ctx,
		`SELECT
			id,
			title,
			user_id,
			created_at,
			updated_at,
		FROM tags WHERE id = $1 AND deleted_at = 0`,
		id,
	)
	if err := row.Scan(
		&tag.Id,
		&tag.Title,
		&tag.UserID,
		&tag.CreatedAt,
		&tag.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &tag, nil
}

func (t *tagStorage) GetByUser(userID int64) ([]*entities.Tag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var tags []*entities.Tag
	rows, err := t.db.Query(
		ctx,
		`SELECT
			id,
			title,
			user_id,
			created_at,
			updated_at,
		FROM tags WHERE user_id = $1 AND deleted_at = 0`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tag entities.Tag
		if err := rows.Scan(
			&tag.Id,
			&tag.Title,
			&tag.UserID,
			&tag.CreatedAt,
			&tag.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tags = append(tags, &tag)
	}

	return tags, nil
}

func (t *tagStorage) Update(tag *entities.UpdateTag) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if _, err := t.db.Exec(
		ctx,
		`UPDATE tags SET title = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2`,
		tag.Title,
		tag.ID,
	); err != nil {
		return err
	}

	return nil
}

func (t *tagStorage) Delete(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if _, err := t.db.Exec(
		ctx,
		`UPDATE tags SET deleted_at = $1 WHERE id = $2`,
		time.Now().Unix(),
		id,
	); err != nil {
		return err
	}

	return nil
}
