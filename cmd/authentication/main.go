package main

import (
	"graphql/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.RouteSetup(router)

	router.Run(":5000")
}
