package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

// Test D1Client creation
func TestNewD1Client(t *testing.T) {
	// Save original env values
	origAccountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	origDatabaseID := os.Getenv("CLOUDFLARE_DATABASE_ID")
	origAPIKey := os.Getenv("CLOUDFLARE_API_KEY")

	// Set test values
	os.Setenv("CLOUDFLARE_ACCOUNT_ID", "test-account-id")
	os.Setenv("CLOUDFLARE_DATABASE_ID", "test-database-id")
	os.Setenv("CLOUDFLARE_API_KEY", "test-api-key")

	// Restore original values after test
	defer func() {
		os.Setenv("CLOUDFLARE_ACCOUNT_ID", origAccountID)
		os.Setenv("CLOUDFLARE_DATABASE_ID", origDatabaseID)
		os.Setenv("CLOUDFLARE_API_KEY", origAPIKey)
	}()

	client := NewD1Client()

	if client.AccountID != "test-account-id" {
		t.Errorf("Expected AccountID %q, got %q", "test-account-id", client.AccountID)
	}
	if client.DatabaseID != "test-database-id" {
		t.Errorf("Expected DatabaseID %q, got %q", "test-database-id", client.DatabaseID)
	}
	if client.APIKey != "test-api-key" {
		t.Errorf("Expected APIKey %q, got %q", "test-api-key", client.APIKey)
	}
	if client.BaseURL != "https://api.cloudflare.com/client/v4" {
		t.Errorf("Expected BaseURL %q, got %q", "https://api.cloudflare.com/client/v4", client.BaseURL)
	}
}

// Test D1Client Query method
func TestD1ClientQuery(t *testing.T) {
	tests := []struct {
		name           string
		sql            string
		params         []string
		mockResponse   D1Response
		mockStatusCode int
		expectError    bool
		errorContains  string
	}{
		{
			name:   "successful query",
			sql:    "SELECT * FROM test",
			params: []string{},
			mockResponse: D1Response{
				Success: true,
				Result: []struct {
					Meta struct {
						ChangedDB      bool    `json:"changed_db"`
						Changes        int     `json:"changes"`
						Duration       float64 `json:"duration"`
						LastRowID      int     `json:"last_row_id"`
						RowsRead       int     `json:"rows_read"`
						RowsWritten    int     `json:"rows_written"`
						ServedByRegion string  `json:"served_by_region"`
					} `json:"meta"`
					Results []map[string]interface{} `json:"results"`
					Success bool                     `json:"success"`
				}{
					{
						Results: []map[string]interface{}{
							{"id": 1, "name": "test"},
						},
						Success: true,
					},
				},
			},
			mockStatusCode: http.StatusOK,
			expectError:    false,
		},
		{
			name:   "query with error response",
			sql:    "SELECT * FROM nonexistent",
			params: []string{},
			mockResponse: D1Response{
				Success: false,
				Errors: []struct {
					Code    int    `json:"code"`
					Message string `json:"message"`
				}{
					{Code: 400, Message: "table not found"},
				},
			},
			mockStatusCode: http.StatusBadRequest,
			expectError:    true,
			errorContains:  "table not found",
		},
		{
			name:   "query with generic failure",
			sql:    "SELECT * FROM test",
			params: []string{},
			mockResponse: D1Response{
				Success: false,
			},
			mockStatusCode: http.StatusBadRequest,
			expectError:    true,
			errorContains:  "D1 request failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Verify request method
				if r.Method != "POST" {
					t.Errorf("Expected POST method, got %s", r.Method)
				}

				// Verify headers
				if r.Header.Get("Content-Type") != "application/json" {
					t.Errorf("Expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
				}
				if r.Header.Get("Authorization") != "Bearer test-key" {
					t.Errorf("Expected Authorization Bearer test-key, got %s", r.Header.Get("Authorization"))
				}

				// Verify request body
				var reqBody D1Request
				if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
					t.Fatalf("Failed to decode request body: %v", err)
				}

				if reqBody.SQL != tt.sql {
					t.Errorf("Expected SQL %q, got %q", tt.sql, reqBody.SQL)
				}

				// Send response
				w.WriteHeader(tt.mockStatusCode)
				json.NewEncoder(w).Encode(tt.mockResponse)
			}))
			defer server.Close()

			// Create client with test server URL
			client := &D1Client{
				AccountID:  "test-account",
				DatabaseID: "test-db",
				APIKey:     "test-key",
				BaseURL:    server.URL,
			}

			// Execute query
			resp, err := client.Query(tt.sql, tt.params...)

			// Check error
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				} else if tt.errorContains != "" && !contains(err.Error(), tt.errorContains) {
					t.Errorf("Expected error to contain %q, got %q", tt.errorContains, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if resp == nil {
					t.Errorf("Expected response, got nil")
				} else if resp.Success != tt.mockResponse.Success {
					t.Errorf("Expected success %v, got %v", tt.mockResponse.Success, resp.Success)
				}
			}
		})
	}
}

