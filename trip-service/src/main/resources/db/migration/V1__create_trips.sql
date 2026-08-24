CREATE TABLE trips(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
    customer_id UUID NOT NULL,
    driver_id UUID,

    pickup_location TEXT NOT NULL,
    dropoff_location TEXT NOT NULL,
    pickup_lat DECIMAL(10,8) NOT NULL,
    pickup_lng DECIMAL(11,8) NOT NULL,
    dropoff_lat DECIMAL(10,8) NOT NULL,
    dropoff_lng DECIMAL(11,8) NOT NULL,
    
    distance DOUBLE PRECISION NOT NULL,
    
    total_amount DOUBLE PRECISION NOT NULL,
    driver_earning DOUBLE PRECISION NOT NULL,
    
    trip_status VARCHAR(20) NOT NULL DEFAULT 'PENDING',
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    accepted_at TIMESTAMPTZ,
    pickedup_at TIMESTAMPTZ,
    completed_at TIMESTAMPTZ,
    cancelled_at TIMESTAMPTZ,
    
    CONSTRAINT check_status CHECK (trip_status IN ('PENDING', 'ACCEPTED', 'ON_TRIP', 'COMPLETED', 'CANCELLED'))
);