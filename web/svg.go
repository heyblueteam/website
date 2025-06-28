package web

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"strings"
)


// Color palette for the art generation
var colorPalette = []string{
	"#00a0d2", // Primary blue
	"#F06C9B", // Pink
	"#FED766", // Yellow
	"#A9F0D1", // Mint green
	"#F86624", // Orange
}

// SVGGenerator generates deterministic abstract art SVGs
type SVGGenerator struct {
	width  int
	height int
}

// NewSVGGenerator creates a new SVG generator
func NewSVGGenerator() *SVGGenerator {
	return &SVGGenerator{
		width:  400,
		height: 300,
	}
}

// GenerateInsightSVG creates a unique SVG based on the insight title
func (sg *SVGGenerator) GenerateInsightSVG(title string) string {
	// Create deterministic seed from title
	hash := sha256.Sum256([]byte(title))
	seed := int64(binary.BigEndian.Uint64(hash[:8]))
	
	// Create seeded random generator for consistent results
	r := rand.New(rand.NewSource(seed))
	
	// Select color palette for this artwork (2-3 colors)
	colors := sg.selectColors(r)
	
	// Build SVG content
	var svg strings.Builder
	
	// SVG header
	svg.WriteString(fmt.Sprintf(`<svg width="%d" height="%d" viewBox="0 0 %d %d" xmlns="http://www.w3.org/2000/svg">`,
		sg.width, sg.height, sg.width, sg.height))
	
	// Background (more opaque to ensure full coverage)
	bgOpacity := 0.1 + r.Float64()*0.1
	svg.WriteString(fmt.Sprintf(`<rect width="100%%" height="100%%" fill="%s" opacity="%.3f"/>`,
		colors[0], bgOpacity))
	
	// Generate soft gradient art
	sg.generateSoftGradientStyle(&svg, r, colors)
	
	// Close SVG
	svg.WriteString("</svg>")
	
	return svg.String()
}

// selectColors picks 2-3 colors from the palette for this artwork
func (sg *SVGGenerator) selectColors(r *rand.Rand) []string {
	// Always include primary blue
	selected := []string{colorPalette[0]}
	
	// Add 1-2 additional colors
	numAdditional := 1 + r.Intn(2) // 1 or 2 additional colors
	
	for i := 0; i < numAdditional; i++ {
		// Pick from remaining colors (skip primary blue at index 0)
		colorIndex := 1 + r.Intn(len(colorPalette)-1)
		color := colorPalette[colorIndex]
		
		// Avoid duplicates
		duplicate := false
		for _, existing := range selected {
			if existing == color {
				duplicate = true
				break
			}
		}
		
		if !duplicate {
			selected = append(selected, color)
		}
	}
	
	return selected
}

// generateSoftGradientStyle creates beautiful soft gradient blends with organic shapes
func (sg *SVGGenerator) generateSoftGradientStyle(svg *strings.Builder, r *rand.Rand, colors []string) {
	// Create multiple overlapping soft gradients for beautiful color blending
	numGradients := 4 + r.Intn(3) // 4-6 gradients for rich blending
	
	for i := 0; i < numGradients; i++ {
		// Random center position (extend beyond canvas for full coverage)
		cx := -40 + r.Float64()*180 // -40% to 140%
		cy := -40 + r.Float64()*180 // -40% to 140%
		
		// Pick colors for this gradient
		color1 := colors[r.Intn(len(colors))]
		color2 := colors[r.Intn(len(colors))]
		
		// Create large, soft radial gradient
		gradientId := fmt.Sprintf("gradient%d", i)
		svg.WriteString(fmt.Sprintf(`<defs><radialGradient id="%s" cx="%.1f%%" cy="%.1f%%" r="90%%">`, 
			gradientId, cx, cy))
		
		// Multiple color stops for smooth transitions
		svg.WriteString(fmt.Sprintf(`<stop offset="0%%" stop-color="%s" stop-opacity="0.25"/>`, color1))
		svg.WriteString(fmt.Sprintf(`<stop offset="30%%" stop-color="%s" stop-opacity="0.15"/>`, color2))
		svg.WriteString(fmt.Sprintf(`<stop offset="60%%" stop-color="%s" stop-opacity="0.08"/>`, color1))
		svg.WriteString(`<stop offset="100%" stop-color="transparent"/>`)
		svg.WriteString(`</radialGradient></defs>`)
		
		// Apply gradient across extended canvas with minimal blur for soft edges
		svg.WriteString(fmt.Sprintf(`<rect x="-50%%" y="-50%%" width="200%%" height="200%%" fill="url(#%s)" filter="blur(0.5px)"/>`,
			gradientId))
	}
	
	// Add a few linear gradients for additional flow
	numLinear := 2 + r.Intn(2) // 2-3 linear gradients
	for i := 0; i < numLinear; i++ {
		// Random angle for gradient direction
		angle := r.Float64() * 360
		
		color1 := colors[r.Intn(len(colors))]
		color2 := colors[r.Intn(len(colors))]
		
		gradientId := fmt.Sprintf("linear%d", i)
		svg.WriteString(fmt.Sprintf(`<defs><linearGradient id="%s" gradientTransform="rotate(%.1f)">`, 
			gradientId, angle))
		svg.WriteString(fmt.Sprintf(`<stop offset="0%%" stop-color="%s" stop-opacity="0.15"/>`, color1))
		svg.WriteString(fmt.Sprintf(`<stop offset="50%%" stop-color="%s" stop-opacity="0.1"/>`, color2))
		svg.WriteString(`<stop offset="100%" stop-color="transparent"/>`)
		svg.WriteString(`</linearGradient></defs>`)
		
		// Apply across full extended canvas
		svg.WriteString(fmt.Sprintf(`<rect x="-50%%" y="-50%%" width="200%%" height="200%%" fill="url(#%s)" filter="blur(0.5px)"/>`,
			gradientId))
	}

	// Add subtle organic shapes
	sg.addOrganicShapes(svg, r, colors)
}

