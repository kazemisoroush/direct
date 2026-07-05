package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHealthReturnsOK checks GET /health responds 200 with a status body.
func TestHealthReturnsOK(t *testing.T) {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	// Act
	New().ServeHTTP(rec, req)

	// Assert
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := rec.Body.String(); got != "{\"status\":\"ok\"}\n" {
		t.Fatalf("body = %q, want status ok", got)
	}
}
