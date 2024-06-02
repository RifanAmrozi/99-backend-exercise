package handlers

import (
	"99-backend-exercise/public-api/client"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateListing(c *gin.Context) {
	var requestBody struct {
		UserID      int    `json:"user_id"`
		ListingType string `json:"listing_type"`
		Price       int    `json:"price"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	listing, err := client.CreateListing(requestBody.UserID, requestBody.ListingType, requestBody.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"listing": listing})
}

func GetListings(c *gin.Context) {
	pageNumStr := c.DefaultQuery("page_num", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")
	userIDStr := c.Query("user_id")

	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil || pageNum <= 0 {
		pageNum = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	var userID *int
	if userIDStr != "" {
		id, err := strconv.Atoi(userIDStr)
		if err == nil {
			userID = &id
		}
	}

	listings, err := client.GetListings(pageNum, pageSize, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":   true,
		"listings": listings,
	})
}
