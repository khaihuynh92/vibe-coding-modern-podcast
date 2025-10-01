
# Phase 0: Outline & Research — Modern Podcast Website

## Decisions
- Static-only deployment with a Node build script that produces `dist/`.
- No external feeds/APIs; 20 episodes embedded as static JSON content.
- On-site playback via a single bundled mock audio file reused across episodes.
- Neutral, modern branding with strong accent color; accessible typography and spacing.

## Rationale
- Minimizes complexity and dependencies; aligns with Static-First principle.
- Ensures offline-friendly behavior and predictable deployments.
- Simplifies privacy/security posture (no third-party trackers, no secrets).

## Alternatives Considered
- Eleventy/Astro static site generators: powerful, but extra tooling overhead for current scope.
- Client-side routing SPA: conflicts with progressive enhancement and SEO needs.

## References
- SEO basics for static sites (titles, descriptions, sitemap, robots).
- Accessibility for audio controls and transcripts (WCAG 2.1 AA guidance).


