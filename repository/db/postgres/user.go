package postgres

import (
	"web3ten0/go-gin-gorm-sample/domain"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) domain.UserRepository {
	return &userRepo{db}
}

func (r *userRepo) GetAll() ([]*domain.User, error) {
	var users []*domain.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *userRepo) GetAllByName(name string) ([]*domain.User, error) {
	var users []*domain.User
	result := r.db.Where("name LIKE ?", "%"+name+"%").Find(&users)
	return users, result.Error
}

func (r *userRepo) GetById(id uint) (*domain.User, error) {
	var user *domain.User
	result := r.db.First(&user, id)
	return user, result.Error
}

func (r *userRepo) Create(user *domain.User) (userID uint, err error) {
	err = r.db.Create(&user).Error
	userID = user.ID
	return
}
