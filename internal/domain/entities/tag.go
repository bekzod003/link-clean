package entities

import "time"

type Tag struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt int64     `json:"-"`
}
