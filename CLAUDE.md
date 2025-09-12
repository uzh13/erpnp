# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

The project uses a Makefile for all development tasks:

- `make build` - Build the main CLI binary to `bin/erpnp`
- `make build-wasm` - Build WebAssembly version to `web/wasm/erpnp.wasm`
- `make test` - Run all Go tests
- `make lint` - Run golangci-lint (configured in `.golangci.yml`)
- `make fmt` - Format all Go code with gofmt
- `make tidy` - Tidy Go module dependencies
- `make run` - Run the processor with --help flag
- `make clean` - Remove built binaries
- `make docker-build` - Build Docker image

The CI pipeline runs: `make tidy`, `make lint`, `make test`, `make build`

## Architecture

This is a Go project for Evolving Resolutive Process Notation (ERPN).

ERPN schema https://github.com/SchemaStore/schemastore/blob/master/src/schemas/json/evolving-resolutive-process-notation-1.0.json

Project contains two main parts:

### CLI Tool (`cmd/processor/`)
- Main entry point provides subcommands: `validate`, `transform`, `add`, `remove`, `sum`, `wasm`
- Currently contains stub implementations for configuration processing
- Built as `bin/erpnp` binary

### WebAssembly Interface (`web/wasm/`)
- Simple WASM wrapper that exposes `helloFromWasm()` function to JavaScript
- HTML test page at `web/wasm/index.html` for testing WASM functionality
- Built with `GOOS=js GOARCH=wasm` to `web/wasm/erpnp.wasm`

### Project Structure
- `cmd/processor/` - CLI application main package
- `web/wasm/` - WebAssembly build target and HTML interface
- Go module: `github.com/uzh13/erpnp`
- Uses Go 1.22.4+

Note: 