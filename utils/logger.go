package utils

import (
	"log/slog"
	"os"
)

func InitLogger() (err error) {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   true,            // 记录日志位置
		Level:       slog.LevelDebug, // 设置日志级别
		ReplaceAttr: nil,
	}))
	slog.SetDefault(l)
	return

}
