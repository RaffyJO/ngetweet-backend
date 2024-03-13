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
	db.DB.Where("user_id = ?", authenticatedUser.ID).Find(&tweets)
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

func AddLike(c *gin.Context) {
	var tweet models.Tweet

	tweetID := c.Param("id")
	if err := db.DB.First(&tweet, tweetID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tweet not found"})
		return
	}

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

	var like models.Like
	if err := db.DB.Where("tweet_id = ? AND user_id = ?", tweet.ID, authenticatedUser.ID).First(&like).Error; err == nil {
		if err := db.DB.Delete(&like).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove like"})
			return
		}

		tweet.Likes--
	} else {
		like = models.Like{
			TweetID: tweet.ID,
			UserID:  authenticatedUser.ID,
		}
		if err := db.DB.Create(&like).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add like"})
			return
		}

		tweet.Likes++
	}

	if err := db.DB.Save(&tweet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tweet"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Like status updated successfully"})
}
