
# Phase 1: Quickstart — Modern Podcast Website

## Prerequisites
- Node.js (Active LTS)

## Build
1. Place content under `site/content/`:
   - `episodes.json` (20 items matching contracts/episode.schema.json)
   - `about.md`, `faq.json`
2. Place templates/static under `site/public/` and assets in `site/public/assets/`.
3. Run the build script to produce `dist/`:
   - Example: `node scripts/build.mjs`

## Preview
- Use any static file server to preview `dist/` locally, e.g.:
  - `npx serve dist` or `python3 -m http.server` from `dist/`.

## Validate
- HTML validation and link check via CI.
- Accessibility check (axe) and Lighthouse CI (mobile) run in CI.


