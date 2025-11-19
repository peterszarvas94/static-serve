# Serve

A simple Go-based static file server for local development and testing.

## Features

- Serve any directory as a static website
- Configurable port and directory
- Simple command-line interface

## Installation

```bash
go install github.com/peterszarvas94/serve@latest
```

## Usage

```bash
serve [options]
```

### Options

- `-d, -dir, --dir <directory>` - Directory to serve (default: current directory)
- `-p, -port, --port <port>` - Port to listen on (default: 8080)

### Examples

```bash
# Serve current directory on port 8080
serve

# Serve a specific directory (all equivalent)
serve -d /path/to/website
serve -dir /path/to/website
serve --dir /path/to/website

# Use a different port (all equivalent)
serve -p 3000
serve -port 3000
serve --port 3000

# Combine options
serve -d ./public -p 9000
serve --dir ./public --port 9000
```

## Building

```bash
go build -o serve main.go
./serve -d ./public -p 8080
```

## Requirements

- Go 1.25.4 or later