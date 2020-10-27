package services

import (
	"github.com/yesseneon/bookstore_users_api/domain/user"
	"github.com/yesseneon/bookstore_users_api/utils/cuserr"
)

func CreateUser(u *user.User) (*user.User, *cuserr.RESTError) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.Create(); err != nil {
		return nil, err
	}

	return u, nil
}

func GetUser(id int) (*user.User, *cuserr.RESTError) {
	u := &user.User{ID: id}
	if err := u.Get(); err != nil {
		return nil, err
	}

	return u, nil
}

func UpdateUser(u *user.User) (*user.User, *cuserr.RESTError) {
	_, err := GetUser(u.ID) // existing user
	if err != nil {
		return nil, err
	}

	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.Update(); err != nil {
		return nil, err
	}

	return u, nil
}

func PartUpdateUser(u *user.User) (*user.User, *cuserr.RESTError) {
	eu, err := GetUser(u.ID) // existing user
	if err != nil {
		return nil, err
	}

	if u.Email != "" {
		if err := u.Validate(); err != nil {
			return nil, err
		}
	}

	if err := eu.PartUpdate(u); err != nil {
		return nil, err
	}

	return eu, nil
}

func DeleteUser(id int) *cuserr.RESTError {
	_, err := GetUser(id) // existing user
	if err != nil {
		return err
	}

	u := &user.User{ID: id}
	return u.Delete()
}
