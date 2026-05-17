BINARY := calculator-mcp
IMAGE  := calculator-mcp

.PHONY: build run clean fmt vet tidy test docker-build docker-run help

build:   ## Compile the binary
	go build -o $(BINARY) .

run: build   ## Build and run the server
	./$(BINARY)

clean:   ## Remove the compiled binary
	rm -f $(BINARY)

fmt:   ## Format all Go source files
	gofmt -w .

vet:   ## Run go vet
	go vet ./...

tidy:   ## Tidy go.mod and go.sum
	go mod tidy

test:   ## Run tests
	go test ./...

docker-build:   ## Build the Docker image
	docker build -t $(IMAGE) .

docker-run:   ## Run the server in Docker on port 8080
	docker run --rm -p 8080:8080 $(IMAGE)

help:   ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*##' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*##"}; {printf "  %-8s %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
