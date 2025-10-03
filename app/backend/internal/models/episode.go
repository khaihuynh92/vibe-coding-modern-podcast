package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

// Episode represents a podcast episode
type Episode struct {
	ID          string    `json:"id"`
	Number      int       `json:"number"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Duration    string    `json:"duration"`
	PublishDate string    `json:"publishDate"`
	ArtworkURL  string    `json:"artworkUrl"`
	ArtworkAlt  string    `json:"artworkAlt,omitempty"`
	AudioURL    string    `json:"audioUrl"`
	Tags        []string  `json:"tags"`
}

// EpisodeService handles episode data operations
type EpisodeService struct {
	episodes []Episode
}

// NewEpisodeService creates a new episode service
func NewEpisodeService() *EpisodeService {
	service := &EpisodeService{}
	if err := service.loadEpisodes(); err != nil {
		// If loading fails, use default episodes
		service.episodes = getDefaultEpisodes()
	}
	return service
}

// GetAll returns all episodes sorted by number (descending)
func (s *EpisodeService) GetAll() []Episode {
	episodes := make([]Episode, len(s.episodes))
	copy(episodes, s.episodes)
	
	// Sort by episode number descending (newest first)
	sort.Slice(episodes, func(i, j int) bool {
		return episodes[i].Number > episodes[j].Number
	})
	
	return episodes
}

// GetByID returns an episode by its ID
func (s *EpisodeService) GetByID(id string) (*Episode, error) {
	for _, episode := range s.episodes {
		if episode.ID == id {
			return &episode, nil
		}
	}
	return nil, fmt.Errorf("episode with ID %s not found", id)
}

// GetFeatured returns the most recent episode as featured
func (s *EpisodeService) GetFeatured() (*Episode, error) {
	if len(s.episodes) == 0 {
		return nil, fmt.Errorf("no episodes available")
	}
	
	// Find the episode with the highest number
	var featured *Episode
	for i := range s.episodes {
		if featured == nil || s.episodes[i].Number > featured.Number {
			featured = &s.episodes[i]
		}
	}
	
	return featured, nil
}

// loadEpisodes loads episodes from the frontend content file
func (s *EpisodeService) loadEpisodes() error {
	// Try to load from the frontend content directory
	contentPath := filepath.Join("..", "..", "frontend", "site", "content", "episodes.json")
	
	// If that doesn't exist, try relative to current working directory
	if _, err := os.Stat(contentPath); os.IsNotExist(err) {
		contentPath = filepath.Join("app", "frontend", "site", "content", "episodes.json")
	}
	
	file, err := os.Open(contentPath)
	if err != nil {
		return fmt.Errorf("failed to open episodes file: %w", err)
	}
	defer file.Close()
	
	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read episodes file: %w", err)
	}
	
	if err := json.Unmarshal(data, &s.episodes); err != nil {
		return fmt.Errorf("failed to parse episodes JSON: %w", err)
	}
	
	return nil
}

// getDefaultEpisodes returns a set of default episodes if loading fails
func getDefaultEpisodes() []Episode {
	return []Episode{
		{
			ID:          "ep001",
			Number:      1,
			Title:       "Welcome to Our Podcast",
			Description: "In our inaugural episode, we introduce ourselves and share what you can expect from this podcast.",
			Duration:    "25:30",
			PublishDate: "2025-01-01",
			ArtworkURL:  "/assets/images/ep001.svg",
			ArtworkAlt:  "Episode 1 artwork",
			AudioURL:    "/assets/audio/mock.mp3",
			Tags:        []string{"introduction", "welcome"},
		},
		{
			ID:          "ep002",
			Number:      2,
			Title:       "Getting Started",
			Description: "We dive into the basics and share some fundamental concepts.",
			Duration:    "32:15",
			PublishDate: "2025-01-08",
			ArtworkURL:  "/assets/images/ep002.svg",
			ArtworkAlt:  "Episode 2 artwork",
			AudioURL:    "/assets/audio/mock.mp3",
			Tags:        []string{"basics", "fundamentals"},
		},
	}
}
