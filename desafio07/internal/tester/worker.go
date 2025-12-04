package tester

import (
	"net/http"
	"sync"
	"time"
)

func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     90 * time.Second,
		},
	}
}

func worker(
	url string,
	client *http.Client,
	jobsChan <-chan int,
	resultsChan chan<- RequestResult,
) {
	for range jobsChan {
		result := executeRequest(url, client)
		resultsChan <- result
	}
}

func executeRequest(url string, client *http.Client) RequestResult {
	start := time.Now()
	resp, err := client.Get(url)
	duration := time.Since(start)

	if err != nil {
		return RequestResult{
			StatusCode: 0,
			Duration:   duration,
			Error:      err,
		}
	}
	defer resp.Body.Close()

	return RequestResult{
		StatusCode: resp.StatusCode,
		Duration:   duration,
		Error:      nil,
	}
}

func RunLoadTest(url string, totalRequests, concurrency int) *Statistics {
	client := NewHTTPClient()

	jobsChan := make(chan int, totalRequests)
	resultsChan := make(chan RequestResult, totalRequests)

	for i := 0; i < totalRequests; i++ {
		jobsChan <- i
	}
	close(jobsChan)

	// Init workers
	var wg sync.WaitGroup
	wg.Add(concurrency)

	start := time.Now()

	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			worker(url, client, jobsChan, resultsChan)
		}()
	}

	// Wait for results
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Collect results
	statistics := NewStatistics()
	for result := range resultsChan {
		statistics.AddRequestResult(result)
	}

	statistics.TotalDuration = time.Since(start)

	return statistics
}
