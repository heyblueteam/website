package demo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandleDemo_HoneypotDetection(t *testing.T) {
	handler := NewHandler()
	
	// Test request with honeypot field filled (bot)
	botRequest := DemoRequest{
		FullName:    "Test User",
		Email:       "test@example.com",
		Company:     "Test Corp",
		JobTitle:    "CEO",
		CompanySize: "50-250",
		UseCase:     "process-automation",
		URL:         "http://spam.com", // Honeypot field filled
	}
	
	body, _ := json.Marshal(botRequest)
	req := httptest.NewRequest("POST", "/api/demo-request", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	handler.Handle(w, req)
	
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 for bot detection, got %d", w.Code)
	}
}

func TestHandleDemo_ValidRequest(t *testing.T) {
	// Skip if EmailIt API key not set
	if os.Getenv("EMAILIT_API_KEY") == "" && os.Getenv("EMAILIT_API_KEY") == "your-emailit-api-key-here" {
		t.Skip("EMAILIT_API_KEY not configured")
	}
	
	handler := NewHandler()
	
	// Valid request
	validRequest := DemoRequest{
		FullName:    "John Doe",
		Email:       "john@example.com",
		Company:     "Example Corp",
		JobTitle:    "VP Engineering",
		CompanySize: "250-1000",
		UseCase:     "project-management",
		Message:     "We need better project tracking",
		URL:         "", // Honeypot field empty (human)
	}
	
	body, _ := json.Marshal(validRequest)
	req := httptest.NewRequest("POST", "/api/demo-request", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	handler.Handle(w, req)
	
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
	
	var response map[string]interface{}
	json.NewDecoder(w.Body).Decode(&response)
	
	if !response["success"].(bool) {
		t.Error("Expected success to be true")
	}
}

func TestHandleDemo_MissingRequiredFields(t *testing.T) {
	handler := NewHandler()
	
	// Request missing required fields
	invalidRequest := DemoRequest{
		FullName: "Test User",
		// Missing other required fields
	}
	
	body, _ := json.Marshal(invalidRequest)
	req := httptest.NewRequest("POST", "/api/demo-request", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	handler.Handle(w, req)
	
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for missing fields, got %d", w.Code)
	}
}

func TestHandleDemo_InvalidEmail(t *testing.T) {
	handler := NewHandler()
	
	// Request with invalid email
	invalidRequest := DemoRequest{
		FullName:    "Test User",
		Email:       "not-an-email",
		Company:     "Test Corp",
		JobTitle:    "CEO",
		CompanySize: "50-250",
		UseCase:     "process-automation",
	}
	
	body, _ := json.Marshal(invalidRequest)
	req := httptest.NewRequest("POST", "/api/demo-request", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	handler.Handle(w, req)
	
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for invalid email, got %d", w.Code)
	}
}