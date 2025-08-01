package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/joho/godotenv"
	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type PageStats struct {
	Path        string // e.g., "pages/company/about.html" 
	MetadataKey string // e.g., "company/about"
	HasMetadata bool   // whether it exists in metadata.json
	Languages   int    // number of language variants (0-16)
}

type DirectoryStats struct {
	Name              string
	TotalPages        int
	PagesWithMetadata int
	MissingPages      []string
	Coverage          float64
}

type CoverageReport struct {
	TotalHTMLPages       int
	PagesWithMetadata    int
	PagesWithoutMetadata int
	OverallCoverage      float64
	DirectoryStats       map[string]DirectoryStats
	MissingPages         []PageStats
	AllPages             []PageStats
}

type MetadataStructure struct {
	Site  interface{}            `json:"site"`
	Pages map[string]interface{} `json:"pages"`
}

type Config struct {
	AddMode    bool
	TargetPage string
	ShowReport bool
}

type MetadataInput struct {
	PagePath    string
	MetadataKey string
	Title       string
	Description string
}

type PageMetadata struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
}

type TranslationTask struct {
	MetadataKey string
	Title       string
	Description string
	TargetLang  string
}

type TranslationResult struct {
	Task         TranslationTask
	Translation  PageMetadata
	InputTokens  int
	OutputTokens int
	Error        error
}

type TranslationStats struct {
	TotalTasks        int
	CompletedTasks    int32
	FailedTasks       int32
	TotalInputTokens  int64
	TotalOutputTokens int64
	StartTime         time.Time
}

// Language names mapping
var languageNames = map[string]string{
	"en":    "English",
	"es":    "Spanish",
	"fr":    "French",
	"de":    "German",
	"it":    "Italian",
	"pt":    "Portuguese",
	"ja":    "Japanese",
	"ko":    "Korean",
	"zh":    "Chinese (Simplified)",
	"zh-TW": "Chinese (Traditional)",
	"ru":    "Russian",
	"nl":    "Dutch",
	"pl":    "Polish",
	"sv":    "Swedish",
	"id":    "Indonesian",
	"km":    "Khmer",
}

const (
	inputTokenCost  = 0.15 // $0.15 per 1M input tokens for gpt-4o-mini
	outputTokenCost = 0.60 // $0.60 per 1M output tokens for gpt-4o-mini
)

func main() {
	// Parse command line arguments
	config := parseFlags()

	fmt.Println("📊 Blue Metadata Coverage Report")
	fmt.Println("=================================")

	// Scan all HTML pages
	htmlPages, err := scanPagesDirectory("pages")
	if err != nil {
		fmt.Printf("❌ Error scanning pages directory: %v\n", err)
		return
	}

	if len(htmlPages) == 0 {
		fmt.Println("❌ No HTML files found in /pages directory")
		return
	}

	fmt.Printf("\n📝 Found %d HTML pages in /pages directory\n", len(htmlPages))

	// Load metadata
	metadataKeys, err := loadMetadata("data/metadata.json")
	if err != nil {
		fmt.Printf("❌ Error loading metadata: %v\n", err)
		return
	}

	fmt.Printf("📋 Found %d page entries in metadata.json\n\n", len(metadataKeys))

	// Analyze coverage
	report := analyzePageCoverage(htmlPages, metadataKeys)

	if config.AddMode {
		// Interactive metadata addition mode
		err := runInteractiveMode(report, config)
		if err != nil {
			fmt.Printf("❌ Error in interactive mode: %v\n", err)
			return
		}
	} else {
		// Generate standard report
		generateReport(report)
	}
}

func scanPagesDirectory(basePath string) ([]string, error) {
	var htmlFiles []string

	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only include .html files, exclude .bak files
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".html") && !strings.HasSuffix(info.Name(), ".bak") {
			htmlFiles = append(htmlFiles, path)
		}

		return nil
	})

	return htmlFiles, err
}

func loadMetadata(metadataPath string) (map[string]bool, error) {
	file, err := os.ReadFile(metadataPath)
	if err != nil {
		return nil, err
	}

	var metadata MetadataStructure
	err = json.Unmarshal(file, &metadata)
	if err != nil {
		return nil, err
	}

	// Extract all page keys
	metadataKeys := make(map[string]bool)
	for pageKey := range metadata.Pages {
		metadataKeys[pageKey] = true
	}

	return metadataKeys, nil
}

