package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type DirectoryStats struct {
	Name        string
	TotalFiles  int
	PresentFiles int
	MissingFiles []string
	Coverage    float64
}

type LanguageStats struct {
	Language      string
	Emoji         string
	TotalFiles    int
	PresentFiles  int
	MissingFiles  []string
	Coverage      float64
	Directories   map[string]DirectoryStats
}

var languageEmojis = map[string]string{
	"de":    "🇩🇪",
	"es":    "🇪🇸", 
	"fr":    "🇫🇷",
	"it":    "🇮🇹",
	"ja":    "🇯🇵",
	"ko":    "🇰🇷",
	"zh":    "🇨🇳",
	"zh-TW": "🇹🇼",
	"ru":    "🇷🇺",
	"pl":    "🇵🇱",
	"pt":    "🇵🇹",
	"nl":    "🇳🇱",
	"sv":    "🇸🇪",
	"id":    "🇮🇩",
	"km":    "🇰🇭",
}

var languageNames = map[string]string{
	"de":    "German",
	"es":    "Spanish",
	"fr":    "French", 
	"it":    "Italian",
	"ja":    "Japanese",
	"ko":    "Korean",
	"zh":    "Chinese",
	"zh-TW": "Chinese (Traditional)",
	"ru":    "Russian",
	"pl":    "Polish",
	"pt":    "Portuguese",
	"nl":    "Dutch",
	"sv":    "Swedish",
	"id":    "Indonesian",
	"km":    "Khmer",
}

func main() {
	fmt.Println("📊 Blue Translation Coverage Report")
	fmt.Println("===================================")
	
	// Scan baseline English content
	baselineFiles, err := scanDirectory("content/en")
	if err != nil {
		fmt.Printf("❌ Error scanning baseline directory: %v\n", err)
		return
	}
	
	if len(baselineFiles) == 0 {
		fmt.Println("❌ No files found in /content/en directory")
		return
	}
	
	fmt.Printf("\n📝 Baseline: English (/content/en) - %d files\n\n", len(baselineFiles))
	
	// Get all language directories
	languages, err := getLanguageDirectories()
	if err != nil {
		fmt.Printf("❌ Error getting language directories: %v\n", err)
		return
	}
	
	// Analyze each language
	var allStats []LanguageStats
	for _, lang := range languages {
		stats := analyzeLanguage(lang, baselineFiles)
		allStats = append(allStats, stats)
	}
	
	// Sort by coverage (highest first)
	sort.Slice(allStats, func(i, j int) bool {
		return allStats[i].Coverage > allStats[j].Coverage
	})
	
	// Generate report
	generateReport(allStats, len(baselineFiles))
}

func scanDirectory(basePath string) (map[string]bool, error) {
	files := make(map[string]bool)
	
	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Only include .md files
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			// Get relative path from base directory
			relPath, err := filepath.Rel(basePath, path)
			if err != nil {
				return err
			}
			files[relPath] = true
		}
		
		return nil
	})
	
	return files, err
}

func getLanguageDirectories() ([]string, error) {
	contentDir := "content"
	entries, err := os.ReadDir(contentDir)
	if err != nil {
		return nil, err
	}
	
	var languages []string
	for _, entry := range entries {
		if entry.IsDir() && entry.Name() != "en" {
			languages = append(languages, entry.Name())
		}
	}
	
	sort.Strings(languages)
	return languages, nil
}

func analyzeLanguage(language string, baselineFiles map[string]bool) LanguageStats {
	langPath := filepath.Join("content", language)
	langFiles, err := scanDirectory(langPath)
	if err != nil {
		// Language directory doesn't exist or is empty
		langFiles = make(map[string]bool)
	}
	
	// Calculate overall stats
	presentCount := 0
	var missingFiles []string
	
	for file := range baselineFiles {
		if langFiles[file] {
			presentCount++
		} else {
			missingFiles = append(missingFiles, file)
		}
	}
	
	coverage := float64(presentCount) / float64(len(baselineFiles)) * 100
	
	// Analyze by directory
	directories := analyzeDirectories(baselineFiles, langFiles)
	
	emoji := languageEmojis[language]
	if emoji == "" {
		emoji = "🌐"
	}
	
	return LanguageStats{
		Language:      language,
		Emoji:         emoji,
		TotalFiles:    len(baselineFiles),
		PresentFiles:  presentCount,
		MissingFiles:  missingFiles,
		Coverage:      coverage,
		Directories:   directories,
	}
}

func analyzeDirectories(baselineFiles, langFiles map[string]bool) map[string]DirectoryStats {
	dirMap := make(map[string]map[string]bool)
	
	// Group files by directory
	for file := range baselineFiles {
		dir := filepath.Dir(file)
		if dir == "." {
			dir = "root"
		}
		
		if dirMap[dir] == nil {
			dirMap[dir] = make(map[string]bool)
		}
		dirMap[dir][file] = true
	}
	
	// Calculate stats for each directory
	directories := make(map[string]DirectoryStats)
	for dir, files := range dirMap {
		presentCount := 0
		var missingFiles []string
		
		for file := range files {
			if langFiles[file] {
				presentCount++
			} else {
				missingFiles = append(missingFiles, filepath.Base(file))
			}
		}
		
		coverage := float64(presentCount) / float64(len(files)) * 100
		
		directories[dir] = DirectoryStats{
			Name:         dir,
			TotalFiles:   len(files),
			PresentFiles: presentCount,
			MissingFiles: missingFiles,
			Coverage:     coverage,
		}
	}
	
	return directories
}

