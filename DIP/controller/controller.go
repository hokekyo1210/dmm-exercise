package controller

import "dip/object"

type UserServiceInterface interface {
	Fetch(int) *object.UserLike
}
