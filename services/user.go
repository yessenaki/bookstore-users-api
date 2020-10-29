package services

import (
	"github.com/yesseneon/bookstore_users_api/domain/user"
	"github.com/yesseneon/bookstore_users_api/utils/encryption"
	"github.com/yesseneon/bookstore_users_api/utils/errors"
)

func CreateUser(u *user.User) (*user.User, *errors.RESTError) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	u.Status = user.StatusActive
	password, err := encryption.EncryptPassword(u.Password)
	if err != nil {
		return nil, errors.InternalServerError()
	}
	u.Password = password

	if err := u.Create(); err != nil {
		return nil, err
	}

	return u, nil
}

func FindUsers(status string) ([]user.User, *errors.RESTError) {
	var u *user.User
	return u.Find(status)
}

func GetUser(id int) (*user.User, *errors.RESTError) {
	u := &user.User{ID: id}
	if err := u.Get(); err != nil {
		return nil, err
	}

	return u, nil
}

func UpdateUser(u *user.User) (*user.User, *errors.RESTError) {
	_, err := GetUser(u.ID)
	if err != nil {
		return nil, err
	}

	if err := u.Update(); err != nil {
		return nil, err
	}

	return u, nil
}

func PartUpdateUser(u *user.User) (*user.User, *errors.RESTError) {
	eu, err := GetUser(u.ID) // existing user
	if err != nil {
		return nil, err
	}

	if err := eu.PartUpdate(u); err != nil {
		return nil, err
	}

	return eu, nil
}

func DeleteUser(id int) *errors.RESTError {
	_, err := GetUser(id)
	if err != nil {
		return err
	}

	u := &user.User{ID: id}
	return u.Delete()
}
