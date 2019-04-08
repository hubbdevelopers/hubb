package controllers

import (
	"fmt"
	"strconv"

	"github.com/hubbdevelopers/hubb/db"
	"github.com/hubbdevelopers/hubb/repositories"

	"github.com/gin-gonic/gin"
)

func GetUserFollowings(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repo := repositories.NewFollowRepository(db.GetORM())
	followings := repo.GetFollowingsByUserID(userID)

	c.JSON(200, gin.H{
		"data": followings,
	})
}

func GetUserFollowers(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repo := repositories.NewFollowRepository(db.GetORM())
	followers := repo.GetFollowersByUserID(userID)

	c.JSON(200, gin.H{
		"data": followers,
	})
}

func FollowUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	followingID, err := strconv.Atoi(c.Param("followingid"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repo := repositories.NewFollowRepository(db.GetORM())
	follow := repo.CreateFollowUser(userID, followingID)

	c.JSON(200, gin.H{
		"data": follow,
	})
}

func UnfollowUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	followingID, err := strconv.Atoi(c.Param("followingid"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repo := repositories.NewFollowRepository(db.GetORM())
	repo.DeleteFollowUser(userID, followingID)

	c.JSON(200, gin.H{
		"data": "deleted",
	})
}

func FollowCommunity(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	followingID, err := strconv.Atoi(c.Param("followingid"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repo := repositories.NewFollowRepository(db.GetORM())
	follow := repo.CreateFollowCommunity(userID, followingID)

	c.JSON(200, gin.H{
		"data": follow,
	})
}

func UnfollowCommunity(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	followingID, err := strconv.Atoi(c.Param("followingid"))
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repo := repositories.NewFollowRepository(db.GetORM())
	repo.DeleteFollowCommunity(userID, followingID)

	c.JSON(200, gin.H{
		"data": "deleted",
	})
}
