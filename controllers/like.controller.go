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

func AddLike(c *gin.Context) {
	var tweet models.Tweet

	tweetID := c.Param("id")
	if err := db.DB.First(&tweet, tweetID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Tweet not found"})
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
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to remove like"})
			return
		}

		tweet.Likes--
	} else {
		like = models.Like{
			TweetID: tweet.ID,
			UserID:  authenticatedUser.ID,
		}
		if err := db.DB.Create(&like).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add like"})
			return
		}

		tweet.Likes++
	}

	if err := db.DB.Save(&tweet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update tweet"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Like status updated successfully"})
}
