package user

import (
	"regexp"
	"strings"
	"time"

	"github.com/yesseneon/bookstore-utils/errors"
)

const (
	StatusActive   = "active"
	StatusInactive = "inactive"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"first_name" gorm:"size:60"`
	LastName  string    `json:"last_name" gorm:"size:60"`
	Email     string    `json:"email" gorm:"unique;size:60"`
	Password  string    `json:"password" gorm:"size:60"`
	Status    string    `json:"status" gorm:"size:30"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) Validate() *errors.RESTError {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	u.Password = strings.TrimSpace(u.Password)

	match := regexp.MustCompile(".+@.+\\..+").Match([]byte(u.Email))
	if u.Email == "" || match == false {
		return errors.BadRequest("Invalid email address")
	}

	if len(u.Password) < 6 || len(u.Password) > 20 {
		return errors.BadRequest("Your password must be 6-20 characters long")
	}

	return nil
}
