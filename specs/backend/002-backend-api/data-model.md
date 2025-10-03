# Data Model: Backend API for Podcast Website

## Overview

This document defines the data models and API contracts for the podcast website backend API. The API serves episode data, static content, and health information to the frontend application.

## Core Entities

### Episode

Represents a podcast episode with complete metadata.

**Attributes:**
- `id` (string): Unique episode identifier (format: "ep001")
- `number` (integer): Episode number (1-based)
- `title` (string): Episode title (max 200 chars)
- `description` (string): Episode description (max 1000 chars)
- `duration` (string): Duration in MM:SS format
- `publishDate` (string): Publish date in YYYY-MM-DD format
- `artworkUrl` (string): URL to episode artwork
- `artworkAlt` (string): Alt text for accessibility
- `audioUrl` (string): URL to episode audio file
- `tags` (array): Array of tag strings

**Example:**
```json
{
  "id": "ep001",
  "number": 1,
  "title": "Welcome to Our Podcast",
  "description": "In our inaugural episode, we introduce ourselves and share what you can expect from this podcast.",
  "duration": "25:30",
  "publishDate": "2025-01-01",
  "artworkUrl": "/assets/images/ep001.svg",
  "artworkAlt": "Episode 1 artwork",
  "audioUrl": "/assets/audio/mock.mp3",
  "tags": ["introduction", "welcome"]
}
```

### AboutContent

Represents structured about page content.

**Attributes:**
- `title` (string): About page title
- `description` (string): Brief description
- `mission` (string): Mission statement
- `whoWeAre` (string): Team information
- `whatWeCover` (array): List of covered topics
- `joinCommunity` (string): Community information

**Example:**
```json
{
  "title": "About Our Podcast",
  "description": "Welcome to our podcast—a space where we explore the art, science, and business of audio storytelling.",
  "mission": "We're dedicated to demystifying the podcasting world and providing actionable insights for creators at every stage of their journey.",
  "whoWeAre": "Our team brings together years of experience in audio production, content creation, and digital media.",
  "whatWeCover": [
    "Production techniques and sound design secrets",
    "Audience growth strategies that actually work"
  ],
  "joinCommunity": "We believe podcasting is better together. Join thousands of creators who tune in each week to level up their craft."
}
```

### FAQItem

Represents a single FAQ question and answer.

**Attributes:**
- `question` (string): FAQ question
- `answer` (string): FAQ answer

**Example:**
```json
{
  "question": "How often do you release new episodes?",
  "answer": "We release a new episode every week, typically on Sundays. Occasionally, we'll drop bonus episodes or special interviews between our regular schedule."
}
```

### FAQContent

Represents the complete FAQ page content.

**Attributes:**
- `items` (array): Array of FAQItem objects

**Example:**
```json
{
  "items": [
    {
      "question": "How often do you release new episodes?",
      "answer": "We release a new episode every week, typically on Sundays."
    },
    {
      "question": "Where can I listen to the podcast?",
      "answer": "Our podcast is available on all major platforms including Apple Podcasts, Spotify, Google Podcasts, and directly on this website."
    }
  ]
}
```

### HealthResponse

Represents health check response.

**Attributes:**
- `status` (string): Health status ("healthy" or "unhealthy")
- `timestamp` (string): Response timestamp in ISO 8601 format
- `version` (string): API version
- `uptime` (string): Server uptime duration
- `system` (object): System information

**Example:**
```json
{
  "status": "healthy",
  "timestamp": "2025-10-03T02:10:32Z",
  "version": "2.0.0",
  "uptime": "4.972645785s",
  "system": {
    "go_version": "go1.25.1",
    "num_goroutines": "5",
    "num_cpu": "8"
  }
}
```

### ErrorResponse

Represents error response.

**Attributes:**
- `error` (string): Error type identifier
- `message` (string): Human-readable error message
- `code` (integer): HTTP status code

**Example:**
```json
{
  "error": "not_found",
  "message": "Episode not found",
  "code": 404
}
```

## API Endpoints

### Episodes
- `GET /api/episodes` → Array of Episode objects
- `GET /api/episodes/featured` → Single Episode object
- `GET /api/episodes/:id` → Single Episode object

### Content
- `GET /api/about` → AboutContent object
- `GET /api/faq` → FAQContent object

### Health & Monitoring
- `GET /health` → HealthResponse object
- `GET /ready` → HealthResponse object

## Data Sources

The API loads data from the following sources with fallback to default content:

1. **Episodes**: `app/frontend/site/content/episodes.json`
2. **About**: `app/frontend/site/content/about.md` (parsed to structured format)
3. **FAQ**: `app/frontend/site/content/faq.json`

## Validation Rules

- Episode IDs must match pattern `^ep\d{3}$`
- Episode numbers must be positive integers
- Duration must match pattern `^\d{1,2}:\d{2}$`
- Publish dates must match pattern `^\d{4}-\d{2}-\d{2}$`
- All required fields must be present
- String lengths must not exceed specified maximums
- Arrays must contain valid items according to their schemas

## Error Handling

All endpoints return appropriate HTTP status codes:
- 200: Success
- 400: Bad Request (invalid parameters)
- 404: Not Found (resource doesn't exist)
- 500: Internal Server Error (server error)

Error responses follow the ErrorResponse schema and include:
- Error type identifier
- Human-readable message
- HTTP status code
