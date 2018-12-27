package models

import (
	"github.com/jinzhu/gorm"
)

type Page struct {
	gorm.Model
	Name      string
	OwnerId   int `gorm:"type:INT(10) UNSIGNED; not null"`
	OwnerType string
	Content   string `sql:"type:text;"`
	Image     string
	Draft     bool `gorm:"default:true"`
}
