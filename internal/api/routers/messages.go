package routers

import (
	messageControllers "github.com/fuadnafiz98/go-postgres/internal/api/controllers"
	"github.com/gin-gonic/gin"
)

func MessageRoutes(route *gin.Engine) {
	messages := route.Group("/messages")

	messages.GET("/", messageControllers.GetAllMessages)
}
