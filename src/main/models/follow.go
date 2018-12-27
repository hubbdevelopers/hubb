package models

import (
	"github.com/jinzhu/gorm"
)

type Follow struct {
	gorm.Model
	UserId        int `gorm:"type:INT(10) UNSIGNED; not null"`
	FollowingId   int `gorm:"type:INT(10) UNSIGNED; not null"`
	FollowingType string
}
