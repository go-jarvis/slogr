package slogr

import (
	"context"
	"log/slog"
	"os"
	"testing"
)

// TestOriginLogger is a test for slog
func TestOriginLogger(t *testing.T) {
	opt := &slog.HandlerOptions{
		// Level: slog.LevelDebug,
		Level: slog.LevelWarn,
	}
	h := slog.NewJSONHandler(os.Stdout, opt)
	log := slog.New(h)
	log = log.With("foo", "bar")
	log.Debug("hello world")
	log.Info("hello world")
	log.Warn("hello world")
	log.Error("hello world")
}

// TestSLogger is a test for SLogger
func TestSLogger(t *testing.T) {
	log := New("debug", "json")
	log = log.With("foo", "bar")
	output(log)
}

func TestLoggerFromContext(t *testing.T) {
	ctx := context.Background()
	log := FromContext(ctx)

	output(log)
}

func output(log Logger) {
	log = log.With("in", "output")
	log.Debug("hello world")
	log.Info("hello world")
	log.Warn("hello world")
	log.Error("hello world")
}

// {"time":"2023-12-22T15:44:48.455032+08:00","level":"DEBUG","msg":"hello world","foo":"bar","in":"output","caller":"/Users/franktang/data/gopath/src/github.com/go-jarvis/slogr/slogr_test.go:31#TestSLogger"}
// {"time":"2023-12-22T15:44:48.455347+08:00","level":"INFO","msg":"hello world","foo":"bar","in":"output","caller":"/Users/franktang/data/gopath/src/github.com/go-jarvis/slogr/slogr_test.go:31#TestSLogger"}
// {"time":"2023-12-22T15:44:48.455352+08:00","level":"WARN","msg":"hello world","foo":"bar","in":"output","caller":"/Users/franktang/data/gopath/src/github.com/go-jarvis/slogr/slogr_test.go:31#TestSLogger"}
// {"time":"2023-12-22T15:44:48.455356+08:00","level":"ERROR","msg":"hello world","foo":"bar","in":"output","caller":"/Users/franktang/data/gopath/src/github.com/go-jarvis/slogr/slogr_test.go:31#TestSLogger"}

func TestSLogger2(t *testing.T) {

	log := FromContext(context.TODO())
	// log.Initialize()
	log = log.With("foo", "bar")

	output(log)
}
