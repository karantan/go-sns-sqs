// Package logger ...
package logger

import (
	"log"
	"os"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New creates package specific logging pipeline.
func New(name string) *zap.SugaredLogger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	level := zap.InfoLevel

	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level)
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()

	return logger.Named(name).Sugar()
}

func NewSentry(dsn string) error {

	err := sentry.Init(sentry.ClientOptions{
		Dsn:                dsn,
		EnableTracing:      true,
		TracesSampleRate:   1.0,
		ProfilesSampleRate: 1.0,
		Debug:              true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	return nil
}
