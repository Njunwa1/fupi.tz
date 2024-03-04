protoc -I ./ \
      --go_out ./golang \
      --go_opt paths=source_relative \
      --go-grpc_out ./golang \
      --go-grpc_opt paths=source_relative \
      ./url/url.proto

protoc -I ./ \
      --go_out ./golang \
      --go_opt paths=source_relative \
      --go-grpc_out ./golang \
      --go-grpc_opt paths=source_relative \
      ./keygen/keygen.proto