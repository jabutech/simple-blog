package post

import (
	"github.com/jabutech/simple-blog/user"
	"gorm.io/gorm"
)

type Repository interface {
	Save(post Post) (Post, error)
	FindAll(user user.User) ([]Post, error)
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

func (r *repository) FindAll(user user.User) ([]Post, error) {
	var posts []Post

	// If user is_admin not 1 (admin)
	if user.IsAdmin != 1 {
		// Find all posts except user_id same with current id
		err := r.db.Preload("User").Where("user_id NOT IN (?)", user.ID).Find(&posts).Error
		if err != nil {
			return posts, err
		}

		return posts, nil
	}

	// If not, find all post
	err := r.db.Preload("User").Find(&posts).Error
	if err != nil {
		return posts, err
	}

	return posts, nil
}

func (r *repository) FindByTitle(title string) (Post, error) {
	var post Post

	err := r.db.Where("title = ?", title).Find(&post).Error
	if err != nil {
		return post, err
	}

	return post, nil
}
