package post

import (
	"time"

	"github.com/jabutech/simple-blog/user"
)

type Post struct {
	Id        string
	UserId    string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      user.User // Relation one to many to table users
}
