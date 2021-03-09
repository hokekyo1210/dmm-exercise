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

func NewUserController(service *service.UserLikeService) *UserContoller {
	return &UserContoller{
		service: service,
	}
}
