package models

import (
	"github.com/jinzhu/gorm"
)

type Community struct {
	gorm.Model
	Name   string
	Pages  []Page `gorm:"polymorphic:Owner;"`
	Follow Follow `gorm:"polymorphic:Following;"`
}
