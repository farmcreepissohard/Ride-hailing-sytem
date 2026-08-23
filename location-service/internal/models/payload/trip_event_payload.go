package payload

type TripEventPayload struct {
	TripID    string  `json:"trip_id"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
