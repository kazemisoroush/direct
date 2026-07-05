package restaurant

import (
	"context"
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/kazemisoroush/direct/backend/internal/domain"
)

// postcodePattern matches a 4-digit Australian postcode inside an address string.
var postcodePattern = regexp.MustCompile(`\b\d{4}\b`)

// DynamoStore reads restaurants from a single DynamoDB table.
type DynamoStore struct {
	table  string
	client *dynamodb.Client
}

// NewDynamoStore builds a DynamoStore over one table.
func NewDynamoStore(client *dynamodb.Client, table string) *DynamoStore {
	return &DynamoStore{table: table, client: client}
}

// ListDeliveringTo returns the restaurants that deliver to the address. At M1/M2 scale a
// scan is fine (one restaurant); a postcode GSI can replace it when the catalogue grows.
func (s *DynamoStore) ListDeliveringTo(ctx context.Context, address string) ([]domain.Restaurant, error) {
	out, err := s.client.Scan(ctx, &dynamodb.ScanInput{TableName: aws.String(s.table)})
	if err != nil {
		return nil, fmt.Errorf("scan restaurants: %w", err)
	}

	var all []domain.Restaurant
	if err := attributevalue.UnmarshalListOfMaps(out.Items, &all); err != nil {
		return nil, fmt.Errorf("unmarshal restaurants: %w", err)
	}

	postcode := postcodePattern.FindString(address)
	// Always return a non-nil slice so the API encodes [] rather than null.
	result := make([]domain.Restaurant, 0, len(all))
	for _, r := range all {
		if deliversTo(r, postcode) {
			result = append(result, r)
		}
	}
	return result, nil
}

// deliversTo reports whether the restaurant delivers to the postcode. A restaurant with no
// listed postcodes is treated as delivering everywhere.
func deliversTo(r domain.Restaurant, postcode string) bool {
	if len(r.DeliversToPostcodes) == 0 {
		return true
	}
	for _, p := range r.DeliversToPostcodes {
		if p == postcode {
			return true
		}
	}
	return false
}
