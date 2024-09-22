package utils

import (
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/middlewares"
	kratoszap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

const (
	LogPath       = "./logs/"
	LogFileName   = "log.log"
	LogMaxSize    = 5
	LogMaxBackups = 5
	LogMaxAge     = 30
	LogCompress   = false
)

func getJsonLogEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogSyncWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   LogPath + LogFileName,
		MaxSize:    LogMaxSize,
		MaxBackups: LogMaxBackups,
		MaxAge:     LogMaxAge,
		Compress:   LogCompress,
	}
	return zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(lumberJackLogger),
		zapcore.AddSync(os.Stdout),
	)
}

func SetJsonLogger() log.Logger {
	writeSyncer := getLogSyncWriter()
	encoder := getJsonLogEncoder()

	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	z := zap.New(core)
	tracing.Server()
	zapLogger := kratoszap.NewLogger(z)
	logger := log.With(
		zapLogger,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
		"x_trace_id", middlewares.GetTraceId(),
		"x_span_id", middlewares.GetSpanId(),
	)
	log.SetLogger(logger)

	return logger
}
