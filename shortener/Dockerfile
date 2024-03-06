#FROM golang:1.18 AS builder
#WORKDIR /usr/src/app
#COPY . .
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shortener ./cmd/server/main.go
#
#FROM scratch
#COPY --from=builder /usr/src/app/shortner ./shortner
#CMD ["./shortner"]

FROM golang:1.19-alpine AS builder
ARG SHORTENER_SERVICE_ADDRESS
WORKDIR /app
COPY . .
# Copy the file from one level up to the current working directory
RUN go build -o ./shortener-service ./cmd/server/main.go
RUN go build -ldflags  "-X main.shortenerServiceAddr=$SHORTENER_SERVICE_ADDRESS" -o ./gateway-service ./cmd/client/main.go
FROM alpine:latest AS shortener-service
WORKDIR /app
COPY --from=builder /app/shortener-service .
EXPOSE 50051
ENTRYPOINT ["./orders-service"]
FROM alpine:latest AS gateway-service
WORKDIR /app
COPY --from=builder /app/gateway-service .
EXPOSE 8080
ENTRYPOINT ["./gateway-service"]