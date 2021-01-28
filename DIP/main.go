package main

import (
	"dip/controller"
)

func main() {

	userController := controller.NewUserMySQLController() //こちらはMySQL用
	userController.Get()

	userController = controller.NewUserPostgresController() //こちらはPostgreSQL用
	userController.Get()

}
