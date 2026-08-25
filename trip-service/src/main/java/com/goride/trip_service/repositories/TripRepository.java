package com.goride.trip_service.repositories;

import java.util.UUID;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import com.goride.trip_service.entities.TripEntity;

public interface TripRepository extends JpaRepository<TripEntity, UUID> {

	@Modifying(clearAutomatically = true)
	@Query(value = """
			UPDATE trips
			SET driver_id = :driver_id,
			    trip_status = 'ACCEPTED',
			    accepted_at = NOW()
			WHERE id = :trip_id
			""", nativeQuery = true)
	int matchingTrip(
			@Param("trip_id") UUID tripId,
			@Param("driver_id") UUID driverId);

	@Modifying(clearAutomatically = true)
	@Query(value = """
			UPDATE trips
			SET trip_status = 'ON_TRIP',
			        pickedup_at = NOW()
			WHERE id = :trip_id AND driver_id = :driver_id AND trip_status = 'ACCEPTED'
			""", nativeQuery = true)
	int startTrip(
			@Param("trip_id") UUID tripId,
			@Param("driver_id") UUID driverId);

	@Modifying(clearAutomatically = true)
	@Query(value = """
			UPDATE trips
			SET trip_status = 'COMPLETED',
			        completed_at = NOW()
			WHERE id = :trip_id AND driver_id = :driver_id AND trip_status = 'ON_TRIP'
			""", nativeQuery = true)
	int completeTrip(
			@Param("trip_id") UUID tripId,
			@Param("driver_id") UUID driverId);

	@Modifying(clearAutomatically = true)
	@Query(value = """
			UPDATE trips
			SET trip_status = 'CANCELLED',
			        cancelled_at = NOW(),
			        cancelled_by = :cancelled_by,
			        cancelled_reason = :cancelled_reason
			WHERE id = :trip_id
			        AND (driver_id = :cancelled_by OR customer_id = :cancelled_by)
			        AND trip_status IN ('PENDING', 'ACCEPTED')
			""", nativeQuery = true)
	int cancelTrip(
			@Param("trip_id") UUID tripId,
			@Param("cancelled_by") UUID cancelledBy,
			@Param("cancelled_reason") String cancelledReason);

}
