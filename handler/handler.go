package handler

import (
	"log/slog"
	"os"
	"strings"

	"github.com/go-jarvis/slogr/leveler"
)

// New returns a new slog.New based on the given format and level
func New(format string, level string) slog.Handler {

	opt := &slog.HandlerOptions{
		Level: leveler.Parse(level),
	}

	switch strings.ToLower(format) {

	case "json":
		return slog.NewJSONHandler(os.Stdout, opt)
	case "text":
		return slog.NewTextHandler(os.Stdout, opt)
	}

	return slog.NewJSONHandler(os.Stdout, opt)
}
