package logger

import (
	"blog-service/constants"
	"context"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

// AppLogger extends the logrus.Logger with additional functionality.
type AppLogger struct {
	*logrus.Logger
}

// NewAppLogger initializes a new AppLogger.
func NewAppLogger(level logrus.Level) *AppLogger {
	logger := logrus.New()
	logger.SetFormatter(
		&logrus.JSONFormatter{
			TimestampFormat:  "02 Jan 06 15:04:04 -0700",
			DisableTimestamp: false,
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "time",
				logrus.FieldKeyLevel: "level",
				logrus.FieldKeyMsg:   "msg",
				logrus.FieldKeyFunc:  "@caller",
				logrus.FieldKeyFile:  "file",
			},
			PrettyPrint: false,
		},
	)
	logger.SetOutput(os.Stdout)
	logger.SetLevel(level) // Set default log level
	return &AppLogger{logger}
}

// WithContext adds the request ID from the context to the log entry.
func (a *AppLogger) WithContext(ctx context.Context) *logrus.Entry {
	// Default to "unknown" if the request ID is not found
	requestID := "unknown"
	if id, ok := ctx.Value(constants.RequestIDKey).(string); ok {
		requestID = id
	}

	// Create a log entry with the request ID
	entry := a.Logger.WithField(string(constants.RequestIDKey), requestID)

	// Add caller information
	_, file, line, ok := runtime.Caller(2) // 2 to skip this function and the caller
	if ok {
		entry = entry.WithField("file", file).WithField("line", line)
	}

	return entry
}

// Info logs an info level message.
func (a *AppLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	a.WithContext(ctx).Infof(msg, args...)
}

// Warn logs a warning level message.
func (a *AppLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	a.WithContext(ctx).Warnf(msg, args...)
}

// Error logs an error level message.
func (a *AppLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	a.WithContext(ctx).Errorf(msg, args...)
}

// Debug logs a debug level formatted message.
func (a *AppLogger) Debug(ctx context.Context, format string, args ...interface{}) {
	a.WithContext(ctx).Debugf(format, args...)
}

// WithField adds a single field to the log entry.
func (a *AppLogger) WithField(key string, value interface{}) *logrus.Entry {
	return a.Logger.WithField(key, value)
}

// WithFields adds multiple fields to the log entry.
func (a *AppLogger) WithFields(fields logrus.Fields) *logrus.Entry {
	return a.Logger.WithFields(fields)
}

// SetLevel allows dynamic adjustment of log level.
//func SetLevel(level logrus.Level) {
//	Log.SetLevel(level)
//}

// Initialize the global logger

var Log = NewAppLogger(logrus.InfoLevel)
