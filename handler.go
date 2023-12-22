package slogr

import (
	"log/slog"
	"os"
	"strings"
)

// Handler returns a new slog.Handler based on the given format and level
func Handler(format string, level string) slog.Handler {

	opt := &slog.HandlerOptions{
		Level: Leveler(strings.ToLower(level)).Leveler(),
	}

	switch strings.ToLower(format) {

	case "json":
		return slog.NewJSONHandler(os.Stdout, opt)
	case "text":
		return slog.NewTextHandler(os.Stdout, opt)
	}

	return slog.NewJSONHandler(os.Stdout, opt)

}
