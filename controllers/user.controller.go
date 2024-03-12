package controllers

import (
	"log"
	"net/http"
	"ngetweet/db"
	"ngetweet/models"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserIndex(c *gin.Context) {
	var users []models.User

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic recovered: %v", r)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		}
	}()

	if err := db.DB.Preload("Tweets").Find(&users).Error; err != nil {
		log.Println("Error querying users:", err)
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func UserCreate(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if user.Name == "" || user.Email == "" || user.Nickname == "" || user.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Make sure all the required data is filled in."})
		return
	}

	var existingUser models.User
	if err := db.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "Email is already registered."})
		return
	}

	if err := db.DB.Where("nickname = ?", user.Nickname).First(&existingUser).Error; err == nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "Nickname is already taken."})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash password."})
		return
	}

	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	result := db.DB.Create(&user)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusExpectationFailed, gin.H{"message": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully added data."})
}
