package leveler

import (
	"fmt"
	"log/slog"
	"strings"
)

// type Leveler string

// var (
// 	LevelDebug Leveler = "debug"
// 	LevelInfo  Leveler = "info"
// 	LevelWarn  Leveler = "warn"
// 	LevelError Leveler = "error"
// 	LevelFatal Leveler = "fatal"
// )

// // Leveler returns the slog.Level for the given level
// func (l Leveler) String() string {
// 	switch l {
// 	case LevelDebug:
// 		return "debug"
// 	case LevelInfo:
// 		return "info"
// 	case LevelWarn:
// 		return "warn"
// 	case LevelError:
// 		return "error"
// 		// case LevelFatal:
// 		// 	return "fatal"
// 	}

// 	return "unknown"
// }

// func (l Leveler) SlogLevel() slog.Level {
// 	switch l {
// 	case LevelDebug:
// 		return slog.LevelDebug
// 	case LevelInfo:
// 		return slog.LevelInfo
// 	case LevelWarn:
// 		return slog.LevelWarn
// 	case LevelError:
// 		return slog.LevelError
// 		// case LevelFatal:
// 		// 	return slog.LevelFatal
// 	}

// 	return slog.LevelDebug
// }

func Parse(lvl string) slog.Level {

	switch strings.ToLower(lvl) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	}

	fmt.Printf("Unsupported level, defaulting to info: %s\n", lvl)
	return slog.LevelInfo
}
