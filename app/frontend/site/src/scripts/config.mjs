/**
 * Configuration for API endpoints and environment settings
 * @module config
 */

// API Configuration
export const API_CONFIG = {
  // Base URL for API (can be overridden by environment variable)
  baseUrl: typeof window !== 'undefined' && window.API_BASE_URL 
    ? window.API_BASE_URL 
    : 'http://localhost:3001/api',
  
  // Timeout for API requests (ms)
  timeout: 10000,
  
  // Retry configuration
  retry: {
    maxAttempts: 3,
    delay: 1000,
    backoff: 2 // exponential backoff multiplier
  }
};

// Endpoint paths
export const ENDPOINTS = {
  featuredEpisode: '/episodes/featured',
  episodes: '/episodes',
  episodeById: (id) => `/episodes/${id}`,
  about: '/about',
  faq: '/faq',
  health: '/health'
};

// Feature flags
export const FEATURES = {
  useFallbackContent: true, // Enable fallback to static content on API failure
  enableLoadingStates: true, // Show loading indicators
  enableErrorUI: true // Show error messages to users
};

// Environment detection
export const ENV = {
  isDevelopment: typeof window !== 'undefined' && window.location.hostname === 'localhost',
  isProduction: typeof window !== 'undefined' && window.location.hostname !== 'localhost'
};

