FROM microservice_builder:latest as builder
RUN CGO_ENABLED=0 go build -o bin/service2 service2/*.go


FROM scratch
COPY --from=builder /github.com/egsam98/microservice/bin/service2 .
ENTRYPOINT ["./service2"]