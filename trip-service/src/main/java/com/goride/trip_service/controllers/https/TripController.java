package com.goride.trip_service.controllers.https;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.goride.trip_service.dto.TripRequestDto;
import com.goride.trip_service.services.TripService;

import lombok.RequiredArgsConstructor;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;

@RestController
@RequestMapping("/api/v1/trips")
@RequiredArgsConstructor
public class TripController {

    private final TripService tripService;

    @PostMapping("/request")
    public String postMethodName(@RequestHeader("X-User-Id") String customerId,
            @RequestBody TripRequestDto tripRequestDto) {
        final String tripId = tripService.handleTripRequest(customerId, tripRequestDto);
        return tripId;
    }

}
