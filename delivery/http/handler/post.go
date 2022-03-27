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
	var query viewmodel.PostTitle
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

func (h *PostHandler) GetById(c *gin.Context) {
	var params viewmodel.PostID
	if err := c.ShouldBindUri(&params); err != nil {
		response.Json(http.StatusBadRequest, nil, err, c)
		return
	}
	post, err := h.postUsecase.GetById(params.ID)
	if err != nil {
		response.Json(http.StatusInternalServerError, nil, err, c)
	} else {
		response.Json(http.StatusOK, post, nil, c)
	}
}
