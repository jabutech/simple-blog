package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/jabutech/simple-blog/helper"
	"github.com/jabutech/simple-blog/user"
	"github.com/jabutech/simple-blog/util"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input RegisterInput) (user.User, error)
	Login(input LoginInput) (user.User, error)
	IsEmailAvailable(email string) (bool, error)
	GenerateToken(user user.User) (string, error)
}

type service struct {
	repository user.Repository
}

func NewService(repository user.Repository) *service {
	return &service{repository}
}

// Function for register data user
func (s *service) Register(input RegisterInput) (user.User, error) {
	// Passing input into object user
	user := user.User{}
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

func (s *service) Login(input LoginInput) (user.User, error) {
	// Get data input from request
	email := input.Email
	password := input.Password

	// Find user with repository
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	// If User id is empty string
	if user.ID == "" {
		return user, errors.New("email or password incorrect")
	}

	// If user is available, compare password hash with password from request use bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("email or password incorrect")
	}

	// If no error, return user
	return user, nil
}

type Claim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func (s *service) GenerateToken(user user.User) (string, error) {
	// Create 1 day
	expirationTime := time.Now().AddDate(0, 0, 1)

	// Create clain for payload token
	claim := Claim{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// Load config
	config, err := util.LoadConfig("../") // "." as location file app.env in root folder
	helper.FatalError("error load config: ", err)

	// Signed token with secret key
	signedToken, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return signedToken, err
	}

	// If success, return token
	return signedToken, nil
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
