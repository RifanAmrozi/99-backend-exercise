package handlers

import (
	"99-backend-exercise/listing-service/services"
	"99-backend-exercise/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateListing(c *gin.Context) {
	userID, err := strconv.Atoi(c.PostForm("user_id"))
	if err != nil || userID <= 0 {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"result": false, "error": "Invalid user ID"})
		return
	}

	price, err := strconv.Atoi(c.PostForm("price"))
	if err != nil || price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"result": false, "error": "Invalid price"})
		return
	}

	listingType := c.PostForm("listing_type")
	if listingType != "rent" && listingType != "sale" {
		c.JSON(http.StatusBadRequest, gin.H{"result": false, "error": "Invalid listing type"})
		return
	}

	listing := models.Listing{
		UserID:      userID,
		Price:       price,
		ListingType: listingType,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := services.CreateListing(&listing); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":  true,
		"listing": listing,
	})
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

	listings, err := services.GetListings(pageNum, pageSize, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":   true,
		"listings": listings,
	})
}
