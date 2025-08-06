.PHONY: build clean test run install

# Build the binary
build:
	go build -o bin/memory main.go

# Clean build artifacts
clean:
	rm -rf bin/

# Run tests
test:
	go test ./...

# Run the program
run:
	go run main.go

# Install the binary to GOPATH/bin
install:
	go install

# Build for multiple platforms
build-all:
	GOOS=linux GOARCH=amd64 go build -o bin/memory-linux-amd64 main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/memory-darwin-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/memory-windows-amd64.exe main.go

# Initialize go modules
init:
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run