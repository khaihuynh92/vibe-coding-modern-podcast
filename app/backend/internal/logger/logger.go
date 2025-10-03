package logger

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Logger wraps logrus.Logger with additional functionality
type Logger struct {
	*logrus.Logger
}

// NewLogger creates a new structured logger
func NewLogger(level string) *Logger {
	logger := logrus.New()

	// Set log level
	switch level {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}

	// Set JSON formatter for production
	if os.Getenv("GO_ENV") == "production" {
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339,
		})
	}

	// Set output to stdout
	logger.SetOutput(os.Stdout)

	return &Logger{Logger: logger}
}

// LogRequest creates a request logger middleware
func (l *Logger) LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Get client IP
		clientIP := c.ClientIP()

		// Get method and status
		method := c.Request.Method
		statusCode := c.Writer.Status()

		// Get body size
		bodySize := c.Writer.Size()

		// Create log entry
		entry := l.WithFields(logrus.Fields{
			"timestamp":  start.Format(time.RFC3339),
			"method":     method,
			"path":       path,
			"query":      raw,
			"status":     statusCode,
			"latency":    latency.String(),
			"latency_ms": float64(latency.Nanoseconds()) / 1000000.0,
			"client_ip":  clientIP,
			"body_size":  bodySize,
			"user_agent": c.Request.UserAgent(),
		})

		// Log based on status code
		if statusCode >= 500 {
			entry.Error("Server error")
		} else if statusCode >= 400 {
			entry.Warn("Client error")
		} else {
			entry.Info("Request completed")
		}
	}
}

// LogError logs an error with context
func (l *Logger) LogError(err error, context map[string]interface{}) {
	entry := l.WithFields(logrus.Fields(context))
	entry.Error(err.Error())
}

// LogInfo logs an info message with context
func (l *Logger) LogInfo(message string, context map[string]interface{}) {
	entry := l.WithFields(logrus.Fields(context))
	entry.Info(message)
}

// LogDebug logs a debug message with context
func (l *Logger) LogDebug(message string, context map[string]interface{}) {
	entry := l.WithFields(logrus.Fields(context))
	entry.Debug(message)
}

// LogWarn logs a warning message with context
func (l *Logger) LogWarn(message string, context map[string]interface{}) {
	entry := l.WithFields(logrus.Fields(context))
	entry.Warn(message)
}

// Global logger instance
var GlobalLogger *Logger

// InitLogger initializes the global logger
func InitLogger(level string) {
	GlobalLogger = NewLogger(level)
}

// GetLogger returns the global logger instance
func GetLogger() *Logger {
	if GlobalLogger == nil {
		GlobalLogger = NewLogger("info")
	}
	return GlobalLogger
}
