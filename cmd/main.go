package main

import (
	"web3ten0/go-gin-gorm-sample/delivery/http/handler"
	"web3ten0/go-gin-gorm-sample/repository/db/postgres"
	"web3ten0/go-gin-gorm-sample/usecase"

	"github.com/gin-contrib/cors"
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
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8001"}
	r.Use(cors.New(config))
	r.GET("/users", userHandler.GetAll)
	r.GET("/users/view/:id", userHandler.GetById)
	r.POST("/users/add", userHandler.Create)
	r.GET("/categories", categoryHandler.GetAll)
	r.GET("/categories/view/:id", categoryHandler.GetById)
	r.GET("/posts", postHandler.GetAll)
	r.GET("/posts/view/:id", postHandler.GetById)
	r.Run()
}
