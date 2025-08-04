package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

// PlaceholderType represents different types of content that need preservation
type PlaceholderType string

const (
	PlaceholderCodeBlock   PlaceholderType = "CB"
	PlaceholderInlineCode  PlaceholderType = "CODE"
	PlaceholderTableCell   PlaceholderType = "CELL"
	PlaceholderURL         PlaceholderType = "URL"
	PlaceholderEmail       PlaceholderType = "EMAIL"
	PlaceholderCallout     PlaceholderType = "CALLOUT"
	PlaceholderLink        PlaceholderType = "LINK"
	PlaceholderLegalRef    PlaceholderType = "LEGALREF"
)

// Placeholder represents a masked piece of content
type Placeholder struct {
	ID           string
	Type         PlaceholderType
	Content      string
	Position     int
	LineNumber   int
	BeforeText   string // Context for recovery
	AfterText    string // Context for recovery
}

// DocumentProcessor handles the parsing and processing of legal documentation
type DocumentProcessor struct {
	parser goldmark.Markdown
}

// NewDocumentProcessor creates a new document processor
func NewDocumentProcessor() *DocumentProcessor {
	md := goldmark.New(
		goldmark.WithExtensions(extension.Table),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)
	
	return &DocumentProcessor{
		parser: md,
	}
}

// ProcessDocument extracts and masks technical content, returning masked content and placeholder map
func (p *DocumentProcessor) ProcessDocument(content string) (string, map[string]Placeholder, error) {
	placeholderMap := make(map[string]Placeholder)
	
	// Extract frontmatter first to preserve it separately
	frontmatter, bodyContent := p.extractFrontmatter(content)
	
	// Pre-process to handle special markdown extensions (callouts)
	bodyContent, calloutMap := p.extractCallouts(bodyContent, placeholderMap)
	
	// Extract code blocks first (they're the largest structures)
	bodyContent, codeBlockMap := p.extractCodeBlocks(bodyContent, placeholderMap)
	
	// Extract inline code
	bodyContent, inlineCodeMap := p.extractInlineCode(bodyContent, placeholderMap)
	
	// Extract internal legal documentation links
	bodyContent, linkMap := p.extractLegalLinks(bodyContent, placeholderMap)
	
	// Extract URLs and emails
	bodyContent = p.extractURLsAndEmails(bodyContent, placeholderMap)
	
	// Extract legal references (section numbers, dates, percentages, amounts)
	bodyContent = p.extractLegalReferences(bodyContent, placeholderMap)
	
	// Merge all placeholder maps
	for k, v := range calloutMap {
		placeholderMap[k] = v
	}
	for k, v := range codeBlockMap {
		placeholderMap[k] = v
	}
	for k, v := range inlineCodeMap {
		placeholderMap[k] = v
	}
	for k, v := range linkMap {
		placeholderMap[k] = v
	}
	
	// Combine frontmatter with processed body
	finalContent := frontmatter + bodyContent
	
	return finalContent, placeholderMap, nil
}

// extractFrontmatter separates YAML frontmatter from body content
func (p *DocumentProcessor) extractFrontmatter(content string) (string, string) {
	if !strings.HasPrefix(content, "---\n") {
		return "", content
	}
	
	// Find the closing --- of the frontmatter
	endIndex := strings.Index(content[4:], "\n---\n")
	if endIndex == -1 {
		return "", content
	}
	
	// Add 4 for the opening ---, and 5 for the closing ---\n
	frontmatterEnd := endIndex + 4 + 5
	frontmatter := content[:frontmatterEnd]
	bodyContent := content[frontmatterEnd:]
	
	return frontmatter, bodyContent
}

// extractCallouts extracts Blue's callout blocks
func (p *DocumentProcessor) extractCallouts(content string, placeholderMap map[string]Placeholder) (string, map[string]Placeholder) {
	calloutRegex := regexp.MustCompile(`(?s)::callout\{[^\}]*\}\n(.*?)\n::`)
	
	matches := calloutRegex.FindAllStringSubmatchIndex(content, -1)
	
	// Process matches in reverse to maintain positions
	for i := len(matches) - 1; i >= 0; i-- {
		match := matches[i]
		fullMatch := content[match[0]:match[1]]
		
		// Generate placeholder
		id := uuid.New().String()
		placeholder := fmt.Sprintf("@@%s##%s##%s@@", PlaceholderCallout, id, PlaceholderCallout)
		
		placeholderMap[placeholder] = Placeholder{
			ID:       id,
			Type:     PlaceholderCallout,
			Content:  fullMatch,
			Position: match[0],
		}
		
		// Replace in content
		content = content[:match[0]] + placeholder + content[match[1]:]
	}
	
	return content, placeholderMap
}

