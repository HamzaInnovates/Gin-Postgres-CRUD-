package controller

import (
	"postgres/config"
	"postgres/models"

	"github.com/gin-gonic/gin"
)

func GetUserData(c *gin.Context) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(200, users)
}
func AddUserData(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(201, user)

}
func UpdateUserData(c *gin.Context) {
	var user models.User
	Userid := c.Param("id")
	if err := config.DB.First(&user, Userid).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(201, user)
}
func DeleteUserData(c *gin.Context) {
	var user models.User
	Userid := c.Param("id")
	if err := config.DB.First(&user, Userid).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	config.DB.Delete(&user)
	c.JSON(201, "User deleted successfuly")
}
func GetUser(c *gin.Context) {
	var user models.User
	Userid := c.Param("id")
	if err := config.DB.First(&user, Userid).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(201, user)
}
