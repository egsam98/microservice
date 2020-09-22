package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"
	service1 "service1/proto/service1"
)

type Service1 struct{}

func (e *Service1) Handle(ctx context.Context, msg *service1.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *service1.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
