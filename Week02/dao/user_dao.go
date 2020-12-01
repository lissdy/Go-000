package dao

import (
	"database/sql"
	"github.com/pkg/errors"
	"week02/model"
)

func SelectUserById(user_id string) (*model.User, error) {
	var user *model.User
	err := fakeDBQuery()
	switch {
	case err == sql.ErrNoRows:
		 return nil, errors.Wrap(err, "can not find user with id: " + user_id)
	case err != nil:
		return nil,  errors.Wrap(err, "Other query error")
	default:
		return user, nil
	}
}

func fakeDBQuery() error{
	return sql.ErrNoRows
}