package service

import (
	"dip/object"
	"dip/repository"
)

type UserLikeService struct {
	repo UserLikeFetcherInterface
}

func (userLikeService *UserLikeService) Fetch(userID int) *object.UserLike {
	return userLikeService.repo.FetchByUserID(userID)
}

func NewUserLikeMySQLService() *UserLikeService {
	return &UserLikeService{
		repo: repository.NewUserLikeMySQLRepository(),
	}
}

func NewUserLikePostgresService() *UserLikeService {
	return &UserLikeService{
		repo: repository.NewUserLikePostgresRepository(),
	}
}
