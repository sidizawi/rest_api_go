package db

import (
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

type Book struct {
	gorm.Model
	Title  string
	Author string
	Rating int
}
