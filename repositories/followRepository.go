package repositories

import (
	"github.com/hubbdevelopers/hubb/db"
	"github.com/hubbdevelopers/hubb/models"
	"github.com/jinzhu/gorm"
)

type FollowRepository interface {
	GetFollowingsByUserID(userID int) *[]models.Follow
	GetFollowersByUserID(userID int) *[]models.Follow
	GetFollowersByCommunityID(communityID int) *[]models.Follow
	CreateFollowUser(userID int, followingUserID int) *models.Follow
	CreateFollowCommunity(userID int, followingCommunityID int) *models.Follow
	DeleteFollowUser(userID int, followingUserID int)
	DeleteFollowCommunity(userID int, followingCommunityID int)
}

func NewFollowRepository(db *gorm.DB) FollowRepository {
	return dbFollowRepository{db: db}
}

type dbFollowRepository struct {
	db *gorm.DB
}

func (repo dbFollowRepository) GetFollowingsByUserID(userID int) *[]models.Follow {
	follows := []models.Follow{}
	repo.db.Where("user_id = ?", userID).Find(&follows)
	return &follows
}

func (repo dbFollowRepository) GetFollowersByUserID(userID int) *[]models.Follow {
	follows := []models.Follow{}
	repo.db.Where("following_id = ?", userID).Where("following_type = ?", "user").Find(&follows)
	return &follows
}

func (repo dbFollowRepository) GetFollowersByCommunityID(communityID int) *[]models.Follow {
	follows := []models.Follow{}
	repo.db.Where("following_id = ?", communityID).Where("following_type = ?", "user").Find(&follows)
	return &follows
}

func (repo dbFollowRepository) CreateFollowUser(userID int, followingUserID int) *models.Follow {
	follow := models.Follow{UserId: userID, FollowingId: followingUserID, FollowingType: "user"}
	repo.db.Create(&follow)
	return &follow
}

func (repo dbFollowRepository) CreateFollowCommunity(userID int, followingCommunityID int) *models.Follow {
	follow := models.Follow{UserId: userID, FollowingId: followingCommunityID, FollowingType: "community"}
	repo.db.Create(&follow)
	return &follow
}

func (dbFollowRepository) DeleteFollowUser(userID int, followingUserID int) {
	orm := db.GetORM()
	follow := models.Follow{}
	orm.Where("user_id = ?", userID).Where("following_id = ?", followingUserID).Where("following_type = ?", "user").First(&follow)
	orm.Delete(&follow)
}

func (repo dbFollowRepository) DeleteFollowCommunity(userID int, followingCommunityID int) {
	follow := models.Follow{}
	repo.db.Where("user_id = ?", userID).Where("following_id = ?", followingCommunityID).Where("following_type = ?", "community").First(&follow)
	repo.db.Delete(&follow)
}
