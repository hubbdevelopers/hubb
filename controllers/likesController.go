package controllers

import (
	"fmt"
	"strconv"

	"github.com/hubbdevelopers/hubb/db"
	"github.com/hubbdevelopers/hubb/repositories"

	"github.com/gin-gonic/gin"
	"github.com/hubbdevelopers/hubb/models"
)

func GetLikes(c *gin.Context) {

	repo := repositories.NewLikeRepository(db.GetORM())
	var likes *[]models.Like

	pageID := c.Query("pageid")
	userID := c.Query("userid")
	if pageID != "" {
		pageIDInt, err := strconv.Atoi(pageID)
		if err != nil {
			fmt.Print(err)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		likes = repo.GetByPageID(pageIDInt)
	} else if userID != "" {
		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			fmt.Print(err)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		likes = repo.GetByUserID(userIDInt)
	} else {
		likes = repo.GetAll()
	}

	c.JSON(200, gin.H{
		"data": likes,
	})

}

type likeStruct struct {
	UserID int `json:"userId" binding:"required"`
	PageID int `json:"pageId" binding:"required"`
}

func CreateLike(c *gin.Context) {
	var json likeStruct
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Print(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	repo := repositories.NewLikeRepository(db.GetORM())
	likes := repo.Create(json.UserID, json.PageID)
	c.JSON(200, gin.H{
		"data": likes,
	})
}

func DeleteLike(c *gin.Context) {
	pageID, err := strconv.Atoi(c.Query("pageid"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userID, err := strconv.Atoi(c.Query("userid"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repo := repositories.NewLikeRepository(db.GetORM())
	repo.Delete(userID, pageID)
	c.JSON(200, gin.H{
		"data": "deleted",
	})
}