func normalizePagePath(filePath string) string {
	// Remove "pages/" prefix and ".html" suffix
	normalized := strings.TrimPrefix(filePath, "pages/")
	normalized = strings.TrimSuffix(normalized, ".html")

	// Special case: index.html -> home
	if normalized == "index" {
		return "home"
	}

	// Special case: directory/index.html -> directory/index or directory
	// For now, keep as directory/index to match metadata structure
	
	return normalized
}

func analyzePageCoverage(htmlPages []string, metadataKeys map[string]bool) CoverageReport {
	var allPages []PageStats
	var missingPages []PageStats
	pagesWithMetadata := 0

	// Analyze each HTML page
	for _, pagePath := range htmlPages {
		metadataKey := normalizePagePath(pagePath)
		hasMetadata := metadataKeys[metadataKey]

		pageStats := PageStats{
			Path:        pagePath,
			MetadataKey: metadataKey,
			HasMetadata: hasMetadata,
			Languages:   0, // TODO: Count language variants
		}

		allPages = append(allPages, pageStats)

		if hasMetadata {
			pagesWithMetadata++
		} else {
			missingPages = append(missingPages, pageStats)
		}
	}

	// Calculate directory-level statistics
	directoryStats := calculateDirectoryStats(allPages)

	// Calculate overall coverage
	totalPages := len(htmlPages)
	overallCoverage := float64(pagesWithMetadata) / float64(totalPages) * 100

	return CoverageReport{
		TotalHTMLPages:       totalPages,
		PagesWithMetadata:    pagesWithMetadata,
		PagesWithoutMetadata: totalPages - pagesWithMetadata,
		OverallCoverage:      overallCoverage,
		DirectoryStats:       directoryStats,
		MissingPages:         missingPages,
		AllPages:             allPages,
	}
}

func calculateDirectoryStats(allPages []PageStats) map[string]DirectoryStats {
	dirMap := make(map[string]map[string]PageStats)

	// Group pages by directory
	for _, page := range allPages {
		dir := filepath.Dir(page.Path)
		// Remove "pages/" prefix from directory path
		dir = strings.TrimPrefix(dir, "pages")
		if dir == "" || dir == "." {
			dir = "root"
		}

		if dirMap[dir] == nil {
			dirMap[dir] = make(map[string]PageStats)
		}
		dirMap[dir][page.Path] = page
	}

	// Calculate stats for each directory
	directoryStats := make(map[string]DirectoryStats)
	for dir, pages := range dirMap {
		totalPages := len(pages)
		pagesWithMetadata := 0
		var missingPages []string

		for _, page := range pages {
			if page.HasMetadata {
				pagesWithMetadata++
			} else {
				// Use just the filename for missing pages list
				missingPages = append(missingPages, filepath.Base(page.Path))
			}
		}

		coverage := float64(pagesWithMetadata) / float64(totalPages) * 100

		directoryStats[dir] = DirectoryStats{
			Name:              dir,
			TotalPages:        totalPages,
			PagesWithMetadata: pagesWithMetadata,
			MissingPages:      missingPages,
			Coverage:          coverage,
		}
	}

	return directoryStats
}

func generateReport(report CoverageReport) {
	// Generate summary
	generateSummary(report)

	// Generate directory breakdown
	generateDirectoryBreakdown(report)

	// Generate missing pages detail
	generateMissingPagesDetail(report)

	// Generate recommendations
	generateRecommendations(report)
}

func generateSummary(report CoverageReport) {
	fmt.Println("📊 Summary:")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("📄 Total HTML pages: %d\n", report.TotalHTMLPages)
	fmt.Printf("✅ Pages with metadata: %d (%.1f%%)\n", report.PagesWithMetadata, report.OverallCoverage)
	fmt.Printf("❌ Pages without metadata: %d (%.1f%%)\n", report.PagesWithoutMetadata, 100-report.OverallCoverage)
	fmt.Println()
}

func generateDirectoryBreakdown(report CoverageReport) {
	fmt.Println("📊 Coverage by Directory:")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Printf("%-25s %-15s %-15s %-10s\n", "Directory", "Total", "With Metadata", "Coverage")
	fmt.Println(strings.Repeat("-", 70))

	// Sort directories by name for consistent output
	var dirNames []string
	for dirName := range report.DirectoryStats {
		dirNames = append(dirNames, dirName)
	}
	sort.Strings(dirNames)

	for _, dirName := range dirNames {
		dirStats := report.DirectoryStats[dirName]
		
		statusIcon := ""
		if dirStats.Coverage == 100 {
			statusIcon = " ✅"
		} else if dirStats.Coverage == 0 {
			statusIcon = " ❌"
		}

		fmt.Printf("📁 %-23s %-15s %-15s %.1f%%%s\n",
			dirName,
			fmt.Sprintf("%d", dirStats.TotalPages),
			fmt.Sprintf("%d", dirStats.PagesWithMetadata),
			dirStats.Coverage,
			statusIcon)
	}
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println()
}

