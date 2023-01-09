package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connection Database
func DatabaseInit() {
	var err error
	
	var PGHOST = os.Getenv("DB_HOST")
	var PGUSER = os.Getenv("DB_USER")
	var PGPASSWORD = os.Getenv("DB_PASSWORD")
	var PGDATABASE = os.Getenv("DB_NAME")
	var PGPORT = os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", PGHOST, PGUSER, PGPASSWORD, PGDATABASE, PGPORT)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Database")
}
