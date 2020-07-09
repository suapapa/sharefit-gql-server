package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Card struct {
	gorm.Model
	Training string
	CurrCnt  int
	TotalCnt int
	Expiry   time.Time
	Users    []User
	CenterID uint
}

// TableName sets the insert table name for this struct type
func (c *Card) TableName() string {
	return "card"
}
