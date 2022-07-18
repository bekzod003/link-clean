package entities

import "time"

type Tag struct {
	Id        int64
	Title     string
	UserID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt int64
}

type UpdateTag struct {
	Id    int64
	Title string
}
