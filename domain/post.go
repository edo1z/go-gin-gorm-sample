package domain

type Post struct {
	Id         int
	UserId     int
	CategoryId int
	Title      string
	Content    string
}

type PostUsecase interface {
}

type PostRepository interface {
}
