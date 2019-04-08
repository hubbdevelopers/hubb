package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	firebase "firebase.google.com/go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hubbdevelopers/hubb/controllers"
	"github.com/hubbdevelopers/hubb/db"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {

	godotenv.Load()

	db.Connect()
	defer db.Close()

	db.Migrate()

	r := gin.Default()

	// r.Use(auth())

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

	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:id/init", controllers.InitializeUser)
	r.PUT("/users/:id/images", controllers.UpdateImage)
	r.PUT("/users/:id/profile", controllers.UpdateProfile)

	r.GET("/users/:id/timeline", controllers.GetTimeline)

	r.GET("/users/:id/followings", controllers.GetUserFollowings)
	r.GET("/users/:id/followers", controllers.GetUserFollowers)
	//r.GET("/communities/:id/followers", controllers.GetCommunityFollowers)
	r.POST("/users/:id/follow/users/:followingid", controllers.FollowUser)
	r.DELETE("/users/:id/follow/users/:followingid", controllers.UnfollowUser)
	r.POST("/users/:id/follow/communities/:followingid", controllers.FollowCommunity)
	r.DELETE("/users/:id/follow/communities/:followingid", controllers.UnfollowCommunity)

	r.GET("/pages", controllers.GetPages)
	r.GET("/pages/:id", controllers.GetPage)
	r.POST("/pages", controllers.CreatePage)
	r.DELETE("/pages/:id", controllers.DeletePage)
	r.PUT("/pages/:id", controllers.UpdatePage)
	r.GET("/recentpages", controllers.GetRecentPages)

	r.GET("/comments", controllers.GetComments)
	r.POST("/comments", controllers.CreateComment)
	r.DELETE("/comments/:id", controllers.DeleteComment)

	r.GET("/likes", controllers.GetLikes)
	r.POST("/likes", controllers.CreateLike)
	r.DELETE("/likes", controllers.DeleteLike)

	r.GET("/communities", controllers.GetCommunities)
	r.GET("/communities/:id", controllers.GetCommunity)
	r.POST("/communities", controllers.CreateCommunity)
	r.DELETE("/communities/:id", controllers.DeleteCommunity)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		opt := option.WithCredentialsFile(os.Getenv("SECRETS_FILE"))
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
		}

		client, err := app.Auth(context.Background())
		if err != nil {
			log.Fatalf("error getting Auth client: %v\n", err)
		}

		fmt.Println(reflect.TypeOf(client))

		authHeader := c.GetHeader("Authorization")
		// fmt.Println(authHeader)

		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		// fmt.Println(idToken)

		if idToken != "" {
			token, err := client.VerifyIDToken(c, idToken)
			if err != nil {
				log.Fatalf("error verifying ID token: %v\n", err)
				c.Next()
			}

			log.Println(token.UID)
			fmt.Println("Authorized!!")

			// TODO
		}

		c.Next()
	}
}
