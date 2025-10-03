package models

import (
	"testing"
)

func TestNewEpisodeService(t *testing.T) {
	service := NewEpisodeService()
	if service == nil {
		t.Fatal("NewEpisodeService returned nil")
	}
	
	episodes := service.GetAll()
	if len(episodes) == 0 {
		t.Error("Expected at least some episodes, got none")
	}
}

func TestGetFeatured(t *testing.T) {
	service := NewEpisodeService()
	
	featured, err := service.GetFeatured()
	if err != nil {
		t.Fatalf("GetFeatured returned error: %v", err)
	}
	
	if featured == nil {
		t.Fatal("GetFeatured returned nil episode")
	}
	
	if featured.ID == "" {
		t.Error("Featured episode has empty ID")
	}
	
	if featured.Title == "" {
		t.Error("Featured episode has empty title")
	}
}

func TestGetByID(t *testing.T) {
	service := NewEpisodeService()
	
	// Test cases
	tests := []struct {
		name        string
		id          string
		expectError bool
	}{
		{
			name:        "Valid ID",
			id:          "ep001",
			expectError: false,
		},
		{
			name:        "Invalid ID",
			id:          "nonexistent",
			expectError: true,
		},
		{
			name:        "Empty ID",
			id:          "",
			expectError: true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			episode, err := service.GetByID(tt.id)
			
			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				if episode != nil {
					t.Error("Expected nil episode but got one")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if episode == nil {
					t.Error("Expected episode but got nil")
				}
				if episode != nil && episode.ID != tt.id {
					t.Errorf("Expected episode ID %s, got %s", tt.id, episode.ID)
				}
			}
		})
	}
}

func TestGetAllSorting(t *testing.T) {
	service := NewEpisodeService()
	episodes := service.GetAll()
	
	if len(episodes) < 2 {
		t.Skip("Need at least 2 episodes to test sorting")
	}
	
	// Check that episodes are sorted by number descending
	for i := 0; i < len(episodes)-1; i++ {
		if episodes[i].Number < episodes[i+1].Number {
			t.Errorf("Episodes not sorted correctly: episode %d comes before episode %d", 
				episodes[i].Number, episodes[i+1].Number)
		}
	}
}

// Benchmark test
func BenchmarkGetAll(b *testing.B) {
	service := NewEpisodeService()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = service.GetAll()
	}
}

func BenchmarkGetFeatured(b *testing.B) {
	service := NewEpisodeService()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = service.GetFeatured()
	}
}
