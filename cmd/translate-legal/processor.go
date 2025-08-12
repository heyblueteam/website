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

// DocumentProcessor handles the parsing and processing of API documentation
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
	
	// Process tables separately before parsing (easier than using AST for tables)
	bodyContent, tableMap := p.processTables(bodyContent, placeholderMap)
	
	// Extract inline code
	bodyContent, inlineCodeMap := p.extractInlineCode(bodyContent, placeholderMap)
	
	// Extract URLs and emails
	bodyContent = p.extractURLsAndEmails(bodyContent, placeholderMap)
	
	// Merge all placeholder maps
	for k, v := range calloutMap {
		placeholderMap[k] = v
	}
	for k, v := range codeBlockMap {
		placeholderMap[k] = v
	}
	for k, v := range tableMap {
		placeholderMap[k] = v
	}
	for k, v := range inlineCodeMap {
		placeholderMap[k] = v
	}
	
	// Combine frontmatter with processed body
	finalContent := frontmatter + bodyContent
	
	return finalContent, placeholderMap, nil
}

// extractFrontmatter separates frontmatter from body content
func (p *DocumentProcessor) extractFrontmatter(content string) (string, string) {
	// Check if content starts with frontmatter
	if !strings.HasPrefix(content, "---") {
		return "", content
	}
	
	// Find the closing frontmatter delimiter
	lines := strings.Split(content, "\n")
	closingIndex := -1
	
	for i := 1; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "---" {
			closingIndex = i
			break
		}
	}
	
	if closingIndex == -1 {
		// No closing delimiter found
		return "", content
	}
	
	// Extract frontmatter (including delimiters) and body
	frontmatterLines := lines[:closingIndex+1]
	bodyLines := lines[closingIndex+1:]
	
	// Process frontmatter to ensure it's translatable
	processedFrontmatter := p.processFrontmatterForTranslation(frontmatterLines)
	
	return processedFrontmatter, strings.Join(bodyLines, "\n")
}

// processFrontmatterForTranslation preserves frontmatter structure exactly
func (p *DocumentProcessor) processFrontmatterForTranslation(lines []string) string {
	// Create a version with only translatable parts exposed
	var translatableVersion []string
	translatableVersion = append(translatableVersion, "---")
	
	for i := 1; i < len(lines)-1; i++ {
		line := strings.TrimSpace(lines[i])
		if strings.HasPrefix(line, "title:") {
			// Extract the title value for translation
			titleValue := strings.TrimSpace(strings.TrimPrefix(line, "title:"))
			translatableVersion = append(translatableVersion, fmt.Sprintf("title: %s", titleValue))
		} else if strings.HasPrefix(line, "description:") {
			// Extract the description value for translation
			descValue := strings.TrimSpace(strings.TrimPrefix(line, "description:"))
			translatableVersion = append(translatableVersion, fmt.Sprintf("description: %s", descValue))
		}
	}
	
	translatableVersion = append(translatableVersion, "---")
	
	// Return the translatable version
	return strings.Join(translatableVersion, "\n") + "\n"
}

