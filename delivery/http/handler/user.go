package handler

import (
	"log"
	"net/http"
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
	users, err := u.userUsecase.GetAll()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		log.Println("users length:", len(users))
		c.String(http.StatusOK, "OK")
	}
}
