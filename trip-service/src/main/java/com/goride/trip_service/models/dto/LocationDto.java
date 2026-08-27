package com.goride.trip_service.models.dto;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class LocationDto {
    private String address;
    private double lat;
    private double lng;
}
