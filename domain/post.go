package domain

import "time"

type Post struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `json:"user_id"`
	CategoryID uint      `json:"category_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type PostUsecase interface {
	GetAll(title string) ([]*Post, error)
	GetById(id uint) (*Post, error)
	Create(post *Post) (postID uint, err error)
}

type PostRepository interface {
	GetAll() ([]*Post, error)
	GetAllByTitle(title string) ([]*Post, error)
	GetById(id uint) (*Post, error)
	Create(post *Post) (postID uint, err error)
}
