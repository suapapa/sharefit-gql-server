package database

import "github.com/jinzhu/gorm"

type Center struct {
	gorm.Model
	Address     string
	Name        string
	PhoneNumber string
	Cards       []Card
}

// TableName sets the insert table name for this struct type
func (c *Center) TableName() string {
	return "center"
}
