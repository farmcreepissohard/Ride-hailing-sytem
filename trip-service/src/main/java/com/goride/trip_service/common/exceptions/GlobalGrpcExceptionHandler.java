package com.goride.trip_service.common.exceptions;

import org.springframework.grpc.server.advice.GrpcAdvice;
import org.springframework.grpc.server.advice.GrpcExceptionHandler;

import io.grpc.Status;
import io.grpc.StatusRuntimeException;
import lombok.extern.slf4j.Slf4j;

@Slf4j
@GrpcAdvice
public class GlobalGrpcExceptionHandler {

    @GrpcExceptionHandler(BaseException.class)
    public StatusRuntimeException handleBaseExceptions(BaseException baseException) {
        log.warn("[gRPC]: Business Logic Error: {}", baseException.getMessage());

        Status status = Status.INTERNAL
                .withDescription(baseException.getErrorCode().getCode() + ": " + baseException.getMessage());
        return status.asRuntimeException();
    }

    @GrpcExceptionHandler(Exception.class)
    public StatusRuntimeException handleBaseExceptions(Exception exception) {
        log.error("[REST] The system experienced an unforeseen malfunction: ", exception);

        return Status.UNKNOWN.withDescription("[gRPC] The system is experiencing issues, please try again later!")
                .asRuntimeException();
    }
}
