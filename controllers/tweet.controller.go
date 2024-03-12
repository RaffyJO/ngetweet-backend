package controllers

import (
	"net/http"
	"ngetweet/db"
	"ngetweet/models"
	"time"

	"github.com/gin-gonic/gin"
)

func TweetIndex(c *gin.Context) {
	var tweets []models.Tweet

	db.DB.Find(&tweets)
	c.JSON(http.StatusOK, gin.H{"data": tweets})
}

func TweetCreate(c *gin.Context) {
	var tweet models.Tweet

	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if tweet.Body == "" || tweet.UserID == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Make sure all the required data is filled in."})
		return
	}

	tweet.CreatedAt = time.Now()
	tweet.UpdatedAt = time.Now()

	result := db.DB.Create(&tweet)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusExpectationFailed, gin.H{"message": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully added data."})
}
