package handler

import (
	"net/http"
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

func (u *UserHandler) GetAll(c *gin.Context) {
	var query viewmodel.GetUsersRquest
	if err := c.ShouldBindQuery(&query); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	users, err := u.userUsecase.GetAll(query.Name)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, users)
	}
}
