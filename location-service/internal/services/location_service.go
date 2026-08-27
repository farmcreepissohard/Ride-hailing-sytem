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
	if id == "" {
		return errors.New("Driver id is required")
	}
	if status == "" {
		return errors.New("Status is required")
	}

	return service.repo.ChangeStatus(id, status)
}

func (service *locationService) UpdateLocation(id string, lng float64, lat float64) error {
	if id == "" {
		return errors.New("Driver id is required")
	}
	if lat < -90 || lat > 90 {
		return errors.New("Invalid latitude value")
	}
	if lng < -180 || lng > 180 {
		return errors.New("Invalid longitude value")
	}

	return service.repo.UpdateLocation(id, lng, lat)
}

func (service *locationService) MatchingNearbyDrivers(lng float64, lat float64, radius float64) ([]dto.MatchingResponse, error) {
	if lat < -90 || lat > 90 {
		return nil, errors.New("Invalid latitude value")
	}
	if lng < -180 || lng > 180 {
		return nil, errors.New("Invalid longitude value")
	}
	if radius < 0 {
		return nil, errors.New("Invalid radius value")
	}

	return service.repo.MatchingNearbyDrivers(lng, lat, radius)
}

func (service *locationService) AcceptTrip(ctx context.Context, tripId string, driverId string) error {
	if ctx == nil {
		return errors.New("Context is required")
	}
	if tripId == "" {
		return errors.New("Trip id is required")
	}
	if driverId == "" {
		return errors.New("Driver id is required")
	}

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
	if ctx == nil {
		return errors.New("Context is required")
	}
	if tripId == "" {
		return errors.New("Trip id is required")
	}
	if driverId == "" {
		return errors.New("Driver id is required")
	}

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
	if ctx == nil {
		return errors.New("Context is required")
	}
	if tripId == "" {
		return errors.New("Trip id is required")
	}
	if driverId == "" {
		return errors.New("Driver id is required")
	}

	if err := service.repo.OutBusy(driverId); err != nil {
		return err
	}

	isSuccess, err := service.grpcClient.CompleteTrip(ctx, tripId, driverId)
	if err != nil {
		return err
	}

	if !isSuccess {
		return errors.New("Can not complete trip, please try again later")
	}

	return nil
}

func (service *locationService) CancelTrip(ctx context.Context, tripId string, driverId string, reason string) error {
	if ctx == nil {
		return errors.New("Context is required")
	}
	if tripId == "" {
		return errors.New("Trip id is required")
	}
	if driverId == "" {
		return errors.New("Driver id is required")
	}
	if reason == "" {
		return errors.New("Unknown reason")
	}

	if err := service.repo.OnCancel(driverId); err != nil {
		return err
	}

	isSuccess, err := service.grpcClient.CancelTrip(ctx, tripId, driverId, reason)
	if err != nil {
		return err
	}

	if !isSuccess {
		return errors.New("Not allowed to cancel this trip")
	}

	return nil
}
