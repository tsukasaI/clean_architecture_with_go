package api

import (
	"clean_architecture_with_go/domain/usecase"
	"clean_architecture_with_go/handler"
	"clean_architecture_with_go/infrastructure"
	"clean_architecture_with_go/repository"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	userRepository := repository.NewUserRepository(infrastructure.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)
	users := router.Group("/users")
	{
		users.GET("", userHandler.GetUsers)
	}

	return router
}
