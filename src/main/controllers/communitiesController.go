package controllers

import (
	"fmt"
	"main/models"

	"github.com/gin-gonic/gin"
)

func GetCommunities(c *gin.Context) {

	userId := c.Query("userid")
	name := c.Query("name")

	var communities []models.Community
	if userId != "" {
		// TODO リレーションを使う
		var communityMembers []models.CommunityMember
		orm.Where("user_id = ?", userId).Find(&communityMembers)

		var communityIds []int
		for _, value := range communityMembers {
			communityIds = append(communityIds, value.CommunityId)
		}

		orm.Where("id in (?)", communityIds).Find(&communities)
	} else if name != "" {
		orm.Where("name = ?", name).Find(&communities)
	} else {
		orm.Find(&communities)
	}

	c.JSON(200, gin.H{
		"data": communities,
	})
}

func GetCommunity(c *gin.Context) {

	id := c.Param("id")
	var community models.Community
	orm.First(&community, id)

	c.JSON(200, gin.H{
		"data": community,
	})
}

type createCommunityStruct struct {
	Name   string `json:"name" binding:"required"`
	UserID int    `json:"userId" binding:"required"`
}

func CreateCommunity(c *gin.Context) {
	var json createCommunityStruct
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Print(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	tx := orm.Begin()
	community := models.Community{Name: json.Name}
	if err := orm.Create(&community).Error; err != nil {
		tx.Rollback()
		return
	}

	communityMember := models.CommunityMember{CommunityId: int(community.ID), UserId: json.UserID, IsOwner: true}
	if err := tx.Create(&communityMember).Error; err != nil {
		tx.Rollback()
		return
	}

	follow := models.Follow{UserId: json.UserID, FollowingId: int(community.ID), FollowingType: "community"}
	orm.Create(&follow)

	tx.Commit()
	c.JSON(200, gin.H{
		"data": community,
	})
}

func DeleteCommunity(c *gin.Context) {
	id := c.Param("id")

	var community models.Community
	orm.First(&community, id)
	orm.Delete(&community)

	c.JSON(200, gin.H{
		"data": community,
	})
}

// type initUser struct {
// 	AccountId  string `json:"account_id"binding:"required"`
// 	Name       string `json:"name" binding:"required"`
// }

// func InitializeUser(c *gin.Context) {
// 	var json initUser
// 	if err := c.ShouldBindJSON(&json); err != nil {
// 		fmt.Print(err)
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	id := c.Param("id")
// 	var user models.User
// 	orm.First(&user, id)

// 	user.AccountId = json.AccountId
// 	user.Name = json.Name

// 	orm.Save(&user)

// 	c.JSON(200, gin.H{
// 		"data": user,
// 	})
// }
