package handler

import (
	"context"
	"github.com/egsam98/microservice/service3/repository/response_repository"

	service3 "github.com/egsam98/microservice/service3/proto/service3"
)

type Service3 struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Service3) Telemetry(ctx context.Context, serialNum *service3.SerialNum, res *service3.Response) error {
	result, err := response_repository.FindBySerialNum(serialNum)
	if err != nil {
		return err
	}
	*res = *result
	return nil
}
