package usecase

import "web3ten0/go-gin-gorm-sample/domain"

type categoryUsecase struct {
	categoryRepo domain.CategoryRepository
}

func NewCategoryUsecase(r domain.CategoryRepository) domain.CategoryUsecase {
	return &categoryUsecase{categoryRepo: r}
}

func (u *categoryUsecase) GetAll() ([]*domain.Category, error) {
	return u.categoryRepo.GetAll()
}
