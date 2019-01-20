package models

import (
	"github.com/jinzhu/gorm"
)

type CommunityMember struct {
	gorm.Model
	UserId      int `gorm:"type:INT(10) UNSIGNED; not null; index"`
	CommunityId int `gorm:"type:INT(10) UNSIGNED; not null; index"`
	IsOwner     bool
}
