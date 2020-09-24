package main

import (
	"github.com/egsam98/microservice/service3/db"
	"github.com/egsam98/microservice/service3/handler"
	service3 "github.com/egsam98/microservice/service3/proto/service3"
	"github.com/egsam98/microservice/service3/subscriber"
	"github.com/joho/godotenv"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"os"
)

func main() {
	_ = godotenv.Load()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.service3"),
		micro.Version("latest"),
		micro.AfterStart(func() error { return db.Init(os.Getenv("SQLITE_DATABASE")) }),
	)

	// Initialise service
	service.Init()

	// Register Handler
	service3.RegisterService3Handler(service.Server(), new(handler.Service3))

	// Register Struct as Subscriber
	micro.RegisterSubscriber(os.Getenv("TOPIC_NAME"), service.Server(), new(subscriber.Service3))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
