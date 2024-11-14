# Go Snippet

A modern web application for sharing code snippets, built with Go. Features secure storage, dynamic templates, and robust middleware.

## Features

- SQLite database with GORM for snippet storage
- Dynamic template rendering with caching
- Structured logging
- Middleware stack using Alice for logging and recovery
- RESTful API endpoints for snippet management

## Tech Stack

- Go 1.21+
- SQLite with GORM
- HTML Templates
- Alice middleware

## Project Structure
```
.
├── cmd/
│   └── web/                
├── internal/
│   └── models/             
└── ui/
    └── templates/          

```

## Getting Started

1. Clone the repository
2. Run `go mod download`
3. Execute `go run ./cmd/web`
4. Visit `http://localhost:8080`

## API Endpoints

- `GET /`: Home page with snippets list
- `GET /snippet/view/{id}`: View specific snippet
- `GET /snippet/create`: Snippet creation form
- `POST /snippet/create`: Create new snippet

## Development

- GORM for database operations
- Template caching for performance
- Structured logging
- Panic recovery middleware
- Request logging middleware

## Contributing

1. Fork repository
2. Create feature branch
3. Commit changes
4. Push to branch
5. Open Pull Request

## License
This project is licensed under the [MIT License](LICENSE).
