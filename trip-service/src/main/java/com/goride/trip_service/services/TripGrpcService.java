package com.goride.trip_service.services;

import java.util.Optional;
import java.util.UUID;

import org.springframework.grpc.server.service.GrpcService;

import com.goride.grpc.CancelTripRequest;
import com.goride.grpc.ResponseStatus;
import com.goride.grpc.TripActionRequest;
import com.goride.grpc.TripResponse;
import com.goride.grpc.UpdateTripInformationGrpc.UpdateTripInformationImplBase;
import com.goride.trip_service.repositories.TripRepository;

import io.grpc.stub.StreamObserver;
import lombok.RequiredArgsConstructor;

@GrpcService
@RequiredArgsConstructor
public class TripGrpcService extends UpdateTripInformationImplBase {

        private final TripRepository tripRepository;

        private void handleResponse(final Optional<UUID> optionalDbResponse,
                        final StreamObserver<TripResponse> responseObserver,
                        final String tripId,
                        final String errorMessage) {
                if (optionalDbResponse.isPresent()) {
                        final UUID dbResponse = optionalDbResponse.get();
                        TripResponse response = TripResponse.newBuilder()
                                        .setTripId(dbResponse.toString())
                                        .setStatus(ResponseStatus.SUCCESS)
                                        .build();

                        responseObserver.onNext(response);
                        responseObserver.onCompleted();
                } else {
                        TripResponse response = TripResponse.newBuilder()
                                        .setTripId(tripId)
                                        .setStatus(ResponseStatus.UNSUCCESS)
                                        .setMessage(errorMessage)
                                        .build();

                        responseObserver.onNext(response);
                        responseObserver.onCompleted();
                }
        }

        @Override
        public void matchTrip(TripActionRequest request, StreamObserver<TripResponse> responseObserver) {
                final Optional<UUID> optionalDbResponse = tripRepository.matchingTrip(
                                UUID.fromString(request.getTripId()),
                                UUID.fromString(request.getDriverId()));

                handleResponse(optionalDbResponse, responseObserver, request.getTripId(), "Trip not found");
        }

        public void startTrip(TripActionRequest request,
                        StreamObserver<com.goride.grpc.TripResponse> responseObserver) {
                final Optional<UUID> optionalDbResponse = tripRepository.startTrip(
                                UUID.fromString(request.getTripId()),
                                UUID.fromString(request.getDriverId()));

                handleResponse(optionalDbResponse, responseObserver, request.getTripId(), "Forbidden on starting trip");
        }

        public void completeTrip(TripActionRequest request,
                        StreamObserver<com.goride.grpc.TripResponse> responseObserver) {
                final Optional<UUID> optionalDbResponse = tripRepository.completeTrip(
                                UUID.fromString(request.getTripId()),
                                UUID.fromString(request.getDriverId()));

                handleResponse(optionalDbResponse, responseObserver, request.getTripId(),
                                "Forbidden on completing trip");
        }

        public void cancelTrip(CancelTripRequest request,
                        StreamObserver<com.goride.grpc.TripResponse> responseObserver) {
                final Optional<UUID> optionalDbResponse = tripRepository.cancelTrip(
                                UUID.fromString(request.getTripId()),
                                UUID.fromString(request.getIdCancelledBy()),
                                request.getReason());

                handleResponse(optionalDbResponse, responseObserver, request.getTripId(),
                                "Forbidden on cancelling trip");
        }
}
