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

func UpdateUser(u user.User, isPartial bool) (*user.User, *cuserr.RESTError) {
	eu, err := GetUser(u.ID) // existing user
	if err != nil {
		return nil, err
	}

	if isPartial == false {
		eu.FirstName = u.FirstName
		eu.LastName = u.LastName
		eu.Email = u.Email
	} else {
		if u.FirstName != "" {
			eu.FirstName = u.FirstName
		}
		if u.LastName != "" {
			eu.LastName = u.LastName
		}
		if u.Email != "" {
			eu.Email = u.Email
		}
	}

	if err := eu.Update(); err != nil {
		return nil, err
	}

	return eu, nil
}
