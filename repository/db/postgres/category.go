package postgres

import (
	"web3ten0/go-gin-gorm-sample/domain"

	"gorm.io/gorm"
)

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) domain.CategoryRepository {
	return &categoryRepo{db}
}

func (r *categoryRepo) GetAll() ([]*domain.Category, error) {
	var categories []*domain.Category
	result := r.db.Find(&categories)
	return categories, result.Error
}
