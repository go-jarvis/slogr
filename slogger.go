package slogr

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/go-jarvis/slogr/handler"
)

// SLogger is a wrapper around slog.Logger
// It provides a default logger with json format and info level
// It also provides `FromContext` function to create a new logger with the given key-value pairs
// It also provides `WithContext` function to get the logger from the context
//
// Level: debug,info,warn,error
// Format: json,text
type SLogger struct {
	Level  string `env:""`
	Format string `env:""`

	logr *slog.Logger
}

var _ Logger = &SLogger{}

// Default returns a new `SLogger` with the default `info` level and `json` format
func Default() Logger {
	return New("debug", "json")
}

// New returns a new `SLogger` with the given level and format
func New(level string, format string) Logger {
	log := &SLogger{
		Level:  level,
		Format: format,
	}
	log.Initialize()

	return log
}

// SetDefaults sets the default values for the logger
func (s *SLogger) SetDefaults() {
	if s.Level == "" {
		s.Level = "debug"
	}

	if s.Format == "" {
		s.Format = "json"
	}
}

// Initialize initializes the logger
func (s *SLogger) Initialize() {
	s.SetDefaults()

	h := handler.New(s.Format, s.Level)
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

func (s *SLogger) With(kvs ...any) Logger {
	return s.with(kvs...)
}

// With returns a new logger with the given key-value pairs
func (s *SLogger) with(kvs ...any) *SLogger {
	logr := s.logr.With(kvs...)
	s.logr = logr

	return s
}

// caller returns the caller of the file and function that called it
func caller() string {
	pc, file, line, _ := runtime.Caller(4)
	funcName := runtime.FuncForPC(pc).Name()

	file = filepath.Base(file) // filename only

	parts := strings.Split(funcName, ".")
	funcName = parts[len(parts)-1]

	return fmt.Sprintf("%s:%d#%s", file, line, funcName)
}
