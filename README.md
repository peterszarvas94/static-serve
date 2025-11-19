# Static Serve

A simple Go-based static file server for local development and testing.

## Features

- Serve any directory as a static website
- Configurable port and directory
- Simple command-line interface

## Installation

```bash
go install github.com/peterszarvas94/static-serve@latest
```

## Usage

```bash
static-serve [options]
```

### Options

- `-d, --dir <directory>` - Directory to serve (default: current directory)
- `-p, --port <port>` - Port to listen on (default: 8080)

### Examples

```bash
# Serve current directory on port 8080
static-serve

# Serve a specific directory
static-serve -d /path/to/website
static-serve --dir /path/to/website

# Use a different port
static-serve -p 3000
static-serve --port 3000

# Combine options
static-serve -d ./public -p 9000
static-serve --dir ./public --port 9000
```

## Building

```bash
go build -o serve main.go
./serve -d ./public -p 8080
```

## Requirements

- Go 1.25.4 or later