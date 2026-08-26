package com.goride.trip_service.models.dto;

import lombok.Data;

@Data
public class TripRequestDto {
    private LocationDto pickupLocation;
    private LocationDto dropoffLocation;
}
