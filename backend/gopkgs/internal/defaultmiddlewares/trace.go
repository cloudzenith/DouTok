package defaultmiddlewares

import (
	"context"
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	kratoserrs "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/google/uuid"
	"time"
)

const (
	TraceIdHeadersKey = "x-trace-id"
	SpanIdHeadersKey  = "x-span-id"
)

type traceIdKey struct{}
type spanIdKey struct{}

func generateId() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func TraceIdInjector() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (i interface{}, err error) {
			meta, ok := metadata.FromServerContext(ctx)
			if !ok {
				return nil, errors.New("metadata not found")
			}

			traceId := meta.Get(TraceIdHeadersKey)
			if traceId == "" {
				traceId, err = generateId()
				if err != nil {
					return nil, err
				}

				metadata.AppendToClientContext(ctx, TraceIdHeadersKey, traceId)
			}
			ctx = context.WithValue(ctx, traceIdKey{}, traceId)
			return handler(ctx, req)
		}
	}
}

func GetTraceId() log.Valuer {
	return func(ctx context.Context) interface{} {
		traceId := ctx.Value(traceIdKey{})
		if traceId == nil {
			return ""
		}

		return traceId
	}
}

func SpanIdInjector() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			spanId, err := generateId()
			if err != nil {
				return nil, err
			}

			ctx = context.WithValue(ctx, spanIdKey{}, spanId)
			return handler(ctx, req)
		}
	}
}

func GetSpanId() log.Valuer {
	return func(ctx context.Context) interface{} {
		spanId := ctx.Value(spanIdKey{})
		if spanId == nil {
			return ""
		}

		return spanId
	}
}

// extractError returns the string of the error
func extractError(err error) (log.Level, string) {
	if err != nil {
		return log.LevelError, fmt.Sprintf("%+v", err)
	}
	return log.LevelInfo, ""
}

// extractArgs returns the string of the req
func extractArgs(req interface{}) string {
	if redacter, ok := req.(logging.Redacter); ok {
		return redacter.Redact()
	}
	if stringer, ok := req.(fmt.Stringer); ok {
		return stringer.String()
	}
	return fmt.Sprintf("%+v", req)
}

func extractArgs2Json(req interface{}) string {
	output, err := sonic.Marshal(req)
	if err != nil {
		return err.Error()
	}

	return string(output)
}

func RequestMonitor() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				code      int32
				reason    string
				kind      string
				operation string
			)
			startTime := time.Now()
			if info, ok := transport.FromServerContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
			}
			reply, err = handler(ctx, req)
			if se := kratoserrs.FromError(err); se != nil {
				code = se.Code
				reason = se.Reason
			}
			level, stack := extractError(err)
			log.Context(ctx).Log(level,
				"kind", "server",
				"component", kind,
				"operation", operation,
				"args", extractArgs(req),
				"json", extractArgs2Json(req),
				"result", extractArgs2Json(reply),
				"code", code,
				"reason", reason,
				"stack", stack,
				"cost", time.Since(startTime).Seconds(),
			)
			return
		}
	}
}
