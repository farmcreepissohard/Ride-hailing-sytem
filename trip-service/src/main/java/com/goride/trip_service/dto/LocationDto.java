package com.goride.trip_service.dto;

import lombok.Data;

@Data
public class LocationDto {
    private String address;
    private double lat;
    private double lng;
}
