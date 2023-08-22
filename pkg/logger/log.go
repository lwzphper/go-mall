package logger

func (l *Logger) Name(name string) {
	l.L.Named(name)
}

func (l *Logger) Debug(msg string, fields ...Field) {
	l.L.Debug(msg, fields...)
}

func (l *Logger) Debugf(msg string, args ...interface{}) {
	l.L.Sugar().Debugf(msg, args)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.L.Info(msg, fields...)
}

func (l *Logger) Infof(msg string, args ...interface{}) {
	l.L.Sugar().Infof(msg, args)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.L.Warn(msg, fields...)
}

func (l *Logger) Warnf(msg string, args ...interface{}) {
	l.L.Sugar().Warnf(msg, args)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.L.Error(msg, fields...)
}

func (l *Logger) Errorf(msg string, args ...interface{}) {
	l.L.Sugar().Errorf(msg, args)
}

func (l *Logger) DPanic(msg string, fields ...Field) {
	l.L.DPanic(msg, fields...)
}

func (l *Logger) DPanicf(msg string, args ...interface{}) {
	l.L.Sugar().DPanicf(msg, args)
}

func (l *Logger) Panic(msg string, fields ...Field) {
	l.L.Panic(msg, fields...)
}

func (l *Logger) Panicf(msg string, args ...interface{}) {
	l.L.Sugar().Panicf(msg, args)
}

func (l *Logger) Fatal(msg string, fields ...Field) {
	l.L.Fatal(msg, fields...)
}

func (l *Logger) Fatalf(msg string, args ...interface{}) {
	l.L.Sugar().Fatalf(msg, args)
}

// Sync 调用内核的Sync方法，刷新所有缓冲的日志条目
// 应用程序应该注意在退出之前调用Sync
func (l *Logger) Sync() error {
	return l.L.Sync()
}
