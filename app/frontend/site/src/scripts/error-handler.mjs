/**
 * Error Handler Module
 * Provides error handling and fallback content strategies
 * @module error-handler
 */

import { FEATURES } from './config.mjs';

/**
 * Load fallback content from embedded JSON
 * @param {string} contentId - ID of the fallback content element
 * @returns {Object|Array|null} Parsed fallback content or null
 */
export function loadFallbackContent(contentId) {
  try {
    const element = document.getElementById(contentId);
    if (!element) {
      console.warn(`Fallback content element not found: ${contentId}`);
      return null;
    }
    
    const content = element.textContent;
    return JSON.parse(content);
  } catch (error) {
    console.error('Failed to load fallback content:', error);
    return null;
  }
}

/**
 * Show error message to user
 * @param {HTMLElement} container - Container element for error message
 * @param {string} message - Error message to display
 * @param {Object} options - Display options
 */
export function showError(container, message, options = {}) {
  if (!FEATURES.enableErrorUI || !container) return;

  const {
    showRetry = true,
    onRetry = null,
    type = 'error' // 'error', 'warning', 'info'
  } = options;

  const errorDiv = document.createElement('div');
  errorDiv.className = `error-message error-${type}`;
  errorDiv.setAttribute('role', 'alert');
  errorDiv.setAttribute('aria-live', 'polite');
  
  errorDiv.innerHTML = `
    <div class="error-content">
      <svg class="error-icon" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <circle cx="12" cy="12" r="10"></circle>
        <line x1="12" y1="8" x2="12" y2="12"></line>
        <line x1="12" y1="16" x2="12.01" y2="16"></line>
      </svg>
      <p class="error-text">${message}</p>
      ${showRetry && onRetry ? '<button class="error-retry" type="button">Try Again</button>' : ''}
    </div>
  `;

  // Add retry handler
  if (showRetry && onRetry) {
    const retryButton = errorDiv.querySelector('.error-retry');
    retryButton.addEventListener('click', () => {
      errorDiv.remove();
      onRetry();
    });
  }

  container.appendChild(errorDiv);
}

/**
 * Show loading state
 * @param {HTMLElement} container - Container element for loading indicator
 * @returns {HTMLElement} Loading element (to be removed later)
 */
export function showLoading(container) {
  if (!FEATURES.enableLoadingStates || !container) return null;

  const loadingDiv = document.createElement('div');
  loadingDiv.className = 'loading-indicator';
  loadingDiv.setAttribute('role', 'status');
  loadingDiv.setAttribute('aria-live', 'polite');
  loadingDiv.setAttribute('aria-label', 'Loading content');
  
  loadingDiv.innerHTML = `
    <div class="loading-spinner"></div>
    <p class="loading-text">Loading episodes...</p>
  `;

  container.appendChild(loadingDiv);
  return loadingDiv;
}

/**
 * Hide loading state
 * @param {HTMLElement} loadingElement - Loading element to remove
 */
export function hideLoading(loadingElement) {
  if (loadingElement && loadingElement.parentNode) {
    loadingElement.remove();
  }
}

/**
 * Handle API error with fallback strategy
 * @param {Error} error - The error that occurred
 * @param {HTMLElement} container - Container for error/content display
 * @param {Function} fallbackFn - Function to call for fallback content
 * @param {Function} retryFn - Function to call on retry
 * @returns {boolean} True if fallback was used, false otherwise
 */
export async function handleAPIError(error, container, fallbackFn, retryFn) {
  console.error('API Error:', error);

  // Try fallback content if enabled
  if (FEATURES.useFallbackContent && fallbackFn) {
    try {
      const fallbackData = await fallbackFn();
      if (fallbackData) {
        console.info('Using fallback content due to API error');
        showError(container, 'Using cached content. Real-time data temporarily unavailable.', {
          type: 'warning',
          showRetry: true,
          onRetry: retryFn
        });
        return true;
      }
    } catch (fallbackError) {
      console.error('Fallback content also failed:', fallbackError);
    }
  }

  // Show error message
  showError(container, 'Unable to load content. Please try again later.', {
    showRetry: true,
    onRetry: retryFn
  });

  return false;
}

