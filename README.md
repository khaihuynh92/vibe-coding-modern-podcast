# Podcast - Modern Audio Storytelling

![CI](https://github.com/yourusername/podsite/workflows/CI/badge.svg)

A sleek, modern podcast website built as a full-stack application with static frontend and RESTful API backend. Features 20 episodes, responsive design, and full accessibility.

## Features

### Frontend
- 🎯 **Static-first** - Deployable to any static host
- ♿ **Accessible** - WCAG 2.1 AA compliant
- 📱 **Responsive** - Mobile-first design
- ⚡ **Fast** - Lighthouse 90+ scores
- 🔒 **Secure** - Strict CSP, no third-party tracking
- 🎨 **Modern UI** - Clean, distinctive design

### Backend
- 🚀 **RESTful API** - Go with Gin framework for high performance
- 🔄 **Stateless** - Horizontally scalable with goroutines
- 🛡️ **Secure** - CORS middleware, security headers, type safety
- 📊 **Health Checks** - Kubernetes-ready with /health and /ready endpoints
- ⚡ **Fast** - Compiled binary with minimal resource usage

## Architecture

This project contains two **completely independent** applications:

- **Frontend** (`app/frontend/`) - Static site with API integration (Node.js/npm)
- **Backend** (`app/backend/`) - RESTful API server (Go/make)

Each application has its own:
- Build system and dependencies
- Testing framework and CI/CD
- Deployment process and Docker configuration
- Documentation and development workflow

## Quick Start

### Prerequisites

- Node.js 18+ (Active LTS) for frontend
- Go 1.21+ for backend

### Build Both Projects

Each project is completely independent and must be built separately:

```bash
# Build frontend
cd app/frontend
npm install
npm run build

# Build backend (in a new terminal)
cd app/backend
make deps
make build
```

### Start Both Services

```bash
# Start frontend (terminal 1)
cd app/frontend
npm run dev

# Start backend (terminal 2)
cd app/backend
make dev
```

This will start:
- Frontend: `http://localhost:3000`
- Backend API: `http://localhost:3001`

### Frontend Only

```bash
cd app/frontend
npm install
npm run build
npm run preview
```

Visit `http://localhost:3000` to view the site.

### Backend Only

```bash
cd app/backend
make deps
make dev
```

API available at `http://localhost:3001`

### Validate

```bash
# Run all validators
npm run validate

# Individual checks
npm run validate:html        # HTML validation
npm run validate:links       # Link checking
npm run validate:a11y        # Accessibility
npm run validate:lighthouse  # Performance & SEO
```

## Project Structure

```
podsite/                      # Project root
├── app/
│   ├── frontend/            # Frontend application (standalone)
│   │   ├── site/
│   │   │   ├── content/     # Episode data and page content
│   │   │   │   ├── episodes.json
│   │   │   │   ├── about.md
│   │   │   │   └── faq.json
│   │   │   ├── public/      # Static templates and assets
│   │   │   │   ├── index.html
│   │   │   │   ├── episodes.html
│   │   │   │   ├── about.html
│   │   │   │   ├── faq.html
│   │   │   │   └── assets/  # CSS, JS, images, audio
│   │   │   └── src/
│   │   │       └── scripts/ # API integration modules
│   │   ├── scripts/         # Build and content generation
│   │   │   └── build.mjs
│   │   ├── tests/           # Validation tests
│   │   ├── dist/            # Build output (generated)
│   │   ├── package.json     # Frontend dependencies
│   │   └── README.md
│   └── backend/             # Backend API application (standalone)
│       ├── cmd/
│       │   └── server/      # Application entry points
│       ├── internal/        # Private application code
│       │   ├── config/      # Configuration management
│       │   ├── handlers/    # HTTP request handlers
│       │   ├── middleware/  # HTTP middleware
│       │   └── models/      # Data models and business logic
│       ├── tests/           # API tests
│       ├── bin/             # Compiled binaries (generated)
│       ├── Dockerfile       # Docker configuration
│       ├── Makefile        # Build automation
│       ├── go.mod          # Go module dependencies
│       └── README.md
├── specs/                   # Feature specifications
│   ├── frontend/            # Frontend specifications
│   │   ├── constitution.md
│   │   └── 001-i-am-building/
│   └── backend/             # Backend specifications
│       └── constitution.md
├── .github/                 # CI/CD workflows
│   └── workflows/
│       ├── frontend-ci.yml  # Frontend CI/CD
│       ├── backend-ci.yml   # Backend CI/CD
│       └── ci.yml           # Orchestrated CI
└── README.md                # This file
```

## Pages

- **Home** (`/`) - Hero and featured episode
- **Episodes** (`/episodes.html`) - All 20 episodes with inline details
- **About** (`/about.html`) - Podcast mission and team
- **FAQ** (`/faq.html`) - Frequently asked questions

## Development

### Independent Project Commands

Each project has its own build system and must be run from its directory:

**Frontend Commands** (from `app/frontend/`):
```bash
npm install          # Install dependencies
npm run build        # Build static site
npm run dev          # Development server
npm run test         # Run all tests
npm run clean        # Clean build artifacts
npm run deploy       # Deploy to hosting platform
```

**Backend Commands** (from `app/backend/`):
```bash
make deps            # Download Go dependencies
make build           # Build Go binary
make dev             # Development server with auto-reload
make test            # Run Go tests
make clean           # Clean build artifacts
make docker-build    # Build Docker image
make docker-run      # Run Docker container
```

### Frontend Development

Content is managed in JSON/Markdown under `app/frontend/site/content/`. The build script validates against JSON schemas and generates static HTML.

**Adding Episodes**: Edit `app/frontend/site/content/episodes.json` following the schema in `specs/frontend/001-i-am-building/contracts/episode.schema.json`.

**Styling**: CSS is in `app/frontend/site/public/assets/css/main.css`. Uses CSS custom properties for theming.

**API Integration**: JavaScript modules in `app/frontend/site/src/scripts/` handle API calls with fallback to static content.

### Backend Development

The backend is a Go API server using the Gin web framework for high performance.

**Adding Endpoints**: 
1. Create handler in `app/backend/internal/handlers/`
2. Add route in `app/backend/cmd/server/main.go`
3. Update `app/backend/README.md`

**Data Source**: Currently reads from `app/frontend/site/content/episodes.json`. Replace with a database for production.

**API Documentation**: See `app/backend/README.md` for available endpoints and Go-specific development practices.

## Constitution Compliance

This project follows the [Podsite Constitution](/.specify/memory/constitution.md) with separate constitutions for:

- **Frontend**: [Frontend Constitution](/specs/frontend/constitution.md) - Static-first, progressive enhancement, performance budgets
- **Backend**: [Backend Constitution](/specs/backend/constitution.md) - API-first design, security, scalability

Current compliance:
- ✅ Static-first architecture
- ✅ Progressive enhancement
- ✅ Performance budgets met
- ✅ WCAG 2.1 AA accessibility
- ✅ Security & privacy (CSP, no tracking)

## CI/CD

GitHub Actions workflow runs on every PR:
- Build validation
- HTML validation
- Link checking
- Accessibility tests (axe)
- Lighthouse CI

## License

MIT

---

**Version**: 1.0.0  
**Built**: 2025-10-01

