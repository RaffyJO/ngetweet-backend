package controllers

import (
	"net/http"
	"ngetweet/db"
	"ngetweet/models"
	
	"github.com/gin-gonic/gin"
)

func GetCategory(c *gin.Context){
	
	var category []models.Category
	db.DB.Find(&category)
	c.JSON(http.StatusOK, gin.H{"data": category})
}

func CreateCategory(c *gin.Context){
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to process request data"})
		return
	}

	if result:= db.DB.Create(&category); result.Error!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category created successfully"})

}