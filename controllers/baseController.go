package controllers

import (
	"github.com/hubbdevelopers/hubb/db"
	"github.com/jinzhu/gorm"
)

var orm *gorm.DB

func Init() {
	orm = db.GetORM()
	//orm.LogMode(true)
}
