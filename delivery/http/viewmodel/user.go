package viewmodel

import "web3ten0/go-gin-gorm-sample/domain"

type UserName struct {
	Name string `form:"name"`
}

type UserID struct {
	ID uint `uri:"id" binding:"required" json:"id"`
}

type UserForCreate struct {
	Name string `json:"name" binding:"required"`
	Msg  string `json:"msg"`
	Age  uint   `json:"age"`
}

func (u UserForCreate) ToModel() *domain.User {
	return &domain.User{
		ID:   0,
		Name: u.Name,
		Msg:  u.Msg,
		Age:  u.Age,
	}
}
