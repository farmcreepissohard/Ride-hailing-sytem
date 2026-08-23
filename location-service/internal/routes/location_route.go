package routes

import (
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(locationController *controllers.LocationController) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.PATCH("/driver", locationController.ChangeStatus)
		v1.PUT("/location", locationController.UpdateLocation)
	}

	return r
}
