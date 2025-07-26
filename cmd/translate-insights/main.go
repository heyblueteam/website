package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)


// TranslationTask represents a single translation job
type TranslationTask struct {
	SourceFile   string
	TargetFile   string
	SourceLang   string
	TargetLang   string
	SourcePath   string
	TargetPath   string
}

// TranslationStats tracks statistics for the translation session
type TranslationStats struct {
	TotalTasks        int
	CompletedTasks    int32  // Use atomic for thread-safe updates
	TotalInputTokens  int64  // Use atomic for thread-safe updates
	TotalOutputTokens int64  // Use atomic for thread-safe updates
	StartTime         time.Time
	mu                sync.Mutex
}

// TranslationResult represents the result of a translation task
type TranslationResult struct {
	Task         TranslationTask
	InputTokens  int
	OutputTokens int
	Error        error
	Duration     time.Duration
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
	// OpenAI pricing per 1M tokens
	inputTokenCost  = 0.4    // $0.4 per 1M input tokens
	outputTokenCost = 1.60   // $1.60 per 1M output tokens
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Warning: .env file not found: %v\n", err)
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: OPENAI_API_KEY not found in environment variables")
		os.Exit(1)
	}

	// Initialize OpenAI client
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	// Get base content directory
	contentDir := filepath.Join(".", "content")

	// Discover available languages
	languages, err := discoverLanguages(contentDir)
	if err != nil {
		fmt.Printf("Error discovering languages: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d languages: %v\n", len(languages), languages)

	// Get English insights
	englishInsights, err := getInsights(filepath.Join(contentDir, "en", "insights"))
	if err != nil {
		fmt.Printf("Error reading English insights: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d English insights\n", len(englishInsights))

	// Build translation queue
	tasks := buildTranslationQueue(contentDir, languages, englishInsights)

	if len(tasks) == 0 {
		fmt.Println("No translations needed. All insights are up to date!")
		return
	}

	// Display translation summary
	fmt.Printf("\n=== Translation Summary ===\n")
	tasksByLang := make(map[string]int)
	for _, task := range tasks {
		tasksByLang[task.TargetLang]++
	}

	for lang, count := range tasksByLang {
		fmt.Printf("%s (%s): %d translations needed\n", languageNames[lang], lang, count)
	}
	fmt.Printf("\nTotal translations required: %d\n", len(tasks))

	// Initialize statistics
	stats := &TranslationStats{
		TotalTasks: len(tasks),
		StartTime:  time.Now(),
	}

	// Number of concurrent workers
	const workerCount = 100

	// Start translation process
	fmt.Printf("\n=== Starting Translation Process ===\n")
	fmt.Printf("Running %d parallel workers...\n", workerCount)
	
	systemPrompt := getSystemPrompt()

	// Create channels
	taskChan := make(chan TranslationTask, len(tasks))
	resultChan := make(chan TranslationResult, len(tasks))

	// Create wait group for workers
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go translationWorker(i, &client, systemPrompt, taskChan, resultChan, &wg)
	}

	// Send all tasks to the channel
	for _, task := range tasks {
		taskChan <- task
	}
	close(taskChan)

	// Start result collector in a separate goroutine
	doneChan := make(chan bool)
	go collectResults(stats, resultChan, len(tasks), doneChan)

	// Wait for all workers to complete
	wg.Wait()
	close(resultChan)

	// Wait for result collector to finish
	<-doneChan

	// Final summary
	totalTime := time.Since(stats.StartTime)
	completed := atomic.LoadInt32(&stats.CompletedTasks)
	inputTokens := atomic.LoadInt64(&stats.TotalInputTokens)
	outputTokens := atomic.LoadInt64(&stats.TotalOutputTokens)
	totalCost := calculateCost(int(inputTokens), int(outputTokens))
	avgTime := totalTime.Seconds() / float64(completed)

	fmt.Printf("\n=== Translation Complete ===\n")
	fmt.Printf("Total tasks: %d\n", stats.TotalTasks)
	fmt.Printf("Completed: %d\n", completed)
	fmt.Printf("Total time: %s\n", formatDuration(totalTime))
	fmt.Printf("Average time: %.1fs/translation\n", avgTime)
	fmt.Printf("Total tokens: %d input, %d output\n", inputTokens, outputTokens)
	fmt.Printf("Total cost: $%.4f\n", totalCost)
}

// translationWorker processes translation tasks from the channel
func translationWorker(id int, client *openai.Client, systemPrompt string, taskChan <-chan TranslationTask, resultChan chan<- TranslationResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range taskChan {
		startTime := time.Now()
		inputTokens, outputTokens, err := translateInsight(client, task, systemPrompt)
		duration := time.Since(startTime)

		result := TranslationResult{
			Task:         task,
			InputTokens:  inputTokens,
			OutputTokens: outputTokens,
			Error:        err,
			Duration:     duration,
		}

		resultChan <- result
	}
}

