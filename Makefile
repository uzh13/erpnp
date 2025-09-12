# Simple Makefile for development
BINARY=bin/erpnp
PKG=./...
WASM_OUT=web/wasm/erpnp.wasm
GO=go


.PHONY: all build build-wasm test lint fmt tidy clean run docker-build


all: build


build:
	$(GO) build -o $(BINARY) ./cmd/processor


build-wasm:
	GOOS=js GOARCH=wasm $(GO) build -o $(WASM_OUT) ./web/wasm


test:
	$(GO) test ./...


lint:
	golangci-lint run || true


fmt:
	gofmt -s -w .


tidy:
	$(GO) mod tidy


run:
	$(GO) run ./cmd/processor --help


clean:
	rm -rf $(BINARY) $(WASM_OUT)


docker-build:
	docker build -t processor:latest .