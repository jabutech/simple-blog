package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
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
