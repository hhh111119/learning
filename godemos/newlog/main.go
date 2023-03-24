package main

import (
	"context"
	"net"
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr)))
	slog.Info("hello", "name", "Al")
	slog.Error("oops", "err", net.ErrClosed, "status", 500)
	slog.LogAttrs(context.Background(), slog.LevelError, "oops",
		slog.Any("err", net.ErrClosed), slog.Int("status", 500))

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr)))
	slog.Info("hello", "name", "Al")
	slog.Error("oops", "err", net.ErrClosed, "status", 500)
	slog.LogAttrs(context.Background(), slog.LevelError, "oops",
		slog.Any("err", net.ErrClosed), slog.Int("status", 500))
}
