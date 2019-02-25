package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/hubbdevelopers/hubb/db"
	"github.com/jinzhu/gorm"
	"google.golang.org/api/option"
)

var orm *gorm.DB

func Init() {
	orm = db.GetORM()
	//orm.LogMode(true)
}

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		if idToken == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		opt := option.WithCredentialsFile(os.Getenv("SECRETS_FILE"))
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Printf("error initializing app: %v\n", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		client, err := app.Auth(context.Background())
		if err != nil {
			log.Printf("error getting Auth client: %v\n", err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		token, err := client.VerifyIDToken(c, idToken)
		if err != nil {
			log.Printf("error verifying ID token: %v\n", err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		log.Println(token.UID)
		fmt.Println("Authorized!!")
		c.Next()
	}
}
