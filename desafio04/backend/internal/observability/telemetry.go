package observability

import (
	"context"
	"log"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func InitTelemetry(ctx context.Context) (shutdown func(context.Context) error, err error) {
	shutdownTracer, err := initTracer(ctx)
	if err != nil {
		return nil, err
	}

	shutdownMeter, err := initMeter(ctx)
	if err != nil {
		return nil, err
	}

	return func(ctx context.Context) error {
		err := shutdownTracer(ctx)
		if err != nil {
			log.Printf("error shutting down tracer: %v", err)
		}

		err = shutdownMeter(ctx)
		if err != nil {
			log.Printf("error shutting down meter: %v", err)
		}

		return nil
	}, nil
}

func initTracer(ctx context.Context) (shutdown func(context.Context) error, err error) {
	res, err := resource.New(ctx, resource.WithAttributes(
		semconv.ServiceName(os.Getenv("OTEL_SERVICE_NAME")),
	))
	if err != nil {
		log.Printf("failed to create resource: %v", err)
		return nil, err
	}

	exporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		log.Printf("failed to create trace exporter: %v", err)
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)

	otel.SetTracerProvider(tp)

	return tp.Shutdown, nil
}

func initMeter(ctx context.Context) (shutdown func(context.Context) error, err error) {
	res, err := resource.New(
		ctx,
		resource.WithAttributes(
			semconv.ServiceName(os.Getenv("OTEL_SERVICE_NAME")),
		),
	)
	if err != nil {
		log.Printf("failed to create resource: %v", err)
		return nil, err
	}

	exporter, err := prometheus.New()
	if err != nil {
		log.Printf("failed to create prometheus exporter: %v", err)
		return nil, err
	}

	mp := metric.NewMeterProvider(
		metric.WithReader(exporter),
		metric.WithResource(res),
	)

	otel.SetMeterProvider(mp)

	return mp.Shutdown, nil
}
