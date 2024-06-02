package handlers

import (
	"99-backend-exercise/public-api/client"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var requestBody struct {
		Name string `json:"name"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := client.CreateUser(requestBody.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
