package service

import "dip/object"

type UserLikeFetcherInterface interface {
	FetchByUserID(int) *object.UserLike
}
