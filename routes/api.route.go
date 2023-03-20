package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hardikkheni/seedsuite-server/controllers"
	"github.com/hardikkheni/seedsuite-server/middlewares"
	"github.com/hardikkheni/seedsuite-server/requests"
)

func LoadApiRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", middlewares.Validate(new(requests.LoginRequest)), controllers.Login)
	}
}
