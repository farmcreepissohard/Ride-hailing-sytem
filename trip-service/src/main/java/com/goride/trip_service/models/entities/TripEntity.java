package com.goride.trip_service.models.entities;

import java.math.BigDecimal;
import java.time.OffsetDateTime;
import java.util.UUID;

import org.hibernate.annotations.CreationTimestamp;

import com.goride.trip_service.enums.TripStatusEnum;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.EnumType;
import jakarta.persistence.Enumerated;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.Table;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Entity
@Table(name = "trips")
@Getter
@NoArgsConstructor
public class TripEntity {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private UUID id;

    @Column(name = "customer_id", nullable = false)
    private UUID customerId;

    @Column(name = "driver_id")
    @Setter
    private UUID driverId;

    @Column(name = "pickup_location", nullable = false, columnDefinition = "TEXT")
    private String pickupLocation;

    @Column(name = "dropoff_location", nullable = false, columnDefinition = "TEXT")
    private String dropoffLocation;

    @Column(name = "pickup_lat", nullable = false)
    private Double pickupLat;

    @Column(name = "pickup_lng", nullable = false)
    private Double pickupLng;

    @Column(name = "dropoff_lat", nullable = false)
    private Double dropoffLat;

    @Column(name = "dropoff_lng", nullable = false)
    private Double dropoffLng;

    @Column(nullable = false)
    private Double distance;

    @Column(name = "total_amount", nullable = false)
    private BigDecimal totalAmount;

    @Column(name = "driver_earning", nullable = false)
    private BigDecimal driverEarning;

    @Column(name = "trip_status", nullable = false, length = 20)
    @Enumerated(EnumType.STRING)
    private TripStatusEnum tripStatus = TripStatusEnum.PENDING;

    @Column(name = "created_at", nullable = false, updatable = false)
    @CreationTimestamp
    private OffsetDateTime createdAt;

    @Column(name = "accepted_at")
    @Setter
    private OffsetDateTime acceptedAt;

    @Column(name = "pickedup_at")
    @Setter
    private OffsetDateTime pickedupAt;

    @Column(name = "completed_at")
    @Setter
    private OffsetDateTime completedAt;

    @Column(name = "cancelled_at")
    @Setter
    private OffsetDateTime cancelledAt;

    @Column(name = "cancelled_by")
    @Setter
    private UUID cancelledBy;

    @Column(name = "cancelled_reason", columnDefinition = "TEXT")
    @Setter
    private String cancelledReason;

    public TripEntity(UUID customerId, String pickupAddress, String dropoffAddress,
            double pickupLatitude, double pickupLongitude, double dropoffLatitude, double dropoffLongitude,
            double distanceInKm, BigDecimal totalAmount, BigDecimal driverEarning) {
        this.customerId = customerId;
        this.pickupLocation = pickupAddress;
        this.dropoffLocation = dropoffAddress;
        this.pickupLat = pickupLatitude;
        this.pickupLng = pickupLongitude;
        this.dropoffLat = dropoffLatitude;
        this.dropoffLng = dropoffLongitude;
        this.distance = distanceInKm;
        this.totalAmount = totalAmount;
        this.driverEarning = driverEarning;
    }

}
