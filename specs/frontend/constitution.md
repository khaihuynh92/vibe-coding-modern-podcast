# Frontend Constitution

## Core Principles

### I. Static-First

- All frontend pages are generated as static assets; integrate with backend APIs to get data.
- Frontend deployed on a static host (e.g., GitHub Pages, Netlify, Vercel static export, Cloudflare Pages).
- A build step is allowed but must produce a self-contained `dist` directory for frontend.

### II. Progressive Enhancement

- Base experience must function without JavaScript; JS enhances, never blocks content.
- Navigation, content, and critical interactions remain usable with network constraints and JS disabled.
- Avoid client-side routing unless it demonstrably improves UX without hurting SEO or accessibility.

### III. Performance Budgets

- JavaScript: < 50KB gzipped for initial load
- CSS: < 20KB gzipped
- Images: < 500KB total per page
- Lighthouse scores: Performance ≥ 90, Accessibility ≥ 90, Best Practices ≥ 90, SEO ≥ 90

### IV. Security & Privacy

- Enforce HTTPS; do not embed secrets, API keys, or tokens in the repo or client.
- Content Security Policy: `default-src 'self'`; disallow `unsafe-eval` and `unsafe-inline`; use nonces when required.
- Apply Subresource Integrity on third-party scripts/styles; minimize third-party dependencies.
- No tracking by default; if analytics are required, obtain consent and use privacy-preserving tools.

## Additional Constraints & Standards

- Tech stack: HTML5, CSS (Sass/PostCSS optional), vanilla JS or minimal frameworks that support static output (e.g., Eleventy, Astro, Vite, Next static export).
- Node.js: use active LTS for build tooling; no server runtime required.
- Browser support: last 2 versions of evergreen browsers; graceful degradation for others.
- SEO: valid HTML, unique titles/descriptions, `robots.txt`, `sitemap.xml`, canonical URLs, Open Graph/Twitter cards where appropriate.
- Assets: fingerprinting/immutable caching; minify CSS/JS; preconnect/preload critical assets; use system fonts or max 2 families with `font-display: swap`.
- Images: responsive `srcset`/sizes; lazy-load non-critical images; compress at build time.
- Content: authored in Markdown/MDX or CMS export that compiles to static files.
- Internationalization: optional; if enabled, use separate static routes per locale.
- Legal: include `privacy.html` and `terms.html` if any data collection occurs.

## Development Workflow & Quality Gates

- Git: feature branches with pull requests; require at least one review.
- CI: on every PR and main, run build, HTML validation, link check (internal/external), lint/format check, and Lighthouse CI (mobile).
- Tests: add unit tests when custom JS exists; smoke test rendering of key pages.
- Accessibility checks: automated axe (or equivalent) in CI; manual keyboard pass pre-release.
- Release: merges to main trigger a production build and deployment to the static host.
- Rollback: keep previous build artifacts for instant rollback.

## Governance

- This constitution supersedes other conventions for frontend development.
- All PRs must include a brief note confirming compliance or calling out justified exceptions.
- Amendments require a documented proposal, reviewer approval, and a version bump with migration notes when relevant.

**Version**: 1.0.0 | **Ratified**: 2025-10-01 | **Last Amended**: 2025-10-01
