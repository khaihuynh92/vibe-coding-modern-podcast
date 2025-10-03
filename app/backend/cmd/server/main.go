package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/podsite/backend/internal/config"
	"github.com/podsite/backend/internal/handlers"
	"github.com/podsite/backend/internal/middleware"
)

// @title Podsite API
// @version 2.0.0
// @description Backend API for Podsite podcast website
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.podsite.com/support
// @contact.email support@podsite.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:3001
// @BasePath /api
func main() {
	// Load configuration
	cfg := config.Load()

	// Set Gin mode based on environment
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create Gin router
	router := gin.New()

	// Add middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())
	router.Use(middleware.CORS(cfg.CORSOrigins))
	router.Use(middleware.Security())

	// Health check endpoints
	router.GET("/health", handlers.HealthCheck)
	router.GET("/ready", handlers.ReadinessCheck)

	// API routes
	api := router.Group("/api")
	{
		episodes := api.Group("/episodes")
		{
			episodes.GET("", handlers.GetEpisodes)
			episodes.GET("/featured", handlers.GetFeaturedEpisode)
			episodes.GET("/:id", handlers.GetEpisodeByID)
		}
	}

	// Swagger documentation (only in development)
	if cfg.Environment != "production" {
		// This will be added when we implement Swagger
		// docs.SwaggerInfo.BasePath = "/api"
		// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Create HTTP server
	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on port %s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Give outstanding requests 30 seconds to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
