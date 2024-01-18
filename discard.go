package slogr

import "context"

// Discard returns a new Discard Logger with the given key-value pairs
// Discard Logger does not log anything
type Discard struct{}

func (d *Discard) With(kvs ...any) Logger {
	return d
}
func (d *Discard) WithContext(ctx context.Context, kvs ...any) (context.Context, Logger) {

	return ctx, d
}

func (d *Discard) Debug(format string, args ...interface{}) {}
func (d *Discard) Info(format string, args ...interface{})  {}
func (d *Discard) Warn(format string, args ...interface{})  {}
func (d *Discard) Error(format string, args ...interface{}) {}
