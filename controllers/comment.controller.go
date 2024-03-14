package controllers

import (
	"net/http"
	"ngetweet/db"
	"ngetweet/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CommentIndex(c *gin.Context) {
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

	var comments []models.Comment
	db.DB.Where("user_id = ?", authenticatedUser.ID).Find(&comments)
	c.JSON(http.StatusOK, gin.H{"data": comments})
}

func AddComment(c *gin.Context) {
	var tweet models.Tweet

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

	tweetID := c.Param("id")
	if err := db.DB.First(&tweet, tweetID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Tweet not found"})
		return
	}

	var request struct {
		Body            string `json:"body"`
		ParentCommentID uint   `json:"parent_comment_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to process request data"})
		return
	}

	// Create object comment
	comment := models.Comment{
		UserID:          authenticatedUser.ID,
		TweetID:         tweet.ID,
		Body:            request.Body,
		ParentCommentID: request.ParentCommentID,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := db.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add comment"})
		return
	}
	tweet.TotalComments++

	if err := db.DB.Save(&tweet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update tweet"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment created successfully"})
}

func DeleteComment(c *gin.Context) {
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

	// Get comment ID
	commentID := c.Param("id")

	var comment models.Comment
	if err := db.DB.First(&comment, commentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Comment not found"})
		return
	}

	// Check user permission comment
	if comment.UserID != authenticatedUser.ID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to delete this comment"})
		return
	}

	// Delete comment
	if err := db.DB.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete comment"})
		return
	}

	// Update tweet total comment
	var tweet models.Tweet
	if err := db.DB.First(&tweet, comment.TweetID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Comment not found"})
		return
	}
	tweet.TotalComments--

	if err := db.DB.Save(&tweet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update tweet"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
