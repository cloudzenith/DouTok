package middlewares

import (
	"context"
	"errors"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/constants"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/google/uuid"
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

			traceId := meta.Get(constants.TraceIdHeadersKey)
			if traceId == "" {
				traceId, err = generateId()
				if err != nil {
					return nil, err
				}

				metadata.AppendToClientContext(ctx, constants.TraceIdHeadersKey, traceId)
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
