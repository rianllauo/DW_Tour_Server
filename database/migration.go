package database

import (
	"dewetour/models"
	"dewetour/pkg/mysql"
	"fmt"
)

// Automatic Migration if Running App
func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Trip{},
		&models.Country{},
		&models.Transaction{},
	)

	if err != nil {
		fmt.Println(err)
		fmt.Println(err.Error())
	}

	fmt.Println("Migration Success")
}