func generateMissingPagesDetail(report CoverageReport) {
	if len(report.MissingPages) == 0 {
		fmt.Println("🎉 All pages have metadata! ✅")
		return
	}

	fmt.Println("❌ Pages Missing Metadata:")
	fmt.Println(strings.Repeat("=", 50))

	// Group missing pages by directory
	missingByDir := make(map[string][]string)
	for _, page := range report.MissingPages {
		dir := filepath.Dir(page.Path)
		dir = strings.TrimPrefix(dir, "pages")
		if dir == "" || dir == "." {
			dir = "root"
		}

		filename := filepath.Base(page.Path)
		missingByDir[dir] = append(missingByDir[dir], filename)
	}

	// Sort directories for consistent output
	var dirNames []string
	for dirName := range missingByDir {
		dirNames = append(dirNames, dirName)
	}
	sort.Strings(dirNames)

	for _, dirName := range dirNames {
		files := missingByDir[dirName]
		sort.Strings(files) // Sort filenames too

		fmt.Printf("\n📁 %s directory:\n", dirName)
		
		if len(files) <= 5 {
			// Show all files if 5 or fewer
			for _, file := range files {
				fmt.Printf("  • %s\n", file)
			}
		} else {
			// Show first 3 and count remaining
			for i := 0; i < 3; i++ {
				fmt.Printf("  • %s\n", files[i])
			}
			fmt.Printf("  • (and %d more...)\n", len(files)-3)
		}
	}
	fmt.Println()
}

func generateRecommendations(report CoverageReport) {
	fmt.Println("🎯 Recommendations:")
	fmt.Println(strings.Repeat("=", 50))

	// Find directories with poor coverage
	var needsAttention []string
	var perfectDirs []string

	for dirName, dirStats := range report.DirectoryStats {
		if dirStats.Coverage == 100 {
			perfectDirs = append(perfectDirs, dirName)
		} else if dirStats.Coverage < 50 {
			needsAttention = append(needsAttention, fmt.Sprintf("%s (%.0f%%)", dirName, dirStats.Coverage))
		}
	}

	if len(perfectDirs) > 0 {
		sort.Strings(perfectDirs)
		fmt.Printf("✅ Perfect coverage: %s\n", strings.Join(perfectDirs, ", "))
	}

	if len(needsAttention) > 0 {
		fmt.Printf("⚠️  Needs attention: %s\n", strings.Join(needsAttention, ", "))
	}

	// Priority recommendations
	fmt.Println("\n💡 Priority Actions:")
	if report.OverallCoverage < 60 {
		fmt.Println("  1. Focus on high-traffic pages first (home, pricing, enterprise)")
		fmt.Println("  2. Complete directories with partial coverage")
		fmt.Println("  3. Consider if all pages actually need metadata")
	}

	fmt.Printf("\n📈 Overall Status: %.1f%% metadata coverage\n", report.OverallCoverage)
	if report.OverallCoverage >= 80 {
		fmt.Println("🎉 Excellent metadata coverage!")
	} else if report.OverallCoverage >= 60 {
		fmt.Println("👍 Good metadata coverage, room for improvement")
	} else {
		fmt.Println("⚠️  Metadata coverage needs significant improvement")
	}

	if len(report.MissingPages) > 0 {
		fmt.Printf("\n💡 Tip: Use 'go run cmd/metadata-coverage.go -add' to interactively add metadata!\n")
	}
}

// parseFlags parses command line arguments
func parseFlags() Config {
	config := Config{
		ShowReport: true,
	}

	args := os.Args[1:]
	for i, arg := range args {
		switch arg {
		case "-add", "--add":
			config.AddMode = true
			config.ShowReport = false
			// Check if next arg is a page path
			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
				config.TargetPage = args[i+1]
			}
		case "-h", "--help", "-help":
			printUsage()
			os.Exit(0)
		}
	}

	return config
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  go run cmd/metadata-coverage.go                    # Show coverage report")
	fmt.Println("  go run cmd/metadata-coverage.go -add               # Interactive metadata addition")
	fmt.Println("  go run cmd/metadata-coverage.go -add pages/file.html # Add metadata for specific page")
	fmt.Println("")
	fmt.Println("Flags:")
	fmt.Println("  -add, --add    Enable interactive metadata addition mode")
	fmt.Println("  -h, --help     Show this help message")
}

