package handlers

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string            `json:"status"`
	Timestamp string            `json:"timestamp"`
	Version   string            `json:"version"`
	Uptime    string            `json:"uptime"`
	System    map[string]string `json:"system"`
}

var startTime = time.Now()

// HealthCheck handles GET /health
// @Summary Health check endpoint
// @Description Returns the health status of the API
// @Tags health
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	uptime := time.Since(startTime)
	
	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Version:   "2.0.0",
		Uptime:    uptime.String(),
		System: map[string]string{
			"go_version":      runtime.Version(),
			"num_goroutines":  string(rune(runtime.NumGoroutine())),
			"num_cpu":         string(rune(runtime.NumCPU())),
		},
	}
	
	c.JSON(http.StatusOK, response)
}

// ReadinessResponse represents the readiness check response
type ReadinessResponse struct {
	Status      string `json:"status"`
	Timestamp   string `json:"timestamp"`
	Database    string `json:"database"`
	ExternalAPI string `json:"external_api"`
}

// ReadinessCheck handles GET /ready
// @Summary Readiness check endpoint
// @Description Returns the readiness status of the API and its dependencies
// @Tags health
// @Produce json
// @Success 200 {object} ReadinessResponse
// @Success 503 {object} ReadinessResponse
// @Router /ready [get]
func ReadinessCheck(c *gin.Context) {
	// Check database connectivity (currently using in-memory data)
	dbStatus := "ok"
	
	// Check external API dependencies (none currently)
	apiStatus := "ok"
	
	// Determine overall status
	status := "ready"
	httpStatus := http.StatusOK
	
	if dbStatus != "ok" || apiStatus != "ok" {
		status = "not_ready"
		httpStatus = http.StatusServiceUnavailable
	}
	
	response := ReadinessResponse{
		Status:      status,
		Timestamp:   time.Now().UTC().Format(time.RFC3339),
		Database:    dbStatus,
		ExternalAPI: apiStatus,
	}
	
	c.JSON(httpStatus, response)
}
