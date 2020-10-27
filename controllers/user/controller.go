package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yesseneon/bookstore_users_api/domain/user"
	"github.com/yesseneon/bookstore_users_api/services"
	"github.com/yesseneon/bookstore_users_api/utils/cuserr"
)

func Create(c *gin.Context) {
	var u *user.User

	if err := c.ShouldBindJSON(&u); err != nil {
		restErr := cuserr.BadRequest("Invalid JSON body")
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

func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restErr := cuserr.BadRequest("User ID must be a number")
		c.JSON(restErr.Status, restErr)
		return
	}

	u, restErr := services.GetUser(id)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, u)
}

func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restErr := cuserr.BadRequest("User ID must be a number")
		c.JSON(restErr.Status, restErr)
		return
	}

	var u user.User
	if err := c.ShouldBindJSON(&u); err != nil {
		restErr := cuserr.BadRequest("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	u.ID = id
	uu, restErr := services.UpdateUser(u, c.Request.Method == http.MethodPatch) // updated user
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, uu)
}
