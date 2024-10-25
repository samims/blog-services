package logger

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

// init initialize the logger with JSON formatting, stdout, and info level
func init() {
	Log = logrus.New()
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.InfoLevel)
}

func WithContext(ctx context.Context) *logrus.Entry {
	if requestID, ok := ctx.Value("request_id").(string); ok {
		return Log.WithField("request_id", requestID)
	}
	return Log.WithField("request_id", "unknown")
}

// SetLevel allows dynamic adjustment of log level
func SetLevel(level logrus.Level) {
	Log.SetLevel(level)
}

// Info logs an info level message
func Info(ctx context.Context, msg string, args ...interface{}) {
	WithContext(ctx).Infof(msg, args...)
}

// Warn logs a warning level message
func Warn(ctx context.Context, msg string, args ...interface{}) {
	WithContext(ctx).Warnf(msg, args...)
}

// Error logs an error level message
func Error(ctx context.Context, msg string, args ...interface{}) {
	WithContext(ctx).Errorf(msg, args...)
}

// Debug logs a debug level formatted message
func Debug(ctx context.Context, format string, args ...interface{}) {
	WithContext(ctx).Debugf(format, args...)
}

// WithField adds a single field to the log entry
func WithField(key string, value interface{}) *logrus.Entry {
	return Log.WithField(key, value)
}

// WithFields adds multiple fields to the log entry
func WithFields(fields logrus.Fields) *logrus.Entry {
	return Log.WithFields(fields)
}
