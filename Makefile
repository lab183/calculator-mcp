BINARY := calculator-mcp

.PHONY: build run clean fmt vet tidy

build:
	go build -o $(BINARY) .

run: build
	./$(BINARY)

clean:
	rm -f $(BINARY)

fmt:
	gofmt -w .

vet:
	go vet ./...

tidy:
	go mod tidy
