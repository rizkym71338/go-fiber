package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Model *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_fiber"))
	if err != nil {
		panic(err)
	}
	Model = db
	Migration()
}
