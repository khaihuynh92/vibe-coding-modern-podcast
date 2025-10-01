/**
 * Featured Episode Module
 * Handles fetching and rendering the featured episode
 * @module featured-episode
 */

import { API } from './api.mjs';
import { loadFallbackContent, handleAPIError, showLoading, hideLoading } from './error-handler.mjs';

/**
 * Render episode data into the featured episode container
 * @param {Object} episode - Episode data
 * @param {HTMLElement} container - Container element
 */
function renderFeaturedEpisode(episode, container) {
  if (!episode || !container) return;

  const publishDate = new Date(episode.publishDate).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  });

  // Create featured episode HTML
  const html = `
    <article class="featured-card">
      <img src="${episode.artworkUrl || '/assets/images/placeholder.svg'}" 
           alt="${episode.title} artwork" 
           class="featured-artwork" 
           width="400" 
           height="400"
           loading="eager"
      >
      <div class="featured-info">
        <p class="episode-label">Episode ${episode.number}</p>
        <h3>${episode.title}</h3>
        <p class="episode-description">${episode.description}</p>
        <p class="episode-meta">
          <time datetime="${episode.publishDate}">${publishDate}</time>
          <span class="duration">${episode.duration}</span>
        </p>
        <audio controls src="${episode.audioUrl || '/assets/audio/mock.mp3'}" preload="none">
          <a href="${episode.audioUrl || '/assets/audio/mock.mp3'}">Download audio</a>
        </audio>
      </div>
    </article>
  `;

  container.innerHTML = '<h2 id="featured-heading">Featured Episode</h2>' + html;
}

/**
 * Load and display the featured episode
 * @param {HTMLElement} container - Container element for the featured episode
 */
export async function loadFeaturedEpisode(container) {
  if (!container) {
    console.error('Featured episode container not found');
    return;
  }

  const loadingElement = showLoading(container);

  const retryLoad = () => loadFeaturedEpisode(container);

  try {
    // Attempt to fetch from API
    const episode = await API.getFeaturedEpisode();
    hideLoading(loadingElement);
    renderFeaturedEpisode(episode, container);
  } catch (error) {
    hideLoading(loadingElement);
    
    // Try fallback content
    const usedFallback = await handleAPIError(
      error,
      container,
      async () => {
        const fallbackData = loadFallbackContent('featured-episode-fallback');
        if (fallbackData) {
          renderFeaturedEpisode(fallbackData, container);
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
 * Initialize featured episode loading
 */
export function initFeaturedEpisode() {
  // Wait for DOM to be ready
  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', () => {
      const container = document.querySelector('.featured-episode');
      if (container) {
        loadFeaturedEpisode(container);
      }
    });
  } else {
    const container = document.querySelector('.featured-episode');
    if (container) {
      loadFeaturedEpisode(container);
    }
  }
}

