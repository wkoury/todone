# Variables
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)
BINARY := todone

.PHONY: all clean test build lint fmt air

all: build

clean:
	rm -rf ./bin/*

build:
	mkdir -p bin
	go build -o bin/${BINARY} cmd/main.go

build-release:
	mkdir -p bin
	go build -o bin/${BINARY} -ldflags="-s -w" -trimpath cmd/main.go

run: build
	./bin/${BINARY}

test:
	go test $(GO_FILES)

lint:
	@echo "Running golangci-lint..."
	golangci-lint run ./...

fmt:
	@echo "Formatting code with gofmt..."
	go fmt $(GO_FILES)

tidy:
	go mod tidy

install: build-release
	sudo cp ./bin/${BINARY} /usr/local/bin/${BINARY}
