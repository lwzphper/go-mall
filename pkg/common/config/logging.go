package config

import (
	"fmt"
	"github.com/lwzphper/go-mall/pkg/common/config/app"
	"github.com/lwzphper/go-mall/pkg/logger"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

func NewDefaultLogging() *Logging {
	return &Logging{
		Name:       "default",
		Driver:     "rotate",
		FileName:   "./log/logging.log",
		Level:      "info",
		MaxAge:     0,
		FileSizeMB: 100,
	}
}

type Logging struct {
	Name       string `toml:"name" yaml:"name" mapstructure:"name" env:"LOGGING_NAME"`
	Driver     string `toml:"driver" yaml:"driver" mapstructure:"driver" env:"LOGGING_DRIVER"`
	FileName   string `toml:"file_name" yaml:"file_name" mapstructure:"file_name" env:"LOGGING_FILE_NAME"`
	Level      string `toml:"level" yaml:"level" mapstructure:"level" env:"LOGGING_LEVEL"`
	FileSizeMB int    `toml:"file_size_mb" yaml:"file_size_mb" mapstructure:"file_size_mb" env:"LOGGING_FILE_SIZE_MB"`
	MaxAge     int    `toml:"max_age" yaml:"max_age" mapstructure:"max_age" env:"LOGGING_MAX_AGE"`
	MaxBackups int    `toml:"max_back_ups" yaml:"max_back_ups" mapstructure:"max_back_ups" env:"LOGGING_MAX_BACK_UPS"`
	Compress   bool   `toml:"compress" yaml:"compress" mapstructure:"compress" env:"LOGGING_COMPRESS"`
}

// InitLogger 初始化日志
func (l *Logging) InitLogger(env app.Env) *logger.Logger {
	level, err := zapcore.ParseLevel(l.Level)
	if err != nil {
		level = logger.InfoLevel // 如果解析出错，设置 info 级别
	}

	var log *logger.Logger
	var writer io.Writer
	// 日志切分
	if l.Driver == "rotate" {
		lCfg := logger.SizeRotateLogConfig{
			Level:      level,
			FileName:   l.FileName,
			MaxSize:    l.FileSizeMB,
			MaxAge:     l.MaxAge,
			MaxBackups: l.MaxBackups,
		}
		writer = logger.GetSizeLogWriter(lCfg.FileName, lCfg.MaxSize, lCfg.MaxBackups, lCfg.MaxAge, lCfg.Compress)
	} else { // 单个文件
		writer, err = os.OpenFile(l.FileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(fmt.Sprintf("cannot open log file: %v", err))
		}
	}

	// 如果开发环境，将错误输出到终端
	if env == app.EnvDevelopment {
		tops := []logger.TeeOption{
			{
				W: os.Stdout,
				Lef: func(lvl logger.Level) bool { // 将全部错误都输出到终端
					return true
				},
			},
			{
				W: writer,
				Lef: func(lvl logger.Level) bool {
					return lvl >= level
				},
			},
		}
		log = logger.NewTee(tops)
	} else {
		log = logger.New(writer, level)
	}

	log.Name(l.Name)

	logger.ResetDefault(log)
	return log
}
