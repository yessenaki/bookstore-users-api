package user

import (
	"errors"
	"fmt"
	"strings"

	"github.com/yesseneon/bookstore_users_api/datasources/postgres/conn"
	"github.com/yesseneon/bookstore_users_api/utils/cuserr"
	"gorm.io/gorm"
)

func (u *User) Create() *cuserr.RESTError {
	res := conn.DB.Create(&u)
	err := res.Error
	if err != nil {
		if strings.Contains(err.Error(), "users_email_key") {
			return cuserr.BadRequest(fmt.Sprintf("Email %s already exists", u.Email))
		}

		return cuserr.InternalServerError("Error while trying to create user")
	}

	return nil
}

func (u *User) Get() *cuserr.RESTError {
	res := conn.DB.First(&u, u.ID)
	err := res.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return cuserr.NotFound("No record found matching the given ID")
		}

		return cuserr.InternalServerError(fmt.Sprintf("Error while trying to get user %d", u.ID))
	}

	return nil
}

func (u *User) Update() *cuserr.RESTError {
	res := conn.DB.Save(&u)
	err := res.Error
	if err != nil {
		if strings.Contains(err.Error(), "users_email_key") {
			return cuserr.BadRequest(fmt.Sprintf("Email %s already exists", u.Email))
		}

		return cuserr.InternalServerError(fmt.Sprintf("Error while trying to update user %d", u.ID))
	}

	return nil
}
