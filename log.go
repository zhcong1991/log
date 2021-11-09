package log

func Debug(tag string, f string, v ...interface{}) {
	_logger.writeLog(DEBUG, tag, f, v...)
}

func Info(tag string, f string, v ...interface{}) {
	_logger.writeLog(INFO, tag, f, v...)
}

func Warn(tag string, f string, v ...interface{}) {
	_logger.writeLog(WARN, tag, f, v...)
}

func Error(tag string, f string, v ...interface{}) {
	_logger.writeLog(ERROR, tag, f, v...)
}

func Fatal(tag string, f string, v ...interface{}) {
	_logger.writeLog(FATAL, tag, f, v...)
}

func Init(cfg *Config) {
	_logger.Init(cfg)
}