package controllers

import (
	"net/http"
	"ngetweet/db"
	"ngetweet/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func UserIndex(c *gin.Context) {
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

	var users []models.User
	db.DB.Where("id = ?", authenticatedUser.ID).Preload("Tweets").Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func Register(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
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

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create token"})
		return
	}

	// Mengembalikan data pengguna dan token dalam respons
	c.JSON(http.StatusOK, gin.H{"data": user, "token": tokenString})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read body"})
		return
	}

	var user models.User
	db.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or password"})
		return
	}

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create token"})
		return
	}

	// Mengembalikan data pengguna dan token dalam respons
	c.JSON(http.StatusOK, gin.H{"data": user, "token": tokenString})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func Following(c *gin.Context) {
	var following models.Followings
	var followers models.Followers
	if err := c.ShouldBindJSON(&following); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	followers.UserID = following.FollowingId
	followers.FollowersgId = following.UserID

	if check := db.DB.First(&following); check.RowsAffected != 0 {
		db.DB.Delete(&following)
		db.DB.Delete(&followers)
		c.JSON(http.StatusOK, gin.H{"message": "Successfully Unfollow data."})
		return
	}

	db.DB.Create(&followers)
	result := db.DB.Create(&following)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusExpectationFailed, gin.H{"message": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Following data."})
}
