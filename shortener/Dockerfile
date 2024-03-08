FROM golang:1.22-alpine AS builder
ARG SHORTENER_SERVICE_ADDRESS
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./shortener-service ./cmd/server/main.go
RUN go build -ldflags  "-X main.shortenerServiceAddr=$SHORTENER_SERVICE_ADDRESS" -o ./gateway-service ./cmd/client/main.go

FROM alpine:latest AS shortener-service
WORKDIR /app
COPY --from=builder /app/shortener-service .
EXPOSE 50051
RUN chmod +x ./shortener-service
ENTRYPOINT ["./shortener-service"]

FROM alpine:latest AS gateway-service
WORKDIR /app
COPY --from=builder /app/gateway-service .
EXPOSE 8080
RUN chmod +x ./gateway-service
ENTRYPOINT ["./gateway-service"]