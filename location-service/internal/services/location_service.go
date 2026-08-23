package services

import (
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/models/dto"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/models/enum"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/repositories"
)

type LocationService interface {
	ChangeStatus(id string, status enum.DriverStatus) error
	UpdateLocation(id string, lng float64, lat float64) error
	MatchingNearbyDrivers(lng float64, lat float64, radius float64) ([]dto.MatchingResponse, error)
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

func (service *locationService) MatchingNearbyDrivers(lng float64, lat float64, radius float64) ([]dto.MatchingResponse, error) {
	return service.repo.MatchingNearbyDrivers(lng, lat, radius)
}
