package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hubbdevelopers/hubb/db"
	"github.com/hubbdevelopers/hubb/repositories"
)

func GetUsers(c *gin.Context) {

	accountId := c.Query("accountid")
	uId := c.Query("uid")

	repo := repositories.NewUserRepository(db.GetORM())

	if accountId != "" {
		user := repo.GetByAccountID(accountId)
		c.JSON(200, gin.H{
			"data": user,
		})
	} else if uId != "" {
		user := repo.GetByUID(uId)
		c.JSON(200, gin.H{
			"data": user,
		})
	} else {
		users := repo.GetAll()
		c.JSON(200, gin.H{
			"data": users,
		})
	}
}

func GetUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repo := repositories.NewUserRepository(db.GetORM())
	user := repo.GetByID(id)

	c.JSON(200, gin.H{
		"data": user,
	})
}

type signupStruct struct {
	Email    string `form:"email" json:"email" xml:"email" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	Uid      string `form:"uid" json:"uid" xml:"uid" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var json signupStruct
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Print(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	repo := repositories.NewUserRepository(db.GetORM())
	user := repo.Create(json.Uid)

	c.JSON(200, gin.H{
		"data": user,
	})
}

type initUserStruct struct {
	AccountID string `json:"account_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
}

func InitializeUser(c *gin.Context) {
	var json initUserStruct
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Print(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repo := repositories.NewUserRepository(db.GetORM())
	user := repo.Initilize(id, json.AccountID, json.Name)

	c.JSON(200, gin.H{
		"data": user,
	})
}

type updateImageStruct struct {
	Image string `json:"image" binding:"required"`
}

func UpdateImage(c *gin.Context) {
	var json updateImageStruct
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Print(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repo := repositories.NewUserRepository(db.GetORM())
	user := repo.UpdateImage(id, json.Image)

	c.JSON(200, gin.H{
		"data": user,
	})
}

type updateProfileStruct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Facebook    string `json:"facebook"`
	Twitter     string `json:"twitter"`
	Instagram   string `json:"instagram"`
	Birthday    string `json:"birthday"`
}

func UpdateProfile(c *gin.Context) {
	var json updateProfileStruct
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Print(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	repo := repositories.NewUserRepository(db.GetORM())
	user := repo.UpdateProfile(id, json.Name, json.Description, json.Homepage, json.Facebook, json.Twitter, json.Instagram, json.Birthday)

	c.JSON(200, gin.H{
		"data": user,
	})
}
