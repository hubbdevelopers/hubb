package repositories

import (
	"github.com/hubbdevelopers/hubb/models"
	"github.com/jinzhu/gorm"
)

type PageRepository interface {
	GetAll() *[]models.Page
	GetRecentPages() *[]models.Page
	GetTimeLine(userID int) *[]models.Page
	GetByUserID(userID int) *[]models.Page
	GetByCommunityID(communityID int) *[]models.Page
	GetByID(id int) *models.Page
	Create(name string, ownerID int, ownerType string) *models.Page
	Update(id int, name string, content string, image string, draft bool) *models.Page
	Delete(id int)
}

func NewPageRepository(db *gorm.DB) PageRepository {
	return dbPageRepository{db: db}
}

type dbPageRepository struct {
	db *gorm.DB
}

func (repo dbPageRepository) GetAll() *[]models.Page {
	pages := []models.Page{}
	repo.db.Find(&pages)
	return &pages
}

func (repo dbPageRepository) GetRecentPages() *[]models.Page {
	pages := []models.Page{}
	repo.db.Order("created_at").Find(&pages)
	return &pages
}

// T0D0 リレーションを使う
func (repo dbPageRepository) GetTimeLine(userID int) *[]models.Page {
	var user models.User
	repo.db.First(&user, userID)

	var follows []models.Follow
	repo.db.Model(&user).Related(&follows)

	timeline := []models.Page{}
	for _, follow := range follows {

		if follow.FollowingType == "user" {
			var followingUser models.User
			repo.db.First(&followingUser, follow.FollowingId)

			var pages []models.Page
			repo.db.Model(&followingUser).Related(&pages, "Pages")

			timeline = append(timeline, pages...)

		} else if follow.FollowingType == "community" {
			var followingCommunity models.Community
			repo.db.First(&followingCommunity, follow.FollowingId)

			var pages []models.Page
			repo.db.Model(&followingCommunity).Related(&pages, "Pages")

			timeline = append(timeline, pages...)
		}

	}
	return &timeline
}

func (repo dbPageRepository) GetByUserID(userID int) *[]models.Page {
	pages := []models.Page{}
	repo.db.Where("owner_id = ?", userID).Where("owner_type = ?", "users").Find(&pages)
	return &pages
}

func (repo dbPageRepository) GetByCommunityID(communityID int) *[]models.Page {
	pages := []models.Page{}
	repo.db.Where("owner_id = ?", communityID).Where("owner_type = ?", "communities").Find(&pages)
	return &pages
}

func (repo dbPageRepository) GetByID(id int) *models.Page {
	page := models.Page{}
	repo.db.First(&page, id)
	return &page
}

func (repo dbPageRepository) Create(name string, ownerID int, ownerType string) *models.Page {
	page := models.Page{Name: name, OwnerId: ownerID, OwnerType: ownerType}
	repo.db.Create(&page)
	return &page
}

func (repo dbPageRepository) Update(id int, name string, content string, image string, draft bool) *models.Page {
	var page models.Page
	repo.db.First(&page, id)

	page.Name = name
	page.Content = content
	page.Image = image
	page.Draft = draft
	repo.db.Save(&page)
	return &page
}

func (repo dbPageRepository) Delete(id int) {
	var page models.Page
	repo.db.First(&page, id)
	repo.db.Delete(&page)
}
