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
	ResetDefault(logger)
	defer Sync()

	// 使用包级函数调用
	Info("file log info")
	Warn("file log warn")
	Error("file log error")
}

// 测试多文件写入。 warn 级别及以下的日志写入 warn.log； warn 级别及以上的写入 error.log
func TestNewTee(t *testing.T) {
	infoFile, err := os.OpenFile("./log/info.log", logFlag, 0644)
	if err != nil {
		panic(err)
	}

	errorFile, err := os.OpenFile("./log/error.log", logFlag, 0644)
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
				return InfoLevel < lvl
			},
		},
	}

	logger := NewTee(tops)
	ResetDefault(logger)
	Debug("file log Debug")
	Info("file log Info")
	Warn("file log Warn")
	Error("file log Error")
}
