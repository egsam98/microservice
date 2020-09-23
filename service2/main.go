package main

import (
	"github.com/egsam98/microservice/service2/publisher"
	"github.com/joho/godotenv"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"os"
)

func main() {
	_ = godotenv.Load()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.service2"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	go publisher.Run(os.Getenv("TOPIC_NAME"))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
