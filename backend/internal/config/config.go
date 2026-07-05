// Package config centralises reading runtime settings from the environment.
package config

import "os"

const defaultServerAddr = ":8080"

// Config holds every runtime setting the backend reads from the environment.
type Config struct {
	Table        string
	JWTIssuer    string
	JWTClientID  string
	Addr         string
	AuthDisabled bool
}

// Load reads the configuration from environment variables.
func Load() Config {
	return Config{
		Table:        os.Getenv("DIRECT_TABLE"),
		JWTIssuer:    os.Getenv("DIRECT_JWT_ISSUER"),
		JWTClientID:  os.Getenv("DIRECT_JWT_CLIENT_ID"),
		Addr:         os.Getenv("DIRECT_ADDR"),
		AuthDisabled: os.Getenv("DIRECT_AUTH_DISABLED") == "true",
	}
}

// ServerAddr returns the local server address, defaulting when unset.
func (c Config) ServerAddr() string {
	if c.Addr == "" {
		return defaultServerAddr
	}
	return c.Addr
}

// AuthEnabled reports whether JWT verification is fully configured.
func (c Config) AuthEnabled() bool {
	return c.JWTIssuer != "" && c.JWTClientID != ""
}
