package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/kazemisoroush/direct/backend/internal/mocks"
)

func okHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusOK) })
}

func TestHealthIsOpen(t *testing.T) {
	// Arrange: /health must pass without the verifier being consulted at all.
	ctrl := gomock.NewController(t)
	verifier := mocks.NewMockTokenVerifier(ctrl)
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	// Act
	NewAuthMiddleware(verifier).Wrap(okHandler()).ServeHTTP(rec, req)

	// Assert
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestMissingTokenIsRejected(t *testing.T) {
	ctrl := gomock.NewController(t)
	verifier := mocks.NewMockTokenVerifier(ctrl)
	req := httptest.NewRequest(http.MethodGet, "/restaurants", nil)
	rec := httptest.NewRecorder()

	NewAuthMiddleware(verifier).Wrap(okHandler()).ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestInvalidTokenIsRejected(t *testing.T) {
	ctrl := gomock.NewController(t)
	verifier := mocks.NewMockTokenVerifier(ctrl)
	verifier.EXPECT().Verify("wrong").Return(errors.New("invalid"))
	req := httptest.NewRequest(http.MethodGet, "/restaurants", nil)
	req.Header.Set("Authorization", "Bearer wrong")
	rec := httptest.NewRecorder()

	NewAuthMiddleware(verifier).Wrap(okHandler()).ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestValidTokenPasses(t *testing.T) {
	ctrl := gomock.NewController(t)
	verifier := mocks.NewMockTokenVerifier(ctrl)
	verifier.EXPECT().Verify("good").Return(nil)
	req := httptest.NewRequest(http.MethodGet, "/restaurants", nil)
	req.Header.Set("Authorization", "Bearer good")
	rec := httptest.NewRecorder()

	NewAuthMiddleware(verifier).Wrap(okHandler()).ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
