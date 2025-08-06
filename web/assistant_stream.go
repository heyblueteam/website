package web

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type StreamRequest struct {
	Message  string `json:"message"`
	ThreadID string `json:"threadId,omitempty"`
}

type StreamChunk struct {
	Content  string `json:"content"`
	ThreadID string `json:"threadId,omitempty"`
	Done     bool   `json:"done"`
	Error    string `json:"error,omitempty"`
}

type AssistantStreamEvent struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

type MessageDelta struct {
	Delta struct {
		Content []struct {
			Index int `json:"index"`
			Type  string `json:"type"`
			Text  *struct {
				Value       string      `json:"value"`
				Annotations []interface{} `json:"annotations"`
			} `json:"text,omitempty"`
		} `json:"content"`
	} `json:"delta"`
}

func HandleAssistantStream(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get API key and Assistant ID from environment
	apiKey := os.Getenv("OPENAI_ASSISTANT_API_KEY")
	assistantID := os.Getenv("OPENAI_ASSISTANT_ID")

	if apiKey == "" || assistantID == "" {
		sendStreamError(w, "AI Assistant not configured")
		return
	}

	// Parse request
	var req StreamRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendStreamError(w, "Invalid request")
		return
	}

	if req.Message == "" {
		sendStreamError(w, "Message is required")
		return
	}

	// Set SSE headers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Create or use existing thread
	threadID := req.ThreadID
	if threadID == "" {
		thread, err := createThread(apiKey)
		if err != nil {
			sendStreamError(w, "Failed to create conversation thread")
			return
		}
		threadID = thread.ID
	}

	// Send initial chunk with threadID
	initialChunk := StreamChunk{
		Content:  "",
		ThreadID: threadID,
		Done:     false,
	}
	sendSSEData(w, initialChunk)

	// Add message to thread
	if err := addMessage(apiKey, threadID, req.Message); err != nil {
		errorChunk := StreamChunk{
			Error: "Failed to send message",
			Done:  true,
		}
		sendSSEData(w, errorChunk)
		fmt.Fprintf(w, "data: [DONE]\n\n")
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
		return
	}

	// Create a streaming run
	runBody := map[string]interface{}{
		"assistant_id": assistantID,
		"stream":       true,
	}

	jsonBody, err := json.Marshal(runBody)
	if err != nil {
		sendStreamError(w, "Failed to process request")
		return
	}

	// Create the run request with streaming
	runReq, err := http.NewRequest("POST",
		fmt.Sprintf("https://api.openai.com/v1/threads/%s/runs", threadID),
		bytes.NewBuffer(jsonBody))
	if err != nil {
		sendStreamError(w, "Failed to create run request")
		return
	}

	runReq.Header.Set("Authorization", "Bearer "+apiKey)
	runReq.Header.Set("Content-Type", "application/json")
	runReq.Header.Set("OpenAI-Beta", "assistants=v2")

	// Make the streaming request
	client := &http.Client{}
	resp, err := client.Do(runReq)
	if err != nil {
		sendStreamError(w, "Failed to connect to assistant")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		errorChunk := StreamChunk{
			Error: fmt.Sprintf("Assistant error: %s", string(body)),
			Done:  true,
		}
		sendSSEData(w, errorChunk)
		fmt.Fprintf(w, "data: [DONE]\n\n")
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
		return
	}

	// Process the SSE stream from OpenAI
	scanner := bufio.NewScanner(resp.Body)
	var currentEvent string
	
	for scanner.Scan() {
		line := scanner.Text()
		
		// Skip empty lines
		if line == "" {
			continue
		}

		// Debug logging disabled - uncomment for troubleshooting
		// fmt.Printf("[Assistant Stream Debug] Line: %s\n", line)

		// Parse event type
		if strings.HasPrefix(line, "event: ") {
			currentEvent = strings.TrimPrefix(line, "event: ")
			// fmt.Printf("[Assistant Stream Debug] Event type: %s\n", currentEvent)
			continue
		}

		// Parse SSE data
		if strings.HasPrefix(line, "data: ") {
			data := strings.TrimPrefix(line, "data: ")
			
			// Skip [DONE] message
			if data == "[DONE]" {
				// fmt.Println("[Assistant Stream Debug] Received [DONE]")
				// Send our own DONE message
				finalChunk := StreamChunk{
					Content:  "",
					ThreadID: threadID,
					Done:     true,
				}
				sendSSEData(w, finalChunk)
				fmt.Fprintf(w, "data: [DONE]\n\n")
				if f, ok := w.(http.Flusher); ok {
					f.Flush()
				}
				break
			}

			// Handle specific event types
			switch currentEvent {
			case "thread.message.delta":
				// Parse the delta JSON
				var deltaData map[string]interface{}
				if err := json.Unmarshal([]byte(data), &deltaData); err != nil {
					// fmt.Printf("[Assistant Stream Debug] Failed to parse delta JSON: %v\n", err)
					continue
				}

				// Extract text from delta.content[0].text.value
				if delta, ok := deltaData["delta"].(map[string]interface{}); ok {
					if content, ok := delta["content"].([]interface{}); ok {
						for _, item := range content {
							if contentItem, ok := item.(map[string]interface{}); ok {
								if contentItem["type"] == "text" {
									if text, ok := contentItem["text"].(map[string]interface{}); ok {
										if value, ok := text["value"].(string); ok {
											// fmt.Printf("[Assistant Stream Debug] Sending text chunk: %s\n", value)
											// Send the text chunk
											chunk := StreamChunk{
												Content:  value,
												ThreadID: threadID,
												Done:     false,
											}
											sendSSEData(w, chunk)
										}
									}
								}
							}
						}
					}
				}

			case "thread.run.completed":
				// fmt.Println("[Assistant Stream Debug] Run completed")
				// Send completion signal
				finalChunk := StreamChunk{
					Content:  "",
					ThreadID: threadID,
					Done:     true,
				}
				sendSSEData(w, finalChunk)
				fmt.Fprintf(w, "data: [DONE]\n\n")
				if f, ok := w.(http.Flusher); ok {
					f.Flush()
				}
				break

			case "thread.run.failed", "thread.run.cancelled", "thread.run.expired":
				// fmt.Printf("[Assistant Stream Debug] Run failed/cancelled/expired: %s\n", currentEvent)
				errorChunk := StreamChunk{
					Error: fmt.Sprintf("Assistant run %s", strings.TrimPrefix(currentEvent, "thread.run.")),
					Done:  true,
				}
				sendSSEData(w, errorChunk)
				fmt.Fprintf(w, "data: [DONE]\n\n")
				if f, ok := w.(http.Flusher); ok {
					f.Flush()
				}
				break

			case "thread.run.requires_action":
				// fmt.Println("[Assistant Stream Debug] Run requires action (function calling not supported)")
				errorChunk := StreamChunk{
					Error: "Assistant requires action (function calling not supported in streaming)",
					Done:  true,
				}
				sendSSEData(w, errorChunk)
				fmt.Fprintf(w, "data: [DONE]\n\n")
				if f, ok := w.(http.Flusher); ok {
					f.Flush()
				}
				break
			}
			
			// Check if we should exit the loop after handling the event
			if currentEvent == "thread.run.completed" || currentEvent == "thread.run.failed" || 
			   currentEvent == "thread.run.cancelled" || currentEvent == "thread.run.expired" || 
			   currentEvent == "thread.run.requires_action" {
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		errorChunk := StreamChunk{
			Error: "Stream interrupted",
			Done:  true,
		}
		sendSSEData(w, errorChunk)
		fmt.Fprintf(w, "data: [DONE]\n\n")
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	}
}

func sendSSEData(w http.ResponseWriter, chunk StreamChunk) {
	data, err := json.Marshal(chunk)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "data: %s\n\n", string(data))
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}

func sendStreamError(w http.ResponseWriter, message string) {
	// For streaming endpoint, we still need to send as SSE if headers are set
	if w.Header().Get("Content-Type") == "text/event-stream" {
		errorChunk := StreamChunk{
			Error: message,
			Done:  true,
		}
		sendSSEData(w, errorChunk)
		fmt.Fprintf(w, "data: [DONE]\n\n")
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	} else {
		// Regular error response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": message,
		})
	}
}