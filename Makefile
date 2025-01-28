# Makefile for Kisumu Lang

.PHONY: all build run-repl test clean docs format lint coverage pre-commit help

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

# Format the code
format:
	gofmt -w .

# Run linters
lint:
	golangci-lint run

# Generate and display test coverage
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

# Run all checks before committing
pre-commit: format lint test

# Display help message
help:
	@echo "Available targets:"
	@echo "  build       - Build the project"
	@echo "  run-repl    - Run the REPL"
	@echo "  test        - Run tests"
	@echo "  coverage    - Generate and display test coverage"
	@echo "  format      - Format the code"
	@echo "  lint        - Run linters"
	@echo "  clean       - Clean up build artifacts"
	@echo "  docs        - Open documentation"
	@echo "  pre-commit  - Run all checks before committing"