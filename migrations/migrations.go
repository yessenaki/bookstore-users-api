package migrations

import (
	"log"

	"github.com/yesseneon/bookstore_users_api/datasources/postgres/conn"
	"github.com/yesseneon/bookstore_users_api/domain/user"
)

func Migrate() {
	conn.DB.AutoMigrate(&user.User{})
	log.Println("Schema(s) successfully migrated")
}
