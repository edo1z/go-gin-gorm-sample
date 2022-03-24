package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	msg := "GetCategories"
	c.String(http.StatusOK, msg)
}
