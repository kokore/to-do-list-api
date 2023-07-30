package logz

import (
	"github.com/hugeman/todolist/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	Info  = "info"
	Debug = "debug"
)

var Logger *zap.Logger

func Init() error {
	if err := initLogger(); err != nil {
		return err
	}
	return nil
}

func initLogger() error {
	logger, err := newLogger()
	if err != nil {
		return err
	}

	Logger = logger

	return nil
}

func newLogger() (*zap.Logger, error) {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeTime = zapcore.ISO8601TimeEncoder

	c := zap.NewProductionConfig()
	c.EncoderConfig = ec
	c.Encoding = config.Config.Log.Format
	c.DisableStacktrace = true
	c.Level.SetLevel(parseLevel(config.Config.Log.Level))

	return c.Build()
}

func parseLevel(level string) zapcore.Level {
	switch level {
	default:
		return zapcore.InfoLevel
	case Debug:
		return zapcore.DebugLevel
	}
}
