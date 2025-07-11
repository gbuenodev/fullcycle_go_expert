package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	// Create a 300ms context for the HTTP request
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	// Create the request with context
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Printf("error creating request: %s\n", err)
		return
	}

	// Execute the HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error making request: %s\n", err)
		return
	}
	defer resp.Body.Close()

	// Decode the response body
	var result struct {
		Bid string `json:"bid"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("error decoding response: %s\n", err)
		return
	}

	// Write the bid to a file
	file, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Printf("error creating file: %s\n", err)
		return
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "Dólar:%s\n", result.Bid)
	if err != nil {
		fmt.Printf("error writing to file: %s\n", err)
		return
	}

	fmt.Printf("Successfully saved exchange to: %s\n", file.Name())
}
