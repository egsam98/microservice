package publisher

import (
	"context"
	go_micro_api_service1 "github.com/egsam98/microservice/service1/proto/service1"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"log"
	"os"
	"strconv"
	"time"
)

var service1 = go_micro_api_service1.NewService1Service("go.micro.api.service1", client.DefaultClient)
var latencySec time.Duration

func Run(topic string) {
	if err := setupLatency(); err != nil {
		log.Fatal("LATENCY_SEC environment must be set (number)")
	}

	log.Printf("Publisher to \"%s\" topic with latency %ds has run\n", topic, latencySec)
	ctx := context.TODO()

	for {
		event := micro.NewEvent(topic, client.DefaultClient)
		result, err := service1.Json(ctx, &go_micro_api_service1.Empty{})
		if err != nil {
			log.Println(err)
			wait()
			continue
		}

		if err := event.Publish(ctx, result); err != nil {
			log.Println(err)
			wait()
			continue
		}

		log.Printf("Event %v has published to \"messages\" topic\n", result)
		wait()
	}
}

func setupLatency() error {
	result, err := strconv.Atoi(os.Getenv("LATENCY_SEC"))
	if err != nil {
		return err
	}
	latencySec = time.Duration(result)
	return nil
}

func wait() {
	time.Sleep(latencySec * time.Second)
}
