package middleware

import (
	"compress/gzip"
	"strings"

	"github.com/gin-gonic/gin"
)

// gzipWriter wraps gin.ResponseWriter with gzip compression
type gzipWriter struct {
	gin.ResponseWriter
	writer *gzip.Writer
}

// Write writes data to the gzip writer
func (w *gzipWriter) Write(data []byte) (int, error) {
	return w.writer.Write(data)
}

// WriteString writes a string to the gzip writer
func (w *gzipWriter) WriteString(s string) (int, error) {
	return w.writer.Write([]byte(s))
}

// Close closes the gzip writer
func (w *gzipWriter) Close() error {
	return w.writer.Close()
}

// Compression returns a Gin middleware for response compression
func Compression() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if client accepts gzip encoding
		if !strings.Contains(c.GetHeader("Accept-Encoding"), "gzip") {
			c.Next()
			return
		}

		// Check if response should be compressed
		contentType := c.GetHeader("Content-Type")
		if !shouldCompress(contentType) {
			c.Next()
			return
		}

		// Create gzip writer
		gzWriter := gzip.NewWriter(c.Writer)
		defer gzWriter.Close()

		// Wrap response writer
		c.Writer = &gzipWriter{
			ResponseWriter: c.Writer,
			writer:         gzWriter,
		}

		// Set headers
		c.Header("Content-Encoding", "gzip")
		c.Header("Vary", "Accept-Encoding")

		c.Next()
	}
}

// shouldCompress determines if content should be compressed
func shouldCompress(contentType string) bool {
	// List of content types that should be compressed
	compressibleTypes := []string{
		"application/json",
		"application/javascript",
		"application/xml",
		"text/css",
		"text/html",
		"text/javascript",
		"text/plain",
		"text/xml",
	}

	for _, t := range compressibleTypes {
		if strings.HasPrefix(contentType, t) {
			return true
		}
	}

	return false
}

// CompressionLevel returns a Gin middleware with custom compression level
func CompressionLevel(level int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if client accepts gzip encoding
		if !strings.Contains(c.GetHeader("Accept-Encoding"), "gzip") {
			c.Next()
			return
		}

		// Check if response should be compressed
		contentType := c.GetHeader("Content-Type")
		if !shouldCompress(contentType) {
			c.Next()
			return
		}

		// Create gzip writer with custom level
		gzWriter, err := gzip.NewWriterLevel(c.Writer, level)
		if err != nil {
			c.Next()
			return
		}
		defer gzWriter.Close()

		// Wrap response writer
		c.Writer = &gzipWriter{
			ResponseWriter: c.Writer,
			writer:         gzWriter,
		}

		// Set headers
		c.Header("Content-Encoding", "gzip")
		c.Header("Vary", "Accept-Encoding")

		c.Next()
	}
}
