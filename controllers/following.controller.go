package controllers

import(
	"ngetweet/db"
	"ngetweet/models"
	"github.com/gin-gonic/gin"
	"net/http"

)
func FollowIndex(c *gin.Context){
	var followings  []models.Followings

	db.DB.Preload("User").Preload("Following").Find(&followings)
	c.JSON(http.StatusOK,gin.H{"followings":followings})
}