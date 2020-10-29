package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yesseneon/bookstore_users_api/domain/user"
	"github.com/yesseneon/bookstore_users_api/services"
	"github.com/yesseneon/bookstore_users_api/utils/errors"
)

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

	c.JSON(http.StatusCreated, u.Marshal(false))
}

func Find(c *gin.Context) {
	users, restErr := services.FindUsers(c.Query("status"))
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	res := make([]interface{}, len(users))
	for i, u := range users {
		res[i] = u.Marshal(c.GetHeader("X-Public") == "true")
	}

	c.JSON(http.StatusOK, users)
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

	c.JSON(http.StatusOK, u.Marshal(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	id, restErr := getUserID(c.Param("id"))
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	var u *user.User
	if err := c.ShouldBindJSON(&u); err != nil {
		restErr := errors.BadRequest("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	u.ID = id
	if c.Request.Method == http.MethodPatch {
		u, restErr = services.PartUpdateUser(u)
	} else {
		u, restErr = services.UpdateUser(u)
	}

	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, u.Marshal(false))
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

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func getUserID(paramID string) (int, *errors.RESTError) {
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return 0, errors.BadRequest("User ID must be a number")
	}

	return id, nil
}
