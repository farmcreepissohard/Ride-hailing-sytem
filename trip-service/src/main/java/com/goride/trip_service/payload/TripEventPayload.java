package com.goride.trip_service.payload;

import lombok.Data;

@Data
public class TripEventPayload {
    final String tripId;
    final double longitude;
    final double latitude;

    public TripEventPayload(final String tripId, final double longitude, final double latitude) {
        this.tripId = tripId;
        this.longitude = longitude;
        this.latitude = latitude;
    }
}
