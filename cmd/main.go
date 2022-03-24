package main

import (
	"web3ten0/go-gin-gorm-sample/delivery/http/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/users", handler.GetUsers)
	r.GET("/categories", handler.GetCategories)
	r.GET("/posts", handler.GetPosts)
	r.Run()
}
