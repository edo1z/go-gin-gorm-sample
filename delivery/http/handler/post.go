package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	msg := "GetPosts"
	c.String(http.StatusOK, msg)
}
