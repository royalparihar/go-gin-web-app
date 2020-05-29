package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/royalparihar/go-gin-web-app/models"
)

type LoginInput struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json: "password" binding:"required"`
}

type CreateUserInput struct {
	Name  string `json:"name" binding:"required"`
	UserName string `json:"user_name" binding:"required"`
	Password string `json: "password" binding:"required"`
}

type UpdateUserInput struct {
	Name  string `json:"name"`
}

type UpdateUserPasswordInput struct {
	Password  string `json:"password" binding:"required"`
}

// GET /user/:id
func FindUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// POST /user
func CreateUser(c *gin.Context) {
	var savedUser models.User
	var input CreateUserInput
	if err := models.DB.Where("user_name = ?", input.UserName).First(&savedUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already created!"})
		return
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name, UserName: input.UserName, Password: input.Password}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})

}

// PATCH /user/:id/setadmin
func CreateAdminUser(c *gin.Context) {

	var user models.User
	id := c.Param("id")
	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Model(&user).Updates("IsAdmin", true)

	c.JSON(http.StatusOK, gin.H{"data": true})

}

// PATCH /user/:id/removeadmin
func RemoveAdminUser(c *gin.Context) {

	var user models.User
	id := c.Param("id")
	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Model(&user).Updates("IsAdmin", false)

	c.JSON(http.StatusOK, gin.H{"data": true})

}
// PATCH /user/:id
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PATCH /user/:id/reset-password
func UpdateUserPassword(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateUserPasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /user/:id
func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Model(&user).Updates("IsDelete", true)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func Login(c *gin.Context) {
	var user models.User
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Where("user_name = ? AND password = ?", input.UserName, input.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}