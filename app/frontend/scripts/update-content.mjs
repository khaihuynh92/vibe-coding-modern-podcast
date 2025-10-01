#!/usr/bin/env node
import { writeFileSync, mkdirSync, readFileSync } from 'fs';
import { join } from 'path';

console.log('🎨 Updating podcast content with real assets...\n');

// Create a simple podcast intro audio (sine wave tone)
function createAudioBuffer() {
  const sampleRate = 44100;
  const duration = 30; // 30 seconds
  const frequency = 440; // A4 note
  const samples = sampleRate * duration;
  const buffer = new ArrayBuffer(44 + samples * 2);
  const view = new DataView(buffer);
  
  // WAV header
  const writeString = (offset, string) => {
    for (let i = 0; i < string.length; i++) {
      view.setUint8(offset + i, string.charCodeAt(i));
    }
  };
  
  writeString(0, 'RIFF');
  view.setUint32(4, 36 + samples * 2, true);
  writeString(8, 'WAVE');
  writeString(12, 'fmt ');
  view.setUint32(16, 16, true);
  view.setUint16(20, 1, true);
  view.setUint16(22, 1, true);
  view.setUint32(24, sampleRate, true);
  view.setUint32(28, sampleRate * 2, true);
  view.setUint16(32, 2, true);
  view.setUint16(34, 16, true);
  writeString(36, 'data');
  view.setUint32(40, samples * 2, true);
  
  // Generate sine wave
  for (let i = 0; i < samples; i++) {
    const sample = Math.sin(2 * Math.PI * frequency * i / sampleRate) * 0.3;
    view.setInt16(44 + i * 2, sample * 32767, true);
  }
  
  return buffer;
}

// Create podcast episode artwork using SVG
function createEpisodeArtwork(episodeNumber, title) {
  const colors = [
    '#6366f1', '#8b5cf6', '#ec4899', '#ef4444', '#f59e0b',
    '#10b981', '#06b6d4', '#3b82f6', '#8b5cf6', '#ec4899'
  ];
  const color = colors[episodeNumber % colors.length];
  
  return `<?xml version="1.0" encoding="UTF-8"?>
<svg width="400" height="400" viewBox="0 0 400 400" xmlns="http://www.w3.org/2000/svg">
  <defs>
    <linearGradient id="grad${episodeNumber}" x1="0%" y1="0%" x2="100%" y2="100%">
      <stop offset="0%" style="stop-color:${color};stop-opacity:1" />
      <stop offset="100%" style="stop-color:${color}80;stop-opacity:1" />
    </linearGradient>
  </defs>
  <rect width="400" height="400" fill="url(#grad${episodeNumber})"/>
  <circle cx="200" cy="150" r="60" fill="white" opacity="0.9"/>
  <polygon points="180,130 180,170 220,150" fill="${color}"/>
  <text x="200" y="250" text-anchor="middle" fill="white" font-family="Arial, sans-serif" font-size="24" font-weight="bold">
    Episode ${episodeNumber}
  </text>
  <text x="200" y="280" text-anchor="middle" fill="white" font-family="Arial, sans-serif" font-size="16" opacity="0.9">
    ${title.length > 30 ? title.substring(0, 30) + '...' : title}
  </text>
  <text x="200" y="320" text-anchor="middle" fill="white" font-family="Arial, sans-serif" font-size="14" opacity="0.7">
    Podcast
  </text>
</svg>`;
}

// Create Open Graph image
function createOGImage() {
  return `<?xml version="1.0" encoding="UTF-8"?>
<svg width="1200" height="630" viewBox="0 0 1200 630" xmlns="http://www.w3.org/2000/svg">
  <defs>
    <linearGradient id="ogGrad" x1="0%" y1="0%" x2="100%" y2="100%">
      <stop offset="0%" style="stop-color:#6366f1;stop-opacity:1" />
      <stop offset="100%" style="stop-color:#8b5cf6;stop-opacity:1" />
    </linearGradient>
  </defs>
  <rect width="1200" height="630" fill="url(#ogGrad)"/>
  <circle cx="300" cy="200" r="80" fill="white" opacity="0.9"/>
  <polygon points="270,170 270,230 350,200" fill="#6366f1"/>
  <text x="600" y="250" text-anchor="middle" fill="white" font-family="Arial, sans-serif" font-size="48" font-weight="bold">
    Modern Audio Storytelling
  </text>
  <text x="600" y="300" text-anchor="middle" fill="white" font-family="Arial, sans-serif" font-size="24" opacity="0.9">
    Weekly podcast exploring the craft, business, and future of podcasting
  </text>
  <text x="600" y="400" text-anchor="middle" fill="white" font-family="Arial, sans-serif" font-size="20" opacity="0.7">
    Listen on all major platforms
  </text>
</svg>`;
}

// Update audio file
console.log('🎵 Creating podcast intro audio...');
const audioBuffer = createAudioBuffer();
writeFileSync(join(process.cwd(), 'site/public/assets/audio/mock.mp3'), Buffer.from(audioBuffer));

// Update episode artwork
console.log('🖼️  Creating episode artwork...');
const episodes = JSON.parse(readFileSync(join(process.cwd(), 'site/content/episodes.json'), 'utf-8'));

episodes.forEach((episode, index) => {
  const svg = createEpisodeArtwork(episode.number, episode.title);
  writeFileSync(join(process.cwd(), `site/public/assets/images/ep${String(episode.number).padStart(3, '0')}.svg`), svg);
});

// Update Open Graph image
console.log('📱 Creating Open Graph image...');
const ogSvg = createOGImage();
writeFileSync(join(process.cwd(), 'app/frontend/site/public/assets/images/og-image.svg'), ogSvg);

// Update episodes.json to use SVG files
console.log('📝 Updating episode references...');
episodes.forEach(episode => {
  episode.artworkUrl = `/assets/images/ep${String(episode.number).padStart(3, '0')}.svg`;
});

writeFileSync(join(process.cwd(), 'site/content/episodes.json'), JSON.stringify(episodes, null, 2));

// Update HTML templates to use SVG
console.log('🔧 Updating HTML templates...');
const templates = ['index.html', 'episodes.html', 'about.html', 'faq.html'];

templates.forEach(template => {
  const filePath = join(process.cwd(), `site/public/${template}`);
  let content = readFileSync(filePath, 'utf-8');
  
  // Update Open Graph image reference
  content = content.replace(/og-image\.jpg/g, 'og-image.svg');
  content = content.replace(/twitter:image.*og-image\.jpg/g, 'twitter:image" content="https://example.com/assets/images/og-image.svg');
  
  writeFileSync(filePath, content);
});

console.log('\n✅ Content updated successfully!');
console.log('🎯 Next steps:');
console.log('   1. Run: node scripts/build.mjs');
console.log('   2. Run: npm run preview');
console.log('   3. View the updated site with real podcast artwork');
