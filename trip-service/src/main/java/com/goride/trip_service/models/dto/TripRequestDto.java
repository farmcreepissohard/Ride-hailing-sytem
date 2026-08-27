package com.goride.trip_service.models.dto;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class TripRequestDto {
    private LocationDto pickupLocation;
    private LocationDto dropoffLocation;
}
