# Podsite Frontend

Modern podcast website built as a static site with API integration.

## 🚀 Quick Start

```bash
# Install dependencies
npm install

# Build the site
npm run build

# Start development server
npm run dev

# Preview built site
npm run start
```

## 📁 Project Structure

```
app/frontend/
├── site/
│   ├── content/          # Static content (fallback data)
│   ├── public/           # HTML templates and assets
│   └── src/              # Source scripts and styles
├── scripts/              # Build scripts
├── tests/                # Test suites
└── dist/                 # Built site (generated)
```

## 🛠️ Available Scripts

### Development
- `npm run dev` - Build and start development server
- `npm run build` - Build the static site
- `npm run clean` - Clean build artifacts
- `npm run preview` - Preview built site

### Testing & Validation
- `npm test` - Run all tests and validation
- `npm run test:html` - Validate HTML
- `npm run test:links` - Check internal/external links
- `npm run test:lighthouse` - Performance audit
- `npm run test:a11y` - Accessibility testing
- `npm run test:api` - API integration tests

### Content Management
- `npm run update-content` - Update episode content
- `npm run update-audio` - Download audio files

### Deployment
- `npm run deploy` - Deploy to configured platform
- `npm run deploy:netlify` - Deploy to Netlify
- `npm run deploy:vercel` - Deploy to Vercel
- `npm run deploy:gh-pages` - Deploy to GitHub Pages

## 🏗️ Architecture

### Static-First Approach
- All pages are pre-built as static HTML
- API integration enhances the experience
- Graceful fallback to embedded content

### Progressive Enhancement
- Base functionality works without JavaScript
- API calls provide dynamic content
- Accessible and performant by default

### Performance Budgets
- JavaScript: < 50KB gzipped
- CSS: < 20KB gzipped
- Images: < 500KB per page
- Lighthouse scores: ≥ 90 across all categories

## 🔧 Configuration

### API Integration
Configure backend API endpoints in `site/src/scripts/config.mjs`:

```javascript
export const API_CONFIG = {
  baseUrl: 'http://localhost:3001/api',
  timeout: 5000,
  retries: 3
};
```

### Build Configuration
The build process is configured in `scripts/build.mjs` and includes:
- Content validation against JSON schemas
- Asset optimization and minification
- Sitemap and robots.txt generation
- Fallback content injection

## 🧪 Testing

### Automated Testing
- HTML validation with `html-validate`
- Link checking with `linkinator`
- Accessibility testing with `axe-core`
- Performance auditing with Lighthouse CI

### Manual Testing
1. Test with JavaScript disabled
2. Test on mobile devices
3. Test with slow network connections
4. Test API failure scenarios

## 🚀 Deployment

### Static Hosting Platforms
The built site (`dist/` folder) can be deployed to:
- **Netlify**: `npm run deploy:netlify`
- **Vercel**: `npm run deploy:vercel`
- **GitHub Pages**: `npm run deploy:gh-pages`
- **Cloudflare Pages**: Upload `dist/` folder

### Environment Variables
No environment variables required for the frontend build.

## 📊 Performance

### Lighthouse Scores
Target scores (all ≥ 90):
- Performance: 90+
- Accessibility: 90+
- Best Practices: 90+
- SEO: 90+

### Bundle Sizes
- Initial JavaScript: ~15KB gzipped
- CSS: ~8KB gzipped
- Images: Optimized with responsive `srcset`

## 🔒 Security

### Content Security Policy
```
default-src 'self';
script-src 'self';
style-src 'self' 'unsafe-inline';
img-src 'self' data:;
```

### Privacy
- No tracking by default
- No third-party analytics
- No cookies or local storage

## 🤝 Contributing

1. Make changes in the `site/` directory
2. Run `npm run build` to build the site
3. Run `npm test` to validate changes
4. Submit a pull request

## 📄 License

MIT License - see LICENSE file for details.
