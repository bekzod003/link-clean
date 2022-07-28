package postgres

import (
	"context"
	"time"

	"github.com/bekzod003/link-clean/pkg/database/client/postgresql"

	"github.com/bekzod003/link-clean/internal/domain/entities"
)

type linkStorage struct {
	db postgresql.Client
}

func NewLinkStorage(db postgresql.Client) *linkStorage {
	return &linkStorage{db: db}
}

func (l *linkStorage) Create(ctx context.Context, link *entities.Link) error {
	_, err := l.db.Exec(
		ctx,
		`INSERT INTO "links"
			(title, url, user_id, tag_id)
		VALUES
			($1, $2, $3, $4)`,
		link.Title,
		link.URL,
		link.UserID,
		link.TagID,
	)

	return err
}

func (l *linkStorage) Get(ctx context.Context, id int64) (*entities.Link, error) {
	var link entities.Link
	err := l.db.QueryRow(
		ctx,
		`SELECT
			id,
			url,
			title,
			description,
			user_id,
			tag_id,
			created_at,
			updated_at
		FROM links WHERE id = $1 AND deleted_at = 0`,
		id,
	).Scan(
		&link.ID,
		&link.URL,
		&link.Title,
		&link.Description,
		&link.UserID,
		&link.TagID,
		&link.CreatedAt,
		&link.UpdatedAt,
	)

	return &link, err
}

func (l *linkStorage) GetByUser(ctx context.Context, userID int64) ([]*entities.Link, error) {
	var links []*entities.Link
	rows, err := l.db.Query(
		ctx,
		`SELECT
			id,
			url,
			title,
			description,
			user_id,
			tag_id,
			created_at,
			updated_at
		FROM links WHERE user_id = $1 AND deleted_at = 0`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var link entities.Link
		if err := rows.Scan(
			&link.ID,
			&link.URL,
			&link.Title,
			&link.Description,
			&link.UserID,
			&link.TagID,
			&link.CreatedAt,
			&link.UpdatedAt,
		); err != nil {
			return nil, err
		}
		links = append(links, &link)
	}

	return links, nil
}

func (l *linkStorage) GetByTag(ctx context.Context, tagID int64) ([]*entities.Link, error) {
	var links []*entities.Link
	rows, err := l.db.Query(
		ctx,
		`SELECT
			id,
			url,
			title,
			description,
			user_id,
			tag_id,
			created_at,
			updated_at
		FROM links WHERE tag_id = $1 AND deleted_at = 0`,
		tagID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var link entities.Link
		if err := rows.Scan(
			&link.ID,
			&link.URL,
			&link.Title,
			&link.Description,
			&link.UserID,
			&link.TagID,
			&link.CreatedAt,
			&link.UpdatedAt,
		); err != nil {
			return nil, err
		}
		links = append(links, &link)
	}

	return links, nil
}

func (l *linkStorage) Update(ctx context.Context, link *entities.Link) error {
	_, err := l.db.Exec(
		ctx,
		`UPDATE "links" SET
			title = $1,
			url = $2,
			user_id = $3,
			tag_id = $4,
			updated_at = $5
		WHERE id = $6`,
		link.Title,
		link.URL,
		link.UserID,
		link.TagID,
		time.Now(),
		link.ID,
	)
	return err
}

func (l *linkStorage) Delete(ctx context.Context, id int64) error {
	_, err := l.db.Exec(
		ctx,
		`UPDATE "links" SET deleted_at = $1 WHERE id = $2`,
		time.Now().Unix(),
		id,
	)
	return err
}
