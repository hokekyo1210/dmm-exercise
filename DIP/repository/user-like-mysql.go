package repository

import "dip/object"

type UserLikeMySQLRepository struct {
}

func (userLikeMySQLRepository *UserLikeMySQLRepository) FetchByUserID(userID int) *object.UserLike {

	//DBからイイね情報を取ってくる処理

	return &object.UserLike{ //仮の値
		UserID:    -1,
		PartnerID: -1,
	}
}

func NewUserLikeMySQLRepository() *UserLikeMySQLRepository {
	return &UserLikeMySQLRepository{}
}
