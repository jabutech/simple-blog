package post

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jabutech/simple-blog/user"
)

type Service interface {
	Create(title CreateOrUpdatePostInput) (Post, error)
	Update(Id CreateOrUpdatePostInput) (Post, error)
	GetPosts(title string, user user.User) ([]Post, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input CreateOrUpdatePostInput) (Post, error) {
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

func (s *service) Update(Input CreateOrUpdatePostInput) (Post, error) {
	// Find post by id
	post, err := s.repository.FindById(Input.UserId)
	if err != nil {
		return post, err
	}

	// If post not available
	if post.Id == "" {
		return post, errors.New("post not found")
	}

	// Passing update request title to object post title
	post.Title = Input.Title

	// Update post
	updatedPost, err := s.repository.Update(post)
	if err != nil {
		return updatedPost, err
	}

	// Success, return updated post
	return updatedPost, nil
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
