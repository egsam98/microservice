FROM microservice_builder:latest as builder
RUN go build -o bin/service3 service3/*.go

ENTRYPOINT ["/github.com/egsam98/microservice/bin/service3"]