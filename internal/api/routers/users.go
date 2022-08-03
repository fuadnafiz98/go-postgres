package routers

import (
	"github.com/fuadnafiz98/go-postgres/internal/api/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {
	users := route.Group("/users")

	users.GET("/", controllers.GetAllUsers)
	users.GET("/add", controllers.AddRandomUser)
}
