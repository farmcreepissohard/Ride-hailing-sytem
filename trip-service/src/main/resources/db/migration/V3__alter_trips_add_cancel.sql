ALTER TABLE trips
ADD COLUMN cancelled_by UUID,
ADD COLUMN cancelled_reason TEXT;