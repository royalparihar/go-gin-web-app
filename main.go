package main

import (
	"github.com/royalparihar/go-gin-web-app/controllers"
	"github.com/royalparihar/go-gin-web-app/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/weather-report/:city", controllers.GetWeatherReport)

	// Books API
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	// User API
	r.GET("/user/:id", controllers.FindUser)
	r.POST("/user", controllers.CreateUser)
	r.PATCH("/user/:id/setadmin", controllers.CreateAdminUser)
	r.PATCH("/user/:id/removeadmin", controllers.RemoveAdminUser)
	r.PATCH("/user/:id", controllers.UpdateUser)
	r.PATCH("/user/:id/reset-password", controllers.UpdateUserPassword)
	r.DELETE("/user/:id", controllers.DeleteUser)
	r.POST("/login", controllers.Login)

	// Run the server
	r.Run()
}
