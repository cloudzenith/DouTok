package launcher

import (
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

//func initTracer(serviceName string, endpoint string) error {
//	exporter, err := otlptrace.New(
//		context.Background(),
//		otlptracegrpc.NewClient(otlptracegrpc.WithEndpoint(endpoint)),
//	)
//	if err != nil {
//		return err
//	}
//
//	tp := trace.NewTracerProvider(
//		trace.WithSampler(trace.ParentBased(trace.TraceIDRatioBased(1.0))),
//		trace.WithBatcher(exporter),
//		trace.WithResource(resource.NewSchemaless(
//			semconv.ServiceNameKey.String(serviceName),
//			attribute.String("exporter", "tempo"),
//		)),
//	)
//
//	otel.SetTracerProvider(tp)
//	log.Infof("Tracing enabled. Exporting spans to %s with service name %s", endpoint, serviceName)
//	return nil
//}

func initTracer(serviceName string, endpoint string) error {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)))
	if err != nil {
		return err
	}

	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.ParentBased(trace.TraceIDRatioBased(1.0))),
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(serviceName),
			attribute.String("exporter", "jaeger"),
		)),
	)

	otel.SetTracerProvider(tp)
	log.Infof("Tracing enabled. Exporting spans to %s with service name %s", endpoint, serviceName)
	return nil
}
