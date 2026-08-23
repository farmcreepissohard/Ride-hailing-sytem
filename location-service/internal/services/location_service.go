package services

import (
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/models/enum"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/repositories"
)

type LocationService interface {
	ChangeStatus(id string, status enum.DriverStatus) error
	UpdateLocation(id string, lng float64, lat float64) error
}

type locationService struct {
	repo repositories.LocationRepository
}

func NewLocationService(repo repositories.LocationRepository) LocationService {
	return &locationService{repo: repo}
}

func (service *locationService) ChangeStatus(id string, status enum.DriverStatus) error {
	return service.repo.ChangeStatus(id, status)
}

func (service *locationService) UpdateLocation(id string, lng float64, lat float64) error {
	return service.repo.UpdateLocation(id, lng, lat)
}