func generateReport(allStats []LanguageStats, totalFiles int) {
	// Generate summary table first
	generateSummaryTable(allStats, totalFiles)
	
	// Then detailed breakdown
	fmt.Println("\n📋 Detailed Breakdown:")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println()
	
	for _, stats := range allStats {
		generateLanguageReport(stats)
	}
	
	// Generate final summary
	fmt.Println("📈 Summary:")
	if len(allStats) > 0 {
		fmt.Printf("  🥇 Best: %s (%.1f%%)\n", getLanguageDisplayName(allStats[0].Language), allStats[0].Coverage)
	}
	if len(allStats) > 1 {
		fmt.Printf("  🥈 Second: %s (%.1f%%)\n", getLanguageDisplayName(allStats[1].Language), allStats[1].Coverage)
	}
	if len(allStats) > 2 {
		fmt.Printf("  🥉 Third: %s (%.1f%%)\n", getLanguageDisplayName(allStats[2].Language), allStats[2].Coverage)
	}
	
	// Find languages that need attention
	var needsAttention []string
	for _, stats := range allStats {
		if stats.Coverage < 10 {
			needsAttention = append(needsAttention, getLanguageDisplayName(stats.Language))
		}
	}
	
	if len(needsAttention) > 0 {
		fmt.Printf("  ⚠️  Needs attention: %s (< 10%%)\n", strings.Join(needsAttention, ", "))
	}
	
	fmt.Printf("\n💾 Total baseline files: %d\n", totalFiles)
	fmt.Printf("🌐 Languages analyzed: %d\n", len(allStats))
}

func generateLanguageReport(stats LanguageStats) {
	langName := getLanguageDisplayName(stats.Language)
	
	fmt.Printf("%s %s (%s) - %.1f%% coverage (%d/%d files)\n", 
		stats.Emoji, langName, stats.Language, stats.Coverage, 
		stats.PresentFiles, stats.TotalFiles)
	
	// Sort directories by name for consistent output
	var dirNames []string
	for dirName := range stats.Directories {
		dirNames = append(dirNames, dirName)
	}
	sort.Strings(dirNames)
	
	// Show directory breakdown
	for _, dirName := range dirNames {
		dirStats := stats.Directories[dirName]
		
		if dirStats.Coverage == 100 {
			fmt.Printf("  📁 %s/ - %.0f%% (%d/%d) ✅\n", 
				dirName, dirStats.Coverage, dirStats.PresentFiles, dirStats.TotalFiles)
		} else if dirStats.Coverage == 0 {
			fmt.Printf("  📁 %s/ - %.0f%% (%d/%d) ❌ Missing entire directory\n", 
				dirName, dirStats.Coverage, dirStats.PresentFiles, dirStats.TotalFiles)
		} else {
			fmt.Printf("  📁 %s/ - %.0f%% (%d/%d) ❌ Missing: %s\n", 
				dirName, dirStats.Coverage, dirStats.PresentFiles, dirStats.TotalFiles,
				formatMissingFiles(dirStats.MissingFiles))
		}
	}
	
	fmt.Println()
}

func formatMissingFiles(missingFiles []string) string {
	if len(missingFiles) == 0 {
		return "none"
	}
	
	// Sort for consistent output
	sort.Strings(missingFiles)
	
	if len(missingFiles) <= 3 {
		return strings.Join(missingFiles, ", ")
	}
	
	// Show first 3 and count
	return fmt.Sprintf("%s, and %d more", 
		strings.Join(missingFiles[:3], ", "), len(missingFiles)-3)
}

func generateSummaryTable(allStats []LanguageStats, totalFiles int) {
	fmt.Println("📊 Coverage Summary:")
	fmt.Println(strings.Repeat("=", 65))
	fmt.Printf("%-4s %-20s %-15s %-10s\n", "", "Language", "Files", "Coverage")
	fmt.Println(strings.Repeat("-", 65))
	
	for i, stats := range allStats {
		rankEmoji := ""
		if i == 0 {
			rankEmoji = "🥇"
		} else if i == 1 {
			rankEmoji = "🥈"
		} else if i == 2 {
			rankEmoji = "🥉"
		}
		
		langName := getLanguageDisplayName(stats.Language)
		filesRatio := fmt.Sprintf("%d/%d", stats.PresentFiles, stats.TotalFiles)
		coverage := fmt.Sprintf("%.1f%%", stats.Coverage)
		
		fmt.Printf("%-4s %s %-18s %-15s %s\n", 
			rankEmoji, stats.Emoji, langName, filesRatio, coverage)
	}
	
	fmt.Println(strings.Repeat("=", 65))
}

func getLanguageDisplayName(langCode string) string {
	if name, exists := languageNames[langCode]; exists {
		return name
	}
	return strings.ToUpper(langCode)
}