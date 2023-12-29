package main

import (
	"context"

	"github.com/go-jarvis/slogr"
)

func main() {

	ctx := context.Background()

	// 打印 json 日志格式
	log := slogr.New("debug", "json")
	log.With("logger", "slogger")
	ctx = slogr.WithContext(ctx, log)
	output(ctx)

	// 屏蔽所有日志
	logd := &slogr.Discard{}
	ctx = slogr.WithContext(ctx, logd)
	output(ctx)
}

func output(ctx context.Context) {
	log := slogr.FromContext(ctx)

	log = log.With("in", "output")
	log.Debug("hello world")
	log.Info("hello world")
	log.Warn("hello world")
	log.Error("hello world")
}
