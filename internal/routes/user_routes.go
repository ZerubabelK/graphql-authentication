package routes

import (
	"graphql/internal/handlers"

	"github.com/gin-gonic/gin"
)

func UserRouteGroup(router *gin.Engine) *gin.RouterGroup {
	return router.Group("/user")

}

func UserRouteSetup(routeGroup *gin.RouterGroup) {
	routeGroup.POST("/me", handlers.GetUserHandler())

	routeGroup.POST("/update/profile_image", handlers.UpdateProfileImageHandler())
}
