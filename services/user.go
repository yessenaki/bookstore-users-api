package services

import (
	"github.com/yesseneon/bookstore_users_api/domain/user"
	"github.com/yesseneon/bookstore_users_api/utils/encryption"
	"github.com/yesseneon/bookstore_utils/errors"
	"golang.org/x/crypto/bcrypt"
)

var UserService userServiceInterface = &userService{}

type userService struct{}

type userServiceInterface interface {
	CreateUser(*user.User) (*user.User, *errors.RESTError)
	FindUsers(string) ([]user.User, *errors.RESTError)
	GetUser(int) (*user.User, *errors.RESTError)
	UpdateUser(*user.User) (*user.User, *errors.RESTError)
	PartUpdateUser(*user.User) (*user.User, *errors.RESTError)
	DeleteUser(int) *errors.RESTError
	LoginUser(user.LoginData) (*user.User, *errors.RESTError)
}

func (srv *userService) CreateUser(u *user.User) (*user.User, *errors.RESTError) {
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

func (srv *userService) FindUsers(status string) ([]user.User, *errors.RESTError) {
	var u *user.User
	return u.Find(status)
}

func (srv *userService) GetUser(id int) (*user.User, *errors.RESTError) {
	u := &user.User{ID: id}
	if err := u.Get(); err != nil {
		return nil, err
	}

	return u, nil
}

func (srv *userService) UpdateUser(u *user.User) (*user.User, *errors.RESTError) {
	_, err := srv.GetUser(u.ID)
	if err != nil {
		return nil, err
	}

	if err := u.Update(); err != nil {
		return nil, err
	}

	return u, nil
}

func (srv *userService) PartUpdateUser(u *user.User) (*user.User, *errors.RESTError) {
	eu, err := srv.GetUser(u.ID) // existing user
	if err != nil {
		return nil, err
	}

	if err := eu.PartUpdate(u); err != nil {
		return nil, err
	}

	return eu, nil
}

func (srv *userService) DeleteUser(id int) *errors.RESTError {
	_, err := srv.GetUser(id)
	if err != nil {
		return err
	}

	u := &user.User{ID: id}
	return u.Delete()
}

func (srv *userService) LoginUser(data user.LoginData) (*user.User, *errors.RESTError) {
	u := &user.User{Email: data.Email}
	if restErr := u.FindByEmail(); restErr != nil {
		return nil, restErr
	}

	// Does the entered password match the stored password?
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(data.Password))
	if err != nil {
		return nil, errors.BadRequest("Password is incorrect")
	}

	return u, nil
}
