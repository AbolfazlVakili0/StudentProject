package database

import (
	"github.com/glebarez/sqlite" 
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	// استفاده از درایور جدید
	DB, err = gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		panic("cannot open database: " + err.Error())
	}
}