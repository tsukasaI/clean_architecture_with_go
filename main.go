package main

import (
	"clean_architecture_with_go/domain/usecase"
	"clean_architecture_with_go/handler"
	"clean_architecture_with_go/repository"

	"gorm.io/driver/postgres"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	// データベースの接続設定
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// ユーザーリポジトリの作成
	userRepository := repository.NewUserRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	// ルーティングの設定
	router := gin.Default()
	router.GET("/users", userHandler.GetUsers)

	// サーバーの起動
	router.Run()
}
