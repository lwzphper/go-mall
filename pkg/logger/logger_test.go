package logger

import (
	"os"
	"testing"
)

var logFlag = os.O_CREATE | os.O_APPEND | os.O_WRONLY

// 测试终端输出
func TestStdout(t *testing.T) {
	defer Sync()
	Info("test string", String("test-key", "test-value"))
}

// 测试文件日志
func TestFileLog(t *testing.T) {
	file, err := os.OpenFile("./log/test.log", logFlag, 0644)
	if err != nil {
		t.Errorf("cannot open log file: %v", err)
		return
	}

	logger := New(file, WarnLevel)
	defer logger.Sync()

	// 使用包级函数调用
	logger.Info("file log info")
	logger.Warn("file log warn")
	logger.Error("file log error")
}

// 测试多文件写入。 warn 级别及以下的日志写入 warn.log； warn 级别及以上的写入 error.log
func TestNewTee(t *testing.T) {
	infoFile, err := os.OpenFile("./log/tee_info.log", logFlag, 0644)
	if err != nil {
		panic(err)
	}

	errorFile, err := os.OpenFile("./log/tee_error.log", logFlag, 0644)
	if err != nil {
		panic(err)
	}

	var tops = []TeeOption{
		{
			W: infoFile,
			Lef: func(lvl Level) bool {
				return lvl <= InfoLevel
			},
		},
		{
			W: errorFile,
			Lef: func(lvl Level) bool {
				return lvl > InfoLevel
			},
		},
	}

	logger := NewTee(tops)
	logger.Debug("file log Debug")
	logger.Info("file log Info")
	logger.Warn("file log Warn")
	logger.Error("file log Error")
}

// 测试日志分割
func TestNewWithSizeRotate(t *testing.T) {
	cfg := SizeRotateLogConfig{
		Level:      DebugLevel,
		FileName:   "./log/size_rotate.log",
		MaxSize:    1,
		MaxAge:     10,
		MaxBackups: 30,
	}
	logger := NewWithSizeRotate(cfg)
	for i := 0; i < 100000; i++ {
		logger.Debug("debug msg", String("rotate", "debug"))
		logger.Info("info msg", String("rotate", "info"))
		logger.Warn("warn msg", String("rotate", "warn"))
		logger.Error("error msg", String("rotate", "error"))
	}
}
