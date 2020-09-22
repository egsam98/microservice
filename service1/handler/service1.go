package handler

import (
	"context"
	service1 "service1/proto/service1"

	"time"
)

type Service1 struct{}

var _ service1.Service1Handler = &Service1{}

var datetime = time.Date(2020, 9, 21, 12, 0, 01, 0, time.UTC).
	Format("02-01-2006,15:04:05")
var response = service1.Response{
	SerialNum: "DDD_B1:B2:15:E2",
	Packet: &service1.Packet{
		Datetime: datetime,
		Temp1:    22.33,
		Temp2:    33.42,
		Height:   223.4,
		Latitude: 33.44122345,
		Logitude: 43.43321233,
		Status:   1,
	},
}

func (e *Service1) Json(ctx context.Context, in *service1.Empty, res *service1.Response) error {
	*res = response
	return nil
}
