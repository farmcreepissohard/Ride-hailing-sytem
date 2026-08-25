package com.goride.trip_service.controllers.grpc;

import org.springframework.grpc.server.service.GrpcService;

import com.goride.grpc.CancelTripRequest;
import com.goride.grpc.ResponseStatus;
import com.goride.grpc.TripActionRequest;
import com.goride.grpc.TripResponse;
import com.goride.grpc.UpdateTripInformationGrpc.UpdateTripInformationImplBase;
import com.goride.trip_service.services.TripService;

import io.grpc.stub.StreamObserver;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;

@Slf4j
@GrpcService
@RequiredArgsConstructor
public class TripGrpcController extends UpdateTripInformationImplBase {
    private final TripService tripService;

    private void handleResponse(final boolean isSuccess,
            final StreamObserver<TripResponse> responseObserver,
            final String tripId,
            final String errorMessage) {

        TripResponse.Builder builder = TripResponse.newBuilder()
                .setTripId(tripId)
                .setStatus(isSuccess ? ResponseStatus.SUCCESS : ResponseStatus.UNSUCCESS);

        if (!isSuccess) {
            builder.setMessage(errorMessage);
        }

        responseObserver.onNext(builder.build());
        responseObserver.onCompleted();
    }

    @Override
    public void matchTrip(TripActionRequest request, StreamObserver<TripResponse> responseObserver) {
        try {
            final boolean isSuccess = tripService.matchTrip(request.getTripId(), request.getDriverId());
            handleResponse(isSuccess, responseObserver, request.getTripId(), "Trip not found");
        } catch (Exception e) {
            System.err.println("MatchTrip] Error: " + e.getMessage());
            e.printStackTrace();

            responseObserver.onError(
                    io.grpc.Status.UNKNOWN
                            .withDescription("Lỗi nội bộ Java: " + e.getMessage())
                            .asRuntimeException());
        }
    }

    public void startTrip(TripActionRequest request,
            StreamObserver<com.goride.grpc.TripResponse> responseObserver) {
        try {
            final boolean isSuccess = tripService.startTrip(request.getTripId(), request.getDriverId());
            handleResponse(isSuccess, responseObserver, request.getTripId(), "Forbidden on starting trip");
        } catch (Exception e) {
            System.err.println("StartTrip] Error: " + e.getMessage());
            e.printStackTrace();

            responseObserver.onError(
                    io.grpc.Status.UNKNOWN
                            .withDescription("Lỗi nội bộ Java: " + e.getMessage())
                            .asRuntimeException());
        }
    }

    public void completeTrip(TripActionRequest request,
            StreamObserver<com.goride.grpc.TripResponse> responseObserver) {
        try {
            final boolean isSuccess = tripService.comepleteTrip(request.getTripId(), request.getDriverId());
            handleResponse(isSuccess, responseObserver, request.getTripId(), "Forbidden on completing trip");
        } catch (Exception e) {
            System.err.println("CompleteTrip] Error: " + e.getMessage());
            e.printStackTrace();

            responseObserver.onError(
                    io.grpc.Status.UNKNOWN
                            .withDescription("Lỗi nội bộ Java: " + e.getMessage())
                            .asRuntimeException());
        }
    }

    public void cancelTrip(CancelTripRequest request,
            StreamObserver<com.goride.grpc.TripResponse> responseObserver) {
        try {
            final boolean isSuccess = tripService.cancelTrip(request.getTripId(), request.getIdCancelledBy(),
                    request.getReason());
            handleResponse(isSuccess, responseObserver, request.getTripId(), "Forbidden on cancelling trip");
        } catch (Exception e) {
            System.err.println("CancelTrip] Error: " + e.getMessage());
            e.printStackTrace();

            responseObserver.onError(
                    io.grpc.Status.UNKNOWN
                            .withDescription("Lỗi nội bộ Java: " + e.getMessage())
                            .asRuntimeException());
        }
    }
}
