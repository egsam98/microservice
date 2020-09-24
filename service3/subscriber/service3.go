package subscriber

import (
	"context"
	gomicroapiservice1 "github.com/egsam98/microservice/service1/proto/service1"
	"github.com/egsam98/microservice/service3/repository/response_repository"
	"log"
)

type Service3 struct{}

func (e *Service3) Handle(_ context.Context, res *gomicroapiservice1.Response) error {
	log.Printf("Received %v\n", res)

	if err := response_repository.Create(res); err != nil {
		log.Println(err)
		return err
	}

	log.Printf("Response %v has been added to database", res)
	return nil
}
