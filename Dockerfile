FROM golang:alpine

RUN apk update && \
    apk add build-base && \
    apk add --no-cache git

WORKDIR /github.com/egsam98/microservice

ENV GO111MODULE=on

COPY . .