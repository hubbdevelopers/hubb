package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hubbdevelopers/hubb/models"
)

func GetPages(c *gin.Context) {

	userId := c.Query("userid")
	communityId := c.Query("communityid")

	var pages []models.Page
	if userId != "" {
		orm.Where("owner_id = ?", userId).Where("owner_type = ?", "users").Find(&pages)
	} else if communityId != "" {
		orm.Where("owner_id = ?", communityId).Where("owner_type = ?", "communities").Find(&pages)
	} else {
		orm.Find(&pages)
	}

	c.JSON(200, gin.H{
		"data": pages,
	})
}

func GetRecentPages(c *gin.Context) {

	var pages []models.Page
	orm.Order("created_at").Find(&pages)

	c.JSON(200, gin.H{
		"data": pages,
	})
}

func GetPage(c *gin.Context) {

	id := c.Param("id")
	var page models.Page
	orm.First(&page, id)
	c.JSON(200, gin.H{
		"data": page,
	})
}

type newPage struct {
	Name      string `json:"name" binding:"required"`
	OwnerId   int    `json:"ownerId" binding:"required"`
	OwnerType string `json:"ownerType" binding:"required"`
}

func CreatePage(c *gin.Context) {
	var json newPage
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Print(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	page := models.Page{Name: json.Name, OwnerId: json.OwnerId, OwnerType: json.OwnerType}
	orm.Create(&page)
	c.JSON(200, gin.H{
		"data": page,
	})
}

func DeletePage(c *gin.Context) {
	id := c.Param("id")

	var page models.Page
	orm.First(&page, id)
	orm.Delete(&page)

	c.JSON(200, gin.H{
		"data": page,
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

	id := c.Param("id")
	var page models.Page
	orm.First(&page, id)

	page.Name = json.Name
	page.Content = json.Content
	page.Image = json.Image
	page.Draft = json.Draft
	orm.Save(&page)

	c.JSON(200, gin.H{
		"data": page,
	})
}
