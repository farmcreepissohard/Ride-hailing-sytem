package com.goride.trip_service.services;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.verifyNoInteractions;
import static org.mockito.Mockito.when;

import java.util.UUID;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import com.goride.trip_service.core.ConfigurationProperties;
import com.goride.trip_service.models.dto.LocationDto;
import com.goride.trip_service.models.dto.TripRequestDto;
import com.goride.trip_service.publishers.TripEventPublisher;
import com.goride.trip_service.repositories.TripRepository;
import com.goride.trip_service.utils.LocationRandom;

@ExtendWith(MockitoExtension.class)
public class TripServiceImplTest {

    @Mock
    private TripRepository tripRepository;

    @Mock
    private TripEventPublisher eventPublisher;

    @Mock
    private ConfigurationProperties configurationProperties;

    @InjectMocks
    private TripServiceImpl tripService;

    // @Nested
    // @DisplayName("handleTripRequest() function test")
    // class HandlTripRequestTests {

    // // @Test
    // // @DisplayName("Successful")
    // // void success() {
    // // final LocationDto pickupLocation =
    // // LocationRandom.getRandomLocationDto("pickup-location");
    // // final LocationDto dropoffLocation =
    // // LocationRandom.getRandomLocationDto("dropoff-location");
    // // final TripRequestDto req = new TripRequestDto(pickupLocation,
    // // dropoffLocation);
    // // final UUID customerId = UUID.randomUUID();

    // //
    // when(configurationProperties.getOsrmBaseUrl()).thenReturn("http://mock-osrm-url.com");

    // // assertDoesNotThrow(() ->
    // tripService.handleTripRequest(customerId.toString(),
    // // req));

    // // verify(eventPublisher, times(1)).publishTripRequest(customerId,
    // // pickupLocation.getLng(),
    // // pickupLocation.getLat());
    // // }
    // }

    @Nested
    @DisplayName("matchTrip() function test")
    class MatchTripTests {

        @Test
        @DisplayName("Successful")
        void success() {
            final UUID tripId = UUID.randomUUID();
            final UUID driverId = UUID.randomUUID();
            when(tripRepository.matchingTrip(tripId, driverId)).thenReturn(1);
            assertTrue(tripService.matchTrip(tripId.toString(), driverId.toString()));
            verify(tripRepository, times(1)).matchingTrip(tripId, driverId);
        }

        @Test
        @DisplayName("Failed when no available trip")
        void shouldFailedWhenNoAvailableTrip() {
            final UUID tripId = UUID.randomUUID();
            final UUID driverId = UUID.randomUUID();
            when(tripRepository.matchingTrip(tripId, driverId)).thenReturn(0);
            assertFalse(tripService.matchTrip(tripId.toString(), driverId.toString()));
            verify(tripRepository, times(1)).matchingTrip(tripId, driverId);
        }

