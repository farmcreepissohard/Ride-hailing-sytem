package main

import (
	"context"
	"log"
	"os"

	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/consumers"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/controllers"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/grpc_client"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/repositories"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/routes"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/services"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/ws"
	"github.com/farmcreepissohard/Ride-hailing-sytem/pkg/pb"
	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnv()

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_DB_ADDRESS"),
		Password: os.Getenv("REDIS_DB_PASSWORD"),
		DB:       0,
	})
	defer rdb.Close()

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		panic("Cannot connect to Reids " + err.Error())
	}

	locationRepo := repositories.NewLocationRepository(rdb)

	//-----------------

	grpcConn, err := grpc.NewClient(os.Getenv("GRPC_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer grpcConn.Close()
	if err != nil {
		panic("Cannot connect to grpc " + err.Error())
	}

	grpcClient := grpc_client.NewTripGrpcClient(pb.NewUpdateTripInformationClient(grpcConn), grpcConn)
	locationService := services.NewLocationService(grpcClient, locationRepo)
	dispatchService := services.NewDispatchService(locationRepo, rdb)

	locationController := controllers.NewLocationController(locationService, dispatchService)

	//-----------------

	hub := ws.NewHub(rdb)
	wsController := controllers.NewWsController(hub, locationService)

	//-----------------

	conn, err := amqp091.Dial(os.Getenv("AMQP_URL"))
	defer conn.Close()
	if err != nil {
		panic("Failed to connect to RabbitMQ: " + err.Error())
	}

	ch, err := conn.Channel()
	defer ch.Close()
	if err != nil {
		panic("Failed to open a channel: " + err.Error())
	}

	queueName := os.Getenv("TRIP_CREATED_QUEUE")
	if _, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		panic("Failed to declare a queue: " + err.Error())
	}

	tripEventHandler := consumers.NewTripEventConsumer(dispatchService)
	tripEventHandler.Consume(ch, queueName)

	r := routes.SetupRouter(routes.RouterDependency{LocationController: locationController, WsController: wsController})
	r.Run(":8081")
}
