package com.goride.trip_service.core;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import lombok.Getter;

@Component
@Getter
public class ConfigurationProperties {
    @Value("${OSRM_URL}")
    private String osrmBaseUrl;

    @Value("${PRICE_PER_KM}")
    private double pricePerKm;

    @Value("${BASE_FARE}")
    private double baseFare;

    @Value("${DRIVER_COMMISSION}")
    private double driverCommission;
}
