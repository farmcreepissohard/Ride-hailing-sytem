package com.goride.trip_service.services;

import java.math.BigDecimal;
import java.util.Locale;
import java.util.UUID;

import org.springframework.stereotype.Service;
import org.springframework.web.client.RestClient;

import com.goride.trip_service.core.ConfigurationProperties;
import com.goride.trip_service.dto.TripRequestDto;
import com.goride.trip_service.entities.TripEntity;
import com.goride.trip_service.repositories.TripRepository;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import tools.jackson.databind.JsonNode;

@Slf4j
@Service
@RequiredArgsConstructor
public class TripService {

    private final TripRepository tripRepository;
    private final ConfigurationProperties configurationProperties;

    private final RestClient restClient = RestClient.create();

    public String handleTripRequest(String customerId, TripRequestDto tripRequestDto) {

        log.info("wtf ok?");

        final double pickupLongitude = tripRequestDto.getPickupLocation().getLng();
        final double pickupLatitude = tripRequestDto.getPickupLocation().getLat();
        final double dropoffLongitude = tripRequestDto.getDropoffLocation().getLng();
        final double dropoffLatitude = tripRequestDto.getDropoffLocation().getLat();

        String url = String.format(Locale.US, "%s/%f,%f;%f,%f?overview=false", configurationProperties.getOsrmBaseUrl(),
                pickupLongitude, pickupLatitude,
                dropoffLongitude, dropoffLatitude);

        log.info("url: {}", url);

        try {
            final JsonNode response = restClient
                    .get()
                    .uri(url)
                    .header("User-Agent", "GoRide-TripService/1.0")
                    .header("Accept-Encoding", "identity")
                    .header("Accept", "application/json")
                    .retrieve()
                    .body(JsonNode.class);

            if (response != null && response.has("routes") && response.get("routes").size() > 0) {
                double distanceInMeters = response.get("routes").get(0).get("distance").asDouble();
                double distanceInKm = distanceInMeters / 1000.0;

                BigDecimal totalAmount = BigDecimal.valueOf(configurationProperties.getBaseFare())
                        .add(BigDecimal.valueOf(Math.max(0, distanceInKm - 3)).multiply(
                                BigDecimal.valueOf(configurationProperties.getPricePerKm())));
                BigDecimal driverEarning = totalAmount.multiply(
                        BigDecimal.valueOf(configurationProperties.getDriverCommission()))
                        .divide(BigDecimal.valueOf(100));

                log.info("Distance: {} km, Total: {}, Driver earning: {}", distanceInKm, totalAmount, driverEarning);

                final TripEntity trip = tripRepository.save(new TripEntity(
                        UUID.fromString(customerId),
                        tripRequestDto.getPickupLocation().getAddress(),
                        tripRequestDto.getDropoffLocation().getAddress(),
                        pickupLatitude,
                        pickupLongitude,
                        dropoffLatitude,
                        dropoffLongitude,
                        distanceInKm,
                        totalAmount,
                        driverEarning));

                return trip.getId().toString();
            } else {
                log.error("No Route founded");
                throw new RuntimeException("No route found");
            }
        } catch (Exception e) {
            log.error("Api calling failed", e);
            throw new RuntimeException("Can not handle trip now");
        }
    }
}
