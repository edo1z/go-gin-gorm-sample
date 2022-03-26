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
	categoryRepo := postgres.NewCategoryRepo(db)
	postRepo := postgres.NewPostRepo(db)

	userUsecase := usecase.NewUserUsecase(userRepo)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	postUsecase := usecase.NewPostUsecase(postRepo)

	userHandler := handler.NewUserHandler(userUsecase)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)
	postHandler := handler.NewPostHandler(postUsecase)

	r := gin.Default()
	r.GET("/users", userHandler.GetAll)
	r.GET("/categories", categoryHandler.GetAll)
	r.GET("/posts", postHandler.GetAll)
	r.Run()
}
