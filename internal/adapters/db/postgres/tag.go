package postgres

import (
	"context"
	"github.com/bekzod003/link-clean/pkg/database/client/postgresql"
	"time"

	"github.com/bekzod003/link-clean/internal/domain/entities"
)

type tagStorage struct {
	db postgresql.Client
}

func NewTagStorage(db postgresql.Client) *tagStorage {
	return &tagStorage{db: db}
}

func (t *tagStorage) Create(ctx context.Context, tag *entities.CreateTag) (id int64, err error) {
	err = t.db.QueryRow(
		ctx,
		"INSERT INTO tags (title, user_id) VALUES ($1, $2) RETURNING id",
		tag.Title,
		tag.UserID,
	).Scan(&id)
	return id, err
}

func (t *tagStorage) Get(ctx context.Context, id int64) (*entities.Tag, error) {
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

func (t *tagStorage) GetByUser(ctx context.Context, userID int64) ([]*entities.Tag, error) {
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

func (t *tagStorage) Update(ctx context.Context, tag *entities.UpdateTag) error {
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

func (t *tagStorage) Delete(ctx context.Context, id int64) error {
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
