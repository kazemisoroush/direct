package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadReadsEnv(t *testing.T) {
	t.Setenv("DIRECT_TABLE", "restaurants")
	t.Setenv("DIRECT_JWT_ISSUER", "https://issuer.example")
	t.Setenv("DIRECT_JWT_CLIENT_ID", "client123")
	t.Setenv("DIRECT_ADDR", ":9000")
	t.Setenv("DIRECT_AUTH_DISABLED", "true")

	cfg := Load()

	assert.Equal(t, "restaurants", cfg.Table)
	assert.Equal(t, "https://issuer.example", cfg.JWTIssuer)
	assert.Equal(t, "client123", cfg.JWTClientID)
	assert.Equal(t, ":9000", cfg.Addr)
	assert.True(t, cfg.AuthDisabled)
}

func TestAuthDisabledOnlyOnExactTrue(t *testing.T) {
	t.Setenv("DIRECT_AUTH_DISABLED", "1")
	assert.False(t, Load().AuthDisabled)
}

func TestServerAddrDefaultsWhenUnset(t *testing.T) {
	assert.Equal(t, ":8080", Config{}.ServerAddr())
	assert.Equal(t, ":9000", Config{Addr: ":9000"}.ServerAddr())
}

func TestAuthEnabledRequiresIssuerAndClient(t *testing.T) {
	assert.False(t, Config{}.AuthEnabled(), "no issuer or client")
	assert.False(t, Config{JWTIssuer: "x"}.AuthEnabled(), "issuer only")
	assert.False(t, Config{JWTClientID: "y"}.AuthEnabled(), "client only")
	assert.True(t, Config{JWTIssuer: "x", JWTClientID: "y"}.AuthEnabled(), "both set")
}
