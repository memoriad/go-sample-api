package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

// InitDb for connect database
func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("mysql", "smile:password@/sampledb?charset=utf8&parseTime=True&loc=Local")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}

	return db
}
