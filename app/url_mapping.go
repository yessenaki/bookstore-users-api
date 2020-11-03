package app

import (
	"github.com/yesseneon/bookstore_users_api/controllers/ping"
	"github.com/yesseneon/bookstore_users_api/controllers/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", user.Create)
	router.GET("/users", user.Find)
	router.GET("/users/:id", user.Get)
	router.PUT("/users/:id", user.Update)
	router.PATCH("/users/:id", user.Update)
	router.DELETE("/users/:id", user.Delete)
	router.POST("/users/login", user.Login)
}
