package domain

type Category struct {
	Id   int
	Name string
}

type CategoryUsecase interface {
	GetAll() ([]*Category, error)
}

type CategoryRepository interface {
	GetAll() ([]*Category, error)
}
