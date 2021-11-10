package log

import (
	"context"
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
	ctx := context.WithValue(context.TODO(), "trace", "abcefg")
	Info("Test", "Test Info")
	InfoCtx(ctx,"Test", "Test Info")
	Error("Test", "Test Error")
	ErrorCtx(ctx,"Test", "Test Error")
	_logger.Close()
}
