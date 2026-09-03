package services

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/models/dto"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/repositories"
	"github.com/redis/go-redis/v9"
)

type DispatchService interface {
	HandlingTrip(tripId string, longitude float64, latitude float64, radius float64) error
	NotifyResponse(tripId string, accepted bool) bool
}

type dispatchService struct {
	repo        repositories.LocationRepository
	redisClient *redis.Client
	mutex       sync.RWMutex
	waitList    map[string]chan bool
}

func NewDispatchService(repo repositories.LocationRepository, redisClient *redis.Client) DispatchService {
	return &dispatchService{repo: repo, redisClient: redisClient, waitList: make(map[string]chan bool)}
}

func (service *dispatchService) HandlingTrip(tripId string, longitude float64, latitude float64, radius float64) error {
	drivers, err := service.repo.MatchingNearbyDrivers(longitude, latitude, radius)
	if err != nil {
		log.Printf("Failed to search Redis: %v", err)
		return err
	}

	responseChan := make(chan bool)
	service.mutex.Lock()
	service.waitList[tripId] = responseChan
	service.mutex.Unlock()

	go func(tripId string, driverList []dto.MatchingResponse) {
		defer func() {
			service.mutex.Lock()
			delete(service.waitList, tripId)
			service.mutex.Unlock()
			close(responseChan)
		}()

		isAccepted := false

		for _, driver := range driverList {

			channelName := "ws_" + driver.DriverId
			jsonData, err := json.Marshal(map[string]interface{}{
				"title":   "NEW_TRIP",
				"trip_id": tripId,
			})
			if err != nil {
				continue
			}
			service.redisClient.Publish(context.Background(), channelName, jsonData)

			select {

			case <-time.After(15 * time.Second):
				log.Printf("Timeout 15s")
			case accepted := <-responseChan:
				if accepted {
					isAccepted = true
					log.Printf("Driver %s accepted", driver.DriverId)
					return
				}
			}
		}
		if !isAccepted {
			log.Printf("Trip %s No drivers accept", tripId)
		}
	}(tripId, drivers)

	return nil
}

func (service *dispatchService) NotifyResponse(tripId string, accepted bool) bool {
	service.mutex.RLock()
	ch, exists := service.waitList[tripId]
	service.mutex.RUnlock()

	if exists {
		ch <- accepted
		return true
	}
	return false
}
