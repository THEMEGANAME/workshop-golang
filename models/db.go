package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func InitDB(){
	var err error
	db,err = gorm.Open("sqlite3","user.db")
	if err != nil {
		panic("failed connecting db")
	}
}