// extractCallouts handles special ::callout blocks
func (p *DocumentProcessor) extractCallouts(content string, placeholderMap map[string]Placeholder) (string, map[string]Placeholder) {
	// Match callouts and process their content
	calloutPattern := regexp.MustCompile(`(?s)(::callout\s*)((?:---\s*[^-]+---\s*)?)(.*?)(::)`)
	localMap := make(map[string]Placeholder)
	
	content = calloutPattern.ReplaceAllStringFunc(content, func(match string) string {
		// Extract the parts of the callout
		matches := calloutPattern.FindStringSubmatch(match)
		if len(matches) != 5 {
			return match
		}
		
		calloutStart := matches[1]  // ::callout
		frontmatter := matches[2]   // ---\nicon: ...\n---
		calloutContent := matches[3] // The actual content to translate
		calloutEnd := matches[4]     // ::
		
		// Create placeholders for the callout structure
		startPlaceholder := p.generatePlaceholder(PlaceholderCallout)
		startUUID := p.extractPlaceholderID(startPlaceholder)
		localMap[startUUID] = Placeholder{
			ID:      startUUID,
			Type:    PlaceholderCallout,
			Content: calloutStart,
		}
		
		endPlaceholder := p.generatePlaceholder(PlaceholderCallout)
		endUUID := p.extractPlaceholderID(endPlaceholder)
		localMap[endUUID] = Placeholder{
			ID:      endUUID,
			Type:    PlaceholderCallout,
			Content: calloutEnd,
		}
		
		// Handle frontmatter if present
		result := startPlaceholder
		if frontmatter != "" {
			fmPlaceholder := p.generatePlaceholder(PlaceholderCallout)
			fmUUID := p.extractPlaceholderID(fmPlaceholder)
			localMap[fmUUID] = Placeholder{
				ID:      fmUUID,
				Type:    PlaceholderCallout,
				Content: frontmatter,
			}
			result += fmPlaceholder
		}
		
		// Leave the content for translation but process any nested technical content
		result += calloutContent + endPlaceholder
		
		return result
	})
	
	return content, localMap
}

// extractCodeBlocks extracts all code blocks and replaces them with placeholders
func (p *DocumentProcessor) extractCodeBlocks(content string, placeholderMap map[string]Placeholder) (string, map[string]Placeholder) {
	// Match code blocks with optional language and labels
	codeBlockPattern := regexp.MustCompile("(?s)```[^`]*```")
	localMap := make(map[string]Placeholder)
	
	content = codeBlockPattern.ReplaceAllStringFunc(content, func(match string) string {
		placeholder := p.generatePlaceholder(PlaceholderCodeBlock)
		uuid := p.extractPlaceholderID(placeholder)
		localMap[uuid] = Placeholder{
			ID:      uuid,
			Type:    PlaceholderCodeBlock,
			Content: match,
		}
		return placeholder
	})
	
	return content, localMap
}

// extractInlineCode extracts inline code snippets
func (p *DocumentProcessor) extractInlineCode(content string, placeholderMap map[string]Placeholder) (string, map[string]Placeholder) {
	inlineCodePattern := regexp.MustCompile("`([^`]+)`")
	localMap := make(map[string]Placeholder)
	
	content = inlineCodePattern.ReplaceAllStringFunc(content, func(match string) string {
		placeholder := p.generatePlaceholder(PlaceholderInlineCode)
		uuid := p.extractPlaceholderID(placeholder)
		localMap[uuid] = Placeholder{
			ID:      uuid,
			Type:    PlaceholderInlineCode,
			Content: match,
		}
		return placeholder
	})
	
	return content, localMap
}

// processTables handles table parsing with column intelligence
func (p *DocumentProcessor) processTables(content string, placeholderMap map[string]Placeholder) (string, map[string]Placeholder) {
	// Simple table detection regex - match tables including trailing newlines
	tablePattern := regexp.MustCompile(`(?m)(^\|.*\|.*$\n)+`)
	localMap := make(map[string]Placeholder)
	
	content = tablePattern.ReplaceAllStringFunc(content, func(table string) string {
		// Preserve trailing newlines
		trailingNewlines := ""
		for strings.HasSuffix(table, "\n") {
			trailingNewlines += "\n"
			table = table[:len(table)-1]
		}
		
		lines := strings.Split(table, "\n")
		if len(lines) < 2 {
			return table + trailingNewlines
		}
		
		// Parse headers
		headers := p.parseTableRow(lines[0])
		translatableCols := p.classifyColumns(headers)
		
		// Process each row
		var processedLines []string
		for i, line := range lines {
			if i == 0 {
				// Header row - never translate
				processedLines = append(processedLines, line)
			} else if strings.Contains(line, "---") {
				// Separator row
				processedLines = append(processedLines, line)
			} else {
				// Data row
				cells := p.parseTableRow(line)
				var processedCells []string
				
				for j, cell := range cells {
					if p.shouldTranslateCell(j, translatableCols, cell) {
						// This cell should be translated
						processedCells = append(processedCells, cell)
					} else {
						// This cell should be preserved
						placeholder := p.generatePlaceholder(PlaceholderTableCell)
						uuid := p.extractPlaceholderID(placeholder)
						localMap[uuid] = Placeholder{
							ID:      uuid,
							Type:    PlaceholderTableCell,
							Content: cell,
						}
						processedCells = append(processedCells, placeholder)
					}
				}
				
				processedLines = append(processedLines, "| "+strings.Join(processedCells, " | ")+" |")
			}
		}
		
		return strings.Join(processedLines, "\n") + trailingNewlines
	})
	
	return content, localMap
}

