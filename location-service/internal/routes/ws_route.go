package routes

import (
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/controllers"
	"github.com/gin-gonic/gin"
)

func WsRoutes(r *gin.Engine, wsController *controllers.WsController) {
	ws := r.Group("api/v1/ws")
	{
		ws.GET("/", wsController.HandleConnection)
	}
}
