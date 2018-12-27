package controllers

import (
	"main/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserFollowings(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	var follows []models.Follow
	orm.Where("user_id = ?", userId).Find(&follows)
	c.JSON(200, gin.H{
		"data": follows,
	})
}

func GetUserFollowers(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	var follows []models.Follow
	orm.Where("following_id = ?", userId).Find(&follows)
	c.JSON(200, gin.H{
		"data": follows,
	})
}

func FollowUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	followingId, _ := strconv.Atoi(c.Param("followingid"))

	follow := models.Follow{UserId: userId, FollowingId: followingId, FollowingType: "user"}
	orm.Create(&follow)
	c.JSON(200, gin.H{
		"data": follow,
	})
}

func UnfollowUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	followingId, _ := strconv.Atoi(c.Param("followingid"))

	var follow models.Follow
	orm.Where("user_id = ?", userId).Where("following_id = ?", followingId).Where("following_type = ?", "user").First(&follow)
	orm.Delete(&follow)
	c.JSON(200, gin.H{
		"data": follow,
	})
}

func FollowCommunity(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	followingId, _ := strconv.Atoi(c.Param("followingid"))

	follow := models.Follow{UserId: userId, FollowingId: followingId, FollowingType: "community"}
	orm.Create(&follow)
	c.JSON(200, gin.H{
		"data": follow,
	})
}

func UnfollowCommunity(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	followingId, _ := strconv.Atoi(c.Param("followingid"))

	var follow models.Follow
	orm.Where("user_id = ?", userId).Where("following_id = ?", followingId).Where("following_type = ?", "community").First(&follow)
	orm.Delete(&follow)
	c.JSON(200, gin.H{
		"data": follow,
	})
}
