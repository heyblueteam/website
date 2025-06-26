package web

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// NavItem represents a navigation item
type NavItem struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Href     string    `json:"href,omitempty"`
	Expanded bool      `json:"expanded,omitempty"`
	Children []NavItem `json:"children,omitempty"`
}

// Navigation holds the complete navigation structure
type Navigation struct {
	Sections []NavItem `json:"sections"`
	Legal    []NavItem `json:"legal"`
}

// Router handles file-based routing for HTML pages
type Router struct {
	pagesDir      string
	layoutsDir    string
	componentsDir string
	navigation    *Navigation
}

// NewRouter creates a new router instance
func NewRouter(pagesDir string) *Router {
	router := &Router{
		pagesDir:      pagesDir,
		layoutsDir:    "layouts",
		componentsDir: "components",
	}
	
	// Load navigation data
	if err := router.loadNavigation(); err != nil {
		log.Printf("Error loading navigation: %v", err)
	}
	
	return router
}

// loadNavigation loads navigation data from JSON file
func (r *Router) loadNavigation() error {
	data, err := os.ReadFile("data/nav.json")
	if err != nil {
		return err
	}
	
	r.navigation = &Navigation{}
	return json.Unmarshal(data, r.navigation)
}

// PageData holds data for template rendering
type PageData struct {
	Title      string
	Content    template.HTML
	Navigation *Navigation
}

// ServeHTTP implements the http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Skip public file requests
	if strings.HasPrefix(req.URL.Path, "/public/") {
		return
	}

	// Get the requested path
	path := req.URL.Path
	
	// Redirect .html URLs to clean URLs
	if strings.HasSuffix(path, ".html") && path != "/" {
		cleanURL := strings.TrimSuffix(path, ".html")
		http.Redirect(w, req, cleanURL, http.StatusMovedPermanently)
		return
	}
	
	// Convert URL path to file path
	var filePath string
	if path == "/" {
		// Root path maps to index.html
		filePath = filepath.Join(r.pagesDir, "index.html")
	} else {
		// Remove leading slash
		cleanPath := strings.TrimPrefix(path, "/")
		
		if strings.HasSuffix(cleanPath, "/") {
			// Directory path, look for index.html
			filePath = filepath.Join(r.pagesDir, cleanPath, "index.html")
		} else {
			// Try as direct .html file first
			filePath = filepath.Join(r.pagesDir, cleanPath+".html")
			
			// If that doesn't exist, try as directory with index.html
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				indexPath := filepath.Join(r.pagesDir, cleanPath, "index.html")
				if _, err := os.Stat(indexPath); err == nil {
					// Redirect to trailing slash version
					http.Redirect(w, req, path+"/", http.StatusMovedPermanently)
					return
				}
			}
		}
	}

	// Check if page file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.NotFound(w, req)
		return
	}

	// Read the page content
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Error reading page", http.StatusInternalServerError)
		log.Printf("Page reading error: %v", err)
		return
	}

	// Prepare template files - start with layout
	templateFiles := []string{
		filepath.Join(r.layoutsDir, "main.html"),
	}
	
	// Auto-scan all component templates
	componentFiles, err := filepath.Glob(filepath.Join(r.componentsDir, "*.html"))
	if err != nil {
		http.Error(w, "Error loading components", http.StatusInternalServerError)
		log.Printf("Component scanning error: %v", err)
		return
	}
	
	// Add all component files
	templateFiles = append(templateFiles, componentFiles...)

	// Create template with custom functions
	tmpl := template.New("main.html").Funcs(template.FuncMap{
		"toJSON": func(v any) template.JS {
			data, _ := json.Marshal(v)
			return template.JS(data)
		},
	})
	
	// Parse all template files  
	tmpl, err = tmpl.ParseFiles(templateFiles...)
	if err != nil {
		http.Error(w, "Error loading templates", http.StatusInternalServerError)
		log.Printf("Template parsing error: %v", err)
		return
	}

	// Prepare page data
	pageData := PageData{
		Title:      r.getPageTitle(path),
		Content:    template.HTML(contentBytes),
		Navigation: r.navigation,
	}

	// Set content type and execute main layout
	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.ExecuteTemplate(w, "main.html", pageData); err != nil {
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
		log.Printf("Template execution error: %v", err)
		return
	}
}

// getPageTitle extracts a title from the URL path
func (r *Router) getPageTitle(path string) string {
	if path == "/" {
		return "Home"
	}
	
	// Remove leading slash and convert to title case
	cleanPath := strings.TrimPrefix(path, "/")
	cleanPath = strings.TrimSuffix(cleanPath, "/")
	
	// Replace slashes with spaces and title case
	parts := strings.Split(cleanPath, "/")
	for i, part := range parts {
		parts[i] = strings.Title(strings.ReplaceAll(part, "-", " "))
	}
	
	return strings.Join(parts, " - ")
}