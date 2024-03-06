
install-protocol:
	@echo "Installing protocol buffer compiler..."
	@brew install protobuf
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@echo "Done."

protoc:
	@echo "Generating protocol buffer files..."
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto
	@echo "Done."