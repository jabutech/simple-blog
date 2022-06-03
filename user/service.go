package user

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input RegisterInput) (User, error)
	IsEmailAvailable(email string) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// Function for register data user
func (s *service) Register(input RegisterInput) (User, error) {
	// Passing input into object user
	user := User{}
	user.Fullname = input.Fullname
	user.Email = input.Email

	// Create new variable isAdmin with default value zero (0) / false
	isAdmin := 0
	// If value `input.IsAdmin` is available / true
	if input.IsAdmin {
		// Change value isAdmin to (1) true
		isAdmin = 1
	}
	// Passing IsAdmin
	user.IsAdmin = isAdmin

	// Generate uuid for id
	id := uuid.New()
	user.ID = id.String()

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	// Passing password with passwordHash
	user.Password = string(passwordHash)

	// Save uset to db, with repository
	newUser, err := s.repository.Save(user)
	// Check if error
	if err != nil {
		return newUser, err
	}

	// If success return new user without error
	return newUser, nil
}

// EmailIsAvailable for check if email already exists or not
func (s *service) IsEmailAvailable(email string) (bool, error) {

	// Find email on db with repository
	user, err := s.repository.FindByEmail(email)
	// If error
	if err != nil {
		return false, err
	}

	// If user.Id must be empty string / not is exist
	if user.ID == "" {
		return false, nil
	}

	// If is exist
	return true, nil
}
