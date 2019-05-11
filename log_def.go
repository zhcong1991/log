package log

import (
	"time"
)

type LogLevel = int32

const (
	LogLevel_Debug LogLevel = iota
	LogLevel_Trace
	LogLevel_Error
	LogLevel_Info
	LogLevel_Fatal
	LogLevel_Count
)

var (
	logName      = [LogLevel_Count]string{"debug", "trace", "error", "info", "fatal"}
	logShortName = [LogLevel_Count]string{"[D]", "[T]", "[E]", "[I]", "[F]"}
)

const (
	LoggerChannel_Console = "console"
	LoggerChannel_File    = "file"
	LoggerChannel_Conn    = "conn"
)

type logRecord struct {
	level     LogLevel
	content   string
	timestamp time.Time
}

type LoggerChannel interface {
	GetType() string
	Init(config string) error
	Write(level LogLevel, data []byte) error
}

type newLoggerChannelFunc func() LoggerChannel
