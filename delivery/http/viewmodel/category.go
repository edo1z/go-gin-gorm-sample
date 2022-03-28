package viewmodel

import "web3ten0/go-gin-gorm-sample/domain"

type CategoryName struct {
	Name string `form:"name"`
}

type CategoryID struct {
	ID uint `uri:"id" binding:"required"`
}

type CategoryForCreate struct {
	Name string `json:"name" binding:"required"`
}

func (c CategoryForCreate) ToModel() *domain.Category {
	return &domain.Category{
		ID:   0,
		Name: c.Name,
	}
}
