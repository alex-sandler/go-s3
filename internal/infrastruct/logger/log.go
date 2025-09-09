package logger

import (
	"context"
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	global       *zap.SugaredLogger
	defaultLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
)

func init() {
	global = New(defaultLevel, os.Stdout)
}

func New(level zapcore.LevelEnabler, w io.Writer, options ...zap.Option) *zap.SugaredLogger {
	if level == nil {
		level = defaultLevel
	}

	cfg := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	enc := zapcore.NewJSONEncoder(cfg)

	return zap.New(zapcore.NewCore(enc, zapcore.AddSync(w), level), options...).Sugar()
}

type ctxLoggerKey struct{}

func WithLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, ctxLoggerKey{}, logger)
}

func FromContext(ctx context.Context) *zap.SugaredLogger {
	l := ctx.Value(ctxLoggerKey{})
	if l != nil {
		return l.(*zap.SugaredLogger)
	}

	return global
}
