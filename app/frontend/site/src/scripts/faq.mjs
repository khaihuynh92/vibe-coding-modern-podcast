/**
 * FAQ Page Module
 * Handles fetching and rendering the FAQ content
 * @module faq
 */

import { API } from './api.mjs';
import { loadFallbackContent, handleAPIError, showLoading, hideLoading } from './error-handler.mjs';

/**
 * Render a single FAQ item
 * @param {Object} item - FAQ item data
 * @param {number} index - Item index
 * @returns {string} HTML string for FAQ item
 */
function renderFAQItem(item, index) {
  return `
    <div class="faq-item" data-faq-index="${index}">
      <button 
        type="button" 
        class="faq-question" 
        aria-expanded="false"
        aria-controls="faq-answer-${index}"
        id="faq-question-${index}"
      >
        <span class="faq-question-text">${item.question}</span>
        <svg class="faq-icon" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <polyline points="6 9 12 15 18 9"></polyline>
        </svg>
      </button>
      <div 
        class="faq-answer" 
        id="faq-answer-${index}" 
        hidden
        role="region"
        aria-labelledby="faq-question-${index}"
      >
        <p>${item.answer}</p>
      </div>
    </div>
  `;
}

/**
 * Render FAQ content into the page
 * @param {Object} faqData - FAQ content data
 * @param {HTMLElement} container - Container element
 */
function renderFAQContent(faqData, container) {
  if (!faqData || !container) return;
  
  // Handle both array format and object with items property
  const faqItems = Array.isArray(faqData) ? faqData : (faqData.items || []);
  if (!Array.isArray(faqItems) || faqItems.length === 0) return;

  // Create FAQ content HTML
  const html = `
    <div class="faq-content">
      <header class="faq-header">
        <h1>Frequently Asked Questions</h1>
        <p class="faq-description">Find answers to common questions about our podcast.</p>
      </header>

      <div class="faq-list">
        ${faqItems.map((item, index) => renderFAQItem(item, index)).join('')}
      </div>
    </div>
  `;

  container.innerHTML = html;

  // Attach event listeners for expand/collapse
  attachFAQListeners(container);
}

/**
 * Attach event listeners to FAQ items
 * @param {HTMLElement} container - Container element
 */
function attachFAQListeners(container) {
  const faqItems = container.querySelectorAll('.faq-item');
  
  faqItems.forEach(item => {
    const questionButton = item.querySelector('.faq-question');
    const answerDiv = item.querySelector('.faq-answer');
    
    questionButton.addEventListener('click', () => {
      const isExpanded = questionButton.getAttribute('aria-expanded') === 'true';
      
      if (isExpanded) {
        // Collapse
        answerDiv.hidden = true;
        questionButton.setAttribute('aria-expanded', 'false');
        item.classList.remove('expanded');
      } else {
        // Expand
        answerDiv.hidden = false;
        questionButton.setAttribute('aria-expanded', 'true');
        item.classList.add('expanded');
      }
    });
  });
}

/**
 * Load and display the FAQ content
 * @param {HTMLElement} container - Container element for the FAQ content
 */
export async function loadFAQContent(container) {
  if (!container) {
    console.error('FAQ content container not found');
    return;
  }

  const loadingElement = showLoading(container);

  const retryLoad = () => loadFAQContent(container);

  try {
    // Attempt to fetch from API
    const faqData = await API.getFAQ();
    hideLoading(loadingElement);
    renderFAQContent(faqData, container);
  } catch (error) {
    hideLoading(loadingElement);
    
    // Try fallback content
    const usedFallback = await handleAPIError(
      error,
      container,
      async () => {
        const fallbackData = loadFallbackContent('faq-fallback');
        if (fallbackData) {
          renderFAQContent(fallbackData, container);
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
      attachFAQListeners(container);
    }
  }
}

/**
 * Initialize FAQ content loading
 */
export function initFAQContent() {
  // Wait for DOM to be ready
  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', () => {
      const container = document.querySelector('.faq-container') || document.querySelector('main');
      if (container) {
        loadFAQContent(container);
      }
    });
  } else {
    const container = document.querySelector('.faq-container') || document.querySelector('main');
    if (container) {
      loadFAQContent(container);
    }
  }
}
