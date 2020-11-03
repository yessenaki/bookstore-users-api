package user

import (
	goerrors "errors"
	"strings"

	"github.com/yesseneon/bookstore_users_api/datasources/postgres/conn"
	"github.com/yesseneon/bookstore_users_api/logger"
	"github.com/yesseneon/bookstore_utils/errors"
	"gorm.io/gorm"
)

func (u *User) Create() *errors.RESTError {
	res := conn.DB.Create(&u)
	err := res.Error
	if err != nil {
		return getDBError(err)
	}

	return nil
}

func (u *User) Find(status string) ([]User, *errors.RESTError) {
	var users []User
	res := conn.DB.Where("status=?", status).Find(&users)
	err := res.Error
	if err != nil {
		return nil, getDBError(err)
	}

	return users, nil
}

func (u *User) Get() *errors.RESTError {
	res := conn.DB.First(&u)
	err := res.Error
	if err != nil {
		return getDBError(err)
	}

	return nil
}

func (u *User) Update() *errors.RESTError {
	res := conn.DB.Save(&u)
	err := res.Error
	if err != nil {
		return getDBError(err)
	}

	return nil
}

func (eu *User) PartUpdate(u *User) *errors.RESTError {
	res := conn.DB.Model(&eu).Updates(u)
	err := res.Error
	if err != nil {
		return getDBError(err)
	}

	return nil
}

func (u *User) Delete() *errors.RESTError {
	res := conn.DB.Delete(&u)
	err := res.Error
	if err != nil {
		return getDBError(err)
	}

	return nil
}

func (u *User) FindByEmail() *errors.RESTError {
	res := conn.DB.Where("email=? AND status=?", u.Email, StatusActive).First(&u)
	err := res.Error
	if err != nil {
		return getDBError(err)
	}

	return nil
}

func getDBError(err error) *errors.RESTError {
	if goerrors.Is(err, gorm.ErrRecordNotFound) {
		return errors.NotFound()
	}

	if strings.Contains(err.Error(), "users_email_key") {
		return errors.BadRequest("This email address already exists")
	}

	logger.Error("Error while trying to execute sql query", err)
	return errors.InternalServerError()
}
