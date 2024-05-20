.PHONY: build test clean run

BINARY_NAME=boilerplate

SOURCES=$(shell find . -name '*.go' -not -path "./vendor/*")

GO_CMD=go
AIR_CMD=air

build: $(SOURCES)
	@echo "Building..."
	$(GO_CMD) build -o $(BINARY_NAME) ./cmd

test: $(SOURCES)
	@echo "Testing..."
	$(GO_CMD) test -v ./...

clean:
	@echo "Cleaning..."
	$(GO_CMD) clean -cache
	rm -f $(BINARY_NAME)

run: build
	@echo "Running..."
	./$(BINARY_NAME)

dev: $(SOURCES)
	@echo "Running in Dev..."
	$(AIR_CMD) -c .air.toml
