package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"github.com/pushpal-api/controllers"
	"gopkg.in/mgo.v2"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "You successfully reached the Pushpal API."})
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	database := session.DB("pushpal")

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/", Index)
		users := v1.Group("/users")
		{
			userController := controllers.NewUserController(database)
			users.GET("/", userController.GetUsers)
			users.GET("/:id", userController.GetUser)
			users.POST("/", userController.CreateUser)
		}

		auth := v1.Group("/auth")
		{
			authController := controllers.NewAuthController(database)
			auth.POST("/", authController.Authentication)
		}
	}
	router.Run(":12345")
}
