package domain

type Post struct {
	Id         int
	UserId     int
	CategoryId int
	Title      string
	Content    string
}

type PostUsecase interface {
	GetAll() ([]*Post, error)
}

type PostRepository interface {
	GetAll() ([]*Post, error)
}
