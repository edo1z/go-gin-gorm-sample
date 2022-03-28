package domain

import "time"

type Category struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryUsecase interface {
	GetAll(name string) ([]*Category, error)
	GetById(id uint) (*Category, error)
	Create(category *Category) (categoryID uint, err error)
}

type CategoryRepository interface {
	GetAll() ([]*Category, error)
	GetAllByName(name string) ([]*Category, error)
	GetById(id uint) (*Category, error)
	Create(category *Category) (categoryID uint, err error)
}