// extractCodeBlocks extracts markdown code blocks
func (p *DocumentProcessor) extractCodeBlocks(content string, placeholderMap map[string]Placeholder) (string, map[string]Placeholder) {
	// Regex to match code blocks with optional language specifier
	codeBlockRegex := regexp.MustCompile("(?s)```[^\n]*\n.*?\n```")
	
	matches := codeBlockRegex.FindAllStringSubmatchIndex(content, -1)
	
	// Process matches in reverse to maintain positions
	for i := len(matches) - 1; i >= 0; i-- {
		match := matches[i]
		codeBlock := content[match[0]:match[1]]
		
		// Generate placeholder
		id := uuid.New().String()
		placeholder := fmt.Sprintf("@@%s##%s##%s@@", PlaceholderCodeBlock, id, PlaceholderCodeBlock)
		
		placeholderMap[placeholder] = Placeholder{
			ID:       id,
			Type:     PlaceholderCodeBlock,
			Content:  codeBlock,
			Position: match[0],
		}
		
		// Replace in content
		content = content[:match[0]] + placeholder + content[match[1]:]
	}
	
	return content, placeholderMap
}

// extractInlineCode extracts inline code snippets
func (p *DocumentProcessor) extractInlineCode(content string, placeholderMap map[string]Placeholder) (string, map[string]Placeholder) {
	// Match inline code that's not part of a code block placeholder
	inlineCodeRegex := regexp.MustCompile("`[^`]+`")
	
	matches := inlineCodeRegex.FindAllStringIndex(content, -1)
	
	// Process matches in reverse
	for i := len(matches) - 1; i >= 0; i-- {
		match := matches[i]
		inlineCode := content[match[0]:match[1]]
		
		// Skip if this is within a placeholder
		if strings.Contains(content[max(0, match[0]-10):min(len(content), match[1]+10)], "@@") {
			continue
		}
		
		// Generate placeholder
		id := uuid.New().String()
		placeholder := fmt.Sprintf("@@%s##%s##%s@@", PlaceholderInlineCode, id, PlaceholderInlineCode)
		
		placeholderMap[placeholder] = Placeholder{
			ID:       id,
			Type:     PlaceholderInlineCode,
			Content:  inlineCode,
			Position: match[0],
		}
		
		// Replace in content
		content = content[:match[0]] + placeholder + content[match[1]:]
	}
	
	return content, placeholderMap
}

// extractLegalLinks extracts internal legal documentation links
func (p *DocumentProcessor) extractLegalLinks(content string, placeholderMap map[string]Placeholder) (string, map[string]Placeholder) {
	// Match markdown links that reference /legal/ paths
	linkRegex := regexp.MustCompile(`\[[^\]]+\]\(/legal/[^\)]+\)`)
	
	matches := linkRegex.FindAllStringIndex(content, -1)
	
	// Process matches in reverse
	for i := len(matches) - 1; i >= 0; i-- {
		match := matches[i]
		link := content[match[0]:match[1]]
		
		// Generate placeholder
		id := uuid.New().String()
		placeholder := fmt.Sprintf("@@%s##%s##%s@@", PlaceholderLink, id, PlaceholderLink)
		
		placeholderMap[placeholder] = Placeholder{
			ID:       id,
			Type:     PlaceholderLink,
			Content:  link,
			Position: match[0],
		}
		
		// Replace in content
		content = content[:match[0]] + placeholder + content[match[1]:]
	}
	
	return content, placeholderMap
}

// extractURLsAndEmails extracts URLs and email addresses
func (p *DocumentProcessor) extractURLsAndEmails(content string, placeholderMap map[string]Placeholder) string {
	// Extract URLs
	urlRegex := regexp.MustCompile(`https?://[^\s\)]+|www\.[^\s\)]+`)
	matches := urlRegex.FindAllStringIndex(content, -1)
	
	for i := len(matches) - 1; i >= 0; i-- {
		match := matches[i]
		url := content[match[0]:match[1]]
		
		id := uuid.New().String()
		placeholder := fmt.Sprintf("@@%s##%s##%s@@", PlaceholderURL, id, PlaceholderURL)
		
		placeholderMap[placeholder] = Placeholder{
			ID:       id,
			Type:     PlaceholderURL,
			Content:  url,
			Position: match[0],
		}
		
		content = content[:match[0]] + placeholder + content[match[1]:]
	}
	
	// Extract emails
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	matches = emailRegex.FindAllStringIndex(content, -1)
	
	for i := len(matches) - 1; i >= 0; i-- {
		match := matches[i]
		email := content[match[0]:match[1]]
		
		id := uuid.New().String()
		placeholder := fmt.Sprintf("@@%s##%s##%s@@", PlaceholderEmail, id, PlaceholderEmail)
		
		placeholderMap[placeholder] = Placeholder{
			ID:       id,
			Type:     PlaceholderEmail,
			Content:  email,
			Position: match[0],
		}
		
		content = content[:match[0]] + placeholder + content[match[1]:]
	}
	
	return content
}

