package usecase

import (
	"web3ten0/go-gin-gorm-sample/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(u domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepo: u}
}

func (u *userUsecase) GetAll(name string) ([]*domain.User, error) {
	if name == "" {
		return u.userRepo.GetAll()
	} else {
		return u.userRepo.GetAllByName(name)
	}
}

func (u *userUsecase) GetById(id uint) (*domain.User, error) {
	return u.userRepo.GetById(id)
}

func (u *userUsecase) Create(user *domain.User) (userID uint, err error) {
	return u.userRepo.Create(user)
}
