package repositories

import (
	"github.com/hubbdevelopers/hubb/db"
	"github.com/hubbdevelopers/hubb/models"
)

type LikeRepository interface {
	GetAll() *[]models.Like
	GetByUserID(userID int) *[]models.Like
	GetByPageID(pageID int) *[]models.Like
	Create(userID, pageID int) *models.Like
	Delete(userID, pageID int)
}

func NewLikeRepository() LikeRepository {
	return dbLikeRepository{}
}

type dbLikeRepository struct{}

func (dbLikeRepository) GetAll() *[]models.Like {
	orm := db.GetORM()
	likes := []models.Like{}
	orm.Find(&likes)
	return &likes
}

func (dbLikeRepository) GetByUserID(userID int) *[]models.Like {
	orm := db.GetORM()
	likes := []models.Like{}
	orm.Where("user_id = ?", userID).Find(&likes)
	return &likes
}

func (dbLikeRepository) GetByPageID(pageID int) *[]models.Like {
	orm := db.GetORM()
	likes := []models.Like{}
	orm.Where("page_id = ?", pageID).Find(&likes)
	return &likes
}

func (dbLikeRepository) Create(userID, pageID int) *models.Like {
	orm := db.GetORM()
	like := models.Like{UserId: userID, PageId: pageID}
	orm.Create(&like)
	return &like
}

func (dbLikeRepository) Delete(userID, pageID int) {
	orm := db.GetORM()
	var like models.Like
	orm.Where("page_id = ?", pageID).Where("user_id = ?", userID).First(&like)
	orm.Delete(&like)
}
