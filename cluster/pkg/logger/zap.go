package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger creates a zap logger instance.
func NewLogger(level, out string) *zap.Logger {
	var lvl zapcore.Level

	// set zap logger level
	if err := lvl.Set(level); err != nil {
		log.Printf("cannot parse log level %s: %s", level, err)

		lvl = zapcore.WarnLevel
	}

	// open `out` file for logs export
	file, err := os.OpenFile(out, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil
	}

	// create zapcore components
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	fileCore := zapcore.NewCore(encoder, zapcore.AddSync(file), lvl)
	cores := []zapcore.Core{
		fileCore,
	}

	// create new zap logger instance
	core := zapcore.NewTee(cores...)

	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
}
