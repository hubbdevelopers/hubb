package repositories

import (
	"github.com/hubbdevelopers/hubb/db"
	"github.com/hubbdevelopers/hubb/models"
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

func NewFollowRepository() FollowRepository {
	return dbFollowRepository{}
}

type dbFollowRepository struct{}

func (dbFollowRepository) GetFollowingsByUserID(userID int) *[]models.Follow {
	orm := db.GetORM()
	follows := []models.Follow{}
	orm.Where("user_id = ?", userID).Find(&follows)
	return &follows
}

func (dbFollowRepository) GetFollowersByUserID(userID int) *[]models.Follow {
	orm := db.GetORM()
	follows := []models.Follow{}
	orm.Where("following_id = ?", userID).Where("following_type = ?", "user").Find(&follows)
	return &follows
}

func (dbFollowRepository) GetFollowersByCommunityID(communityID int) *[]models.Follow {
	orm := db.GetORM()
	follows := []models.Follow{}
	orm.Where("following_id = ?", communityID).Where("following_type = ?", "user").Find(&follows)
	return &follows
}

func (dbFollowRepository) CreateFollowUser(userID int, followingUserID int) *models.Follow {
	orm := db.GetORM()
	follow := models.Follow{UserId: userID, FollowingId: followingUserID, FollowingType: "user"}
	orm.Create(&follow)
	return &follow
}

func (dbFollowRepository) CreateFollowCommunity(userID int, followingCommunityID int) *models.Follow {
	orm := db.GetORM()
	follow := models.Follow{UserId: userID, FollowingId: followingCommunityID, FollowingType: "community"}
	orm.Create(&follow)
	return &follow
}

func (dbFollowRepository) DeleteFollowUser(userID int, followingUserID int) {
	orm := db.GetORM()
	follow := models.Follow{}
	orm.Where("user_id = ?", userID).Where("following_id = ?", followingUserID).Where("following_type = ?", "user").First(&follow)
	orm.Delete(&follow)
}

func (dbFollowRepository) DeleteFollowCommunity(userID int, followingCommunityID int) {
	orm := db.GetORM()
	follow := models.Follow{}
	orm.Where("user_id = ?", userID).Where("following_id = ?", followingCommunityID).Where("following_type = ?", "community").First(&follow)
	orm.Delete(&follow)
}
