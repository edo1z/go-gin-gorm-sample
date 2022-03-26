package domain

type Post struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id"`
	CategoryId int    `json:"category_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}

type PostUsecase interface {
	GetAll(title string) ([]*Post, error)
}

type PostRepository interface {
	GetAll() ([]*Post, error)
	GetAllByTitle(title string) ([]*Post, error)
}
