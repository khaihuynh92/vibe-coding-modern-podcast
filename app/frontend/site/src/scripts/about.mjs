/**
 * About Page Module
 * Handles fetching and rendering the about page content
 * @module about
 */

import { API } from './api.mjs';
import { loadFallbackContent, handleAPIError, showLoading, hideLoading } from './error-handler.mjs';

/**
 * Render about content into the page
 * @param {Object} aboutData - About content data
 * @param {HTMLElement} container - Container element
 */
function renderAboutContent(aboutData, container) {
  if (!aboutData || !container) return;

  // Create about content HTML
  const html = `
    <div class="about-content">
      <header class="about-header">
        <h1>${aboutData.title}</h1>
        <p class="about-description">${aboutData.description}</p>
      </header>

      <section class="about-mission">
        <h2>Our Mission</h2>
        <p>${aboutData.mission}</p>
      </section>

      <section class="about-team">
        <h2>Who We Are</h2>
        <p>${aboutData.whoWeAre}</p>
      </section>

      <section class="about-topics">
        <h2>What We Cover</h2>
        <ul class="topics-list">
          ${aboutData.whatWeCover.map(topic => `<li>${topic}</li>`).join('')}
        </ul>
      </section>

      <section class="about-community">
        <h2>Join Our Community</h2>
        <p>${aboutData.joinCommunity}</p>
      </section>
    </div>
  `;

  container.innerHTML = html;
}

/**
 * Load and display the about content
 * @param {HTMLElement} container - Container element for the about content
 */
export async function loadAboutContent(container) {
  if (!container) {
    console.error('About content container not found');
    return;
  }

  const loadingElement = showLoading(container);

  const retryLoad = () => loadAboutContent(container);

  try {
    // Attempt to fetch from API
    const aboutData = await API.getAbout();
    hideLoading(loadingElement);
    renderAboutContent(aboutData, container);
  } catch (error) {
    hideLoading(loadingElement);
    
    // Try fallback content
    const usedFallback = await handleAPIError(
      error,
      container,
      async () => {
        const fallbackData = loadFallbackContent('about-fallback');
        if (fallbackData) {
          renderAboutContent(fallbackData, container);
          return fallbackData;
        }
        return null;
      },
      retryLoad
    );

    if (!usedFallback) {
      // If no fallback available, keep existing static content
      console.info('Keeping pre-rendered content due to API failure');
    }
  }
}

/**
 * Initialize about content loading
 */
export function initAboutContent() {
  // Wait for DOM to be ready
  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', () => {
      const container = document.querySelector('.about-container') || document.querySelector('main');
      if (container) {
        loadAboutContent(container);
      }
    });
  } else {
    const container = document.querySelector('.about-container') || document.querySelector('main');
    if (container) {
      loadAboutContent(container);
    }
  }
}
