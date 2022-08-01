package log

import (
	"context"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	Init(&Config{
		LogLevel: "debug",
		Writers: []WriterConfig{
			{
				Type:     "console",
				LogLevel: "debug",
			},
			{
				Type:      "file",
				LogLevel:  "info",
				FilePath:  "./log/service.info",
				SplitUnit: "hour",
			},
			{
				Type:      "file",
				LogLevel:  "warn",
				FilePath:  "./log/service.wf",
				SplitUnit: "hour",
			},
		},
	})
	ctx := context.WithValue(context.TODO(), "trace", "abcefg")
	for i := 0; i < 700; i++ {
		Info("Test", "Test Info")
		InfoCtx(ctx, "Test", "Test Info")
		Error("Test", "Test Error")
		ErrorCtx(ctx, "Test", "Test Error")
		time.Sleep(5 * time.Second)
		i++
	}
	_logger.Close()
}
