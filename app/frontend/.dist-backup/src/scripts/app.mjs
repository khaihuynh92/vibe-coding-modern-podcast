// Progressive enhancement for episode details and audio playback
(function() {
  'use strict';

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

  // Handle play buttons with simple audio player
  let currentAudio = null;
  const playButtons = document.querySelectorAll('.btn-play');
  
  playButtons.forEach(btn => {
    btn.addEventListener('click', function() {
      const audioUrl = this.dataset.audio;
      
      // Stop current audio if playing
      if (currentAudio && !currentAudio.paused) {
        currentAudio.pause();
        currentAudio.currentTime = 0;
      }
      
      // Create and play new audio
      currentAudio = new Audio(audioUrl);
      currentAudio.play().catch(err => {
        console.error('Playback failed:', err);
      });
      
      this.textContent = 'Playing...';
      
      currentAudio.addEventListener('ended', () => {
        this.textContent = 'Play';
      });
      
      currentAudio.addEventListener('pause', () => {
        this.textContent = 'Play';
      });
    });
  });
})();

