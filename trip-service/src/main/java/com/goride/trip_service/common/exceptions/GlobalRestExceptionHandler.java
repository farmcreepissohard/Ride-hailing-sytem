package com.goride.trip_service.common.exceptions;

import java.time.LocalDateTime;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

import com.goride.trip_service.common.errors.ErrorCode;
import com.goride.trip_service.common.errors.ErrorResponse;

import lombok.extern.slf4j.Slf4j;

@Slf4j
@RestControllerAdvice
public class GlobalRestExceptionHandler {

    private ErrorResponse buildErrorResponse(ErrorCode errorCode) {
        return new ErrorResponse(LocalDateTime.now(), errorCode.getMessage(), errorCode.getHttpStatus().value());
    }

    @ExceptionHandler(BaseException.class)
    public ResponseEntity<ErrorResponse> handleBaseExceptions(BaseException baseException) {
        log.warn("[REST] Business Logic Error: {}", baseException.getMessage());

        ErrorCode errorCode = baseException.getErrorCode();
        return ResponseEntity.status(errorCode.getHttpStatus()).body(buildErrorResponse(errorCode));
    }

    @ExceptionHandler(Exception.class)
    public ResponseEntity<ErrorResponse> handleUnwantedException(Exception exception) {
        log.error("[REST] The system experienced an unforeseen malfunction: ", exception);

        ErrorResponse response = new ErrorResponse(LocalDateTime.now(),
                "[REST] The system is experiencing issues, please try again later!",
                HttpStatus.INTERNAL_SERVER_ERROR.value());
        return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body(response);
    }
}