        @Test
        @DisplayName("Failed when empty tripId")
        void shouldFailedWhenEmptyTripId() {
            final String tripId = "";
            final UUID driverId = UUID.randomUUID();
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.matchTrip(tripId, driverId.toString()));
            assertEquals("Trip id is required", exception.getMessage());
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when null tripId")
        void shouldFailedWhenNullTripId() {
            final String tripId = null;
            final UUID driverId = UUID.randomUUID();
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.matchTrip(tripId, driverId.toString()));
            assertEquals("Trip id is required", exception.getMessage());
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when invalid UUID tripId")
        void shouldFailedWhenInvalidTripId() {
            final String tripId = "asdf";
            final UUID driverId = UUID.randomUUID();
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.matchTrip(tripId, driverId.toString()));
            assertTrue(exception.getMessage().contains("Invalid UUID string"));
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when empty driverId")
        void shouldFailedWhenEmptyDriverId() {
            final UUID tripId = UUID.randomUUID();
            final String driverId = "";
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.matchTrip(tripId.toString(), driverId));
            assertEquals("Driver id is required", exception.getMessage());
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when null driverId")
        void shouldFailedWhenNullDriverId() {
            final UUID tripId = UUID.randomUUID();
            final String driverId = null;
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.matchTrip(tripId.toString(), driverId));
            assertEquals("Driver id is required", exception.getMessage());
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when invalid UUID driverId")
        void shouldFailedWhenInvalidDriverId() {
            final UUID tripId = UUID.randomUUID();
            final String driverId = "asdf";
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.matchTrip(tripId.toString(), driverId));
            assertTrue(exception.getMessage().contains("Invalid UUID string"));
            verifyNoInteractions(tripRepository);
        }
    }

    @Nested
    @DisplayName("startTrip() function test")
    class StartTripTests {

        @Test
        @DisplayName("Successful")
        void success() {
            final UUID tripId = UUID.randomUUID();
            final UUID driverId = UUID.randomUUID();
            when(tripRepository.startTrip(tripId, driverId)).thenReturn(1);
            assertTrue(tripService.startTrip(tripId.toString(), driverId.toString()));
            verify(tripRepository, times(1)).startTrip(tripId, driverId);
        }

        @Test
        @DisplayName("Failed when no modifying in database")
        void shouldFailedWhenDatabaseNoUpdate() {
            final UUID tripId = UUID.randomUUID();
            final UUID driverId = UUID.randomUUID();
            when(tripRepository.startTrip(tripId, driverId)).thenReturn(0);
            assertFalse(tripService.startTrip(tripId.toString(), driverId.toString()));
            verify(tripRepository, times(1)).startTrip(tripId, driverId);
        }

        @Test
        @DisplayName("Failed when empty tripId")
        void shouldFailedWhenEmptyTripId() {
            final String tripId = "";
            final UUID driverId = UUID.randomUUID();
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.startTrip(tripId, driverId.toString()));
            assertEquals("Trip id is required", exception.getMessage());
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when null tripId")
        void shouldFailedWhenNullTripId() {
            final String tripId = null;
            final UUID driverId = UUID.randomUUID();
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.startTrip(tripId, driverId.toString()));
            assertEquals("Trip id is required", exception.getMessage());
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when invalid UUID tripId")
        void shouldFailedWhenInvalidTripId() {
            final String tripId = "asdf";
            final UUID driverId = UUID.randomUUID();
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.startTrip(tripId, driverId.toString()));
            assertTrue(exception.getMessage().contains("Invalid UUID string"));
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when empty driverId")
        void shouldFailedWhenEmptyDriverId() {
            final UUID tripId = UUID.randomUUID();
            final String driverId = "";
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.startTrip(tripId.toString(), driverId));
            assertEquals("Driver id is required", exception.getMessage());
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when null driverId")
        void shouldFailedWhenNullDriverId() {
            final UUID tripId = UUID.randomUUID();
            final String driverId = null;
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.startTrip(tripId.toString(), driverId));
            assertEquals("Driver id is required", exception.getMessage());
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when invalid UUID driverId")
        void shouldFailedWhenInvalidDriverId() {
            final UUID tripId = UUID.randomUUID();
            final String driverId = "asdf";
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.startTrip(tripId.toString(), driverId));
            assertTrue(exception.getMessage().contains("Invalid UUID string"));
            verifyNoInteractions(tripRepository);
        }
    }

    @Nested
    @DisplayName("completeTrip() function test")
    class CompleteTripTests {

        @Test
        @DisplayName("Successful")
        void success() {
            final UUID tripId = UUID.randomUUID();
            final UUID driverId = UUID.randomUUID();
            when(tripRepository.completeTrip(tripId, driverId)).thenReturn(1);
            assertTrue(tripService.completeTrip(tripId.toString(), driverId.toString()));
            verify(tripRepository, times(1)).completeTrip(tripId, driverId);
        }

        @Test
        @DisplayName("Failed when no modifying in database")
        void shouldFailedWhenDatabaseNoUpdate() {
            final UUID tripId = UUID.randomUUID();
            final UUID driverId = UUID.randomUUID();
            when(tripRepository.completeTrip(tripId, driverId)).thenReturn(0);
            assertFalse(tripService.completeTrip(tripId.toString(), driverId.toString()));
            verify(tripRepository, times(1)).completeTrip(tripId, driverId);
        }

        @Test
        @DisplayName("Failed when empty tripId")
        void shouldFailedWhenEmptyTripId() {
            final String tripId = "";
            final UUID driverId = UUID.randomUUID();
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.completeTrip(tripId, driverId.toString()));
            assertEquals("Trip id is required", exception.getMessage());
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when null tripId")
        void shouldFailedWhenNullTripId() {
            final String tripId = null;
            final UUID driverId = UUID.randomUUID();
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.completeTrip(tripId, driverId.toString()));
            assertEquals("Trip id is required", exception.getMessage());
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when invalid UUID tripId")
        void shouldFailedWhenInvalidTripId() {
            final String tripId = "asdf";
            final UUID driverId = UUID.randomUUID();
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.completeTrip(tripId, driverId.toString()));
            assertTrue(exception.getMessage().contains("Invalid UUID string"));
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when empty driverId")
        void shouldFailedWhenEmptyDriverId() {
            final UUID tripId = UUID.randomUUID();
            final String driverId = "";
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.completeTrip(tripId.toString(), driverId));
            assertEquals("Driver id is required", exception.getMessage());
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when null driverId")
        void shouldFailedWhenNullDriverId() {
            final UUID tripId = UUID.randomUUID();
            final String driverId = null;
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.completeTrip(tripId.toString(), driverId));
            assertEquals("Driver id is required", exception.getMessage());
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when invalid UUID driverId")
        void shouldFailedWhenInvalidDriverId() {
            final UUID tripId = UUID.randomUUID();
            final String driverId = "asdf";
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.completeTrip(tripId.toString(), driverId));
            assertTrue(exception.getMessage().contains("Invalid UUID string"));
            verifyNoInteractions(tripRepository);
        }
    }

    @Nested
    @DisplayName("cancelTrip() function test")
    class CancelTripTests {

        @Test
        @DisplayName("Successful")
        void success() {
            final UUID tripId = UUID.randomUUID();
            final UUID cancelledBy = UUID.randomUUID();
            final String reason = "hello world";
            when(tripRepository.cancelTrip(tripId, cancelledBy, reason)).thenReturn(1);
            assertTrue(tripService.cancelTrip(tripId.toString(), cancelledBy.toString(), reason));
            verify(tripRepository, times(1)).cancelTrip(tripId, cancelledBy, reason);
        }

        @Test
        @DisplayName("Failed when no modifying in database")
        void shouldFailedWhenDatabaseNoUpdate() {
            final UUID tripId = UUID.randomUUID();
            final UUID cancelledBy = UUID.randomUUID();
            final String reason = "hello world";
            when(tripRepository.cancelTrip(tripId, cancelledBy, reason)).thenReturn(0);
            assertFalse(tripService.cancelTrip(tripId.toString(), cancelledBy.toString(), reason));
            verify(tripRepository, times(1)).cancelTrip(tripId, cancelledBy, reason);
        }

        @Test
        @DisplayName("Failed when empty tripId")
        void shouldFailedWhenEmptyTripId() {
            final String tripId = "";
            final UUID cancelledBy = UUID.randomUUID();
            final String reason = "hello world";
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.cancelTrip(tripId, cancelledBy.toString(), reason));
            assertEquals("Trip id is required", exception.getMessage());
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when null tripId")
        void shouldFailedWhenNullTripId() {
            final String tripId = null;
            final UUID cancelledBy = UUID.randomUUID();
            final String reason = "hello world";
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.cancelTrip(tripId, cancelledBy.toString(), reason));
            assertEquals("Trip id is required", exception.getMessage());
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when invalid UUID tripId")
        void shouldFailedWhenInvalidTripId() {
            final String tripId = "asdf";
            final UUID cancelledBy = UUID.randomUUID();
            final String reason = "hello world";
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.cancelTrip(tripId, cancelledBy.toString(), reason));
            assertTrue(exception.getMessage().contains("Invalid UUID string"));
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when empty canceller")
        void shouldFailedWhenEmptyCanceller() {
            final UUID tripId = UUID.randomUUID();
            final String cancelledBy = "";
            final String reason = "hello world";
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.cancelTrip(tripId.toString(), cancelledBy, reason));
            assertEquals("Canceller id is required", exception.getMessage());
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when null canceller")
        void shouldFailedWhenNullCanceller() {
            final UUID tripId = UUID.randomUUID();
            final String cancelledBy = null;
            final String reason = "hello world";
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.cancelTrip(tripId.toString(), cancelledBy, reason));
            assertEquals("Canceller id is required", exception.getMessage());
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when invalid UUID canceller")
        void shouldFailedWhenInvalidCanceller() {
            final UUID tripId = UUID.randomUUID();
            final String cancelledBy = "asdf";
            final String reason = "hello world";
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.cancelTrip(tripId.toString(), cancelledBy, reason));
            assertTrue(exception.getMessage().contains("Invalid UUID string"));
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when empty reason")
        void shouldFailedWhenEmptyReason() {
            final UUID tripId = UUID.randomUUID();
            final UUID cancelledBy = UUID.randomUUID();
            final String reason = "";
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.cancelTrip(tripId.toString(), cancelledBy.toString(), reason));
            assertEquals("Unknown reason", exception.getMessage());
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when null reason")
        void shouldFailedWhenNullReason() {
            final UUID tripId = UUID.randomUUID();
            final UUID cancelledBy = UUID.randomUUID();
            final String reason = null;
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.cancelTrip(tripId.toString(), cancelledBy.toString(), reason));
            assertEquals("Unknown reason", exception.getMessage());
            verifyNoInteractions(tripRepository);
        }
    }

    @Nested
    @DisplayName("noDriverFound() function tests")
    class NoDriverFoundTests {

        @Test
        @DisplayName("successful")
        void shouldSuccess() {
            final UUID tripId = UUID.randomUUID();
            when(tripRepository.noDriverFound(tripId)).thenReturn(1);
            assertTrue(tripService.noDriverFound(tripId.toString()));
            verify(tripRepository, times(1)).noDriverFound(tripId);
        }

        @Test
        @DisplayName("Failed when no modifying in database")
        void shouldFailedWhenDatabaseNoUpdate() {
            final UUID tripId = UUID.randomUUID();
            when(tripRepository.noDriverFound(tripId)).thenReturn(0);
            assertFalse(tripService.noDriverFound(tripId.toString()));
            verify(tripRepository, times(1)).noDriverFound(tripId);
        }

        @Test
        @DisplayName("Failed when empty tripId")
        void shouldFailedWhenEmptyTripId() {
            final String tripId = "";
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.noDriverFound(tripId));
            assertEquals("Trip id is required", exception.getMessage());
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when null tripId")
        void shouldFailedWhenNullTripId() {
            final String tripId = null;
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.noDriverFound(tripId));
            assertEquals("Trip id is required", exception.getMessage());
            verifyNoInteractions(tripRepository);
        }

        @Test
        @DisplayName("Failed when invalid UUID tripId")
        void shouldFailedWhenInvalidTripId() {
            final String tripId = "asdf";
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> tripService.noDriverFound(tripId));
            assertTrue(exception.getMessage().contains("Invalid UUID string"));
            verifyNoInteractions(tripRepository);
        }
    }
}
