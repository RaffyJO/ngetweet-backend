package controllers

import (
	"net/http"
	"ngetweet/db"
	"ngetweet/models"

	"github.com/gin-gonic/gin"
)

func FollowIndex(c *gin.Context) {
	var followings []models.Followings

	db.DB.Preload("User").Preload("Following").Find(&followings)
	c.JSON(http.StatusOK, gin.H{"followings": followings})
}
