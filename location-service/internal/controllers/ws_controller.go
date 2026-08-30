package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/models/dto"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/models/enum"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/services"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/transport"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgraded = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WsController struct {
	hub             *transport.Hub
	locationService services.LocationService
}

func NewWsController(hub *transport.Hub, locationService services.LocationService) *WsController {
	return &WsController{hub: hub, locationService: locationService}
}

func (c *WsController) HandleConnection(ctx *gin.Context) {
	conn, err := upgraded.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("Failed in Upgrade:", err)
		return
	}
	defer conn.Close()

	driverId := ctx.Query("user_id")
	c.hub.Add(driverId, conn)
	defer c.hub.Remove(driverId, conn)

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
				if err := c.locationService.UpdateLocation(driverId, locationRequestDto.Longitude, locationRequestDto.Latitude); err != nil {
					log.Println("Update location failed: ", err.Error())
					break
				}
			}

		}
	}
}
