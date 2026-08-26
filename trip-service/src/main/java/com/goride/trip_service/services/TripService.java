package com.goride.trip_service.services;

import com.goride.trip_service.models.dto.TripRequestDto;

public interface TripService {
    public String handleTripRequest(final String customerId,
            final TripRequestDto tripRequestDto);

    public boolean matchTrip(final String tripId, final String driverId);

    public boolean startTrip(final String tripId, final String driverId);

    public boolean comepleteTrip(final String tripId, final String driverId);

    public boolean cancelTrip(final String tripId, final String cancelledBy, final String reason);
}
