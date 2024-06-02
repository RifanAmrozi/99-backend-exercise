package main

import (
	"99-backend-exercise/public-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/public-api/listings", handlers.GetListings)
	router.POST("/public-api/users", handlers.CreateUser)
	router.POST("/public-api/listings", handlers.CreateListing)

	router.Run(":8080")
}
