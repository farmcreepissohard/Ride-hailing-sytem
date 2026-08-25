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
		v1.POST("/trips/:tripId/accept", locationController.AcceptTrip)
		v1.POST("/trips/:tripId/start", locationController.StartTrip)
		v1.POST("/trips/:tripId/complete", locationController.CompleteTrip)
		v1.POST("/trips/:tripId/cancel", locationController.CancelTrip)
	}

	return r
}
