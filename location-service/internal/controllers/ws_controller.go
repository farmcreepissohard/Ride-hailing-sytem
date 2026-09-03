package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/models/dto"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/models/enum"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/services"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/ws"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgraded = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WsController struct {
	hub             *ws.Hub
	locationService services.LocationService
}

func NewWsController(hub *ws.Hub, locationService services.LocationService) *WsController {
	return &WsController{hub: hub, locationService: locationService}
}

func (wsController *WsController) HandleConnection(c *gin.Context) {
	conn, err := upgraded.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed in Upgrade:", err)
		return
	}
	defer conn.Close()

	driverId := c.Query("user_id")
	wsController.hub.Add(driverId, conn)
	defer wsController.hub.Remove(driverId, conn)

	for {
		var req dto.WebsocketRequest

		if err := conn.ReadJSON(&req); err != nil {
			log.Println("Reading Json failed: ", err.Error())
			break
		}

		switch req.Action {
		case enum.UpdateLocation:
			{
				payload := req.Payload
				var locationRequestDto dto.LocationRequest
				if err := json.Unmarshal(payload, &locationRequestDto); err != nil {
					log.Println("Bad json format", err.Error())
					break
				}
				if err := wsController.locationService.UpdateLocation(driverId, locationRequestDto.Longitude, locationRequestDto.Latitude); err != nil {
					log.Println("Update location failed: ", err.Error())
					break
				}
			}

		}
	}
}
