package domain

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryUsecase interface {
	GetAll(name string) ([]*Category, error)
	GetById(id int) (*Category, error)
}

type CategoryRepository interface {
	GetAll() ([]*Category, error)
	GetAllByName(name string) ([]*Category, error)
	GetById(id int) (*Category, error)
}