// Test HealthChecker creation
func TestNewHealthChecker(t *testing.T) {
	d1Client := &D1Client{}
	checker := NewHealthChecker(d1Client)

	if checker.d1Client != d1Client {
		t.Errorf("Expected d1Client to be set")
	}
	if checker.cache == nil {
		t.Errorf("Expected cache to be initialized")
	}
	if checker.lastCheckTime == nil {
		t.Errorf("Expected lastCheckTime to be initialized")
	}
}

// Test HealthChecker initialization
func TestHealthCheckerInitialize(t *testing.T) {
	// Create mock D1 client
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request
		var reqBody D1Request
		json.NewDecoder(r.Body).Decode(&reqBody)

		// Mock different responses based on SQL
		response := D1Response{Success: true}

		if contains(reqBody.SQL, "CREATE TABLE") || contains(reqBody.SQL, "CREATE INDEX") {
			// Schema creation
			response.Result = []struct {
				Meta struct {
					ChangedDB      bool    `json:"changed_db"`
					Changes        int     `json:"changes"`
					Duration       float64 `json:"duration"`
					LastRowID      int     `json:"last_row_id"`
					RowsRead       int     `json:"rows_read"`
					RowsWritten    int     `json:"rows_written"`
					ServedByRegion string  `json:"served_by_region"`
				} `json:"meta"`
				Results []map[string]interface{} `json:"results"`
				Success bool                     `json:"success"`
			}{
				{Success: true},
			}
		} else if contains(reqBody.SQL, "DELETE FROM") {
			// Cleanup
			response.Result = []struct {
				Meta struct {
					ChangedDB      bool    `json:"changed_db"`
					Changes        int     `json:"changes"`
					Duration       float64 `json:"duration"`
					LastRowID      int     `json:"last_row_id"`
					RowsRead       int     `json:"rows_read"`
					RowsWritten    int     `json:"rows_written"`
					ServedByRegion string  `json:"served_by_region"`
				} `json:"meta"`
				Results []map[string]interface{} `json:"results"`
				Success bool                     `json:"success"`
			}{
				{
					Meta: struct {
						ChangedDB      bool    `json:"changed_db"`
						Changes        int     `json:"changes"`
						Duration       float64 `json:"duration"`
						LastRowID      int     `json:"last_row_id"`
						RowsRead       int     `json:"rows_read"`
						RowsWritten    int     `json:"rows_written"`
						ServedByRegion string  `json:"served_by_region"`
					}{
						ChangedDB: true,
						Changes:   5,
					},
					Success: true,
				},
			}
		} else if contains(reqBody.SQL, "SELECT service_name") {
			// Load historical data
			response.Result = []struct {
				Meta struct {
					ChangedDB      bool    `json:"changed_db"`
					Changes        int     `json:"changes"`
					Duration       float64 `json:"duration"`
					LastRowID      int     `json:"last_row_id"`
					RowsRead       int     `json:"rows_read"`
					RowsWritten    int     `json:"rows_written"`
					ServedByRegion string  `json:"served_by_region"`
				} `json:"meta"`
				Results []map[string]interface{} `json:"results"`
				Success bool                     `json:"success"`
			}{
				{
					Results: []map[string]interface{}{
						{
							"service_name": "API",
							"status":       "up",
							"checked_at":   time.Now().Format(time.RFC3339),
						},
						{
							"service_name": "Website",
							"status":       "down",
							"checked_at":   time.Now().Add(-1 * time.Hour).Format(time.RFC3339),
						},
					},
					Success: true,
				},
			}
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer mockServer.Close()

	d1Client := &D1Client{
		BaseURL:    mockServer.URL,
		AccountID:  "test",
		DatabaseID: "test",
		APIKey:     "test",
	}

	checker := NewHealthChecker(d1Client)
	err := checker.Initialize()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify cache was populated
	cacheCount := 0
	checker.cache.Range(func(key, value interface{}) bool {
		cacheCount++
		return true
	})

	if cacheCount == 0 {
		t.Errorf("Expected cache to be populated, but it's empty")
	}
}

// Test CheckService method
func TestCheckService(t *testing.T) {
	tests := []struct {
		name           string
		service        Service
		serverResponse int
		expectStatus   string
	}{
		{
			name: "service up with 200 response",
			service: Service{
				Name:           "Test Service",
				URL:            "https://example.com",
				HealthEndpoint: "",
			},
			serverResponse: http.StatusOK,
			expectStatus:   "up",
		},
		{
			name: "service down with 500 response",
			service: Service{
				Name:           "Test Service",
				URL:            "https://example.com",
				HealthEndpoint: "",
			},
			serverResponse: http.StatusInternalServerError,
			expectStatus:   "down",
		},
		{
			name: "service with health endpoint",
			service: Service{
				Name:           "Test Service",
				URL:            "https://example.com",
				HealthEndpoint: "https://example.com/health",
			},
			serverResponse: http.StatusOK,
			expectStatus:   "up",
		},
		{
			name: "white label files special case",
			service: Service{
				Name:           "White Label Files",
				URL:            "https://wl-files.example.com",
				HealthEndpoint: "",
			},
			serverResponse: http.StatusNotFound,
			expectStatus:   "up",
		},
		{
			name: "service down with 404 response",
			service: Service{
				Name:           "Regular Service",
				URL:            "https://example.com",
				HealthEndpoint: "",
			},
			serverResponse: http.StatusNotFound,
			expectStatus:   "down",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test server for the service
			serviceServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.serverResponse)
			}))
			defer serviceServer.Close()

			// Update service URLs to point to test server
			testService := tt.service
			testService.URL = serviceServer.URL
			if testService.HealthEndpoint != "" {
				testService.HealthEndpoint = serviceServer.URL + "/health"
			}

			// Create mock D1 server
			d1Server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				response := D1Response{Success: true}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(response)
			}))
			defer d1Server.Close()

			d1Client := &D1Client{
				BaseURL:    d1Server.URL,
				AccountID:  "test",
				DatabaseID: "test",
				APIKey:     "test",
			}

			checker := NewHealthChecker(d1Client)
			result := checker.CheckService(testService)

			if result.Status != tt.expectStatus {
				t.Errorf("Expected status %q, got %q", tt.expectStatus, result.Status)
			}
			if result.ServiceName != testService.Name {
				t.Errorf("Expected service name %q, got %q", testService.Name, result.ServiceName)
			}
			if result.ServiceURL != testService.URL {
				t.Errorf("Expected service URL %q, got %q", testService.URL, result.ServiceURL)
			}
		})
	}
}

