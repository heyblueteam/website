package web

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// HTMLService handles HTML page pre-rendering
type HTMLService struct {
	cache           *HTMLCache
	pagesDir        string
	layoutsDir      string
	componentsDir   string
	markdownService *MarkdownService
}

// NewHTMLService creates a new HTML service
func NewHTMLService(pagesDir, layoutsDir, componentsDir string, markdownService *MarkdownService) *HTMLService {
	return &HTMLService{
		cache:           NewHTMLCache(),
		pagesDir:        pagesDir,
		layoutsDir:      layoutsDir,
		componentsDir:   componentsDir,
		markdownService: markdownService,
	}
}

// PreRenderAllHTMLPages pre-renders all HTML pages in the pages directory
func (hs *HTMLService) PreRenderAllHTMLPages(navigationService *NavigationService, seoService *SEOService) error {
	count := 0

	// List of pages to exclude from pre-rendering (dynamic content)
	excludedPages := []string{
		"/platform/status", // Dynamic status page (truly dynamic - status changes)
		// Note: /insights is now pre-rendered with insights data baked in
	}

	// Walk through all HTML files in pages directory
	err := filepath.WalkDir(hs.pagesDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip non-HTML files
		if !strings.HasSuffix(path, ".html") {
			return nil
		}

		// Generate URL path for this file
		urlPath := hs.generateURLPath(path)

		// Check if this page should be excluded
		if hs.isExcluded(urlPath, excludedPages) {
			// Skipping excluded page
			return nil
		}

		// Get file info for modification time
		info, err := d.Info()
		if err != nil {
			log.Printf("Warning: could not get file info for %s: %v", path, err)
			return nil // Continue processing other files
		}

		// Pre-render the HTML page
		html, err := hs.renderHTMLPage(path, urlPath, navigationService, seoService)
		if err != nil {
			log.Printf("Warning: failed to pre-render %s: %v", path, err)
			return nil // Continue processing other files
		}

		// Cache the pre-rendered content
		cachedContent := &CachedContent{
			HTML:        html,
			Frontmatter: nil, // HTML pages don't have frontmatter
			ModTime:     info.ModTime(),
			FilePath:    path,
		}

		hs.cache.Set(urlPath, cachedContent)
		count++

		// Progress tracking removed for cleaner output

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to walk pages directory: %w", err)
	}

	// Pre-rendering complete
	return nil
}

// renderHTMLPage renders a single HTML page with templates
func (hs *HTMLService) renderHTMLPage(filePath, urlPath string, navigationService *NavigationService, seoService *SEOService) (string, error) {
	// Read the HTML file
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	// Prepare page data
	pageData := hs.preparePageData(urlPath, "", false, nil, navigationService, seoService)

	// Create template for page content
	contentTmpl := template.New("page-content").Funcs(templateFuncs)

	// Load component templates
	componentFiles, err := hs.loadComponentTemplates()
	if err != nil {
		return "", fmt.Errorf("failed to load components: %w", err)
	}

	// Parse component files first
	if len(componentFiles) > 0 {
		contentTmpl, err = contentTmpl.ParseFiles(componentFiles...)
		if err != nil {
			return "", fmt.Errorf("failed to parse components: %w", err)
		}
	}

	// Parse the page content
	contentTmpl, err = contentTmpl.Parse(string(contentBytes))
	if err != nil {
		return "", fmt.Errorf("failed to parse page template: %w", err)
	}

	// Execute the page template
	var renderedContent strings.Builder
	if err = contentTmpl.Execute(&renderedContent, pageData); err != nil {
		return "", fmt.Errorf("failed to execute page template: %w", err)
	}

	// Now render with main layout
	templateFiles := []string{
		filepath.Join(hs.layoutsDir, "main.html"),
	}
	templateFiles = append(templateFiles, componentFiles...)

	// Create main template
	mainTmpl := template.New("main.html").Funcs(templateFuncs)
	mainTmpl, err = mainTmpl.ParseFiles(templateFiles...)
	if err != nil {
		return "", fmt.Errorf("failed to parse main template: %w", err)
	}

	// Prepare final page data with rendered content
	finalPageData := hs.preparePageData(urlPath, template.HTML(renderedContent.String()), false, nil, navigationService, seoService)

	// Execute main template
	var finalHTML strings.Builder
	if err := mainTmpl.ExecuteTemplate(&finalHTML, "main.html", finalPageData); err != nil {
		return "", fmt.Errorf("failed to execute main template: %w", err)
	}

	return finalHTML.String(), nil
}

// loadComponentTemplates loads all component template files
func (hs *HTMLService) loadComponentTemplates() ([]string, error) {
	componentFiles, err := filepath.Glob(filepath.Join(hs.componentsDir, "*.html"))
	if err != nil {
		return nil, err
	}
	return componentFiles, nil
}

