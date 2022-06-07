package user

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindAll() ([]User, error)
	FindById(id string) (User, error)
	FindByEmail(email string) (User, error)
	FindByFullname(fullname string) ([]User, error)
}

type repository struct {
	db *gorm.DB
}

// Instance
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Save for save data user to db
func (r *repository) Save(user User) (User, error) {
	// Create new user on db
	err := r.db.Save(&user).Error
	// If err return object data user, with error
	if err != nil {
		return user, err
	}

	// If success return new data user, with no error
	return user, nil
}

func (r *repository) FindAll() ([]User, error) {
	var users []User

	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) FindById(id string) (User, error) {
	var user User

	// Find user by email
	err := r.db.Where("id = ?", id).Find(&user).Error
	// If err return object data user, with error
	if err != nil {
		return user, err
	}

	// If success return new data user, with no error
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	// Find user by email
	err := r.db.Where("email = ?", email).Find(&user).Error
	// If err return object data user, with error
	if err != nil {
		return user, err
	}

	// If success return new data user, with no error
	return user, nil
}

func (r *repository) FindByFullname(fullname string) ([]User, error) {
	strFullname := fmt.Sprint("%" + fullname + "%")
	var users []User
	// Find user like fullname
	err := r.db.Where("fullname LIKE ?", strFullname).Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}