// addOrganicShapes adds subtle organic blob shapes to the composition
func (sg *SVGGenerator) addOrganicShapes(svg *strings.Builder, r *rand.Rand, colors []string) {
	numShapes := 2 + r.Intn(3) // 2-4 organic shapes
	
	for i := 0; i < numShapes; i++ {
		// Random position across the canvas
		centerX := r.Float64() * float64(sg.width)
		centerY := r.Float64() * float64(sg.height)
		
		// Random size (fairly large but subtle)
		baseRadius := 80 + r.Float64()*120 // 80-200px radius
		
		// Generate organic blob shape using multiple control points
		var pathData strings.Builder
		pathData.WriteString(fmt.Sprintf("M %.1f %.1f", centerX, centerY-baseRadius))
		
		// Create smooth organic shape with 6-8 control points
		numPoints := 6 + r.Intn(3)
		angleStep := 360.0 / float64(numPoints)
		
		for j := 0; j < numPoints; j++ {
			angle := float64(j) * angleStep * math.Pi / 180
			
			// Add some randomness to radius for organic feel
			radiusVariation := 0.7 + r.Float64()*0.6 // 0.7-1.3 multiplier
			currentRadius := baseRadius * radiusVariation
			
			x := centerX + currentRadius*math.Cos(angle)
			y := centerY + currentRadius*math.Sin(angle)
			
			// Create smooth curves using quadratic bezier
			if j == 0 {
				pathData.WriteString(fmt.Sprintf(" Q %.1f %.1f %.1f %.1f", x, y, x, y))
			} else {
				// Control point for smooth curve
				prevAngle := float64(j-1) * angleStep * math.Pi / 180
				prevRadius := baseRadius * (0.7 + r.Float64()*0.6)
				cpX := centerX + prevRadius*0.8*math.Cos(prevAngle+angleStep*math.Pi/360)
				cpY := centerY + prevRadius*0.8*math.Sin(prevAngle+angleStep*math.Pi/360)
				
				pathData.WriteString(fmt.Sprintf(" Q %.1f %.1f %.1f %.1f", cpX, cpY, x, y))
			}
		}
		pathData.WriteString(" Z") // Close the path
		
		// Pick a color and create gradient for this shape
		shapeColor := colors[r.Intn(len(colors))]
		shapeGradientId := fmt.Sprintf("shapeGrad%d", i)
		
		// Create radial gradient for the organic shape
		svg.WriteString(fmt.Sprintf(`<defs><radialGradient id="%s" cx="50%%" cy="50%%" r="70%%">`, shapeGradientId))
		svg.WriteString(fmt.Sprintf(`<stop offset="0%%" stop-color="%s" stop-opacity="0.08"/>`, shapeColor))
		svg.WriteString(fmt.Sprintf(`<stop offset="70%%" stop-color="%s" stop-opacity="0.04"/>`, shapeColor))
		svg.WriteString(`<stop offset="100%" stop-color="transparent"/>`)
		svg.WriteString(`</radialGradient></defs>`)
		
		// Draw the organic shape
		svg.WriteString(fmt.Sprintf(`<path d="%s" fill="url(#%s)" filter="blur(1px)"/>`, 
			pathData.String(), shapeGradientId))
	}
}


// GenerateSVGDataURL creates a complete data URL for the SVG
func (sg *SVGGenerator) GenerateSVGDataURL(title string) string {
	svgContent := sg.GenerateInsightSVG(title)
	encoded := base64.StdEncoding.EncodeToString([]byte(svgContent))
	return fmt.Sprintf("data:image/svg+xml;base64,%s", encoded)
}