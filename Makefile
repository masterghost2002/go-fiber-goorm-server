.PHONY: all build run clean build-prod run-prod

# Default target
all: build

# Build the application for development
build:
	go build -o ./bin/my-go-server ./cmd/server/server.go

# Run the application in development mode with watch using Air
run:
	go run ./cmd/server/server.go

# Clean build artifacts
clean:
	rm -rf ./bin

# Build the application for production
build-prod:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/my-go-server-prod ./cmd/server/main.go

# Run the application in production
run-prod:
	./bin/my-go-server-prod