// Test cache operations
func TestAddToCache(t *testing.T) {
	checker := &HealthChecker{
		cache:         &sync.Map{},
		lastCheckTime: &sync.Map{},
	}

	// Add multiple check results
	now := time.Now()
	results := []CheckResult{
		{
			ServiceName: "API",
			Status:      "up",
			CheckedAt:   now,
		},
		{
			ServiceName: "API",
			Status:      "down",
			CheckedAt:   now.Add(5 * time.Minute),
		},
		{
			ServiceName: "Website",
			Status:      "up",
			CheckedAt:   now,
		},
	}

	for _, result := range results {
		checker.addToCache(result)
	}

	// Verify cache contents
	apiKey := fmt.Sprintf("API:%s", now.Format("2006-01-02"))
	if cached, ok := checker.cache.Load(apiKey); ok {
		checks := cached.([]CheckResult)
		if len(checks) != 2 {
			t.Errorf("Expected 2 checks for API, got %d", len(checks))
		}
	} else {
		t.Errorf("Expected API checks in cache")
	}

	// Verify last check time
	if lastCheck, ok := checker.lastCheckTime.Load("API"); ok {
		checkTime := lastCheck.(time.Time)
		expectedTime := now.Add(5 * time.Minute)
		if !checkTime.Equal(expectedTime) {
			t.Errorf("Expected last check time %v, got %v", expectedTime, checkTime)
		}
	} else {
		t.Errorf("Expected last check time for API")
	}
}

