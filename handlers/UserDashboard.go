package handlers

import (
	"backend-assignment/database"
	"backend-assignment/models"

	"github.com/gin-gonic/gin"
)

func UserDash(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "User dashboard",
	})
}

func Profile(c *gin.Context) {
	var profile models.User
	database.DB.First(&profile, 1)

	c.JSON(200, gin.H{
		"Profile": profile,
	})

}

func LogOut(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg":    "Logout Succesful",
		"status": "True",
	})
}
