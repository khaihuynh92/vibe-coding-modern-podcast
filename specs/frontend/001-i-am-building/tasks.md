
# Tasks: Modern Podcast Website

**Input**: Design documents from `/Users/minhkhaihuynh/Documents/Git/podsite/specs/001-i-am-building/`
**Prerequisites**: research.md (available), data-model.md (available), quickstart.md (available)

## Execution Flow (main)
```
1. Load plan.md from feature directory
   → If not found: continue using research.md, data-model.md, quickstart.md
2. Load optional design documents:
   → data-model.md: Extract entities → model/content tasks
   → research.md: Extract decisions → setup and constraints tasks
   → quickstart.md: Extract build/validation scenarios → test tasks
3. Generate tasks by category with TDD ordering
4. Number tasks sequentially (T001, T002...)
5. Mark [P] for tasks on different files with no dependencies
6. Provide dependency notes and a parallel execution example
```

## Format: `[ID] [P?] Description`
- **[P]**: Can run in parallel (different files, no dependencies)
- Include exact file paths in descriptions

## Phase 3.1: Setup
- [x] T001 Create project structure per plan under `app/frontend/` root
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/{content,public,src/{styles,scripts,templates},public/assets/{audio,images,css,js},scripts}`
- [x] T002 Initialize Node workspace with `package.json` and scripts
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/package.json` (scripts: build, preview, validate)
- [x] T003 [P] Add `.gitignore` entries for `dist/` and tooling caches
      → `/Users/minhkhaihuynh/Documents/Git/podsite/.gitignore`

## Phase 3.2: Tests First (TDD) — CI & Site Validation
**CRITICAL: These checks/configs MUST exist and initially fail before implementation**
- [x] T004 [P] Create HTML validation config
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/.htmlvalidate.json`
- [x] T005 [P] Create link checking config (internal/external)
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/linkinator.config.json`
- [x] T006 [P] Create Lighthouse CI config (mobile)
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/lighthouserc.json`
- [x] T007 [P] Create axe accessibility test runner for built pages
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/tests/accessibility/axe.mjs`
- [x] T008 Configure CI workflow to run build + validators on PR/main
      → `/Users/minhkhaihuynh/Documents/Git/podsite/.github/workflows/ci.yml`

## Phase 3.3: Core Implementation (ONLY after tests are failing)
- [x] T009 Create JSON Schemas for content contracts
      → `/Users/minhkhaihuynh/Documents/Git/podsite/specs/frontend/001-i-am-building/contracts/episode.schema.json`
      → `/Users/minhkhaihuynh/Documents/Git/podsite/specs/frontend/001-i-am-building/contracts/page-content.schema.json`
- [x] T010 [P] Add 20 episodes matching Episode schema (used as fallback content only)
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/content/episodes.json`
- [x] T011 [P] Add About page content
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/content/about.md`
- [x] T012 [P] Add FAQ content (array of Q/A objects)
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/content/faq.json`
- [x] T013 Create Node build script to output static `dist/`
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/scripts/build.mjs` (copy/minify, inject content into templates, emit sitemap/robots)
- [x] T014 Create landing page template with featured episode section
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/public/index.html`
- [x] T015 Create episodes page template with list + inline details pattern
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/public/episodes.html`
- [x] T016 Create About page template
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/public/about.html`
- [x] T017 Create FAQ page template
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/public/faq.html`
- [x] T018 [P] Add assets: bundled mock audio, base images, CSS entry
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/public/assets/audio/mock.mp3`
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/public/assets/images/`
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/public/assets/css/main.css`
- [x] T019 [P] Add minimal progressive enhancement script (inline expansion + audio)
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/src/scripts/app.mjs`
- [x] T020 Implement sitemap.xml and robots.txt generation in build
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/public/robots.txt`
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/public/sitemap.xml` (generated to `dist/`)
- [x] T021 Add SEO/meta and CSP in templates (no third-party scripts)
      → update `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/public/*.html`

## Phase 3.3.1: API Integration (NEW)
- [x] T027 [P] Create API service module for episode endpoints
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/src/scripts/api.mjs`
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/src/scripts/config.mjs`
- [x] T028 [P] Implement featured episode API integration
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/src/scripts/featured-episode.mjs`
- [x] T029 [P] Implement episodes list API integration
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/src/scripts/episodes-list.mjs`
- [x] T030 [P] Add error handling and fallback content for API failures
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/src/scripts/error-handler.mjs`
- [x] T031 [P] Create API integration tests
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/tests/api/integration.mjs`

