package demo

import "time"

// DemoRequest represents a demo form submission
type DemoRequest struct {
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	Company     string `json:"company"`
	JobTitle    string `json:"jobTitle"`
	Phone       string `json:"phone"`
	CompanySize string `json:"companySize"`
	UseCase     string `json:"useCase"`
	Message     string `json:"message"`
	// Honeypot field - if filled, it's likely a bot
	URL string `json:"url"`
}

// EmailPayload represents an email to be sent via EmailIt
type EmailPayload struct {
	From    string                 `json:"from"`
	To      string                 `json:"to"`
	ReplyTo string                 `json:"reply_to,omitempty"`
	Subject string                 `json:"subject"`
	HTML    string                 `json:"html"`
	Text    string                 `json:"text,omitempty"`
	Headers map[string]string      `json:"headers,omitempty"`
}

// Config holds configuration for the demo package
type Config struct {
	EmailItAPIKey   string
	EmailItFromEmail string
	EmailItFromName  string
	NotificationEmail string
}

// RateLimitEntry tracks request timestamps
type RateLimitEntry struct {
	Timestamps []time.Time
}