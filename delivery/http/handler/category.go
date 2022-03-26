package handler

import (
	"log"
	"net/http"
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
	categories, err := h.categoryUsecase.GetAll()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		log.Println("categories length:", len(categories))
		c.String(http.StatusOK, "OK")
	}
}
