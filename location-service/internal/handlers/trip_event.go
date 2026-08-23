package handlers

import (
	"encoding/json"
	"log"

	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/models/payload"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/services"
	"github.com/rabbitmq/amqp091-go"
)

type TripEventHandler struct {
	service services.LocationService
}

func NewTripEventHanlder(service *services.LocationService) *TripEventHandler {
	return &TripEventHandler{service: *service}
}

func (h *TripEventHandler) Consume(ch *amqp091.Channel, queueName string) {
	msgs, err := ch.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	log.Printf("Waiting for messages in queue: %s", queueName)

	go func() {
		for d := range msgs {
			log.Printf("📩 Received a message: %s", d.Body)

			var payload payload.TripEventPayload
			if err := json.Unmarshal(d.Body, &payload); err != nil {
				log.Printf("JSON Unmarshal error: %v", err)
				d.Nack(false, false)
				continue
			}

			drivers, err := h.service.MatchingNearbyDrivers(payload.Longitude, payload.Latitude, 3.0)
			if err != nil {
				log.Printf("Failed to search Redis: %v", err)
				d.Nack(false, true)
				continue
			}

			log.Printf("Found %d drivers for Trip %s", len(drivers), payload.TripID)

			d.Ack(false)
		}

	}()

}
