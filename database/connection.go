package database

import (
	"github.com/kedarnathpc/go-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("user:root@/go_auth"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database :(")
	}
	DB = connection
	connection.AutoMigrate(&models.User{})
}
