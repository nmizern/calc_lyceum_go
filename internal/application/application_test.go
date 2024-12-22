package application_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nmizern/calc_lyceum_go/internal/application"
)

func TestCalcHandler(t *testing.T) {
	// тестовый сервер
	handler := http.HandlerFunc(application.CalcHandler)
	srv := httptest.NewServer(handler)
	defer srv.Close()

	tests := []struct {
		name           string
		expression     string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid expression",
			expression:     `1+2`,
			expectedStatus: http.StatusOK,
			expectedBody:   `"result"`,
		},
		{
			name:           "Invalid symbol",
			expression:     `1+2a`,
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   `"error":"Expression is not valid"`,
		},
		{
			name:           "Division by zero",
			expression:     `10/0`,
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   `"error":"Expression is not valid"`,
		},
		{
			name:           "Mismatched brackets",
			expression:     `(1+2`,
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   `"error":"Expression is not valid"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			requestBody, _ := json.Marshal(map[string]string{
				"expression": tt.expression,
			})

			req, err := http.NewRequest(http.MethodPost, srv.URL, bytes.NewReader(requestBody))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			// запрос к серверу
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("Failed to do request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}

			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Failed to read response body: %v", err)
			}
			bodyStr := string(bodyBytes)

			if !contains(bodyStr, tt.expectedBody) {
				t.Errorf("Expected body to contain %q, got %q", tt.expectedBody, bodyStr)
			}
		})
	}
}

func contains(body, sub string) bool {
	return bytes.Contains([]byte(body), []byte(sub))
}
