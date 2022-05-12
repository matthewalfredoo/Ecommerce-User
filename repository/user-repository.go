package repository

import (
	"Ecommerce-User/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type UserRepository interface {
	FindByEmail(email string) model.User
	IsDuplicateEmail(email string) (conn *gorm.DB)
	InsertUser(user model.User) model.User
	VerifyCredential(email string, password string) interface{}
	UpdateUser(user model.User) model.User
	ProfileUser(userID string) model.User
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(conn *gorm.DB) UserRepository {
	return &userRepository{connection: conn}
}

// IsDuplicateEmail checks if the email is already in the database
func (db *userRepository) IsDuplicateEmail(email string) (conn *gorm.DB) {
	var user model.User
	return db.connection.Where("email = ?", email).Take(&user)
}

// InsertUser is invoked when a user registers an account
func (db *userRepository) InsertUser(user model.User) model.User {
	user.Password = HashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userRepository) FindByEmail(email string) model.User {
	var user model.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

// VerifyCredential is invoked when a user logs in
// This function checks if the email exists in the database
func (db *userRepository) VerifyCredential(email string, password string) interface{} {
	var user model.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userRepository) ProfileUser(userID string) model.User {
	var user model.User
	db.connection.Where("id = ?", userID).Take(&user)
	return user
}

// UpdateUser is invoked when a user updates their profile
func (db *userRepository) UpdateUser(user model.User) model.User {
	var tempUser model.User
	db.connection.Find(&tempUser, user.ID)

	if user.Password != "" {
		user.Password = HashAndSalt([]byte(user.Password))
	} else {
		user.Password = tempUser.Password
	}
	user.Role = tempUser.Role

	db.connection.Save(&user)
	return user
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Print(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
