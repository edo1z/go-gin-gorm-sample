package domain

type User struct {
	Id   int
	Name string
	Msg  string
	Age  int
}

type UserUsecase interface {
	GetAll() ([]*User, error)
}

type UserRepository interface {
	GetAll() ([]*User, error)
}
