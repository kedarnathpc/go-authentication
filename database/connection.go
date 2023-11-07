package database

import (
	"github.com/kedarnathpc/go-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect initializes the database connection.
func Connect() {
	// Establish a connection to the MySQL database.
	connection, err := gorm.Open(mysql.Open("user:root@/go_auth"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database :(")
	}

	// Store the database connection for use throughout the application.
	DB = connection

	// Auto-migrate the User model to create the corresponding table in the database.
	connection.AutoMigrate(&models.User{})
}
