package db

import (
	"main/models"
)

func Migrate() {
	db.AutoMigrate(&models.User{}, &models.Community{}, &models.CommunityMember{}, &models.Page{}, &models.Comment{}, &models.Like{}, &models.Follow{})
	db.Model(&models.Comment{}).AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT")
	db.Model(&models.Comment{}).AddForeignKey("page_id", "pages(id)", "CASCADE", "RESTRICT")
	db.Model(&models.Like{}).AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT")
	db.Model(&models.Like{}).AddForeignKey("page_id", "pages(id)", "CASCADE", "RESTRICT")
	db.Model(&models.CommunityMember{}).AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT")
	db.Model(&models.CommunityMember{}).AddForeignKey("community_id", "communities(id)", "CASCADE", "RESTRICT")
	db.Model(&models.Follow{}).AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT")
}
