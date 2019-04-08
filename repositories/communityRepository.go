package repositories

import (
	"github.com/hubbdevelopers/hubb/models"
	"github.com/jinzhu/gorm"
)

type CommunityRepository interface {
	GetAll() *[]models.Community
	GetByUserID(userID int) *[]models.Community
	GetByName(name string) *[]models.Community
	GetByID(id int) *models.Community
	Create(userID int, name string) *models.Community
	Delete(id int)
}

func NewCommunityRepository(db *gorm.DB) CommunityRepository {
	return dbCommunityRepository{db: db}
}

type dbCommunityRepository struct {
	db *gorm.DB
}

func (repo dbCommunityRepository) GetAll() *[]models.Community {
	communities := []models.Community{}
	repo.db.Find(&communities)
	return &communities
}

func (repo dbCommunityRepository) GetByUserID(userID int) *[]models.Community {
	communities := []models.Community{}
	communityMembers := []models.CommunityMember{}
	// TODO リレーションを使う
	repo.db.Where("user_id = ?", userID).Find(&communityMembers)

	var communityIds []int
	for _, value := range communityMembers {
		communityIds = append(communityIds, value.CommunityId)
	}

	repo.db.Where("id in (?)", communityIds).Find(&communities)
	return &communities
}

func (repo dbCommunityRepository) GetByName(communityName string) *[]models.Community {
	communities := []models.Community{}
	repo.db.Where("name = ?", communityName).Find(&communities)
	return &communities
}

func (repo dbCommunityRepository) GetByID(id int) *models.Community {
	community := models.Community{}
	repo.db.First(&community, id)
	return &community
}

func (repo dbCommunityRepository) Create(userID int, name string) *models.Community {
	tx := repo.db.Begin()
	community := models.Community{Name: name}
	if err := repo.db.Create(&community).Error; err != nil {
		tx.Rollback()
		return nil
	}

	communityMember := models.CommunityMember{CommunityId: int(community.ID), UserId: userID, IsOwner: true}
	if err := tx.Create(&communityMember).Error; err != nil {
		tx.Rollback()
		return nil
	}

	follow := models.Follow{UserId: userID, FollowingId: int(community.ID), FollowingType: "community"}
	if err := tx.Create(&follow).Error; err != nil {
		tx.Rollback()
		return nil
	}

	tx.Commit()

	return &community
}

func (repo dbCommunityRepository) Delete(id int) {
	var community models.Community
	repo.db.First(&community, id)
	repo.db.Delete(&community)
}
