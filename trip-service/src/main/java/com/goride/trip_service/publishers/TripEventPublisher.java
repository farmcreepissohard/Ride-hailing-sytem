package com.goride.trip_service.publishers;

import java.util.UUID;

import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import com.goride.trip_service.payload.TripEventPayload;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;

@Slf4j
@Component
@RequiredArgsConstructor
public class TripEventPublisher {
    private final RabbitTemplate rabbitTemplate;

    @Value("${TRIP_CREATED_QUEUE}")
    private final String tripCreatedQueue;

    public void publishTripRequest(final UUID tripId, final double longitude, final double latitude) {
        final TripEventPayload message = new TripEventPayload(tripId.toString(), longitude, latitude);
        log.info("[RabbitMQ] send trip request: {}", message);
        rabbitTemplate.convertAndSend(tripCreatedQueue, message);
    }
}
