package main

import (
	"dip/controller"
	"dip/repository"
	"dip/service"
)

func main() {

	userLikeMySQLRepository := repository.NewUserLikeMySQLRepository()
	userLikeMySQLService := service.NewUserLikeService(userLikeMySQLRepository)
	userMySQLController := controller.NewUserController(userLikeMySQLService)
	userMySQLController.Get()

	userLikePostgresRepository := repository.NewUserLikePostgresRepository()
	userLikePostgresService := service.NewUserLikeService(userLikePostgresRepository)
	userPostgresController := controller.NewUserController(userLikePostgresService)
	userPostgresController.Get()

	//新しいDBに対応する場合はリポジトリを新規実装するだけで良い！

}
