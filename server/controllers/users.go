package controllers

import (
	"bookstore/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := models.User{Username: input.Username, Password: input.Password, ConfirmPassword: input.ConfirmPassword}

	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

