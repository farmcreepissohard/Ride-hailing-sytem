package com.goride.trip_service.utils;

import java.util.Random;

import com.goride.trip_service.models.dto.LocationDto;

public class LocationRandom {
    private static final Random random = new Random();

    public static double getRandomLatitude() {
        return -90 + (random.nextDouble() * 180);
    }

    public static double getRandomLongitude() {
        return -180 + (random.nextDouble() * 360);
    }

    public static LocationDto getRandomLocationDto(final String locationName) {
        return new LocationDto(locationName, getRandomLatitude(), getRandomLongitude());
    }
}
