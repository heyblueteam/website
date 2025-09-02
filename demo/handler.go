package demo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// Handler manages demo request submissions
type Handler struct {
	rateLimiter  *RateLimiter
	emailService *EmailService
	config       *Config
}

// NewHandler creates a new demo handler instance
func NewHandler() *Handler {
	config := &Config{
		EmailItAPIKey:     os.Getenv("EMAILIT_API_KEY"),
		EmailItFromEmail:  os.Getenv("EMAILIT_FROM_EMAIL"),
		EmailItFromName:   os.Getenv("EMAILIT_FROM_NAME"),
		NotificationEmail: os.Getenv("NOTIFICATION_EMAIL"),
	}
	
	// Set defaults if not provided
	if config.EmailItFromEmail == "" {
		config.EmailItFromEmail = "enterprise@blue.cc"
	}
	if config.EmailItFromName == "" {
		config.EmailItFromName = "Blue Enterprise Team"
	}
	if config.NotificationEmail == "" {
		config.NotificationEmail = "manny@blue.cc"
	}
	
	return &Handler{
		rateLimiter:  NewRateLimiter(),
		emailService: NewEmailService(config),
		config:       config,
	}
}

// Handle processes demo form submissions
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	// Enable CORS for the endpoint
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	// Handle preflight requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	
	// Only accept POST requests
	if r.Method != "POST" {
		h.sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Parse JSON body
	var req DemoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	// Validate and sanitize the request
	if err := ValidateRequest(&req); err != nil {
		switch err {
		case ErrBotDetected:
			// Silently accept but don't process (to not tip off bots)
			h.sendSuccess(w)
			log.Printf("Bot detected from IP: %s", h.getClientIP(r))
			return
		case ErrInvalidEmail:
			h.sendError(w, "Please enter a valid email address", http.StatusBadRequest)
			return
		case ErrMissingRequired:
			h.sendError(w, "Please fill in all required fields", http.StatusBadRequest)
			return
		default:
			h.sendError(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	
	// Get client IP for rate limiting
	clientIP := h.getClientIP(r)
	
	// Check rate limit
	if !h.rateLimiter.Allow(clientIP) {
		h.sendError(w, "Too many requests. Please try again later.", http.StatusTooManyRequests)
		log.Printf("Rate limit exceeded for IP: %s", clientIP)
		return
	}
	
	// Send emails (with retry logic)
	if err := h.emailService.SendDemoEmails(&req, clientIP); err != nil {
		log.Printf("Failed to send emails: %v", err)
		// Still return success to user to avoid revealing email issues
		h.sendSuccess(w)
		return
	}
	
	// Log successful submission
	log.Printf("Demo request received from %s at %s (IP: %s)", req.Company, req.Email, clientIP)
	
	// Send success response
	h.sendSuccess(w)
}

// getClientIP extracts the client IP address from the request
func (h *Handler) getClientIP(r *http.Request) string {
	// Check X-Forwarded-For header (for proxies/load balancers)
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// Take the first IP if there are multiple
		parts := strings.Split(forwarded, ",")
		if len(parts) > 0 {
			return strings.TrimSpace(parts[0])
		}
	}
	
	// Check X-Real-IP header
	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}
	
	// Fall back to RemoteAddr
	addr := r.RemoteAddr
	// Remove port if present
	if idx := strings.LastIndex(addr, ":"); idx != -1 {
		addr = addr[:idx]
	}
	
	return addr
}

// sendSuccess sends a success response
func (h *Handler) sendSuccess(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Thank you for your interest! We'll contact you within 24 hours.",
	})
}

// sendError sends an error response
func (h *Handler) sendError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": false,
		"error":   message,
	})
}

// GetStats returns handler statistics (useful for monitoring)
func (h *Handler) GetStats() map[string]interface{} {
	ipCount, totalCount := h.rateLimiter.GetStats()
	return map[string]interface{}{
		"rate_limit": map[string]interface{}{
			"unique_ips":     ipCount,
			"total_requests": totalCount,
			"max_per_ip":     MaxRequestsPerIPPerHour,
			"max_total":      MaxTotalRequestsPerHour,
		},
		"config": map[string]interface{}{
			"from_email":        h.config.EmailItFromEmail,
			"from_name":         h.config.EmailItFromName,
			"notification_email": h.config.NotificationEmail,
			"api_key_set":       h.config.EmailItAPIKey != "",
		},
	}
}

// ValidateConfig checks if the handler is properly configured
func (h *Handler) ValidateConfig() error {
	if h.config.EmailItAPIKey == "" {
		return fmt.Errorf("EMAILIT_API_KEY environment variable is not set")
	}
	return nil
}