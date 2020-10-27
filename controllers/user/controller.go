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
	id, restErr := getUserID(c.Param("id"))
	if restErr != nil {
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
	id, restErr := getUserID(c.Param("id"))
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	var u *user.User
	if err := c.ShouldBindJSON(&u); err != nil {
		restErr := cuserr.BadRequest("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	u.ID = id
	uu, restErr := services.UpdateUser(u) // updated user
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, uu)
}

func PartUpdate(c *gin.Context) {
	id, restErr := getUserID(c.Param("id"))
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	var u *user.User
	if err := c.ShouldBindJSON(&u); err != nil {
		restErr := cuserr.BadRequest("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	u.ID = id
	uu, restErr := services.PartUpdateUser(u) // partially updated user
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, uu)
}

func Delete(c *gin.Context) {
	id, restErr := getUserID(c.Param("id"))
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	if restErr := services.DeleteUser(id); restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, struct{ Status string }{Status: "Deleted"})
}

func getUserID(paramID string) (int, *cuserr.RESTError) {
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return 0, cuserr.BadRequest("User ID must be a number")
	}

	return id, nil
}
