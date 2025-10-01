# Podcast - Modern Audio Storytelling

![CI](https://github.com/yourusername/podsite/workflows/CI/badge.svg)

A sleek, modern podcast website built as a static site with progressive enhancement. Features 20 episodes, responsive design, and full accessibility.

## Features

- 🎯 **Static-first** - No runtime server required
- ♿ **Accessible** - WCAG 2.1 AA compliant
- 📱 **Responsive** - Mobile-first design
- ⚡ **Fast** - Lighthouse 90+ scores
- 🔒 **Secure** - Strict CSP, no third-party tracking
- 🎨 **Modern UI** - Clean, distinctive design

## Quick Start

### Prerequisites

- Node.js 18+ (Active LTS)

### Build

```bash
# Install dependencies
npm install

# Build the site (from root)
npm run build

# Or build from frontend directory
cd app/frontend
npm install
npm run build

# Preview locally
npm run preview
```

Visit `http://localhost:3000` to view the site.

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
podsite/
├── app/
│   └── frontend/          # Frontend application (standalone)
│       ├── site/          # Static site content and templates
│       │   ├── content/   # Episode data and page content
│       │   │   ├── episodes.json
│       │   │   ├── about.md
│       │   │   └── faq.json
│       │   ├── public/    # Static templates and assets
│       │   │   ├── index.html
│       │   │   ├── episodes.html
│       │   │   ├── about.html
│       │   │   ├── faq.html
│       │   │   └── assets/
│       │   └── src/
│       │       ├── styles/
│       │       └── scripts/
│       ├── scripts/       # Build and content generation scripts
│       │   └── build.mjs  # Build script
│       ├── tests/         # Validation tests
│       ├── dist/          # Build output (generated)
│       ├── package.json   # Frontend dependencies & scripts
│       ├── package-lock.json
│       ├── .htmlvalidate.json
│       ├── lighthouserc.json
│       └── linkinator.config.json
├── specs/                 # Feature specifications
│   ├── frontend/          # Frontend specifications
│   │   └── 001-i-am-building/ # Modern Podcast Website
│   └── backend/           # Backend specifications
├── .github/               # CI/CD workflows
└── package.json           # Workspace configuration
```

## Pages

- **Home** (`/`) - Hero and featured episode
- **Episodes** (`/episodes.html`) - All 20 episodes with inline details
- **About** (`/about.html`) - Podcast mission and team
- **FAQ** (`/faq.html`) - Frequently asked questions

## Development

Content is managed in JSON/Markdown under `app/frontend/site/content/`. The build script validates against JSON schemas and generates static HTML.

### Adding Episodes

Edit `app/frontend/site/content/episodes.json` following the schema in `specs/frontend/001-i-am-building/contracts/episode.schema.json`.

### Styling

CSS is in `app/frontend/site/public/assets/css/main.css`. Uses CSS custom properties for theming.

### Progressive Enhancement

JavaScript in `app/frontend/site/public/assets/js/app.js` enhances with inline episode details and audio playback. Site remains functional without JS.

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

