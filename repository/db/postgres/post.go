package postgres

import (
	"web3ten0/go-gin-gorm-sample/domain"

	"gorm.io/gorm"
)

type postRepo struct {
	db *gorm.DB
}

func NewPostRepo(db *gorm.DB) domain.PostRepository {
	return &postRepo{db}
}

func (r *postRepo) GetAll() ([]*domain.Post, error) {
	var posts []*domain.Post
	result := r.db.Find(&posts)
	return posts, result.Error
}

func (r *postRepo) GetAllByTitle(title string) ([]*domain.Post, error) {
	var posts []*domain.Post
	result := r.db.Where("title LIKE ?", "%"+title+"%").Find(&posts)
	return posts, result.Error
}
