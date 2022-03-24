package usecase

import "web3ten0/go-gin-gorm-sample/domain"

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(u domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepo: u}
}

func (u *userUsecase) GetAll() ([]*domain.User, error) {
	return u.userRepo.GetAll()
}
