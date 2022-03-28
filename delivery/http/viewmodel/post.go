package viewmodel

import "web3ten0/go-gin-gorm-sample/domain"

type PostTitle struct {
	Title string `form:"title"`
}

type PostID struct {
	ID uint `uri:"id" binding:"required"`
}

type PostForCreate struct {
	UserID     uint   `json:"user_id" binding:"required"`
	CategoryID uint   `json:"category_id" binding:"required"`
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

func (p PostForCreate) ToModel() *domain.Post {
	return &domain.Post{
		ID:         0,
		UserID:     p.UserID,
		CategoryID: p.CategoryID,
		Title:      p.Title,
		Content:    p.Content,
	}
}
