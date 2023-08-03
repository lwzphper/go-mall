package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"time"
)

type LevelEnablerFunc func(lvl Level) bool

type TeeOption struct {
	W   io.Writer
	Lef LevelEnablerFunc
}

// NewTee 创建写多个log文件的logger。根据不同的日志级别，写入不同的文件
func NewTee(tops []TeeOption, opts ...Option) *Logger {
	var cores []zapcore.Core
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}

	for _, top := range tops {
		if top.W == nil {
			panic("the writer is nil")
		}

		lv := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return top.Lef(lvl)
		})

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg.EncoderConfig),
			zapcore.AddSync(top.W),
			lv,
		)
		cores = append(cores, core)
	}
	return &Logger{
		l: zap.New(zapcore.NewTee(cores...), opts...),
	}
}
