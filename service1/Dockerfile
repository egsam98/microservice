FROM golang:alpine as builder

RUN apk update && \
    apk add build-base && \
    apk add --no-cache git

WORKDIR /github.com/egsam98/microservice
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 go build -o service1-service *.go


FROM scratch

COPY --from=builder /github.com/egsam98/microservice .
EXPOSE 8080
ENTRYPOINT ["./service1-service"]