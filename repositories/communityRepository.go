package repositories

import (
	"github.com/hubbdevelopers/hubb/db"
	"github.com/hubbdevelopers/hubb/models"
)

type CommunityRepository interface {
	GetAll() *[]models.Community
	GetByUserID(userID int) *[]models.Community
	GetByName(name string) *[]models.Community
	GetByID(id int) *models.Community
	Create(userID int, name string) *models.Community
	Delete(id int)
}

func NewCommunityRepository() CommunityRepository {
	return dbCommunityRepository{}
}

type dbCommunityRepository struct{}

func (dbCommunityRepository) GetAll() *[]models.Community {
	orm := db.GetORM()
	communities := []models.Community{}
	orm.Find(&communities)
	return &communities
}

func (dbCommunityRepository) GetByUserID(userID int) *[]models.Community {
	orm := db.GetORM()
	communities := []models.Community{}
	communityMembers := []models.CommunityMember{}
	// TODO リレーションを使う
	orm.Where("user_id = ?", userID).Find(&communityMembers)

	var communityIds []int
	for _, value := range communityMembers {
		communityIds = append(communityIds, value.CommunityId)
	}

	orm.Where("id in (?)", communityIds).Find(&communities)
	return &communities
}

func (dbCommunityRepository) GetByName(communityName string) *[]models.Community {
	orm := db.GetORM()
	communities := []models.Community{}
	orm.Where("name = ?", communityName).Find(&communities)
	return &communities
}

func (dbCommunityRepository) GetByID(id int) *models.Community {
	orm := db.GetORM()
	community := models.Community{}
	orm.First(&community, id)
	return &community
}

func (dbCommunityRepository) Create(userID int, name string) *models.Community {
	orm := db.GetORM()
	tx := orm.Begin()
	community := models.Community{Name: name}
	if err := orm.Create(&community).Error; err != nil {
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

func (dbCommunityRepository) Delete(id int) {
	orm := db.GetORM()
	var community models.Community
	orm.First(&community, id)
	orm.Delete(&community)
}
