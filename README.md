# Go Snippet

A modern, secure web application for sharing code snippets, built with Go. Features secure storage, dynamic templates, and robust middleware with TLS support.

## Features

- SQLite database with GORM for snippet storage
- Dynamic template rendering with caching
- Structured logging using slog
- Middleware stack using Alice for logging and recovery
- RESTful API endpoints for snippet management
- Secure session management with SCS
- TLS/HTTPS support with modern cipher suites
- Configurable server timeouts

## Tech Stack

- Go 1.21+
- SQLite with GORM
- HTML Templates
- Alice middleware
- SCS session management
- TLS for secure communication

## Project Structure

```
.
├── cmd/
│   └── web/                # Application entrypoint and server configuration
├── internal/
│   └── models/            # Database models and business logic
├── ui/
│   └── templates/         # HTML templates
└── tls/                   # TLS certificates
    ├── cert.pem
    └── key.pem
```

## Getting Started

1. Clone the repository
2. Run `go mod download`
3. Generate TLS certificates (for development):
   ```bash
   mkdir tls
   cd tls
   go run /usr/local/go/src/crypto/tls/generate_cert.go --host=localhost
   ```
4. Execute `go run ./cmd/web`
5. Visit `https://localhost:4000`

## Security Features

- HTTPS/TLS encryption
- Modern cipher suite preferences (X25519, P256)
- Secure session configuration
- Server timeouts:
  - Idle timeout: 1 minute
  - Read timeout: 5 seconds
  - Write timeout: 10 seconds

## API Endpoints

- `GET /`: Home page with snippets list
- `GET /snippet/view/{id}`: View specific snippet
- `GET /snippet/create`: Snippet creation form
- `POST /snippet/create`: Create new snippet

## Development

- GORM for database operations
- Template caching for performance
- Structured logging with slog
- Panic recovery middleware
- Request logging middleware
- Connection pool configuration
- Secure session management

## Contributing

1. Fork repository
2. Create feature branch
3. Commit changes
4. Push to branch
5. Open Pull Request

## License

This project is licensed under the [MIT License](LICENSE).
