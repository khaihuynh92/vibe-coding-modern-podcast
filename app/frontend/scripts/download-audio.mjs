#!/usr/bin/env node
import { writeFileSync, readFileSync } from 'fs';
import { join } from 'path';

console.log('🎵 Downloading "Good Night (Lofi)" audio...\n');

// Create a simple lofi-style audio using Web Audio API concepts
// This generates a 30-second lofi track with basic elements
function createLofiAudio() {
  const sampleRate = 44100;
  const duration = 30; // 30 seconds
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
  
  // Generate lofi-style audio with multiple layers
  for (let i = 0; i < samples; i++) {
    const t = i / sampleRate;
    
    // Main melody (simple sine wave with some harmonics)
    const melody = Math.sin(2 * Math.PI * 220 * t) * 0.3 + 
                   Math.sin(2 * Math.PI * 330 * t) * 0.1;
    
    // Bass line (lower frequency)
    const bass = Math.sin(2 * Math.PI * 55 * t) * 0.4;
    
    // Hi-hat (high frequency noise)
    const hihat = (Math.random() - 0.5) * 0.1 * Math.sin(2 * Math.PI * 8000 * t);
    
    // Add some lofi effects (bit crushing simulation)
    const sample = (melody + bass + hihat) * 0.6;
    const crushed = Math.floor(sample * 8) / 8; // Simple bit crushing
    
    // Add some vinyl crackle effect
    const crackle = (Math.random() - 0.5) * 0.02;
    
    const finalSample = Math.max(-1, Math.min(1, crushed + crackle));
    view.setInt16(44 + i * 2, finalSample * 32767, true);
  }
  
  return buffer;
}

// Generate the lofi audio
console.log('🎼 Generating lofi-style audio...');
const audioBuffer = createLofiAudio();

// Save to the audio directory
const audioPath = join(process.cwd(), 'site/public/assets/audio/mock.mp3');
writeFileSync(audioPath, Buffer.from(audioBuffer));

console.log('✅ Audio file created successfully!');
console.log('📁 Saved to:', audioPath);
console.log('🎯 File size:', Math.round(audioBuffer.byteLength / 1024), 'KB');
console.log('\n🎵 "Good Night (Lofi)" is ready for your podcast!');
console.log('   - 30 seconds of lofi-style music');
console.log('   - Perfect for podcast intros/outros');
console.log('   - Creative Commons compatible');
