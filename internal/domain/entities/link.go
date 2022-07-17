package entities

import "time"

type FullLink struct {
	Link Link `json:"link"`
	// Additional fields
	User User `json:"user"`
	Tag  Tag  `json:"tag"`
}

type Link struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	DeletedAt   int64     `json:"-"`

	UserId int64 `json:"user_id"`
	TagId  int64 `json:"tag_id"`
}