// Test GetCurrentStatus
func TestGetCurrentStatus(t *testing.T) {
	checker := &HealthChecker{
		cache:         &sync.Map{},
		lastCheckTime: &sync.Map{},
	}

	// Add test data
	now := time.Now()
	checker.addToCache(CheckResult{
		ServiceName: "Website",
		Status:      "up",
		CheckedAt:   now.Add(-2 * time.Minute),
	})
	checker.addToCache(CheckResult{
		ServiceName: "API",
		Status:      "down",
		CheckedAt:   now.Add(-30 * time.Minute),
	})

	// Get current status
	statuses := checker.GetCurrentStatus()

	// Verify we have status for all monitored services
	if len(statuses) != len(monitoredServices) {
		t.Errorf("Expected %d statuses, got %d", len(monitoredServices), len(statuses))
	}

	// Check specific statuses
	for _, status := range statuses {
		switch status.Name {
		case "Website":
			if status.Status != "up" {
				t.Errorf("Expected Website status 'up', got %q", status.Status)
			}
			if status.LastChecked == "Never" {
				t.Errorf("Expected Website to have last check time")
			}
		case "API":
			if status.Status != "down" {
				t.Errorf("Expected API status 'down', got %q", status.Status)
			}
		}
	}
}

// Test GetHistoricalData
func TestGetHistoricalData(t *testing.T) {
	checker := &HealthChecker{
		cache:         &sync.Map{},
		lastCheckTime: &sync.Map{},
	}

	// Add test data for multiple days
	now := time.Now()
	for i := 0; i < 5; i++ {
		date := now.AddDate(0, 0, -i)
		// Add multiple checks per day
		for j := 0; j < 10; j++ {
			status := "up"
			if j < 2 { // 20% down
				status = "down"
			}
			checker.addToCache(CheckResult{
				ServiceName: "API",
				Status:      status,
				CheckedAt:   date.Add(time.Duration(j) * time.Minute),
			})
		}
	}

	histories := checker.GetHistoricalData()

	// Verify we have history for all services
	if len(histories) != len(monitoredServices) {
		t.Errorf("Expected %d histories, got %d", len(monitoredServices), len(histories))
	}

	// Check API history
	for _, history := range histories {
		if history.Name == "API" {
			// Should have 90 days of data
			if len(history.Days) != 90 {
				t.Errorf("Expected 90 days of data, got %d", len(history.Days))
			}

			// Check recent days have correct uptime
			for i := 0; i < 5; i++ {
				day := history.Days[89-i]
				if day.Uptime != 80.0 {
					t.Errorf("Expected day %d uptime 80%%, got %.2f%%", i, day.Uptime)
				}
				if day.Status != "outage" {
					t.Errorf("Expected day %d status 'outage', got %q", i, day.Status)
				}
			}
		}
	}
}

// Test hasRecentCheck
func TestHasRecentCheck(t *testing.T) {
	tests := []struct {
		name         string
		queryResult  []map[string]interface{}
		expectRecent bool
	}{
		{
			name: "has recent check",
			queryResult: []map[string]interface{}{
				{"count": float64(5)},
			},
			expectRecent: true,
		},
		{
			name: "no recent check",
			queryResult: []map[string]interface{}{
				{"count": float64(0)},
			},
			expectRecent: false,
		},
		{
			name:         "empty result",
			queryResult:  []map[string]interface{}{},
			expectRecent: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				response := D1Response{
					Success: true,
					Result: []struct {
						Meta struct {
							ChangedDB      bool    `json:"changed_db"`
							Changes        int     `json:"changes"`
							Duration       float64 `json:"duration"`
							LastRowID      int     `json:"last_row_id"`
							RowsRead       int     `json:"rows_read"`
							RowsWritten    int     `json:"rows_written"`
							ServedByRegion string  `json:"served_by_region"`
						} `json:"meta"`
						Results []map[string]interface{} `json:"results"`
						Success bool                     `json:"success"`
					}{
						{
							Results: tt.queryResult,
							Success: true,
						},
					},
				}

				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(response)
			}))
			defer mockServer.Close()

			d1Client := &D1Client{
				BaseURL:    mockServer.URL,
				AccountID:  "test",
				DatabaseID: "test",
				APIKey:     "test",
			}

			checker := NewHealthChecker(d1Client)
			result := checker.hasRecentCheck()

			if result != tt.expectRecent {
				t.Errorf("Expected hasRecentCheck %v, got %v", tt.expectRecent, result)
			}
		})
	}
}


