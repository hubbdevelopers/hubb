package models

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	UserId int `gorm:"type:INT(10) UNSIGNED; not null; index"`
	PageId int `gorm:"type:INT(10) UNSIGNED; not null; index"`
	Text   string
}
