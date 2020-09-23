FROM golang:alpine as builder

RUN apk update && \
    apk add build-base && \
    apk add --no-cache git

WORKDIR /github.com/egsam98/microservice
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 go build -o service2-service *.go


FROM scratch

COPY --from=builder /github.com/egsam98/microservice/service2-service .
EXPOSE 8080
ENTRYPOINT ["./service2-service"]