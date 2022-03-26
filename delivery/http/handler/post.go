package handler

import (
	"net/http"
	"web3ten0/go-gin-gorm-sample/delivery/http/handler/response"
	"web3ten0/go-gin-gorm-sample/delivery/http/viewmodel"
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
	var query viewmodel.GetPostRquest
	if err := c.ShouldBindQuery(&query); err != nil {
		response.Json(http.StatusBadRequest, nil, err, c)
		return
	}
	posts, err := h.postUsecase.GetAll(query.Title)
	if err != nil {
		response.Json(http.StatusInternalServerError, nil, err, c)
	} else {
		response.Json(http.StatusOK, posts, nil, c)
	}
}
