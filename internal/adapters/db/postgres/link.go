package postgres

import (
	"context"
	"time"

	"github.com/bekzod003/link-clean/internal/domain/entities"
	"github.com/jackc/pgx/v4/pgxpool"
)

type linkStorage struct {
	db *pgxpool.Pool
}

func NewLinkStorage(db *pgxpool.Pool) *linkStorage {
	return &linkStorage{db: db}
}

func (l *linkStorage) Create(link *entities.Link) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

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

func (l *linkStorage) Get(id int) (*entities.Link, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

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
			updated_at,
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

func (l *linkStorage) GetByUser(userID int) ([]*entities.Link, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

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
			updated_at,
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

func (l *linkStorage) GetByTag(tagID int) ([]*entities.Link, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

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
			updated_at,
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

func (l *linkStorage) Update(link *entities.Link) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

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

func (l *linkStorage) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := l.db.Exec(
		ctx,
		`UPDATE "links" SET deleted_at = $1 WHERE id = $2`,
		time.Now().Unix(),
		id,
	)
	return err
}
