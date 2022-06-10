package post

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jabutech/simple-blog/user"
)

type Service interface {
	Create(title CreatePostInput) (Post, error)
	Update(Id UpdatePostInput) (Post, error)
	Delete(Id string) error
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

func (s *service) Update(Input UpdatePostInput) (Post, error) {
	// Find post by id
	post, err := s.repository.FindById(Input.PostId)
	if err != nil {
		return post, err
	}

	// If post not available
	if post.Id == "" {
		return post, errors.New("post not found")
	}

	// If post user_id not same with current user is loggedin that requested update
	if post.UserId != Input.UserId {
		return post, errors.New("do not have access to this post")
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

func (s *service) Delete(Id string, userIsAdmin bool) (bool, error) {
	// If status is_admin current user loggedin not `true`
	if !userIsAdmin {
		// Return error
		return false, errors.New("access not allowed")
	}

	// Find post
	post, err := s.repository.FindById(Id)
	if err != nil {
		return false, err
	}

	// If post not available
	if post.Id == "" {
		return false, errors.New("post not found")
	}

	// Delete post
	err = s.repository.Delete(post)
	if err != nil {
		return false, err
	}

	// Success, return true for notif success delete
	return true, nil

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
