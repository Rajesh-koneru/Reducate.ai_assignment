package handlers

import (
	"backend-assignment/database"
	"backend-assignment/models"

	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//user signing handler for user registration

//Input data struct

type userDetails struct {
	UserName string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// signin handler
func Signin(c *gin.Context) {

	var data userDetails
	c.ShouldBindJSON(&data)
	fmt.Println("user received", data.UserName)

	var existingUser models.User

	database.DB.Where(
		"email = ?",
		data.Email,
	).First(&existingUser)

	fmt.Println(existingUser.Email)

	if existingUser.ID != 0 {
		c.JSON(400, gin.H{
			"error": "Email already exists",
		})
		return

	}

	//hashing user password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(data.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {

		c.JSON(500, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	//creating a user objects
	user := models.User{
		Name:     data.UserName,
		Email:    data.Email,
		Password: string(hashedPassword),
		Role:     "user",
	}

	//saving user to the database
	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Insertion failed",
		})
	} else {

		c.JSON(201, gin.H{
			"message": "User created successfully.",
		})
	}

}