// extractLegalReferences extracts legal-specific references like section numbers, dates, amounts
func (p *DocumentProcessor) extractLegalReferences(content string, placeholderMap map[string]Placeholder) string {
	// Extract section numbers (e.g., 1.1, 2.3.4)
	sectionRegex := regexp.MustCompile(`\b\d+(\.\d+)+\b`)
	matches := sectionRegex.FindAllStringIndex(content, -1)
	
	for i := len(matches) - 1; i >= 0; i-- {
		match := matches[i]
		section := content[match[0]:match[1]]
		
		// Skip if already in a placeholder
		if strings.Contains(content[max(0, match[0]-10):min(len(content), match[1]+10)], "@@") {
			continue
		}
		
		id := uuid.New().String()
		placeholder := fmt.Sprintf("@@%s##%s##%s@@", PlaceholderLegalRef, id, PlaceholderLegalRef)
		
		placeholderMap[placeholder] = Placeholder{
			ID:       id,
			Type:     PlaceholderLegalRef,
			Content:  section,
			Position: match[0],
		}
		
		content = content[:match[0]] + placeholder + content[match[1]:]
	}
	
	// Extract monetary amounts (e.g., $200, $1,000)
	amountRegex := regexp.MustCompile(`\$[\d,]+`)
	matches = amountRegex.FindAllStringIndex(content, -1)
	
	for i := len(matches) - 1; i >= 0; i-- {
		match := matches[i]
		amount := content[match[0]:match[1]]
		
		if strings.Contains(content[max(0, match[0]-10):min(len(content), match[1]+10)], "@@") {
			continue
		}
		
		id := uuid.New().String()
		placeholder := fmt.Sprintf("@@%s##%s##%s@@", PlaceholderLegalRef, id, PlaceholderLegalRef)
		
		placeholderMap[placeholder] = Placeholder{
			ID:       id,
			Type:     PlaceholderLegalRef,
			Content:  amount,
			Position: match[0],
		}
		
		content = content[:match[0]] + placeholder + content[match[1]:]
	}
	
	// Extract percentages (e.g., 25%, 42%)
	percentRegex := regexp.MustCompile(`\b\d+%`)
	matches = percentRegex.FindAllStringIndex(content, -1)
	
	for i := len(matches) - 1; i >= 0; i-- {
		match := matches[i]
		percent := content[match[0]:match[1]]
		
		if strings.Contains(content[max(0, match[0]-10):min(len(content), match[1]+10)], "@@") {
			continue
		}
		
		id := uuid.New().String()
		placeholder := fmt.Sprintf("@@%s##%s##%s@@", PlaceholderLegalRef, id, PlaceholderLegalRef)
		
		placeholderMap[placeholder] = Placeholder{
			ID:       id,
			Type:     PlaceholderLegalRef,
			Content:  percent,
			Position: match[0],
		}
		
		content = content[:match[0]] + placeholder + content[match[1]:]
	}
	
	return content
}

// ValidateTranslation checks if the translation preserved all placeholders
func (p *DocumentProcessor) ValidateTranslation(original, translated string) error {
	// Extract all placeholders from original
	placeholderRegex := regexp.MustCompile(`@@[A-Z]+##[a-f0-9-]+##[A-Z]+@@`)
	originalPlaceholders := placeholderRegex.FindAllString(original, -1)
	translatedPlaceholders := placeholderRegex.FindAllString(translated, -1)
	
	// Check if all placeholders are present
	if len(originalPlaceholders) != len(translatedPlaceholders) {
		return fmt.Errorf("placeholder count mismatch: original has %d, translated has %d",
			len(originalPlaceholders), len(translatedPlaceholders))
	}
	
	// Create a map to check presence
	placeholderMap := make(map[string]bool)
	for _, p := range originalPlaceholders {
		placeholderMap[p] = true
	}
	
	// Check each translated placeholder
	for _, p := range translatedPlaceholders {
		if !placeholderMap[p] {
			return fmt.Errorf("placeholder %s not found in original", p)
		}
		delete(placeholderMap, p)
	}
	
	// Check if any placeholders are missing
	if len(placeholderMap) > 0 {
		var missing []string
		for p := range placeholderMap {
			missing = append(missing, p)
		}
		return fmt.Errorf("missing placeholders in translation: %v", missing)
	}
	
	return nil
}

// RecoverPlaceholders attempts to recover missing placeholders
func (p *DocumentProcessor) RecoverPlaceholders(original, translated string) string {
	// This is a fallback method to try to preserve content structure
	// In practice, it's better to retry the translation than to use this
	return translated
}

// RestoreContent replaces placeholders with their original content
func (p *DocumentProcessor) RestoreContent(content string, placeholderMap map[string]Placeholder) string {
	// Sort placeholders by position to restore in order
	for placeholder, data := range placeholderMap {
		content = strings.ReplaceAll(content, placeholder, data.Content)
	}
	
	return content
}

// Helper functions
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}