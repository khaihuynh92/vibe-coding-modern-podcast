package models

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

// AboutContent represents the about page content
type AboutContent struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Mission     string `json:"mission"`
	WhoWeAre    string `json:"whoWeAre"`
	WhatWeCover []string `json:"whatWeCover"`
	JoinCommunity string `json:"joinCommunity"`
}

// FAQItem represents a single FAQ item
type FAQItem struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

// FAQContent represents the FAQ page content
type FAQContent struct {
	Items []FAQItem `json:"items"`
}

// ContentService handles static content operations
type ContentService struct {
	aboutContent *AboutContent
	faqContent   *FAQContent
}

// NewContentService creates a new content service
func NewContentService() *ContentService {
	service := &ContentService{}
	service.loadAboutContent()
	service.loadFAQContent()
	return service
}

// GetAbout returns the about page content
func (s *ContentService) GetAbout() *AboutContent {
	return s.aboutContent
}

// GetFAQ returns the FAQ page content
func (s *ContentService) GetFAQ() *FAQContent {
	return s.faqContent
}

// loadAboutContent loads about content from the frontend content file
func (s *ContentService) loadAboutContent() {
	// Try to load from the frontend content directory
	contentPath := filepath.Join("..", "..", "frontend", "site", "content", "about.md")
	
	// If that doesn't exist, try relative to current working directory
	if _, err := os.Stat(contentPath); os.IsNotExist(err) {
		contentPath = filepath.Join("app", "frontend", "site", "content", "about.md")
	}
	
	file, err := os.Open(contentPath)
	if err != nil {
		// Use default content if file doesn't exist
		s.aboutContent = getDefaultAboutContent()
		return
	}
	defer file.Close()
	
	data, err := io.ReadAll(file)
	if err != nil {
		s.aboutContent = getDefaultAboutContent()
		return
	}
	
	// Parse markdown content and convert to structured data
	s.aboutContent = parseAboutMarkdown(string(data))
}

// loadFAQContent loads FAQ content from the frontend content file
func (s *ContentService) loadFAQContent() {
	// Try to load from the frontend content directory
	contentPath := filepath.Join("..", "..", "frontend", "site", "content", "faq.json")
	
	// If that doesn't exist, try relative to current working directory
	if _, err := os.Stat(contentPath); os.IsNotExist(err) {
		contentPath = filepath.Join("app", "frontend", "site", "content", "faq.json")
	}
	
	file, err := os.Open(contentPath)
	if err != nil {
		// Use default content if file doesn't exist
		s.faqContent = getDefaultFAQContent()
		return
	}
	defer file.Close()
	
	data, err := io.ReadAll(file)
	if err != nil {
		s.faqContent = getDefaultFAQContent()
		return
	}
	
	if err := json.Unmarshal(data, &s.faqContent); err != nil {
		s.faqContent = getDefaultFAQContent()
		return
	}
}

// parseAboutMarkdown parses markdown content and converts to structured data
func parseAboutMarkdown(content string) *AboutContent {
	// This is a simple parser for the specific markdown structure
	// In a production app, you might want to use a proper markdown parser
	
	// For now, return the default content
	// In a real implementation, you'd parse the markdown properly
	return getDefaultAboutContent()
}

// getDefaultAboutContent returns default about content
func getDefaultAboutContent() *AboutContent {
	return &AboutContent{
		Title:       "About Our Podcast",
		Description: "Welcome to our podcast—a space where we explore the art, science, and business of audio storytelling.",
		Mission:     "We're dedicated to demystifying the podcasting world and providing actionable insights for creators at every stage of their journey. Whether you're just starting out or looking to scale your existing show, we've got you covered.",
		WhoWeAre:    "Our team brings together years of experience in audio production, content creation, and digital media. We're passionate about the power of voice and the unique intimacy that podcasting offers.",
		WhatWeCover: []string{
			"Production techniques and sound design secrets",
			"Audience growth strategies that actually work",
			"Monetization approaches for sustainable podcasting",
			"Industry insights from leading voices in audio",
			"Technical know-how without the jargon",
		},
		JoinCommunity: "We believe podcasting is better together. Join thousands of creators who tune in each week to level up their craft. Subscribe on your favorite platform and never miss an episode.",
	}
}

// getDefaultFAQContent returns default FAQ content
func getDefaultFAQContent() *FAQContent {
	return &FAQContent{
		Items: []FAQItem{
			{
				Question: "How often do you release new episodes?",
				Answer:   "We release a new episode every week, typically on Sundays. Occasionally, we'll drop bonus episodes or special interviews between our regular schedule.",
			},
			{
				Question: "Where can I listen to the podcast?",
				Answer:   "Our podcast is available on all major platforms including Apple Podcasts, Spotify, Google Podcasts, and directly on this website. Choose whichever platform works best for you!",
			},
			{
				Question: "Can I suggest a topic or guest?",
				Answer:   "Absolutely! We love hearing from our listeners. Send us your topic ideas or guest suggestions through our contact form or social media channels. We read every message.",
			},
			{
				Question: "Do you have transcripts available?",
				Answer:   "Yes, we provide full transcripts for accessibility. You can find them on each episode's page, usually within 48 hours of release.",
			},
			{
				Question: "How can I support the podcast?",
				Answer:   "The best way to support us is to subscribe, rate, and review on your podcast platform of choice. Sharing episodes with friends who might enjoy them also helps us grow. We also have sponsorship and partnership opportunities available.",
			},
			{
				Question: "Do you take advertising or sponsorships?",
				Answer:   "We do work with select sponsors whose products and services align with our audience's interests. All sponsorships are clearly disclosed, and we only partner with brands we believe in.",
			},
			{
				Question: "Can I use clips from your podcast?",
				Answer:   "Short clips for educational or commentary purposes under fair use are fine. For commercial use or longer excerpts, please contact us for permission. Always provide attribution.",
			},
			{
				Question: "How do I contact the hosts?",
				Answer:   "You can reach us through our contact form, email (listed in episode show notes), or on social media. We try to respond to all messages within a few business days.",
			},
		},
	}
}
