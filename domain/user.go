package domain

type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Msg  string `json:"msg"`
	Age  int    `json:"age"`
}

type UserUsecase interface {
	GetAll(name string) ([]*User, error)
	GetById(id int) (*User, error)
}

type UserRepository interface {
	GetAll() ([]*User, error)
	GetAllByName(name string) ([]*User, error)
	GetById(id int) (*User, error)
}
