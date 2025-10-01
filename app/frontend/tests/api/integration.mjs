#!/usr/bin/env node
/**
 * API Integration Tests
 * Tests the API integration with error handling and fallback behavior
 */

import { API, checkAPIHealth, APIError } from '../../site/src/scripts/api.mjs';
import { API_CONFIG, ENDPOINTS } from '../../site/src/scripts/config.mjs';

let testsPassed = 0;
let testsFailed = 0;

/**
 * Test result logger
 */
function logTest(name, passed, error = null) {
  if (passed) {
    console.log(`✅ ${name}`);
    testsPassed++;
  } else {
    console.error(`❌ ${name}`);
    if (error) {
      console.error(`   Error: ${error.message}`);
    }
    testsFailed++;
  }
}

/**
 * Test API configuration
 */
function testConfiguration() {
  console.log('\n📋 Testing API Configuration...\n');

  try {
    logTest(
      'API_CONFIG.baseUrl is defined',
      typeof API_CONFIG.baseUrl === 'string' && API_CONFIG.baseUrl.length > 0
    );

    logTest(
      'API_CONFIG.timeout is a number',
      typeof API_CONFIG.timeout === 'number' && API_CONFIG.timeout > 0
    );

    logTest(
      'ENDPOINTS.featuredEpisode is defined',
      typeof ENDPOINTS.featuredEpisode === 'string'
    );

    logTest(
      'ENDPOINTS.episodes is defined',
      typeof ENDPOINTS.episodes === 'string'
    );

    logTest(
      'ENDPOINTS.episodeById is a function',
      typeof ENDPOINTS.episodeById === 'function'
    );
  } catch (error) {
    logTest('Configuration tests', false, error);
  }
}

/**
 * Test API service methods exist
 */
function testAPIServiceInterface() {
  console.log('\n🔌 Testing API Service Interface...\n');

  try {
    logTest(
      'API.getFeaturedEpisode is a function',
      typeof API.getFeaturedEpisode === 'function'
    );

    logTest(
      'API.getEpisodes is a function',
      typeof API.getEpisodes === 'function'
    );

    logTest(
      'API.getEpisodeById is a function',
      typeof API.getEpisodeById === 'function'
    );

    logTest(
      'checkAPIHealth is a function',
      typeof checkAPIHealth === 'function'
    );
  } catch (error) {
    logTest('API interface tests', false, error);
  }
}

/**
 * Test API error handling
 */
async function testErrorHandling() {
  console.log('\n🛡️  Testing Error Handling...\n');

  // Save original config
  const originalBaseUrl = API_CONFIG.baseUrl;
  const originalTimeout = API_CONFIG.timeout;
  const originalRetryAttempts = API_CONFIG.retry.maxAttempts;

  try {
    // Test with invalid URL (should fail gracefully)
    API_CONFIG.baseUrl = 'http://invalid-url-that-does-not-exist.test';
    API_CONFIG.timeout = 1000;
    API_CONFIG.retry.maxAttempts = 1;

    try {
      await API.getFeaturedEpisode();
      logTest('API handles network errors', false);
    } catch (error) {
      logTest(
        'API handles network errors',
        error instanceof Error
      );
    }

    // Test APIError class
    const apiError = new APIError('Test error', 404, { message: 'Not found' });
    logTest(
      'APIError instance has correct properties',
      apiError.status === 404 && apiError.data.message === 'Not found'
    );

  } catch (error) {
    logTest('Error handling tests', false, error);
  } finally {
    // Restore original config
    API_CONFIG.baseUrl = originalBaseUrl;
    API_CONFIG.timeout = originalTimeout;
    API_CONFIG.retry.maxAttempts = originalRetryAttempts;
  }
}

/**
 * Test API health check
 */
async function testHealthCheck() {
  console.log('\n🏥 Testing API Health Check...\n');

  try {
    const isHealthy = await checkAPIHealth();
    logTest(
      'checkAPIHealth returns a boolean',
      typeof isHealthy === 'boolean'
    );

    // Note: May fail if backend is not running, which is expected
    if (!isHealthy) {
      console.log('   ℹ️  Backend API is not available (this is expected if backend is not running)');
    }
  } catch (error) {
    logTest('Health check tests', false, error);
  }
}

/**
 * Test endpoint URL construction
 */
function testEndpointConstruction() {
  console.log('\n🔗 Testing Endpoint Construction...\n');

  try {
    const episodeId = '123';
    const episodeUrl = ENDPOINTS.episodeById(episodeId);
    
    logTest(
      'episodeById constructs correct URL',
      episodeUrl === `/episodes/${episodeId}`
    );

    logTest(
      'Episode URL contains episode ID',
      episodeUrl.includes(episodeId)
    );
  } catch (error) {
    logTest('Endpoint construction tests', false, error);
  }
}

/**
 * Run all tests
 */
async function runTests() {
  console.log('🧪 API Integration Tests\n');
  console.log('═'.repeat(50));

  testConfiguration();
  testAPIServiceInterface();
  testEndpointConstruction();
  await testHealthCheck();
  await testErrorHandling();

  // Summary
  console.log('\n' + '═'.repeat(50));
  console.log('\n📊 Test Summary:\n');
  console.log(`   Passed: ${testsPassed}`);
  console.log(`   Failed: ${testsFailed}`);
  console.log(`   Total:  ${testsPassed + testsFailed}`);

  if (testsFailed > 0) {
    console.log('\n❌ Some tests failed\n');
    process.exit(1);
  } else {
    console.log('\n✅ All tests passed\n');
    process.exit(0);
  }
}

// Run tests
runTests().catch((error) => {
  console.error('\n💥 Test runner failed:', error);
  process.exit(1);
});

