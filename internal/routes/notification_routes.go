package routes

import (
	"graphql/internal/handlers"

	"github.com/gin-gonic/gin"
)

func NotificationRouteGroup(router *gin.Engine) *gin.RouterGroup {
	return router.Group("/notification")

}

func NotificationRouteSetup(routeGroup *gin.RouterGroup) {

	routeGroup.POST("/like", handlers.LikeNotificationHandler())

	routeGroup.POST("/comment", handlers.CommentNotificationHandler())

	routeGroup.POST("/admin/recipe", handlers.AdminRecipeNotificationHandler())
}
