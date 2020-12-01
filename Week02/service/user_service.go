package service

import (
	"week02/dao"
	"week02/model"
)

func FindUserById(user_id string) (*model.User, error) {
	user, err := dao.SelectUserById("1")
	if err != nil {
		return nil, err
	}
	return user, nil
}