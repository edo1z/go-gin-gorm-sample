package usecase

import "web3ten0/go-gin-gorm-sample/domain"

type categoryUsecase struct {
	categoryRepo domain.CategoryRepository
}

func NewCategoryUsecase(r domain.CategoryRepository) domain.CategoryUsecase {
	return &categoryUsecase{categoryRepo: r}
}

func (u *categoryUsecase) GetAll(name string) ([]*domain.Category, error) {
	if name == "" {
		return u.categoryRepo.GetAll()
	} else {
		return u.categoryRepo.GetAllByName(name)
	}
}

func (u *categoryUsecase) GetById(id uint) (*domain.Category, error) {
	return u.categoryRepo.GetById(id)
}

func (u *categoryUsecase) Create(category *domain.Category) (categoryID uint, err error) {
	return u.categoryRepo.Create(category)
}
