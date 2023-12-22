package slogr

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"
	"strings"
)

type SLogger struct {
	Level  string `env:""`
	Format string `env:""`

	logr *slog.Logger
}

func Default() *SLogger {
	log := &SLogger{}
	log.Initialize()

	return log
}

type sloggerKeyType int

var sloggerKey = sloggerKeyType(0)

// WithContext returns a new context with the given logger
func WithContext(ctx context.Context, log *SLogger) context.Context {
	return context.WithValue(ctx, sloggerKey, log)
}

// FromContext returns the logger from the given context
// if none is found, it returns the default logger: json format, info level
func FromContext(ctx context.Context) *SLogger {
	log, ok := ctx.Value(sloggerKey).(*SLogger)
	if ok {
		return log
	}

	return Default()
}

// SetDefaults sets the default values for the logger
func (s *SLogger) SetDefaults() {
	if s.Level == "" {
		s.Level = "info"
	}

	if s.Format == "" {
		s.Format = "json"
	}
}

// Initialize initializes the logger
func (s *SLogger) Initialize() {
	s.SetDefaults()

	h := Handler(s.Format, s.Level)
	s.logr = slog.New(h)
}

// output print logs a message with the given level
func (s *SLogger) output(l slog.Level, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)

	caller := caller()
	s.logr.Log(context.Background(), l, msg, "caller", caller)
}

func (s *SLogger) Debug(format string, args ...interface{}) {
	s.output(slog.LevelDebug, format, args...)
}

func (s *SLogger) Info(format string, args ...interface{}) {
	s.output(slog.LevelInfo, format, args...)
}

func (s *SLogger) Warn(format string, args ...interface{}) {
	s.output(slog.LevelWarn, format, args...)
}

func (s *SLogger) Error(format string, args ...interface{}) {
	s.output(slog.LevelError, format, args...)
}

// With returns a new logger with the given key-value pairs
func (s *SLogger) With(kvs ...any) *SLogger {
	logr := s.logr.With(kvs...)
	s.logr = logr

	return s
}

// caller returns the caller of the file and function that called it
func caller() string {
	pc, file, line, _ := runtime.Caller(4)
	funcName := runtime.FuncForPC(pc).Name()

	// file = filepath.Base(file) // filename only

	parts := strings.Split(funcName, ".")
	funcName = parts[len(parts)-1]

	return fmt.Sprintf("%s:%d#%s", file, line, funcName)
}
