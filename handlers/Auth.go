
package handlers

import (
	"backend-assignment/database"
	"backend-assignment/models"
	"backend-assignment/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {

	var input LoginInput

	c.ShouldBindJSON(&input)

	var user models.User

	result := database.DB.Where(
		"email = ?",
		input.Email,
	).First(&user)

	if result.Error != nil {

		c.JSON(401, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	)

	if err != nil {

		c.JSON(401, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	token, err1 := utils.GenerateJWT(
		user.Email,
		user.Role,
	)

	if err1 != nil {

		c.JSON(500, gin.H{
			"error": "Failed to generate token",
		})

		return
	}

	c.JSON(200, gin.H{
		"token": token,
		"messge":"Login success",
	})
}