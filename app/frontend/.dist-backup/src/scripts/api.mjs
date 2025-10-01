/**
 * API Service Module
 * Handles all HTTP requests to the backend API with retry logic and error handling
 * @module api
 */

import { API_CONFIG, ENDPOINTS } from './config.mjs';

/**
 * Custom error class for API errors
 */
export class APIError extends Error {
  constructor(message, status, data) {
    super(message);
    this.name = 'APIError';
    this.status = status;
    this.data = data;
  }
}

/**
 * Sleep utility for retry delays
 */
const sleep = (ms) => new Promise(resolve => setTimeout(resolve, ms));

/**
 * Make an API request with retry logic
 * @param {string} endpoint - API endpoint path
 * @param {Object} options - Fetch options
 * @param {number} attempt - Current retry attempt
 * @returns {Promise<Object>} Response data
 */
async function makeRequest(endpoint, options = {}, attempt = 1) {
  const url = `${API_CONFIG.baseUrl}${endpoint}`;
  const controller = new AbortController();
  const timeoutId = setTimeout(() => controller.abort(), API_CONFIG.timeout);

  try {
    const response = await fetch(url, {
      ...options,
      signal: controller.signal,
      headers: {
        'Content-Type': 'application/json',
        ...options.headers
      }
    });

    clearTimeout(timeoutId);

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}));
      throw new APIError(
        errorData.message || `HTTP ${response.status}: ${response.statusText}`,
        response.status,
        errorData
      );
    }

    return await response.json();
  } catch (error) {
    clearTimeout(timeoutId);

    // Don't retry on client errors (4xx) or if max attempts reached
    if (
      error instanceof APIError && error.status >= 400 && error.status < 500 ||
      attempt >= API_CONFIG.retry.maxAttempts
    ) {
      throw error;
    }

    // Retry with exponential backoff
    const delay = API_CONFIG.retry.delay * Math.pow(API_CONFIG.retry.backoff, attempt - 1);
    console.warn(`API request failed (attempt ${attempt}/${API_CONFIG.retry.maxAttempts}), retrying in ${delay}ms...`, error);
    
    await sleep(delay);
    return makeRequest(endpoint, options, attempt + 1);
  }
}

/**
 * API Service object with methods for each endpoint
 */
export const API = {
  /**
   * Get the featured episode
   * @returns {Promise<Object>} Featured episode data
   */
  async getFeaturedEpisode() {
    return makeRequest(ENDPOINTS.featuredEpisode);
  },

  /**
   * Get all episodes
   * @returns {Promise<Array>} Array of episode objects
   */
  async getEpisodes() {
    return makeRequest(ENDPOINTS.episodes);
  },

  /**
   * Get a specific episode by ID
   * @param {string|number} id - Episode ID
   * @returns {Promise<Object>} Episode data
   */
  async getEpisodeById(id) {
    return makeRequest(ENDPOINTS.episodeById(id));
  }
};

/**
 * Check if the API is available
 * @returns {Promise<boolean>} True if API is reachable
 */
export async function checkAPIHealth() {
  try {
    const controller = new AbortController();
    const timeoutId = setTimeout(() => controller.abort(), 5000);
    
    const response = await fetch(`${API_CONFIG.baseUrl}/health`, {
      signal: controller.signal
    });
    
    clearTimeout(timeoutId);
    return response.ok;
  } catch (error) {
    console.warn('API health check failed:', error);
    return false;
  }
}

