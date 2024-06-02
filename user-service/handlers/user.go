package handlers

import (
	"99-backend-exercise/models"
	"99-backend-exercise/user-service/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	user := models.User{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := services.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"user":   user,
	})
}

func GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": false, "error": "Invalid user ID"})
		return
	}

	user, err := services.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"user":   user,
	})
}
