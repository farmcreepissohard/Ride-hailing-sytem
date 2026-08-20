package com.goride.trip_service.repositories;

import java.util.UUID;

import org.springframework.data.jpa.repository.JpaRepository;

import com.goride.trip_service.entities.TripEntity;

public interface TripRepository extends JpaRepository<TripEntity, UUID> {

}
