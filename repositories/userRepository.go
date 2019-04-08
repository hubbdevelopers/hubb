package repositories

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/hubbdevelopers/hubb/models"
)

type UserRepository interface {
	GetAll() *[]models.User
	GetByID(id int) *models.User
	GetByUID(uid string) *models.User
	GetByAccountID(accountID string) *models.User
	Create(uid string) *models.User
	Initilize(id int, accountID string, name string) *models.User
	UpdateImage(id int, image string) *models.User
	UpdateProfile(id int, name string, description string, homepage string, facebook string, twitter string, instagram string, birthday string) *models.User
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return dbUserRepository{db: db}
}

type dbUserRepository struct {
	db *gorm.DB
}

func (repo dbUserRepository) GetAll() *[]models.User {
	users := []models.User{}
	repo.db.Find(&users)
	return &users
}

func (repo dbUserRepository) GetByID(id int) *models.User {
	user := models.User{}
	repo.db.First(&user, id)
	return &user
}

func (repo dbUserRepository) GetByUID(uid string) *models.User {
	user := models.User{}
	repo.db.Where("uid = ?", uid).First(&user)
	return &user
}

func (repo dbUserRepository) GetByAccountID(accountID string) *models.User {
	user := models.User{}
	repo.db.Where("account_id = ?", accountID).First(&user)
	return &user
}

func (repo dbUserRepository) Create(uid string) *models.User {
	user := models.User{UID: uid}
	repo.db.Create(&user)
	return &user
}

func (repo dbUserRepository) Initilize(id int, accountID string, name string) *models.User {
	user := models.User{}
	repo.db.First(&user, id)

	user.AccountId = accountID
	user.Name = name

	repo.db.Save(&user)
	return &user
}

func (repo dbUserRepository) UpdateImage(id int, image string) *models.User {
	user := models.User{}
	repo.db.First(&user, id)

	user.Image = image

	repo.db.Save(&user)
	return &user
}

func (repo dbUserRepository) UpdateProfile(id int, name string, description string, homepage string, facebook string, twitter string, instagram string, birthday string) *models.User {
	user := models.User{}
	repo.db.First(&user, id)

	if name != "" {
		user.Name = name
	}

	if description != "" {
		user.Description = description
	}

	if homepage != "" {
		user.Homepage = homepage
	}

	if facebook != "" {
		user.Facebook = facebook
	}

	if twitter != "" {
		user.Twitter = twitter
	}

	if instagram != "" {
		user.Instagram = instagram
	}

	if birthday != "" {
		layout := "2006-01-02"
		t, _ := time.Parse(layout, birthday)
		user.Birthday = &t
	}

	repo.db.Save(&user)
	return &user
}
