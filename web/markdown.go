package web

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// MarkdownService handles markdown processing
type MarkdownService struct {
	markdown goldmark.Markdown
	cache    *MarkdownCache
}

// NewMarkdownService creates a new markdown service
func NewMarkdownService() *MarkdownService {
	// Configure Goldmark with extensions
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			NewYouTubeExtension(),
			NewCalloutExtension(),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
			html.WithUnsafe(), // Allow raw HTML (including video tags)
		),
	)

	return &MarkdownService{
		markdown: md,
		cache:    NewMarkdownCache(),
	}
}

// Convert converts markdown content to HTML
func (ms *MarkdownService) Convert(markdownContent []byte) (string, error) {
	var htmlBuffer strings.Builder
	if err := ms.markdown.Convert(markdownContent, &htmlBuffer); err != nil {
		return "", err
	}
	return htmlBuffer.String(), nil
}

// ProcessMarkdownFile reads a markdown file, parses frontmatter, and converts to HTML
func (ms *MarkdownService) ProcessMarkdownFile(filePath string, seoService *SEOService) (string, *Frontmatter, error) {
	// Read file
	mdBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", nil, err
	}

	// Parse frontmatter
	var markdownContent []byte
	frontmatter, markdownContent, err := seoService.ParseFrontmatter(mdBytes)
	if err != nil {
		// Continue without frontmatter
		markdownContent = mdBytes
	}

	// Convert to HTML
	html, err := ms.Convert(markdownContent)
	if err != nil {
		return "", nil, err
	}

	return html, frontmatter, nil
}

// fileTask represents a file to be processed
type fileTask struct {
	path     string
	lang     string
	modTime  time.Time
}

// PreRenderAllMarkdown pre-renders all markdown files in the content directory
func (ms *MarkdownService) PreRenderAllMarkdown(contentService *ContentService, seoService *SEOService) error {
	// Collect all markdown files first
	var fileTasks []fileTask
	
	// Walk through all markdown files in content directory for each language
	for _, lang := range SupportedLanguages {
		contentDir := filepath.Join("content", lang)

		// Skip if language directory doesn't exist
		if _, err := os.Stat(contentDir); os.IsNotExist(err) {
			log.Printf("Skipping language %s: directory not found", lang)
			continue
		}

		err := filepath.WalkDir(contentDir, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}

			// Skip non-markdown files
			if !strings.HasSuffix(path, ".md") {
				return nil
			}

			// Get file info for modification time
			info, err := d.Info()
			if err != nil {
				log.Printf("Warning: could not get file info for %s: %v", path, err)
				return nil // Continue processing other files
			}

			fileTasks = append(fileTasks, fileTask{
				path:    path,
				lang:    lang,
				modTime: info.ModTime(),
			})

			return nil
		})

		if err != nil {
			return fmt.Errorf("failed to walk content directory for %s: %w", lang, err)
		}
	}

	// Process files in parallel using worker pool
	const numWorkers = 20
	taskChan := make(chan fileTask, len(fileTasks))
	resultChan := make(chan int, len(fileTasks))
	errorChan := make(chan error, numWorkers)

	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range taskChan {
				// Generate URL path for this file (removes language from path)
				urlPath := ms.generateURLPath(task.path)

				// Process the markdown file
				html, frontmatter, err := ms.ProcessMarkdownFile(task.path, seoService)
				if err != nil {
					log.Printf("Warning: failed to process %s: %v", task.path, err)
					continue // Continue processing other files
				}

				// Detect if the HTML contains code blocks
				needsCodeHighlight := DetectCodeBlocks(html)

				// Cache the pre-rendered content with language-specific key
				cachedContent := &CachedContent{
					HTML:               html,
					Frontmatter:        frontmatter,
					ModTime:            task.modTime,
					FilePath:           task.path,
					NeedsCodeHighlight: needsCodeHighlight,
				}

				// Use language-specific cache key
				cacheKey := task.lang + ":" + urlPath
				ms.cache.Set(cacheKey, cachedContent)
				resultChan <- 1
			}
		}()
	}

	// Send all tasks to workers
	for _, task := range fileTasks {
		taskChan <- task
	}
	close(taskChan)

	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	// Count processed files
	count := 0
	for range resultChan {
		count++
	}

	// Check for any errors
	select {
	case err := <-errorChan:
		if err != nil {
			return err
		}
	default:
	}

	// Pre-rendering complete
	return nil
}

// generateURLPath converts a file path to a clean URL path
func (ms *MarkdownService) generateURLPath(filePath string) string {
	// Remove content/ prefix and .md suffix
	urlPath := strings.TrimPrefix(filePath, "content/")
	urlPath = strings.TrimSuffix(urlPath, ".md")

	// Remove language prefix (e.g., "en/" or "es/")
	for _, lang := range SupportedLanguages {
		if strings.HasPrefix(urlPath, lang+"/") {
			urlPath = strings.TrimPrefix(urlPath, lang+"/")
			break
		}
	}

	// Handle index files
	urlPath = strings.TrimSuffix(urlPath, "/index")

	// Clean URL parts: remove number prefixes and normalize
	urlParts := strings.Split(urlPath, "/")
	for i, part := range urlParts {
		if part != "" {
			// Use the same cleaning logic as the utils
			urlParts[i] = CleanID(part)
		}
	}

	// Reconstruct URL with leading slash
	cleanURL := "/" + strings.Join(urlParts, "/")

	// Handle root case
	if cleanURL == "/" {
		return "/"
	}

	return cleanURL
}

// GetCachedContent retrieves pre-rendered content from cache
func (ms *MarkdownService) GetCachedContent(urlPath string) (*CachedContent, bool) {
	return ms.cache.Get(urlPath)
}

// GetCachedContentForLang retrieves pre-rendered content from cache for a specific language
func (ms *MarkdownService) GetCachedContentForLang(urlPath, lang string) (*CachedContent, bool) {
	// Try language-specific content first
	cacheKey := lang + ":" + urlPath
	if content, found := ms.cache.Get(cacheKey); found {
		return content, true
	}

	// Fall back to English if specific language not found
	if lang != DefaultLanguage {
		englishKey := DefaultLanguage + ":" + urlPath
		return ms.cache.Get(englishKey)
	}

	return nil, false
}

// GetAllCachedContent returns all cached content (for search indexing)
func (ms *MarkdownService) GetAllCachedContent() map[string]*CachedContent {
	return ms.cache.GetAll()
}

// GetCachedContentByLanguage returns all cached content for a specific language
func (ms *MarkdownService) GetCachedContentByLanguage(lang string) map[string]*CachedContent {
	return ms.cache.GetByLanguage(lang)
}

// GetCacheSize returns the number of cached items
func (ms *MarkdownService) GetCacheSize() int {
	return ms.cache.Size()
}
