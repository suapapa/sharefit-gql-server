package database

import "github.com/jinzhu/gorm"

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
