# Backend Constitution

## Core Principles

### I. API-First Design

- All backend functionality exposed through well-defined REST APIs or GraphQL endpoints.
- APIs designed for frontend consumption with clear contracts and documentation.
- Backend services are stateless and horizontally scalable.

### II. Security & Privacy

- Enforce HTTPS for all API endpoints; no HTTP in production.
- Implement proper authentication and authorization for protected endpoints.
- No secrets, API keys, or tokens in code; use environment variables or secure secret management.
- Apply rate limiting and input validation on all endpoints.
- Log security events and monitor for suspicious activity.

### III. Performance & Scalability

- API response times: < 200ms for 95th percentile
- Database queries: < 100ms for 95th percentile
- Support horizontal scaling through load balancing
- Implement caching strategies (Redis, CDN) where appropriate
- Monitor and alert on performance degradation

### IV. Data Integrity & Reliability

- All data operations are ACID compliant where applicable
- Implement proper error handling and logging
- Use database transactions for multi-step operations
- Regular automated backups with tested restore procedures
- Implement circuit breakers for external service dependencies

## Additional Constraints & Standards

- **Tech Stack**: Go 1.21+ with Gin web framework for HTTP routing and middleware
- **Project Structure**: Follow Go standard project layout (`cmd/`, `internal/`, `pkg/`, `api/`)
- **Dependency Management**: Go modules (`go.mod`) with semantic versioning
- **Database**: In-memory structures for development; PostgreSQL for production
- **HTTP Framework**: Gin for routing, middleware, JSON handling, and request validation
- **Testing**: Go's built-in testing framework with table-driven tests, benchmarks, and race detection
- **Logging**: Structured logging with `logrus` or `zap`, JSON format for production
- **Error Handling**: Proper error wrapping with `fmt.Errorf` and custom error types
- **Configuration**: Environment-based config with `viper` or built-in `os.Getenv`
- **API Documentation**: Swagger/OpenAPI with `swaggo/gin-swagger` annotations
- **Code Quality**: `gofmt`, `golangci-lint`, `go vet`, and `staticcheck` in CI pipeline
- **Performance**: Built-in profiling with `net/http/pprof` for performance analysis
- **Containerization**: Multi-stage Docker builds with Alpine Linux base images
- **Health Checks**: `/health` and `/ready` endpoints for Kubernetes liveness/readiness probes

## Development Workflow & Quality Gates

- Git: feature branches with pull requests; require at least one review.
- CI: on every PR and main, run tests, linting, security scans, and API contract validation.
- Testing: minimum 80% code coverage for business logic
- Security: automated dependency scanning, SAST tools, and regular security audits
- API versioning: maintain backward compatibility for at least 2 major versions
- Documentation: keep API documentation up-to-date with code changes

## Data Management

- Personal data: comply with GDPR, CCPA, and other applicable privacy regulations
- Data retention: implement policies for data lifecycle management
- Encryption: encrypt sensitive data at rest and in transit
- Access control: implement principle of least privilege for data access
- Audit trails: log all data access and modifications

## Integration Standards

- Frontend contracts: maintain clear API contracts in `/specs/` directory
- Error responses: consistent error format across all endpoints
- CORS: properly configured for frontend domain(s)
- Webhooks: implement for real-time updates where needed
- Third-party integrations: use official SDKs and follow best practices

## Governance

- This constitution supersedes other conventions for backend development.
- All PRs must include a brief note confirming compliance or calling out justified exceptions.
- Amendments require a documented proposal, reviewer approval, and a version bump with migration notes when relevant.

**Version**: 2.0.0 | **Ratified**: 2025-10-01 | **Last Amended**: 2025-10-02 | **Major Change**: Migration from Node.js to Go
