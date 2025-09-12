# ERPNP - Evolving Resolutive Process Notation Processor

A Go-based CLI tool and WebAssembly library for processing ERPN files.

## Quick Start

```bash
# Build CLI
make build

# Run CLI
./bin/erpnp --help

# Build WebAssembly version
make build-wasm
```

## Commands

- `validate` - Validate ERPNP files
- `transform` - Transform configurations
- `add` - Add nodes (stub)
- `remove` - Remove nodes (stub)
- `sum` - Sum fields (stub)
- `wasm` - WebAssembly build info

## Development

```bash
make test    # Run tests
make lint    # Run linting
make fmt     # Format code
```

See `CLAUDE.md` for detailed development guidance.