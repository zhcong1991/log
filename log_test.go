package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	Init(&Config{
		LogLevel: DEBUG,
		Writes: []WriterConfig{
			{
				Type:     "console",
				LogLevel: DEBUG,
			},
			{
				Type:      "file",
				LogLevel:  INFO,
				FilePath:  "./log/service.info",
				SplitUnit: "hour",
			},
			{
				Type:     "file",
				LogLevel: WARN,
				FilePath:  "./log/service.wf",
				SplitUnit: "day",
			},
		},
	})
	Debug("Test", "Test Debug")
	Warn("Test", "Test Trace")
	Error("Test", "Test Error")
	Info("Test", "Test Info")
	Fatal("Test", "Test Fatal")
	_logger.Close()
}
