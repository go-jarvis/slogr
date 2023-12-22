package slogr

import (
	"context"
	"fmt"
	"log/slog"
	"os"
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

func (s *SLogger) log(l slog.Level, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	s.logr.Log(context.Background(), l, msg)
}

func (s *SLogger) Debug(format string, args ...interface{}) {
	s.log(slog.LevelDebug, format, args...)
}

func (s *SLogger) Info(format string, args ...interface{}) {
	s.log(slog.LevelInfo, format, args...)
}

func (s *SLogger) Warn(format string, args ...interface{}) {
	s.log(slog.LevelWarn, format, args...)
}

func (s *SLogger) Error(format string, args ...interface{}) {
	s.log(slog.LevelError, format, args...)
}

func (s *SLogger) With(kvs ...any) *SLogger {
	logr := s.logr.With(kvs...)
	s.logr = logr

	return s
}
