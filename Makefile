
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