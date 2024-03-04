
install-protocol:
	@echo "Installing protocol buffer compiler..."
	@brew install protobuf
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@echo "Done."