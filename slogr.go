package slogr

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strings"
)

type SLogger struct {
	Level  string `env:""`
	Format string `env:""`

	logr *slog.Logger
}

type sloggerKeyType int

var sloggerKey = sloggerKeyType(0)

func WithContext(ctx context.Context, log *SLogger) context.Context {
	return context.WithValue(ctx, sloggerKey, log)
}

func FromContext(ctx context.Context) *SLogger {
	return ctx.Value(sloggerKey).(*SLogger)
}

func (s *SLogger) SetDefaults() {
	if s.Level == "" {
		s.Level = "info"
	}

	if s.Format == "" {
		s.Format = "json"
	}
}

func (s *SLogger) Initialize() {
	s.SetDefaults()

	opt := &slog.HandlerOptions{
		Level: Leveler(strings.ToLower(s.Level)).Leveler(),
	}

	var h *slog.JSONHandler
	if s.Format == "json" {
		h = slog.NewJSONHandler(os.Stdout, opt)
	}

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
