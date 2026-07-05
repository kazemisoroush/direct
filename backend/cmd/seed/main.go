// Command seed writes restaurants (with menus) into the DynamoDB table out-of-band.
// The API Lambda is read-only; seeding runs from a developer/operator with write creds:
//
//	go run ./cmd/seed --table <DIRECT_TABLE> --file cmd/seed/hills-kebabs.json
package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/kazemisoroush/direct/backend/internal/domain"
)

func main() {
	table := flag.String("table", os.Getenv("DIRECT_TABLE"), "DynamoDB restaurants table name")
	file := flag.String("file", "", "path to a JSON file with a list of restaurants")
	flag.Parse()

	if *table == "" || *file == "" {
		log.Fatal("usage: seed --table <name> --file <restaurants.json>")
	}

	raw, err := os.ReadFile(*file)
	if err != nil {
		log.Fatalf("read %s: %v", *file, err)
	}

	var restaurants []domain.Restaurant
	if err := json.Unmarshal(raw, &restaurants); err != nil {
		log.Fatalf("parse %s: %v", *file, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	awsCfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("load AWS config: %v", err)
	}
	client := dynamodb.NewFromConfig(awsCfg)

	for _, r := range restaurants {
		item, err := attributevalue.MarshalMap(r)
		if err != nil {
			log.Fatalf("marshal %s: %v", r.ID, err)
		}
		if _, err := client.PutItem(ctx, &dynamodb.PutItemInput{TableName: aws.String(*table), Item: item}); err != nil {
			log.Fatalf("put %s: %v", r.ID, err)
		}
		log.Printf("seeded %s — %s (%d menu items)", r.ID, r.Name, len(r.Menu))
	}
	log.Printf("done: %d restaurant(s) seeded into %s", len(restaurants), *table)
}
