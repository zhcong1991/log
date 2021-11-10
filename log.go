package log

import "context"

func Debug(tag string, f string, v ...interface{}) {
	_logger.writeLog(DEBUG, tag, f, v...)
}

func DebugCtx(ctx context.Context, tag string, f string, v ...interface{}) {
	_logger.writeLogCtx(DEBUG, ctx, tag, f, v...)
}

func Info(tag string, f string, v ...interface{}) {
	_logger.writeLog(INFO, tag, f, v...)
}

func InfoCtx(ctx context.Context, tag string, f string, v ...interface{}) {
	_logger.writeLogCtx(INFO, ctx, tag, f, v...)
}

func Warn(tag string, f string, v ...interface{}) {
	_logger.writeLog(WARN, tag, f, v...)
}

func WarnCtx(ctx context.Context, tag string, f string, v ...interface{}) {
	_logger.writeLogCtx(WARN, ctx, tag, f, v...)
}

func Error(tag string, f string, v ...interface{}) {
	_logger.writeLog(ERROR, tag, f, v...)
}

func ErrorCtx(ctx context.Context, tag string, f string, v ...interface{}) {
	_logger.writeLogCtx(ERROR, ctx, tag, f, v...)
}

func Fatal(tag string, f string, v ...interface{}) {
	_logger.writeLog(FATAL, tag, f, v...)
}

func FatalCtx(ctx context.Context, tag string, f string, v ...interface{}) {
	_logger.writeLogCtx(FATAL, ctx, tag, f, v...)
}

func Init(cfg *Config) {
	_logger.Init(cfg)
}

func Close() {
	_logger.Close()
}