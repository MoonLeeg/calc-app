package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		body           string
		expectedStatus int
		expectedBody   *Response
	}{
		{
			name:           "Valid Expression",
			method:         http.MethodPost,
			body:           `{"expression": "2 + 2"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   &Response{Result: "4"},
		},
		{
			name:           "Invalid Method",
			method:         http.MethodGet,
			body:           `{"expression": "2 + 2"}`,
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Empty Expression",
			method:         http.MethodPost,
			body:           `{"expression": ""}`,
			expectedStatus: http.StatusUnprocessableEntity,
		},
		{
			name:           "Invalid JSON",
			method:         http.MethodPost,
			body:           `{"invalid json"}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Division by Zero",
			method:         http.MethodPost,
			body:           `{"expression": "1/0"}`,
			expectedStatus: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/calculate", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			CalculateHandler(w, req)

			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, resp.StatusCode)
			}

			if tt.expectedBody != nil {
				var gotResp Response
				if err := json.Unmarshal(body, &gotResp); err != nil {
					t.Fatalf("Failed to unmarshal response body: %v", err)
				}

				if gotResp.Result != tt.expectedBody.Result {
					t.Errorf("Expected result %s, got %s", tt.expectedBody.Result, gotResp.Result)
				}
			}
		})
	}
}
