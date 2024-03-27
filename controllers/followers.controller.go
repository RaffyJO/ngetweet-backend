package controllers

import (
	"net/http"
	"ngetweet/db"
	"ngetweet/models"

	"github.com/gin-gonic/gin"
)

func FollowersIndex(c *gin.Context) {
	var followers []models.Followers

	db.DB.Preload("User").Preload("Followers").Find(&followers)
	c.JSON(http.StatusOK, gin.H{"followers": followers})
}
