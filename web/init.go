package web

import (
	"fmt"
	"log"
	"path/filepath"
	"time"
)

// NewRouter creates a new router instance with all services initialized
func NewRouter(pagesDir string, logger *Logger) *Router {
	// Initialize SEO service
	seoService := NewSEOService()
	if err := seoService.LoadData(); err != nil {
		log.Printf("⚠️  Error loading SEO data: %v", err)
	}

	// Initialize services
	markdownService := NewMarkdownService()
	contentService := NewContentService("content")
	navigationService := NewNavigationService(seoService)
	htmlService := NewHTMLService(pagesDir, "layouts", "components", markdownService)
	schemaService := NewSchemaService(seoService.metadata, "https://blue.cc")
	
	// Set schema service on HTML service
	htmlService.SetSchemaService(schemaService)

	router := &Router{
		pagesDir:          pagesDir,
		layoutsDir:        "layouts",
		componentsDir:     "components",
		contentDir:        "content",
		markdownService:   markdownService,
		contentService:    contentService,
		navigationService: navigationService,
		htmlService:       htmlService,
		seoService:        seoService,
		schemaService:     schemaService,
		statusChecker:     nil, // Will be set by SetStatusChecker
		tocExcludedPaths: []string{ // These pages will not show toc
			"/changelog",
			"/roadmap",
			"/platform/status",
		},
		logger: logger,
	}

	// Pre-render all markdown content at startup
	renderStart := time.Now()
	logger.Log(LogRender, "📝", "Starting", "Markdown pre-rendering")
	if err := markdownService.PreRenderAllMarkdown(contentService, seoService); err != nil {
		logger.Log(LogError, "⚠️", "Failed", fmt.Sprintf("Pre-render markdown: %v", err))
	} else {
		fileCount := markdownService.GetCacheSize()
		logger.Log(LogRender, "✅", "Completed", fmt.Sprintf("%d Markdown files cached", fileCount), time.Since(renderStart))
	}

	// Pre-render all HTML pages at startup
	htmlStart := time.Now()
	logger.Log(LogRender, "🌐", "Starting", "HTML pre-rendering")
	if err := htmlService.PreRenderAllHTMLPages(navigationService, seoService); err != nil {
		logger.Log(LogError, "⚠️", "Failed", fmt.Sprintf("Pre-render HTML: %v", err))
	} else {
		pageCount := htmlService.GetCacheSize()
		logger.Log(LogRender, "✅", "Completed", fmt.Sprintf("%d pages across %d languages", pageCount, len(SupportedLanguages)), time.Since(htmlStart))
	}

	// Generate search index with pre-rendered content
	searchStart := time.Now()
	logger.Log(LogIndex, "🔍", "Starting", "Search index generation")
	itemCount, err := GenerateSearchIndexWithCaches(markdownService, htmlService, logger)
	if err != nil {
		logger.Log(LogError, "⚠️", "Failed", fmt.Sprintf("Generate search index: %v", err))
	} else {
		logger.Log(LogIndex, "✅", "Completed", fmt.Sprintf("%d items indexed", itemCount), time.Since(searchStart))
	}

	// Link checker moved to main.go for parallel execution

	return router
}

// loadComponentTemplates loads all component template files
func (r *Router) loadComponentTemplates() ([]string, error) {
	componentFiles, err := filepath.Glob(filepath.Join(r.componentsDir, "*.html"))
	if err != nil {
		return nil, err
	}
	return componentFiles, nil
}

// SetStatusChecker sets the status checker for the router
func (r *Router) SetStatusChecker(checker *HealthChecker) {
	r.statusChecker = checker
	// Also set it on the HTML service for status page rendering
	if r.htmlService != nil {
		r.htmlService.SetStatusChecker(checker)
	}
}