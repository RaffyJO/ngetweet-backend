package controllers

import(
	"ngetweet/db"
	"ngetweet/models"
	"github.com/gin-gonic/gin"
	"net/http"

)
func FollowIndex(c *gin.Context){
	user, exists := c.Get("user")
	println(user)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		return
	}

	authenticatedUser, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get authenticated user"})
		return
	}

	var followings  []models.Followings

	db.DB.Where("user_id = ?", authenticatedUser.ID).Preload("User").Preload("Following").Find(&followings)
	c.JSON(http.StatusOK,gin.H{"followings":followings})
}