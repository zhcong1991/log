package channel

import (
	"os"
	"github.com/beanwc/log"
)

type ConsoleChannel struct {
}

func (channel *ConsoleChannel) GetType() string {
	return log.LoggerChannel_Console
}

func (channel *ConsoleChannel) Init(config string) error {
	return nil
}

func (channel *ConsoleChannel) Write(level log.LogLevel, data []byte) error {
	if level == log.LogLevel_Error || level == log.LogLevel_Fatal {
		os.Stderr.Write(data)
	} else {
		os.Stdout.Write(data)
	}
	return nil
}

func init() {
	log.Register(log.LoggerChannel_Console, func() log.LoggerChannel {
		return new(ConsoleChannel)
	})
}
