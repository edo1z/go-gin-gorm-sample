package usecase

import "web3ten0/go-gin-gorm-sample/domain"

type postUsecase struct {
	postRepo domain.PostRepository
}

func NewPostUsecase(r domain.PostRepository) domain.PostUsecase {
	return &postUsecase{postRepo: r}
}

func (u *postUsecase) GetAll(title string) ([]*domain.Post, error) {
	if title == "" {
		return u.postRepo.GetAll()
	} else {
		return u.postRepo.GetAllByTitle(title)
	}
}

func (u *postUsecase) GetById(id uint) (*domain.Post, error) {
	return u.postRepo.GetById(id)
}

func (u *postUsecase) Create(post *domain.Post) (postID uint, err error) {
	return u.postRepo.Create(post)
}
