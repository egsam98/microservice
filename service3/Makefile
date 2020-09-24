
GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
    
	protoc --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/service3/service3.proto
    

.PHONY: build
build: proto

	go build -o service3 *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t service3:latest
