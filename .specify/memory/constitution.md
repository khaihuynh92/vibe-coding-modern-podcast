# Podsite Constitution

## Overview

This is the main constitution for the Podsite project, a modern podcast platform built as a monorepo with separate frontend and backend applications. Each application has its own detailed constitution that must be followed.

## Project Architecture

### Monorepo Structure

- **Frontend**: Static site built with modern web technologies, deployed to CDN
- **Backend**: API services providing data and functionality to the frontend
- **Specifications**: Shared contracts and documentation in `/specs/`
- **Independent Deployment**: Each application can be deployed and scaled independently

### Core Principles

1. **Separation of Concerns**: Frontend and backend have distinct responsibilities
2. **Contract-Driven Development**: Well-defined APIs and interfaces between systems
3. **Independent Scalability**: Each application can scale based on its specific needs
4. **Shared Standards**: Common development practices and quality gates

## Application Constitutions

### Frontend Constitution
- **Location**: `/specs/frontend/constitution.md`
- **Scope**: Static site generation, progressive enhancement, performance, accessibility
- **Deployment**: Static hosting (GitHub Pages, Netlify, Vercel, Cloudflare Pages)

### Backend Constitution
- **Location**: `/specs/backend/constitution.md`
- **Scope**: API design, data management, security, scalability
- **Deployment**: Cloud services (AWS, GCP, Azure) or containerized platforms

## Shared Development Standards

### Git Workflow
- Feature branches with pull requests
- Minimum one reviewer approval required
- All changes must reference applicable constitution compliance

### Quality Gates
- Automated testing and validation in CI/CD
- Security scanning and dependency updates
- Performance monitoring and alerting
- Documentation must be kept current

### Governance
- This main constitution supersedes other project conventions
- Application-specific constitutions take precedence for their respective domains
- Amendments require documented proposal, review, and version bump
- All PRs must include constitution compliance notes

## Compliance

All development work must comply with:
1. This main constitution
2. The applicable application constitution (frontend or backend)
3. Any additional project-specific standards

**Version**: 2.0.0 | **Ratified**: 2025-10-01 | **Last Amended**: 2025-10-01