package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

type Level = zapcore.Level
type Field = zap.Field
type Option = zap.Option

type RotateOptions struct {
	FileName   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
}

type Logger struct {
	l             *zap.Logger
	RotateOptions RotateOptions
	level         Level
}

// New create a new logger (not support log rotating).
func New(writer io.Writer, level Level, opts ...Option) *Logger {
	if writer == nil {
		panic("the writer is nil")
	}
	cfg := getProductionCfg()

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(writer),
		level,
	)
	return &Logger{
		l:     zap.New(core, opts...),
		level: level,
	}
}

// NewWithRotate create a new logger support log rotating.
func NewWithRotate(writer io.Writer, level Level, logRotate RotateOptions, opts ...Option) *Logger {
	if writer == nil {
		panic("the writer is nil")
	}
	cfg := getProductionCfg()

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   top.Filename,
		MaxSize:    top.Ropt.MaxSize,
		MaxBackups: top.Ropt.MaxBackups,
		MaxAge:     top.Ropt.MaxAge,
		Compress:   top.Ropt.Compress,
	})

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(w),
		level,
	)
	return &Logger{
		l:     zap.New(core, opts...),
		level: level,
	}
}

// 参考：https://www.cnblogs.com/Me1onRind/p/10918863.html
func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		filename+".%Y%m%d%H", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}
	return hook
}

// 获取配置项
func getProductionCfg() zap.Config {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	return cfg
}

var std = New(os.Stderr, InfoLevel)

// 使用method value语法将std实例的各个方法以包级函数的形式暴露给用户，简化用户对logger实例的获取
var (
	Info   = std.Info
	Warn   = std.Warn
	Error  = std.Error
	DPanic = std.DPanic
	Panic  = std.Panic
	Fatal  = std.Fatal
	Debug  = std.Debug
	Sync   = std.Sync
)

// ResetDefault not safe for concurrent use
func ResetDefault(l *Logger) {
	std = l
	Info = std.Info
	Warn = std.Warn
	Error = std.Error
	DPanic = std.DPanic
	Panic = std.Panic
	Fatal = std.Fatal
	Debug = std.Debug
}

func Default() *Logger {
	return std
}
