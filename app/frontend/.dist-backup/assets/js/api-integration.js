/**
 * API Integration Main Entry Point
 * This file is loaded in HTML pages to initialize API functionality
 */

// Import API modules
import { initFeaturedEpisode } from '../../src/scripts/featured-episode.mjs';
import { initEpisodesList } from '../../src/scripts/episodes-list.mjs';

// Initialize based on current page
(function() {
  'use strict';

  // Initialize featured episode on landing page
  if (document.querySelector('.featured-episode')) {
    initFeaturedEpisode();
  }

  // Initialize episodes list on episodes page
  if (document.querySelector('.episodes-container')) {
    initEpisodesList();
  }

  // Global audio player event listener
  document.addEventListener('play-episode', (event) => {
    const { episodeId, audioUrl } = event.detail;
    console.log('Playing episode:', episodeId, audioUrl);
    
    // Find or create audio player
    const audioPlayer = document.querySelector('#global-audio-player') || createAudioPlayer();
    
    if (audioUrl) {
      const source = audioPlayer.querySelector('source');
      if (source) {
        source.src = audioUrl;
        audioPlayer.load();
        audioPlayer.play();
      }
    }
  });

  function createAudioPlayer() {
    const audio = document.createElement('audio');
    audio.id = 'global-audio-player';
    audio.controls = true;
    audio.innerHTML = '<source src="" type="audio/mpeg">';
    document.body.appendChild(audio);
    return audio;
  }
})();

