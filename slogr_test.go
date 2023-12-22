package slogr

import (
	"log/slog"
	"os"
	"testing"
)

func TestXxx(t *testing.T) {
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

func TestSLogger(t *testing.T) {
	log := &SLogger{
		Level:  "info",
		Format: "json",
	}
	log.Initialize()
	log = log.With("foo", "bar")
	log.Debug("hello world")
	log.Info("hello world")
	log.Warn("hello world")
	log.Error("hello world")
}
