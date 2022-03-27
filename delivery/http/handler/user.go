package handler

import (
	"net/http"
	"web3ten0/go-gin-gorm-sample/delivery/http/handler/response"
	"web3ten0/go-gin-gorm-sample/delivery/http/viewmodel"
	"web3ten0/go-gin-gorm-sample/domain"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(uu domain.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: uu}
}

func (h *UserHandler) GetAll(c *gin.Context) {
	var query viewmodel.UserName
	if err := c.ShouldBindQuery(&query); err != nil {
		response.Json(http.StatusBadRequest, nil, err, c)
		return
	}
	users, err := h.userUsecase.GetAll(query.Name)
	if err != nil {
		response.Json(http.StatusInternalServerError, nil, err, c)
	} else {
		response.Json(http.StatusOK, users, nil, c)
	}
}

func (h *UserHandler) GetById(c *gin.Context) {
	var params viewmodel.UserID
	if err := c.ShouldBindUri(&params); err != nil {
		response.Json(http.StatusBadRequest, nil, err, c)
		return
	}
	user, err := h.userUsecase.GetById(params.ID)
	if err != nil {
		response.Json(http.StatusInternalServerError, nil, err, c)
	} else {
		response.Json(http.StatusOK, user, nil, c)
	}
}
