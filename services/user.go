package services

import (
	"github.com/yesseneon/bookstore_users_api/domain/user"
	"github.com/yesseneon/bookstore_users_api/utils/cuserr"
)

func GetUser(id int) (*user.User, *cuserr.RESTError) {
	u := &user.User{ID: id}
	if err := u.Get(); err != nil {
		return nil, err
	}

	return u, nil
}

func CreateUser(u *user.User) (*user.User, *cuserr.RESTError) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.Create(); err != nil {
		return nil, err
	}

	return u, nil
}
