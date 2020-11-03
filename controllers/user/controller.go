package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yesseneon/bookstore_oauth_lib/oauth"
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

	u, restErr := services.UserService.CreateUser(u)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusCreated, u.Marshal(false))
}

func Find(c *gin.Context) {
	users, restErr := services.UserService.FindUsers(c.Query("status"))
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
	if restErr := oauth.AuthenticateUser(c.Request); restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	if callerID := oauth.GetCallerID(c.Request); callerID == 0 {
		restErr := errors.RESTError{
			Status:  http.StatusUnauthorized,
			Message: "Resource not available",
		}
		c.JSON(restErr.Status, restErr)
		return
	}

	id, restErr := getUserID(c.Param("id"))
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	u, restErr := services.UserService.GetUser(id)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	if oauth.GetCallerID(c.Request) == u.ID {
		c.JSON(http.StatusOK, u.Marshal(false))
		return
	}

	c.JSON(http.StatusOK, u.Marshal(oauth.IsPublic(c.Request)))
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
		u, restErr = services.UserService.PartUpdateUser(u)
	} else {
		u, restErr = services.UserService.UpdateUser(u)
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

	if restErr := services.UserService.DeleteUser(id); restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func Login(c *gin.Context) {
	var data user.LoginData
	if err := c.ShouldBindJSON(&data); err != nil {
		restErr := errors.BadRequest("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	u, restErr := services.UserService.LoginUser(data)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, u.Marshal(false))
}

func getUserID(paramID string) (int, *errors.RESTError) {
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return 0, errors.BadRequest("User ID must be a number")
	}

	return id, nil
}
