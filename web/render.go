package web

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// RenderHTML renders HTML pages (cached or dynamic)
func (r *Router) RenderHTML(w http.ResponseWriter, req *http.Request, filePath, path, lang string) {
	// Try cached version first
	if cachedContent, found := r.htmlService.GetCachedContentForLang(path, lang); found {
		// Found in cache - use pre-rendered content
		w.Header().Set("Content-Type", "text/html")
		
		// For SPA requests, include code highlight flag in header
		if req.Header.Get("X-Requested-With") == "XMLHttpRequest" {
			if cachedContent.NeedsCodeHighlight {
				w.Header().Set("X-Needs-Code-Highlight", "true")
			}
		}
		
		_, err := w.Write([]byte(cachedContent.HTML))
		if err != nil {
			// Check if the error is due to client disconnect
			if strings.Contains(err.Error(), "broken pipe") || strings.Contains(err.Error(), "connection reset by peer") {
				// Client disconnected, silently return
				return
			}
			// Log other write errors
			log.Printf("Error writing cached content: %v", err)
		}
		return
	}

	// Not in cache, process on-demand (fallback for dynamic pages like status)
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Error reading page", http.StatusInternalServerError)
		log.Printf("Page reading error: %v", err)
		return
	}

	// Process the page content through templates
	processedContent, err := r.processPageContent(contentBytes, path, lang)
	if err != nil {
		http.Error(w, "Error processing page", http.StatusInternalServerError)
		log.Printf("Page processing error: %v", err)
		return
	}

	// Render with main layout
	pageData := r.preparePageData(path, template.HTML(processedContent), false, nil, r.navigationService.GetNavigationForPathWithLanguage(path, lang), lang)
	if err := r.RenderWithTemplate(w, processedContent, &pageData, lang); err != nil {
		log.Printf("Template rendering error: %v", err)
	}
}

// RenderMarkdown renders markdown content with layout
func (r *Router) RenderMarkdown(w http.ResponseWriter, req *http.Request, path, lang string) {
	var contentBytes []byte
	var frontmatter *Frontmatter
	var cachedContent *CachedContent

	// Check cached markdown first
	if cached, found := r.markdownService.GetCachedContentForLang(path, lang); found {
		// Found in cache - use pre-rendered content
		cachedContent = cached
		contentBytes = []byte(cached.HTML)
		frontmatter = cached.Frontmatter
		
		// For SPA requests, we need to pass the code highlight flag
		if req.Header.Get("X-Requested-With") == "XMLHttpRequest" && cached.NeedsCodeHighlight {
			w.Header().Set("X-Needs-Code-Highlight", "true")
		}
	} else {
		// Not in cache, try to find and process markdown file (fallback)
		markdownPath, mdErr := r.contentService.FindMarkdownFileForLang(path, lang)
		if mdErr != nil {
			// Before serving 404, check if this is a directory URL that should redirect
			if firstItemURL := r.navigationService.GetFirstItemInDirectory(path); firstItemURL != "" {
				http.Redirect(w, req, firstItemURL, http.StatusMovedPermanently)
				return
			}
			r.Render404(w, req)
			return
		}

		// Process markdown file on-demand
		htmlContent, fm, err := r.markdownService.ProcessMarkdownFile(markdownPath, r.seoService)
		if err != nil {
			http.Error(w, "Error processing markdown file", http.StatusInternalServerError)
			log.Printf("Markdown processing error: %v", err)
			return
		}

		contentBytes = []byte(htmlContent)
		frontmatter = fm
	}

	// Prepare page data and render with template
	pageData := r.preparePageDataWithCache(path, template.HTML(contentBytes), true, frontmatter, r.navigationService.GetNavigationForPathWithLanguage(path, lang), lang, cachedContent)
	
	// Set content type and render
	w.Header().Set("Content-Type", "text/html")
	if err := r.RenderWithTemplate(w, contentBytes, &pageData, lang); err != nil {
		log.Printf("Markdown template rendering error: %v", err)
	}
}

// RenderWithTemplate applies the main layout template to content
func (r *Router) RenderWithTemplate(w http.ResponseWriter, content []byte, pageData *PageData, lang string) error {
	// Prepare template files - start with layout
	templateFiles := []string{
		filepath.Join(r.layoutsDir, "main.html"),
	}

	// Auto-scan all component templates
	componentFiles, err := r.loadComponentTemplates()
	if err != nil {
		http.Error(w, "Error loading components", http.StatusInternalServerError)
		return err
	}

	// Add all component files
	templateFiles = append(templateFiles, componentFiles...)

	// Create template with custom functions including language-specific translation
	tmpl := template.New("main.html").Funcs(getTemplateFuncs(lang))

	// Parse all template files
	tmpl, err = tmpl.ParseFiles(templateFiles...)
	if err != nil {
		http.Error(w, "Error loading templates", http.StatusInternalServerError)
		return err
	}

	// Update page data with final content
	pageData.Content = template.HTML(content)

	// Execute main layout
	if err := tmpl.ExecuteTemplate(w, "main.html", pageData); err != nil {
		// Check if the error is due to client disconnect (broken pipe)
		if strings.Contains(err.Error(), "broken pipe") || strings.Contains(err.Error(), "connection reset by peer") {
			// Client disconnected, this is expected during rapid navigation
			// Silently return without logging
			return nil
		}
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
		return err
	}

	return nil
}

