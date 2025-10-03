package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupHealthTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Add test routes
	router.GET("/health", HealthCheck)
	router.GET("/ready", ReadinessCheck)

	return router
}

func TestHealthCheck(t *testing.T) {
	router := setupHealthTestRouter()

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var health HealthResponse
	err := json.Unmarshal(w.Body.Bytes(), &health)
	assert.NoError(t, err)

	// Check required fields
	assert.Equal(t, "healthy", health.Status)
	assert.NotEmpty(t, health.Timestamp)
	assert.Equal(t, "2.0.0", health.Version)
	assert.NotEmpty(t, health.Uptime)
	assert.NotNil(t, health.System)

	// Check system information
	assert.Contains(t, health.System, "go_version")
	assert.Contains(t, health.System, "num_goroutines")
	assert.Contains(t, health.System, "num_cpu")
}

func TestReadinessCheck(t *testing.T) {
	router := setupHealthTestRouter()

	req, _ := http.NewRequest("GET", "/ready", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var readiness ReadinessResponse
	err := json.Unmarshal(w.Body.Bytes(), &readiness)
	assert.NoError(t, err)

	// Check required fields
	assert.Equal(t, "ready", readiness.Status)
	assert.NotEmpty(t, readiness.Timestamp)
	assert.Equal(t, "ok", readiness.Database)
	assert.Equal(t, "ok", readiness.ExternalAPI)
}

func TestHealthCheckResponseFormat(t *testing.T) {
	router := setupHealthTestRouter()

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var health HealthResponse
	err := json.Unmarshal(w.Body.Bytes(), &health)
	assert.NoError(t, err)

	// Verify field types
	assert.IsType(t, "", health.Status)
	assert.IsType(t, "", health.Timestamp)
	assert.IsType(t, "", health.Version)
	assert.IsType(t, "", health.Uptime)
	assert.IsType(t, map[string]string{}, health.System)

	// Check timestamp format (should be RFC3339)
	_, err = time.Parse(time.RFC3339, health.Timestamp)
	assert.NoError(t, err, "Timestamp should be in RFC3339 format")

	// Check uptime format (should be a duration string)
	_, err = time.ParseDuration(health.Uptime)
	assert.NoError(t, err, "Uptime should be a valid duration string")
}

func TestReadinessCheckResponseFormat(t *testing.T) {
	router := setupHealthTestRouter()

	req, _ := http.NewRequest("GET", "/ready", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var readiness ReadinessResponse
	err := json.Unmarshal(w.Body.Bytes(), &readiness)
	assert.NoError(t, err)

	// Verify field types
	assert.IsType(t, "", readiness.Status)
	assert.IsType(t, "", readiness.Timestamp)
	assert.IsType(t, "", readiness.Database)
	assert.IsType(t, "", readiness.ExternalAPI)

	// Check timestamp format
	_, err = time.Parse(time.RFC3339, readiness.Timestamp)
	assert.NoError(t, err, "Timestamp should be in RFC3339 format")
}

func TestHealthCheckConsistency(t *testing.T) {
	router := setupHealthTestRouter()

	// Test health endpoint multiple times to ensure consistency
	for i := 0; i < 3; i++ {
		req, _ := http.NewRequest("GET", "/health", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var health HealthResponse
		err := json.Unmarshal(w.Body.Bytes(), &health)
		assert.NoError(t, err)
		assert.Equal(t, "healthy", health.Status)
		assert.Equal(t, "2.0.0", health.Version)
	}
}

func TestReadinessCheckConsistency(t *testing.T) {
	router := setupHealthTestRouter()

	// Test readiness endpoint multiple times to ensure consistency
	for i := 0; i < 3; i++ {
		req, _ := http.NewRequest("GET", "/ready", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var readiness ReadinessResponse
		err := json.Unmarshal(w.Body.Bytes(), &readiness)
		assert.NoError(t, err)
		assert.Equal(t, "ready", readiness.Status)
		assert.Equal(t, "ok", readiness.Database)
		assert.Equal(t, "ok", readiness.ExternalAPI)
	}
}

func TestHealthCheckPerformance(t *testing.T) {
	router := setupHealthTestRouter()

	start := time.Now()

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	duration := time.Since(start)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Less(t, duration, 100*time.Millisecond, "Health check should respond quickly")
}

func TestReadinessCheckPerformance(t *testing.T) {
	router := setupHealthTestRouter()

	start := time.Now()

	req, _ := http.NewRequest("GET", "/ready", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	duration := time.Since(start)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Less(t, duration, 100*time.Millisecond, "Readiness check should respond quickly")
}
