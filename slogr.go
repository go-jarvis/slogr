package slogr

import "context"

// Logger is the interface for slogr
// It provides `With` function to create a new logger with the given key-value pairs
// It provides `Debug`, `Info`, `Warn`, `Error` functions to log messages
type Logger interface {
	With(kv ...any) Logger
	WithContext(ctx context.Context, ky ...any) (context.Context, Logger)
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
}

type loggerKeyType int

var loggerKey = loggerKeyType(0)

// WithContext returns a new context with the given logger
func WithContext(ctx context.Context, log Logger) context.Context {
	return context.WithValue(ctx, loggerKey, log)
}

// FromContext returns the logger from the given context
// if none is found, it returns the default logger: json format, info level
// if the given context is nil or not Logger in context, default it returns a Discard logger, it does not log anything
func FromContext(ctx context.Context) Logger {
	if ctx == nil {
		return &Discard{}
	}

	log, ok := ctx.Value(loggerKey).(Logger)
	if ok {
		return log
	}

	return &Discard{}
}
