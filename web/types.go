package web

import (
	"html/template"
	"time"
)

// InsightData represents an insight for template rendering
type InsightData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Slug        string `json:"slug"`
	PNGPath     string `json:"png_path"`
	Date        string `json:"date"`
	URL         string `json:"url"`
}

// StatusPageData represents status information for template rendering
type StatusPageData struct {
	Services  []ServiceHistory `json:"services"`
	Generated time.Time        `json:"generated"`
}

// PageData holds data for template rendering
type PageData struct {
	Title              string
	Content            template.HTML
	Navigation         *Navigation
	PageMeta           *PageMetadata
	SiteMeta           *SiteMetadata
	Description        string
	Keywords           []string
	IsMarkdown         bool
	Frontmatter        *Frontmatter
	TOC                []TOCEntry
	CustomerNumber     int
	Insights           []InsightData
	Path               string
	SchemaData         template.JS
	Language           string
	LanguageLocale     string
	SupportedLanguages []string
	StatusData         *StatusPageData
	NeedsCodeHighlight bool
	ForceAuthState     bool
}