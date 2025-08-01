package web

import (
	"os"
	"path/filepath"
	"testing"
)

// TestContentServiceFindMarkdownFile tests finding markdown files
func TestContentServiceFindMarkdownFile(t *testing.T) {
	testDir := t.TempDir()
	contentDir := filepath.Join(testDir, "content")
	
	// Create test directory structure with language subdirectory
	langDir := filepath.Join(contentDir, "en") // Default language
	os.MkdirAll(filepath.Join(langDir, "about"), 0755)
	os.MkdirAll(filepath.Join(langDir, "blog"), 0755)
	
	// Create test files in language directory
	createTestFile(t, filepath.Join(langDir, "about.md"), "# About")
	createTestFile(t, filepath.Join(langDir, "blog", "index.md"), "# Blog")
	createTestFile(t, filepath.Join(langDir, "blog", "post.md"), "# Post")
	
	cs := NewContentService(contentDir)
	
	tests := []struct {
		name        string
		path        string
		wantFile    string
		shouldExist bool
	}{
		{
			name:        "Direct file match",
			path:        "/about",
			wantFile:    filepath.Join(contentDir, "en", "about.md"),
			shouldExist: true,
		},
		{
			name:        "Directory with index",
			path:        "/blog",
			wantFile:    filepath.Join(contentDir, "en", "blog", "index.md"),
			shouldExist: true,
		},
		{
			name:        "Nested file",
			path:        "/blog/post",
			wantFile:    filepath.Join(contentDir, "en", "blog", "post.md"),
			shouldExist: true,
		},
		{
			name:        "Path with trailing slash",
			path:        "/about/",
			wantFile:    filepath.Join(contentDir, "en", "about.md"),
			shouldExist: true,
		},
		{
			name:        "Non-existent file",
			path:        "/nonexistent",
			shouldExist: false,
		},
		{
			name:        "Non-existent nested file",
			path:        "/blog/nonexistent",
			shouldExist: false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := cs.FindMarkdownFile(tt.path)
			
			if tt.shouldExist {
				if err != nil {
					t.Errorf("Expected file to exist but got error: %v", err)
				}
				if file != tt.wantFile {
					t.Errorf("Expected file %q, got %q", tt.wantFile, file)
				}
			} else {
				if err == nil {
					t.Errorf("Expected error for non-existent file, but got file: %s", file)
				}
				if !os.IsNotExist(err) {
					t.Errorf("Expected os.ErrNotExist, got %v", err)
				}
			}
		})
	}
}

// TestContentServiceFindNumberedFiles tests finding files with numeric prefixes
func TestContentServiceFindNumberedFiles(t *testing.T) {
	// ContentService uses BaseDir from ContentTypes directly (not relative to contentDir)
	// So we need to create the exact directory structure expected
	testDir := t.TempDir()
	
	// Change to test directory so relative paths work
	originalWd, _ := os.Getwd()
	os.Chdir(testDir)
	defer os.Chdir(originalWd)
	
	// Create the expected directory structure with language subdirectory
	docsDir := filepath.Join("content", "en", "docs")
	os.MkdirAll(docsDir, 0755)
	
	// Create docs structure with numbered files
	createTestFile(t, filepath.Join(docsDir, "01.introduction.md"), "# Introduction")
	createTestFile(t, filepath.Join(docsDir, "02.getting-started.md"), "# Getting Started")
	createTestFile(t, filepath.Join(docsDir, "10.advanced-features.md"), "# Advanced")
	
	// Create numbered subdirectory
	numberedDir := filepath.Join(docsDir, "03.api")
	os.MkdirAll(numberedDir, 0755)
	createTestFile(t, filepath.Join(numberedDir, "index.md"), "# API")
	createTestFile(t, filepath.Join(numberedDir, "01.authentication.md"), "# Auth")
	
	// Initialize ContentService with "content" as expected
	cs := NewContentService("content")
	
	tests := []struct {
		name        string
		path        string
		wantFile    string
		shouldExist bool
	}{
		{
			name:        "Find numbered file by slug",
			path:        "/docs/introduction",
			wantFile:    filepath.Join("content", "en", "docs", "01.introduction.md"),
			shouldExist: true,
		},
		{
			name:        "Find numbered file with hyphens",
			path:        "/docs/getting-started",
			wantFile:    filepath.Join("content", "en", "docs", "02.getting-started.md"),
			shouldExist: true,
		},
		{
			name:        "Find numbered directory index",
			path:        "/docs/api",
			wantFile:    filepath.Join("content", "en", "docs", "03.api", "index.md"),
			shouldExist: true,
		},
		{
			name:        "Find file in numbered directory",
			path:        "/docs/api/authentication",
			wantFile:    filepath.Join("content", "en", "docs", "03.api", "01.authentication.md"),
			shouldExist: true,
		},
		{
			name:        "Two-digit numbered file",
			path:        "/docs/advanced-features",
			wantFile:    filepath.Join("content", "en", "docs", "10.advanced-features.md"),
			shouldExist: true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := cs.FindMarkdownFile(tt.path)
			
			if tt.shouldExist {
				if err != nil {
					t.Errorf("Expected file to exist but got error: %v", err)
				}
				if file != tt.wantFile {
					t.Errorf("Expected file %q, got %q", tt.wantFile, file)
				}
			} else {
				if err == nil {
					t.Errorf("Expected error for non-existent file, but got file: %s", file)
				}
			}
		})
	}
}