// parseTableRow splits a table row into cells
func (p *DocumentProcessor) parseTableRow(row string) []string {
	row = strings.Trim(row, "|")
	cells := strings.Split(row, "|")
	for i := range cells {
		cells[i] = strings.TrimSpace(cells[i])
	}
	return cells
}

// extractURLsAndEmails finds and masks URLs and email addresses
func (p *DocumentProcessor) extractURLsAndEmails(content string, placeholderMap map[string]Placeholder) string {
	// Email pattern
	emailPattern := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	content = emailPattern.ReplaceAllStringFunc(content, func(match string) string {
		placeholder := p.generatePlaceholder(PlaceholderEmail)
		uuid := p.extractPlaceholderID(placeholder)
		placeholderMap[uuid] = Placeholder{
			ID:      uuid,
			Type:    PlaceholderEmail,
			Content: match,
		}
		return placeholder
	})
	
	// URL pattern (http/https)
	urlPattern := regexp.MustCompile(`https?://[^\s\)]+`)
	content = urlPattern.ReplaceAllStringFunc(content, func(match string) string {
		placeholder := p.generatePlaceholder(PlaceholderURL)
		uuid := p.extractPlaceholderID(placeholder)
		placeholderMap[uuid] = Placeholder{
			ID:      uuid,
			Type:    PlaceholderURL,
			Content: match,
		}
		return placeholder
	})
	
	return content
}

// classifyColumns determines which columns should be translated
func (p *DocumentProcessor) classifyColumns(headers []string) []int {
	translatableCols := []int{}
	
	translatableHeaders := map[string]bool{
		"description": true,
		"purpose":     true,
		"notes":       true,
		"explanation": true,
		"details":     true,
		"usage":       true,
		"message":     true,
	}
	
	for i, header := range headers {
		headerLower := strings.ToLower(strings.TrimSpace(header))
		if translatableHeaders[headerLower] {
			translatableCols = append(translatableCols, i)
		}
	}
	
	return translatableCols
}

// shouldTranslateCell determines if a specific cell should be translated
func (p *DocumentProcessor) shouldTranslateCell(colIndex int, translatableCols []int, content string) bool {
	// Check if this column is marked as translatable
	for _, col := range translatableCols {
		if col == colIndex {
			// Special case: if content looks like a technical identifier, don't translate
			if p.isTechnicalIdentifier(content) {
				return false
			}
			return true
		}
	}
	
	// Special handling for "Required" column - only translate Yes/No
	contentLower := strings.ToLower(strings.TrimSpace(content))
	if strings.Contains(contentLower, "yes") || strings.Contains(contentLower, "no") ||
	   strings.Contains(contentLower, "✅") || strings.Contains(contentLower, "❌") {
		return true
	}
	
	return false
}

