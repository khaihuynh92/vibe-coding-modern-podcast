/**
 * Episodes List Module
 * Handles fetching and rendering the episodes list
 * @module episodes-list
 */

import { API } from './api.mjs';
import { loadFallbackContent, handleAPIError, showLoading, hideLoading } from './error-handler.mjs';

/**
 * Render a single episode item
 * @param {Object} episode - Episode data
 * @returns {string} HTML string for episode item
 */
function renderEpisodeItem(episode) {
  const publishDate = new Date(episode.publishDate).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  });

  return `
    <article class="episode-item" data-episode-id="${episode.id || episode.number}">
      <div class="episode-summary">
        <img 
          src="${episode.artworkUrl || '/assets/images/placeholder.svg'}" 
          alt="${episode.title} artwork" 
          class="episode-artwork"
          loading="lazy"
          width="80"
          height="80"
        />
        <div class="episode-info">
          <h3 class="episode-title">
            <span class="episode-number">Episode ${episode.number}</span>
            ${episode.title}
          </h3>
          <p class="episode-meta">
            <time datetime="${episode.publishDate}">${publishDate}</time>
            <span class="episode-duration">${episode.duration}</span>
          </p>
        </div>
        <button 
          type="button" 
          class="episode-toggle" 
          aria-expanded="false"
          aria-controls="episode-details-${episode.number}"
        >
          <span class="visually-hidden">Show details</span>
          <svg class="icon-chevron" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <polyline points="6 9 12 15 18 9"></polyline>
          </svg>
        </button>
      </div>
      <div 
        class="episode-details" 
        id="episode-details-${episode.number}" 
        hidden
      >
        <p class="episode-description">${episode.description}</p>
        <div class="episode-actions">
          <button 
            type="button" 
            class="btn-play" 
            data-episode-id="${episode.id || episode.number}"
            ${episode.audioUrl ? `data-audio-url="${episode.audioUrl}"` : ''}
          >
            <svg class="icon-play" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <polygon points="5 3 19 12 5 21 5 3"></polygon>
            </svg>
            Play Episode
          </button>
        </div>
      </div>
    </article>
  `;
}

/**
 * Render episodes list into the container
 * @param {Array} episodes - Array of episode objects
 * @param {HTMLElement} container - Container element
 */
function renderEpisodesList(episodes, container) {
  if (!episodes || !Array.isArray(episodes) || !container) {
    console.error('Invalid episodes data or container');
    return;
  }

  // Sort episodes by number (descending - newest first)
  const sortedEpisodes = [...episodes].sort((a, b) => b.number - a.number);

  // Render episodes
  const episodesHTML = sortedEpisodes.map(episode => renderEpisodeItem(episode)).join('');
  
  // Find or create the episodes-list container
  let listContainer = container.querySelector('.episodes-list');
  if (!listContainer) {
    listContainer = document.createElement('div');
    listContainer.className = 'episodes-list';
    container.appendChild(listContainer);
  }
  
  listContainer.innerHTML = episodesHTML;

  // Attach event listeners for expand/collapse
  attachEpisodeListeners(listContainer);
}

/**
 * Attach event listeners to episode items
 * @param {HTMLElement} container - Container element
 */
function attachEpisodeListeners(container) {
  // Toggle episode details
  const toggleButtons = container.querySelectorAll('.episode-toggle');
  toggleButtons.forEach(button => {
    button.addEventListener('click', (e) => {
      const item = e.target.closest('.episode-item');
      const details = item.querySelector('.episode-details');
      const isExpanded = button.getAttribute('aria-expanded') === 'true';

      if (isExpanded) {
        details.hidden = true;
        button.setAttribute('aria-expanded', 'false');
      } else {
        details.hidden = false;
        button.setAttribute('aria-expanded', 'true');
      }
    });
  });

  // Play button handlers
  const playButtons = container.querySelectorAll('.btn-play');
  playButtons.forEach(button => {
    button.addEventListener('click', (e) => {
      const episodeId = button.getAttribute('data-episode-id');
      const audioUrl = button.getAttribute('data-audio-url');
      
      // Dispatch custom event for audio playback
      const event = new CustomEvent('play-episode', {
        detail: { episodeId, audioUrl },
        bubbles: true
      });
      button.dispatchEvent(event);
    });
  });
}

/**
 * Load and display the episodes list
 * @param {HTMLElement} container - Container element for the episodes list
 */
export async function loadEpisodesList(container) {
  if (!container) {
    console.error('Episodes list container not found');
    return;
  }

  const loadingElement = showLoading(container);

  const retryLoad = () => loadEpisodesList(container);

  try {
    // Attempt to fetch from API
    const episodes = await API.getEpisodes();
    hideLoading(loadingElement);
    renderEpisodesList(episodes, container);
  } catch (error) {
    hideLoading(loadingElement);
    
    // Try fallback content
    const usedFallback = await handleAPIError(
      error,
      container,
      async () => {
        const fallbackData = loadFallbackContent('episodes-list-fallback');
        if (fallbackData) {
          renderEpisodesList(fallbackData, container);
          return fallbackData;
        }
        return null;
      },
      retryLoad
    );

    if (!usedFallback) {
      // If no fallback available, keep existing static content
      console.info('Keeping pre-rendered content due to API failure');
      
      // Still attach listeners to pre-rendered content
      const listContainer = container.querySelector('.episodes-list') || container;
      attachEpisodeListeners(listContainer);
    }
  }
}

/**
 * Initialize episodes list loading
 */
export function initEpisodesList() {
  // Wait for DOM to be ready
  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', () => {
      const container = document.querySelector('.episodes-container');
      if (container) {
        loadEpisodesList(container);
      }
    });
  } else {
    const container = document.querySelector('.episodes-container');
    if (container) {
      loadEpisodesList(container);
    }
  }
}

