FROM golang:1.22-alpine AS builder
WORKDIR /usr/src/app
COPY . .
# Copy the file from one level up to the current working directory
RUN go build -o keygen ./cmd/main.go

FROM alpine:latest
COPY --from=builder /usr/src/app/keygen .
RUN chmod +x ./keygen
CMD ["./keygen"]