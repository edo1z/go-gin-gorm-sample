package handler

import (
	"net/http"
	"web3ten0/go-gin-gorm-sample/domain"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postUsecase domain.PostUsecase
}

func NewPostHandler(pu domain.PostUsecase) *PostHandler {
	return &PostHandler{postUsecase: pu}
}

func (h *PostHandler) GetAll(c *gin.Context) {
	posts, err := h.postUsecase.GetAll()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, posts)
	}
}
