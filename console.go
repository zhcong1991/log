package log

import (
	"os"
)

type ConsoleChannel struct {
}

func (channel *ConsoleChannel) GetType() string {
	return LoggerChannel_Console
}

func (channel *ConsoleChannel) Init(config string) error {
	return nil
}

func (channel *ConsoleChannel) Write(level LogLevel, data []byte) error {
	if level == LogLevel_Error || level == LogLevel_Fatal {
		os.Stderr.Write(data)
	} else {
		os.Stdout.Write(data)
	}
	return nil
}

func init() {
	Register(LoggerChannel_Console, func() LoggerChannel {
		return new(ConsoleChannel)
	})
}
