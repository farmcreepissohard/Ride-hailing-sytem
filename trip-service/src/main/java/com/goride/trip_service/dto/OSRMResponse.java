package com.goride.trip_service.dto;

import java.util.List;

import lombok.Data;

@Data
public class OSRMResponse {
    private List<Route> routes;

    @Data
    public static class Route {
        private double distance;
    }
}
