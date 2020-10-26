package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yesseneon/bookstore_users_api/domain/user"
	"github.com/yesseneon/bookstore_users_api/services"
	"github.com/yesseneon/bookstore_users_api/utils/errors"
)

func Get(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func Create(c *gin.Context) {
	var u *user.User

	if err := c.ShouldBindJSON(&u); err != nil {
		restErr := errors.BadRequest("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	u, restErr := services.CreateUser(u)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusCreated, u)
}
