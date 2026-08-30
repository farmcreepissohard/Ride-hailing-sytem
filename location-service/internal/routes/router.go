package routes

import (
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/controllers"
	"github.com/gin-gonic/gin"
)

type RouterDependency struct {
	LocationController *controllers.LocationController
	WsController       *controllers.WsController
}

func SetupRouter(deps RouterDependency) *gin.Engine {
	r := gin.Default()

	LocationRoutes(r, deps.LocationController)
	WsRoutes(r, deps.WsController)

	return r
}
