package user

import (
	"regexp"
	"strings"
	"time"

	"github.com/yesseneon/bookstore_users_api/utils/cuserr"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) Validate() *cuserr.RESTError {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	match := regexp.MustCompile(".+@.+\\..+").Match([]byte(u.Email))

	if u.Email == "" || match == false {
		return cuserr.BadRequest("Invalid email address")
	}

	return nil
}
