package dto

import (
	"encoding/json"

	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/models/enum"
)

type WebsocketRequest struct {
	Action  enum.WebsocketActionEnum `json:"action"`
	Payload json.RawMessage          `json:"payload"`
}
