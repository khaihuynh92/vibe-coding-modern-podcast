
# Phase 1: Data Model — Modern Podcast Website

## Entities

### Episode
- id: string (unique)
- number: integer (1..n)
- title: string (1..120)
- description: string (1..500)
- duration: string (e.g., "42:35")
- publishDate: string (ISO 8601 date)
- artworkUrl: string (relative path)
- artworkAlt: string (alt text)
- audioUrl: string (relative path to bundled mock audio)
- tags: array<string>

### PageContent
- title: string
- body: markdown | HTML-safe string

### FAQItem
- question: string
- answer: string (markdown | HTML-safe string)

## Validation Rules
- Episode.title, description, duration, publishDate, artworkUrl, artworkAlt are required.
- Episode.audioUrl required (for bundled mock playback).
- FAQItem.question and answer are required.


