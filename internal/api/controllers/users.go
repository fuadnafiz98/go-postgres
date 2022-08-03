package controllers

import (
	"fmt"
	"net/http"

	"github.com/fuadnafiz98/go-postgres/internal/database"
	"github.com/fuadnafiz98/go-postgres/internal/database/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User

	database.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func AddRandomUser(c *gin.Context) {
	user := models.User{Username: "fuadnafiz98", Email: "nafizfuad98@gmail.com"}

	result := database.DB.Create(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully!"})
}
