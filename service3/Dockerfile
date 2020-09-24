FROM microservice_builder:latest as builder
RUN CGO_ENABLED=0 go build -o bin/service3 service3/*.go


FROM scratch
COPY --from=builder /github.com/egsam98/microservice/bin/service3 .
ENTRYPOINT ["./service3"]