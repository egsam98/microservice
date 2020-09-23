FROM microservice_builder:latest as builder
RUN CGO_ENABLED=0 go build -o bin/service1 service1/*.go


FROM scratch
COPY --from=builder /github.com/egsam98/microservice/bin/service1 .
ENTRYPOINT ["./service1"]