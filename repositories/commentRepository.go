package repositories

import (
	"github.com/hubbdevelopers/hubb/db"
	"github.com/hubbdevelopers/hubb/models"
)

type CommentRepository interface {
	GetByPageID(pageID int) *[]models.Comment
	GetByUserID(userID int) *[]models.Comment
	Create(userID, pageID int, text string) *models.Comment
	Delete(id int)
}

func NewCommentRepository() CommentRepository {
	return dbCommentRepository{}
}

type dbCommentRepository struct{}

func (dbCommentRepository) GetByUserID(userID int) *[]models.Comment {
	orm := db.GetORM()
	comments := []models.Comment{}
	orm.Where("user_id = ?", userID).Find(&comments)
	return &comments
}

func (dbCommentRepository) GetByPageID(pageID int) *[]models.Comment {
	orm := db.GetORM()
	comments := []models.Comment{}
	orm.Where("pager_id = ?", pageID).Find(&comments)
	return &comments
}

func (dbCommentRepository) Create(userID, pageID int, text string) *models.Comment {
	orm := db.GetORM()
	comment := models.Comment{Text: text, UserId: userID, PageId: pageID}
	orm.Create(&comment)
	return &comment
}

func (dbCommentRepository) Delete(id int) {
	orm := db.GetORM()
	var comment models.Comment
	orm.First(&comment, id)
	orm.Delete(&comment)
}
