package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hubbdevelopers/hubb/controllers"
	"github.com/hubbdevelopers/hubb/db"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	db.Connect()
	defer db.Close()

	db.Migrate()

	controllers.Init()

	r := setupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))

	// TODO この API に認証をかけると Login と Signup ができない
	r.GET("/users", controllers.GetUsers)

	authenticated := r.Group("/")
	authenticated.Use(controllers.Authenticate())
	{
		authenticated.GET("/users/:id", controllers.GetUser)
		authenticated.POST("/users", controllers.CreateUser)
		authenticated.PUT("/users/:id/init", controllers.InitializeUser)
		authenticated.PUT("/users/:id/images", controllers.UpdateImage)
		authenticated.PUT("/users/:id/profile", controllers.UpdateProfile)

		authenticated.GET("/users/:id/timeline", controllers.GetTimeline)

		authenticated.GET("/users/:id/followings", controllers.GetUserFollowings)
		authenticated.GET("/users/:id/followers", controllers.GetUserFollowers)
		//authenticated.GET("/communities/:id/followers", controllers.GetCommunityFollowers)
		authenticated.POST("/users/:id/follow/users/:followingid", controllers.FollowUser)
		authenticated.DELETE("/users/:id/follow/users/:followingid", controllers.UnfollowUser)
		authenticated.POST("/users/:id/follow/communities/:followingid", controllers.FollowCommunity)
		authenticated.DELETE("/users/:id/follow/communities/:followingid", controllers.UnfollowCommunity)

		authenticated.GET("/pages", controllers.GetPages)
		authenticated.GET("/pages/:id", controllers.GetPage)
		authenticated.POST("/pages", controllers.CreatePage)
		authenticated.DELETE("/pages/:id", controllers.DeletePage)
		authenticated.PUT("/pages/:id", controllers.UpdatePage)

		authenticated.GET("/comments", controllers.GetComments)
		authenticated.POST("/comments", controllers.CreateComment)
		authenticated.DELETE("/comments/:id", controllers.DeleteComment)

		authenticated.GET("/likes", controllers.GetLikes)
		authenticated.POST("/likes", controllers.CreateLike)
		authenticated.DELETE("/likes", controllers.DeleteLike)

		authenticated.GET("/communities", controllers.GetCommunities)
		authenticated.GET("/communities/:id", controllers.GetCommunity)
		authenticated.POST("/communities", controllers.CreateCommunity)
		authenticated.DELETE("/communities/:id", controllers.DeleteCommunity)
	}

	return r
}
