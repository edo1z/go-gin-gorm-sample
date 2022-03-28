package domain

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Msg       string    `json:"msg"`
	Age       uint      `json:"age"`
	ImgUrl    string    `json:"img_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserUsecase interface {
	GetAll(name string) ([]*User, error)
	GetById(id uint) (*User, error)
	Create(user *User) (userID uint, err error)
}

type UserRepository interface {
	GetAll() ([]*User, error)
	GetAllByName(name string) ([]*User, error)
	GetById(id uint) (*User, error)
	Create(user *User) (userID uint, err error)
}
