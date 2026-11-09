# 3V30OStudios Go Server

A Go-based backend server for the 3V30OStudios platform.

## Project Structure

```
.
├── cmd/
│   └── server/          # Main application entry point
│       └── main.go
├── internal/
│   ├── api/            # API handlers and utilities
│   │   └── handler.go
│   └── config/         # Configuration management
│       └── config.go
├── pkg/
│   └── models/         # Data models
│       └── models.go
└── go.mod             # Go module definition
```

## Getting Started

### Prerequisites

- Go 1.24.9 or higher

### Installation

```bash
# Install dependencies
go mod tidy

# Build the server
go build -o bin/server ./cmd/server
```

### Running the Server

```bash
# Run directly
go run cmd/server/main.go

# Or use the built binary
./bin/server
```

The server will start on port 8080 by default.

### Environment Variables

- `PORT` - Server port (default: 8080)
- `HOST` - Server host (default: 0.0.0.0)
- `LOG_LEVEL` - Logging level (default: info)

## API Endpoints

- `GET /` - Root endpoint, returns welcome message
- `GET /health` - Health check endpoint

## Development

```bash
# Format code
go fmt ./...

# Run tests
go test ./...

# Build for production
go build -ldflags="-s -w" -o bin/server ./cmd/server
```

## License

Part of the 3V30OStudios ecosystem.
