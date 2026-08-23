package main

import (
	"context"
	"log"
	"os"

	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/controllers"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/repositories"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/routes"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/services"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// r := gin.Default()

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 		"status":  "success!",
	// 	})
	// })

	loadEnv()

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_DB_ADDRESS"),
		Password: os.Getenv("REDIS_DB_PASSWORD"),
		DB:       0,
	})

	defer rdb.Close()

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic("Cannot connect to Reids" + err.Error())
	}

	locationRepo := repositories.NewLocationRepository(rdb)
	locationService := services.NewLocationService(locationRepo)
	locationController := controllers.NewLocationController(locationService)

	r := routes.SetupRouter(locationController)
	r.Run(":8081")
}
