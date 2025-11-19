# Static Serve

A simple Go-based static file server for local development and testing.

## Features

- Serve any directory as a static website
- Configurable port and directory
- Simple command-line interface

## Usage

```bash
go run main.go [options]
```

### Options

- `-d <directory>` - Directory to serve (default: current directory)
- `-p <port>` - Port to listen on (default: 8080)

### Examples

```bash
# Serve current directory on port 8080
go run main.go

# Serve a specific directory
go run main.go -d /path/to/website

# Use a different port
go run main.go -p 3000

# Combine options
go run main.go -d ./public -p 9000
```

## Building

```bash
go build -o serve main.go
./serve -d ./public -p 8080
```

## Requirements

- Go 1.25.4 or later