// processPageContent processes HTML page content as a template
func (r *Router) processPageContent(contentBytes []byte, path, lang string) ([]byte, error) {
	// Process all HTML pages as templates to enable template variables
	pageData := r.preparePageData(path, "", false, nil, r.navigationService.GetNavigationForPathWithLanguage(path, lang), lang)

	// Create a template for the page content with language-specific functions
	contentTmpl := template.New("page-content").Funcs(getTemplateFuncs(lang))

	// Auto-scan all component templates for page content parsing
	componentFiles, err := r.loadComponentTemplates()
	if err != nil {
		return nil, err
	}

	// Parse component files first
	if len(componentFiles) > 0 {
		contentTmpl, err = contentTmpl.ParseFiles(componentFiles...)
		if err != nil {
			return nil, err
		}
	}

	// Then parse the page content
	contentTmpl, err = contentTmpl.Parse(string(contentBytes))
	if err != nil {
		return nil, err
	}

	// Execute the page template with the page data
	var renderedContent strings.Builder
	if err = contentTmpl.Execute(&renderedContent, pageData); err != nil {
		return nil, err
	}

	return []byte(renderedContent.String()), nil
}

// Render404 serves a custom 404 page or falls back to default
func (r *Router) Render404(w http.ResponseWriter, req *http.Request) {
	// Extract language from the request URL
	lang, _ := extractLanguageFromPath(req.URL.Path)
	if lang == "" {
		lang = detectPreferredLanguage(req)
	}
	// Try to serve custom 404.html page
	notFoundPath := filepath.Join(r.pagesDir, "404.html")
	if _, err := os.Stat(notFoundPath); err == nil {
		// Custom 404 page exists, we'll set status when we write the response

		// Read the 404 page content
		contentBytes, err := os.ReadFile(notFoundPath)
		if err != nil {
			// Fallback to default if we can't read the file
			http.NotFound(w, req)
			return
		}

		// Prepare page data for 404 page (before processing templates)
		pageData := r.preparePageData("/404", "", false, nil, r.navigationService.GetNavigationForPathWithLanguage("/404", lang), lang)

		// Create a template for the 404 page content
		contentTmpl := template.New("404-content").Funcs(getTemplateFuncs(lang))

		// Auto-scan all component templates for 404 content parsing
		componentFiles, err := r.loadComponentTemplates()
		if err != nil {
			http.NotFound(w, req)
			return
		}

		// Parse component files first
		if len(componentFiles) > 0 {
			contentTmpl, err = contentTmpl.ParseFiles(componentFiles...)
			if err != nil {
				http.NotFound(w, req)
				return
			}
		}

		// Then parse the 404 page content
		contentTmpl, err = contentTmpl.Parse(string(contentBytes))
		if err != nil {
			http.NotFound(w, req)
			return
		}

		// Execute the 404 content template with the page data
		var renderedContent strings.Builder
		if err = contentTmpl.Execute(&renderedContent, pageData); err != nil {
			http.NotFound(w, req)
			return
		}

		// Now prepare template files for main layout
		templateFiles := []string{
			filepath.Join(r.layoutsDir, "main.html"),
		}
		templateFiles = append(templateFiles, componentFiles...)

		// Create template with custom functions for main layout
		tmpl := template.New("main.html").Funcs(getTemplateFuncs(lang))

		// Parse all template files
		tmpl, err = tmpl.ParseFiles(templateFiles...)
		if err != nil {
			http.NotFound(w, req)
			return
		}

		// Update page data with rendered content
		pageData.Content = template.HTML(renderedContent.String())

		// Set 404 status and execute main layout template
		w.WriteHeader(http.StatusNotFound)
		if err := tmpl.ExecuteTemplate(w, "main.html", pageData); err != nil {
			// Check if the error is due to client disconnect (broken pipe)
			if strings.Contains(err.Error(), "broken pipe") || strings.Contains(err.Error(), "connection reset by peer") {
				// Client disconnected, this is expected during rapid navigation
				// Silently return without logging
				return
			}
			// Headers already written, just log the error
			log.Printf("Error executing 404 template: %v", err)
			return
		}
	} else {
		// Custom 404 page doesn't exist, use default
		http.NotFound(w, req)
	}
}