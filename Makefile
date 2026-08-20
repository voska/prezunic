VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -s -w -X main.version=$(VERSION)
BIN     := bin/prezunic

.PHONY: build test lint clean fmt vet install ci

build:
	go build -ldflags "$(LDFLAGS)" -o $(BIN) ./cmd/prezunic

test:
	go test -race ./...

lint:
	golangci-lint run ./...

fmt:
	gofmt -w .

vet:
	go vet ./...

clean:
	rm -rf bin/

ci: fmt lint vet test build

install: build
	@mkdir -p ~/.local/bin
	cp $(BIN) ~/.local/bin/prezunic
	@echo "installed -> ~/.local/bin/prezunic"