// TestContentServiceCaseInsensitiveSearch tests case-insensitive file matching
func TestContentServiceCaseInsensitiveSearch(t *testing.T) {
	testDir := t.TempDir()
	
	// Change to test directory so relative paths work
	originalWd, _ := os.Getwd()
	os.Chdir(testDir)
	defer os.Chdir(originalWd)
	
	// Create the expected directory structure with language subdirectory
	docsDir := filepath.Join("content", "en", "docs")
	os.MkdirAll(docsDir, 0755)
	
	// Create files with different casing
	createTestFile(t, filepath.Join(docsDir, "01.API-Guide.md"), "# API Guide")
	createTestFile(t, filepath.Join(docsDir, "02.User_Guide.md"), "# User Guide")
	
	cs := NewContentService("content")
	
	tests := []struct {
		name        string
		path        string
		shouldExist bool
	}{
		{
			name:        "Lowercase search for uppercase file",
			path:        "/docs/api-guide",
			shouldExist: true,
		},
		{
			name:        "Search with underscore for hyphenated file",
			path:        "/docs/user-guide",
			shouldExist: true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := cs.FindMarkdownFile(tt.path)
			
			if tt.shouldExist {
				if err != nil {
					t.Errorf("Expected file to exist but got error: %v", err)
				}
				if file == "" {
					t.Error("Expected to find a file but got empty string")
				}
			} else {
				if err == nil {
					t.Errorf("Expected error for non-existent file, but got file: %s", file)
				}
			}
		})
	}
}

// TestFindNumberedDirectory tests finding directories with numeric prefixes
func TestFindNumberedDirectory(t *testing.T) {
	testDir := t.TempDir()
	contentDir := filepath.Join(testDir, "content")
	docsDir := filepath.Join(contentDir, "docs")
	
	// Create numbered directories
	os.MkdirAll(filepath.Join(docsDir, "01.getting-started"), 0755)
	os.MkdirAll(filepath.Join(docsDir, "02.api"), 0755)
	os.MkdirAll(filepath.Join(docsDir, "10.advanced"), 0755)
	
	// Create a regular file to ensure it doesn't match directory searches
	createTestFile(t, filepath.Join(docsDir, "03.not-a-dir.md"), "# File")
	
	cs := NewContentService(contentDir)
	
	tests := []struct {
		name         string
		dir          string
		segment      string
		wantDir      string
		shouldExist  bool
	}{
		{
			name:        "Find simple numbered directory",
			dir:         docsDir,
			segment:     "getting-started",
			wantDir:     filepath.Join(docsDir, "01.getting-started"),
			shouldExist: true,
		},
		{
			name:        "Find directory with short name",
			dir:         docsDir,
			segment:     "api",
			wantDir:     filepath.Join(docsDir, "02.api"),
			shouldExist: true,
		},
		{
			name:        "Two-digit numbered directory",
			dir:         docsDir,
			segment:     "advanced",
			wantDir:     filepath.Join(docsDir, "10.advanced"),
			shouldExist: true,
		},
		{
			name:        "File should not match directory search",
			dir:         docsDir,
			segment:     "not-a-dir",
			shouldExist: false,
		},
		{
			name:        "Non-existent directory",
			dir:         docsDir,
			segment:     "nonexistent",
			shouldExist: false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir, err := cs.findNumberedDirectory(tt.dir, tt.segment)
			
			if tt.shouldExist {
				if err != nil {
					t.Errorf("Expected directory to exist but got error: %v", err)
				}
				if dir != tt.wantDir {
					t.Errorf("Expected directory %q, got %q", tt.wantDir, dir)
				}
			} else {
				if err == nil {
					t.Errorf("Expected error for non-existent directory, but got: %s", dir)
				}
			}
		})
	}
}

// TestGenerateFilePatterns tests the pattern generation for file matching
func TestGenerateFilePatterns(t *testing.T) {
	tests := []struct {
		name      string
		segment   string
		extension string
		want      []string
	}{
		{
			name:      "Simple segment with extension",
			segment:   "introduction",
			extension: ".md",
			want: []string{
				"*introduction.md",
				"*introduction.md",
				"*introduction.md",
			},
		},
		{
			name:      "Hyphenated segment",
			segment:   "getting-started",
			extension: ".md",
			want: []string{
				"*getting-started.md",
				"*getting started.md",
				"*getting_started.md",
			},
		},
		{
			name:      "Directory pattern (no extension)",
			segment:   "api-docs",
			extension: "",
			want: []string{
				"*api-docs",
				"*api docs",
				"*api_docs",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			patterns := GenerateFilePatterns(tt.segment, tt.extension)
			
			if len(patterns) != len(tt.want) {
				t.Errorf("Expected %d patterns, got %d", len(tt.want), len(patterns))
			}
			
			for i, pattern := range patterns {
				if i < len(tt.want) && pattern != tt.want[i] {
					t.Errorf("Pattern %d: expected %q, got %q", i, tt.want[i], pattern)
				}
			}
		})
	}
}

// TestContentServiceEdgeCases tests edge cases and error conditions
func TestContentServiceEdgeCases(t *testing.T) {
	testDir := t.TempDir()
	contentDir := filepath.Join(testDir, "content")
	
	cs := NewContentService(contentDir)
	
	tests := []struct {
		name string
		path string
	}{
		{
			name: "Empty path",
			path: "",
		},
		{
			name: "Root path",
			path: "/",
		},
		{
			name: "Multiple slashes",
			path: "//docs//test//",
		},
		{
			name: "Path with dots",
			path: "../../../etc/passwd",
		},
		{
			name: "Very long path",
			path: "/" + string(make([]byte, 1000)),
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := cs.FindMarkdownFile(tt.path)
			
			// All edge cases should return not found error
			if !os.IsNotExist(err) {
				t.Errorf("Expected os.ErrNotExist for edge case %q, got %v", tt.path, err)
			}
		})
	}
}

