package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/podsite/backend/internal/models"
	"github.com/stretchr/testify/assert"
)

func setupContentTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Add test routes
	api := router.Group("/api")
	{
		api.GET("/about", GetAbout)
		api.GET("/faq", GetFAQ)
	}

	return router
}

func TestGetAbout(t *testing.T) {
	router := setupContentTestRouter()

	req, _ := http.NewRequest("GET", "/api/about", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var about models.AboutContent
	err := json.Unmarshal(w.Body.Bytes(), &about)
	assert.NoError(t, err)

	// Check that about content has all required fields
	assert.NotEmpty(t, about.Title)
	assert.NotEmpty(t, about.Description)
	assert.NotEmpty(t, about.Mission)
	assert.NotEmpty(t, about.WhoWeAre)
	assert.NotNil(t, about.WhatWeCover)
	assert.NotEmpty(t, about.JoinCommunity)

	// Check that WhatWeCover is not empty
	assert.Greater(t, len(about.WhatWeCover), 0)
}

func TestGetFAQ(t *testing.T) {
	router := setupContentTestRouter()

	req, _ := http.NewRequest("GET", "/api/faq", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var faq models.FAQContent
	err := json.Unmarshal(w.Body.Bytes(), &faq)
	assert.NoError(t, err)

	// Check that FAQ content has items
	assert.NotNil(t, faq.Items)
	assert.Greater(t, len(faq.Items), 0)

	// Check that each FAQ item has required fields
	for _, item := range faq.Items {
		assert.NotEmpty(t, item.Question)
		assert.NotEmpty(t, item.Answer)
	}
}

func TestGetAboutResponseFormat(t *testing.T) {
	router := setupContentTestRouter()

	req, _ := http.NewRequest("GET", "/api/about", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var about models.AboutContent
	err := json.Unmarshal(w.Body.Bytes(), &about)
	assert.NoError(t, err)

	// Verify specific field types and content
	assert.IsType(t, "", about.Title)
	assert.IsType(t, "", about.Description)
	assert.IsType(t, "", about.Mission)
	assert.IsType(t, "", about.WhoWeAre)
	assert.IsType(t, []string{}, about.WhatWeCover)
	assert.IsType(t, "", about.JoinCommunity)

	// Check that strings are not empty
	assert.Greater(t, len(about.Title), 0)
	assert.Greater(t, len(about.Description), 0)
	assert.Greater(t, len(about.Mission), 0)
	assert.Greater(t, len(about.WhoWeAre), 0)
	assert.Greater(t, len(about.JoinCommunity), 0)
}

func TestGetFAQResponseFormat(t *testing.T) {
	router := setupContentTestRouter()

	req, _ := http.NewRequest("GET", "/api/faq", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var faq models.FAQContent
	err := json.Unmarshal(w.Body.Bytes(), &faq)
	assert.NoError(t, err)

	// Verify structure
	assert.IsType(t, []models.FAQItem{}, faq.Items)

	// Check that we have a reasonable number of FAQ items
	assert.GreaterOrEqual(t, len(faq.Items), 5) // Should have at least 5 FAQ items

	// Check that each item has proper content
	for i, item := range faq.Items {
		assert.NotEmpty(t, item.Question, "FAQ item %d should have a question", i)
		assert.NotEmpty(t, item.Answer, "FAQ item %d should have an answer", i)
		assert.Greater(t, len(item.Question), 10, "FAQ question %d should be substantial", i)
		assert.Greater(t, len(item.Answer), 20, "FAQ answer %d should be substantial", i)
	}
}

func TestContentEndpointsConsistency(t *testing.T) {
	router := setupContentTestRouter()

	// Test about endpoint multiple times to ensure consistency
	for i := 0; i < 3; i++ {
		req, _ := http.NewRequest("GET", "/api/about", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var about models.AboutContent
		err := json.Unmarshal(w.Body.Bytes(), &about)
		assert.NoError(t, err)
		assert.NotEmpty(t, about.Title)
	}

	// Test FAQ endpoint multiple times to ensure consistency
	for i := 0; i < 3; i++ {
		req, _ := http.NewRequest("GET", "/api/faq", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var faq models.FAQContent
		err := json.Unmarshal(w.Body.Bytes(), &faq)
		assert.NoError(t, err)
		assert.Greater(t, len(faq.Items), 0)
	}
}
