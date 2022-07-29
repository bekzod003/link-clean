package entities

import "time"

type Tag struct {
	ID        int64
	Title     string
	UserID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt int64
}

type UpdateTag struct {
	ID    int64
	Title string
}
