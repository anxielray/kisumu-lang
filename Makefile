# Makefile for Kisumu Lang

.PHONY: all build run-repl test clean docs

# Default task
all: build

# Build the project
build:
	@echo "Building Kisumu Lang CLI..."
	go build -o bin/kisumu ./cmd/kisumu || (echo "Build failed. Please check errors above." && exit 1)
	go build -o bin/repl ./cmd/repl || (echo "Build failed for REPL. Please check errors above." && exit 1)

# Run the REPL
run-repl:
	./bin/repl

# Run tests
test:
	go test ./...

# Clean up build artifacts
clean:
	rm -rf bin
	rm -rf coverage.out

# Generate documentation
docs:
	@echo "Opening documentation..."
	@open ./docs/README.md || xdg-open ./docs/README.md || echo "Please manually open the docs in your browser."
