#!/usr/bin/env node
import { readFileSync } from 'fs';
import { glob } from 'glob';
import axe from 'axe-core';
import { JSDOM } from 'jsdom';

const htmlFiles = await glob('dist/**/*.html');

let hasErrors = false;

for (const file of htmlFiles) {
  const html = readFileSync(file, 'utf-8');
  const dom = new JSDOM(html, { runScripts: 'dangerously' });
  const { window } = dom;

  // Inject axe-core
  const script = window.document.createElement('script');
  script.textContent = axe.source;
  window.document.head.appendChild(script);

  try {
    const results = await window.axe.run(window.document, {
      runOnly: ['wcag2a', 'wcag2aa', 'wcag21a', 'wcag21aa']
    });

    if (results.violations.length > 0) {
      hasErrors = true;
      console.error(`❌ ${file}: ${results.violations.length} accessibility violations`);
      results.violations.forEach(v => {
        console.error(`  - [${v.impact}] ${v.id}: ${v.description}`);
        console.error(`    Help: ${v.helpUrl}`);
      });
    } else {
      console.log(`✅ ${file}: No accessibility violations`);
    }
  } catch (err) {
    console.error(`Error testing ${file}:`, err.message);
    hasErrors = true;
  }

  dom.window.close();
}

process.exit(hasErrors ? 1 : 0);

