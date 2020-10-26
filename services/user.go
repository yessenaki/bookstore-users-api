package services

import (
	"github.com/yesseneon/bookstore_users_api/domain/user"
	"github.com/yesseneon/bookstore_users_api/utils/errors"
)

func GetUser(id int) (*user.User, *errors.RESTError) {
	return nil, nil
}

func CreateUser(u *user.User) (*user.User, *errors.RESTError) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.Create(); err != nil {
		return nil, err
	}

	return u, nil
}
