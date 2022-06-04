package auth

import (
	"github.com/jabutech/simple-blog/user"
	"gorm.io/gorm"
)

type Repository interface {
	Save(user user.User) (user.User, error)
	FindByEmail(email string) (user.User, error)
}

type repository struct {
	db *gorm.DB
}

// Instance
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Save for save data user to db
func (r *repository) Save(user user.User) (user.User, error) {
	// Create new user on db
	err := r.db.Save(&user).Error
	// If err return object data user, with error
	if err != nil {
		return user, err
	}

	// If success return new data user, with no error
	return user, nil
}

func (r *repository) FindByEmail(email string) (user.User, error) {
	var user user.User

	// Find user by email
	err := r.db.Where("email = ?", email).Find(&user).Error
	// If err return object data user, with error
	if err != nil {
		return user, err
	}

	// If success return new data user, with no error
	return user, nil
}
