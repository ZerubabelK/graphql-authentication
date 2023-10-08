package routes

import (
	"graphql/internal/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRouteGroup(router *gin.Engine) *gin.RouterGroup {
	return router.Group("/auth")
}

func AuthRouteSetup(routeGroup *gin.RouterGroup) {
	routeGroup.POST("/login", handlers.LoginHandler())
	routeGroup.POST("/register", handlers.RegisterHandler())
}
