package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/podsite/backend/internal/models"
)

var episodeService = models.NewEpisodeService()

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// GetEpisodes handles GET /api/episodes
// @Summary Get all episodes
// @Description Returns a list of all podcast episodes
// @Tags episodes
// @Produce json
// @Success 200 {array} models.Episode
// @Failure 500 {object} ErrorResponse
// @Router /episodes [get]
func GetEpisodes(c *gin.Context) {
	episodes := episodeService.GetAll()
	c.JSON(http.StatusOK, episodes)
}

// GetFeaturedEpisode handles GET /api/episodes/featured
// @Summary Get featured episode
// @Description Returns the most recent episode as the featured episode
// @Tags episodes
// @Produce json
// @Success 200 {object} models.Episode
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /episodes/featured [get]
func GetFeaturedEpisode(c *gin.Context) {
	episode, err := episodeService.GetFeatured()
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "not_found",
			Message: "No featured episode available",
			Code:    http.StatusNotFound,
		})
		return
	}
	
	c.JSON(http.StatusOK, episode)
}

// GetEpisodeByID handles GET /api/episodes/:id
// @Summary Get episode by ID
// @Description Returns a specific episode by its ID
// @Tags episodes
// @Param id path string true "Episode ID"
// @Produce json
// @Success 200 {object} models.Episode
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /episodes/{id} [get]
func GetEpisodeByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "bad_request",
			Message: "Episode ID is required",
			Code:    http.StatusBadRequest,
		})
		return
	}
	
	// Handle numeric IDs (episode numbers)
	if num, err := strconv.Atoi(id); err == nil {
		id = "ep" + padNumber(num)
	}
	
	episode, err := episodeService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "not_found",
			Message: "Episode not found",
			Code:    http.StatusNotFound,
		})
		return
	}
	
	c.JSON(http.StatusOK, episode)
}

// padNumber pads a number with leading zeros to 3 digits
func padNumber(num int) string {
	if num < 10 {
		return "00" + strconv.Itoa(num)
	} else if num < 100 {
		return "0" + strconv.Itoa(num)
	}
	return strconv.Itoa(num)
}
