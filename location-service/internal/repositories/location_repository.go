package repositories

import (
	"context"
	"errors"

	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/models/dto"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/models/enum"
	"github.com/redis/go-redis/v9"
)

type LocationRepository interface {
	ChangeStatus(id string, status enum.DriverStatus) error
	UpdateLocation(id string, lng float64, lat float64) error
	MatchingNearbyDrivers(lng float64, lat float64, radius float64) ([]dto.MatchingResponse, error)
}

type locationRepository struct {
	redisClient *redis.Client
}

func NewLocationRepository(redisClient *redis.Client) LocationRepository {
	return &locationRepository{redisClient: redisClient}
}

func (repository *locationRepository) ChangeStatus(id string, status enum.DriverStatus) error {
	rdb := repository.redisClient
	ctx := context.Background()

	if status == enum.StatusOnline {
		if err := rdb.SRem(ctx, "offline_drivers", id, 0).Err(); err != nil {
			return err
		}
		if err := rdb.SAdd(ctx, "online_drivers", id, 0).Err(); err != nil {
			return err
		}
	} else if status == enum.StatusOffline {
		if err := rdb.SRem(ctx, "online_drivers", id, 0).Err(); err != nil {
			return err
		}
		if err := rdb.SAdd(ctx, "offline_drivers", id, 0).Err(); err != nil {
			return err
		}
		if err := rdb.ZRem(ctx, "location_drivers", &redis.GeoLocation{Name: id}).Err(); err != nil {
			return err
		}
	}

	return nil
}

func (repository *locationRepository) UpdateLocation(id string, lng float64, lat float64) error {
	rdb := repository.redisClient
	backgroundContext := context.Background()

	if existed := rdb.SIsMember(backgroundContext, "online_drivers", id); !existed.Val() {
		return errors.New("Driver is not online")
	}

	if err := rdb.GeoAdd(backgroundContext, "location_drivers", &redis.GeoLocation{Name: id, Longitude: lng, Latitude: lat}).Err(); err != nil {
		return err
	}

	return nil
}

func (repository *locationRepository) MatchingNearbyDrivers(lng float64, lat float64, radius float64) ([]dto.MatchingResponse, error) {
	rdb := repository.redisClient
	backgroundContext := context.Background()

	res := rdb.GeoSearch(backgroundContext, "location_drivers", &redis.GeoSearchQuery{Latitude: lat, Longitude: lng, Radius: radius, RadiusUnit: "km"})
	if err := res.Err(); err != nil {
		return nil, err
	}

	var drivers []dto.MatchingResponse
	for _, v := range res.Val() {
		drivers = append(drivers, dto.MatchingResponse{DriverId: v})
	}

	return drivers, nil
}
