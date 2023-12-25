package routes

import (
	"graphql/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RouteSetup(router *gin.Engine) {
	router.MaxMultipartMemory = 8 << 20
	authRouteGroup := AuthRouteGroup(router)
	AuthRouteSetup(authRouteGroup)

	userRouteGroup := UserRouteGroup(router)
	userRouteGroup.Use(middleware.AuthMiddleware())
	UserRouteSetup(userRouteGroup)

	notificationRouteGroup := NotificationRouteGroup(router)
	NotificationRouteSetup(notificationRouteGroup)
}
