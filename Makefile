
install-protoc-gen:
	@echo "Installing protocol buffer compiler..."
	@brew install protobuf
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@echo "Done."

grpc-gateway:
	@echo "Get gRPC gateway libs..."
	@go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	@go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	@echo "Done."

protoc:
	@echo "Generating protocol buffer files..."
	@cd proto && protoc --go_out=../proto/golang --go_opt=paths=source_relative \
              --go-grpc_out=../proto/golang --go-grpc_opt=paths=source_relative \
              --grpc-gateway_out=../proto/golang --grpc-gateway_opt paths=source_relative \
              --grpc-gateway_opt generate_unbound_methods=true \
              ./**/*.proto
	@echo "Done."

rabbit-start:
	@echo "Starting RabbitMQ..."
	@docker run -d --hostname fupitz-rabbitmq --name fupitz-rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq
	@echo "Done."

redis-start:
	@echo "Starting Redis..."
	@docker run -d --name fupitz-redis -p 6379:6379 redis
	@echo "Done."

jaegar-start:
	@echo "Starting Jaegar..."
	@docker run -d --name jaeger \
       -e METRICS_STORAGE_TYPE=prometheus \
       -e PROMETHEUS_SERVER_URL=http://localhost:9090 \
       -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
       -e PROMETHEUS_QUERY_SUPPORT_SPANMETRICS_CONNECTOR=true \
       -e PROMETHEUS_QUERY_NAMESPACE=- \
	   -e PROMETHEUS_QUERY_DURATION_UNIT=- \
       -e PROMETHEUS_QUERY_NORMALIZE_CALLS=true \
       -e PROMETHEUS_QUERY_NORMALIZE_DURATION=true \
       -p 16686:16686 -p 14268:14268 -p 9411:9411 \
       jaegertracing/all-in-one:latest
	@echo "Done."

prometheus-start:
	@echo "Starting Prometheus..."
	@docker run -d --name prometheus \
		-v prometheus-data:/prometheus -p 9090:9090 \
		prom/prometheus
	@echo "Done."

otel-collector:
	@echo "Starting otel_collector..."
	@docker run -d --name otel-collector otel/opentelemetry-collector-contrib:latest