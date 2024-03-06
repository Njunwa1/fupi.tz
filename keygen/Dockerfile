FROM golang:1.18 AS builder
WORKDIR /usr/src/app
COPY . .
# Copy the file from one level up to the current working directory
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o keygen ./cmd/main.go

FROM scratch
COPY --from=builder /usr/src/app/keygen ./keygen
CMD ["./keygen"]