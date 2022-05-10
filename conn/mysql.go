package conn

import (
	"Ecommerce-User/model"
	"errors"
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
	if err = db.AutoMigrate(&model.User{}); err == nil && db.Migrator().HasTable(&model.User{}) {
		if err := db.First(&model.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&model.User{
				Nama:     "Risky Saputra",
				Alamat:   "Jl. Kebon Jeruk No. 1",
				Email:    "risky@gmail.com",
				NomorHP:  "081212121212",
				Password: "12345678",
				Role:     "seller",
			})
			db.Create(&model.User{
				Nama:     "Kevin Nainggolan",
				Alamat:   "Jl. Kebon Jeruk No. 3",
				Email:    "kevin@gmail.com",
				NomorHP:  "081212121313",
				Password: "12345678",
				Role:     "seller",
			})
			db.Create(&model.User{
				Nama:     "Ricky Metal",
				Alamat:   "Jl. Kebon Jeruk No. -1",
				Email:    "ricky@gmail.com",
				NomorHP:  "081212122390",
				Password: "12345678",
				Role:     "seller",
			})
			db.Create(&model.User{
				Nama:     "Rut Ferwati Feronika Scintya Leonardo Sitepu",
				Alamat:   "Jl. Kebon Jeruk No. 1000",
				Email:    "ferwati@gmail.com",
				NomorHP:  "08116161669",
				Password: "12345678",
				Role:     "customer",
			})
		}
	}
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close a connection database")
	}
	dbSQL.Close()
}
