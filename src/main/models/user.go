package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UID         string
	AccountId   string
	Name        string
	Pages       []Page `gorm:"polymorphic:Owner;"`
	Comments    []Comment
	Image       string
	Description string `sql:"type:text;"`
	Homepage    string
	Twitter     string
	Facebook    string
	Instagram   string
	Birthday    *time.Time `sql:"type:date"`
	Follow      Follow     `gorm:"polymorphic:Following;"`
}