// collectResults collects results from workers and updates progress
func collectResults(stats *TranslationStats, resultChan <-chan TranslationResult, totalTasks int, doneChan chan<- bool) {
	for result := range resultChan {
		completed := atomic.AddInt32(&stats.CompletedTasks, 1)
		
		// Update token counts atomically
		if result.Error == nil {
			atomic.AddInt64(&stats.TotalInputTokens, int64(result.InputTokens))
			atomic.AddInt64(&stats.TotalOutputTokens, int64(result.OutputTokens))
		}

		// Calculate progress
		progress := float64(completed) / float64(totalTasks) * 100
		elapsed := time.Since(stats.StartTime)
		avgTime := elapsed.Seconds() / float64(completed)
		
		// Get current totals
		inputTokens := atomic.LoadInt64(&stats.TotalInputTokens)
		outputTokens := atomic.LoadInt64(&stats.TotalOutputTokens)
		currentCost := calculateCost(int(inputTokens), int(outputTokens))

		// Display progress
		fmt.Printf("\r[%d/%d] %.1f%% | Avg: %.1fs | Elapsed: %s | Tokens: %d in, %d out | Cost: $%.4f", 
			completed, totalTasks, progress, avgTime, formatDuration(elapsed), 
			inputTokens, outputTokens, currentCost)

		// Show errors on new line
		if result.Error != nil {
			fmt.Printf("\n  ERROR [%s -> %s]: %v\n", 
				result.Task.SourceFile, result.Task.TargetLang, result.Error)
		}
	}
	
	// Clear the progress line and move to next line
	fmt.Println()
	doneChan <- true
}

func discoverLanguages(contentDir string) ([]string, error) {
	entries, err := os.ReadDir(contentDir)
	if err != nil {
		return nil, err
	}

	var languages []string
	for _, entry := range entries {
		if entry.IsDir() && entry.Name() != "en" {
			// Check if it has an insights directory
			insightsPath := filepath.Join(contentDir, entry.Name(), "insights")
			if _, err := os.Stat(insightsPath); err == nil {
				languages = append(languages, entry.Name())
			}
		}
	}

	return languages, nil
}

func getInsights(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var insights []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".md") {
			insights = append(insights, entry.Name())
		}
	}

	return insights, nil
}

func buildTranslationQueue(contentDir string, languages []string, englishInsights []string) []TranslationTask {
	var tasks []TranslationTask

	for _, lang := range languages {
		targetInsightsDir := filepath.Join(contentDir, lang, "insights")
		
		// Get existing translations
		existingInsights, _ := getInsights(targetInsightsDir)
		existingMap := make(map[string]bool)
		for _, insight := range existingInsights {
			existingMap[insight] = true
		}

		// Find missing translations
		for _, insight := range englishInsights {
			if !existingMap[insight] {
				task := TranslationTask{
					SourceFile: insight,
					TargetFile: insight,
					SourceLang: "en",
					TargetLang: lang,
					SourcePath: filepath.Join(contentDir, "en", "insights", insight),
					TargetPath: filepath.Join(contentDir, lang, "insights", insight),
				}
				tasks = append(tasks, task)
			}
		}
	}

	return tasks
}

func getSystemPrompt() string {
	return `You are a professional translator for the Blue website, a B2B SaaS process management platform. You must translate the ENTIRE markdown file including all body content.

WHAT TO TRANSLATE:
1. The "title" field value in the frontmatter
2. The "description" field value in the frontmatter  
3. ALL BODY CONTENT after the closing --- of the frontmatter
4. All headings, paragraphs, lists, and text content

NEVER TRANSLATE:
- "Blue" (product name)
- Email addresses (support@blue.cc, sales@blue.cc, etc.)
- URLs and links
- Image filenames
- Code blocks
- Technical terms (API, GraphQL, webhook, OAuth, etc.)
- Person names
- Field names like "title:" and "description:" (only translate their values)

IMPORTANT: You MUST translate the ENTIRE document body, not just the frontmatter. Every paragraph, heading, and text element must be translated to the target language.

FORMAT:
- Return the complete translated markdown file
- Preserve all markdown formatting (headers, lists, links, etc.)
- Keep the same structure as the original
- Do not add any explanations or comments`
}

