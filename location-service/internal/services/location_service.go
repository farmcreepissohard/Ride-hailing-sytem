package services

import (
	"context"
	"errors"

	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/grpc_client"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/models/dto"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/models/enum"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/repositories"
)

type LocationService interface {
	ChangeStatus(id string, status enum.DriverStatus) error
	UpdateLocation(id string, lng float64, lat float64) error
	MatchingNearbyDrivers(lng float64, lat float64, radius float64) ([]dto.MatchingResponse, error)

	AcceptTrip(ctx context.Context, tripId string, driverId string) error
	StartTrip(ctx context.Context, tripId string, driverId string) error
	CompleteTrip(ctx context.Context, tripId string, driverId string) error
	CancelTrip(ctx context.Context, tripId string, driverId string, reason string) error
}

type locationService struct {
	grpcClient grpc_client.TripGrpcClient
	repo       repositories.LocationRepository
}

func NewLocationService(grpcClient grpc_client.TripGrpcClient, repo repositories.LocationRepository) LocationService {
	return &locationService{grpcClient: grpcClient, repo: repo}
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

func (service *locationService) AcceptTrip(ctx context.Context, tripId string, driverId string) error {
	if err := service.repo.OnReady(driverId); err != nil {
		return err
	}

	isSuccess, err := service.grpcClient.MatchTrip(ctx, tripId, driverId)
	if err != nil {
		return err
	}

	if !isSuccess {
		return errors.New("Trip already taken or forbidden")
	}

	return nil
}

func (service *locationService) StartTrip(ctx context.Context, tripId string, driverId string) error {
	if err := service.repo.OnBusy(driverId); err != nil {
		return err
	}

	isSuccess, err := service.grpcClient.StartTrip(ctx, tripId, driverId)
	if err != nil {
		return err
	}

	if !isSuccess {
		return errors.New("Can not start trip, please try again later")
	}

	return nil
}

func (service *locationService) CompleteTrip(ctx context.Context, tripId string, driverId string) error {
	if err := service.repo.OutBusy(driverId); err != nil {
		return err
	}

	isSuccess, err := service.grpcClient.CompleteTrip(ctx, tripId, driverId)
	if err != nil {
		return err
	}

	if !isSuccess {
		return errors.New("Can not start trip, please try again later")
	}

	return nil
}

func (service *locationService) CancelTrip(ctx context.Context, tripId string, driverId string, reason string) error {
	if err := service.repo.OnCancel(driverId); err != nil {
		return err
	}

	isSuccess, err := service.grpcClient.CancelTrip(ctx, tripId, driverId, reason)
	if err != nil {
		return err
	}

	if !isSuccess {
		return errors.New("Can not start trip, please try again later")
	}

	return nil
}
