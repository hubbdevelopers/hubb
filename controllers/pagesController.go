package controllers

import (
	"fmt"

	"github.com/hubbdevelopers/hubb/db"
	"github.com/hubbdevelopers/hubb/repositories"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hubbdevelopers/hubb/models"
)

func GetPages(c *gin.Context) {

	userID := c.Query("userid")
	communityID := c.Query("communityid")

	repo := repositories.NewPageRepository(db.GetORM())

	var pages *[]models.Page
	if userID != "" {
		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			fmt.Print(err)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		pages = repo.GetByUserID(userIDInt)
	} else if communityID != "" {
		communityIDInt, err := strconv.Atoi(communityID)
		if err != nil {
			fmt.Print(err)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		pages = repo.GetByCommunityID(communityIDInt)
	} else {
		pages = repo.GetAll()
	}

	c.JSON(200, gin.H{
		"data": pages,
	})
}

func GetRecentPages(c *gin.Context) {

	repo := repositories.NewPageRepository(db.GetORM())
	pages := repo.GetRecentPages()

	c.JSON(200, gin.H{
		"data": pages,
	})
}

func GetTimeline(c *gin.Context) {

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repo := repositories.NewPageRepository(db.GetORM())
	pages := repo.GetTimeLine(userID)

	c.JSON(200, gin.H{
		"data": pages,
	})
}

func GetPage(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	repo := repositories.NewPageRepository(db.GetORM())
	page := repo.GetByID(id)
	c.JSON(200, gin.H{
		"data": page,
	})
}

type newPage struct {
	Name      string `json:"name" binding:"required"`
	OwnerID   int    `json:"ownerId" binding:"required"`
	OwnerType string `json:"ownerType" binding:"required"`
}

func CreatePage(c *gin.Context) {
	var json newPage
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Print(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	repo := repositories.NewPageRepository(db.GetORM())
	page := repo.Create(json.Name, json.OwnerID, json.OwnerType)
	c.JSON(200, gin.H{
		"data": page,
	})
}

func DeletePage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repo := repositories.NewPageRepository(db.GetORM())
	repo.Delete(id)

	c.JSON(200, gin.H{
		"data": "deleted",
	})
}

type updatePageStruct struct {
	Draft   bool   `json:"draft"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

func UpdatePage(c *gin.Context) {
	var json updatePageStruct
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

	repo := repositories.NewPageRepository(db.GetORM())
	page := repo.Update(id, json.Name, json.Content, json.Image, json.Draft)

	c.JSON(200, gin.H{
		"data": page,
	})
}
