package repositories

import (
	"github.com/hubbdevelopers/hubb/db"
	"github.com/hubbdevelopers/hubb/models"
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

func NewPageRepository() PageRepository {
	return dbPageRepository{}
}

type dbPageRepository struct{}

func (dbPageRepository) GetAll() *[]models.Page {
	orm := db.GetORM()
	pages := []models.Page{}
	orm.Find(&pages)
	return &pages
}

func (dbPageRepository) GetRecentPages() *[]models.Page {
	orm := db.GetORM()
	pages := []models.Page{}
	orm.Order("created_at").Find(&pages)
	return &pages
}

// T0D0 リレーションを使う
func (dbPageRepository) GetTimeLine(userID int) *[]models.Page {
	orm := db.GetORM()
	var user models.User
	orm.First(&user, userID)

	var follows []models.Follow
	orm.Model(&user).Related(&follows)

	timeline := []models.Page{}
	for _, follow := range follows {

		if follow.FollowingType == "user" {
			var followingUser models.User
			orm.First(&followingUser, follow.FollowingId)

			var pages []models.Page
			orm.Model(&followingUser).Related(&pages, "Pages")

			timeline = append(timeline, pages...)

		} else if follow.FollowingType == "community" {
			var followingCommunity models.Community
			orm.First(&followingCommunity, follow.FollowingId)

			var pages []models.Page
			orm.Model(&followingCommunity).Related(&pages, "Pages")

			timeline = append(timeline, pages...)
		}

	}
	return &timeline
}

func (dbPageRepository) GetByUserID(userID int) *[]models.Page {
	orm := db.GetORM()
	pages := []models.Page{}
	orm.Where("owner_id = ?", userID).Where("owner_type = ?", "users").Find(&pages)
	return &pages
}

func (dbPageRepository) GetByCommunityID(communityID int) *[]models.Page {
	orm := db.GetORM()
	pages := []models.Page{}
	orm.Where("owner_id = ?", communityID).Where("owner_type = ?", "communities").Find(&pages)
	return &pages
}

func (dbPageRepository) GetByID(id int) *models.Page {
	orm := db.GetORM()
	page := models.Page{}
	orm.First(&page, id)
	return &page
}

func (dbPageRepository) Create(name string, ownerID int, ownerType string) *models.Page {
	orm := db.GetORM()
	page := models.Page{Name: name, OwnerId: ownerID, OwnerType: ownerType}
	orm.Create(&page)
	return &page
}

func (dbPageRepository) Update(id int, name string, content string, image string, draft bool) *models.Page {
	orm := db.GetORM()
	var page models.Page
	orm.First(&page, id)

	page.Name = name
	page.Content = content
	page.Image = image
	page.Draft = draft
	orm.Save(&page)
	return &page
}

func (dbPageRepository) Delete(id int) {
	orm := db.GetORM()
	var page models.Page
	orm.First(&page, id)
	orm.Delete(&page)
}
