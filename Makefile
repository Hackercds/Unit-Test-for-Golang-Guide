.PHONY: test cover htmlcover bench vet lint build test-fast clean all

# Run all tests with verbose output
test:
	go test -v ./...

# Run tests with coverage summary
cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

# Generate HTML coverage report
htmlcover: cover
	go tool cover -html=coverage.out -o coverage.html

# Run all benchmarks
bench:
	go test -bench=. -benchmem ./...

# Static analysis
vet:
	go vet ./...

# Lint check (requires golangci-lint)
lint:
	golangci-lint run ./...

# Verify compilation
build:
	go build ./...

# Run tests without verbose output (faster feedback)
test-fast:
	go test ./...

# Clean generated files
clean:
	rm -f coverage.out coverage.html

# Full pipeline
all: vet test cover bench build
