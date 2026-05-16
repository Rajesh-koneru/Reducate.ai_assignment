package handlers

import (
	"github.com/gin-gonic/gin"

	"backend-assignment/database"

	"backend-assignment/models"
	"fmt"
)

func AdminDash(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "welcome to the Admin panel",
	})
}

func GetUsers(c *gin.Context) {
	var users []models.User

	database.DB.Find(&users)

	fmt.Println("all users ", users)

	c.JSON(200, gin.H{
		"data": users,
	})

}

// deleting user details
type data struct {
	Email string `json:"email"`
}

func DeleteUsers(c *gin.Context) {

	var details data
	var user models.User

	err := c.ShouldBindJSON(&details)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	result := database.DB.Where("Email = ?", details.Email).Delete(&user)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"msg": "User not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "User deleted",
	})
}
