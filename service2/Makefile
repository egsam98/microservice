
GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:

	go build -o service2-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t service2-service:latest
