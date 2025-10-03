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

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Add test routes
	api := router.Group("/api")
	{
		episodes := api.Group("/episodes")
		{
			episodes.GET("", GetEpisodes)
			episodes.GET("/featured", GetFeaturedEpisode)
			episodes.GET("/:id", GetEpisodeByID)
		}
	}

	return router
}

func TestGetEpisodes(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/episodes", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var episodes []models.Episode
	err := json.Unmarshal(w.Body.Bytes(), &episodes)
	assert.NoError(t, err)
	assert.NotEmpty(t, episodes)

	// Check that episodes are sorted by number descending
	for i := 1; i < len(episodes); i++ {
		assert.GreaterOrEqual(t, episodes[i-1].Number, episodes[i].Number)
	}
}

func TestGetFeaturedEpisode(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/episodes/featured", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var episode models.Episode
	err := json.Unmarshal(w.Body.Bytes(), &episode)
	assert.NoError(t, err)
	assert.NotEmpty(t, episode.ID)
	assert.NotEmpty(t, episode.Title)
}

func TestGetEpisodeByID(t *testing.T) {
	router := setupTestRouter()

	tests := []struct {
		name           string
		episodeID      string
		expectedStatus int
		expectError    bool
	}{
		{
			name:           "Valid episode ID",
			episodeID:      "ep001",
			expectedStatus: http.StatusOK,
			expectError:    false,
		},
		{
			name:           "Valid numeric ID",
			episodeID:      "1",
			expectedStatus: http.StatusOK,
			expectError:    false,
		},
		{
			name:           "Invalid episode ID",
			episodeID:      "ep999",
			expectedStatus: http.StatusNotFound,
			expectError:    true,
		},
		{
			name:           "Invalid episode ID format",
			episodeID:      "invalid",
			expectedStatus: http.StatusNotFound,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/api/episodes/"+tt.episodeID, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if !tt.expectError {
				var episode models.Episode
				err := json.Unmarshal(w.Body.Bytes(), &episode)
				assert.NoError(t, err)
				assert.NotEmpty(t, episode.ID)
			} else {
				var errorResp ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &errorResp)
				assert.NoError(t, err)
				assert.NotEmpty(t, errorResp.Error)
			}
		})
	}
}

func TestGetEpisodesResponseFormat(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/episodes", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var episodes []models.Episode
	err := json.Unmarshal(w.Body.Bytes(), &episodes)
	assert.NoError(t, err)

	// Check that at least one episode has all required fields
	if len(episodes) > 0 {
		episode := episodes[0]
		assert.NotEmpty(t, episode.ID)
		assert.NotEmpty(t, episode.Number)
		assert.NotEmpty(t, episode.Title)
		assert.NotEmpty(t, episode.Description)
		assert.NotEmpty(t, episode.Duration)
		assert.NotEmpty(t, episode.PublishDate)
		assert.NotEmpty(t, episode.ArtworkURL)
		assert.NotEmpty(t, episode.AudioURL)
		assert.NotNil(t, episode.Tags)
	}
}

func TestGetFeaturedEpisodeResponseFormat(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/episodes/featured", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var episode models.Episode
	err := json.Unmarshal(w.Body.Bytes(), &episode)
	assert.NoError(t, err)

	// Check that featured episode has all required fields
	assert.NotEmpty(t, episode.ID)
	assert.NotEmpty(t, episode.Number)
	assert.NotEmpty(t, episode.Title)
	assert.NotEmpty(t, episode.Description)
	assert.NotEmpty(t, episode.Duration)
	assert.NotEmpty(t, episode.PublishDate)
	assert.NotEmpty(t, episode.ArtworkURL)
	assert.NotEmpty(t, episode.AudioURL)
	assert.NotNil(t, episode.Tags)
}