// Test formatDuration helper
func TestFormatDuration(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		expected string
	}{
		{
			name:     "less than a minute",
			duration: 30 * time.Second,
			expected: "Less than a minute ago",
		},
		{
			name:     "1 minute",
			duration: 1 * time.Minute,
			expected: "1 minute ago",
		},
		{
			name:     "multiple minutes",
			duration: 15 * time.Minute,
			expected: "15 minutes ago",
		},
		{
			name:     "1 hour",
			duration: 1 * time.Hour,
			expected: "1 hour ago",
		},
		{
			name:     "multiple hours",
			duration: 5 * time.Hour,
			expected: "5 hours ago",
		},
		{
			name:     "more than a day",
			duration: 36 * time.Hour,
			expected: "More than a day ago",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatDuration(tt.duration)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

// Test getDayDataForService
func TestGetDayDataForService(t *testing.T) {
	checker := &HealthChecker{
		cache:         &sync.Map{},
		lastCheckTime: &sync.Map{},
	}

	date := time.Now()

	// Test with no data - should assume operational
	dayData := checker.getDayDataForService("API", date)
	if dayData.Status != "operational" {
		t.Errorf("Expected status 'operational' for no data, got %q", dayData.Status)
	}
	if dayData.Uptime != 100.0 {
		t.Errorf("Expected uptime 100%% for no data, got %.2f%%", dayData.Uptime)
	}

	// Add some checks with mixed status
	for i := 0; i < 100; i++ {
		status := "up"
		if i < 5 { // 5% down
			status = "down"
		}
		checker.addToCache(CheckResult{
			ServiceName: "API",
			Status:      status,
			CheckedAt:   date.Add(time.Duration(i) * time.Minute),
		})
	}

	// Test with degraded performance
	dayData = checker.getDayDataForService("API", date)
	if dayData.Status != "degraded" {
		t.Errorf("Expected status 'degraded' for 95%% uptime, got %q", dayData.Status)
	}
	if dayData.Uptime != 95.0 {
		t.Errorf("Expected uptime 95%%, got %.2f%%", dayData.Uptime)
	}
}

// Test concurrent operations
func TestConcurrentCacheOperations(t *testing.T) {
	checker := &HealthChecker{
		cache:         &sync.Map{},
		lastCheckTime: &sync.Map{},
	}

	// Perform concurrent writes
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(n int) {
			for j := 0; j < 100; j++ {
				checker.addToCache(CheckResult{
					ServiceName: fmt.Sprintf("Service%d", n),
					Status:      "up",
					CheckedAt:   time.Now(),
				})
			}
			done <- true
		}(i)
	}

	// Wait for all goroutines
	for i := 0; i < 10; i++ {
		<-done
	}

	// Verify data integrity
	count := 0
	checker.cache.Range(func(key, value interface{}) bool {
		checks := value.([]CheckResult)
		if len(checks) != 100 {
			t.Errorf("Expected 100 checks, got %d", len(checks))
		}
		count++
		return true
	})

	if count != 10 {
		t.Errorf("Expected 10 cache entries, got %d", count)
	}
}

// Test CheckAllServices with concurrency
func TestCheckAllServices(t *testing.T) {
	// Create a counter for service calls
	var callCount sync.Map

	// Create test server that counts calls
	serviceServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Count calls per path
		count := 0
		if val, ok := callCount.Load(r.URL.Path); ok {
			count = val.(int)
		}
		callCount.Store(r.URL.Path, count+1)

		w.WriteHeader(http.StatusOK)
	}))
	defer serviceServer.Close()

	// Mock D1 server
	d1Server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := D1Response{Success: true}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer d1Server.Close()

	// Override monitoredServices temporarily
	originalServices := monitoredServices
	monitoredServices = []Service{
		{Name: "Service1", URL: serviceServer.URL + "/1", HealthEndpoint: ""},
		{Name: "Service2", URL: serviceServer.URL + "/2", HealthEndpoint: ""},
		{Name: "Service3", URL: serviceServer.URL + "/3", HealthEndpoint: ""},
	}
	defer func() {
		monitoredServices = originalServices
	}()

	d1Client := &D1Client{
		BaseURL:    d1Server.URL,
		AccountID:  "test",
		DatabaseID: "test",
		APIKey:     "test",
	}

	checker := NewHealthChecker(d1Client)
	checker.CheckAllServices()

	// Verify all services were checked
	totalCalls := 0
	callCount.Range(func(key, value interface{}) bool {
		totalCalls += value.(int)
		return true
	})

	if totalCalls != 3 {
		t.Errorf("Expected 3 service calls, got %d", totalCalls)
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && strings.Contains(s, substr))
}
