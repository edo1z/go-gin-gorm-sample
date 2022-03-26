package domain

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryUsecase interface {
	GetAll() ([]*Category, error)
}

type CategoryRepository interface {
	GetAll() ([]*Category, error)
}
