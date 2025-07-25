package web

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestLoadNavigation tests loading navigation from JSON file
func TestLoadNavigation(t *testing.T) {
	// Create temporary test directory
	testDir := t.TempDir()
	dataDir := filepath.Join(testDir, "data")
	os.MkdirAll(dataDir, 0755)

	// Create test navigation file
	navData := &Navigation{
		Sections: []NavItem{
			{
				ID:   "platform",
				Name: "Platform",
				Href: "/platform",
			},
			{
				ID:   "solutions",
				Name: "Solutions",
				Children: []NavItem{
					{
						ID:   "by-role",
						Name: "By Role",
						Href: "/solutions/by-role",
					},
				},
			},
		},
	}

	navJSON, _ := json.MarshalIndent(navData, "", "  ")
	os.WriteFile(filepath.Join(dataDir, "nav.json"), navJSON, 0644)

	// Change working directory temporarily
	oldWd, _ := os.Getwd()
	os.Chdir(testDir)
	defer os.Chdir(oldWd)

	// Test loading
	seoService := NewSEOService()
	navService := &NavigationService{seoService: seoService}
	err := navService.LoadNavigation()
	if err != nil {
		t.Fatalf("Failed to load navigation: %v", err)
	}

	// Verify navigation loaded
	if navService.navigation == nil {
		t.Fatal("Navigation not loaded")
	}

	if len(navService.navigation.Sections) != 2 {
		t.Errorf("Expected 2 sections, got %d", len(navService.navigation.Sections))
	}

	// Verify first section
	if navService.navigation.Sections[0].ID != "platform" {
		t.Errorf("Expected first section ID 'platform', got %q", navService.navigation.Sections[0].ID)
	}

	// Verify section with children
	if len(navService.navigation.Sections[1].Children) != 1 {
		t.Errorf("Expected 1 child in solutions section, got %d", len(navService.navigation.Sections[1].Children))
	}
}

// TestGenerateContentNavigation tests generating navigation from content directory
func TestGenerateContentNavigation(t *testing.T) {
	// Create test directory structure
	testDir := t.TempDir()

	// Change to test directory so relative paths work
	oldWd, _ := os.Getwd()
	os.Chdir(testDir)
	defer os.Chdir(oldWd)

	// Use relative path as expected by the navigation code
	contentDir := "content/docs"
	os.MkdirAll(contentDir, 0755)

	// Create subdirectory with dot-separated numeric prefix
	subDir := filepath.Join(contentDir, "01.getting-started")
	os.MkdirAll(subDir, 0755)

	// Create _dir.yml for custom title
	dirYml := `title: Getting Started Guide`
	os.WriteFile(filepath.Join(subDir, "_dir.yml"), []byte(dirYml), 0644)

	// Create markdown files with dot-separated prefixes
	introMd := `---
title: Introduction to Blue
description: Learn about Blue
---
# Introduction`
	os.WriteFile(filepath.Join(subDir, "01.introduction.md"), []byte(introMd), 0644)

	quickstartMd := `---
title: Quick Start
---
# Quick Start`
	os.WriteFile(filepath.Join(subDir, "02.quickstart.md"), []byte(quickstartMd), 0644)

	// Create file at root level with dot-separated prefix
	overviewMd := `---
title: Documentation Overview
---
# Overview`
	os.WriteFile(filepath.Join(contentDir, "00.overview.md"), []byte(overviewMd), 0644)

	// Test navigation generation
	seoService := NewSEOService()
	navService := &NavigationService{seoService: seoService}

	nav, err := navService.GenerateContentNavigation(contentDir, "/docs")
	if err != nil {
		t.Fatalf("Failed to generate navigation: %v", err)
	}

	// Verify structure
	if len(nav.Sections) != 2 {
		t.Errorf("Expected 2 sections, got %d", len(nav.Sections))
	}

	// Find sections by name for robust testing
	var gettingStartedSection *NavItem
	var overviewSection *NavItem

	for i := range nav.Sections {
		switch nav.Sections[i].Name {
		case "Getting Started Guide":
			gettingStartedSection = &nav.Sections[i]
		case "Documentation Overview":
			overviewSection = &nav.Sections[i]
		}
	}

	if gettingStartedSection == nil {
		t.Fatal("Getting Started Guide section not found")
	}
	if overviewSection == nil {
		t.Fatal("Documentation Overview section not found")
	}

	// Check children in getting-started directory
	if len(gettingStartedSection.Children) != 2 {
		t.Errorf("Expected 2 children in getting-started, got %d", len(gettingStartedSection.Children))
	}

	// Check child ordering
	if gettingStartedSection.Children[0].Name != "Introduction to Blue" {
		t.Errorf("Expected first child to be 'Introduction to Blue', got %q", gettingStartedSection.Children[0].Name)
	}

	// Check href generation - numeric prefixes should be cleaned
	expectedHref := "/docs/getting-started/introduction"
	if gettingStartedSection.Children[0].Href != expectedHref {
		t.Errorf("Expected href %q, got %q", expectedHref, gettingStartedSection.Children[0].Href)
	}

	// Check IDs contain content type prefix
	if !strings.HasPrefix(overviewSection.ID, "docs-") {
		t.Errorf("Expected ID to have 'docs-' prefix, got %q", overviewSection.ID)
	}
}

