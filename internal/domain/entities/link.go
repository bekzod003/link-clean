package entities

import "time"

type FullLink struct {
	Link Link `json:"link"`
	// Additional fields
	User User `json:"user"`
	Tag  Tag  `json:"tag"`
}

type Link struct {
	ID          int64     `json:"id"`
	URL         string    `json:"url"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	DeletedAt   int64     `json:"-"`

	UserID int64 `json:"user_id"`
	TagID  int64 `json:"tag_id"`
}
