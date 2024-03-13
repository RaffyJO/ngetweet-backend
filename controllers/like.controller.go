package controllers

import (
	"net/http"
	"ngetweet/db"
	"ngetweet/models"

	"github.com/gin-gonic/gin"
)

func LikeIndex(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		return
	}

	authenticatedUser, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get authenticated user"})
		return
	}

	var likes []models.Like
	db.DB.Where("user_id = ?", authenticatedUser.ID).Find(&likes)
	c.JSON(http.StatusOK, gin.H{"data": likes})
}
