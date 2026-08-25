package payload

type TripEventPayload struct {
	TripID    string  `json:"tripId"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
