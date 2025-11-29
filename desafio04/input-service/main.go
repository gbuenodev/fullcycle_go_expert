package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func main() {
	ctx := context.Background()

	// Initialize OTEL with Zipkin
	shutdown, err := initTelemetry(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize telemetry: %v", err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Printf("Error during telemetry shutdown: %v", err)
		}
	}()

	// Create HTTP handler
	http.HandleFunc("/weather", handleWeatherRequest)

	// Start server
	port := getEnv("PORT", "8000")
	log.Printf("Input service starting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func handleWeatherRequest(w http.ResponseWriter, r *http.Request) {
	// Only accept POST
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var reqBody struct {
		CEP string `json:"cep"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if reqBody.CEP == "" {
		http.Error(w, "missing cep field", http.StatusBadRequest)
		return
	}
	cep := reqBody.CEP

	// Create traced HTTP client
	client := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	// Prepare request to backend
	backendURL := getEnv("BACKEND_URL", "http://backend:3000/weather")
	backendReqBody := map[string]string{"cep": cep}
	jsonBody, err := json.Marshal(backendReqBody)
	if err != nil {
		http.Error(w, "failed to create request", http.StatusInternalServerError)
		return
	}

	// Create request with context (trace propagation happens here)
	req, err := http.NewRequestWithContext(r.Context(), "POST", backendURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		http.Error(w, "failed to create request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// Make request to backend (trace context propagates via headers)
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("backend request failed: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read backend response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "failed to read backend response", http.StatusInternalServerError)
		return
	}

	// Forward backend response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func initTelemetry(ctx context.Context) (func(context.Context) error, error) {
	// Create resource
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(getEnv("OTEL_SERVICE_NAME", "input-service")),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// Create Zipkin exporter
	exporter, err := zipkin.New(getEnv("ZIPKIN_ENDPOINT", "http://zipkin:9411/api/v2/spans"))
	if err != nil {
		return nil, fmt.Errorf("failed to create zipkin exporter: %w", err)
	}

	// Create trace provider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)

	// Set global tracer provider
	otel.SetTracerProvider(tp)

	return tp.Shutdown, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
