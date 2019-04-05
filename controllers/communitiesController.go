package controllers

import (
	"fmt"
	"strconv"

	"github.com/hubbdevelopers/hubb/repositories"

	"github.com/gin-gonic/gin"
	"github.com/hubbdevelopers/hubb/models"
)

func GetCommunities(c *gin.Context) {

	repo := repositories.NewCommunityRepository()
	userID := c.Query("userid")
	name := c.Query("name")

	var communities *[]models.Community
	if userID != "" {
		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			fmt.Print(err)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		communities = repo.GetByUserID(userIDInt)
	} else if name != "" {
		communities = repo.GetByName(name)
	} else {
		communities = repo.GetAll()
	}

	c.JSON(200, gin.H{
		"data": communities,
	})
}

func GetCommunity(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	repo := repositories.NewCommunityRepository()
	community := repo.GetByID(id)

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
	repo := repositories.NewCommunityRepository()
	community := repo.Create(json.UserID, json.Name)

	c.JSON(200, gin.H{
		"data": community,
	})
}

func DeleteCommunity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repo := repositories.NewCommunityRepository()
	repo.Delete(id)

	c.JSON(200, gin.H{
		"data": "deleted",
	})
}
