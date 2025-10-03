package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/podsite/backend/internal/models"
)

var contentService = models.NewContentService()

// GetAbout handles GET /api/about
// @Summary Get about page content
// @Description Returns the about page content including mission, team info, and what we cover
// @Tags content
// @Produce json
// @Success 200 {object} models.AboutContent
// @Failure 500 {object} ErrorResponse
// @Router /about [get]
func GetAbout(c *gin.Context) {
	content := contentService.GetAbout()
	c.JSON(http.StatusOK, content)
}

// GetFAQ handles GET /api/faq
// @Summary Get FAQ page content
// @Description Returns the FAQ page content with all questions and answers
// @Tags content
// @Produce json
// @Success 200 {object} models.FAQContent
// @Failure 500 {object} ErrorResponse
// @Router /faq [get]
func GetFAQ(c *gin.Context) {
	content := contentService.GetFAQ()
	c.JSON(http.StatusOK, content)
}
