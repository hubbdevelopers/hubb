package controllers

import (
	"fmt"
	"main/models"

	"github.com/gin-gonic/gin"
)

func GetComments(c *gin.Context) {

	pageId := c.Query("pageid")

	var comments []models.Comment

	if pageId != "" {
		orm.Where("page_id = ?", pageId).Find(&comments)

	} else {
		comments = []models.Comment{}
		orm.Find(&comments)
	}

	c.JSON(200, gin.H{
		"data": comments,
	})

}

type createCommentStruct struct {
	Text   string `json:"text" binding:"required"`
	UserId int    `json:"userId" binding:"required"`
	PageId int    `json:"pageId" binding:"required"`
}

func CreateComment(c *gin.Context) {
	var json createCommentStruct
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Print(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	comment := models.Comment{Text: json.Text, UserId: json.UserId, PageId: json.PageId}
	orm.Create(&comment)
	c.JSON(200, gin.H{
		"data": comment,
	})
}

func DeleteComment(c *gin.Context) {
	id := c.Param("id")

	var comment models.Comment
	orm.First(&comment, id)
	orm.Delete(&comment)

	c.JSON(200, gin.H{
		"data": comment,
	})
}
