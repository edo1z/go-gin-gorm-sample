package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	msg := "GetUsers"
	c.String(http.StatusOK, msg)
}
