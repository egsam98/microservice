package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	service2 "service2/proto/service2"
)

type Service2 struct{}

func (e *Service2) Handle(ctx context.Context, msg *service2.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *service2.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
