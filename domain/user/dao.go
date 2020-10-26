package user

import (
	"fmt"
	"strings"

	"github.com/yesseneon/bookstore_users_api/datasources/postgres/conn"
	"github.com/yesseneon/bookstore_users_api/utils/errors"
)

func (u *User) Get() *errors.RESTError {
	return nil
}

func (u *User) Create() *errors.RESTError {
	result := conn.DB.Create(&u)
	err := result.Error
	if err != nil {
		if strings.Contains(err.Error(), "users_email_key") {
			return errors.BadRequest(fmt.Sprintf("Email %s already exists", u.Email))
		}

		return errors.InternalServerError("Error while trying to create user")
	}

	return nil
}
