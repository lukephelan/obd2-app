# Output binary name
BINARY_NAME = obd2-tui

# Build the application
.PHONY: build
build:
	go build -o $(BINARY_NAME) ./cmd

# Run the application after building
.PHONY: run
run: build
	./$(BINARY_NAME)

# Clean up the binary
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)

# Rebuild from scratch
.PHONY: rebuild
rebuild: clean build

# Run unit tests
.PHONY: test coverage
test:
	go test ./...

coverage:
	go test -coverprofile=coverage.out ./... && go tool cover -func=coverage.out
