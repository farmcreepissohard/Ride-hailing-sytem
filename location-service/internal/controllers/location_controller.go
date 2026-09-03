package controllers

import (
	"net/http"

	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/models/dto"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/services"
	"github.com/gin-gonic/gin"
)

type LocationController struct {
	service  services.LocationService
	dispatch services.DispatchService
}

func NewLocationController(service services.LocationService, dispatch services.DispatchService) *LocationController {
	return &LocationController{service: service, dispatch: dispatch}
}

func (controller *LocationController) ChangeStatus(c *gin.Context) {
	id := c.GetHeader("X-Driver-Id")
	if id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req dto.ChangeStatusRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.service.ChangeStatus(id, req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"driverId": id, "status": req.Status})
}

func (controller *LocationController) UpdateLocation(c *gin.Context) {
	id := c.GetHeader("X-Driver-Id")
	if id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req dto.LocationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.service.UpdateLocation(id, req.Longitude, req.Latitude); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func (controller *LocationController) AcceptTrip(c *gin.Context) {
	id := c.GetHeader("X-Driver-Id")
	if id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	tripId := c.Param("tripId")
	if tripId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid trip id"})
		return
	}

	if err := controller.service.AcceptTrip(c.Request.Context(), tripId, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	controller.dispatch.NotifyResponse(tripId, true)

	c.JSON(http.StatusOK, gin.H{"message": "accept successfully"})
}

func (controller *LocationController) DeclineTrip(c *gin.Context) {
	id := c.GetHeader("X-Driver-Id")
	if id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	tripId := c.Param("tripId")
	if tripId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid trip id"})
		return
	}

	controller.dispatch.NotifyResponse(tripId, false)

	c.JSON(http.StatusOK, gin.H{"message": "decline successfully"})
}

func (controller *LocationController) StartTrip(c *gin.Context) {
	id := c.GetHeader("X-Driver-Id")
	if id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	tripId := c.Param("tripId")
	if tripId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid trip id"})
		return
	}

	if err := controller.service.StartTrip(c.Request.Context(), tripId, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "start successfully"})
}

func (controller *LocationController) CompleteTrip(c *gin.Context) {
	id := c.GetHeader("X-Driver-Id")
	if id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	tripId := c.Param("tripId")
	if tripId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid trip id"})
		return
	}

	if err := controller.service.CompleteTrip(c.Request.Context(), tripId, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "complete successfully"})
}

func (controller *LocationController) CancelTrip(c *gin.Context) {
	id := c.GetHeader("X-Driver-Id")
	if id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	tripId := c.Param("tripId")
	if tripId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid trip id"})
		return
	}

	var req dto.CancelReasonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := controller.service.CancelTrip(c.Request.Context(), tripId, id, req.Reason); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "cancel successfully"})
}
