package main

import (
	"fmt"

	"backend-assignment/handlers"

	"github.com/gin-gonic/gin"

	"backend-assignment/Middlewares"
	"backend-assignment/database"
)

func main() {
	database.Connect()
	database.CreateUserTable()

	r := gin.Default()

	r.POST("/login", handlers.Login)
	r.POST("Signin", handlers.Signin)

	// Admin Routes
	admin := r.Group("/admin")
	admin.Use(Middlewares.AuthMiddleware("admin"))
	{
		admin.GET("/dashboard", handlers.AdminDash)
		admin.GET("/users", handlers.GetUsers)
		admin.DELETE("/delete", handlers.DeleteUsers)

	}
	// User Routes
	user := r.Group("/user")
	user.Use(Middlewares.AuthMiddleware("user"))
	{
		user.GET("/dashboard", handlers.UserDash)
		user.GET("Profile", handlers.Profile)

	}

	r.GET("/logout", handlers.LogOut)

	fmt.Println("http://localhost:8000")

	r.Run(":8000")

}
