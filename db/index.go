package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func InitDB() {
	dbUrl := os.Getenv("DATABASE_URL")
	var err error
	DBConn, err = gorm.Open(postgres.Open(dbUrl))
	if err != nil {
		fmt.Println("Failed to connect to database")
		panic("Failed to connect to database")
	}

	err = DBConn.Exec("CREATE EXTENSION IF NOT EXISTS\"uuid-ossp\"").Error
	if err != nil {
		fmt.Println("Can't install uuid extension")
		panic(err)
	}

	err = DBConn.AutoMigrate(&User{}, &SearchSettings{})
	if err != nil {
		fmt.Println("Can't do migrations...")
		panic(err)
	}

	fmt.Println("Connected to Database...")

}

func GetDB() *gorm.DB {
	return DBConn
}
