package controller

import "dip/service"

type UserContoller struct {
	service UserServiceInterface
}

func (userController *UserContoller) Get() {
	userID := 1
	userController.service.Fetch(userID)

	//処理する
}

func NewUserMySQLController() *UserContoller {
	return &UserContoller{
		service: service.NewUserLikeMySQLService(),
	}
}

func NewUserPostgresController() *UserContoller {
	return &UserContoller{
		service: service.NewUserLikePostgresService(),
	}
}
