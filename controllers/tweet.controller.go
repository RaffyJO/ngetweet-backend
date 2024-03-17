package controllers

import (
	"net/http"
	"ngetweet/db"
	"ngetweet/models"
	"time"

	"github.com/gin-gonic/gin"
)

func TweetIndex(c *gin.Context) {
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

	var tweets []models.Tweet
	db.DB.Where("user_id = ?", authenticatedUser.ID).Preload("UserLikes").Preload("Comments").Find(&tweets)
	c.JSON(http.StatusOK, gin.H{"data": tweets})
}

func TweetCreate(c *gin.Context) {
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

	var tweet models.Tweet
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	tweet.UserID = authenticatedUser.ID
	if tweet.Body == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Make sure the tweet body is filled in"})
		return
	}

	tweet.CreatedAt = time.Now()
	tweet.UpdatedAt = time.Now()

	result := db.DB.Create(&tweet)
	if result.Error != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"message": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully added tweet"})
}

func DeleteTweet(c *gin.Context) {
	// Get logged user
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

	// Get tweet ID
	tweetID := c.Param("id")

	var tweet models.Tweet
	if err := db.DB.First(&tweet, tweetID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Tweet not found"})
		return
	}

	// Check user permission tweet
	if tweet.UserID != authenticatedUser.ID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to delete this tweet"})
		return
	}

	// Delete tweet
	if err := db.DB.Delete(&tweet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete tweet"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tweet deleted successfully"})
}
