package usecase

import "web3ten0/go-gin-gorm-sample/domain"

type postUsecase struct {
	postRepo domain.PostRepository
}

func NewPostUsecase(r domain.PostRepository) domain.PostUsecase {
	return &postUsecase{postRepo: r}
}

func (u *postUsecase) GetAll() ([]*domain.Post, error) {
	return u.postRepo.GetAll()
}