// TestSortNavItems tests navigation item sorting
func TestSortNavItems(t *testing.T) {
	navService := &NavigationService{}

	items := []NavItem{
		{Name: "Third", OriginalID: "03.third"},
		{Name: "First", OriginalID: "01.first"},
		{Name: "No Prefix", OriginalID: "no-prefix"},
		{Name: "Second", OriginalID: "02.second"},
	}

	navService.sortNavItems(items)

	// Check order - numeric prefixed items should come first in order
	expectedOrder := []string{"First", "Second", "Third", "No Prefix"}
	for i, expected := range expectedOrder {
		if items[i].Name != expected {
			t.Errorf("Expected item %d to be %q, got %q", i, expected, items[i].Name)
		}
	}
}

// TestGetNavigationForPath tests navigation retrieval based on path
func TestGetNavigationForPath(t *testing.T) {
	// Create navigation service with test data
	navService := &NavigationService{
		navigation: &Navigation{
			Sections: []NavItem{
				{ID: "platform", Name: "Platform", Href: "/platform"},
			},
		},
		docsNavigation: &Navigation{
			Sections: []NavItem{
				{ID: "docs-intro", Name: "Introduction", Href: "/docs/intro"},
			},
		},
		apiNavigation: &Navigation{
			Sections: []NavItem{
				{ID: "api-overview", Name: "Overview", Href: "/api/overview"},
			},
		},
		legalNavigation: &Navigation{
			Sections: []NavItem{
				{ID: "legal-terms", Name: "Terms", Href: "/legal/terms"},
			},
		},
	}

	tests := []struct {
		name             string
		path             string
		expectedSections int
		expandedSection  string
	}{
		{
			name:             "Docs path",
			path:             "/docs/intro",
			expectedSections: 4, // platform + docs + api + legal
			expandedSection:  "documentation",
		},
		{
			name:             "API path",
			path:             "/api/overview",
			expectedSections: 4,
			expandedSection:  "api-reference",
		},
		{
			name:             "Legal path",
			path:             "/legal/terms",
			expectedSections: 4,
			expandedSection:  "legal",
		},
		{
			name:             "Other path",
			path:             "/platform",
			expectedSections: 4,
			expandedSection:  "", // No section expanded
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nav := navService.GetNavigationForPath(tt.path)

			if len(nav.Sections) != tt.expectedSections {
				t.Errorf("Expected %d sections, got %d", tt.expectedSections, len(nav.Sections))
			}

			// Check if correct section is expanded
			for _, section := range nav.Sections {
				if section.ID == tt.expandedSection {
					if !section.Expanded {
						t.Errorf("Expected section %s to be expanded", tt.expandedSection)
					}
				} else if section.Expanded {
					t.Errorf("Section %s should not be expanded", section.ID)
				}
			}
		})
	}
}

