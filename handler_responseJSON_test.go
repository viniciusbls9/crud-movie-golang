package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespondWithError(t *testing.T) {
	w := httptest.NewRecorder()
	code := http.StatusInternalServerError
	message := "Internal Server Error"

	respondWithError(w, code, message)

	if w.Code != code {
		t.Errorf("Expected status code %d, but got %d", code, w.Code)
	}

	var response struct {
		Error string `json:"error"`
	}
	err := parseJSON(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to parse JSON response: %v", err)
	}

	if response.Error != message {
		t.Errorf("Expected error message '%s', but got '%s'", message, response.Error)
	}
}

func parseJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func TestRespondWithJSON(t *testing.T) {
	w := httptest.NewRecorder()

	payload := struct {
		Message string `json:"message"`
	}{
		Message: "Hello, World!",
	}

	respondWithJSON(w, http.StatusOK, payload)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	expectedContentType := "application/json"
	if contentType != expectedContentType {
		t.Errorf("Expected Content-Type %s, but got %s", expectedContentType, contentType)
	}

	var response struct {
		Message string `json:"message"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to parse JSON response: %v", err)
	}

	expectedMessage := "Hello, World!"
	if response.Message != expectedMessage {
		t.Errorf("Expected message '%s', but got '%s'", expectedMessage, response.Message)
	}
}
