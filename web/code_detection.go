package web

import (
	"strings"
)

// DetectCodeBlocks checks if HTML content contains code blocks or inline code
// Returns true if any code elements are found (both block and inline)
func DetectCodeBlocks(html string) bool {
	// Check for code blocks (pre > code)
	if strings.Contains(html, "<pre>") && strings.Contains(html, "<code>") {
		return true
	}
	
	// Check for standalone code tags (inline code)
	if strings.Contains(html, "<code>") {
		return true
	}
	
	// Check for language-specific code blocks that might use classes
	if strings.Contains(html, "class=\"language-") {
		return true
	}
	
	// Check for highlight.js specific classes
	if strings.Contains(html, "class=\"hljs") {
		return true
	}
	
	return false
}