// Command seed loads restaurants from a JSON file and POSTs each to the Direct API.
// API-first: data enters the catalogue through the API, never by writing to storage directly.
//
//	go run ./cmd/seed --api https://<api-id>.execute-api.<region>.amazonaws.com \
//	  --token <cognito-access-token> --file cmd/seed/hills-kebabs.json
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/kazemisoroush/direct/backend/internal/domain"
)

func main() {
	api := flag.String("api", "", "Direct API base URL")
	token := flag.String("token", "", "Cognito access token (sent as a Bearer token)")
	file := flag.String("file", "", "path to a JSON file with a list of restaurants")
	flag.Parse()

	if *api == "" || *token == "" || *file == "" {
		log.Fatal("usage: seed --api <url> --token <bearer> --file <restaurants.json>")
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

	client := &http.Client{Timeout: 15 * time.Second}
	endpoint := strings.TrimRight(*api, "/") + "/restaurants"

	for _, r := range restaurants {
		body, err := json.Marshal(r)
		if err != nil {
			log.Fatalf("marshal %s: %v", r.ID, err)
		}
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
		if err != nil {
			log.Fatalf("build request for %s: %v", r.ID, err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+*token)

		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("POST %s: %v", r.ID, err)
		}
		payload, _ := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if resp.StatusCode != http.StatusCreated {
			log.Fatalf("POST %s: unexpected status %d: %s", r.ID, resp.StatusCode, payload)
		}
		log.Printf("seeded %s — %s via API (%d menu items)", r.ID, r.Name, len(r.Menu))
	}
	log.Printf("done: %d restaurant(s) POSTed to %s", len(restaurants), endpoint)
}
