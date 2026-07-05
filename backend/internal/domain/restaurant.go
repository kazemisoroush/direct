// Package domain holds the core Direct types shared across the backend.
package domain

// Restaurant is a place a customer can order from directly. DeliversToPostcodes lists the
// postcodes it delivers to; an empty list means it delivers everywhere (no restriction yet).
type Restaurant struct {
	ID                  string   `json:"id" dynamodbav:"id"`
	Name                string   `json:"name" dynamodbav:"name"`
	Suburb              string   `json:"suburb" dynamodbav:"suburb"`
	Address             string   `json:"address" dynamodbav:"address"`
	DeliversToPostcodes []string `json:"deliversToPostcodes" dynamodbav:"deliversToPostcodes"`
}
