package post

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jabutech/simple-blog/user"
)

type Service interface {
	Create(title CreatePostInput) (Post, error)
	GetPosts(title string, user user.User) ([]Post, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input CreatePostInput) (Post, error) {
	// Passing input into object post
	post := Post{}
	post.Title = input.Title
	post.UserId = input.UserId
	// Generate uuid for post id
	id := uuid.New()
	post.Id = id.String()

	// Check title is exist
	checkTitle, err := s.repository.TitleIsExist(input.Title)
	if err != nil {
		return checkTitle, err
	}

	// If title id is not empty
	if checkTitle.Id != "" {
		return checkTitle, errors.New("title already exists")
	}

	// Save to db
	newPost, err := s.repository.Save(post)
	if err != nil {
		return newPost, err
	}

	// Success, return newPost
	return newPost, nil

}

func (s *service) GetPosts(title string, user user.User) ([]Post, error) {
	// If parameter title not empty string
	if title != "" {
		posts, err := s.repository.FindByTitle(title)
		if err != nil {
			return posts, err
		}

		return posts, nil
	}

	// Find all posts
	posts, err := s.repository.FindAll(user)
	if err != nil {
		return posts, err
	}

	return posts, nil

}