// runInteractiveMode handles the interactive metadata addition workflow
func runInteractiveMode(report CoverageReport, config Config) error {
	fmt.Println("\n🎯 Interactive Metadata Addition Mode")
	fmt.Println("=====================================")

	if len(report.MissingPages) == 0 {
		fmt.Println("🎉 All pages already have metadata! Nothing to add.")
		return nil
	}

	// Load environment variables for OpenAI
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("Warning: .env file not found: %v\n", err)
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("❌ Error: OPENAI_API_KEY not found in environment variables")
		fmt.Println("Please set your OpenAI API key in .env file or environment")
		return fmt.Errorf("OPENAI_API_KEY not found")
	}

	// Initialize OpenAI client
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	// Get target languages
	targetLanguages, err := getTargetLanguages()
	if err != nil {
		return fmt.Errorf("error reading target languages: %v", err)
	}

	var pagesToProcess []PageStats
	if config.TargetPage != "" {
		// Process specific page
		for _, page := range report.MissingPages {
			if page.Path == config.TargetPage || strings.HasSuffix(page.Path, config.TargetPage) {
				pagesToProcess = append(pagesToProcess, page)
				break
			}
		}
		if len(pagesToProcess) == 0 {
			return fmt.Errorf("page not found or already has metadata: %s", config.TargetPage)
		}
	} else {
		// Process all missing pages
		pagesToProcess = report.MissingPages
		fmt.Printf("Found %d pages missing metadata\n\n", len(pagesToProcess))
	}

	// Process each page
	addedCount := 0
	for i, page := range pagesToProcess {
		fmt.Printf("Processing page %d/%d: %s\n", i+1, len(pagesToProcess), page.Path)
		fmt.Println(strings.Repeat("─", 60))

		// Get metadata input from user
		input, shouldContinue, err := promptForMetadata(page)
		if err != nil {
			return err
		}
		if !shouldContinue {
			break
		}

		// Translate metadata
		translations, err := translateMetadata(input, targetLanguages, &client)
		if err != nil {
			fmt.Printf("❌ Translation failed: %v\n", err)
			continue
		}

		// Add to metadata.json
		err = addPageMetadata("data/metadata.json", input.MetadataKey, translations)
		if err != nil {
			fmt.Printf("❌ Failed to save metadata: %v\n", err)
			continue
		}

		fmt.Printf("✅ Added metadata for %s\n\n", page.Path)
		addedCount++
	}

	fmt.Printf("🎉 Interactive mode complete! Added metadata for %d pages.\n", addedCount)
	return nil
}

// promptForMetadata gets metadata input from user
func promptForMetadata(page PageStats) (MetadataInput, bool, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Page: %s\n", filepath.Base(page.Path))
	fmt.Printf("Metadata Key: %s\n", page.MetadataKey)
	fmt.Printf("URL: /%s\n\n", page.MetadataKey)

	// Get title with retry loop
	var title string
	for {
		fmt.Print("📝 Title (e.g., \"Blue Enterprise - Advanced Process Management\"): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return MetadataInput{}, false, err
		}
		title = strings.TrimSpace(input)
		if title != "" {
			break
		}
		fmt.Println("❌ Title cannot be empty. Please try again.")
	}

	// Get description with retry loop
	var description string
	for {
		fmt.Print("📝 Description (1-2 sentences): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return MetadataInput{}, false, err
		}
		description = strings.TrimSpace(input)
		if description != "" {
			break
		}
		fmt.Println("❌ Description cannot be empty. Please try again.")
	}

	// Ask if user wants to continue
	fmt.Print("\nContinue to next page? (y/n/q to quit): ")
	continueInput, err := reader.ReadString('\n')
	if err != nil {
		return MetadataInput{}, false, err
	}
	continueChoice := strings.ToLower(strings.TrimSpace(continueInput))
	
	shouldContinue := true
	if continueChoice == "n" || continueChoice == "q" || continueChoice == "quit" {
		shouldContinue = false
	}

	return MetadataInput{
		PagePath:    page.Path,
		MetadataKey: page.MetadataKey,
		Title:       title,
		Description: description,
	}, shouldContinue, nil
}

