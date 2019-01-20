package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hubbdevelopers/hubb/models"
)

func GetLikes(c *gin.Context) {

	pageId := c.Query("pageid")
	userId := c.Query("userid")

	var likes []models.Like

	if pageId != "" {
		orm.Where("page_id = ?", pageId).Find(&likes)
	} else if userId != "" {
		orm.Where("user_id = ?", userId).Find(&likes)
	} else {
		likes = []models.Like{}
		orm.Find(&likes)
	}

	c.JSON(200, gin.H{
		"data": likes,
	})

}

type likeStruct struct {
	UserId int `json:"userId" binding:"required"`
	PageId int `json:"pageId" binding:"required"`
}

func CreateLike(c *gin.Context) {
	var json likeStruct
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Print(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var likeCount []models.Like
	orm.Where("page_id = ?", json.PageId).Where("user_id = ?", json.UserId).Find(&likeCount)

	if len(likeCount) > 0 {
		c.JSON(421, gin.H{
			"error": "already exists",
		})
	} else {
		like := models.Like{UserId: json.UserId, PageId: json.PageId}
		orm.Create(&like)
		c.JSON(200, gin.H{
			"data": like,
		})
	}
}

func DeleteLike(c *gin.Context) {
	pageId := c.Query("pageid")
	userId := c.Query("userid")

	var like models.Like
	orm.Where("page_id = ?", pageId).Where("user_id = ?", userId).First(&like)
	orm.Delete(&like)
	c.JSON(200, gin.H{
		"data": like,
	})
}
