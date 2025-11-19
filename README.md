# Server

A simple Go-based static file server for local development and testing.

## Features

- Serve any directory as a static website
- Configurable port and directory
- Simple command-line interface

## Installation

```bash
go install github.com/peterszarvas94/server@latest
```

## Usage

```bash
server [options]
```

### Options

- `-d, -dir, --dir <directory>` - Directory to serve (default: current directory)
- `-p, -port, --port <port>` - Port to listen on (default: 8080)

### Examples

```bash
# Serve current directory on port 8080
server

# Serve a specific directory (all equivalent)
server -d /path/to/website
server -dir /path/to/website
server --dir /path/to/website

# Use a different port (all equivalent)
server -p 3000
server -port 3000
server --port 3000

# Combine options
server -d ./public -p 9000
server --dir ./public --port 9000
```

## Building

```bash
go build -o serve main.go
./serve -d ./public -p 8080
```

## Requirements

- Go 1.25.4 or later