// getTargetLanguages reads supported languages from web/languages.go
func getTargetLanguages() ([]string, error) {
	data, err := os.ReadFile("web/languages.go")
	if err != nil {
		return nil, fmt.Errorf("failed to read web/languages.go: %v", err)
	}

	content := string(data)
	start := strings.Index(content, "var SupportedLanguages = []string{")
	if start == -1 {
		return nil, fmt.Errorf("SupportedLanguages not found in web/languages.go")
	}

	end := strings.Index(content[start:], "}")
	if end == -1 {
		return nil, fmt.Errorf("end of SupportedLanguages not found")
	}

	sliceContent := content[start : start+end]
	var languages []string
	lines := strings.Split(sliceContent, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "\"") {
			if endQuote := strings.Index(line[1:], "\""); endQuote != -1 {
				lang := line[1 : endQuote+1]
				if lang != "en" { // Exclude English since it's the source
					languages = append(languages, lang)
				}
			}
		}
	}

	return languages, nil
}

// translateMetadata translates metadata to all target languages
func translateMetadata(input MetadataInput, languages []string, client *openai.Client) (map[string]PageMetadata, error) {
	fmt.Printf("🌐 Translating to %d languages using OpenAI...\n", len(languages))

	// Create translation tasks
	var tasks []TranslationTask
	for _, lang := range languages {
		tasks = append(tasks, TranslationTask{
			MetadataKey: input.MetadataKey,
			Title:       input.Title,
			Description: input.Description,
			TargetLang:  lang,
		})
	}

	// Initialize statistics
	stats := &TranslationStats{
		TotalTasks: len(tasks),
		StartTime:  time.Now(),
	}

	// Process translations
	translations := make(map[string]PageMetadata)
	
	// Add English version
	translations["en"] = PageMetadata{
		Title:       input.Title,
		Description: input.Description,
		Keywords:    generateKeywords(input.Title, input.Description),
	}

	// Concurrent translation
	const workerCount = 10
	taskChan := make(chan TranslationTask, len(tasks))
	resultChan := make(chan TranslationResult, len(tasks))

	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go translationWorker(i, client, taskChan, resultChan, &wg)
	}

	// Send tasks
	for _, task := range tasks {
		taskChan <- task
	}
	close(taskChan)

	// Collect results
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		if result.Error != nil {
			atomic.AddInt32(&stats.FailedTasks, 1)
			fmt.Printf("❌ Translation failed for %s: %v\n", result.Task.TargetLang, result.Error)
			continue
		}

		atomic.AddInt32(&stats.CompletedTasks, 1)
		atomic.AddInt64(&stats.TotalInputTokens, int64(result.InputTokens))
		atomic.AddInt64(&stats.TotalOutputTokens, int64(result.OutputTokens))
		translations[result.Task.TargetLang] = result.Translation

		completed := atomic.LoadInt32(&stats.CompletedTasks)
		failed := atomic.LoadInt32(&stats.FailedTasks)
		percentage := float64(completed+failed) / float64(stats.TotalTasks) * 100
		fmt.Printf("\r[%d/%d] %.1f%% | Completed: %d | Failed: %d", completed+failed, stats.TotalTasks, percentage, completed, failed)
	}

	fmt.Println()

	// Print summary
	completed := atomic.LoadInt32(&stats.CompletedTasks)
	failed := atomic.LoadInt32(&stats.FailedTasks)
	inputTokens := atomic.LoadInt64(&stats.TotalInputTokens)
	outputTokens := atomic.LoadInt64(&stats.TotalOutputTokens)
	totalCost := float64(inputTokens)/1_000_000*inputTokenCost + float64(outputTokens)/1_000_000*outputTokenCost

	fmt.Printf("✅ Translation complete: %d succeeded, %d failed\n", completed, failed)
	fmt.Printf("💰 Estimated cost: $%.4f (%d input, %d output tokens)\n", totalCost, inputTokens, outputTokens)

	return translations, nil
}

// translationWorker processes translation tasks
func translationWorker(id int, client *openai.Client, taskChan <-chan TranslationTask, resultChan chan<- TranslationResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range taskChan {
		result := translateEntry(client, task)
		resultChan <- result
	}
}

