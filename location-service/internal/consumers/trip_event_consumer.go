package consumers

import (
	"encoding/json"
	"log"

	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/models/payload"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/services"
	"github.com/rabbitmq/amqp091-go"
)

type TripEventConsumer struct {
	service services.DispatchService
}

func NewTripEventConsumer(service services.DispatchService) *TripEventConsumer {
	return &TripEventConsumer{service: service}
}

func (h *TripEventConsumer) Consume(ch *amqp091.Channel, queueName string) {
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

			if err := h.service.HandlingTrip(payload.TripID, payload.Longitude, payload.Latitude, 3.0); err != nil {
				d.Nack(false, true)
				continue
			}

			d.Ack(false)
		}

	}()

}
