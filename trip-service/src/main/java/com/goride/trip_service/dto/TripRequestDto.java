package com.goride.trip_service.dto;

import lombok.Data;

@Data
public class TripRequestDto {
    private LocationDto pickupLocation;
    private LocationDto dropoffLocation;
}