// TestGetFirstItemInDirectory tests finding first item in directory navigation
func TestGetFirstItemInDirectory(t *testing.T) {
	// Create navigation service with test data
	navService := &NavigationService{
		docsNavigation: &Navigation{
			Sections: []NavItem{
				{
					ID:   "getting-started",
					Name: "Getting Started",
					Children: []NavItem{
						{ID: "intro", Name: "Introduction", Href: "/docs/getting-started/intro"},
						{ID: "setup", Name: "Setup", Href: "/docs/getting-started/setup"},
					},
				},
				{
					ID:   "guides",
					Name: "Guides",
					Href: "/docs/guides",
				},
			},
		},
		apiNavigation: &Navigation{
			Sections: []NavItem{
				{ID: "overview", Name: "Overview", Href: "/api/overview"},
				{ID: "auth", Name: "Authentication", Href: "/api/auth"},
			},
		},
	}

	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "Root docs path",
			path:     "/docs",
			expected: "/docs/getting-started/intro",
		},
		{
			name:     "Docs subdirectory",
			path:     "/docs/getting-started",
			expected: "/docs/getting-started/intro",
		},
		{
			name:     "Root API path",
			path:     "/api",
			expected: "/api/overview",
		},
		{
			name:     "Non-content path",
			path:     "/platform",
			expected: "",
		},
		{
			name:     "Non-existent directory",
			path:     "/docs/non-existent",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := navService.GetFirstItemInDirectory(tt.path)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

// TestProcessDirectory tests processing a content directory
func TestProcessDirectory(t *testing.T) {
	// Create test directory structure
	testDir := t.TempDir()

	// Change to test directory so relative paths work
	oldWd, _ := os.Getwd()
	os.Chdir(testDir)
	defer os.Chdir(oldWd)

	contentDir := "content/docs"
	dirPath := filepath.Join(contentDir, "01.guides")
	os.MkdirAll(dirPath, 0755)

	// Create _dir.yml
	dirYml := `title: User Guides`
	os.WriteFile(filepath.Join(dirPath, "_dir.yml"), []byte(dirYml), 0644)

	// Create subdirectory with dot-separated prefix
	subDir := filepath.Join(dirPath, "02.advanced")
	os.MkdirAll(subDir, 0755)

	// Create files with dot-separated prefixes
	basicMd := `---
title: Basic Guide
---
# Basic`
	os.WriteFile(filepath.Join(dirPath, "01.basic.md"), []byte(basicMd), 0644)

	advancedMd := `---
title: Advanced Topics
---
# Advanced`
	os.WriteFile(filepath.Join(subDir, "01.performance.md"), []byte(advancedMd), 0644)

	// Test processing
	seoService := NewSEOService()
	navService := &NavigationService{seoService: seoService}

	navItem, err := navService.processDirectory(dirPath, "01.guides", "/docs", contentDir)
	if err != nil {
		t.Fatalf("Failed to process directory: %v", err)
	}

	// Verify structure
	if navItem.Name != "User Guides" {
		t.Errorf("Expected name 'User Guides', got %q", navItem.Name)
	}

	if navItem.ID != "docs-guides" {
		t.Errorf("Expected ID 'docs-guides', got %q", navItem.ID)
	}

	// Check children count (1 file + 1 subdirectory)
	if len(navItem.Children) != 2 {
		t.Errorf("Expected 2 children, got %d", len(navItem.Children))
	}

	// Verify file child
	if navItem.Children[0].Name != "Basic Guide" {
		t.Errorf("Expected first child name 'Basic Guide', got %q", navItem.Children[0].Name)
	}

	// Verify subdirectory child
	if navItem.Children[1].Name != "Advanced" {
		t.Errorf("Expected second child name 'Advanced', got %q", navItem.Children[1].Name)
	}

	// Check nested children
	if len(navItem.Children[1].Children) != 1 {
		t.Errorf("Expected 1 nested child, got %d", len(navItem.Children[1].Children))
	}
}

// TestNavItemJSON tests JSON serialization of navigation items
func TestNavItemJSON(t *testing.T) {
	navItem := NavItem{
		ID:       "test",
		Name:     "Test Item",
		Href:     "/test",
		External: true,
		Expanded: true,
		Children: []NavItem{
			{ID: "child", Name: "Child", Href: "/child"},
		},
		OriginalID: "01-test", // This should not appear in JSON
	}

	// Marshal to JSON
	data, err := json.Marshal(navItem)
	if err != nil {
		t.Fatalf("Failed to marshal NavItem: %v", err)
	}

	// Check that JSON is valid
	if !json.Valid(data) {
		t.Error("Invalid JSON generated")
	}

	// Verify OriginalID is not in JSON output
	jsonStr := string(data)
	if contains := strings.Contains(jsonStr, "OriginalID"); contains {
		t.Error("OriginalID should not appear in JSON output")
	}

	// Unmarshal and verify
	var decoded NavItem
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal NavItem: %v", err)
	}

	if decoded.ID != navItem.ID {
		t.Errorf("Expected ID %q, got %q", navItem.ID, decoded.ID)
	}

	if decoded.External != navItem.External {
		t.Errorf("Expected External %v, got %v", navItem.External, decoded.External)
	}

	if len(decoded.Children) != 1 {
		t.Errorf("Expected 1 child, got %d", len(decoded.Children))
	}
}

// TestNavigationServiceIntegration tests the full navigation service
func TestNavigationServiceIntegration(t *testing.T) {
	// Create comprehensive test structure
	testDir := t.TempDir()

	// Create data directory
	dataDir := filepath.Join(testDir, "data")
	os.MkdirAll(dataDir, 0755)

	// Create nav.json
	navData := &Navigation{
		Sections: []NavItem{
			{ID: "home", Name: "Home", Href: "/"},
		},
	}
	navJSON, _ := json.Marshal(navData)
	os.WriteFile(filepath.Join(dataDir, "nav.json"), navJSON, 0644)

	// Create content directories
	for _, contentType := range []string{"docs", "api", "legal"} {
		dir := filepath.Join(testDir, "content", contentType)
		os.MkdirAll(dir, 0755)

		// Add a test file
		testMd := `---
title: ` + contentType + ` Test
---
# Test`
		os.WriteFile(filepath.Join(dir, "test.md"), []byte(testMd), 0644)
	}

	// Change working directory
	oldWd, _ := os.Getwd()
	os.Chdir(testDir)
	defer os.Chdir(oldWd)

	// Create navigation service
	seoService := NewSEOService()
	navService := NewNavigationService(seoService)

	// Verify all navigations are loaded
	if navService.navigation == nil {
		t.Error("Static navigation not loaded")
	}

	if navService.docsNavigation == nil {
		t.Error("Docs navigation not generated")
	}

	if navService.apiNavigation == nil {
		t.Error("API navigation not generated")
	}

	if navService.legalNavigation == nil {
		t.Error("Legal navigation not generated")
	}

	// Test combined navigation
	fullNav := navService.GetNavigationForPath("/docs/test")

	// Should have static + 3 dynamic sections
	if len(fullNav.Sections) < 4 {
		t.Errorf("Expected at least 4 sections, got %d", len(fullNav.Sections))
	}
}