// isTechnicalIdentifier checks if content looks like a technical identifier
func (p *DocumentProcessor) isTechnicalIdentifier(content string) bool {
	// Check for patterns that indicate technical content
	technicalPatterns := []string{
		`^[A-Z_]+$`,           // ALL_CAPS
		`^[a-z]+[A-Z]`,        // camelCase
		`^\w+_\w+$`,           // snake_case
		`^[A-Z][A-Z0-9_]*$`,   // CONSTANT_CASE
		`\$\{.*\}`,            // Template variables
		`^/.*/$`,              // Regex patterns
	}
	
	for _, pattern := range technicalPatterns {
		if matched, _ := regexp.MatchString(pattern, strings.TrimSpace(content)); matched {
			return true
		}
	}
	
	return false
}

// generatePlaceholder creates a unique placeholder
func (p *DocumentProcessor) generatePlaceholder(placeholderType PlaceholderType) string {
	id := uuid.New().String()
	return fmt.Sprintf("@@%s##%s##%s@@", placeholderType, id, placeholderType)
}

// extractPlaceholderID extracts just the UUID from a full placeholder
func (p *DocumentProcessor) extractPlaceholderID(placeholder string) string {
	// Extract UUID from format @@TYPE##UUID##TYPE@@
	parts := strings.Split(strings.Trim(placeholder, "@"), "##")
	if len(parts) >= 2 {
		return parts[1]
	}
	return ""
}

// ValidateTranslation checks if all placeholders are preserved correctly
func (p *DocumentProcessor) ValidateTranslation(original, translated string) error {
	originalPlaceholders := p.extractPlaceholders(original)
	translatedPlaceholders := p.extractPlaceholders(translated)
	
	// Check count
	if len(originalPlaceholders) != len(translatedPlaceholders) {
		return fmt.Errorf("placeholder count mismatch: original=%d, translated=%d", 
			len(originalPlaceholders), len(translatedPlaceholders))
	}
	
	// Check exact matches
	for _, placeholder := range originalPlaceholders {
		found := false
		for _, transPlaceholder := range translatedPlaceholders {
			if placeholder == transPlaceholder {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("missing or modified placeholder: %s", placeholder)
		}
	}
	
	return nil
}

// extractPlaceholders finds all placeholders in content
func (p *DocumentProcessor) extractPlaceholders(content string) []string {
	placeholderPattern := regexp.MustCompile(`@@[A-Z]+##[a-f0-9-]+##[A-Z]+@@`)
	return placeholderPattern.FindAllString(content, -1)
}

// RestoreContent replaces all placeholders with their original content
func (p *DocumentProcessor) RestoreContent(content string, placeholderMap map[string]Placeholder) string {
	result := content
	
	// Restore placeholders by reconstructing the full placeholder format
	for uuid, data := range placeholderMap {
		fullPlaceholder := fmt.Sprintf("@@%s##%s##%s@@", data.Type, uuid, data.Type)
		result = strings.ReplaceAll(result, fullPlaceholder, data.Content)
	}
	
	return result
}

// RecoverPlaceholders attempts to recover corrupted placeholders
func (p *DocumentProcessor) RecoverPlaceholders(original, corrupted string) string {
	// Extract all original placeholders
	originalPlaceholders := p.extractPlaceholders(original)
	
	// Try to find and fix corrupted placeholders
	result := corrupted
	
	for _, placeholder := range originalPlaceholders {
		if !strings.Contains(result, placeholder) {
			// Try to find partial matches or translations
			parts := strings.Split(placeholder, "##")
			if len(parts) == 3 {
				id := parts[1]
				// Look for the ID in various forms
				patterns := []string{
					fmt.Sprintf("@@%s##%s##%s@@", ".*", id, ".*"),  // Any type with same ID
					fmt.Sprintf("@@.*##%s##.*@@", id),              // More flexible
					id,                                               // Just the ID
				}
				
				for _, pattern := range patterns {
					re := regexp.MustCompile(pattern)
					if matches := re.FindAllString(result, -1); len(matches) > 0 {
						// Replace the corrupted version with the correct one
						result = strings.Replace(result, matches[0], placeholder, 1)
						break
					}
				}
			}
		}
	}
	
	return result
}