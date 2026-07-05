package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// fakeVerifier accepts exactly one token value.
type fakeVerifier struct{ valid string }

func (f fakeVerifier) Verify(token string) error {
	if token == f.valid {
		return nil
	}
	return errors.New("invalid")
}

func okHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusOK) })
}

func TestHealthIsOpen(t *testing.T) {
	mw := NewAuthMiddleware(fakeVerifier{valid: "good"})
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	mw.Wrap(okHandler()).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("health status = %d, want 200 (no token required)", rec.Code)
	}
}

func TestMissingTokenIsRejected(t *testing.T) {
	mw := NewAuthMiddleware(fakeVerifier{valid: "good"})
	req := httptest.NewRequest(http.MethodGet, "/restaurants", nil)
	rec := httptest.NewRecorder()

	mw.Wrap(okHandler()).ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401 for missing token", rec.Code)
	}
}

func TestInvalidTokenIsRejected(t *testing.T) {
	mw := NewAuthMiddleware(fakeVerifier{valid: "good"})
	req := httptest.NewRequest(http.MethodGet, "/restaurants", nil)
	req.Header.Set("Authorization", "Bearer wrong")
	rec := httptest.NewRecorder()

	mw.Wrap(okHandler()).ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401 for invalid token", rec.Code)
	}
}

func TestValidTokenPasses(t *testing.T) {
	mw := NewAuthMiddleware(fakeVerifier{valid: "good"})
	req := httptest.NewRequest(http.MethodGet, "/restaurants", nil)
	req.Header.Set("Authorization", "Bearer good")
	rec := httptest.NewRecorder()

	mw.Wrap(okHandler()).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200 for valid token", rec.Code)
	}
}
