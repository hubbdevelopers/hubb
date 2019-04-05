package repositories

import (
	"time"

	"github.com/hubbdevelopers/hubb/db"
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

func NewUserRepository() UserRepository {
	return dbUserRepository{}
}

type dbUserRepository struct{}

func (dbUserRepository) GetAll() *[]models.User {
	orm := db.GetORM()
	users := []models.User{}
	orm.Find(&users)
	return &users
}

func (dbUserRepository) GetByID(id int) *models.User {
	orm := db.GetORM()
	user := models.User{}
	orm.First(&user, id)
	return &user
}

func (dbUserRepository) GetByUID(uid string) *models.User {
	orm := db.GetORM()
	user := models.User{}
	orm.Where("uid = ?", uid).First(&user)
	return &user
}
func (dbUserRepository) GetByAccountID(accountID string) *models.User {
	orm := db.GetORM()
	user := models.User{}
	orm.Where("account_id = ?", accountID).First(&user)
	return &user
}

func (dbUserRepository) Create(uid string) *models.User {
	orm := db.GetORM()
	user := models.User{UID: uid}
	orm.Create(&user)
	return &user
}

func (dbUserRepository) Initilize(id int, accountID string, name string) *models.User {
	orm := db.GetORM()
	user := models.User{}
	orm.First(&user, id)

	user.AccountId = accountID
	user.Name = name

	orm.Save(&user)
	return &user
}

func (dbUserRepository) UpdateImage(id int, image string) *models.User {
	orm := db.GetORM()
	user := models.User{}
	orm.First(&user, id)

	user.Image = image

	orm.Save(&user)
	return &user
}

func (dbUserRepository) UpdateProfile(id int, name string, description string, homepage string, facebook string, twitter string, instagram string, birthday string) *models.User {
	orm := db.GetORM()
	user := models.User{}
	orm.First(&user, id)

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

	orm.Save(&user)
	return &user
}
