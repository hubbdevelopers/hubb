package db

import (
	"fmt"
	"log"
	"os"
	"time"

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

	numRetry := 5
	for i := 0; i <= numRetry; i++ {
		db, err = gorm.Open("mysql", user+":"+password+"@tcp("+host+":3306)/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			if i == numRetry {
				panic(fmt.Sprintf("Failed to connect to DB: %v", err))
			} else {
				log.Printf("An error occured while connecting to DB. Going to retry: %v\n", err)
				time.Sleep(3 * time.Second)
			}
		}
	}

	return db
}

func Close() {
	db.Close()
}

func GetORM() *gorm.DB {
	return db
}
