package controllers

import (
	"fmt"
	"main/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	accountId := c.Query("accountid")
	uId := c.Query("uid")

	if accountId != "" {
		var user models.User
		orm.Where("account_id = ?", accountId).First(&user)
		c.JSON(200, gin.H{
			"data": user,
		})
	} else if uId != "" {
		var user models.User
		orm.Where("uid = ?", uId).First(&user)
		c.JSON(200, gin.H{
			"data": user,
		})
	} else {
		users := []models.User{}
		orm.Find(&users)

		c.JSON(200, gin.H{
			"data": users,
		})
	}
}

func GetUser(c *gin.Context) {

	id := c.Param("id")
	var user models.User
	orm.First(&user, id)

	c.JSON(200, gin.H{
		"data": user,
	})
}

func GetTimeline(c *gin.Context) {

	id := c.Param("id")
	fmt.Println(id)

	var user models.User
	orm.First(&user, id)

	fmt.Println(user)
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
	fmt.Println(timeline)

	c.JSON(200, gin.H{
		"data": timeline,
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

	user := models.User{UID: json.Uid}
	orm.Create(&user)
	c.JSON(200, gin.H{
		"data": user,
	})
}

type initUserStruct struct {
	AccountId string `json:"account_id"binding:"required"`
	Name      string `json:"name" binding:"required"`
}

func InitializeUser(c *gin.Context) {
	var json initUserStruct
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Print(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	var user models.User
	orm.First(&user, id)

	user.AccountId = json.AccountId
	user.Name = json.Name

	orm.Save(&user)

	c.JSON(200, gin.H{
		"data": user,
	})
}

type updateImageStruct struct {
	Image string `json:"image"binding:"required"`
}

func UpdateImage(c *gin.Context) {
	var json updateImageStruct
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Print(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	var user models.User
	orm.First(&user, id)

	user.Image = json.Image

	orm.Save(&user)

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

	id := c.Param("id")
	var user models.User
	orm.First(&user, id)

	if json.Name != "" {
		user.Name = json.Name
	}

	if json.Description != "" {
		user.Description = json.Description
	}

	if json.Homepage != "" {
		user.Homepage = json.Homepage
	}

	if json.Facebook != "" {
		user.Facebook = json.Facebook
	}

	if json.Twitter != "" {
		user.Twitter = json.Twitter
	}

	if json.Instagram != "" {
		user.Instagram = json.Instagram
	}

	if json.Birthday != "" {
		layout := "2006-01-02"
		t, _ := time.Parse(layout, json.Birthday)
		user.Birthday = &t
	}

	orm.Save(&user)

	c.JSON(200, gin.H{
		"data": user,
	})
}
