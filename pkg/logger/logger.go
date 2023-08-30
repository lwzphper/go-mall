package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

type Level = zapcore.Level
type Field = zap.Field
type Option = zap.Option

type Logger struct {
	L     *zap.Logger
	level Level
}

// New create a new logger (not support log rotating).
func New(writer io.Writer, level Level, opts ...Option) *Logger {
	if writer == nil {
		panic("the writer is nil")
	}

	core := zapcore.NewCore(
		getEncoder(),
		zapcore.AddSync(writer),
		level,
	)
	opts = append(opts, zap.AddCaller())
	return &Logger{
		L:     zap.New(core, opts...),
		level: level,
	}
}

// SizeRotateLogConfig 日志大小分割
type SizeRotateLogConfig struct {
	Level      Level  `json:"level"`       // Level 最低日志等级
	FileName   string `json:"file_name"`   // FileName 日志文件位置
	MaxSize    int    `json:"max_size"`    // MaxSize 进行切割之前，日志文件的最大大小(MB为单位)，默认为100MB
	MaxAge     int    `json:"max_age"`     // MaxAge 是根据文件名中编码的时间戳保留旧日志文件的最大天数。
	MaxBackups int    `json:"max_backups"` // MaxBackups 是要保留的旧日志文件的最大数量。默认是保留所有旧的日志文件（尽管 MaxAge 可能仍会导致它们被删除。）
	Compress   bool   `json:"compress"`    // Compress 是否压缩保存
}

// NewWithSizeRotate create a new logger support log rotating.
func NewWithSizeRotate(lCfg SizeRotateLogConfig, opts ...Option) *Logger {
	opts = append(opts, zap.AddCaller())
	w := GetSizeLogWriter(lCfg.FileName, lCfg.MaxSize, lCfg.MaxBackups, lCfg.MaxAge, lCfg.Compress)
	core := zapcore.NewCore(
		getEncoder(),
		zapcore.AddSync(w),
		lCfg.Level,
	)

	//logger := zap.New(core, opts...)
	//zap.ReplaceGlobals(logger) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可

	return &Logger{
		L:     zap.New(core, opts...),
		level: lCfg.Level,
	}
}

// GetSizeLogWriter 负责日志写入的位置
func GetSizeLogWriter(filename string, maxsize, maxBackup, maxAge int, compress bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,  // 文件位置
		MaxSize:    maxsize,   // 进行切割之前,日志文件的最大大小(MB为单位)
		MaxAge:     maxAge,    // 保留旧文件的最大天数
		MaxBackups: maxBackup, // 保留旧文件的最大个数
		Compress:   compress,  // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

type LevelEnablerFunc func(lvl Level) bool

type TeeOption struct {
	W   io.Writer
	Lef LevelEnablerFunc
}

// NewTee 创建写多个log文件的logger。根据不同的日志级别，写入不同的文件
func NewTee(tops []TeeOption, opts ...Option) *Logger {
	// 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
	opts = append(opts, zap.AddCaller())

	var cores []zapcore.Core
	for _, top := range tops {
		if top.W == nil {
			panic("the writer is nil")
		}

		// 使用闭包主要解决：go 闭包变量延迟绑定问题
		// golang的for range机制相当于对for循环做了优化，会额外创建一个新的 v2 变量存储切片中的元素，
		// 循环中使用的这个变量 v2的值 会在每一次迭代被重新赋值而覆盖，赋值时也会触发拷贝，而其本身地址不会变，
		// 因为始终是同一变量；也就引出了另一个需要注意的for range用法中的指针问题
		// blog link: https://blog.csdn.net/qq_39618369/article/details/121546942
		// link: https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-for-range/
		lvl := func(top TeeOption) zapcore.LevelEnabler {
			return zap.LevelEnablerFunc(func(level zapcore.Level) bool {
				return top.Lef(level)
			})
		}(top)

		core := zapcore.NewCore(getEncoder(), zapcore.AddSync(top.W), lvl)
		cores = append(cores, core)
	}
	return &Logger{
		L: zap.New(zapcore.NewTee(cores...), opts...),
	}
}

// 获取配置项
func getEncoder() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	encodeConfig.TimeKey = "time"
	// 将Level序列化为全大写字符串。例如，将info level序列化为INFO。
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 以 package/file:行 的格式 序列化调用程序，从完整路径中删除除最后一个目录外的所有目录。
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}

var std = New(os.Stderr, InfoLevel)

// 使用method value语法将std实例的各个方法以包级函数的形式暴露给用户，简化用户对logger实例的获取
var (
	Debug  = std.Debug
	Info   = std.Info
	Warn   = std.Warn
	Error  = std.Error
	DPanic = std.DPanic
	Panic  = std.Panic
	Fatal  = std.Fatal
	Sync   = std.Sync
)

// ResetDefault not safe for concurrent use
func ResetDefault(l *Logger) {
	std = l
	Debug = std.Debug
	Info = std.Info
	Warn = std.Warn
	Error = std.Error
	DPanic = std.DPanic
	Panic = std.Panic
	Fatal = std.Fatal
}

func NewDefaultLogger() *Logger {
	return std
}
