package dto

import "github.com/farmcreepissohard/Ride-hailing-sytem/internal/models/enum"

type ChangeStatusRequestDTO struct {
	Status enum.DriverStatus `json:"status" binding:"required,oneof=online offline"`
}
