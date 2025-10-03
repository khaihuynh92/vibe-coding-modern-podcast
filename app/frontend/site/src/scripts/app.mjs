// Progressive enhancement for episode details and audio playback
(function() {
  'use strict';

  // Global audio player state
  let currentAudio = null;
  let currentPlayButton = null;

  // Handle inline episode details toggle
  const detailButtons = document.querySelectorAll('.btn-details');
  detailButtons.forEach(btn => {
    btn.addEventListener('click', function() {
      const card = this.closest('.episode-card');
      const details = card.querySelector('.episode-details');
      const isHidden = details.hasAttribute('hidden');
      
      if (isHidden) {
        details.removeAttribute('hidden');
        this.textContent = 'Hide Details';
        this.setAttribute('aria-expanded', 'true');
      } else {
        details.setAttribute('hidden', '');
        this.textContent = 'Details';
        this.setAttribute('aria-expanded', 'false');
      }
    });
  });

  // Handle play buttons with enhanced audio player
  const playButtons = document.querySelectorAll('.btn-play');
  
  playButtons.forEach(btn => {
    btn.addEventListener('click', function() {
      const audioUrl = this.dataset.audio || this.dataset.audioUrl;
      
      if (!audioUrl) {
        console.error('No audio URL found for play button');
        return;
      }
      
      // If clicking the same button, toggle play/pause
      if (currentPlayButton === this && currentAudio) {
        if (currentAudio.paused) {
          currentAudio.play();
        } else {
          currentAudio.pause();
        }
        return;
      }
      
      // Stop current audio if playing
      if (currentAudio && !currentAudio.paused) {
        currentAudio.pause();
        currentAudio.currentTime = 0;
        if (currentPlayButton) {
          currentPlayButton.textContent = 'Play';
          currentPlayButton.classList.remove('playing');
        }
      }
      
      // Create and play new audio
      currentAudio = new Audio(audioUrl);
      currentPlayButton = this;
      
      // Update button state
      this.textContent = 'Playing...';
      this.classList.add('playing');
      
      currentAudio.play().catch(err => {
        console.error('Playback failed:', err);
        this.textContent = 'Play';
        this.classList.remove('playing');
      });
      
      currentAudio.addEventListener('ended', () => {
        this.textContent = 'Play';
        this.classList.remove('playing');
        currentPlayButton = null;
      });
      
      currentAudio.addEventListener('pause', () => {
        this.textContent = 'Play';
        this.classList.remove('playing');
      });
    });
  });

  // Listen for custom play-episode events from dynamically loaded content
  document.addEventListener('play-episode', function(event) {
    const { episodeId, audioUrl } = event.detail;
    
    if (audioUrl) {
      // Find the play button for this episode and trigger it
      const playButton = document.querySelector(`[data-episode-id="${episodeId}"] .btn-play`);
      if (playButton) {
        playButton.dataset.audioUrl = audioUrl;
        playButton.click();
      }
    }
  });

  // Handle keyboard navigation for play buttons
  document.addEventListener('keydown', function(event) {
    if (event.target.classList.contains('btn-play') && (event.key === 'Enter' || event.key === ' ')) {
      event.preventDefault();
      event.target.click();
    }
  });
})();

