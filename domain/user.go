package domain

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Msg  string `json:"msg"`
	Age  int    `json:"age"`
}

type UserUsecase interface {
	GetAll() ([]*User, error)
}

type UserRepository interface {
	GetAll() ([]*User, error)
}
