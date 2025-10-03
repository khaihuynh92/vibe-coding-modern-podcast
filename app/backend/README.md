# Podsite Backend

REST API backend for the Podsite podcast website, built with Go and Gin.

## 🚀 Quick Start

```bash
# Download dependencies
make deps

# Start development server (with auto-reload)
make dev

# Build and run production server
make build && make run

# Run tests
make test
```

## 📁 Project Structure

```
app/backend/
├── cmd/
│   └── server/          # Application entry points
│       └── main.go      # Main server application
├── internal/            # Private application code
│   ├── config/          # Configuration management
│   ├── handlers/        # HTTP request handlers
│   ├── middleware/      # HTTP middleware
│   └── models/          # Data models and business logic
├── pkg/                 # Public library code (if any)
├── tests/               # Test suites
├── bin/                 # Compiled binaries (generated)
├── Dockerfile           # Docker configuration
├── Makefile            # Build automation
└── go.mod              # Go module dependencies
```

## 🛠️ Available Make Commands

### Development
- `make dev` - Start development server with auto-reload
- `make run` - Build and run production server
- `make build` - Build the application binary

### Testing
- `make test` - Run all tests
- `make test-coverage` - Run tests with coverage report
- `make test-race` - Run tests with race detection
- `make benchmark` - Run benchmarks

### Code Quality
- `make fmt` - Format code with gofmt
- `make lint` - Run linter (requires golangci-lint)
- `make vet` - Run go vet

### Dependencies
- `make deps` - Download dependencies
- `make deps-update` - Update dependencies
- `make deps-verify` - Verify dependencies

### Docker
- `make docker-build` - Build Docker image
- `make docker-run` - Run Docker container
- `make docker-dev` - Run Docker container in development mode

### Utilities
- `make clean` - Clean build artifacts
- `make health` - Check if server is running
- `make help` - Show all available commands

## 🌐 API Endpoints

### Health Check
```
GET /health
```
Returns server health status.

### Episodes
```
GET /api/episodes
```
Returns all podcast episodes.

```
GET /api/episodes/featured
```
Returns the featured episode.

```
GET /api/episodes/:id
```
Returns a specific episode by ID.

## 🏗️ Architecture

### RESTful API Design
- Clean, predictable URL structure following REST conventions
- Proper HTTP status codes and methods
- JSON request/response format with structured error handling
- CORS middleware for secure frontend integration
- Gin framework for high-performance HTTP routing

### Go-Specific Features
- Goroutines for concurrent request handling
- Structured logging with configurable levels
- Graceful shutdown with context cancellation
- Built-in profiling endpoints for performance monitoring
- Compile-time type safety and error handling

### Data Models
Episodes are structured with the following Go struct:
```go
type Episode struct {
    ID          string   `json:"id"`
    Number      int      `json:"number"`
    Title       string   `json:"title"`
    Description string   `json:"description"`
    Duration    string   `json:"duration"`
    PublishDate string   `json:"publishDate"`
    ArtworkURL  string   `json:"artworkUrl"`
    ArtworkAlt  string   `json:"artworkAlt,omitempty"`
    AudioURL    string   `json:"audioUrl"`
    Tags        []string `json:"tags"`
}
```

## 🔧 Configuration

### Environment Variables
Configure the application using environment variables:

```env
PORT=3001
GO_ENV=development
CORS_ORIGINS=http://localhost:3000
LOG_LEVEL=info
```

### CORS Configuration
The API is configured to accept requests from the frontend:
- Development: `http://localhost:3000`
- Production: Configure in environment variables

## 🧪 Testing

### Test Structure
- Unit tests for individual functions
- Integration tests for API endpoints
- Health check tests

### Running Tests
```bash
# Run all tests
npm test

# Run tests with coverage
npm run test:coverage

# Run tests in watch mode
npm run test:watch
```

## 🐳 Docker Support

### Building the Image
```bash
npm run docker:build
```

### Running the Container
```bash
# Production mode
npm run docker:run

# Development mode with volume mounting
npm run docker:dev
```

### Docker Compose (Optional)
```yaml
version: '3.8'
services:
  backend:
    build: ./app/backend
    ports:
      - "3001:3001"
    environment:
      - NODE_ENV=production
      - PORT=3001
```

## 🚀 Deployment

### Platform Options
- **Heroku**: Add `Procfile` with `web: npm start`
- **Railway**: Automatic deployment from Git
- **Render**: Connect repository and set build command
- **AWS/Azure/GCP**: Use Docker image

### Health Checks
The `/health` endpoint provides:
- Server status
- Uptime information
- Memory usage
- Response time

### Production Considerations
- Set `NODE_ENV=production`
- Configure proper CORS origins
- Set up monitoring and logging
- Use process managers (PM2, etc.)

## 📊 Performance

### Response Times
- Health check: < 10ms
- Episode list: < 100ms
- Single episode: < 50ms

### Scalability
- Stateless design for horizontal scaling
- In-memory data for fast responses
- Lightweight Node.js runtime

## 🔒 Security

### CORS Policy
Configured to accept requests only from authorized origins.

### Input Validation
- Request parameter validation
- Sanitized error responses
- No sensitive data exposure

### Docker Security
- Non-root user in container
- Minimal Alpine Linux base image
- Health checks for container monitoring

## 🤝 Contributing

1. Make changes in the `src/` directory
2. Add tests for new functionality
3. Run `npm test` to validate changes
4. Run `npm run validate` to check syntax
5. Submit a pull request

## 📄 License

MIT License - see LICENSE file for details.