package controllers

import (
	"net/http"

	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/models/dto"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/services"
	"github.com/gin-gonic/gin"
)

type LocationController struct {
	service services.LocationService
}

func NewLocationController(service services.LocationService) *LocationController {
	return &LocationController{service: service}
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
