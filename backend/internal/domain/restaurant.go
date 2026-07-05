// Package domain holds the core Direct types shared across the backend.
package domain

// Restaurant is a place a customer can order from directly. DeliversToPostcodes lists the
// postcodes it delivers to; an empty list means it delivers everywhere (no restriction yet).
// Menu is populated on the detail read (GET /restaurants/{id}) and omitted from list results.
type Restaurant struct {
	ID                  string     `json:"id" dynamodbav:"id"`
	Name                string     `json:"name" dynamodbav:"name"`
	Suburb              string     `json:"suburb" dynamodbav:"suburb"`
	Address             string     `json:"address" dynamodbav:"address"`
	Phone               string     `json:"phone,omitempty" dynamodbav:"phone,omitempty"`
	DeliversToPostcodes []string   `json:"deliversToPostcodes" dynamodbav:"deliversToPostcodes"`
	Menu                []MenuItem `json:"menu,omitempty" dynamodbav:"menu,omitempty"`
}

// MenuItem is one orderable item. PriceCents is the direct in-store price in cents (integer
// money, no floats) — the price the customer pays without the delivery-platform markup.
type MenuItem struct {
	ID          string `json:"id" dynamodbav:"id"`
	Name        string `json:"name" dynamodbav:"name"`
	Description string `json:"description,omitempty" dynamodbav:"description,omitempty"`
	PriceCents  int64  `json:"priceCents" dynamodbav:"priceCents"`
	Category    string `json:"category" dynamodbav:"category"`
}
