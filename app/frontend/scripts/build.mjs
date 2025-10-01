#!/usr/bin/env node
import { readFileSync, writeFileSync, mkdirSync, cpSync, existsSync } from 'fs';
import { join, dirname } from 'path';
import { fileURLToPath } from 'url';
import Ajv from 'ajv';
import addFormats from 'ajv-formats';

const __dirname = dirname(fileURLToPath(import.meta.url));
const rootDir = join(__dirname, '..');
const siteDir = join(rootDir, 'site');
const distDir = join(rootDir, 'dist');

// Initialize validator
const ajv = new Ajv({ allErrors: true });
addFormats(ajv);

console.log('🏗️  Building podcast site...\n');

// Clean and create dist directory
if (existsSync(distDir)) {
  cpSync(distDir, join(rootDir, '.dist-backup'), { recursive: true, force: true });
}
mkdirSync(distDir, { recursive: true });

// Load and validate content
console.log('📦 Loading content...');
const episodesRaw = readFileSync(join(siteDir, 'content', 'episodes.json'), 'utf-8');
const episodes = JSON.parse(episodesRaw);

const aboutMd = readFileSync(join(siteDir, 'content', 'about.md'), 'utf-8');
const faqRaw = readFileSync(join(siteDir, 'content', 'faq.json'), 'utf-8');
const faq = JSON.parse(faqRaw);

// Validate episodes against schema
const episodeSchema = JSON.parse(
  readFileSync(join(rootDir, '../../specs/frontend/001-i-am-building/contracts/episode.schema.json'), 'utf-8')
);
const validateEpisode = ajv.compile(episodeSchema);

let validationErrors = false;
episodes.forEach((ep, idx) => {
  if (!validateEpisode(ep)) {
    console.error(`❌ Episode ${idx + 1} validation failed:`, validateEpisode.errors);
    validationErrors = true;
  }
});

if (validationErrors) {
  console.error('\n❌ Content validation failed. Fix errors and rebuild.');
  process.exit(1);
}

console.log(`✅ Validated ${episodes.length} episodes`);

// Helper to convert markdown to simple HTML
function mdToHtml(md) {
  return md
    .split('\n\n')
    .map(para => {
      if (para.startsWith('# ')) return `<h1>${para.slice(2)}</h1>`;
      if (para.startsWith('## ')) return `<h2>${para.slice(3)}</h2>`;
      if (para.startsWith('### ')) return `<h3>${para.slice(4)}</h3>`;
      if (para.startsWith('- ')) {
        const items = para.split('\n').map(line => 
          line.startsWith('- ') ? `<li>${line.slice(2).replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>')}</li>` : ''
        ).join('');
        return `<ul>${items}</ul>`;
      }
      return `<p>${para.replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>')}</p>`;
    })
    .join('\n');
}

// Copy static assets
console.log('📁 Copying assets...');
cpSync(join(siteDir, 'public'), distDir, { recursive: true });

// Copy source scripts for API integration
mkdirSync(join(distDir, 'src', 'scripts'), { recursive: true });
cpSync(join(siteDir, 'src', 'scripts'), join(distDir, 'src', 'scripts'), { recursive: true });

// Read templates
const indexTemplate = readFileSync(join(siteDir, 'public', 'index.html'), 'utf-8');
const episodesTemplate = readFileSync(join(siteDir, 'public', 'episodes.html'), 'utf-8');
const aboutTemplate = readFileSync(join(siteDir, 'public', 'about.html'), 'utf-8');
const faqTemplate = readFileSync(join(siteDir, 'public', 'faq.html'), 'utf-8');

// Build index.html with featured episode fallback data only
console.log('📄 Building pages...');
const featured = episodes[episodes.length - 1]; // Latest episode

// Prepare fallback content for API integration (used only on API failure)
const featuredFallback = JSON.stringify(featured);

const indexHtml = indexTemplate
  .replace('{{FEATURED_EPISODE_FALLBACK}}', featuredFallback);

writeFileSync(join(distDir, 'index.html'), indexHtml);

// Build episodes.html with fallback data only (no pre-rendered content)
// Prepare fallback content for API integration (used only on API failure)
const episodesListFallback = JSON.stringify(episodes);

const episodesHtml = episodesTemplate
  .replace('{{EPISODES_LIST_FALLBACK}}', episodesListFallback);

writeFileSync(join(distDir, 'episodes.html'), episodesHtml);

// Build about.html
const aboutHtml = aboutTemplate.replace('{{ABOUT_CONTENT}}', mdToHtml(aboutMd));
writeFileSync(join(distDir, 'about.html'), aboutHtml);

// Build faq.html
const faqItems = faq
  .map((item, idx) => `
    <details class="faq-item">
      <summary>${item.question}</summary>
      <p>${item.answer}</p>
    </details>
  `)
  .join('\n');

const faqHtml = faqTemplate.replace('{{FAQ_ITEMS}}', faqItems);
writeFileSync(join(distDir, 'faq.html'), faqHtml);

// Generate sitemap.xml
console.log('🗺️  Generating sitemap...');
const baseUrl = 'https://example.com'; // Replace with actual domain
const sitemap = `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url><loc>${baseUrl}/</loc><priority>1.0</priority></url>
  <url><loc>${baseUrl}/episodes.html</loc><priority>0.9</priority></url>
  <url><loc>${baseUrl}/about.html</loc><priority>0.7</priority></url>
  <url><loc>${baseUrl}/faq.html</loc><priority>0.6</priority></url>
</urlset>`;

writeFileSync(join(distDir, 'sitemap.xml'), sitemap);

// Generate robots.txt
const robotsTxt = `User-agent: *
Allow: /

Sitemap: ${baseUrl}/sitemap.xml
`;
writeFileSync(join(distDir, 'robots.txt'), robotsTxt);

console.log('\n✅ Build complete! Output in dist/\n');