func translateInsight(client *openai.Client, task TranslationTask, systemPrompt string) (int, int, error) {
	// Read source file
	content, err := os.ReadFile(task.SourcePath)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to read source file: %w", err)
	}

	// Strip out category and date before sending to OpenAI
	contentForTranslation, originalLines := stripPreservableFields(string(content))

	// Prepare user prompt with stripped content
	userPrompt := fmt.Sprintf("Translate the following ENTIRE markdown document from English to %s. Translate ALL content including the body text, not just the frontmatter:\n\n%s", 
		languageNames[task.TargetLang], contentForTranslation)

	// Create chat completion request
	ctx := context.Background()
	params := openai.ChatCompletionNewParams{
		Model: openai.ChatModelGPT4oMini,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemPrompt),
			openai.UserMessage(userPrompt),
		},
		Temperature: openai.Float(0.3), // Lower temperature for more consistent translations
	}

	completion, err := client.Chat.Completions.New(ctx, params)
	if err != nil {
		return 0, 0, fmt.Errorf("OpenAI API error: %w", err)
	}

	if len(completion.Choices) == 0 {
		return 0, 0, fmt.Errorf("no response from OpenAI")
	}

	translatedContent := completion.Choices[0].Message.Content

	// Restore the original category and date fields
	translatedContent = restorePreservableFields(translatedContent, originalLines)

	// Ensure target directory exists
	targetDir := filepath.Dir(task.TargetPath)
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return 0, 0, fmt.Errorf("failed to create target directory: %w", err)
	}

	// Write translated file
	if err := os.WriteFile(task.TargetPath, []byte(translatedContent), 0644); err != nil {
		return 0, 0, fmt.Errorf("failed to write translated file: %w", err)
	}

	// Get usage information
	inputTokens := int(completion.Usage.PromptTokens)
	outputTokens := int(completion.Usage.CompletionTokens)

	return inputTokens, outputTokens, nil
}

// stripPreservableFields removes category and date from content before translation
func stripPreservableFields(content string) (strippedContent string, originalLines []string) {
	parts := strings.SplitN(content, "---", 3)
	if len(parts) < 3 {
		return content, nil
	}

	frontmatterStr := parts[1]
	body := parts[2]

	// Store original lines that contain category and date
	originalLines = []string{}
	
	// Process frontmatter line by line to preserve exact formatting
	var newFrontmatterLines []string
	lines := strings.Split(strings.TrimSpace(frontmatterStr), "\n")
	
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		// Check if this line contains category or date
		if strings.HasPrefix(trimmed, "category:") || strings.HasPrefix(trimmed, "date:") {
			// Store the exact original line
			originalLines = append(originalLines, line)
		} else if trimmed != "" {
			// Keep other lines
			newFrontmatterLines = append(newFrontmatterLines, line)
		}
	}

	// Reconstruct content without category and date
	if len(newFrontmatterLines) > 0 {
		strippedContent = fmt.Sprintf("---\n%s\n---\n%s", strings.Join(newFrontmatterLines, "\n"), body)
	} else {
		strippedContent = fmt.Sprintf("---\n---\n%s", body)
	}
	
	return strippedContent, originalLines
}

// restorePreservableFields adds back the original category and date fields
func restorePreservableFields(translatedContent string, originalLines []string) string {
	parts := strings.SplitN(translatedContent, "---", 3)
	if len(parts) < 3 || originalLines == nil {
		return translatedContent
	}

	frontmatterStr := parts[1]
	body := parts[2]

	// Process translated frontmatter line by line
	lines := strings.Split(strings.TrimSpace(frontmatterStr), "\n")
	
	// Fix any accidentally translated field names
	translations := map[string]string{
		"título": "title", "titre": "title", "titel": "title", "titolo": "title", 
		"tytuł": "title", "заголовок": "title", "タイトル": "title", "제목": "title", 
		"标题": "title", "標題": "title",
		"descripción": "description", "beschreibung": "description", "descrizione": "description", 
		"descrição": "description", "opis": "description", "описание": "description", 
		"説明": "description", "설명": "description", "描述": "description",
	}
	
	var fixedLines []string
	for _, line := range lines {
		fixedLine := line
		for wrong, correct := range translations {
			if strings.Contains(line, wrong+":") {
				fixedLine = strings.Replace(line, wrong+":", correct+":", 1)
				break
			}
		}
		if strings.TrimSpace(fixedLine) != "" {
			fixedLines = append(fixedLines, fixedLine)
		}
	}

	// Add back the original category and date lines
	fixedLines = append(fixedLines, originalLines...)

	// Reconstruct with restored fields
	return fmt.Sprintf("---\n%s\n---\n%s", strings.Join(fixedLines, "\n"), body)
}

func calculateCost(inputTokens, outputTokens int) float64 {
	inputCost := float64(inputTokens) / 1_000_000 * inputTokenCost
	outputCost := float64(outputTokens) / 1_000_000 * outputTokenCost
	return inputCost + outputCost
}

func formatDuration(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60

	if h > 0 {
		return fmt.Sprintf("%dh %dm %ds", h, m, s)
	} else if m > 0 {
		return fmt.Sprintf("%dm %ds", m, s)
	}
	return fmt.Sprintf("%ds", s)
}