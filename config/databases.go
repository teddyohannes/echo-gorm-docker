package config

import (
	"echo-gorm-docker/entity"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

//SetupDatabaseConnection is creating a new connection to our database
// func SetupDatabasesConnection() {
// 	errEnv := godotenv.Load()
// 	if errEnv != nil {
// 		panic("Failed to load env file")
// 	}

// 	dbUser := os.Getenv("DB_USER")
// 	dbPass := os.Getenv("DB_PASS")
// 	dbHost := os.Getenv("DB_HOST")
// 	dbName := os.Getenv("DB_NAME")

// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		panic("Failed to create a connection to database")
// 	}

// 	db.AutoMigrate(&entity.Book{}, &entity.User{})
// }

//CloseDatabaseConnection method is closing a connection between your app and your db
// func CloseDatabaseConnection(db *gorm.DB) {
// 	dbSQL, err := db.DB()
// 	if err != nil {
// 		panic("Failed to close connection from database")
// 	}
// 	dbSQL.Close()
// }

func Init() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	connect_string := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err = gorm.Open(mysql.Open(connect_string), &gorm.Config{})
	// defer db.Close()
	if err != nil {
		panic("DB Connection Error")
	}
	db.AutoMigrate(&entity.Book{}, &entity.User{})
	return db
}

func DbManager() *gorm.DB {
	return db
}
