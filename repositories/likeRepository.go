package repositories

import (
	"github.com/hubbdevelopers/hubb/models"
	"github.com/jinzhu/gorm"
)

type LikeRepository interface {
	GetAll() *[]models.Like
	GetByUserID(userID int) *[]models.Like
	GetByPageID(pageID int) *[]models.Like
	Create(userID, pageID int) *models.Like
	Delete(userID, pageID int)
}

func NewLikeRepository(db *gorm.DB) LikeRepository {
	return dbLikeRepository{db: db}
}

type dbLikeRepository struct {
	db *gorm.DB
}

func (repo dbLikeRepository) GetAll() *[]models.Like {
	likes := []models.Like{}
	repo.db.Find(&likes)
	return &likes
}

func (repo dbLikeRepository) GetByUserID(userID int) *[]models.Like {
	likes := []models.Like{}
	repo.db.Where("user_id = ?", userID).Find(&likes)
	return &likes
}

func (repo dbLikeRepository) GetByPageID(pageID int) *[]models.Like {
	likes := []models.Like{}
	repo.db.Where("page_id = ?", pageID).Find(&likes)
	return &likes
}

func (repo dbLikeRepository) Create(userID, pageID int) *models.Like {
	like := models.Like{UserId: userID, PageId: pageID}
	repo.db.Create(&like)
	return &like
}

func (repo dbLikeRepository) Delete(userID, pageID int) {
	var like models.Like
	repo.db.Where("page_id = ?", pageID).Where("user_id = ?", userID).First(&like)
	repo.db.Delete(&like)
}
