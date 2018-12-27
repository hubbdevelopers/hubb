package db

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Connect() *gorm.DB {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DATABASE")
	host := os.Getenv("MYSQL_HOST")

	db, err = gorm.Open("mysql", user+":"+password+"@tcp("+host+":3306)/"+dbname+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func Close() {
	db.Close()
}

func GetORM() *gorm.DB {
	return db
}
