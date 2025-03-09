# Output binary name
BINARY_NAME = obd2-tui

# Build the application
build:
	go build -o $(BINARY_NAME) ./cmd

# Run the application after building
run: build
	./$(BINARY_NAME)

# Clean up the binary
clean:
	rm -f $(BINARY_NAME)

# Rebuild from scratch
rebuild: clean build
