package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name        string
	Password    string
	PhoneNumber string
	CardID      uint
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "user"
}

func CreateNewUser(name, password, phoneNumber string, cardID uint) (*User, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:        name, // sholud be unique
		PhoneNumber: phoneNumber,
		CardID:      cardID,
		Password:    hashedPassword,
	}, nil
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GetUserByUsername check if a user exists in database by given username
func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := SharefitDB.Where("name = ?", username).First(&user).Error; err != nil {
		return nil, fmt.Errorf("no such a user, %s", username)
	}
	return &user, nil
}
