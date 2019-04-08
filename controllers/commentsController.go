package controllers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hubbdevelopers/hubb/db"
	"github.com/hubbdevelopers/hubb/models"
	"github.com/hubbdevelopers/hubb/repositories"
)

func GetComments(c *gin.Context) {

	pageID := c.Query("pageid")
	userID := c.Query("userid")

	repo := repositories.NewCommentRepository(db.GetORM())
	var comments *[]models.Comment

	if pageID != "" {
		pageIDInt, err := strconv.Atoi(pageID)
		if err != nil {
			fmt.Print(err)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		comments = repo.GetByPageID(pageIDInt)
	} else if userID != "" {
		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			fmt.Print(err)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		comments = repo.GetByUserID(userIDInt)
	} else {
		c.JSON(400, gin.H{"error": errors.New("get comments need pageid or userid")})
		return
	}

	c.JSON(200, gin.H{
		"data": comments,
	})

}

type createCommentStruct struct {
	Text   string `json:"text" binding:"required"`
	UserID int    `json:"userId" binding:"required"`
	PageID int    `json:"pageId" binding:"required"`
}

func CreateComment(c *gin.Context) {
	var json createCommentStruct
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Print(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	repo := repositories.NewCommentRepository(db.GetORM())
	comment := repo.Create(json.UserID, json.PageID, json.Text)

	c.JSON(200, gin.H{
		"data": comment,
	})
}

func DeleteComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repo := repositories.NewCommentRepository(db.GetORM())
	repo.Delete(id)

	c.JSON(200, gin.H{
		"data": "deleted",
	})
}