// translateEntry translates a single metadata entry
func translateEntry(client *openai.Client, task TranslationTask) TranslationResult {
	ctx := context.Background()

	systemPrompt := `You are a professional translator for Blue, a B2B SaaS process management platform.
You will receive a page title and description to translate.
Return ONLY a valid JSON object with "title" and "description" keys.
Do not include markdown formatting, backticks, or any other text.
Maintain professional tone appropriate for enterprise software.
Keep "Blue" unchanged in all languages.
Preserve technical terms and maintain SEO-friendly length.
Ensure translations are natural and culturally appropriate.`

	userPrompt := fmt.Sprintf(`Translate this page metadata from English to %s:

Title: %s
Description: %s

Return only valid JSON with "title" and "description" keys. No markdown formatting.`,
		languageNames[task.TargetLang], task.Title, task.Description)

	params := openai.ChatCompletionNewParams{
		Model: openai.ChatModelGPT4oMini,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemPrompt),
			openai.UserMessage(userPrompt),
		},
		Temperature: openai.Float(0.3),
		MaxTokens:   openai.Int(500),
	}

	completion, err := client.Chat.Completions.New(ctx, params)
	if err != nil {
		return TranslationResult{
			Task:  task,
			Error: err,
		}
	}

	if len(completion.Choices) == 0 {
		return TranslationResult{
			Task:  task,
			Error: fmt.Errorf("no response from OpenAI"),
		}
	}

	// Parse the JSON response
	var translatedMetadata struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	responseText := completion.Choices[0].Message.Content
	
	// Clean up response text - remove markdown formatting if present
	responseText = strings.TrimSpace(responseText)
	if strings.HasPrefix(responseText, "```json") {
		responseText = strings.TrimPrefix(responseText, "```json")
	}
	if strings.HasPrefix(responseText, "```") {
		responseText = strings.TrimPrefix(responseText, "```")
	}
	if strings.HasSuffix(responseText, "```") {
		responseText = strings.TrimSuffix(responseText, "```")
	}
	responseText = strings.TrimSpace(responseText)
	
	err = json.Unmarshal([]byte(responseText), &translatedMetadata)
	if err != nil {
		return TranslationResult{
			Task:  task,
			Error: fmt.Errorf("failed to parse translation JSON: %v (response: %s)", err, responseText),
		}
	}

	return TranslationResult{
		Task: task,
		Translation: PageMetadata{
			Title:       translatedMetadata.Title,
			Description: translatedMetadata.Description,
			Keywords:    generateKeywords(translatedMetadata.Title, translatedMetadata.Description),
		},
		InputTokens:  int(completion.Usage.PromptTokens),
		OutputTokens: int(completion.Usage.CompletionTokens),
		Error:        nil,
	}
}

// generateKeywords creates SEO keywords from title and description
func generateKeywords(title, description string) []string {
	// Simple keyword extraction - could be enhanced
	keywords := []string{"blue", "process management", "workflow", "automation"}
	
	// Extract key terms from title (basic implementation)
	titleWords := strings.Fields(strings.ToLower(title))
	for _, word := range titleWords {
		word = strings.Trim(word, ".,!?-")
		if len(word) > 3 && word != "blue" {
			keywords = append(keywords, word)
		}
	}
	
	return keywords
}

// addPageMetadata adds metadata to the JSON file
func addPageMetadata(metadataPath string, pageKey string, translations map[string]PageMetadata) error {
	// Create backup
	backupPath := fmt.Sprintf("%s.backup.%d", metadataPath, time.Now().Unix())
	if err := copyFile(metadataPath, backupPath); err != nil {
		fmt.Printf("Warning: Failed to create backup: %v\n", err)
	}

	// Load existing metadata
	file, err := os.ReadFile(metadataPath)
	if err != nil {
		return err
	}

	var metadata MetadataStructure
	err = json.Unmarshal(file, &metadata)
	if err != nil {
		return err
	}

	// Add new page entry
	if metadata.Pages == nil {
		metadata.Pages = make(map[string]interface{})
	}

	pageEntry := make(map[string]interface{})
	for lang, meta := range translations {
		pageEntry[lang] = map[string]interface{}{
			"title":       meta.Title,
			"description": meta.Description,
			"keywords":    meta.Keywords,
		}
	}
	metadata.Pages[pageKey] = pageEntry

	// Save back to file
	data, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return err
	}

	// Atomic write
	tempPath := metadataPath + ".tmp"
	err = os.WriteFile(tempPath, data, 0644)
	if err != nil {
		return err
	}

	return os.Rename(tempPath, metadataPath)
}

// copyFile creates a backup copy of a file
func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0644)
}