package repositories

import (
	"github.com/hubbdevelopers/hubb/models"
	"github.com/jinzhu/gorm"
)

type CommentRepository interface {
	GetByPageID(pageID int) *[]models.Comment
	GetByUserID(userID int) *[]models.Comment
	Create(userID, pageID int, text string) *models.Comment
	Delete(id int)
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return dbCommentRepository{db: db}
}

type dbCommentRepository struct {
	db *gorm.DB
}

func (repo dbCommentRepository) GetByUserID(userID int) *[]models.Comment {
	comments := []models.Comment{}
	repo.db.Where("user_id = ?", userID).Find(&comments)
	return &comments
}

func (repo dbCommentRepository) GetByPageID(pageID int) *[]models.Comment {
	comments := []models.Comment{}
	repo.db.Where("page_id = ?", pageID).Find(&comments)
	return &comments
}

func (repo dbCommentRepository) Create(userID, pageID int, text string) *models.Comment {
	comment := models.Comment{Text: text, UserId: userID, PageId: pageID}
	repo.db.Create(&comment)
	return &comment
}

func (repo dbCommentRepository) Delete(id int) {
	var comment models.Comment
	repo.db.First(&comment, id)
	repo.db.Delete(&comment)
}
