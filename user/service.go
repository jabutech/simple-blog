package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	Register(input RegisterInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

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
