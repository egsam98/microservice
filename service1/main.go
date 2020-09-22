package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"service1/handler"
	service1 "service1/proto/service1"
	"service1/subscriber"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.service1"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	service1.RegisterService1Handler(service.Server(), new(handler.Service1))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.api.service1", service.Server(), new(subscriber.Service1))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
