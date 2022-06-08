package post

import "time"

type Post struct {
	Id        string
	UserId    string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
