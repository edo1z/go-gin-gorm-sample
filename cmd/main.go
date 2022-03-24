package main

import (
	"web3ten0/go-gin-gorm-sample/delivery/http/handler"
	"web3ten0/go-gin-gorm-sample/repository/db/postgres"
	"web3ten0/go-gin-gorm-sample/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	db := postgres.DbOpen()
	defer postgres.DbClose(db)

	userRepo := postgres.NewUserRepo(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	r := gin.Default()
	r.GET("/users", userHandler.GetAll)
	r.GET("/categories", handler.GetCategories)
	r.GET("/posts", handler.GetPosts)
	r.Run()
}
