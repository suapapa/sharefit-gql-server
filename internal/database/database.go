package database

import (
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	SharefitDB *gorm.DB
)

func InitDB() error {
	var err error
	SharefitDB, err = gorm.Open("sqlite3", "sharefit.db")
	if err != nil {
		return err
	}

	return nil
}
