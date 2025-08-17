package applog

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/hirukiyo/gin-sample/internal/app"

	"github.com/MatusOllah/slogcolor"
	"github.com/gin-gonic/gin"
)

func SetLogger(cfg *app.AppEnvironment) {
	var h slog.Handler
	logLevel := AppLogLevel(cfg)
	if cfg.AppMode == gin.ReleaseMode {
		h = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: logLevel})
	} else {
		h = slogcolor.NewHandler(os.Stdout, &slogcolor.Options{
			Level:       logLevel,
			TimeFormat:  time.RFC3339,
			SrcFileMode: slogcolor.ShortFile,
		})
	}
	slog.SetDefault(slog.New(h))
}

func AppLogLevel(cfg *app.AppEnvironment) slog.Level {
	appLogLevel := slog.LevelInfo
	switch strings.ToLower(cfg.AppLogLevel) {
	case "debug":
		appLogLevel = slog.LevelDebug
	case "info":
		appLogLevel = slog.LevelInfo
	case "warn":
		appLogLevel = slog.LevelWarn
	case "error":
		appLogLevel = slog.LevelError
	}
	return appLogLevel
}

func Debug(ctx context.Context, msg string, args ...any) {
	log(ctx, slog.LevelDebug, msg, args...)
}

func Info(ctx context.Context, msg string, args ...any) {
	log(ctx, slog.LevelInfo, msg, args...)
}

func Warn(ctx context.Context, msg string, args ...any) {
	log(ctx, slog.LevelWarn, msg, args...)
}

func Error(ctx context.Context, msg string, args ...any) {
	log(ctx, slog.LevelError, msg, args...)
}

func log(ctx context.Context, level slog.Level, msg string, args ...any) {
	logger := slog.Default()
	if !logger.Enabled(ctx, level) {
		return
	}

	var pcs [1]uintptr
	runtime.Callers(3, pcs[:]) // Skip(3): Callers, log, Debug[f]/Info[f]/Warn[f]/Error[f]/Fatal[f]

	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	r.Add(args...)

	_ = logger.Handler().Handle(ctx, r)
}
