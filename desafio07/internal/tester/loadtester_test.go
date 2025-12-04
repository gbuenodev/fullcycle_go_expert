package tester

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

// createTestServer cria um servidor HTTP de teste
func createTestServer(statusCode int, delay time.Duration) *httptest.Server {
	var counter int32

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if delay > 0 {
			time.Sleep(delay)
		}

		atomic.AddInt32(&counter, 1)

		w.WriteHeader(statusCode)
		fmt.Fprintf(w, "Request #%d", atomic.LoadInt32(&counter))
	}))
}

func TestRunLoadTest_BasicFunctionality(t *testing.T) {
	server := createTestServer(200, 0)
	defer server.Close()

	stats := RunLoadTest(server.URL, 10, 2)

	if stats.TotalRequests != 10 {
		t.Errorf("Expected 10 requests, got %d", stats.TotalRequests)
	}

	if stats.SuccessCount != 10 {
		t.Errorf("Expected 10 successful requests, got %d", stats.SuccessCount)
	}
}

func TestRunLoadTest_ExactRequestCount(t *testing.T) {
	server := createTestServer(200, 0)
	defer server.Close()

	stats := RunLoadTest(server.URL, 17, 5)

	if stats.TotalRequests != 17 {
		t.Errorf("Expected exactly 17 requests, got %d", stats.TotalRequests)
	}
}

func TestRunLoadTest_HandlesErrors(t *testing.T) {
	stats := RunLoadTest("http://localhost:99999", 5, 2)

	if stats.ErrorCount != 5 {
		t.Errorf("Expected 5 errors, got %d", stats.ErrorCount)
	}
}
