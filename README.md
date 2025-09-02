# Social Media Backend API

A REST API for social media functionality built with Go, Fiber, PostgreSQL, and Redis.

## Features

- **User Management**: Create, read, update, delete users
- **Posts**: Create and manage user posts
- **Friendships**: Add and manage friendships between users
- **Real-time Notifications**: WebSocket-based notification system
- **Swagger Documentation**: Interactive API documentation
- **Caching**: Redis-based caching for performance
- **Database**: PostgreSQL with GORM ORM

## Quick Start

### Prerequisites

- Docker and Docker Compose
- Go 1.21+ (for local development)

### Option 1: Full Docker Setup (Recommended)

```bash
# Start all services (PostgreSQL, Redis, Go app)
docker-compose up --build

# Access the application
# API: http://localhost:3015
# Swagger UI: http://localhost:3015/swagger/
```

### Option 2: Local Development

```bash
# Start only database services
docker-compose up postgres redis -d

# Run the Go application locally
go run cmd/main.go

# Access at http://localhost:3015
```

## API Endpoints

### Users
- `POST /socio/users` - Create a new user
- `GET /socio/users` - Get all users
- `GET /socio/users/{id}` - Get user by ID
- `DELETE /socio/users/{id}` - Delete user

### Posts
- `POST /socio/users/{id}/posts` - Create a post for user
- `GET /socio/users/{id}/posts` - Get user's posts
- `DELETE /socio/users/{id}/posts/{post_id}` - Delete a post

### Friendships
- `POST /socio/friends` - Add a friend
- `GET /socio/friends/{id}` - Get user's friends
- `DELETE /socio/friends/{id}?f_id={friend_id}` - Remove friendship (requires friend ID as query parameter)

## Swagger Documentation

Interactive API documentation is available at:
```
http://localhost:3015/swagger/
```

The Swagger UI provides:
- **Complete API endpoint documentation** for all 10 endpoints
- **Request/response schemas** with DTO models
- **Try-it-out functionality** for testing APIs directly
- **Model definitions** for all data structures
- **Parameter documentation** including path, query, and body parameters

## Environment Variables

- `DATABASE_DSN` - PostgreSQL connection string (default: `host=postgres user=manager database=postgres sslmode=disable password=password`)
- `REDIS_URL` - Redis connection URL (default: `redis://redis:6379`)

## Development

### Adding New Endpoints

1. Create controller function with Swagger annotations:
```go
// MyEndpoint godoc
// @Summary Endpoint summary
// @Description Detailed description
// @Tags tag-name
// @Accept json
// @Produce json
// @Param id path string true "ID parameter"
// @Success 200 {object} dto.MyResponse
// @Router /my-endpoint/{id} [get]
func MyEndpoint(c *fiber.Ctx) error {
    // Implementation
}
```

2. Add route in `routes/` directory
3. Regenerate Swagger docs:
```bash
swag init -g cmd/main.go
```

### Project Structure

```
├── cmd/                    # Application entry points
├── controllers/            # HTTP handlers
├── docs/                   # Generated Swagger documentation
├── internals/              # Internal packages
│   ├── cache/             # Redis cache client
│   ├── config/            # Configuration and migrations
│   ├── database/          # Database connection
│   ├── dto/               # Data transfer objects
│   ├── notifications/     # Real-time notifications
│   ├── server/            # HTTP server setup
│   └── validator/         # Request validation
├── models/                 # Database models
├── routes/                 # Route definitions
├── services/               # Business logic
├── docker-compose.yaml     # Docker services
├── Dockerfile             # Go app container
└── go.mod                 # Go dependencies
```

## Technologies Used

- **Go 1.21** - Programming language
- **Fiber v2** - Web framework
- **PostgreSQL** - Primary database
- **Redis** - Caching and sessions
- **GORM** - ORM for database operations
- **Swagger/OpenAPI** - API documentation
- **Docker** - Containerization