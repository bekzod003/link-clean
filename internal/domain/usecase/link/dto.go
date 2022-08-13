package link

import "time"

type Link struct{}

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"-"`
}