// generateURLPath converts a file path to a clean URL path
func (hs *HTMLService) generateURLPath(filePath string) string {
	// Remove pages/ prefix and .html suffix
	urlPath := strings.TrimPrefix(filePath, hs.pagesDir+"/")
	urlPath = strings.TrimSuffix(urlPath, ".html")

	// Handle index files
	if strings.HasSuffix(urlPath, "/index") {
		urlPath = strings.TrimSuffix(urlPath, "/index")
	}

	// Add leading slash
	if urlPath == "" {
		return "/"
	}
	return "/" + urlPath
}

// isExcluded checks if a URL path should be excluded from pre-rendering
func (hs *HTMLService) isExcluded(urlPath string, excludedPages []string) bool {
	for _, excluded := range excludedPages {
		if urlPath == excluded {
			return true
		}
	}
	return false
}

// GetCachedContent retrieves pre-rendered HTML content from cache
func (hs *HTMLService) GetCachedContent(urlPath string) (*CachedContent, bool) {
	return hs.cache.Get(urlPath)
}

// GetAllCachedContent returns all cached HTML content (for search indexing)
func (hs *HTMLService) GetAllCachedContent() map[string]*CachedContent {
	return hs.cache.GetAll()
}

// GetCacheSize returns the number of cached HTML items
func (hs *HTMLService) GetCacheSize() int {
	return hs.cache.Size()
}

// preparePageData creates PageData with metadata for the given path
func (hs *HTMLService) preparePageData(path string, content template.HTML, isMarkdown bool, frontmatter *Frontmatter, navigationService *NavigationService, seoService *SEOService) PageData {
	// Get metadata from SEO service
	title, description, keywords, pageMeta, siteMeta := seoService.PreparePageMetadata(path, isMarkdown, frontmatter)

	// Prepare insights data - only include if on insights page
	var insights []InsightData
	if path == "/insights" && hs.markdownService != nil {
		// Get all cached insights from MarkdownService
		cachedInsights := hs.markdownService.GetAllCachedContent()

		// Initialize SVG generator
		svgGen := NewSVGGenerator()

		for urlPath, content := range cachedInsights {
			if strings.HasPrefix(urlPath, "/insights/") && content.Frontmatter != nil {
				// Extract category from tags or category field
				category := content.Frontmatter.Category
				if category == "" && len(content.Frontmatter.Tags) > 0 {
					// Fallback to first tag if category is not set
					category = content.Frontmatter.Tags[0]
				}

				// Generate unique SVG for this insight based on title
				svgDataURL := svgGen.GenerateSVGDataURL(content.Frontmatter.Title)

				insights = append(insights, InsightData{
					Title:       content.Frontmatter.Title,
					Description: content.Frontmatter.Description,
					Category:    category,
					Slug:        content.Frontmatter.Slug,
					SVGData:     svgDataURL,
					Date:        content.Frontmatter.Date,
					URL:         urlPath,
				})
			}
		}

		// Sort insights by date in reverse chronological order (newest first)
		sort.Slice(insights, func(i, j int) bool {
			// Parse dates, handle missing/invalid dates
			dateI, errI := time.Parse("2006-01-02", insights[i].Date)
			dateJ, errJ := time.Parse("2006-01-02", insights[j].Date)

			// If both dates are invalid, maintain original order
			if errI != nil && errJ != nil {
				return false
			}

			// If one date is invalid, put it at the end
			if errI != nil {
				return false
			}
			if errJ != nil {
				return true
			}

			// Both dates are valid, sort newest first
			return dateI.After(dateJ)
		})

		// Insights loaded for page
	}

	// Extract table of contents
	toc := make([]TOCEntry, 0)
	tocExcludedPaths := []string{"/changelog", "/roadmap", "/platform/status"}

	isExcluded := false
	for _, excludedPath := range tocExcludedPaths {
		if path == excludedPath {
			isExcluded = true
			break
		}
	}

	if !isExcluded && string(content) != "" {
		var err error
		if isMarkdown {
			// For markdown content, extract H2 elements from converted HTML
			toc, err = ExtractH2TOC(string(content))
		} else {
			// For HTML pages, extract from section elements
			toc, err = ExtractHTMLTOC(string(content))
		}

		if err != nil {
			log.Printf("Error extracting TOC for path=%s: %v", path, err)
		} else {
			// TOC extracted
		}
	} else if isExcluded {
		// TOC generation skipped
	}

	// Return PageData with all components
	return PageData{
		Title:          title,
		Content:        content,
		Navigation:     navigationService.GetNavigationForPath(path),
		PageMeta:       pageMeta,
		SiteMeta:       siteMeta,
		Description:    description,
		Keywords:       keywords,
		IsMarkdown:     isMarkdown,
		Frontmatter:    frontmatter,
		TOC:            toc,
		CustomerNumber: 17000,
		Insights:       insights,
		Path:           path,
	}
}
