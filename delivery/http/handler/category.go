package handler

import (
	"net/http"
	"web3ten0/go-gin-gorm-sample/delivery/http/handler/response"
	"web3ten0/go-gin-gorm-sample/delivery/http/viewmodel"
	"web3ten0/go-gin-gorm-sample/domain"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryUsecase domain.CategoryUsecase
}

func NewCategoryHandler(cu domain.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{categoryUsecase: cu}
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	var query viewmodel.GetCategoryRquest
	if err := c.ShouldBindQuery(&query); err != nil {
		response.Json(http.StatusBadRequest, nil, err, c)
		return
	}
	categories, err := h.categoryUsecase.GetAll(query.Name)
	if err != nil {
		response.Json(http.StatusInternalServerError, nil, err, c)
	} else {
		response.Json(http.StatusOK, categories, nil, c)
	}
}
