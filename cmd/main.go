package main

import (
	"net/http"

	database "github.com/fuadnafiz98/go-postgres/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	database.ConnectDatabase()

	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "ok"})
	})

	server.Run()
}
