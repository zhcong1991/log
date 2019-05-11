package log

func Debug(tag string, v ...interface{}) {
	logger.writeLog(LogLevel_Debug, tag, v...)
}

func Debugf(tag string, f string, v ...interface{}) {

}

func Trace(tag string, v ...interface{}) {
	logger.writeLog(LogLevel_Trace, tag, v...)
}

func Error(tag string, v ...interface{}) {
	logger.writeLog(LogLevel_Error, tag, v...)
}

func Info(tag string, v ...interface{}) {
	logger.writeLog(LogLevel_Info, tag, v...)
}

func Fatal(tag string, v ...interface{}) {
	logger.writeLog(LogLevel_Fatal, tag, v...)
}

func SetChannel(channelType string, config string) error {
	return logger.setChannel(channelType, config)
}

func init() {
	err := SetChannel("console", "")
	if err != nil {
		panic(err)
	}
}
