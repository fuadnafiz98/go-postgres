package controllers

import (
	"net/http"

	"github.com/fuadnafiz98/go-postgres/internal/database"
	"github.com/fuadnafiz98/go-postgres/internal/database/models"
	"github.com/gin-gonic/gin"
)

func GetAllMessages(c *gin.Context) {
	var messages []models.Message

	database.DB.Find(&messages)

	c.JSON(http.StatusOK, gin.H{"data": messages})
}
