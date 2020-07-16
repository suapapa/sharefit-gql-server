package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"

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
	SharefitDB.LogMode(true)

	return nil
}

func Migrate() error {
	m := gormigrate.New(SharefitDB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "202007162111",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&User{}, &Card{}, &Center{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable(&User{}, &Card{}, &Center{}).Error
			},
		},
	})
	if err := m.Migrate(); err != nil {
		return fmt.Errorf("Could not migrate: %w", err)
	}
	log.Printf("Migration did run successfully")
	return nil
}
