package conn

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	dsn := "user=postgres password=postgres dbname=bookstore_users port=5432 sslmode=disable TimeZone=Asia/Almaty"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("DB connected")
}
