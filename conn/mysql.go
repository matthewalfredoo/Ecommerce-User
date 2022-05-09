package conn

import (
	"Ecommerce-User/model"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load environment file (.env)")
	}
	dbUser := os.Getenv("DB_USER") //Get the DB_USER data from .env file
	dbPass := os.Getenv("DB_PASS") //Get the DB_PASS data from .env file
	dbHost := os.Getenv("DB_HOST") //Get the DB_HOST data from .env file
	dbName := os.Getenv("DB_NAME") //Get the DB_NAME data from .env file

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName) //Create the DSN string which will be used to connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&model.User{})
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close a connection database")
	}
	dbSQL.Close()
}
