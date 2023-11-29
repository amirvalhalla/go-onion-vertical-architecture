unit_test:
	go test ./... -v

fmt:
	gofumpt -l -w .

devtools:
	@echo "Installing devtools"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.1
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.12
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.12
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install github.com/bufbuild/buf/cmd/buf@v1.25.0
	go install mvdan.cc/gofumpt@latest
	go install github.com/rakyll/statik@v0.1


check:
	golangci-lint run --timeout=20m0s