## Phase 3.4: Integration
- [x] T022 Wire build to validate content against JSON Schemas
      → validate `episodes.json`, `faq.json` prior to HTML generation
- [x] T032 Integrate API calls into landing page for featured episode
      → update `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/public/index.html`
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/public/assets/js/api-integration.js`
- [x] T033 Integrate API calls into episodes page for episodes list
      → update `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/public/episodes.html`
- [x] T034 Update build script to include API integration scripts
      → update `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/scripts/build.mjs`
- [x] T023 Run build and ensure CI validators now pass locally
      → `node app/frontend/scripts/build.mjs` then run validators

## Phase 3.5: Polish
- [x] T024 [P] Optimize images and add responsive `srcset` + lazy loading
      → update images under `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/public/assets/images/`
- [x] T025 [P] Open Graph/Twitter meta on all pages
      → update `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/public/*.html`
- [x] T026 [P] Update README with quickstart and CI badges
      → `/Users/minhkhaihuynh/Documents/Git/podsite/README.md`
- [x] T035 [P] Add API endpoint configuration and environment setup
      → `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/src/scripts/config.mjs`
- [x] T036 [P] Add loading states and error UI for API calls
      → update `/Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/public/assets/css/main.css`

## Dependencies
- Setup (T001-T003) before Tests and Core
- Tests (T004-T008) before Core (T009-T021)
- Contracts (T009) before content files (T010-T012) and build (T013)
- API Integration (T027-T031) before page integration (T032-T033)
- Build (T013) before CI pass checks (T023)
- Core pages (T014-T017) before polish meta (T025)
- API service (T027) before specific API implementations (T028-T029)
- Error handling (T030) before page integration (T032-T033)

## Parallel Example
```
# Launch independent content tasks together once contracts exist:
Task: "T010 Add 20 mocked episodes" → /Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/content/episodes.json
Task: "T011 Add About page content" → /Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/content/about.md
Task: "T012 Add FAQ content" → /Users/minhkhaihuynh/Documents/Git/podsite/app/frontend/site/content/faq.json

# Launch API integration tasks in parallel:
Task: "T027 Create API service module" → /app/frontend/site/src/scripts/api.mjs
Task: "T028 Implement featured episode API" → /app/frontend/site/src/scripts/featured-episode.mjs
Task: "T029 Implement episodes list API" → /app/frontend/site/src/scripts/episodes-list.mjs
Task: "T030 Add error handling" → /app/frontend/site/src/scripts/error-handler.mjs

# Launch validators setup in parallel:
Task: "T004 Create HTML validation config" → /app/frontend/.htmlvalidate.json
Task: "T005 Create link checking config" → /app/frontend/linkinator.config.json
Task: "T006 Create Lighthouse CI config" → /app/frontend/lighthouserc.json
Task: "T007 Create axe test runner" → /app/frontend/tests/accessibility/axe.mjs
```

## Notes
- [P] tasks = different files, no dependencies
- Ensure initial CI fails before implementing pages (enforces TDD gates)
- Keep CSP strict: `default-src 'self'`; no secrets; apply SRI if adding third-party assets
- Commit after each task; small, reversible changes

## Validation Checklist
*GATE: Checked by main() before returning*

- [x] All entities in data-model have corresponding content/tasks
- [x] Tests/checkers are defined before implementation
- [x] Parallel tasks are independent and do not touch the same files
- [x] Each task specifies exact file path
- [x] Tasks respect Static-First and Progressive Enhancement principles

---

## Implementation Summary

**Status**: ✅ **ALL TASKS COMPLETED** (T001-T036)

**Core Deliverables**:
- ✅ 4 static HTML pages (Landing, Episodes, About, FAQ)
- ✅ 20 mocked episodes with full metadata
- ✅ Build system with content validation
- ✅ CI/CD pipeline (GitHub Actions)
- ✅ Accessibility, performance, and SEO optimized
- ✅ Progressive enhancement with vanilla JS

**API Integration** (T027-T036):
- ✅ API service module with retry logic and error handling
- ✅ Featured episode API integration with fallback
- ✅ Episodes list API integration with fallback
- ✅ Comprehensive error handling and graceful degradation
- ✅ API integration tests (14/14 passing)
- ✅ Loading states and error UI components
- ✅ Build script updated with fallback content injection
- ✅ Configuration module for API endpoints

**Build Output**: `app/frontend/dist/` directory ready for deployment
**Next Steps**: 
1. Start backend API server (if available)
2. Run `npm run preview` to view locally with API integration
3. Test fallback behavior by stopping backend API


