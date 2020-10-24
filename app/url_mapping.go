package app

import (
	"github.com/yesseneon/bookstore_users_api/controllers/ping"
	"github.com/yesseneon/bookstore_users_api/controllers/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:id", user.Get)
	router.POST("/users", user.Create)
}
