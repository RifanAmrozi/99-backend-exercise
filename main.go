package main

import (
	"99-backend-exercise/db"
	"99-backend-exercise/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()

	router := gin.Default()

	router.POST("/users", handlers.CreateUser)
	router.GET("/users/:id", handlers.GetUserByID)
	router.POST("/listings", handlers.CreateListing)
	router.GET("/listings", handlers.GetListings)

	router.Run(":8080")
}
