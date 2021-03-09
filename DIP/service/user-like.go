package service

import (
	"dip/object"
)

type UserLikeService struct {
	repo UserLikeFetcherInterface
}

func (userLikeService *UserLikeService) Fetch(userID int) *object.UserLike {
	return userLikeService.repo.FetchByUserID(userID)
}

func NewUserLikeService(repo UserLikeFetcherInterface) *UserLikeService {
	return &UserLikeService{
		repo: repo,
	}
}
