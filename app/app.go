package app

import (
	"github.com/gin-gonic/gin"
	"github.com/yesseneon/bookstore_users_api/datasources/postgres/conn"
	"github.com/yesseneon/bookstore_users_api/migrations"
)

var router = gin.Default()

func StartApp() {
	conn.InitDB()
	migrations.Migrate()

	mapUrls()
	router.Run(":8080")
}
