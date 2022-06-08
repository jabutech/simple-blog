package post

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(post Post) (Post, error)
	FindByTitle(title string) (Post, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(post Post) (Post, error) {
	// Create new post
	err := r.db.Save(&post).Error
	if err != nil {
		return post, err
	}

	return post, nil
}

func (r *repository) FindByTitle(title string) (Post, error) {
	var post Post

	err := r.db.Where("title = ?", title).Find(&post).Error
	if err != nil {
		return post, err
	}

	return post, nil
}
