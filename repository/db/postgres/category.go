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

func (r *categoryRepo) GetAllByName(name string) ([]*domain.Category, error) {
	var categories []*domain.Category
	result := r.db.Where("name LIKE ?", "%"+name+"%").Find(&categories)
	return categories, result.Error
}

func (r *categoryRepo) GetById(id uint) (*domain.Category, error) {
	var category *domain.Category
	result := r.db.First(&category, id)
	return category, result.Error
}

func (r *categoryRepo) Create(category *domain.Category) (categoryID uint, err error) {
	err = r.db.Create(&category).Error
	categoryID = category.ID
	return
}
