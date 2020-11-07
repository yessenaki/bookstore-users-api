package app

import (
	"github.com/gin-gonic/gin"
	"github.com/yesseneon/bookstore-users-api/datasources/postgres/conn"
	"github.com/yesseneon/bookstore-users-api/migrations"
)

var router = gin.Default()

func StartApp() {
	conn.InitDB()
	migrations.Migrate()

	mapUrls()
	router.Run(":8080")
}
